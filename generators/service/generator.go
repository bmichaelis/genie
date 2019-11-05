package service

import (
	"bufio"
	"fmt"
	"genie/util"
	"os"
	"path/filepath"
	"text/template"

	"github.com/gobuffalo/packr/v2"
)

const NAME = "service"

var terminal = util.NewTerminal()

type Generator struct {
	Responses *Responses
}

func (g *Generator) GetName() string {
	return NAME
}

func (g *Generator) AskQuestions() error {
	var err error
	g.Responses, err = NewSurvey().Start()
	if err != nil {
		return err
	}
	return nil
}

func (g *Generator) Execute() {
	g.deleteDir()
	g.writeFiles()
}

func (g *Generator) deleteDir() {
	responses := g.Responses
	if responses.DeleteDir {
		s := terminal.ShowBusy(fmt.Sprintf("Deleting directory and its contents in %s...", responses.ServicePath()))
		if _, err := os.Stat(responses.ServicePath()); !os.IsNotExist(err) {
			_ = os.RemoveAll(responses.ServicePath())
		}
		s.Stop()
	}
}

func (g *Generator) writeFiles() {
	responses := g.Responses
	box := packr.New(NAME, "./_templates/")
	err := box.Walk(func(path string, file packr.File) error {
		fullOutputPath := fmt.Sprintf("%s/%s", responses.ServicePath(), path)
		// create directory path
		dir := filepath.Dir(fullOutputPath)
		if _, err := os.Stat(fullOutputPath); os.IsNotExist(err) {
			_ = os.MkdirAll(dir, os.ModePerm)
		}

		content := file.String()

		t, err := template.New(path).Parse(content)
		if err != nil {
			return err
		}
		f, err := os.Create(fullOutputPath)
		if err != nil {
			return err
		}
		defer f.Close()
		w := bufio.NewWriter(f)
		_ = t.Execute(w, map[string]interface{}{
			g.GetName(): g.Responses,
		})
		return w.Flush()
	})
	if err != nil {
		panic(err)
	}
}

func NewGenerator() *Generator {
	return &Generator{}
}
