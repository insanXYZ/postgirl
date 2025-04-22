package util

import "encoding/json"

func JsonMarshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func JsonMarshalString(v any) (string, error) {
	b, err := json.MarshalIndent(v, "", " ")
	return string(b), err
}

func JsonUnmarshal(data []byte, dst any) error {
	return json.Unmarshal(data, dst)
}
