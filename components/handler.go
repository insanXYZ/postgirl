package components

import (
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"

	"postgirl/components/common"
	"postgirl/internal/cache"
	"postgirl/lib"
	"postgirl/model"
	"postgirl/util"
)

func (r *RequestResponsePanel) HandlerSend() {
	var paramsMap model.ParamsMap
	var headersMap model.HeadersMap
	var sliceParams []string
	var valInputUrl string
	var bodyReader io.Reader

	// reset response and show loading information
	r.response.Reset()
	r.response.Loading <- true
	defer func() {
		r.response.Loading <- false
	}()

	headers, params, body := r.attribute.GetText()

	// process params json and headers json to map
	err := util.JsonUnmarshal([]byte(params), &paramsMap)
	if err != nil {
		common.ShowNotification(&common.NotificationConfig{
			Message: model.ErrInvalidFormatParams,
		})
		return
	}

	err = util.JsonUnmarshal([]byte(headers), &headersMap)
	if err != nil {
		common.ShowNotification(&common.NotificationConfig{
			Message: model.ErrInvalidFormatHeaders,
		})
		return
	}

	url, err := util.ParseUrl(r.inputBar.Url)
	if err != nil {
		common.ShowNotification(&common.NotificationConfig{
			Message: model.ErrInvalidFormatUrl,
		})
		return
	}

	// merge params from url and text area
	for i, v := range url.Params {
		for _, val := range v {
			if !slices.Contains(paramsMap[i], val) {
				paramsMap[i] = append(paramsMap[i], val)
			}
		}
	}

	// assign params to url
	valInputUrl = url.CleanUrl

	if len(paramsMap) != 0 {
		valInputUrl += "?"
	}

	for i, v := range paramsMap {
		for _, s := range v {
			sliceParams = append(sliceParams, fmt.Sprintf("%v=%v", i, s))
		}
	}

	valInputUrl += strings.Join(sliceParams, "&")

	// process body

	if r.attribute.BodyTypeSelected != model.NONE {

		var mapBody model.BodyMap
		var err error

		if r.attribute.BodyTypeSelected != model.XML {
			err := util.JsonUnmarshal([]byte(body), &mapBody)
			if err != nil {
				common.ShowNotification(&common.NotificationConfig{
					Message: model.ErrInvalidFormatBody,
				})
				return
			}
		}

		switch r.attribute.BodyTypeSelected {
		case model.BodyOptions[1]: // form data
			bodyReader, r.attribute.BodyTypeSelected, err = util.CreateReaderFormDataType(mapBody)
		case model.BodyOptions[2]: // x-www-form-urlencoded
			bodyReader = util.CreateReaderXWWWFormUrlencodedType(mapBody)
		case model.BodyOptions[3]: // json
			bodyReader, err = util.CreateReaderJsonType(mapBody)
		case model.BodyOptions[4]: // xml
			bodyReader, err = util.CreateReaderXmlType(body)
		}

		if err != nil {
			common.ShowNotification(&common.NotificationConfig{
				Message: model.ErrInvalidFormatBody,
			})
			return
		}
	}

	lib.Tview.UpdateDraw(func() {
		r.inputBar.InputUrl.SetText(valInputUrl)
		p, _ := util.JsonMarshalString(paramsMap)
		h, _ := util.JsonMarshalString(headersMap)
		r.attribute.SetTextHeaders(h)
		r.attribute.SetTextParams(p)
	})

	req := &model.Request{
		Method: r.inputBar.Method,
		Url:    valInputUrl,
		Attribute: &model.Attribute{
			Params:   paramsMap,
			Headers:  headersMap,
			BodyType: r.attribute.BodyTypeSelected,
			Body:     bodyReader,
		},
	}

	defer func() {
		*r.currentRequest = *req
		err := cache.CacheRequests.Save()
		if err != nil {
			common.ShowNotification(&common.NotificationConfig{
				Message: model.ErrSaveCache,
			})
		}
	}()

	res, err := util.NewRequest(req)
	if err != nil {
		common.ShowNotification(&common.NotificationConfig{
			Message: err.Error(),
		})
		return
	}

	defer res.Body.Close()

	var headerJson string
	var resBody string

	b, err := io.ReadAll(res.Body)
	if err != nil {
		common.ShowNotification(&common.NotificationConfig{
			Message: model.ErrReadResponseBody,
		})
		return
	}
	resBody = string(b)

	headerJson, err = util.JsonMarshalString(res.Header)
	if err != nil {
		common.ShowNotification(&common.NotificationConfig{
			Message: model.ErrReadHeader,
		})
		return
	}
	r.response.StatusCode = strconv.Itoa(res.StatusCode)

	lib.Tview.UpdateDraw(func() {
		r.response.SetBodyText(resBody)
		r.response.SetHeaderText(headerJson)
	})
}
