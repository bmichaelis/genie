package charts

import (
	"encoding/json"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
)

type Responses struct {
	Enable         bool
	ArtifactoryUrl string
	Author         string
	Email          string
}

type Survey struct {
	Responses *Responses
}

func (s *Survey) Start() *Responses {
	color.Yellow("\nCharts\n------------------------------------------------------\n")

	responses := s.Responses
	if err := survey.AskOne(&survey.Confirm{
		Message: "Add Kubernetes Helm Charts?",
		Default: false,
	}, &responses.Enable); err != nil {
		panic(err)
	}

	if responses.Enable {
		if err := survey.AskOne(&survey.Input{
			Message: "Artifactory Url?",
		}, &responses.ArtifactoryUrl, survey.WithValidator(survey.Required)); err != nil {
			panic(err)
		}

		if err := survey.AskOne(&survey.Input{
			Message: "Domain Url (ex. roboncode.com)?",
		}, &responses.ArtifactoryUrl, survey.WithValidator(survey.Required)); err != nil {
			panic(err)
		}

		if err := survey.AskOne(&survey.Input{
			Message: "Author's name?",
		}, &responses.Author, survey.WithValidator(survey.Required)); err != nil {
			panic(err)
		}

		if err := survey.AskOne(&survey.Input{
			Message: "Author's email?",
		}, &responses.Email, survey.WithValidator(survey.Required)); err != nil {
			panic(err)
		}
	}

	var b, _ = json.MarshalIndent(responses, "", "   ")
	fmt.Println("\ncharts", string(b))

	return responses
}

func NewSurvey() *Survey {
	return &Survey{
		Responses: &Responses{},
	}
}
