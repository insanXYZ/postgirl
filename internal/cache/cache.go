package cache

import (
	"slices"

	"postgirl/model"
	"postgirl/util"

	"github.com/rivo/tview"
)

var CacheRequests *cacheRequests

func init() {
	CacheRequests = newCacheRequest()
}

type fields struct {
	panel   *tview.Flex    `json:"-"`
	request *model.Request `json:"request,omitempty"`
}

type cacheRequests struct {
	caches    map[string]*fields
	listLabel []string
}

func newCacheRequest() *cacheRequests {
	return &cacheRequests{
		caches: make(map[string]*fields),
	}
}

func (c *cacheRequests) Create(label string, createPanelFunc func(*model.Request) *tview.Flex) {
	if _, ok := c.caches[label]; !ok {
		r := &model.Request{
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
		c.caches[label] = new(fields)
		c.caches[label].request = r
		c.caches[label].panel = createPanelFunc(r)
		c.listLabel = append(c.listLabel, label)
	}
}

func (c *cacheRequests) GetRequest(label string) *model.Request {
	return c.caches[label].request
}

func (c *cacheRequests) GetPanel(label string) *tview.Flex {
	return c.caches[label].panel
}

func (c *cacheRequests) GetList() []string {
	return c.listLabel
}

func (c *cacheRequests) DeleteMap(label string) {
	delete(c.caches, label)
}

func (c *cacheRequests) DeleteList(index int) {
	c.listLabel = slices.Delete(c.listLabel, index, index+1)
}

func (c *cacheRequests) GetMap() map[string]*model.Request {
	m := make(map[string]*model.Request)

	for i, v := range c.caches {
		m[i] = v.request
	}

	return m
}

func (c *cacheRequests) Save() error {
	b, err := util.JsonMarshal(c.GetMap())
	if err != nil {
		return err
	}

	s := util.Encode(b)

	return util.WriteFile(".cache", []byte(s))
}
