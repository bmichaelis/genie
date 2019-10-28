package main

import (
	"genie/internal/generators"
	"genie/internal/generators/charts"
	"genie/internal/generators/service"
)

func main() {
	g := generators.GetInstance()
	g.Add(service.NewGenerator())
	g.Add(charts.NewGenerator())
	g.Run()
}
