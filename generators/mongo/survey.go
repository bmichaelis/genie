package mongo

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"strings"
)

type Responses struct {
	Enable          bool
	CollectionName  string
	CollectionTitle string
	Database        string
	Username        string
	Password        string
	Port            string
}

type Survey struct {
	Responses *Responses
}

func (s *Survey) Start() *Responses {
	color.Yellow("\nMongo\n------------------------------------------------------\n")

	responses := s.Responses
	if err := survey.AskOne(&survey.Confirm{
		Message: "Enable mongo?",
		Default: true,
	}, &responses.Enable); err != nil {
		panic(err)
	}

	if responses.Enable {

		if err := survey.AskOne(&survey.Input{
			Message: "Database?",
		}, &responses.Database); err != nil {
			panic(err)
		}

		if err := survey.AskOne(&survey.Input{
			Message: "Username?",
		}, &responses.Username); err != nil {
			panic(err)
		}

		if err := survey.AskOne(&survey.Input{
			Message: "Password?",
		}, &responses.Password); err != nil {
			panic(err)
		}

		if err := survey.AskOne(&survey.Input{
			Message: "Port?",
			Default: "27017",
		}, &responses.Port); err != nil {
			panic(err)
		}

		if err := survey.AskOne(&survey.Input{
			Message: "Collection name?",
		}, &responses.CollectionName, survey.WithValidator(survey.Required)); err != nil {
			panic(err)
		}
		responses.CollectionTitle = strings.Title(responses.CollectionName)

	}
	return responses
}

func NewSurvey() *Survey {
	return &Survey{
		Responses: &Responses{},
	}
}
