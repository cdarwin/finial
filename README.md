# finial

A tool to convert `docker-compose.yml` into salt stack pillar data intended for
`docker.compose-ng`.

## Design

The tool accepts a filename option, `-f <filename>` of a valid
`docker-compose.yml` document and a filename option `-o <output>` to be used as
pillar data to the `docker.compose-ng` salt stack formula.

```shell
$ finial -f docker-compose.yml -o compose.sls
```

## Docker Build Image

Use the included Dockerfile to create a go build image

```shell
$ docker build -t finial:build .
```

## Fetch Dependencies

```shell
$ mkdir -p ~/workspace/go_src
$ docker run --rm -v "$PWD":/usr/src/finial -v ~/workspace/go_src:/go/src -w /usr/src/finial finial:build go get gopkg.in/yaml.v2
```

## Build `finial`

I develop and test on OS X but I obviously build in a Linux container

```shell
$ docker run --rm -v "$PWD":/usr/src/finial -v ~/workspace/go_src:/go/src -w /usr/src/finial -e GOOS=darwin -e GOARCH=amd64 finial:build go build -v -o finial-darwin-amd64
```
