package main

import (
	"io/ioutil"
	"log"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

type DockerComposeV3 struct {
	Version  string
	Services map[string]Service
	Volumes  map[string]Volume
}

type Service struct {
	Build   string
	Image   string
	Restart string
	Ports   []string
	Volumes []string
	Links   []string
}

type Volume struct {
}

func visualize(c *cli.Context) {
	dat, err := ioutil.ReadFile("docker-compose.yml")
	check(err)

	dc := DockerComposeV3{}
	err = yaml.Unmarshal(dat, &dc)
	log.Printf("Version %s", dc.Version)
	log.Printf("Volumes %s", dc.Services["web"].Volumes)
	log.Printf("Volumes %s", dc.Volumes)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
