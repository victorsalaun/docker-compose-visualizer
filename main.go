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
		},
	}

	app.Run(os.Args)
}
