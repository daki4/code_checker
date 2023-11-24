package piston

import "github.com/google/uuid"

type SubmissionResponse struct {
	SubmissionId uuid.UUID `json:"submissionId"`
	Output       string    `json:"output"`
}
