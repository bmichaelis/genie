package internal

import (
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"os"
	"reflect"
	"regexp"
	"strings"
)

type Survey struct {
	Responses Responses
}

type Responses struct {
	Source     string
	Namespace  string
	Package    string
	DeleteDir  bool
	GrpcPort   string
	EnableHttp bool
	HttpPort   string
}

func getSourceDirectory() string {
	// see if they are using go path and attempt to build the target dir from it
	// otherwise use cwd
	var targetPath string
	gopath := os.Getenv("GOPATH")

	cwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	if len(gopath) > 0 {
		targetPath = os.ExpandEnv("$GOPATH/src")

		if _, err := os.Stat(targetPath); os.IsNotExist(err) {
			targetPath = cwd
		}
	} else {
		targetPath = cwd
	}
	return targetPath
}

func (s *Survey) Start() {
	responses := s.Responses

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

	s.Responses = responses
}

func (r *Responses) PackageName(val interface{}) error {
	value := reflect.ValueOf(val)
	invalid, err := regexp.MatchString("[\\W0-9A-Z]+", value.String())
	if err != nil {
		return err
	}
	if invalid {
		return errors.New("package name can only contain lowercase letters and underscores")
	}
	return nil
}

func (r *Responses) ServicePath() string {
	if r.Namespace == "" {
		return strings.ToLower(fmt.Sprintf("%s/%s", r.Source, r.Package))
	}
	return strings.ToLower(fmt.Sprintf("%s/%s/%s", r.Source, r.Namespace, r.Package))
}

func NewSurvey() Survey {
	return Survey{
		Responses: Responses{
			Source: getSourceDirectory(),
		},
	}
}
