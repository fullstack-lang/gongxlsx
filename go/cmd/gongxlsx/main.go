package main

import (
	"flag"
	"log"

	gongxlsx_go "github.com/fullstack-lang/gongxlsx/go"
	gongxlsx_fullstack "github.com/fullstack-lang/gongxlsx/go/fullstack"
	gongxlsx_models "github.com/fullstack-lang/gongxlsx/go/models"
	gongxlsx_probe "github.com/fullstack-lang/gongxlsx/go/probe"
	gongxlsx_static "github.com/fullstack-lang/gongxlsx/go/static"

	gongdoc_load "github.com/fullstack-lang/gongdoc/go/load"
)

var (
	logDBFlag  = flag.Bool("logDB", false, "log mode for db")
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	compareFlag = flag.String("compare", "sampleFile", "compare to the other file")

	diagrams         = flag.Bool("diagrams", true, "parse/analysis go/models and go/diagrams")
	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")
)

func main() {

	log.SetPrefix("gongxlsx: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := gongxlsx_static.ServeStaticFiles(*logGINFlag)

	// setup stack
	stage, backRepo := gongxlsx_fullstack.NewStackInstance(r, "gongxlsx")

	nbArgs := flag.Args()
	if len(nbArgs) < 1 {
		log.Panic("there should be at least one file argument")
	}

	sampleXLFile := new(gongxlsx_models.XLFile).Stage(stage)
	sampleXLFile.Open(stage, flag.Arg(0))

	if len(nbArgs) > 1 {
		sampleXLFile2 := new(gongxlsx_models.XLFile).Stage(stage)
		sampleXLFile2.Open(stage, flag.Arg(1))
	}

	if *compareFlag == "sampleFile" {
		log.Println("no file to compare")
	} else {
		fileToCompare := new(gongxlsx_models.XLFile).Stage(stage)
		fileToCompare.Open(stage, *compareFlag)

	}

	stage.Commit()

	gongxlsx_probe.NewProbe(r, gongxlsx_go.GoModelsDir, "gongxlsx", stage, backRepo)

	gongdoc_load.Load(
		"gongxlsx",
		"github.com/fullstack-lang/gongxlsx/go/models",
		gongxlsx_go.GoModelsDir,
		gongxlsx_go.GoDiagramsDir,
		r,
		*embeddedDiagrams,
		&stage.Map_GongStructName_InstancesNb)

	log.Printf("Server ready serve on localhost:8080")
	r.Run()
}
