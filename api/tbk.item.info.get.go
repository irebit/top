package api

type TaobaoTbkItemInfoGetResponse struct {
	TbkItemInfoGetResponse *TbkItemInfoResults `json:"tbk_item_info_get_response"`
}

type TbkItemInfoResults struct {
	Results *NTbkItem `json:"results"`
}

type NTbkItem struct {
	NTbkItem []*TbkItemInfo `json:"n_tbk_item"`
}

type TbkItemInfo struct {
	CatName         string              `json:"cat_name"`
	NumIid          int                 `json:"num_iid"`
	Title           string              `json:"title"`
	PictUrl         string              `json:"pict_url"`
	SmallImages     map[string][]string `json:"small_images"`
	ReservePrice    string              `json:"reserve_price"`
	XkFinalPrice    string              `json:"zk_final_price"`
	UserType        int                 `json:"user_type"`
	Provcity        string              `json:"provcity"`
	ItemUrl         string              `json:"item_url"`
	SellerId        int                 `json:"seller_id"`
	Volume          int                 `json:"volume"`
	Nick            string              `json:"nick"`
	cat_leaf_name   string              `json:"cat_leaf_name"`
	is_prepay       string              `json:"is_prepay"`
	ShopDsr         int                 `json:"shop_dsr"`
	ratesum         int                 `json:"ratesum"`
	IRfdRate        bool                `json:"i_rfd_rate"`
	HGoodRate       bool                `json:"h_good_rate"`
	HPayRate30      bool                `json:"h_pay_rate30"`
	FreeShipment    bool                `json:"free_shipment"`
	MaterialLibType string              `json:"material_lib_type"`
}
