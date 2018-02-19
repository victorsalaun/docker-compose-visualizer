# Docker-compose Visualizer

## Install

    go get github.com/victorsalaun/docker-compose-visualizer

## Usage

    Usage:
      visualize [options]

    Options:
      --input-file, --i                  Path to a docker-compose input file [default: "./docker-compose.yml"]
      --output-file, --o                 Path to a dot output file [default: "./docker-compose.dot"]


Execute GraphViz 

    execute_dot.sh svg .\examples\docker-compose

## Dev

Install vendor package:

    govendor fetch github.com/urfave/cli