package shared

import (
	"encoding/json"
	"fmt"
)

func PrintStruct(payload interface{}, message string) {
	jsonState, _ := json.Marshal(payload)

	logger := GetLogger()
	logger.Log(fmt.Sprintf(message, string(jsonState)))
}


