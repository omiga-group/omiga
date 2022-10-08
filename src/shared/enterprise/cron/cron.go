package cron

import (
	"strings"
	"time"

	timeex "github.com/omiga-group/omiga/src/shared/enterprise/time"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type CronService interface {
	GetCron() *cron.Cron
	Close()
}

type cronService struct {
	logger     *zap.SugaredLogger
	timeHelper timeex.TimeHelper
	cron       *cron.Cron
}

func NewCronService(
	logger *zap.SugaredLogger,
	timeHelper timeex.TimeHelper) (CronService, error) {
	instance := &cronService{
		logger:     logger,
		timeHelper: timeHelper,
	}

	instance.cron = cron.New(
		cron.WithSeconds(),
		cron.WithLogger(instance),
		cron.WithChain(
			cron.DelayIfStillRunning(instance),
		))

	instance.cron.Start()

	return instance, nil
}

func (cs *cronService) GetCron() *cron.Cron {
	return cs.cron
}

func (cs *cronService) Close() {
	if cs.cron != nil {
		cronCtx := cs.cron.Stop()

		// Wait until all running jobs exit gracefully
		cs.timeHelper.WaitUntilCancelled(cronCtx)

		cs.cron = nil
	}
}

func (cs *cronService) Info(msg string, keysAndValues ...interface{}) {
}

func (cs *cronService) Error(err error, msg string, keysAndValues ...interface{}) {
	keysAndValues = formatTimes(keysAndValues)

	cs.logger.Errorf(
		formatString(len(keysAndValues)+2),
		append([]interface{}{msg, "error", err}, keysAndValues...)...)
}

func formatString(numKeysAndValues int) string {
	var sb strings.Builder

	sb.WriteString("%s")

	if numKeysAndValues > 0 {
		sb.WriteString(", ")
	}

	for i := 0; i < numKeysAndValues/2; i++ {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString("%v=%v")
	}

	return sb.String()
}

func formatTimes(keysAndValues []interface{}) []interface{} {
	var formattedArgs []interface{}

	for _, arg := range keysAndValues {
		if t, ok := arg.(time.Time); ok {
			arg = t.Format(time.RFC3339)
		}
		formattedArgs = append(formattedArgs, arg)
	}

	return formattedArgs
}
