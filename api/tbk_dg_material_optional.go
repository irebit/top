package api

type TaobaoTbkDgMaterialOptional struct {
	TbkDgMaterialOptionalResponse *TbkDgMaterialOptionalResponse `json:"tbk_dg_material_optional_response"`
}

type TbkDgMaterialOptionalResponse struct {
	TotalResults int      `json:"total_results"`
	ResultList   *MapData `json:"result_list"`
	RequestId    string   `json:"request_id"`
}

type MapData struct {
	MapData []*GoodsItem `json:"map_data"`
}

type GoodsItem struct {
	CategoryId           int                 `json:"category_id"`
	CategoryName         string              `json:"category_name"`
	CommissionRate       string              `json:"commission_rate"`
	CommissionType       string              `json:"commission_type"`
	CouponAmount         interface{}         `json:"coupon_amount"`
	CouponEndTime        string              `json:"coupon_end_time"`
	CouponId             string              `json:"coupon_id"`
	CouponInfo           string              `json:"coupon_info"`
	CouponRemainCount    int                 `json:"coupon_remain_count"`
	CouponShareUrl       string              `json:"coupon_share_url"`
	CouponStartFee       string              `json:"coupon_start_fee"`
	CouponStartTime      string              `json:"coupon_start_time"`
	CouponTotalCount     int                 `json:"coupon_total_count"`
	IncludeDxjh          string              `json:"include_dxjh"`
	IncludeMkt           string              `json:"include_mkt"`
	InfoDxjh             string              `json:"info_dxjh"`
	ItemDescription      string              `json:"item_description"`
	ItemId               int64               `json:"item_id"`
	ItemUrl              string              `json:"item_url"`
	LevelOneCategoryId   int                 `json:"level_one_category_id"`
	LevelOneCategoryName string              `json:"level_one_category_name"`
	Nick                 string              `json:"nick"`
	NumIid               int                 `json:"num_iid"`
	PictUrl              string              `json:"pict_url"`
	Provcity             string              `json:"provcity"`
	ReservePrice         string              `json:"reserve_price"`
	SellerId             int                 `json:"seller_id"`
	ShopDsr              int                 `json:"shop_dsr"`
	ShopTitle            string              `json:"shop_title"`
	ShortTitle           string              `json:"short_title"`
	SmallImages          map[string][]string `json:"small_images"`
	Title                string              `json:"title"`
	TkTotalCommi         string              `json:"tk_total_commi"`
	TkTotalSales         string              `json:"tk_total_sales"`
	Url                  string              `json:"url"`
	UserType             int                 `json:"user_type"`
	Volume               int                 `json:"volume"`
	WhiteImage           string              `json:"white_image"`
	XId                  string              `json:"x_id"`
	XkFinalPrice         string              `json:"zk_final_price"`
}
