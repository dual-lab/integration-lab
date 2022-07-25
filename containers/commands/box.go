package commands

import (
	"github.com/dual-lab/integration-lab/containers/common"
	"github.com/dual-lab/integration-lab/containers/images"
	"github.com/urfave/cli/v2"
)


var imageType string

var box *cli.Command = &cli.Command{
  Name: "box",
  Aliases: []string{"b"},
  Usage: "Box a particular image selected with the image flag",
  Flags: []cli.Flag{
    &cli.StringFlag{
      Name: "image",
      Required: true,
      Usage: "supported image to build",
      Destination: &imageType,
    },
  },
  Action: func(ctx *cli.Context) error {
    switch imageType {
    case "helloworld":
        return boxImage(images.NewHelloworldBuilder())
    }
    return cli.Exit("Image type not supported", 55)
  },
}

func boxImage(builderFile common.BuilderFile) error {
  err := common.BuildMain(builderFile)

  if err != nil {
    return cli.Exit("Error on boxing the image", 43)
  } else {
    return nil
  }
}
