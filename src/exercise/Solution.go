package exercise

import (
	"time"

	"github.com/google/uuid"
)

type Solution struct {
	Id             uuid.UUID     `json:"id" gorm:"primaryKey"`
	Exercise       Exercise      `gorm:"foreignKey:Id"`
	ExerciseId     uuid.UUID     `json:"exerciseId"`
	User           string        `json:"user"`
	Language       string        `json:"language"`
	Code           string        `json:"code"`
	IsHidden       bool          `json:"isHidden"`
	Score          int           `json:"score"`
	SubmissionTime time.Time     `json:"submissionTime"`
	TestResults    []TestOutcome `json:"testResults" gorm:"foreignKey:SolutionId"`
}

func (solution *Solution) GenerateScore() {
	IsCorrectCounter := 0
	for _, test := range solution.TestResults {
		if test.IsCorrect {
			IsCorrectCounter++
		}
	}
	solution.Score = 100 / len(solution.TestResults) * IsCorrectCounter
}
