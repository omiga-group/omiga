package subscribers

import (
	"context"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/omiga-group/omiga/src/shared/enterprise/cron"
	"github.com/omiga-group/omiga/src/shared/enterprise/security/authentication/passwordgeneration/totp"
	"github.com/omiga-group/omiga/src/venue/rain-processor/configuration"
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

	// Run at every minute from 0 through 59.
	// if _, err := cronService.GetCron().AddJob("* 0/1 * * * *", instance); err != nil {
	// 	return nil, err
	// }

	go instance.Run()

	return instance, nil
}

func (rtps *rainTradingPairSubscriber) Run() {
	playwrightInstance, err := playwright.Run()
	if err != nil {
		rtps.logger.Errorf("Failed to start Playwright instance. Error: %v", err)

		return
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
		rtps.logger.Errorf("Failed to start Chromium browser instance. Error: %v", err)

		return
	}

	defer func() {
		if err := browserInstance.Close(); err != nil {
			rtps.logger.Errorf("Failed to close Chromium browser instance. Error: %v", err)
		}
	}()

	signinPageInstance, err := browserInstance.NewPage()
	if err != nil {
		rtps.logger.Errorf("Failed to start signin page instance. Error: %v", err)

		return
	}

	defer func() {
		if err := signinPageInstance.Close(); err != nil {
			rtps.logger.Errorf("Failed to close signin page instance. Error: %v", err)
		}
	}()

	headersChannel := make(chan map[string]string)

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
		rtps.logger.Errorf("Failed to navigate to /signin page. Error: %v", err)

		return
	}

	if err = signinPageInstance.Fill("input[name=\"email\"]", rtps.rainConfig.Username); err != nil {
		rtps.logger.Errorf("Failed to fill email field. Error: %v", err)

		return
	}

	if err = signinPageInstance.Fill("input[name=\"password\"]", rtps.rainConfig.Password); err != nil {
		rtps.logger.Errorf("Failed to fill password field. Error: %v", err)

		return
	}

	if err = signinPageInstance.Click("text=Sign In"); err != nil {
		rtps.logger.Errorf("Failed to click on Sign In button. Error: %v", err)

		return
	}

	totpCode, err := rtps.totpHelper.GenerateCodeUsingCurrentTime(rtps.rainConfig.TotpSecret)
	if err != nil {
		rtps.logger.Errorf("Failed to generate TOTP code. Error: %v", err)

		return
	}

	if err = signinPageInstance.Fill("input[id=\"code\"]", totpCode); err != nil {
		rtps.logger.Errorf("Failed to fill code field. Error: %v", err)

		return
	}

	if err = signinPageInstance.Click("text=Submit"); err != nil {
		rtps.logger.Errorf("Failed to click on Submit button. Error: %v", err)

		return
	}

	select {
	case headers := <-headersChannel:
		client, err := rainv1.NewClientWithResponses(rtps.rainConfig.BaseUrl)
		if err != nil {
			rtps.logger.Errorf("Failed to create client with response. Error: %v", err)

			return
		}

		response, err := client.GetAllCoinsWithResponse(rtps.ctx, func(ctx context.Context, req *http.Request) error {
			for key, value := range headers {
				req.Header.Set(key, value)
			}

			return nil
		})
		if err != nil {
			rtps.logger.Errorf("Failed to call getAllCoins endpoint. Error: %v", err)

			return
		}

		if response.HTTPResponse.StatusCode != 200 {
			rtps.logger.Errorf("Failed to call getAllCoins endpoint. Return status code is %d", response.HTTPResponse.StatusCode)

			return
		}

		if response.JSON200 == nil {
			rtps.logger.Errorf("Returned JSON object is nil")

			return
		}

	case <-time.After(rtps.timeout):
		rtps.logger.Errorf("Timed out, failed to receive headers required to call getAllCoins endpoint.")
	}
}
