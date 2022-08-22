package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"time"

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

func MultipartDataUpload(inputFiles, otherOptions map[string]string) (*bytes.Buffer, string, error) {
	if len(inputFiles) == 0 {
		return nil, "", Error("no files provided")
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for key, val := range otherOptions {
		err := writer.WriteField(key, val)
		if err != nil {
			return nil, "", WrapError("set mime fields", err)
		}
	}

	for name, filePath := range inputFiles {
		if len(filePath) == 0 {
			continue
		}
		file, err := os.Open(filePath)
		if err != nil {
			return nil, "", WrapError(fmt.Sprintf("open file for upload %s", filePath), err)
		}

		part, err := writer.CreateFormFile(name, filePath)
		if err != nil {
			file.Close()
			return nil, "", WrapError(fmt.Sprintf("create form file %s", filePath), err)
		}
		n, err := io.Copy(part, file)
		file.Close()
		writer.Close()
		if err != nil {
			return nil, "", WrapError(fmt.Sprintf("copy file to part %s", filePath), err)
		}
		if n == 0 {
			return nil, "", Error(fmt.Sprintf("no data copied from file %s", filePath))
		}
	}

	return body, writer.FormDataContentType(), nil
}

func UnmarshalToKeyValueString(body []byte) map[string]string {
	result := map[string]string{}
	conv := map[string]json.RawMessage{}
	err := json.Unmarshal(body, &conv)
	if err != nil {
		return result
	}
	for k := range conv {
		result[k] = string(conv[k])
	}

	return result
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

func RandomBetween(start, end int) int {
	if start == end {
		return start
	}
	rand.Seed(time.Now().Unix())
	min, max := start, end
	if start > end {
		min, max = end, start
	}
	return rand.Intn(max-min+1) + min
}
