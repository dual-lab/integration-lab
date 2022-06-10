package main

import (
	"log"
	"os"
	"github.com/urfave/cli/v2"
)

// Replaced at linking time with repo version
var version = "0.0.0"

func init() {
  log.SetPrefix("ILC-CLI")
}

func main() {
	var app = &cli.App{
		Name:  "containers",
    Version: version,
		Usage: "build image for integration lab",
	}

  if err := app.Run(os.Args); err != nil {
    log.Fatal(err)
  }

}
