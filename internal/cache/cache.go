package cache

import (
	"postgirl/model"
	"slices"
)

var (
	CacheRequests *cacheRequests
)

func init() {
	CacheRequests = newCacheRequest()
}

type cacheRequests struct {
	caches map[string]*model.Request
	listLabel []string
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
		c.listLabel = append(c.listLabel, label)
	}
}

func (c *cacheRequests) GetRequest(label string) *model.Request {
	return c.caches[label]
}

func (c *cacheRequests) GetList() []string {
	return c.listLabel
}

func (c *cacheRequests) DeleteMap(label string) {
	delete(c.caches, label)
}

func (c *cacheRequests) DeleteList(index int){
	c.listLabel = slices.Delete(c.listLabel , index ,index+1)
}

func (c *cacheRequests) GetMap() map[string]*model.Request {
	return c.caches
}
