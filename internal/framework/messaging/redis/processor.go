package redis

import (
	"github.com/abc-valera/flugo-api/internal/domain/repository"
	"github.com/abc-valera/flugo-api/internal/domain/service"
	"github.com/hibiken/asynq"
)

const (
	queueCritical = "critical"
	queueDefault  = "default"
)

type redisTaskProcessor struct {
	log      service.Logger
	server   *asynq.Server
	userRepo repository.UserRepository
}

func newRedisTaskProcessor(log service.Logger, redisOpt *asynq.RedisClientOpt, userRepo repository.UserRepository) *redisTaskProcessor {
	return &redisTaskProcessor{
		log: log,
		server: asynq.NewServer(
			redisOpt,
			asynq.Config{
				Queues: map[string]int{
					queueCritical: 10,
					queueDefault:  5,
				},
			},
		),
		userRepo: userRepo,
	}
}

func (p *redisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()

	// TODO write error handler

	mux.HandleFunc(taskSendVerifyEmail, p.ProccessTaskSendVerifyEmail)

	return p.server.Start(mux)
}
