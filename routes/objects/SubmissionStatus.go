// SPDX-License-Identifier: EUPL-1.2

package objects

// SubmissionStatus represents the status of message.
type SubmissionStatus struct {
	// Status represents the status of the message: `sent`, `error`.
	Status string `json:"status" example:"sent"`
	// Info represents the additional information about the status.
	Info string `json:"info,omitempty" example:"error message"`
}
