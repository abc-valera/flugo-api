package framework

import "time"

type Frameworks struct {
	PasswordFramework PasswordFramework
	TokenFramework    TokenFramework
}

func NewFrameworks(accessDuration, refreshDuration time.Duration) *Frameworks {
	return &Frameworks{
		PasswordFramework: newPasswordFramework(),
		TokenFramework:    newTokenFramework(accessDuration, refreshDuration),
	}
}
