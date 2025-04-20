package util

import "encoding/json"

func Marshal(v any) (string, error) {
	var res string

	b, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		return res, err
	}
	res = string(b)
	return res, nil
}
