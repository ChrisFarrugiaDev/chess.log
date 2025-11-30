package mailer

import (
	"chess_log/go_backend/internal/logger"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"go.uber.org/zap"
)

func TwilioEmailSender(subject, emailTo string) error {

	htmlContent := "<h2>Hello World!</h2>"
	plainTextContent := "hello world"

	twilioSenderEmail := os.Getenv("TWILIO_SENDER_EMAIL")
	from := mail.NewEmail("ChessLog", twilioSenderEmail)
	to := mail.NewEmail("User", emailTo)

	twilioApiKey := os.Getenv("TWILIO_API_KEY")
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(twilioApiKey)

	// Send the email.
	response, err := client.Send(message)
	if err != nil {
		logger.Debug("Failed to send reset password email", zap.Error(err))
		return err
	}

	logger.Debug("Email sent successfully", zap.Int("status_code", response.StatusCode))
	return nil
}
