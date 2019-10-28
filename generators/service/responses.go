package service

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
)

type Responses struct {
	Source     string
	Namespace  string
	Package    string
	PACKAGE    string
	DeleteDir  bool
	GrpcPort   string
	EnableHttp bool
	HttpPort   string
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

func NewResponses() *Responses {
	return &Responses{
		Source: getSourceDirectory(),
	}
}
