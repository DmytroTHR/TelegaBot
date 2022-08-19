package helpers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

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

func DefaultHeader() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}

func ReadFromRequest(request *http.Request) ([]byte, error) {
	log.Debugln("making request on API")
	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return nil, WrapError("receive response", err)
	}
	defer res.Body.Close()

	if err != nil {
		var buf []byte
		n, _ := request.Body.Read(buf)
		if n > 0 {
			log.Warnln("BODY:\t", string(buf))
		}

		return nil, WrapError(fmt.Sprintf("http %s", request.Method), err)
	}

	return getResponseData(res.Body)
}

func getResponseData(response io.ReadCloser) ([]byte, error) {
	buf := bytes.Buffer{}
	_, err := io.Copy(&buf, response)
	if err != nil {
		return nil, WrapError("read response data", err)
	}

	return buf.Bytes(), nil
}

func WrapError(withMessage string, err error) error {
	log.Errorf("%s: %s\n", withMessage, err)

	return fmt.Errorf("%s: %w", withMessage, err)
}

func Error(message string) error {
	log.Errorf(message)

	return fmt.Errorf(message)
}
