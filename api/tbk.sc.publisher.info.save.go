package api

// {
//     "tbk_sc_publisher_info_save_response":{
//         "data":{
//             "relation_id":40232,
//             "account_name":"xxx",
//             "special_id":32304,
//             "desc":"绑定成功"
//         }
//     }
// }

type TaobaoTbkScPublisherInfoSave struct {
	TbkScPublisherInfoSaveResponse *TbkScPublisherInfoSaveResponse `json:"tbk_sc_publisher_info_save_response"`
}

type TbkScPublisherInfoSaveResponse struct {
	Data *PublisherInfo `json:"data"`
}

type PublisherInfo struct {
	RelationId  int    `json:"relation_id"`
	AccountName string `json:"account_name"`
	SpecialId   int    `json:"special_id"`
	Desc        string `json:"desc"`
}
