package subscribers

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/omiga-group/omiga/src/shared/enterprise/security/authentication/passwordgeneration/totp"
	"github.com/omiga-group/omiga/src/venue/rain-processor/configuration"
	"github.com/omiga-group/omiga/src/venue/rain-processor/mappers"
	rainv1 "github.com/omiga-group/omiga/src/venue/rain-processor/rainclient/v1"
	"github.com/omiga-group/omiga/src/venue/shared/repositories"
	"github.com/playwright-community/playwright-go"
	"go.uber.org/zap"
)

type RainTradingPairSubscriber interface {
}

type rainTradingPairSubscriber struct {
	ctx                   context.Context
	logger                *zap.SugaredLogger
	rainConfig            configuration.RainConfig
	tradingPairRepository repositories.TradingPairRepository
	totpHelper            totp.TotpHelper
	baseUrl               *url.URL
	websiteUrl            *url.URL
	timeout               time.Duration
	cachedHeaders         map[string]string
}

func NewRainTradingPairSubscriber(
	ctx context.Context,
	logger *zap.SugaredLogger,
	rainConfig configuration.RainConfig,
	cronService cron.CronService,
	tradingPairRepository repositories.TradingPairRepository,
	totpHelper totp.TotpHelper) (RainTradingPairSubscriber, error) {

	baseUrl, err := url.Parse(rainConfig.BaseUrl)
	if err != nil {
		return nil, err
	}

	websiteUrl, err := url.Parse(rainConfig.WebsiteUrl)
	if err != nil {
		return nil, err
	}

	timeout, err := time.ParseDuration(rainConfig.Timeout)
	if err != nil {
		return nil, err
	}

	instance := &rainTradingPairSubscriber{
		ctx:                   ctx,
		logger:                logger,
		rainConfig:            rainConfig,
		tradingPairRepository: tradingPairRepository,
		totpHelper:            totpHelper,
		baseUrl:               baseUrl,
		websiteUrl:            websiteUrl,
		timeout:               timeout,
	}

	// Run at every 5th minute from 0 through 59..
	if _, err := cronService.GetCron().AddJob("* 0/5 * * * *", instance); err != nil {
		return nil, err
	}

	return instance, nil
}

func (rtps *rainTradingPairSubscriber) Run() {
	if rtps.cachedHeaders != nil {
		if err := rtps.saveTradingPairs(rtps.cachedHeaders); err == nil {
			return
		} else {
			rtps.cachedHeaders = nil
		}
	}

	if headers, err := rtps.getRequiredHeaders(); err != nil {
		rtps.cachedHeaders = nil
		rtps.logger.Error(err)

		return
	} else {
		if err := rtps.saveTradingPairs(headers); err != nil {
			rtps.cachedHeaders = nil
			rtps.logger.Error(err)

			return
		}

		rtps.cachedHeaders = headers
	}
}

func (rtps *rainTradingPairSubscriber) getRequiredHeaders() (map[string]string, error) {
	headersChannel := make(chan map[string]string)
	defer close(headersChannel)

	playwrightInstance, err := playwright.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to start Playwright instance. Error: %v", err)
	}

	defer func() {
		if err := playwrightInstance.Stop(); err != nil {
			rtps.logger.Errorf("Failed to stop Playwright instance. Error: %v", err)
		}
	}()

	browserInstance, err := playwrightInstance.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: &rtps.rainConfig.Headless,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to start Chromium browser instance. Error: %v", err)
	}

	defer func() {
		if err := browserInstance.Close(); err != nil {
			rtps.logger.Errorf("Failed to close Chromium browser instance. Error: %v", err)
		}
	}()

	signinPageInstance, err := browserInstance.NewPage()
	if err != nil {
		return nil, fmt.Errorf("failed to start signin page instance. Error: %v", err)
	}

	defer func() {
		if err := signinPageInstance.Close(); err != nil {
			rtps.logger.Errorf("Failed to close signin page instance. Error: %v", err)
		}
	}()

	signinPageInstance.On("request", func(request playwright.Request) {
		configurationUrl := *rtps.baseUrl
		configurationUrl.Path = path.Join(rtps.baseUrl.Path, "configuration")

		if strings.EqualFold(request.URL(), configurationUrl.String()) {
			go func(headersChannel chan map[string]string) {
				headers := request.Headers()

				headersChannel <- headers
			}(headersChannel)
		}
	})

	signinUrl := *rtps.websiteUrl
	signinUrl.Path = path.Join(rtps.websiteUrl.Path, "signin")

	if _, err = signinPageInstance.Goto(signinUrl.String()); err != nil {
		return nil, fmt.Errorf("failed to navigate to /signin page. Error: %v", err)
	}

	if err = signinPageInstance.Fill("input[name=\"email\"]", rtps.rainConfig.Username); err != nil {
		return nil, fmt.Errorf("failed to fill email field. Error: %v", err)
	}

	if err = signinPageInstance.Fill("input[name=\"password\"]", rtps.rainConfig.Password); err != nil {
		return nil, fmt.Errorf("failed to fill password field. Error: %v", err)
	}

	if err = signinPageInstance.Click("text=Sign In"); err != nil {
		return nil, fmt.Errorf("failed to click on Sign In button. Error: %v", err)
	}

	totpCode, err := rtps.totpHelper.GenerateCodeUsingCurrentTime(rtps.rainConfig.TotpSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to generate TOTP code. Error: %v", err)
	}

	if err = signinPageInstance.Fill("input[id=\"code\"]", totpCode); err != nil {
		return nil, fmt.Errorf("failed to fill code field. Error: %v", err)
	}

	if err = signinPageInstance.Click("text=Submit"); err != nil {
		return nil, fmt.Errorf("failed to click on Submit button. Error: %v", err)
	}

	select {
	case headers := <-headersChannel:
		return headers, nil
	case <-time.After(rtps.timeout):
		return nil, fmt.Errorf("timed out, failed to receive headers required to call getAllCoins endpoint")
	}
}

func (rtps *rainTradingPairSubscriber) saveTradingPairs(headers map[string]string) error {
	client, err := rainv1.NewClientWithResponses(rtps.rainConfig.BaseUrl)
	if err != nil {
		return fmt.Errorf("failed to create client with response. Error: %v", err)
	}

	response, err := client.GetAllCoinsWithResponse(rtps.ctx, func(ctx context.Context, req *http.Request) error {
		for key, value := range headers {
			req.Header.Set(key, value)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to call getAllCoins endpoint. Error: %v", err)
	}

	if response.HTTPResponse.StatusCode != 200 {
		return fmt.Errorf("failed to call getAllCoins endpoint. Return status code is %d", response.HTTPResponse.StatusCode)
	}

	if response.JSON200 == nil {
		return fmt.Errorf("returned JSON object is nil")
	}

	if err = rtps.tradingPairRepository.CreateTradingPairs(
		rtps.ctx,
		rtps.rainConfig.Id,
		mappers.RainCoinsToTradingPairs(response.JSON200.Coins)); err != nil {
		return fmt.Errorf("failed to create trading pairs. Error: %v", err)
	}

	return nil
}