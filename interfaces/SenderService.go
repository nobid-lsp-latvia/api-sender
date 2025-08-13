// SPDX-License-Identifier: EUPL-1.2

package interfaces

import (
	"git.zzdats.lv/edim/api-sender/routes/requests"
	"git.zzdats.lv/edim/api-sender/routes/responses"

	"azugo.io/azugo"
)

type SenderService interface {
	Send(ctx *azugo.Context, to string, subject string, content *[]requests.SendMessageContentRequest, from *requests.SendFromRequest, resp *responses.SendMessageResponse) error
}
