FROM golang:1.9.4-alpine

RUN apk update && \
    apk add graphviz

RUN mkdir -p /go/src/github.com/victorsalaun/docker-compose-visualizer
COPY . /go/src/github.com/victorsalaun/docker-compose-visualizer
WORKDIR /go/src/github.com/victorsalaun/docker-compose-visualizer

RUN go build -o main -o render

VOLUME /workdir
WORKDIR /workdir

CMD ["/go/src/github.com/victorsalaun/docker-compose-visualizer/render", "render"]