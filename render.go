package main

import (
	"fmt"
	"io/ioutil"
	"github.com/awalterschulze/gographviz"
	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
	"os"
	"os/exec"
	"strings"
	"log"
	"bytes"
)

func convertDockerComposeToDots(c *cli.Context) {
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

	if !c.Bool("no-volumes") {
		for volumeKey := range dc.Volumes {
			graph.AddNode(project, nodify(volumeKey), map[string]string{
				"shape": "folder",
			})
		}
	}

	if !c.Bool("no-services") {
		for serviceKey, serviceValue := range dc.Services {
			graph.AddNode(project, nodify(serviceKey), map[string]string{
				"shape": "component",
			})

			if !c.Bool("no-builds") {
				if serviceValue.Build != "" {
					graph.AddNode(project, nodify(serviceValue.Build), map[string]string{
						"shape": "folder",
					})

					edge := gographviz.Edge{
						Dir: true,
						Src: nodify(serviceValue.Build),
						Dst: nodify(serviceKey),
						Attrs: map[gographviz.Attr]string{
							gographviz.Attr("label"): "build",
						},
					}
					graph.Edges.Add(&edge)
				}
			}

			if !c.Bool("no-ports") {
				for portIndex := range serviceValue.Ports {
					graph.AddNode(project, nodify(serviceValue.Ports[portIndex]), map[string]string{
						"shape": "circle",
					})

					edge := gographviz.Edge{
						Dir: true,
						Src: nodify(serviceValue.Ports[portIndex]),
						Dst: nodify(serviceKey),
						Attrs: map[gographviz.Attr]string{
							gographviz.Attr("label"): "port",
						},
					}
					graph.Edges.Add(&edge)
				}
			}

			if !c.Bool("no-volumes") {
				for volumeIndex := range serviceValue.Volumes {
					graph.AddNode(project, nodify(serviceValue.Volumes[volumeIndex]), map[string]string{
						"shape": "folder",
					})

					edge := gographviz.Edge{
						Dir: true,
						Src: nodify(serviceValue.Volumes[volumeIndex]),
						Dst: nodify(serviceKey),
						Attrs: map[gographviz.Attr]string{
							gographviz.Attr("label"): "volume",
						},
					}
					graph.Edges.Add(&edge)
				}
			}

			if !c.Bool("no-links") {
				for linkIndex := range serviceValue.Links {
					edge := gographviz.Edge{
						Dir: true,
						Src: nodify(serviceKey),
						Dst: nodify(serviceValue.Links[linkIndex]),
						Attrs: map[gographviz.Attr]string{
							gographviz.Attr("label"): "link",
						},
					}
					graph.Edges.Add(&edge)
				}
			}

		}
	}

	fmt.Print(graph)

	file, err := os.Create(c.String("output-dot-file"))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write([]byte(graph.String()))

	drawGraphviz(c)
}

func drawGraphviz(c *cli.Context) {
	filesPart := strings.Split(c.String("output-graph-file"), ".")
	cmd := exec.Command("/usr/bin/dot", "-T"+"", filesPart[len(filesPart)-1], c.String("output-dot-file"), "-o", c.String("output-graph-file"))
	cmd.Stdin = strings.NewReader("some input")

	var out bytes.Buffer

	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {

		log.Fatal(err)

	}

	fmt.Printf("in all caps: %q\n", out.String())
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func nodify(s string) string {
	return "\"" + strings.Replace(s, "-", "_", -1) + "\""
}
