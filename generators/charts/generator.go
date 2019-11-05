package charts

import (
	"bufio"
	"fmt"
	"genie/generators"
	"genie/generators/service"
	"github.com/gobuffalo/packr/v2"
	"os"
	"path/filepath"
	"text/template"
)

const (
	NAME  = "charts"
	LEFT  = "{{{"
	RIGHT = "}}}"
)

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
	if g.Responses.Enable {
		g.writeFiles()
	}
}

func (g *Generator) Finalize() {}

func (g *Generator) writeFiles() {
	s := generators.GetInstance().Find(service.NAME).(*service.Generator)
	box := packr.New(NAME, "./_templates")
	err := box.Walk(func(path string, file packr.File) error {
		fullOutputPath := fmt.Sprintf("%s/charts/%s/%s", s.Responses.ServicePath(), s.Responses.Package, path)
		// create directory path
		dir := filepath.Dir(fullOutputPath)
		if _, err := os.Stat(fullOutputPath); os.IsNotExist(err) {
			_ = os.MkdirAll(dir, os.ModePerm)
		}

		content := file.String()

		t, err := template.New(path).Delims(LEFT, RIGHT).Parse(content)
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
			s.GetName(): s.Responses,
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
