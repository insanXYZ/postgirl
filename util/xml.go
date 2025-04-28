package util

import "github.com/clbanning/mxj/v2"

func XmlUnmarshal(data []byte) (mxj.Map, error) {
	return mxj.NewMapXml(data)
}
