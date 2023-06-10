package redis

import (
	"context"

	"github.com/abc-valera/flugo-api/internal/domain"
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
				ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
					code := domain.ErrorCode(err)
					if code == "" {
						code = domain.CodeInternal
					}
					log.Error("PROCESS TASK FAILED",
						"code", code,
						"error", err,
						"task", task.Type(),
						"payload:", string(task.Payload()))
				}),
				Logger: &customAsynqLogger{log},
			},
		),
		userRepo: userRepo,
	}
}

func (p *redisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()

	mux.HandleFunc(taskSendVerifyEmail, p.ProccessTaskSendVerifyEmail)

	return p.server.Start(mux)
}

type customAsynqLogger struct {
	service.Logger
}

func (l *customAsynqLogger) Debug(args ...interface{}) {
	l.Logger.Debug("ASYNQ", "debug", args)
}

func (l *customAsynqLogger) Info(args ...interface{}) {
	l.Logger.Info("ASYNQ", "info", args)
}

func (l *customAsynqLogger) Warn(args ...interface{}) {
	l.Logger.Warn("ASYNQ", "warn", args)
}

func (l *customAsynqLogger) Error(args ...interface{}) {
	l.Logger.Error("ASYNQ", "err", args)
}

func (l *customAsynqLogger) Fatal(args ...interface{}) {
	l.Logger.Fatal("ASYNQ", "fatal", args)
}
