package main

import (
  "flag"
  "io/ioutil"
  "log"
  "gopkg.in/yaml.v2"
)

var (
  filename = flag.String("f", "docker-compose.yml", "absolute path to docker-compose document")
  document interface{}
)

func main() {
  flag.Parse()

  file, err := ioutil.ReadFile(*filename)
  if err != nil {
    log.Fatal(err)
  }
  err = yaml.Unmarshal(file, &document)
  if err != nil {
    log.Fatal(err)
  }
}
