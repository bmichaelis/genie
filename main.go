package main

import (
	"bufio"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/gobuffalo/packr/v2"
	"grpcgen/internal"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

const templateSuffix = ".tmpl"

var qs = []*survey.Question{
	{
		Name:      "Package",
		Prompt:    &survey.Input{Message: "Name of package?"},
		Validate:  survey.Required,
		Transform: survey.ToLower,
	},
}

func main() {
	answers := internal.NewAnswers()

	if err := survey.Ask(qs, &answers); err != nil {
		panic(err)
	}

	// Remove directory
	if _, err := os.Stat(answers.Package); !os.IsNotExist(err) {
		_ = os.RemoveAll(answers.Package)
	}

	box := packr.New("templates", "./templates")
	err := box.Walk(func(path string, file packr.File) error {
		// TODO: clear individual files

		fullOutputPath := fmt.Sprintf("%s/%s", answers.Package, path)
		if strings.Contains(fullOutputPath, templateSuffix) {
			fullOutputPath = strings.TrimSuffix(fullOutputPath, filepath.Ext(fullOutputPath))
		}

		// create directory path
		dir := filepath.Dir(fullOutputPath)
		if _, err := os.Stat(fullOutputPath); os.IsNotExist(err) {
			_ = os.MkdirAll(dir, os.ModePerm)
		}

		t, err := template.New(path).Parse(file.String())
		if err != nil {
			return err
		}
		f, err := os.Create(fullOutputPath)
		if err != nil {
			return err
		}
		defer f.Close()
		w := bufio.NewWriter(f)
		_ = t.Execute(w, &answers)
		return w.Flush()
	})
	if err != nil {
		panic(err)
	}
	if err := os.Chdir(answers.Package); err != nil {
		panic(err)
	}
	cmd := exec.Command("chmod", "+x", "generate.sh")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	cmd = exec.Command("go", "generate")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	cmd = exec.Command("go", "fmt", "./...")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
