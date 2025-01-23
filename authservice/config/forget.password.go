package config

import "fmt"

// ConfigForgetBody constructs the reset password email body with a default password
func ConfigFogetBody(token string, id string) (string, string) {

	subject := "Reset Password"
	body := fmt.Sprintf(
		`
	<h1>Reset Password</h1>
	<hr>
	<p>Your password has been reset to the following temporary password:</p>
	<p><strong>%s</strong></p>
	<p>For security, please update your password after logging in.</p>
	<hr>
	
	`, token)

	return subject, body
}
