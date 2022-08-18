package config

import (
	"fmt"
	"net/url"

	"github.com/DmytroTHR/telegabot/pkg/helpers"
)

var log = helpers.Logger()

type Config struct {
	apiToken string
	apiHost  string
}

func NewConfig(host, token string) (*Config, error) {
	if len(token) == 0 {
		return nil, helpers.Error("empty API token provided")
	}
	if len(host) == 0 {
		return nil, helpers.Error("empty API host provided")
	}
	hostURL, err := url.Parse(host)
	if err != nil {
		return nil, helpers.WrapError(fmt.Sprintf("unable to parse host %s", host), err)
	}

	return &Config{
		apiToken: token,
		apiHost:  hostURL.String(),
	}, nil
}

func (conf *Config) FullAPIPath(method string) (*url.URL, error) {
	log.Println("Prepare to make request on:", method)
	return url.Parse(fmt.Sprintf("%sbot%s/%s", conf.apiHost, conf.apiToken, method))
}
