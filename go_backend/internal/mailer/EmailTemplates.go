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

func GeneratePasswordResetEmail(link string) (string, string) {

	plainBody := fmt.Sprintf(
		"You requested to reset your ChessLog password.\n\n"+
			"Click the link below to choose a new password:\n\n"+
			"%s\n\n"+
			"This link is valid for 10 minutes.\n\n"+
			"If you did not request a password reset, please ignore this email.",
		link,
	)

	htmlBody := fmt.Sprintf(`
	<h1>Password Reset Request</h1>

	<p>You recently requested to reset your <b>ChessLog</b> password.</p>

	<p>Click the button below to create a new password:</p>

	<p>
		<a href="%s"
		   style="display:inline-block;
				  padding:12px 20px;
				  background:#E11D48;
				  color:white;
				  border-radius:6px;
				  text-decoration:none;
				  font-size:16px;">
			Reset Password
		</a>
	</p>

	<p>If the button doesn’t work, use this link:</p>
	<p><a href="%s">%s</a></p>

	<p>This link will expire in <b>10 minutes</b>.</p>

	<br/>
	<p>If you did not request this, you can safely ignore this email.</p>
`, link, link, link)

	return plainBody, htmlBody
}
