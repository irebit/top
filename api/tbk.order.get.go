package api

type TaobaoTbkOrderGetResponse struct {
	TbkOrderGetResponse *TbkOrderResult `json:"tbk_order_get_response"`
}

type TbkOrderResult struct {
	Results *NTbkOrder `json:"results"`
}

type NTbkOrder struct {
	NTbkOrder *[]TbkOrderItem `json:"n_tbk_order"`
}

type TbkOrderItem struct {
	TradeParentId       int64  `json:"trade_parent_id"`
	TradeId             int64  `json:"trade_id"`
	NumIid              int64  `json:"num_iid"`
	ItemTitle           string `json:"item_title"`
	ItemNum             int    `json:"item_num"`
	Price               string `json:"price"`
	PayPrice            string `json:"pay_price"`
	SellerNick          string `json:"seller_nick"`
	SellerShopTitle     string `json:"seller_shop_title"`
	Commission          string `json:"commission"`
	CommissionRate      string `json:"commission_rate"`
	Unid                string `json:"unid"`
	CreateTime          string `json:"create_time"`
	EarningTime         string `json:"earning_time"`
	TkStatus            int    `json:"tk_status"`
	Tk3rdType           string `json:"tk3rd_type"`
	Tk3rdPubId          int    `json:"tk3rd_pub_id"`
	OrderType           string `json:"order_type"`
	IncomeRate          string `json:"income_rate"`
	PubSharePreFee      string `json:"pub_share_pre_fee"`
	SubsidyRate         string `json:"subsidy_rate"`
	SubsidyType         string `json:"subsidy_type"`
	TerminalType        string `json:"terminal_type"`
	AuctionCategory     string `json:"auction_category"`
	SiteId              string `json:"site_id"`
	SiteName            string `json:"site_name"`
	AdzoneId            string `json:"adzone_id"`
	AdzoneName          string `json:"adzone_name"`
	AlipayTotalPrice    string `json:"alipay_total_price"`
	TotalCommissionRate string `json:"total_commission_rate"`
	TotalCommissionFee  string `json:"total_commission_fee"`
	SubsidyFee          string `json:"subsidy_fee"`
	RelationId          int    `json:"relation_id"`
	SpecialId           int    `json:"special_id"`
	ClickTime           string `json:"click_time"`
}
