package mail

import (
	"context"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/skyrocket-qy/erx"
)

type MailService interface {
	SendMail(ctx context.Context, to, sbj, content string) (err error)
}

var _ MailService = &MailServiceImpl{}

type MailServiceImpl struct {
	Token   string
	SrcAddr string
}

func NewMailServiceImpl(sendGridToken, srcAddr string) MailService {
	return &MailServiceImpl{
		Token:   sendGridToken,
		SrcAddr: srcAddr,
	}
}

func (m *MailServiceImpl) SendMail(ctx context.Context, to, sbj, content string) (
	err error,
) {
	from := mail.NewEmail("no-reply", m.SrcAddr)
	subject := sbj
	toMail := mail.NewEmail("to", to)
	plainTextContent := ""
	htmlContent := content
	message := mail.NewSingleEmail(from, subject, toMail, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(m.Token)

	_, err = client.Send(message)
	if err != nil {
		return erx.W(err, "send mail failed reason")
	}

	return nil
}
