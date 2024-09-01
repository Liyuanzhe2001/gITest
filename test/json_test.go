package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONUnmarshal(t *testing.T) {
	headers := "{\"a\":\"a\",\"b\":\"b\"}"
	var headersMap map[string]string
	err := json.Unmarshal([]byte(headers), &headersMap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(headersMap)
}
