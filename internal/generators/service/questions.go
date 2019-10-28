package service

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"strings"
)

type Questions struct {
}

func (*Questions) Ask() *Responses {
	color.Yellow("\nService\n------------------------------------------------------\n")

	responses := NewResponses()
	if err := survey.AskOne(&survey.Input{
		Message: "Namespace? (ex. github.com/myrepo). Leave empty to skip.",
	}, &responses.Namespace); err != nil {
		panic(err)
	}

	if err := survey.AskOne(&survey.Input{
		Message: "Package name?",
	}, &responses.Package, survey.WithValidator(survey.Required), survey.WithValidator(responses.PackageName)); err != nil {
		panic(err)
	}

	responses.PACKAGE = strings.ToUpper(responses.Package)

	if err := survey.AskOne(&survey.Confirm{
		Message: fmt.Sprintf("Delete directory if exists (%s)", responses.ServicePath()),
		Default: false,
	}, &responses.DeleteDir); err != nil {
		panic(err)
	}

	if err := survey.AskOne(&survey.Input{
		Message: "gRPC port?",
		Default: "8080",
	}, &responses.GrpcPort); err != nil {
		panic(err)
	}

	if err := survey.AskOne(&survey.Confirm{
		Message: "Enable HTTP endpoint?",
		Default: true,
	}, &responses.EnableHttp); err != nil {
		panic(err)
	}

	if responses.EnableHttp {
		if err := survey.AskOne(&survey.Input{
			Message: "HTTP port?",
			Default: "3000",
		}, &responses.HttpPort); err != nil {
			panic(err)
		}
	}
	return responses
}

func NewQuestions() *Questions {
	return &Questions{}
}
