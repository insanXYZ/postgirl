package util

import "encoding/xml"

func XmlUnmarshal(data []byte, dst any) error {
	return xml.Unmarshal(data, &dst)
}
