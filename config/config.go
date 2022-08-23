package config

import (
	"fmt"
	"net/url"

	"github.com/DmytroTHR/telegabot/pkg/helpers"
)

var log = helpers.Logger()

type Config struct {
	apiToken       string
	apiHost        string
	UpdateTimeout  int
	UpdateMsgLimit int
}

func NewConfig(host, token string, timeout, limit int) (*Config, error) {
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
	if timeout == 0 {
		timeout = 10
		log.Warnln("Set default value 10s for Update timeout. Check if env TELEGA_TIMEOUT is present.")
	}
	if limit == 0 {
		limit = 10
		log.Warnln("Set default value 10 for Update message limit. Check if env TELEGA_MSG_LIMIT is present.")
	}

	return &Config{
		apiToken:       token,
		apiHost:        hostURL.String(),
		UpdateTimeout:  timeout,
		UpdateMsgLimit: limit,
	}, nil
}

func (conf *Config) FullAPIPath(method string) (*url.URL, error) {
	log.Debugln("Prepare to make request on:", method)
	return url.Parse(fmt.Sprintf("%sbot%s/%s", conf.apiHost, conf.apiToken, method))
}

func (conf *Config) FullAPIFilePath(pathToFile string) (*url.URL, error) {
	log.Debugln("Prepare to make request on:", pathToFile)
	return url.Parse(fmt.Sprintf("%sfile/bot%s/%s", conf.apiHost, conf.apiToken, pathToFile))
}
