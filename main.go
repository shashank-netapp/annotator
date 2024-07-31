package main

import (
	"fmt"
	"github.com/theshashankpal/api-extractor/loader"
	"github.com/theshashankpal/api-extractor/markers"
)

var src = "/Users/shashank/Library/CloudStorage/OneDrive-NetAppInc/Documents/Astra/TRID-POLLING-NEW/trident/storage_drivers/ontap/api/ontap_rest.go"

// Define your marker structure.
type ApiCollector struct {
	OrchestratorFunction []string `marker:"orchestratorFunction,optional"`
	Api                  string   `marker:"api"`
	Method               string   `marker:"method"`
}

func main() {

	def, err := markers.MakeDefinition("apiCollector", markers.DescribesFunc, ApiCollector{})
	if err != nil {
		panic(err)
	}
	// Create a new registry.
	reg := &markers.Registry{}
	err = reg.Register(def)
	if err != nil {
		return
	}

	// Create a new collector.
	col := &markers.Collector{Registry: reg}

	//// Load your Go package.
	pack, err := loader.LoadRoots("/Users/shashank/Library/CloudStorage/OneDrive-NetAppInc/Documents/Astra/TRID-POLLING-NEW/trident/storage_drivers/ontap/api")
	for _, pkg := range pack {
		// Iterate over all types in the package and process markers.
		markers.EachToken(col, pkg, func(info *markers.TypeInfo) {
			// Process markers for each type.
			for _, markerVals := range info.Markers {
				for _, val := range markerVals {
					myMarker, ok := val.(ApiCollector)
					if !ok {
						continue
					}
					fmt.Printf("Type: %s, Marker Field: %s\n", info.Name, myMarker.Api)
				}
			}
		}, func(funcInfo *markers.FuncInfo) {
			for key, markerVals := range funcInfo.Markers {
				for _, val := range markerVals {
					myMarker, ok := val.(ApiCollector)
					if !ok {
						continue
					}
					fmt.Printf("Marker Function Name: %s\n", funcInfo.Name)
					fmt.Printf("Marker Name: %s\n", key)
					fmt.Printf("Marker Field: %s\n", myMarker.Api)
					fmt.Printf("Marker Field: %s\n", myMarker.OrchestratorFunction)
					fmt.Printf("Marker Field: %s\n", myMarker.Method)
				}
			}
		})
	}

}
