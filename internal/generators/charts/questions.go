package charts

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
)

type Questions struct {
}

func (*Questions) Ask() *Responses {
	color.Yellow("\nCharts\n------------------------------------------------------\n")

	responses := NewResponses()
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
	return responses
}

func NewQuestions() *Questions {
	return &Questions{}
}
