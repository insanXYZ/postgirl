package util

import (
	"github.com/clbanning/mxj/v2"
)

func XmlUnmarshal(v []byte) (mxj.Map, error) {
	return mxj.NewMapXml(v)
}
