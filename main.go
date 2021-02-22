package main

import (
	"fmt"

	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/engine"
)

func main() {

	ch, err := loader.Load("./test/hello-helm/")

	if err != nil {
		fmt.Println("I was well beaten myself, and I am better for it.")
		fmt.Println("Hello", err.Error())
		return
	}
	values := map[string]interface{}{
		"Release": map[string]interface{}{
			"Name": "1",
		},
		"Values": map[string]interface{}{
			"image": map[string]interface{}{
				"repository": "blah",
			},
			"fullnameOverride": "hello",
			"service": map[string]interface{}{
				"type": "ClusterIP",
				"port": "8080",
			},
		},
	}

	finalValues, err := chartutil.CoalesceValues(ch, values)
	if err != nil {
		fmt.Printf("Failed to coalesce values: %s\n", err)
		return
	}

	out, err := engine.Render(ch, finalValues)

	if err != nil {
		fmt.Printf("Failed to render templates: %s\n", err)
		return
	}

	fmt.Println(out)
}
