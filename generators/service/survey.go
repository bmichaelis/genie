package service

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/go-openapi/inflect"
)

type Survey struct {
	Responses *Responses
}

func (s *Survey) Start() (*Responses, error) {
	color.Yellow("\nService\n------------------------------------------------------\n")
	var resp = s.Responses

	// Question 1
	var repositoryUrl string
	if err := survey.AskOne(&survey.Input{
		Message: "Repository Url? (ex. https://github.com/roboncode/time-tracker)",
	}, &repositoryUrl, survey.WithValidator(survey.Required)); err != nil {
		return nil, err
	}

	r1, _ := regexp.Compile("^(https?://)?")
	resp.RepositoryPath = r1.ReplaceAllString(repositoryUrl, "")

	r2, _ := regexp.Compile("([\\w-]+)$")
	foundStrings := r2.FindAllString(resp.RepositoryPath, 1)
	if len(foundStrings) > 0 {
		resp.Package = inflect.Underscore(foundStrings[0])
	}

	// Question 2
	var resource string
	if err := survey.AskOne(&survey.Input{
		Message: "Resource (ex. TimeTracker)? ",
	}, &resource, survey.WithValidator(survey.Required), nil); err != nil {
		return nil, err
	}
	singleResource := inflect.Singularize(resource)
	pluralResource := inflect.Pluralize(resource)
	resp.Resource = inflect.Capitalize(singleResource)
	resp.Resources = inflect.Capitalize(pluralResource)
	resp.EnvVar = strings.ToUpper(inflect.Underscore(singleResource))
	resp.HttpResource = inflect.Dasherize(pluralResource)

	if err := survey.AskOne(&survey.Confirm{
		Message: fmt.Sprintf("Delete directory if exists (%s)", s.Responses.ServicePath()),
		Default: false,
	}, &s.Responses.DeleteDir); err != nil {
		return nil, err
	}

	if err := survey.AskOne(&survey.Input{
		Message: "gRPC port?",
		Default: "8080",
	}, &s.Responses.GrpcPort); err != nil {
		return nil, err
	}

	if err := survey.AskOne(&survey.Confirm{
		Message: "Enable HTTP endpoint?",
		Default: true,
	}, &s.Responses.EnableHttp); err != nil {
		return nil, err
	}

	if s.Responses.EnableHttp {
		if err := survey.AskOne(&survey.Input{
			Message: "HTTP port?",
			Default: "3000",
		}, &s.Responses.HttpPort); err != nil {
			return nil, err
		}
	}

	var b, _ = json.MarshalIndent(s.Responses, "", "   ")
	fmt.Println("\nservice", string(b))

	return s.Responses, nil
}

func NewSurvey() *Survey {
	return &Survey{
		Responses: NewResponses(),
	}
}
