package internal

import (
	"fmt"
	"os"
	"strings"
)

type Answers struct {
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

func (a *Answers) ServicePath() string {
	if a.Namespace == "" {
		return strings.ToLower(fmt.Sprintf("%s/%s", a.Source, a.Package))
	}
	return strings.ToLower(fmt.Sprintf("%s/%s/%s", a.Source, a.Namespace, a.Package))
}

func NewAnswers() Answers {
	return Answers{
		Source: getSourceDirectory(),
	}
}
