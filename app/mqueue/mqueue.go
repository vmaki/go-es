package mqueue

import (
	"context"
	"github.com/hibiken/asynq"
)

type MQueue struct {
	ctx context.Context
}

func NewMQueue(ctx context.Context) *MQueue {
	return &MQueue{
		ctx: ctx,
	}
}

func (l *MQueue) Register() *asynq.ServeMux {
	mux := asynq.NewServeMux()

	RegisterApiTasks(mux)

	return mux
}
