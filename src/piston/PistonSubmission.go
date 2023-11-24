package piston

import "github.com/google/uuid"

type PistonSubmission struct {
	Id        uuid.UUID `json:"id"`
	Language  string    `json:"language"`
	Code      string    `json:"code"`
	TestInput string    `json:"testInput"`
}
