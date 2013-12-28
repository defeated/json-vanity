package main

import (
	"bytes"
	"encoding/json"
)

func JsonPrettyPrint(input string) string {
	var buf bytes.Buffer
	json.Indent(&buf, []byte(input), "", " ")
	return buf.String()
}
