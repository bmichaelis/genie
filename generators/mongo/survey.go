package mongo

import (
	"encoding/json"
	"fmt"
	"genie/generators"
	"genie/generators/service"
	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/go-openapi/inflect"
	"strings"
)

type Responses struct {
	Enable     bool
	Collection string
	Database   string
	Username   string
	Password   string
	Address    string
}

type Survey struct {
	Responses *Responses
}

func (s *Survey) Start() *Responses {
	color.Yellow("\nMongo\n------------------------------------------------------\n")

	responses := s.Responses
	if err := survey.AskOne(&survey.Confirm{
		Message: "Use mongo?",
		Default: false,
	}, &responses.Enable); err != nil {
		panic(err)
	}

	if responses.Enable {

		if err := survey.AskOne(&survey.Input{
			Message: "Database name?",
			Default: "default",
		}, &responses.Database); err != nil {
			panic(err)
		}

		if err := survey.AskOne(&survey.Input{
			Message: "Username? (leave empty to skip)",
		}, &responses.Username); err != nil {
			panic(err)
		}

		if err := survey.AskOne(&survey.Input{
			Message: "Password? (leave empty to skip)",
		}, &responses.Password); err != nil {
			panic(err)
		}

		if err := survey.AskOne(&survey.Input{
			Message: "Mongo address?",
			Default: "localhost:27017",
		}, &responses.Address); err != nil {
			panic(err)
		}

		s := generators.GetInstance().Find(service.NAME).(*service.Generator)
		pluralResource := inflect.Pluralize(s.Responses.Resource)
		responses.Collection = strings.ToLower(inflect.Underscore(pluralResource))

	}

	var b, _ = json.MarshalIndent(responses, "", "   ")
	fmt.Println("\nmongo", string(b))

	return responses
}

func NewSurvey() *Survey {
	return &Survey{
		Responses: &Responses{},
	}
}
