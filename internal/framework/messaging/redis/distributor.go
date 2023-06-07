package redis

import (
	"github.com/abc-valera/flugo-api/internal/domain/service"
	"github.com/hibiken/asynq"
)

type redisTaskDistributor struct {
	log    service.Logger
	client *asynq.Client
}

func newRedisTaskDistributor(log service.Logger, redisOpt *asynq.RedisClientOpt) *redisTaskDistributor {
	return &redisTaskDistributor{
		log:    log,
		client: asynq.NewClient(redisOpt),
	}
}
