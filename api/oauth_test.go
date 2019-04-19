package api

import (
	"log"
	"testing"

	"github.com/irebit/top/app"
)

func TestGetRedirectUrl(t *testing.T) {
	o := &Oauth{
		&app.AppInfo{"25795647", "1e179ef2e5d370ede41fabf8be7d4d6b"},
		"https://127.0.0.1",
		"wap",
	}

	// log.Println(o.GetRedirectUri("asdfsadf"))

	var resp *app.TokenInfo
	// _, err, oauthErr := o.GetToken("IfUTBaZ0vFOXXcNU7wFIG6Rt4409811", &resp)

	_, err, oauthErr := o.RefreshToken("6202510d3e78847f87ZZb4cd7fba0c24135079c00070bcd2200674832579", &resp)

	if err != nil {
		log.Println(err)
	}

	if oauthErr != nil {
		log.Println(oauthErr)
	}

	log.Println(resp)
}
