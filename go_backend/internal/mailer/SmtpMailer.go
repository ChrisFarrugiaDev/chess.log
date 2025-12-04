package mailer

import (
	"github.com/wneessen/go-mail"
)

type SmtpMailer struct {
	client *mail.Client
	from   string
}

// Create a new mailer instance (used once during app startup)
func NewSmtpMailer(host string, port int, username, password, from string) (*SmtpMailer, error) {
	opts := []mail.Option{
		mail.WithPort(port),
		mail.WithUsername(username),
		mail.WithPassword(password),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
	}

	// If using SSL port 465 → enable implicit TLS
	if port == 465 {
		opts = append(opts,
			mail.WithSSL(), // Enable implicit SSL
			mail.WithTLSPortPolicy(mail.TLSMandatory), // Require TLS
		)
	} else {
		// Port 587 → STARTTLS
		opts = append(opts,
			mail.WithTLSPortPolicy(mail.TLSOpportunistic), // Upgrade to TLS
		)
	}

	client, err := mail.NewClient(host, opts...)
	if err != nil {
		return nil, err
	}

	return &SmtpMailer{
		client: client,
		from:   from,
	}, nil
}

// Create a message container (HTML + text supported)
func (m *SmtpMailer) CreateMessage(to, subject, plainBody, htmlBody string) (*mail.Msg, error) {
	message := mail.NewMsg()

	if err := message.From(m.from); err != nil {
		return nil, err
	}

	if err := message.To(to); err != nil {
		return nil, err
	}

	message.Subject(subject)

	message.SetBodyString(mail.TypeTextPlain, plainBody)

	if htmlBody != "" {
		message.AddAlternativeString(mail.TypeTextHTML, htmlBody)
	}

	return message, nil
}

// Send the message
func (m *SmtpMailer) Send(msg *mail.Msg) error {
	return m.client.DialAndSend(msg)
}
