package api

type TaobaoTbkDgPunishOrderGetResponse struct {
	TbkDgPunishOrderGetResponse *TbkPunishOrderResult `json:"tbk_dg_punish_order_get_response"`
}

type TbkPunishOrderResult struct {
	Result *TbkPunishOrderData `json:"result"`
}

type TbkPunishOrderData struct {
	Data         *TbkPunishOrderDataResults `json:"data"`
	BizErrorDesc string                     `json:biz_error_desc"`
	BizErrorCode int                        `json:biz_error_code"`
	ResultMsg    string                     `json:result_msg"`
	ResultCode   int                        `json:result_code"`
}

type TbkPunishOrderDataResults struct {
	PageNo     string                    `json:"page_no"`
	PageSize   string                    `json:"page_size"`
	TotalCount string                    `json:"total_count"`
	Results    *TbkPunishOrderDataResult `json:"results"`
}

type TbkPunishOrderDataResult struct {
	Result *[]TbkPunishOrderItem `json:"result"`
}

type TbkPunishOrderItem struct {
	RelationId        int    `json:"relation_id"`
	SettleMonth       int    `json:"settle_month"`
	PunishStatus      string `json:"punish_status"`
	ViolationType     string `json:"violation_type"`
	TkTradeCreateTime string `json:"tk_trade_create_time"`
	TbTradeId         int64  `json:"tb_trade_id"`
	TbTradeParentId   int64  `json:"tb_trade_parent_id"`
	TkAdzoneId        int    `json:"tk_adzone_id"`
	TkSiteId          int    `json:"tk_site_id"`
	TkPubId           string `json:"tk_pub_id"`
	SpecialId         int    `json:"special_id"`
	UnionId           string `json:"union_id"`
}
