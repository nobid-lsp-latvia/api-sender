// SPDX-License-Identifier: EUPL-1.2

package phone

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"git.zzdats.lv/edim/api-sender/interfaces"
	"git.zzdats.lv/edim/api-sender/routes/requests"
	"git.zzdats.lv/edim/api-sender/routes/responses"
	"git.zzdats.lv/edim/api-sender/status"
	"github.com/nobid-lsp-latvia/lx-go-jsondb"

	"azugo.io/azugo"
	"azugo.io/core"
	"azugo.io/core/http"
)

type inst struct {
	app    *core.App
	config *Configuration
	store  jsondb.Store
}

func newPhoneService(app *core.App, config *Configuration, store jsondb.Store) (interfaces.SenderService, error) {
	return &inst{
		app:    app,
		config: config,
		store:  store,
	}, nil
}

func (i *inst) httpClient(ctx *azugo.Context, path string) (http.Client, []http.RequestOption, error) {
	baseURL, err := url.JoinPath(i.config.PhoneURL, path)
	if err != nil {
		return nil, nil, err
	}

	client := ctx.HTTPClient().WithBaseURL(baseURL)

	return client, []http.RequestOption{}, nil
}

func (i *inst) send(ctx *azugo.Context, path string, apiKey string, sender string, number string, text string) (string, error) {
	client, params, err := i.httpClient(ctx, path)
	if err != nil {
		return "", err
	}

	params = append(params, http.WithQueryArg("api-key", apiKey))
	params = append(params, http.WithQueryArg("sender", sender))
	params = append(params, http.WithQueryArg("number", number))
	params = append(params, http.WithQueryArg("text", text))

	if i.config.Debug {
		params = append(params, http.WithQueryArg("flag-debug", "1"))
	}

	response, err := client.Get("", params...)
	if err != nil {
		return "", err
	}

	responseBody := string(response)
	re := regexp.MustCompile(`\d+`)
	numberStr := re.FindString(responseBody)

	numberInt, err := strconv.Atoi(numberStr)
	if err != nil {
		return "", err
	}
	// need to check response if Response code is greater than 100, it is the message SMS ID.
	if numberInt < 100 {
		err := jsondb.ExecError{
			Code:    "unsuccessful",
			Message: "sms submit failed with code: " + responseBody,
		}

		return "", err
	}

	return responseBody, nil
}

func (i *inst) Send(ctx *azugo.Context, to string, subject string, content *[]requests.SendMessageContentRequest, from *requests.SendFromRequest, resp *responses.SendMessageResponse) error {
	senderName := i.config.SenderPhoneName

	if from != nil && from.Name != "" {
		senderName = strings.TrimSpace(from.Name)
	}

	if senderName != i.config.SenderPhoneName {
		err := jsondb.ExecError{
			Code:    "bad_phone",
			Message: "incorrect sender phone",
		}

		return err
	}

	saveSubmissionData := &requests.SaveSubmissionData{
		To: to,
	}
	saveSubmissionData.From = &requests.SendFromRequest{}
	saveSubmissionData.From.Name = senderName
	saveSubmissionData.Subject = subject
	saveSubmissionData.Content = content

	submissionData := &responses.SubmissionData{}

	if err := i.store.Exec(ctx, "sender.save_submission_data", saveSubmissionData, &submissionData); err != nil {
		return err
	}

	updateSubmissionData := &requests.UpdateSubmissionData{}
	updateSubmissionData.TrackingID = submissionData.TrackingID
	messageText := ""

	for _, contentData := range *content {
		if contentData.MessageType == "text/plain" {
			if messageText == "" {
				messageText = contentData.MessageValue
			} else {
				messageText = messageText + " " + contentData.MessageValue
			}
		}
	}

	result, err := i.send(ctx, "/send", i.config.PhoneAPIKey, senderName, to, messageText)
	updateSubmissionData.MessageID = result

	if err != nil {
		status.UpdateSubmissionStatus(i.store, ctx, updateSubmissionData, "error", fmt.Sprintf("failed to send message: %s", err))

		return err
	}

	status.UpdateSubmissionStatus(i.store, ctx, updateSubmissionData, "sent", "")
	ctx.Log().Debug("Test mail successfully delivered.")

	resp.ID = updateSubmissionData.ID

	return nil
}
