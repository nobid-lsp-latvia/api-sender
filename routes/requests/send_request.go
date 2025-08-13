// SPDX-License-Identifier: EUPL-1.2

package requests

// SendRequest represents the request to send an email or SMS.
type SendRequest struct {
	// To represents list of recipients.
	To *[]SendToRequest `json:"to"`
	SendContent
}

// SendContent represents the content of the email or SMS.
type SendContent struct {
	// From represents the sender email or sms information.
	From *SendFromRequest `json:"from,omitempty"`
	// Subject represents the subject of the email.
	Subject string `json:"subject" example:"Subject text"`
	// Content represents the content of the email or SMS.
	Content *[]SendMessageContentRequest `json:"content"`
}

// SendToRequest represents the email or SMS address.
type SendToRequest struct {
	// Email represents the email address.
	Email string `json:"email,omitempty" example:"example@email.com"`
	// PhoneNumber represents the phone number.
	PhoneNumber string `json:"phoneNumber,omitempty" example:"+1234567890"`
}

// SendFromRequest represents the sender of the email or SMS.
type SendFromRequest struct {
	// Name represents the name of the sender.
	Name string `json:"name,omitempty" example:"Service"`
	SendToRequest
}

// SendMessageContentRequest represents the content of the email or SMS.
type SendMessageContentRequest struct {
	// MessageType represents the type of the message, either `text/html` or `text/plain`.
	MessageType string `json:"type" example:"text/html"`
	// MessageValue represents the value of the message.
	MessageValue string `json:"value" example:"<p>Test email message</p>"`
}
