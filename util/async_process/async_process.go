package async_process

import (
    "database/sql"
    "github.com/nexsoft-git/nexcommon/dao"
    "github.com/nexsoft-git/nexcommon/limited_scheduler"
)

func NewAsyncProcessUtil(
    sql *sql.DB,
    dao dao.DBQueueDAO,
    queueGrolog AsyncProcessConfig,
    queueNexlog AsyncProcessConfig,
) AsyncProcessUtil {
    result := &asyncProcessUtil{}

    result.queueGrologQueue = limited_scheduler.NewLimitedSchedulerFromConfig(
        limited_scheduler.QueueConfig{
            NumberOfListener:   queueGrolog.NumberOfListener,
            QueueType:          config.LimitedAsyncProcess.ExampleConfig.QueueType,
            IdleConnectionTime: queueGrolog.IdleConnectionTime,
        }, sql, dao, queueGrolog.NumberOfPopMessage,
    )

    result.queueNexlogQueue = limited_scheduler.NewLimitedSchedulerFromConfig(
        limited_scheduler.QueueConfig{
            NumberOfListener:   queueNexlog.NumberOfListener,
            QueueType:          config.LimitedAsyncProcess.ExampleConfig.QueueType,
            IdleConnectionTime: queueNexlog.IdleConnectionTime,
        }, sql, dao, queueNexlog.NumberOfPopMessage,
    )

    return result
}

type asyncProcessUtil struct {
    queueGrologQueue limited_scheduler.QueueLimitedScheduler
    queueNexlogQueue limited_scheduler.QueueLimitedScheduler
}

func (a *asyncProcessUtil) GetQueueGrologQueue() limited_scheduler.QueueLimitedScheduler {
    return a.queueGrologQueue
}

func (a *asyncProcessUtil) GetQueueNexlogQueue() limited_scheduler.QueueLimitedScheduler {
    return a.queueNexlogQueue
}

func (a *asyncProcessUtil) Close() {
    a.queueGrologQueue.Listener.StopListen()
    a.queueNexlogQueue.Listener.StopListen()
}