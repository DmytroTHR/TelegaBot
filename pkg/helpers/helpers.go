package helpers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"
)

const DefaultContentType = "application/json"

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

func Logger() *logrus.Logger {
	return log
}

func ReadFromRequest(addr *url.URL, httpMethod, body string) ([]byte, error) {
	var res *http.Response
	var err error

	switch httpMethod {
	case http.MethodGet:
		res, err = http.Get(addr.String())
	case http.MethodPost:
		res, err = http.Post(addr.String(), DefaultContentType, strings.NewReader(body))
	default:
		return nil, fmt.Errorf("%s method is not allowed on %s", httpMethod, addr.String())
	}

	if err != nil {
		if len(body) > 0 {
			log.Warnln("BODY:\t", body)
		}

		return nil, WrapError(fmt.Sprintf("http %s on %s", httpMethod, addr.String()), err)
	}

	return unmarshalledResponse(res.Body)
}

func unmarshalledResponse(response io.ReadCloser) ([]byte, error) {
	defer response.Close()

	buf := bytes.Buffer{}
	_, err := io.Copy(&buf, response)
	if err != nil {
		return nil, WrapError("read response data", err)
	}

	return buf.Bytes(), nil
}

func WrapError(withMessage string, err error) error {
	log.Errorf("%s: %v\n", withMessage, err)

	return fmt.Errorf("%s: %w", withMessage, err)
}

func Error(message string) error {
	log.Errorf(message)

	return fmt.Errorf(message)
}
