package main

import (
	"bufio"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/gobuffalo/packr/v2"
	"github.com/janeczku/go-spinner"
	"grpcgen/internal"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

const templateSuffix = ".tmpl"

var spinnerCharSet = []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}
var spinnerSpeed = time.Millisecond * 50

func askQuestions() *internal.Answers {
	answers := internal.NewAnswers()

	if err := survey.AskOne(&survey.Input{
		Message: "What namespace is this service under? (ex. github.com/myrepo). Leave empty to skip.",
	}, &answers.Namespace); err != nil {
		panic(err)
	}

	if err := survey.AskOne(&survey.Input{
		Message: "Name of package?",
	}, &answers.Package, survey.WithValidator(survey.Required)); err != nil {
		panic(err)
	}

	if err := survey.AskOne(&survey.Confirm{
		Message: fmt.Sprintf("Delete directory if exists (%s)", answers.ServicePath()),
		Default: false,
	}, &answers.DeleteDir); err != nil {
		panic(err)
	}

	return &answers
}

func setWorkingDir(answers *internal.Answers) {
	if err := os.Chdir(answers.ServicePath()); err != nil {
		panic(err)
	}
}

func deleteDir(answers *internal.Answers) {
	if answers.DeleteDir {
		s := spinner.StartNew(fmt.Sprintf("Deleting directory and its contents in %s...", answers.ServicePath()))
		s.SetSpeed(spinnerSpeed)
		s.SetCharset(spinnerCharSet)
		if _, err := os.Stat(answers.ServicePath()); !os.IsNotExist(err) {
			_ = os.RemoveAll(answers.ServicePath())
		}
		s.Stop()
	}
}

func generateFiles(answers *internal.Answers) {
	box := packr.New("templates", "./templates")
	err := box.Walk(func(path string, file packr.File) error {
		fullOutputPath := fmt.Sprintf("%s/%s", answers.ServicePath(), path)
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
		_ = t.Execute(w, map[string]string{
			"Package": answers.Package,
		})
		return w.Flush()
	})
	if err != nil {
		panic(err)
	}
}

func goChmod() {
	cmd := exec.Command("chmod", "+x", "generate.sh")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func goGenerate() {
	s := spinner.StartNew("go generate...")
	s.SetSpeed(spinnerSpeed)
	s.SetCharset(spinnerCharSet)
	cmd := exec.Command("go", "generate")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	s.Stop()
}

func goFmt() {
	s := spinner.StartNew("go fmt")
	s.SetSpeed(spinnerSpeed)
	s.SetCharset(spinnerCharSet)
	cmd := exec.Command("go", "fmt", "./...")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	s.Stop()
}

func printInstructions(answers *internal.Answers) {
	fmt.Println("\nGeneration successful!!")
	fmt.Println("------------------------------------------------")
	fmt.Printf("cd %s; go run cmd/main.go\n\n", answers.ServicePath())
}

func main() {
	answers := askQuestions()
	deleteDir(answers)
	generateFiles(answers)
	setWorkingDir(answers)
	goChmod()
	goGenerate()
	goFmt()
	printInstructions(answers)
}
