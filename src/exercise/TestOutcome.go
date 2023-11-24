package exercise

import "github.com/google/uuid"

type TestOutcome struct {
	Id             uint64      `json:"id" gorm:"primaryKey"`
	SolutionId     uuid.UUID   `json:"solutionId"`
	ControlTest    ControlTest `json:"testId" gorm:"foreignKey:id"`
	Input          string      `json:"input"`
	Output         string      `json:"output"`
	ExpectedOutput string      `json:"expectedOutput"`
	IsResultHidden bool        `json:"isHidden"`
	IsCorrect      bool        `json:"isCorrect"`
}
