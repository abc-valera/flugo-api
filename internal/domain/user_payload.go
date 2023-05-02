package domain

import (
	"time"
)

var ErrExpiredToken = NewErrWithMsg(CodeUnauthenticated, "The token has expired")

type Payload struct {
	Username  string    `json:"username"`
	IsRefresh bool      `json:"is_refresh"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(username string, isRefresh bool, duration time.Duration) *Payload {
	return &Payload{
		Username:  username,
		IsRefresh: isRefresh,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
