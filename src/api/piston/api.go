package piston

import (
	gopiston "github.com/milindmadhukar/go-piston"
	"yordanmitev.me/code-checker/piston"
)

type Piston struct {
	// Piston is a struct that represents a piston code engine.
	*gopiston.Client
}

// private instance of piston, passed by value, not reference
var pistonConn Piston

// start a web server that sends and listens for requests from piston
func Start() {
	// client := gopiston.New("", http.DefaultClient, "http://localhost:2000/api/v2/piston/")
	pistonConn.Client = gopiston.CreateDefaultClient()
	// fmt.Printf("%s", client.GetLanguages())
	// pistonConn.Client = gopiston.New("", http.DefaultClient, "http://localhost:2000")
}

// get the supported languages from piston
func (p Piston) GetInstalledLanguages() []string {
	return *p.Client.GetLanguages()
}

// send a test to piston and get a response
func (p Piston) SendTest(submission piston.PistonSubmission) (piston.SubmissionResponse, error) {
	// p.Client.Execute
	output, err := p.Client.Execute(submission.Language, "", // Passing language. Since no version is specified, it uses the latest supported version.
		[]gopiston.Code{
			{Content: submission.Code},
		}, // Passing Code.
		gopiston.Stdin(submission.TestInput), // Passing input as "hello world".
	)
	if err != nil {
		return piston.SubmissionResponse{}, err
	}
	return piston.SubmissionResponse{
		SubmissionId: submission.Id,
		Output:       output.GetOutput(),
	}, nil
}

func GetPiston() Piston {
	return pistonConn
}
