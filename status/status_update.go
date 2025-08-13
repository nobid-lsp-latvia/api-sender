// SPDX-License-Identifier: EUPL-1.2

package status

import (
	"git.zzdats.lv/edim/api-sender/routes/requests"
	"git.zzdats.lv/edim/api-sender/routes/responses"

	"azugo.io/azugo"
	jsondb "github.com/nobid-lsp-latvia/lx-go-jsondb"
)

func UpdateSubmissionStatus(store jsondb.Store, ctx *azugo.Context, submissionData *requests.UpdateSubmissionData, status string, errorMessage string) {
	res := &responses.SubmissionData{}
	submissionData.Status = status

	if submissionData.Status == "error" {
		submissionData.Info = errorMessage
	}

	if err := store.Exec(ctx, "sender.update_submission_status", submissionData, &res); err != nil {
		return
	}
}
