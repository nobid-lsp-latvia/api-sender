// SPDX-License-Identifier: EUPL-1.2

package routes

import (
	"git.zzdats.lv/edim/api-sender/routes/objects"
	"git.zzdats.lv/edim/api-sender/routes/requests"
	"git.zzdats.lv/edim/api-sender/routes/responses"

	"azugo.io/azugo"
)

// @operationId send message
// @title Send message
// @description Allows to send an e-mail or text
// @param SendRequest body requests.SendRequest true "E-mail or text data"
// @success 200 SendResponse responses.SendResponse "Sending result"
// @failure 400 string string "Bad request"
// @failure 401 {empty} "Unauthorized"
// @failure 403 {empty} "Forbidden"
// @failure 422 string string "Invalid request"
// @failure 500 string string "Internal server error"
// @resource Sending e-mails or texts
// @route /1.0/send [post].
func (r *router) send(ctx *azugo.Context) {
	req := &requests.SendRequest{}
	resp := &responses.SendResponse{Messages: &[]responses.SendMessageResponse{}}

	if err := ctx.Body.JSON(req); err != nil {
		ctx.Error(err)

		return
	}

	for _, toRequest := range *req.To {
		itemResp := &responses.SendMessageResponse{}
		if toRequest.Email != "" {
			// response := &responses.SendResponse{}
			if err := r.MailClient().Send(ctx, toRequest.Email, req.Subject, req.Content, req.From, itemResp); err != nil {
				ctx.Error(err)

				return
			}

			*resp.Messages = append(*resp.Messages, *itemResp)

			continue
		}

		// sms sending
		if toRequest.PhoneNumber != "" {
			if err := r.PhoneClient().Send(ctx, toRequest.PhoneNumber, req.Subject, req.Content, req.From, itemResp); err != nil {
				ctx.Error(err)

				return
			}

			*resp.Messages = append(*resp.Messages, *itemResp)

			continue
		}
	}

	ctx.JSON(resp)
}

// @operationId gets message status
// @title Send message
// @description Allows to send an e-mail or text
// @param trackingId path string true "Message tracking ID"
// @success 200 SubmissionData responses.SubmissionData "Sending result"
// @failure 400 string string "Bad request"
// @failure 401 {empty} "Unauthorized"
// @failure 403 {empty} "Forbidden"
// @failure 422 string string "Invalid request"
// @failure 500 string string "Internal server error"
// @resource Gets submission status
// @route /1.0/{trackingId} [get].
func (r *router) getStatus(ctx *azugo.Context) {
	req := &objects.TrackingID{}
	req.ID = ctx.Params.String("trackingId")

	resp := &responses.SubmissionData{}
	if err := r.Store().Exec(ctx, "sender.get_submission_status", req, &resp); err != nil {
		ctx.Error(err)

		return
	}

	ctx.JSON(resp)
}
