package exercise

import (
	"github.com/google/uuid"
	"yordanmitev.me/code-checker/piston"
)

type ControlTest struct {
	Id             uint64    `json:"id" gorm:"primaryKey; autoIncrement:1"`
	ExerciseId     uuid.UUID `json:"exerciseId" gorm:"foreignKey:Id"`
	Input          string    `json:"input"`
	ExpectedOutput string    `json:"output"`
	IsHidden       bool      `json:"isHidden"`
}

func (t ControlTest) CompareData(piston.SubmissionResponse) bool {
	return true
}
