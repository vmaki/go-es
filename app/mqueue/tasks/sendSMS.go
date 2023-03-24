package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	asynq2 "go-es/internal/pkg/asynq"
)

var SendSMSType = "send:sms"

type SendSMSPayload struct {
	Phone    string
	Code     int
	WorkTime int64
}

func SendSMSTask(req SendSMSPayload) error {
	payload, _ := json.Marshal(req)
	task := asynq.NewTask(SendSMSType, payload)

	err := asynq2.EnqueueIn(task, req.WorkTime)
	if err != nil {
		return err
	}

	return nil
}

func SendSMSTaskHandle(ctx context.Context, t *asynq.Task) error {
	var p SendSMSPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("发短信失败, err: : %v: %w", err, asynq.SkipRetry)
	}

	fmt.Printf("开始发送短信. 手机号码: %s, 短信内容: %d\n", p.Phone, p.Code)

	return nil
}
