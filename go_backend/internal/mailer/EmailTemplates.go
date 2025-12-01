package mailer

import "fmt"

func GenerateEmailValidationTemp(link string) (string, string) {

	plainBody := fmt.Sprintf(
		"Welcome to ChessLog!\n\n"+
			"Please verify your email address by clicking the link below:\n\n"+
			"%s\n\n"+
			"This link is valid for 15 minutes.\n\n"+
			"Thank you for joining ChessLog!",
		link,
	)

	htmlBody := fmt.Sprintf(`
	<h1>Welcome to <b>ChessLog</b>!</h1>
	<p>Thank you for creating an account.</p>
	<p>Please verify your email address by clicking the button below:</p>
	
	<p>
		<a href="%s" 
		   style="display:inline-block;padding:12px 20px;background:#4B7BF5;
				  color:white;border-radius:6px;text-decoration:none;font-size:16px;">
			Verify Email
		</a>
	</p>

	<p>If the button doesn’t work, click this link:</p>
	<p><a href="%s">%s</a></p>

	<p>This link will expire in <b>15 minutes</b>.</p>
`, link, link, link)

	return plainBody, htmlBody
}
