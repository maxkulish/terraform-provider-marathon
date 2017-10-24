package marathon

import (
	"time"

	"github.com/ContainerLabs/go-marathon"
)

type config struct {
	config                   marathon.Config
	Client                   marathon.Marathon
	DefaultDeploymentTimeout time.Duration
}

func (c *config) loadAndValidate() error {
	client, err := marathon.NewClient(c.config)
	c.Client = client
	return err
}
