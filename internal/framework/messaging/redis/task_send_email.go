package redis

import (
	"context"
	"encoding/json"

	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/hibiken/asynq"
)

const taskSendVerifyEmail = "task:send_verify_email"

type payloadSendVerifyEmail struct {
	Email string `json:"email"`
}

func (p *redisTaskDistributor) DistributeTaskSendVerifyEmail(
	c context.Context,
	payload *payloadSendVerifyEmail,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return domain.NewInternalError("DistributeTaskSendVerifyEmail", err)
	}

	task := asynq.NewTask(taskSendVerifyEmail, jsonPayload, opts...)
	info, err := p.client.EnqueueContext(c, task)
	if err != nil {
		return domain.NewInternalError("DistributeTaskSendVerifyEmail", err)
	}

	p.log.Info("ENQUEUED TASK",
		"type", task.Type(),
		"queue", info.Queue,
		"max_retry", info.MaxRetry,
	)

	return nil
}

func (p *redisTaskProcessor) ProccessTaskSendVerifyEmail(c context.Context, task *asynq.Task) error {
	payload := new(payloadSendVerifyEmail)
	if err := json.Unmarshal(task.Payload(), payload); err != nil {
		return asynq.SkipRetry
	}

	// send email..

	p.log.Info("PROCESSED TASK",
		"type", task.Type(),
		"email", payload.Email,
	)

	return nil
}
