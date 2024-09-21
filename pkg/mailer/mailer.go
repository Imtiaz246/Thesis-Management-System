package mailer

import (
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/spf13/viper"
	"net/smtp"
	"sync"
)

type Mailer struct {
	sourceName string
	mailerServiceInfo
	sync.RWMutex
}

type mailerServiceInfo struct {
	serviceEmail    string
	servicePass     string
	serviceSMTPAuth string
	serviceSMTPPort string
}

func (m *Mailer) getSmtpAuth() smtp.Auth {
	m.RLock()
	defer m.RUnlock()
	return smtp.PlainAuth("", m.serviceEmail, m.servicePass, m.serviceSMTPAuth)
}

func (m *Mailer) getSourceInfo() string {
	m.RLock()
	defer m.RUnlock()
	return fmt.Sprintf("%s <%s>", m.sourceName, m.serviceEmail)
}

func (m *Mailer) getSmtpServer() string {
	m.RLock()
	defer m.RUnlock()
	return fmt.Sprintf("%s:%s", m.serviceSMTPAuth, m.serviceSMTPPort)
}

type mailingOption struct {
	to           []string
	cc           []string
	bcc          []string
	subject      string
	htmlTemplate []byte
	attachments  []string
}

type OptsHandler func(opts *mailingOption)

func NewMailer(conf *viper.Viper) *Mailer {
	return &Mailer{
		sourceName: conf.GetString("mailer.source_name"),
		mailerServiceInfo: mailerServiceInfo{
			serviceEmail:    conf.GetString("mailer.service_email"),
			servicePass:     conf.GetString("mailer.service_pass"),
			serviceSMTPAuth: conf.GetString("mailer.smtp_auth"),
			serviceSMTPPort: conf.GetString("mailer.smtp_port"),
		},
		RWMutex: sync.RWMutex{},
	}
}

func (m *Mailer) SendMail(optsHandlers ...OptsHandler) error {
	opts := &mailingOption{}
	for _, optsHandler := range optsHandlers {
		optsHandler(opts)
	}
	e := email.Email{
		From:    m.getSourceInfo(),
		Subject: opts.subject,
		HTML:    opts.htmlTemplate,
		To:      opts.to,
		Cc:      opts.cc,
		Bcc:     opts.bcc,
	}

	for _, f := range opts.attachments {
		_, err := e.AttachFile(f)
		if err != nil {
			return err
		}
	}

	return e.Send(m.getSmtpServer(), m.getSmtpAuth())
}

func WithReceivers(receivers []string) OptsHandler {
	return func(opts *mailingOption) {
		opts.to = receivers
	}
}

func WithReceiver(receiver string) OptsHandler {
	return func(opts *mailingOption) {
		opts.to = append(opts.to, receiver)
	}
}

func WithCC(cc []string) OptsHandler {
	return func(opts *mailingOption) {
		opts.cc = cc
	}
}

func WithBCC(bcc []string) OptsHandler {
	return func(opts *mailingOption) {
		opts.bcc = bcc
	}
}

func WithHTMLTemplate(htmlTemplate []byte) OptsHandler {
	return func(opts *mailingOption) {
		opts.htmlTemplate = htmlTemplate
	}
}

func WithAttachments(attachments []string) OptsHandler {
	return func(opts *mailingOption) {
		opts.attachments = attachments
	}
}

func WithSubject(subject string) OptsHandler {
	return func(opts *mailingOption) {
		opts.subject = subject
	}
}
