package service

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"regexp"
)

type Responses struct {
	GoSourcePath   string
	RepositoryPath string // https://github.com/roboncode/awesomesauce-api (input)
	Resource       string // AwesomeSauce (input)
	Resources      string // AwesomeSauces (formatted from Resource)
	Package        string // awesomesauce_api (extracted from RepositoryPath and formatted)
	EnvVar         string // AWESOME_SAUCE (formatted from Resource)
	HttpResource   string // awesome-sauce (formatted from Resource)
	DeleteDir      bool
	GrpcPort       string
	EnableHttp     bool
	HttpPort       string
}

func (r *Responses) ValidateResource(val interface{}) error {
	value := reflect.ValueOf(val)
	invalid, err := regexp.MatchString("[[A-Z]\\w+]+", value.String())
	if err != nil {
		return err
	}
	if invalid {
		return errors.New("resource name must be start with uppercase letter and CamelCase")
	}
	return nil
}

func (r *Responses) ValidatePackageName(val interface{}) error {
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
	return fmt.Sprintf("%s/%s", r.GoSourcePath, r.RepositoryPath)
}

func getGoSourcePath() string {
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

func NewResponses() *Responses {
	return &Responses{
		GoSourcePath: getGoSourcePath(),
	}
}
