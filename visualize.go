package main

import (
	"fmt"
	"io/ioutil"
	"github.com/awalterschulze/gographviz"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
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
	graph.SetName("docker_compose")
	graph.SetDir(true)

	for volumeKey := range dc.Volumes {
		graph.AddNode(project, nodify(volumeKey), map[string]string{
			"shape": "folder",
		})
	}

	for serviceKey, serviceValue := range dc.Services {
		graph.AddNode(project, nodify(serviceKey), map[string]string{
			"shape": "component",
		})

		if serviceValue.Build != "" {
			graph.AddNode(project, nodify(serviceValue.Build), map[string]string{
				"shape": "folder",
			})

			edge := gographviz.Edge{}
			edge.Dir = true
			edge.Src = nodify(serviceValue.Build)
			edge.Dst = nodify(serviceKey)
			graph.Edges.Add(&edge)
		}

		for portIndex := range serviceValue.Ports {
			graph.AddNode(project, nodify(serviceValue.Ports[portIndex]), map[string]string{
				"shape": "circle",
			})

			edge := gographviz.Edge{}
			edge.Dir = true
			edge.Src = nodify(serviceValue.Ports[portIndex])
			edge.Dst = nodify(serviceKey)
			graph.Edges.Add(&edge)
		}

		for volumeIndex := range serviceValue.Volumes {
			graph.AddNode(project, nodify(serviceValue.Volumes[volumeIndex]), map[string]string{
				"shape": "folder",
			})

			edge := gographviz.Edge{}
			edge.Dir = true
			edge.Src = nodify(serviceValue.Volumes[volumeIndex])
			edge.Dst = nodify(serviceKey)
			graph.Edges.Add(&edge)
		}

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
	return "\"" + strings.Replace(s, "-", "_", -1) + "\""
}
