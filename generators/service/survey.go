package service

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"strings"
)

type Survey struct {
	Responses *Responses
}

func (s *Survey) Start() *Responses {
	color.Yellow("\nService\n------------------------------------------------------\n")

	if err := survey.AskOne(&survey.Input{
		Message: "Namespace? (ex. github.com/myrepo). Leave empty to skip.",
	}, &s.Responses.Namespace); err != nil {
		panic(err)
	}

	if err := survey.AskOne(&survey.Input{
		Message: "Package name?",
	}, &s.Responses.Package, survey.WithValidator(survey.Required), survey.WithValidator(s.Responses.PackageName)); err != nil {
		panic(err)
	}

	s.Responses.PACKAGE = strings.ToUpper(s.Responses.Package)

	if err := survey.AskOne(&survey.Confirm{
		Message: fmt.Sprintf("Delete directory if exists (%s)", s.Responses.ServicePath()),
		Default: false,
	}, &s.Responses.DeleteDir); err != nil {
		panic(err)
	}

	if err := survey.AskOne(&survey.Input{
		Message: "gRPC port?",
		Default: "8080",
	}, &s.Responses.GrpcPort); err != nil {
		panic(err)
	}

	if err := survey.AskOne(&survey.Confirm{
		Message: "Enable HTTP endpoint?",
		Default: true,
	}, &s.Responses.EnableHttp); err != nil {
		panic(err)
	}

	if s.Responses.EnableHttp {
		if err := survey.AskOne(&survey.Input{
			Message: "HTTP port?",
			Default: "3000",
		}, &s.Responses.HttpPort); err != nil {
			panic(err)
		}
	}
	return s.Responses
}

func NewSurvey() *Survey {
	return &Survey{
		Responses: NewResponses(),
	}
}
