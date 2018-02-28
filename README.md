# Docker-compose Visualizer

## Run

Run docker image with default values giving workdir

    docker run -it -v $(pwd):/workdir victorsalaun/docker-compose-visualizer
    
## Install

    go get github.com/victorsalaun/docker-compose-visualizer

## Usage

    Usage:
      render [options]

    Options:
      --input-file, --i                  Path to a docker-compose input file [default: "./docker-compose.yml"]
      --output-dot-file, --o             Path to a dot output file [default: "./docker-compose.dot"]
      --output-graph-file, --o           Path to a dot output file [default: "./docker-compose.png"]

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

Build docker image

    docker build -t victorsalaun/docker-compose-visualizer .
    
Run docker image

    docker run -it -v $(pwd):/workdir victorsalaun/docker-compose-visualizer
