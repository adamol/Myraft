package shared

import (
	"net/http"
	"fmt"
	"bytes"
)

type HttpClient struct {
	JsonParser
}

func (c HttpClient) SendRequestRpc(method string, url string, reqJson []byte) *http.Response {
	logger := GetLogger()
	logger.Log(fmt.Sprintf("sending request to %v", url))
	logger.Log(fmt.Sprintf("with payload %v", string(reqJson)))

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqJson))

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	r, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	return r
}
