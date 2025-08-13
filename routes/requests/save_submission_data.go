// SPDX-License-Identifier: EUPL-1.2

package requests

import "git.zzdats.lv/edim/api-sender/routes/objects"

type SaveSubmissionData struct {
	To string `json:"to"`
	objects.SubmissionStatus
	SendContent
}
