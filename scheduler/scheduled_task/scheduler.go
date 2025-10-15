package scheduled_task

import (
	"nexsoft.co.id/example/config"
	"nexsoft.co.id/example/server_config"
)

func ScheduledTask(
	config config.Configuration,
	serverAttribute server_config.ServerAttribute,
) {
	serverAttribute.Scheduler.RegisterScheduler( 
		&removeUnactiveWardesProfileTask{},
	)

	serverAttribute.Scheduler.StartSchedulerCron()
}