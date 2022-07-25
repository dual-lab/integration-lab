package main

import (
	"os"

	"github.com/containers/buildah"
	"github.com/containers/storage/pkg/unshare"
	"github.com/dual-lab/integration-lab/containers/commands"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// Replaced at linking time with repo version
var version = "0.0.0"

func init() {
  logrus.SetFormatter(&logrus.TextFormatter{
    FullTimestamp: true,
  })
  logrus.SetLevel(logrus.InfoLevel)

}

func main() {
	if buildah.InitReexec() {
		return
	}
	unshare.MaybeReexecUsingUserNamespace(false)

	var app = &cli.App{
		Name:    "containers",
		Version: version,
		Authors: []*cli.Author{
			{
				Name:  "dmike16",
				Email: "dual-lab@yandex.com",
			},
		},
		Copyright: "Copyright 2022 dmike16. All rights reserved.",
		Usage:     "build image for integration lab",
		Commands:  commands.Commands,
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}

}
