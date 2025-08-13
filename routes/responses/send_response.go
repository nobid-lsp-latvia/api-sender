// SPDX-License-Identifier: EUPL-1.2

package responses

import "git.zzdats.lv/edim/api-sender/routes/objects"

// SendResponse represents sent message response.
type SendResponse struct {
	// Messages represents list of SendMessageResponse.
	Messages *[]SendMessageResponse `json:"messages"`
}

// SendMessageResponse represents sent message response.
type SendMessageResponse struct {
	objects.TrackingID
}
