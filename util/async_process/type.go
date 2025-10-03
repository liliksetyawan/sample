package async_process

import (
    "time"
    "github.com/nexsoft-git/nexcommon/limited_scheduler"
)

type AsyncProcessUtil interface {
    GetQueueGrologQueue() limited_scheduler.QueueLimitedScheduler
    GetQueueNexlogQueue() limited_scheduler.QueueLimitedScheduler
    Close()
}

type AsyncProcessConfig struct {
    NumberOfListener       int
    NumberOfPopMessage     int
    IdleConnectionTime     time.Duration
}
