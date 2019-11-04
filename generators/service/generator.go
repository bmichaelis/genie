package service

import (
	"bufio"
	"fmt"
	"genie/util"
	"os"
	"os/exec"
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

func (g *Generator) Run() {
	g.printHeader()
	g.Responses = NewSurvey().Start()
	g.deleteDir()
	g.writeFiles()
}

func (g *Generator) Finalize() {
	g.setWorkingDir()
	g.changeFileMode()
	// g.execGoModule111()
	// g.execGoModVendor()
	g.execGoGenerate()
	g.execGoFmt()
	g.printInstructions()
}

func (*Generator) printHeader() {
	header := util.NewHeader()
	header.Print()
}

func (g *Generator) setWorkingDir() {
	responses := g.Responses
	path := responses.ServicePath()
	if err := os.Chdir(path); err != nil {
		panic(err)
	}
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

func (*Generator) changeFileMode() {
	cmd := exec.Command("chmod", "+x", "generate.sh")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func (*Generator) execGoModule111() {
	s := terminal.ShowBusy("export GOMODULE111=on...")
	cmd := exec.Command("export", "GOMODULE111=on")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	s.Stop()
}

func (*Generator) execGoModVendor() {
	s := terminal.ShowBusy("go mod vendor...")
	cmd := exec.Command("go", "mod", "vendor")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	s.Stop()
}

func (*Generator) execGoGenerate() {
	s := terminal.ShowBusy("go generate...")
	cmd := exec.Command("go", "generate")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	s.Stop()
}

func (*Generator) execGoFmt() {
	s := terminal.ShowBusy("go fmt")
	cmd := exec.Command("go", "fmt", "./...")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	s.Stop()
}

func (g *Generator) printInstructions() {
	responses := g.Responses
	fmt.Println("\nService generation complete")
	fmt.Println("------------------------------------------------")
	fmt.Println("In terminal #1, to run the server...")
	fmt.Printf("cd %s\n", responses.ServicePath())
	fmt.Printf("export GOMODULE111=on\n")
	fmt.Printf("go mod vendor\n")
	fmt.Printf("go run cmd/main.go\n\n")
	fmt.Println("In terminal #2, to run the client...")
	fmt.Printf("go run test/main.go\n\n")
}

func NewGenerator() *Generator {
	return &Generator{}
}
