package main

import (
	"docker-sdk/client"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	saveCommand := cli.Command{
		Name:        "save",
		Description: "保存镜像",
		Action: func(context *cli.Context) error {
			var err error
			c, err := client.NewDockerClient(context)
			c.Save(context)
			return err
		},
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:     "images",
				Aliases:  []string{"i"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "path",
				Aliases:  []string{"p"},
				Required: true,
			},
		},
	}
	app := &cli.App{
		Name:  "docker-tools",
		Usage: "docker镜像工具",
		Commands: []*cli.Command{
			&saveCommand,
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "host",
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
