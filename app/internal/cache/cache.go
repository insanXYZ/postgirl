package cache

import "postgirl/app/model"

var CacheRequests *cacheRequests

func init() {
	CacheRequests = NewCacheRequest()
}

type cacheRequests struct {
	caches map[string]model.Request
}

func NewCacheRequest() *cacheRequests {
	return &cacheRequests{
		caches: make(map[string]model.Request),
	}
}

func (c *cacheRequests) create(index string) {
	if _, ok := c.caches[index]; !ok {
		r := model.Request{}
		r.Attribute.Params = make(map[string]string)
		r.Attribute.Headers = make(map[string]string)

		c.caches[index] = r
	}
}
