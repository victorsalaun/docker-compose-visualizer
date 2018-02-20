# Docker-compose Visualizer

## Install

    go get github.com/victorsalaun/docker-compose-visualizer

## Usage

    Usage:
      visualize [options]

    Options:
      --input-file, --i                  Path to a docker-compose input file [default: "./docker-compose.yml"]
      --output-file, --o                 Path to a dot output file [default: "./docker-compose.dot"]

    Flags:
      --no-builds                        Disable displaying builds
      --no-links                         Disable displaying links
      --no-ports                         Disable displaying ports
      --no-services                      Disable displaying services
      --no-volumes                       Disable displaying volumes

Execute GraphViz 

    execute_dot.sh svg .\examples\docker-compose

## Dev

Install vendor package:

    govendor fetch github.com/urfave/cli