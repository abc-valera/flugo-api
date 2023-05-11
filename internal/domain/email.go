package domain

// Note: rename to EmailPackage
type EmailPackage interface {
	SendEmail(subject, content string, to []string, attchFiles []string) error
}
