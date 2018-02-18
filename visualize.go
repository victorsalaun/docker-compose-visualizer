package main

import (
	"io/ioutil"
	"github.com/awalterschulze/gographviz"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"fmt"
	"strings"
	"os"
)

func visualize(c *cli.Context) {
	var (
		err     error
		graph   *gographviz.Graph
		project string
	)

	data, err := ioutil.ReadFile(c.String("input-file"))
	check(err)

	dc := DockerComposeV3{}
	err = yaml.Unmarshal(data, &dc)

	// Create directed graph
	graph = gographviz.NewGraph()
	graph.SetName(project)
	graph.SetDir(true)

	for name := range dc.Services {
		graph.AddNode(project, nodify(name), map[string]string{
			"label": fmt.Sprintf(name),
			"shape": "component",
		})
	}

	fmt.Print(graph)

	file, err := os.Create(c.String("output-file"))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write([]byte(graph.String()))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func nodify(s string) string {
	return strings.Replace(s, "-", "_", -1)
}
