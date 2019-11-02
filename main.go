package main

import (
	"genie/generators"
	"genie/generators/charts"
	"genie/generators/mongo"
	"genie/generators/service"
)

func main() {
	g := generators.GetInstance()
	g.Add(service.NewGenerator())
	g.Add(charts.NewGenerator())
	g.Add(mongo.NewGenerator())
	g.Run()
}
