// SPDX-License-Identifier: EUPL-1.2

package mail

import (
	"crypto/tls"
	"fmt"
	"strings"

	"git.zzdats.lv/edim/api-sender/interfaces"
	"git.zzdats.lv/edim/api-sender/routes/requests"
	"git.zzdats.lv/edim/api-sender/routes/responses"
	"git.zzdats.lv/edim/api-sender/status"
	"github.com/nobid-lsp-latvia/lx-go-jsondb"

	"azugo.io/azugo"
	"azugo.io/core"
	"github.com/wneessen/go-mail"
)

type inst struct {
	app    *core.App
	config *Configuration
	store  jsondb.Store
}

func newMailService(app *core.App, config *Configuration, store jsondb.Store) (interfaces.SenderService, error) {
	return &inst{
		app:    app,
		config: config,
		store:  store,
	}, nil
}

func (i *inst) Send(ctx *azugo.Context, to string, subject string, content *[]requests.SendMessageContentRequest, from *requests.SendFromRequest, resp *responses.SendMessageResponse) error {
	senderEmail := i.config.SenderMail
	senderName := i.config.SenderMailName

	if from != nil && from.Name != "" {
		senderName = strings.TrimSpace(from.Name)
	}

	if from != nil && from.Email != "" {
		senderEmail = strings.TrimSpace(from.Email)
	}

	if senderEmail != i.config.SenderMail {
		err := jsondb.ExecError{
			Code:    "bad_email",
			Message: "incorrect sender email address",
		}

		return err
	}

	saveSubmissionData := &requests.SaveSubmissionData{
		To: to,
	}
	saveSubmissionData.From = &requests.SendFromRequest{
		Name: senderName,
	}
	saveSubmissionData.From.Email = senderEmail
	saveSubmissionData.Subject = subject
	saveSubmissionData.Content = content

	submissionData := &responses.SubmissionData{}

	if err := i.store.Exec(ctx, "sender.save_submission_data", saveSubmissionData, &submissionData); err != nil {
		return err
	}

	updateSubmissionData := &requests.UpdateSubmissionData{}
	updateSubmissionData.TrackingID = submissionData.TrackingID

	message := mail.NewMsg()
	if err := message.FromFormat(senderName, senderEmail); err != nil {
		status.UpdateSubmissionStatus(i.store, ctx, updateSubmissionData, "error", fmt.Sprintf("failed to set FROM address: %s", err))

		return err
	}

	if err := message.To(to); err != nil {
		status.UpdateSubmissionStatus(i.store, ctx, updateSubmissionData, "error", fmt.Sprintf("failed to set TO address: %s", err))

		return err
	}

	message.Subject(subject)

	for index, contentData := range *content {
		messageType := mail.TypeTextPlain
		if contentData.MessageType == "text/html" {
			messageType = mail.TypeTextHTML
		}

		if index == 0 {
			message.SetBodyString(messageType, contentData.MessageValue)
		} else {
			message.AddAlternativeString(messageType, contentData.MessageValue)
		}
	}

	// Deliver the mails via SMTP
	client, err := mail.NewClient(i.config.MailHost, mail.WithPort(i.config.MailPort),
		mail.WithSMTPAuth(mail.SMTPAuthLogin), mail.WithTLSPortPolicy(mail.TLSOpportunistic),
	)
	if err != nil {
		status.UpdateSubmissionStatus(i.store, ctx, updateSubmissionData, "error", fmt.Sprintf("failed to create new mail delivery client: %s", err))

		return err
	}

	err = client.SetTLSConfig(&tls.Config{InsecureSkipVerify: i.config.MailSkipVerify}) //nolint:gosec
	if err != nil {
		status.UpdateSubmissionStatus(i.store, ctx, updateSubmissionData, "error", fmt.Sprintf("failed to set client tls config: %s", err))

		return err
	}

	client.SetSSL(i.config.MailSSL)

	// Set username and password if provided
	if i.config.MailUser != "" && i.config.MailPassword != "" {
		client.SetUsername(i.config.MailUser)
		client.SetPassword(i.config.MailPassword)
	}

	if err := client.DialAndSend(message); err != nil {
		status.UpdateSubmissionStatus(i.store, ctx, updateSubmissionData, "error", fmt.Sprintf("failed to deliver mail: %s", err))

		return err
	}

	status.UpdateSubmissionStatus(i.store, ctx, updateSubmissionData, "sent", "")
	ctx.Log().Debug("Test mail successfully delivered.")

	resp.ID = updateSubmissionData.ID

	return nil
}
