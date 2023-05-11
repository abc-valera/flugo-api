package pkg

import (
	"fmt"
	"net/smtp"

	"github.com/abc-valera/flugo-api/internal/domain"
	"github.com/jordan-wright/email"
)

const (
	senderName        = "Flugo"
	smtpAuth          = "smtp.gmail.com"
	smtpServerAddress = "smtp.gmail.com:587"
)

type goMail struct {
	senderName     string
	senderAddress  string
	senderPassword string
}

func newEmailPackage(address, password string) domain.EmailPackage {
	return &goMail{
		senderName:     senderName,
		senderAddress:  address,
		senderPassword: password,
	}
}

func (f *goMail) SendEmail(subject, content string, to []string, attachFiles []string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", f.senderName, f.senderAddress)
	e.Subject = subject
	e.HTML = []byte(content)
	e.To = to
	for _, f := range attachFiles {
		_, err := e.AttachFile(f)
		if err != nil {
			return domain.NewInternalError("Send email: failed to attach file: "+f+" Err: ", err)
		}
	}

	auth := smtp.PlainAuth("", f.senderAddress, f.senderPassword, smtpAuth)
	return e.Send(smtpServerAddress, auth)
}
