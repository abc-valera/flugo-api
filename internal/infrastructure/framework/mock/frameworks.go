package mock

import (
	"github.com/abc-valera/flugo-api/internal/infrastructure/framework"
)

// Frameworks contains mock framework structs
type Frameworks struct {
	PasswordFramework *PasswordFramework
	TokenFramework    *TokenFramework
	EmailFramework    *EmailFramework
}

// Returns two structs:
//   - frameworks with interfaces which is used to init services
//   - frameworks with mock structs which is used to reference mocks
func NewFrameworks() (*framework.Frameworks, *Frameworks) {
	pf := new(PasswordFramework)
	tf := new(TokenFramework)
	ef := new(EmailFramework)
	return &framework.Frameworks{
			PasswordFramework: pf,
			TokenFramework:    tf,
			EmailFramework:    ef,
		}, &Frameworks{
			PasswordFramework: pf,
			TokenFramework:    tf,
			EmailFramework:    ef,
		}
}
