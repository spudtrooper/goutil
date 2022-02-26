package request

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/check"
)

func MustFormatString(x interface{}) string {
	if x == nil {
		return "nil"
	}
	b, err := json.Marshal(x)
	check.Err(err)
	res, err := PrettyPrintJSON(b)
	check.Err(err)
	return res
}

func MustPrettyPrintJSON(b []byte) string {
	res, err := PrettyPrintJSON(b)
	check.Err(err)
	return res
}

func PrettyPrintJSON(b []byte) (string, error) {
	b = []byte(strings.TrimSpace(string(b)))
	if len(b) == 0 {
		return "", nil
	}
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, b, "", "\t"); err != nil {
		return "", errors.Errorf("json.Indent: payload=%q: %v", string(b), err)
	}
	return prettyJSON.String(), nil
}

// https://stackoverflow.com/questions/28595664/how-to-stop-json-marshal-from-escaping-and/28596225
func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(t); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
