package api

import (
	"net/url"

	"github.com/irebit/top/app"
	"github.com/irebit/top/enum"
	"github.com/irebit/top/top_error"
)

var (
	TaobaoOauthUrl = "https://oauth.m.taobao.com/authorize"
	TaobaoTokenUrl = "https://oauth.m.taobao.com/token"
)

type Oauth struct {
	AppInfo     *app.AppInfo `json:"app_info"`
	CallbackUrl string       `json:"callbackUrl"`
	View        enum.View    `json:"view"`
}

func (o *Oauth) GetRedirectUri(state string) string {
	params := url.Values{}
	params.Set("client_id", o.AppInfo.AppKey)
	params.Set("response_type", "code")
	params.Set("view", string(o.View))
	params.Set("redirect_uri", o.CallbackUrl)
	params.Set("state", state)
	return TaobaoOauthUrl + "?" + params.Encode()
}

func (o *Oauth) GetToken(code string) (*app.TokenInfo, *top_error.OauthErr, error) {
	params := url.Values{}
	params.Set("client_id", o.AppInfo.AppKey)
	params.Set("client_secret", o.AppInfo.AppSecret)
	params.Set("code", code)
	params.Set("redirect_uri", o.CallbackUrl)
	params.Set("grant_type", "authorization_code")

	resp := &app.TokenInfo{}
	respErr := &top_error.OauthErr{}

	response := &Response{
		TaobaoTokenUrl,
		params,
		resp,
		respErr,
		nil,
	}
	_, err := response.Get()

	return resp, respErr, err
}

func (o *Oauth) RefreshToken(refreshToken string) (*app.TokenInfo, *top_error.OauthErr, error) {
	params := url.Values{}
	params.Set("client_id", o.AppInfo.AppKey)
	params.Set("client_secret", o.AppInfo.AppSecret)
	params.Set("refresh_token", refreshToken)
	params.Set("grant_type", "refresh_token")

	resp := &app.TokenInfo{}
	respErr := &top_error.OauthErr{}

	response := &Response{
		TaobaoTokenUrl,
		params,
		resp,
		respErr,
		nil,
	}
	_, err := response.Get()
	return resp, respErr, err
}
