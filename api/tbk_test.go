package api

import (
	"log"
	"testing"

	"github.com/irebit/top/top_error"

	"github.com/irebit/top/utils"
)

func TestTbk(t *testing.T) {
	TaobaoAppID := "xxxxxxx"
	TaobaoAppSecret := "xxxxxxxxxxxxxxxxxxxxx"
	tbk := &TBK{}
	tbk.SetAppInfo(TaobaoAppID, TaobaoAppSecret)
	tbk.SetHashMethod(utils.HASH_MD5)

	// tbk.SetParams(
	// 	"taobao.tbk.dg.material.optional",
	// 	"",
	// 	map[string]string{
	// 		"adzone_id": "103720350305",
	// 		"cat":       "16,18",
	// 	})

	// tbk.SetParams(
	// 	"taobao.itemcats.get",
	// 	"",
	// 	map[string]string{
	// 		"parent_cid": "0",
	// 	})

	params :=
		map[string]interface{}{
			"method":      "taobao.tbk.dg.material.optional",
			"adzone_id":   "103720350305",
			"material_id": "3756",
		}

	var resp *taobao_tbk_dg_material_optional
	var respErr *top_error.TbkErrResponse
	_, err := tbk.GetResponse(params, resp, respErr)
	log.Println(err)
	log.Println(resp, respErr)

}
