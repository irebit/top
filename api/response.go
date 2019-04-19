package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type TopRequestMethod func(string, url.Values, interface{}, interface{}) ([]byte, error)

type Response struct {
	Url     string
	Params  url.Values
	Resp    interface{}
	RespErr interface{}
	Method  TopRequestMethod
}

func (t *Response) Get() ([]byte, error) {
	if t.Method == nil {
		t.Method = PostForm
	}
	return t.Method(t.Url, t.Params, t.Resp, t.RespErr)
}

func PostForm(url string, params url.Values, resp interface{}, respErr interface{}) ([]byte, error) {
	response, progErr := http.PostForm(url, params)

	if progErr != nil {
		fmt.Println(progErr.Error())
		return nil, progErr
	}
	data, progErr := ioutil.ReadAll(response.Body)
	if progErr != nil {
		return data, progErr
	}

	progErr = json.Unmarshal(data, &respErr)

	if progErr != nil {
		return data, progErr
	}

	err := json.Unmarshal(data, &resp)
	return data, err
}

func Get(url string, params url.Values, resp interface{}, respErr interface{}) ([]byte, error) {
	response, progErr := http.Get(url)

	if progErr != nil {
		fmt.Println(progErr.Error())
		return nil, progErr
	}
	data, progErr := ioutil.ReadAll(response.Body)
	if progErr != nil {
		return data, progErr
	}
	progErr = json.Unmarshal(data, &respErr)

	if progErr != nil {
		return data, progErr
	}

	err := json.Unmarshal(data, &resp)
	return data, err
}
