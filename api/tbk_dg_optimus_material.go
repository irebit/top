package api

type TaobaoTbkDgOptimusMaterialOptional struct {
	TbkDgOptimusMaterialResponse *TbkDgOptimusMaterialResponse `json:"tbk_dg_optimus_material_response"`
}

type TbkDgOptimusMaterialResponse struct {
	// TotalResults int      `json:"total_results"`
	ResultList *MapData `json:"result_list"`
	RequestId  string   `json:"request_id"`
}


