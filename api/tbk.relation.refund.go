package api

type TaobaoTbkRelationRefundResponse struct {
	TbkRelationRefundResponse *TbkRelationRefundResult `json:"tbk_relation_refund_response"`
}

type TbkRelationRefundResult struct {
	Result *TbkRelationRefundData `json:"result"`
}

type TbkRelationRefundData struct {
	Data         *TbkRelationRefundDataResults `json:"data"`
	BizErrorDesc string                        `json:biz_error_desc"`
	BizErrorCode int                           `json:biz_error_code"`
	ResultMsg    string                        `json:result_msg"`
	ResultCode   int                           `json:result_code"`
}

type TbkRelationRefundDataResults struct {
	PageNo     string                       `json:"page_no"`
	PageSize   string                       `json:"page_size"`
	TotalCount string                       `json:"total_count"`
	Results    *TbkRelationRefundDataResult `json:"results"`
}

type TbkRelationRefundDataResult struct {
	Result *[]TbkRelationRefundItem `json:"result"`
}

type TbkRelationRefundItem struct {
	TbTradeParentId             int64  `json:"tb_trade_parent_id"`
	RelationId                  int64  `json:"relation_id"`
	Tk3rdPubId                  int64  `json:"tk3rd_pub_id"`
	TkPubId                     int64  `json:"tk_pub_id"`
	TkSubsidyFeeRefund3rdPub    string `json:"tk_subsidy_fee_refund3rd_pub"`
	TkCommissionFeeRefund3rdPub string `json:"tk_commission_fee_refund3rd_pub"`
	TkSubsidyFeeRefundPub       string `json:"tk_subsidy_fee_refund_pub"`
	TkCommissionFeeRefundPub    string `json:"tk_commission_fee_refund_pub"`
	TkRefundSuitTime            string `json:"tk_refund_suit_time"`
	TkRefundTime                string `json:"tk_refund_time"`
	EarningTime                 string `json:"earning_time"`
	TbTradeCreateTime           string `json:"tb_trade_create_time"`
	RefundStatus                int    `json:"refund_status"`
	TbAuctionTitle              string `json:"tb_auction_title"`
	TbTradeId                   int64  `json:"tb_trade_id"`
	SpecialId                   int64  `json:"special_id"`
	RefundFee                   string `json:"refund_fee"`
}
