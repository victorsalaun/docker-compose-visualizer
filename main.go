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
			Name:   "render",
			Action: convertDockerComposeToDots,
			Usage:  "Visualize docker-compose file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "input-file, i",
					Usage: "Path to a docker-compose input file [default: \"./docker-compose.yml\"]",
					Value: "./docker-compose.yml",
				},
				cli.StringFlag{
					Name:  "output-dot-file, d",
					Usage: "Path to a dot output file [default: \"./docker-compose.dot\"]",
					Value: "./docker-compose.dot",
				},
				cli.StringFlag{
					Name:  "output-graph-file, g",
					Usage: "Path to a graph output file [default: \"./docker-compose.svg\"]",
					Value: "./docker-compose.svg",
				},
				// disabling flags
				cli.BoolFlag{
					Name:  "no-builds",
					Usage: "Disable displaying builds",
				},
				cli.BoolFlag{
					Name:  "no-links",
					Usage: "Disable displaying links",
				},
				cli.BoolFlag{
					Name:  "no-ports",
					Usage: "Disable displaying ports",
				},
				cli.BoolFlag{
					Name:  "no-services",
					Usage: "Disable displaying services",
				},
				cli.BoolFlag{
					Name:  "no-volumes",
					Usage: "Disable displaying volumes",
				},
			},
		},
	}

	app.Run(os.Args)
}

// DockerComposeV3 struct represents the docker-compose file structure, https://docs.docker.com/compose/compose-file/#compose-file-structure-and-examples
type DockerComposeV3 struct {
	Version  string             `yaml:"version"`
	Services map[string]Service `yaml:"services"`
	Volumes  map[string]Volume  `yaml:"volumes"`
	Networks map[string]Network `yaml:"networks"`
}

// Service struct is a first level structure of DockerComposeV3, https://docs.docker.com/compose/compose-file/#service-configuration-reference
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

// Volume struct is a first level structure of DockerComposeV3, https://docs.docker.com/compose/compose-file/#volumes
type Volume struct {
}

// Network struct is a first level structure of DockerComposeV3, https://docs.docker.com/compose/compose-file/#network-configuration-reference
type Network struct {
}
