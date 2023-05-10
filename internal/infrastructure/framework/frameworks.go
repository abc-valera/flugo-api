package framework

import "time"

type Frameworks struct {
	PasswordFramework PasswordFramework
	TokenFramework    TokenFramework
	EmailFramework    EmailFramework
}

func NewFrameworks(
	accessDuration, refreshDuration time.Duration,
	senderAddress, senderPassword string,
) *Frameworks {
	return &Frameworks{
		PasswordFramework: newPasswordFramework(),
		TokenFramework:    newTokenFramework(accessDuration, refreshDuration),
		EmailFramework:    newEmailFramework(senderAddress, senderPassword),
	}
}
