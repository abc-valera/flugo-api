package service

type EmailSender interface {
	SendEmail(subject, content string, to []string, attchFiles []string) error
}
