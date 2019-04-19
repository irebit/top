package top_error

import (
	"strconv"
)

type TbkError struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`
}

type TbkErrResponse struct {
	ErrResponse *TbkError `json:"error_response"`
}

func (t *TbkErrResponse) Error() string {
	return "Code:" + strconv.Itoa(t.ErrResponse.Code) + ", Msg:" + t.ErrResponse.Msg + ", SubCode:" + t.ErrResponse.SubCode + " , SubMsg: " + t.ErrResponse.SubMsg
}

func (t *TbkErrResponse) IsNil() bool {
	if t.ErrResponse == nil {
		return true
	}
	return false
}
