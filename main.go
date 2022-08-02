package main

import (
	"github.com/hashicorp-dev-advocates/waypoint-client/pkg/client"
	"log"
	"os"
)

var token string

func main() {

  token = os.Getenv("WAYPOINT_TOKEN")
  if token == "" {
    log.Fatal("WAYPOINT_TOKEN environment variable not set")
  }

  // create a client
  conf := client.DefaultConfig()
  conf.Token = token
  conf.Address = "localhost:9701"

  wp, err := client.New(conf)
  if err != nil {
    log.Fatal(err)
  }
}
