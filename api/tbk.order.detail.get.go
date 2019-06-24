package api

type TaobaoTbkOrderDetailsGetResponse struct {
	TbkOrderDetailsGetResponse *TbkOrderDetailsData `json:"tbk_order_details_get_response"`
}

type TbkOrderDetailsData struct {
	Data *TbkOrderDetailsResults `json:"data"`
}

type TbkOrderDetailsResults struct {
	Results       *TbkOrderDetailsPublisherOrderDto `json:"results"`
	HasPre        bool                              `json:"has_pre"`
	PositionIndex string                            `json:"position_index"`
	HasNext       bool                              `json:"has_next"`
	PageNo        int                               `json:"page_no"`
	PageSize      int                               `json:"page_size"`
}
type TbkOrderDetailsPublisherOrderDto struct {
	PublisherOrderDto *[]TbkOrderDetailsItem `json:"publisher_order_dto"`
}

type TbkOrderDetailsItem struct {
	TbPaidTime          string `json:"tb_paid_time"`
	TkPaidTime          string `json:"tk_paid_time"`
	PayPrice            string `json:"pay_price"`
	PubShareFee         string `json:"pub_share_fee"`
	TradeId             string `json:"trade_id"`
	TkOrderRole         int    `json:"tk_order_role"`
	TkEarningTime       string `json:"tk_earning_time"`
	AdzoneId            int    `json:"adzone_id"`
	PubShareRate        string `json:"pub_share_rate"`
	Unid                int    `json:"unid"`
	RefundTag           int    `json:"refund_tag"`
	SubsidyRate         string `json:"subsidy_rate"`
	TkTotalRate         string `json:"tk_total_rate"`
	ItemCategoryName    string `json:"item_category_name"`
	SellerNick          string `json:"seller_nick"`
	PubId               int    `json:"pub_id"`
	AlimamaRate         string `json:"alimama_rate"`
	SubsidyType         string `json:"subsidy_type"`
	ItemImg             string `json:"item_img"`
	PubSharePreFee      string `json:"pub_share_pre_fee"`
	AlipayTotalPrice    string `json:"alipay_total_price"`
	ItemTitle           string `json:"item_title"`
	SiteName            string `json:"site_name"`
	ItemNum             int    `json:"item_num"`
	SubsidyFee          string `json:"subsidy_fee"`
	AlimamaShareFee     string `json:"alimama_share_fee"`
	TradeParentId       string `json:"trade_parent_id"`
	OrderType           string `json:"order_type"`
	TkCreateTime        string `json:"tk_create_time"`
	FlowSource          string `json:"flow_source"`
	TerminalType        string `json:"terminal_type"`
	ClickTime           string `json:"click_time"`
	TkStatus            int    `json:"tk_status"`
	ItemPrice           string `json:"item_price"`
	ItemId              int    `json:"item_id"`
	AdzoneName          string `json:"adzone_name"`
	TotalCommissionRate string `json:"total_commission_rate"`
	ItemLink            string `json:"item_link"`
	SiteId              int    `json:"site_id"`
	SellerShopTitle     string `json:"seller_shop_title"`
	IncomeRate          string `json:"income_rate"`
	TotalCommissionFee  string `json:"total_commission_fee"`
}
