package messaging

import "context"

type MessagingBroker interface {
	StartTaskProcessor() error
	SendVerifyEmailTask(c context.Context, to string) error
}
