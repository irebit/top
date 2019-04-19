package app

type AppInfo struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
}

type TokenInfo struct {
	W1ExpiresIn           int64  `json:"w1_expires_in"`
	RefreshTokenValidTime int64  `json:"refresh_token_valid_time"`
	TaobaoUserNick        string `json:"taobao_user_nick"`
	ReExpiresIn           int64  `json:"re_expires_in"`
	ExpireTime            int64  `json:"expire_time"`
	TokenType             string `json:"token_type"`
	AccessToken           string `json:"access_token"`
	TaobaoOpenUid         string `json:"taobao_open_uid"`
	W1Valid               int64  `json:"w1_valid"`
	RefreshToken          string `json:"refresh_token"`
	W2ExpiresIn           int64  `json:"w2_expires_in"`
	W2Valid               int64  `json:"w2_valid"`
	R1ExpiresIn           int64  `json:"r1_expires_in"`
	R2ExpiresIn           int64  `json:"r2_expires_in"`
	R2Valid               int64  `json:"r2_valid"`
	R1Valid               int64  `json:"r1_valid"`
	TaobaoUserId          string `json:"taobao_user_id"`
	ExpiresIn             int64  `json:"expires_in"`
}
