// SPDX-License-Identifier: EUPL-1.2

package requests

import "git.zzdats.lv/edim/api-sender/routes/objects"

type UpdateSubmissionData struct {
	objects.TrackingID
	objects.SubmissionStatus
	MessageID string `json:"messageId,omitempty"`
}
