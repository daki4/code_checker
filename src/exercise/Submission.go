package exercise

import (
	"time"

	"github.com/google/uuid"
	pistonapi "yordanmitev.me/code-checker/api/piston"
	piston "yordanmitev.me/code-checker/piston"
	"strings"
)

type Submission struct {
	Id             uuid.UUID `json:"id" gorm:"primaryKey"`
	ExerciseId 	   uuid.UUID `json:"-" gorm:"references:"`
	Username       string    `json:"username"`
	Language       string    `json:"language"`
	Code           string    `json:"code"`
	SubmissionTime time.Time `json:"submissionTime"`
	Exercise       Exercise  `json:"-" gorm:"references:ExerciseId"`
}

func (s Submission) RunTests() Solution {
	solution := Solution{
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
		data.Output = strings.Trim(data.Output, "\n")
		outcome := test.CompareData(data)
		id, _ := uuid.NewUUID()
		solution.TestResults = append(solution.TestResults, TestOutcome{
			Id:             id,
			SolutionId:     s.Id,
			ControlTest:    test,
			Input:          test.Input,
			IsCorrect:      outcome,
			Output:         data.Output,
			ExpectedOutput: test.ExpectedOutput,
		})
	}
	return solution
}

func (s Submission) Strip(test ControlTest) piston.PistonSubmission {
	return piston.PistonSubmission{
		Id:        s.Id,
		Language:  s.Language,
		Code:      s.Code,
		TestInput: test.Input,
	}
}
