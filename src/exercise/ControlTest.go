package exercise

import (
	"github.com/google/uuid"
	"yordanmitev.me/code-checker/piston"
)

type ControlTest struct {
	Id             uuid.UUID `json:"id" gorm:"primaryKey"`
	ExerciseId     uuid.UUID `json:"exerciseId" gorm:"foreignKey:Id"`
	Exercise       Exercise  `json:"-" gorm:"references`
	Input          string    `json:"input"`
	ExpectedOutput string    `json:"expectedOutput"`
	IsHidden       bool      `json:"isHidden"`
}

func (t ControlTest) CompareData(submission piston.SubmissionResponse) bool {
	return t.ExpectedOutput == submission.Output
}
