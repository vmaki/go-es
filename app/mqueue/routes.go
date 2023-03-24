package mqueue

import (
	"github.com/hibiken/asynq"
	"go-es/app/mqueue/tasks"
)

func RegisterApiTasks(mux *asynq.ServeMux) {
	mux.HandleFunc(tasks.SendSMSType, tasks.SendSMSTaskHandle)
}
