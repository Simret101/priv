package config

import "fmt"

// ConfigBody generates the subject and body for an email verification message.
func ConfigBody(token string) (string, string) {
    // Define the email subject
    subject := "Email Verification"
    
    // Format the email body with the provided token
    body := fmt.Sprintf(
        `
        <h2>Verify Your Email</h2>
        <hr>
        <p>Click the link below to verify your email:</p>
        <a href="http://localhost:8080/api/verify-email/%s">Verify Email</a>
        `, token)

    // Return the subject and body
    return subject, body
}
