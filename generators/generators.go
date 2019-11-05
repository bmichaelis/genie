package generators

import (
	"fmt"
	"genie/generators/service"
	"genie/util"
	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"os"
	"os/exec"
)

const UserCancelSurveyMessage = "Hey where did you go? I am so lonely now 😪"
const UserCancelProjectMessage = "👍 Okay...better luck next time!"

var terminal = util.NewTerminal()

type Generatorer interface {
	GetName() string
	AskQuestions() error
	Execute()
}

type Generatorser interface {
	Add(generator Generatorer) Generatorer
	Find(name string) Generatorer
	Run()
}

type Generators struct {
	Store map[string]Generatorer
	Order []Generatorer
}

func (g *Generators) Add(generator Generatorer) Generatorer {
	g.Order = append(g.Order, generator)
	g.Store[generator.GetName()] = generator
	return generator
}

func (g *Generators) Find(name string) Generatorer {
	return g.Store[name]
}

func (g *Generators) Run() {
	g.printHeader()

	var err error
	for _, gen := range g.Order {
		if err = gen.AskQuestions(); err != nil {
			fmt.Println(UserCancelSurveyMessage)
			break
		}
	}

	if err == nil {
		fmt.Println("")
		s := g.Find(service.NAME).(*service.Generator)

		var execute bool
		err := survey.AskOne(&survey.Confirm{
			Message: fmt.Sprintf("Your service will be created at: %s, continue?", s.Responses.ServicePath()),
			Default: true,
		}, &execute)

		if execute {
			for _, gen := range g.Order {
				gen.Execute()
			}

			g.setWorkingDir()
			g.changeFileMode()
			// g.execGoModule111()
			// g.execGoModVendor()
			g.execGoGenerate()
			g.execGoFmt()
			g.printInstructions()
		} else if err != nil {
			fmt.Println(UserCancelSurveyMessage)
		} else {
			fmt.Println(UserCancelProjectMessage)
		}
	}
}

func (*Generators) printHeader() {
	header := util.NewHeader()
	header.Print()
}

func (g *Generators) setWorkingDir() {
	s := g.Find(service.NAME).(*service.Generator)
	path := s.Responses.ServicePath()
	if err := os.Chdir(path); err != nil {
		panic(err)
	}
}

func (*Generators) changeFileMode() {
	cmd := exec.Command("chmod", "+x", "generate.sh")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func (*Generators) execGoModule111() {
	s := terminal.ShowBusy("export GOMODULE111=on...")
	cmd := exec.Command("export", "GOMODULE111=on")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	s.Stop()
}

func (*Generators) execGoModVendor() {
	s := terminal.ShowBusy("go mod vendor...")
	cmd := exec.Command("go", "mod", "vendor")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	s.Stop()
}

func (*Generators) execGoGenerate() {
	s := terminal.ShowBusy("go generate...")
	cmd := exec.Command("go", "generate")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	s.Stop()
}

func (*Generators) execGoFmt() {
	s := terminal.ShowBusy("go fmt")
	cmd := exec.Command("go", "fmt", "./...")
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	s.Stop()
}

func (g *Generators) printInstructions() {
	s := g.Find(service.NAME).(*service.Generator)
	fmt.Println("\n🧞 Service generation complete")
	fmt.Println("------------------------------------------------")
	fmt.Println("In terminal #1, to run the server...")
	color.HiGreen(fmt.Sprintf("cd %s\n", s.Responses.ServicePath()))
	color.HiGreen("export GOMODULE111=on\n")
	color.HiGreen("go mod vendor\n")
	color.HiGreen("go run cmd/main.go\n\n")
	fmt.Println("In terminal #2, to run the client...")
	color.HiGreen("go run test/main.go\n\n")
}

var instance = &Generators{
	Store: map[string]Generatorer{},
}

func GetInstance() Generatorser {
	return instance
}
