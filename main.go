package main

import (
	"os"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "docker-compose-visualizer"
	app.EnableBashCompletion = true

	app.Commands = []cli.Command{
		{
			Name:   "export",
			Action: export,
			Usage:  "Export docker-compose file as draw.io file",
		},
		{
			Name:   "visualize",
			Action: visualize,
			Usage:  "Visualize docker-compose file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "input-file, i",
					Usage: "Path to a docker-compose input file [default: \"./docker-compose.yml\"]",
					Value: "./docker-compose.yml",
				},
				cli.StringFlag{
					Name:  "output-file, o",
					Usage: "Path to a dot output file [default: \"./docker-compose.dot\"]",
					Value: "./docker-compose.dot",
				},
			},
		},
	}

	app.Run(os.Args)
}

type DockerComposeV3 struct {
	Version  string             `yaml:"version"`
	Services map[string]Service `yaml:"services"`
	Volumes  map[string]Volume  `yaml:"volumes"`
	Networks map[string]Network `yaml:"networks"`
}

type Service struct {
	Build         string   `yaml:"build"`
	CapAdd        []string `yaml:"cap_add"`
	Command       []string `yaml:"command"`
	ContainerName string   `yaml:"container_name"`
	DependsOn     []string `yaml:"depends_on"`
	Image         string   `yaml:"image"`
	Links         []string `yaml:"links"`
	Networks      []string `yaml:"networks"`
	Ports         []string `yaml:"ports"`
	Restart       string   `yaml:"restart"`
	Volumes       []string `yaml:"volumes"`
	VolumesFrom   []string `yaml:"volumes_from"`
}

type Volume struct {
}

type Network struct {
}
