package exercise

import (
	"time"

	"github.com/google/uuid"
	pistonapi "yordanmitev.me/code-checker/api/piston"
	piston "yordanmitev.me/code-checker/piston"
)

type Submission struct {
	Id             uuid.UUID `json:"id" gorm:"primaryKey"`
	Username       string    `json:"username"`
	Language       string    `json:"language"`
	Code           string    `json:"code"`
	SubmissionTime time.Time `json:"submissionTime"`
	Exercise       Exercise  `gorm:"foreignKey:Id"`
}

func (s Submission) RunTests() Solution {
	Solution := Solution{
		Id:             s.Id,
		Exercise:       s.Exercise,
		User:           s.Username,
		SubmissionTime: s.SubmissionTime,
		Language:       s.Language,
		Code:           s.Code,
		IsHidden:       false,
		Score:          0,
	}
	for _, test := range s.Exercise.Tests {
		//
		pistonConn := pistonapi.GetPiston()
		data, _ := pistonConn.SendTest(s.Strip(test))
		outcome := test.CompareData(data)
		Solution.TestResults = append(Solution.TestResults, TestOutcome{
			ControlTest:    test,
			IsCorrect:      outcome,
			Output:         data.Output,
			ExpectedOutput: test.ExpectedOutput,
		})
	}
	return Solution
}

func (s Submission) Strip(test ControlTest) piston.PistonSubmission {
	return piston.PistonSubmission{
		Language:  s.Language,
		Code:      s.Code,
		TestInput: test.Input,
	}
}
