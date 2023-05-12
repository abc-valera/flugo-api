package service

// Note: rename to EmailService
type EmailService interface {
	SendEmail(subject, content string, to []string, attchFiles []string) error
}
