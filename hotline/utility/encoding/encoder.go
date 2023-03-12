package encoding

import "encoding/json"

func JsonToStruct(jsonString string, structOutput interface{}) error {
	return json.Unmarshal([]byte(jsonString), structOutput)
}
