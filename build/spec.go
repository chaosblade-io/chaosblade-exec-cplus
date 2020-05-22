package main

import (
	"log"
	"os"

	"github.com/chaosblade-io/chaosblade-spec-go/spec"
	"github.com/chaosblade-io/chaosblade-spec-go/util"

	"github.com/chaosblade-io/chaosblade-exec-cplus/module"
)

func main() {
	if len(os.Args) != 2 {
		log.Panicln("less yaml file path")
	}
	err := util.CreateYamlFile(getModels(), os.Args[1])
	if err != nil {
		log.Panicf("create yaml file error, %v", err)
	}
}

// getModels returns experiment models in the project
func getModels() *spec.Models {
	modelCommandSpecs := []spec.ExpModelCommandSpec{
		module.NewCPlusCommandModelSpec(),
	}
	prepareModel := spec.ExpPrepareModel{
		PrepareType: "cplus",
		PrepareFlags: []spec.ExpFlag{
			{
				Name: "port",
				Desc: "server port to be listening",
			},
			{
				Name: "ip",
				Desc: "The ip address bound to the service",
			},
		},
		PrepareRequired: true,
	}
	specModels := make([]*spec.Models, 0)
	for _, modeSpec := range modelCommandSpecs {
		specModel := util.ConvertSpecToModels(modeSpec, prepareModel, "host")
		specModels = append(specModels, specModel)
	}
	return util.MergeModels(specModels...)
}
