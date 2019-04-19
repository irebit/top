package api

// {
//     "tbk_tpwd_create_response":{
//         "data":{
//             "model":"￥AADPOKFz￥"
//         }
//     }
// }

type TaobaoTbkTpwdCreate struct {
	TbkTpwdCreateResponse *TbkTpwdCreateResponse `json:"tbk_tpwd_create_response"`
}

type TbkTpwdCreateResponse struct {
	Data *TpwdData `json:"data"`
}

type TpwdData struct {
	Model string `json:"model"`
}
