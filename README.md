# finial

A tool to convert `docker-compose.yml` into salt stach pillar data intended for
`docker.compose-ng`.

## Design

The tool accepts a filename option, `-f <filename>` of a a valid
`docker-compose.yml` document and a filename option `-o <output>` to be used as
pillar data to the `docker.compose-ng` salt stack formula.

```shell
$ finial -f docker-compose.yml -o compose.sls
```
