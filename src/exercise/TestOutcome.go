package exercise

import "github.com/google/uuid"

type TestOutcome struct {
	Id             uuid.UUID   `json:"id" gorm:"primaryKey"`
	SolutionId     uuid.UUID   `json:"-"`
	ControlTest    ControlTest `json:"controlTest" gorm:"foreignKey:id"`
	Input          string      `json:"input"`
	Output         string      `json:"output"`
	ExpectedOutput string      `json:"expectedOutput"`
	IsResultHidden bool        `json:"isHidden"`
	IsCorrect      bool        `json:"isCorrect"`
}
