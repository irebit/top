package top_error

import (
	"fmt"
)

type OauthErr struct {
	ErrorDescription string `json:"error_description"`
	ErrorTitle       string `json:"error"`
}

func (o *OauthErr) Error() string {
	return fmt.Sprintf("error_description:%s,error:%s", o.ErrorDescription, o.ErrorTitle)
}

func (o *OauthErr) IsNil() bool {
	if o.ErrorTitle == "" && o.ErrorDescription == "" {
		return true
	}
	return false
}
