// SPDX-License-Identifier: EUPL-1.2

package responses

import (
	"git.zzdats.lv/edim/api-sender/routes/objects"
	"git.zzdats.lv/edim/api-sender/util"
)

// SubmissionData represents the response of the submission data.
type SubmissionData struct {
	objects.TrackingID
	objects.SubmissionStatus
	// SentOn represents the date when the submission was sent.
	SentOn *util.Date `json:"sendOn,omitempty" example:"2020-08-24"`
}
