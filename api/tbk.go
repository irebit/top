package api

import (
	"errors"
	"net/url"

	"github.com/irebit/top/app"
	"github.com/irebit/top/utils"
)

// const TaobaoUrl = "http://gw.api.taobao.com/router/rest"
const TaobaoUrl = "https://eco.taobao.com/router/rest"

type TBK struct {
	appInfo    *app.AppInfo
	reqUrl     string
	values     url.Values
	hashMethod utils.HashMethod
}

func (t *TBK) SetAppInfo(appKey string, secret string) {
	t.appInfo = &app.AppInfo{appKey, secret}
}

func (t *TBK) SetHashMethod(name utils.HashMethodName) {
	if name == utils.HASH_MD5 {
		t.hashMethod = utils.Hmac
	} else {
		t.hashMethod = utils.Md5
	}
	t.SetValue("sign_method", string(name))
}

func (t *TBK) SetReqUrl(reqUrl string) {
	t.reqUrl = reqUrl
}

func (t *TBK) GetReqUrl() string {
	return t.reqUrl
}

func (t *TBK) SetValue(key, value string) {
	if t.values == nil {
		t.values = url.Values{}
	}
	t.values.Set(key, value)
}

func (t *TBK) GetValue(key string) (string, error) {
	if t.values == nil {
		return "", errors.New("values are not set")
	}
	return t.values.Get(key), nil
}

func (t *TBK) GetValues() url.Values {
	return t.values
}

func (t *TBK) DelValue(key string) {
	t.values.Del(key)
}

func (t *TBK) SetParams(params map[string]interface{}) {
	t.SetReqUrl(TaobaoUrl)
	t.SetValue("format", "json")
	t.SetValue("v", "2.0")
	t.SetValue("app_key", t.appInfo.AppKey)
	t.SetValue("timestamp", utils.GetCurrentDateTime())

	for k, v := range params {
		t.SetValue(k, utils.GetString(v))
	}
}

func (t *TBK) GetResponse(params map[string]interface{}, resp interface{}, respErr interface{}) ([]byte, error) {

	//method session
	if t.hashMethod == nil {
		t.SetHashMethod(utils.MD5)
	}

	t.SetParams(params)

	t.DelValue("sign")
	t.SetValue("sign", utils.Sign(t.appInfo.AppSecret, t.values, t.hashMethod))

	response := &Response{
		t.reqUrl,
		t.GetValues(),
		&resp,
		&respErr,
		nil,
	}
	data, err := response.Get()
	return data, err
}
