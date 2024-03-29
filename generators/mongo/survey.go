package mongo

import (
	"encoding/json"
	"fmt"
	"genie/generators"
	"genie/generators/service"
	"regexp"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/go-openapi/inflect"
)

type Responses struct {
	Enable      bool
	Collection  string
	Database    string
	Credentials string
	Username    string
	Password    string
	Address     string
	Port        string
}

type Survey struct {
	Responses *Responses
}

func (s *Survey) Start() (*Responses, error) {
	color.Yellow("\nMongo\n------------------------------------------------------\n")

	responses := s.Responses
	if err := survey.AskOne(&survey.Confirm{
		Message: "Use mongo?",
		Default: false,
	}, &responses.Enable); err != nil {
		return nil, err
	}

	if responses.Enable {

		if err := survey.AskOne(&survey.Input{
			Message: "Database name?",
			Default: "default",
		}, &responses.Database); err != nil {
			return nil, err
		}

		if err := survey.AskOne(&survey.Input{
			Message: "Username? (leave empty to skip)",
		}, &responses.Username); err != nil {
			return nil, err
		}

		if responses.Username != "" {
			if err := survey.AskOne(&survey.Input{
				Message: "Password? (required)",
			}, &responses.Password); err != nil {
				return nil, err
			}
			responses.Credentials = fmt.Sprintf("%s:%s@", responses.Username, responses.Password)
		}

		if err := survey.AskOne(&survey.Input{
			Message: "Mongo address?",
			Default: "localhost:27017",
		}, &responses.Address); err != nil {
			return nil, err
		}

		r, _ := regexp.Compile("\\d+")
		responses.Port = r.FindString(responses.Address)

		s := generators.GetInstance().Find(service.NAME).(*service.Generator)
		pluralResource := inflect.Pluralize(s.Responses.Resource)
		responses.Collection = strings.ToLower(inflect.Underscore(pluralResource))
	}

	var b, _ = json.MarshalIndent(responses, "", "   ")
	fmt.Println("\nmongo", string(b))

	return responses, nil
}

func NewSurvey() *Survey {
	return &Survey{
		Responses: &Responses{},
	}
}
