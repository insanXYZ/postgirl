package components

import (
	"fmt"
	"io"
	"postgirl/app/components/common"
	"postgirl/app/lib"
	"postgirl/app/model"
	"postgirl/app/util"
	"slices"
	"strings"

	"github.com/rivo/tview"
)

type InputBar struct {
	Method   int
	Dropdown *tview.DropDown

	Url      string
	InputUrl *tview.InputField

	Root *tview.Flex
}

func (r *RequestResponsePanel) NewInputUrl() {
	inputBar := &InputBar{}

	methodDropdown := common.CreateDropdown(&common.DropdownConfig{
		Options: model.Methods,
		SelectedFunc: func(_ string, index int) {
			inputBar.Method = index
		},
		CurrentOptions: r.currentModel.Method,
	})
	inputBar.Dropdown = methodDropdown

	inputUrl := common.CreateInputField(&common.InputFieldConfig{
		Placeholder: "Enter URL",
		DefaultText: r.currentModel.Url,
		ChangedFunc: func(text string) {
			inputBar.Url = text
		},
	})
	inputBar.InputUrl = inputUrl

	submitButton := tview.NewButton("send")
	submitButton.SetSelectedFunc(func() {
		r.submit <- true
	})

	flex := tview.NewFlex()
	flex.SetBorder(true)
	flex.AddItem(methodDropdown, 9, 1, false)
	flex.AddItem(inputUrl, 0, 1, false)
	flex.AddItem(tview.NewBox(), 2, 1, false)
	flex.AddItem(submitButton, 6, 1, false)

	inputBar.Root = flex
	r.inputBar = inputBar
}

func (r *RequestResponsePanel) listenChan() {
	for {
		select {
		case <-r.submit:
			var paramsMap model.ParamsMap
			var headersMap model.HeadersMap
			var sliceParams []string
			var valInputUrl string
			var bodyReader io.Reader
			var err error

			headers, params, body := r.attribute.GetText()

			// process params json and headers json to map
			err = util.JsonUnmarshal([]byte(params), &paramsMap)
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

			// process body

			if r.attribute.BodyTypeSelected != model.NONE {

				var xmlbody any
				var mapBody model.BodyMap

				if r.attribute.BodyTypeSelected == model.XML {
					err = util.XmlUnmarshal([]byte(body), &xmlbody)
				} else {
					err = util.JsonUnmarshal([]byte(body), &mapBody)
				}

				if err != nil {
					panic(err.Error()) //should be use notification
				}

				switch r.attribute.BodyTypeSelected {
				case model.BodyOptions[1]: //form data
					bodyReader, r.attribute.BodyTypeSelected, err = util.CreateReaderFormDataType(mapBody)
				case model.BodyOptions[2]: // x-www-form-urlencoded
					bodyReader = util.CreateReaderXWWWFormUrlencodedType(mapBody)
				case model.BodyOptions[3]: // json
					bodyReader, err = util.CreateReaderJsonType(mapBody)
				case model.BodyOptions[4]: //xml
					bodyReader, err = util.CreateReaderXmlType(xmlbody)
				}
			}

			if err != nil {
				panic(err.Error()) //should be use notification
			}

			//merge params from url and text area
			for i, v := range url.Params {
				for _, val := range v {
					if !slices.Contains(paramsMap[i], val) {
						paramsMap[i] = append(paramsMap[i], val)
					}
				}
			}

			valInputUrl = url.CleanUrl + "?"

			for i, v := range paramsMap {
				for _, s := range v {
					sliceParams = append(sliceParams, fmt.Sprintf("%v=%v", i, s))
				}
			}

			valInputUrl += strings.Join(sliceParams, "&")

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

			util.NewRequest(req)
		}
	}
}
