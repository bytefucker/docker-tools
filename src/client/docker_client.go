package client

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"io"
	"log"
	"os"
	"path"
)

type DockerClient struct {
	client *client.Client
}

func NewDockerClient(cli *cli.Context) (*DockerClient, error) {
	var c *client.Client
	var err error
	host := cli.String("host")
	if host == "" {
		c, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	} else {
		c, err = client.NewClientWithOpts(client.WithHost("tcp://10.231.50.28:2375"))
	}
	if err != nil {
		log.Panic("connect docker error ", err)
	}
	return &DockerClient{
		client: c,
	}, err
}

func (c *DockerClient) Save(cli *cli.Context) (err error) {
	images := cli.StringSlice("images")
	filePath := cli.String("path")
	for _, image := range images {
		c.pull(image)
	}
	ctx := context.Background()
	reader, err := c.client.ImageSave(ctx, images)
	defer reader.Close()
	if err != nil {
		log.Println("read image error", err)
		return err
	}
	file, err := os.Create(path.Join(filePath, "image.tar.gz"))
	defer file.Close()
	if err != nil {
		log.Println("create file error", err)
		return err
	}
	return err
}

//拉取镜像
func (c *DockerClient) pull(image string) {
	ctx := context.Background()
	reader, err := c.client.ImagePull(ctx, image, types.ImagePullOptions{})
	defer reader.Close()
	if err != nil {
		logrus.Warnf("pull image %s failed ", image)
		return
	}
	io.Copy(os.Stdout, reader)
}
