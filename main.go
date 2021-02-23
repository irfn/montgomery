package main

import (
	"fmt"
	"os"
	"strings"

	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/engine"

	"github.com/go-git/go-git/v5"
)

func loadchart(repo string) (*chart.Chart, error) {
	urlParts := strings.Split(repo, "/")
	tmpdir := fmt.Sprintf("/tmp/%s", urlParts[len(urlParts)-1])
	os.RemoveAll(tmpdir)
	_, err := git.PlainClone(tmpdir, false, &git.CloneOptions{
		URL:      repo,
		Progress: os.Stdout,
	})
	if err != nil {
		return nil, err
	}
	return loader.Load(tmpdir)
}

func render(ch *chart.Chart, values map[string]interface{}) (map[string]string, error) {
	finalValues, err := chartutil.CoalesceValues(ch, values)
	if err != nil {
		return nil, err
	}
	return engine.Render(ch, finalValues)
}

func main() {
	ch, err := loadchart("https://github.com/irfn/hello-helm")

	if err != nil {
		fmt.Println("I was well beaten myself, and I am better for it.")
		fmt.Println("Hello", err.Error())
		return
	}
	//Following is specific to a Chart.
	//Multiple Charts may be supported via convention
	//eg: Base values should always be in a particular manner.
	//additional values may be overridden via application specific configuration.
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
				"port": "8081",
			},
		},
	}

	out, err := render(ch, values)
	if err != nil {
		fmt.Printf("Failed to render templates: %s\n", err)
		return
	}

	fmt.Println(out["hello/templates/deployment.yaml"])
	fmt.Println(out["hello/templates/service.yaml"])
}
