package service

import "context"

type MessagingBroker interface {
	SendVerifyEmailTask(c context.Context, to string) error
}
