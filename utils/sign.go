package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"sort"
	"strings"
)

type HashMethod func(key, str string) string

type HashMethodName string

const MD5 HashMethodName = "md5"
const HASH_MD5 HashMethodName = "hmac"

func Sign(key string, params url.Values, hashMethod HashMethod) string {
	joinedParams := ""
	keys := make([]string, 0)
	for key, _ := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for i := range keys {
		joinedParams += keys[i] + params[keys[i]][0]
	}
	return hashMethod(key, joinedParams)
}

func Hmac(key, data string) string {
	hmac := hmac.New(md5.New, []byte(key))
	hmac.Write([]byte(data))
	return strings.ToUpper(hex.EncodeToString(hmac.Sum([]byte(""))))
}

func Md5(key, data string) string {
	h := md5.New()
	h.Write([]byte(key + data + key))
	return strings.ToUpper(hex.EncodeToString(h.Sum([]byte(""))))
}
