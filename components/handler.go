package components

import (
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"

	"postgirl/components/common"
	"postgirl/internal/log"
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

	bodyType := r.attribute.BodyTypeSelected

	// reset response and show loading information
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

		if r.attribute.BodyTypeSelected == model.XML {
			mapBody, err = util.XmlUnmarshal([]byte(body))
		} else {
			err = util.JsonUnmarshal([]byte(body), &mapBody)
		}

		if err != nil {
			common.ShowNotification(&common.NotificationConfig{
				Message: model.ErrInvalidFormatBody,
			})
			return
		}

		bodyReader, bodyType, err = util.CreateReader(bodyType, mapBody)

		if err != nil {
			common.ShowNotification(&common.NotificationConfig{
				Message: err.Error(),
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
			Params:     paramsMap,
			Headers:    headersMap,
			BodyString: body,
			BodyType:   bodyType,
			Body:       bodyReader,
		},
	}

	defer func() {
		req.Attribute.BodyType = r.attribute.BodyTypeSelected
		*r.currentRequest = *req
		SaveCache()
	}()

	res, bodyReader, err := util.NewRequest(req)
	if err != nil {
		common.ShowNotification(&common.NotificationConfig{
			Message: err.Error(),
		})
		return
	}

	defer res.Body.Close()

	var headerJson string
	var resBody string

	b, err := io.ReadAll(bodyReader)
	if err != nil {
		common.ShowNotification(&common.NotificationConfig{
			Message: model.ErrReadResponseBody,
		})
		return
	}

	headerJson, err = util.JsonMarshalString(res.Header)
	if err != nil {
		common.ShowNotification(&common.NotificationConfig{
			Message: model.ErrReadResponseHeader,
		})
		return
	}
	r.response.StatusCode = strconv.Itoa(res.StatusCode)

	resBody = string(b)

	contentType := res.Header.Get("Content-Type")

	if strings.Contains(contentType, model.CONTENT_TYPE_JSON) {
		var unm any

		if err := util.JsonUnmarshal(b, &unm); err == nil {
			s, err := util.JsonMarshalString(unm)
			if err == nil {
				resBody = s
			}
		}

	} else if strings.Contains(contentType, model.CONTENT_TYPE_XML) {
		m, err := util.XmlUnmarshal(b)
		if err != nil {
			lib.Tview.Stop()
			fmt.Println(err.Error())
		}

		b, err := m.XmlIndent("", " ")
		if err != nil {
			lib.Tview.Stop()
			fmt.Println(err.Error())
		}

		resBody = string(b)
	}

	log.AddLog(fmt.Sprintf("%v %v %v", model.Methods[req.Method], req.Url, res.Status))

	lib.Tview.UpdateDraw(func() {
		r.response.SetBodyText(resBody)
		r.response.SetHeaderText(headerJson)
	})
}
