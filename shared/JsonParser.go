package shared

import (
	"encoding/json"
	"io"
)

type JsonParser struct {}

func (_ JsonParser) ParseStruct(payload interface{}) []byte {
	reqJson, err := json.Marshal(payload)

	if err != nil {
		panic(err)
	}

	return reqJson
}

func (_ JsonParser) ParseBody(body io.ReadCloser, req interface{}) {
	defer body.Close()

	if err := json.NewDecoder(body).Decode(&req); err != nil {
		logger := GetLogger()
		logger.Log(err.Error())
	}

	PrintStruct(req, "parsed body: %v")
}
