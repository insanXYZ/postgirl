package cache

import (
	"postgirl/app/model"
)

var (
	CacheRequests *cacheRequests

	//model.request values
	DefaultParams  = make(map[string][]string)
	DefaultHeaders = map[string]string{
		"User-Agent":      "Postgirl/v1",
		"Accept":          "*/*",
		"Accept-Encoding": "gzip,deflate,br",
		"Connection":      "keep-alive",
	}
)

func init() {
	CacheRequests = newCacheRequest()
}

type cacheRequests struct {
	caches map[string]*model.Request
}

func newCacheRequest() *cacheRequests {
	return &cacheRequests{
		caches: make(map[string]*model.Request),
	}
}

func (c *cacheRequests) Create(label string) {
	if _, ok := c.caches[label]; !ok {
		r := model.Request{}
		r.Method = model.GET
		r.Attribute.Params = DefaultParams
		r.Attribute.Headers = DefaultHeaders

		c.caches[label] = &r
	}
}

func (c *cacheRequests) Get(label string) *model.Request {
	return c.caches[label]
}
