package main

import (
	"docker-tools/client"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	saveCommand := cli.Command{
		Name:        "save",
		Description: "保存镜像",
		Action: func(cli *cli.Context) error {
			c, err := client.NewDockerClient(cli)
			if err != nil {
				return err
			}
			err = c.Save(cli)
			if err != nil {
				return err
			}
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:     "images",
				Aliases:  []string{"i"},
				Required: true,
				Usage:    "镜像版本号或者版本号文件",
			},
			&cli.StringFlag{
				Name:        "path",
				Aliases:     []string{"p"},
				DefaultText: "./",
				Usage:       "镜像包保存位置",
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
				Name:  "host",
				Usage: "镜像中心地址host。默认为本地连接，远程连接:tcp://10.231.50.28:2375",
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
