package cache

import (
	"postgirl/app/model"
)

var (
	CacheRequests *cacheRequests
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
		r := model.Request{
			Attribute: new(model.Attribute),
		}
		r.Method = model.GET
		r.Attribute.Params = make(map[string][]string)
		r.Attribute.Headers = map[string]string{
			"User-Agent":      "Postgirl/v1",
			"Accept":          "*/*",
			"Accept-Encoding": "gzip,deflate,br",
			"Connection":      "keep-alive",
		}

		c.caches[label] = &r
	}
}

func (c *cacheRequests) Get(label string) *model.Request {
	return c.caches[label]
}
