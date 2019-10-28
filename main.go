package main

import (
	"bufio"
	"fmt"
	"genny/internal"
	"github.com/gobuffalo/packr/v2"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

const templateSuffix = ".tmpl"

var terminal = internal.NewTerminal()

func printHeader() {
	header := internal.NewHeader()
	header.Print()
}

func askQuestions() *internal.Responses {
	survey := internal.NewSurvey()
	survey.Start()
	return &survey.Responses
}

func setWorkingDir(answers *internal.Responses) {
	if err := os.Chdir(answers.ServicePath()); err != nil {
		panic(err)
	}
}

func deleteDir(answers *internal.Responses) {
	if answers.DeleteDir {
		s := terminal.ShowBusy(fmt.Sprintf("Deleting directory and its contents in %s...", answers.ServicePath()))
		if _, err := os.Stat(answers.ServicePath()); !os.IsNotExist(err) {
			_ = os.RemoveAll(answers.ServicePath())
		}
		s.Stop()
	}
}

func generateFiles(answers *internal.Responses) {
	box := packr.New("defaultService", "./templates/service/default")
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
		_ = t.Execute(w, &answers)
		return w.Flush()
	})
	if err != nil {
		panic(err)
	}
}

func changeFileMode() {
	cmd := exec.Command("chmod", "+x", "generate.sh")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func execGoGenerate() {
	s := terminal.ShowBusy("go generate...")
	cmd := exec.Command("go", "generate")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	s.Stop()
}

func execGoFmt() {
	s := terminal.ShowBusy("go fmt")
	cmd := exec.Command("go", "fmt", "./...")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	s.Stop()
}

func printInstructions(answers *internal.Responses) {
	fmt.Println("\nService generation complete")
	fmt.Println("------------------------------------------------")
	fmt.Printf("cd %s; go run cmd/main.go\n\n", answers.ServicePath())
}

func main() {
	printHeader()
	responses := askQuestions()
	deleteDir(responses)
	generateFiles(responses)
	setWorkingDir(responses)
	changeFileMode()
	execGoGenerate()
	execGoFmt()
	printInstructions(responses)
}
