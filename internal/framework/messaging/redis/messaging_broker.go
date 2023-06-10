package redis

import (
	"context"
	"time"

	"github.com/abc-valera/flugo-api/internal/domain/repository"
	"github.com/abc-valera/flugo-api/internal/domain/service"
	"github.com/hibiken/asynq"
)

type messagingBroker struct {
	distributor *redisTaskDistributor
	processor   *redisTaskProcessor
}

func NewMessagingBroker(
	redisUrl,
	redisUser,
	redisPass string,
	userRepo repository.UserRepository,
	log service.Logger,
) service.MessagingBroker {
	redisOpt := &asynq.RedisClientOpt{
		Addr:     redisUrl,
		Username: redisUser,
		Password: redisPass,
	}

	proc := newRedisTaskProcessor(log, redisOpt, userRepo)
	proc.log.Info("Starting task processor")
	go proc.Start()

	return &messagingBroker{
		distributor: newRedisTaskDistributor(log, redisOpt),
		processor:   proc,
	}
}

func (m *messagingBroker) SendVerifyEmailTask(c context.Context, to string) error {
	taskPayload := &payloadSendVerifyEmail{
		Email: to,
	}

	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(5 * time.Second),
		asynq.Queue(queueCritical),
	}

	if err := m.distributor.DistributeTaskSendVerifyEmail(c, taskPayload, opts...); err != nil {
		return err
	}

	return nil
}
