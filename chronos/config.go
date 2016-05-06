package chronos

import (
	"github.com/behance/go-chronos/chronos"
	"time"
)

type config struct {
	URL                      string
	RequestTimeout           int
	DefaultDeploymentTimeout time.Duration
	Debug                    bool

	Client chronos.Chronos
}

func (c *config) loadAndValidate() error {

	config := chronos.Config{
		URL:            c.URL,
		Debug:          c.Debug,
		RequestTimeout: c.RequestTimeout,
	}

	client, err := chronos.NewClient(config)
	c.Client = client
	return err
}
