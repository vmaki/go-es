package asynq

import (
	"github.com/hibiken/asynq"
	"go-es/common/uniqueid"
	"sync"
	"time"
)

var (
	once   sync.Once
	Client *asynq.Client
	Srv    *asynq.Server
)

func ConnectAsynq(address string, username string, password string, db int) {
	r := asynq.RedisClientOpt{Addr: address, Username: username, Password: password, DB: db}

	once.Do(func() {
		Client = asynq.NewClient(r)

		Srv = asynq.NewServer(
			r,
			asynq.Config{
				Concurrency: 10,
				Queues: map[string]int{
					"critical": 6,
					"default":  3,
					"low":      1,
				},
			},
		)
	})
}

func EnqueueIn(task *asynq.Task, workTime int64) (err error) {
	taskID := uniqueid.GenSn(uniqueid.SnPrefixAsynq)

	if workTime == 0 {
		_, err = Client.Enqueue(task, asynq.MaxRetry(3), asynq.TaskID(taskID))
		return err
	}

	_, err = Client.Enqueue(
		task,
		asynq.MaxRetry(3),
		asynq.TaskID(taskID),
		asynq.ProcessIn(time.Duration(workTime)*time.Second),
	)

	return err
}
