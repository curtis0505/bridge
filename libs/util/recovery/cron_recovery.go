package recovery

import (
	"fmt"
	"github.com/curtis0505/bridge/libs/logger/v2"
	"github.com/curtis0505/bridge/libs/util"
	cron "github.com/robfig/cron/v3"
)

func CronRecovery(mod string) cron.JobWrapper {
	return func(j cron.Job) cron.Job {
		return cron.FuncJob(func() {
			defer func() {
				if err := recover(); err != nil {
					trace := fmt.Sprintf("panic: %+v \n\n%s", err, string(Stack(3)))

					logger.Error("GinRecovery", logger.BuildLogInput().WithData("panic", trace))

					message := util.NewMessage().
						SetZone(mod).
						SetMessageType(util.MessageTypePanic).
						SetTitle(util.TitlePanic).
						AddTextParagraphWidget(trace)

					message.SendMessage()
				}
			}()
			j.Run()
		})
	}
}
