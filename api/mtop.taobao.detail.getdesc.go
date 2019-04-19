package api

/*
{
	"api": "mtop.taobao.detail.getdesc",
	"v": "6.0",
	"ret": ["SUCCESS::调用成功"],
	"data": {
		"sellerId": "2187138540",
		"pcDescContent": "<p><img align=\"absmiddle\" src=\"//img.alicdn.com/imgextra/i1/2187138540/O1CN012CxMrPuI67ifsD4_!!2187138540.jpg\" size=\"790x160\"><img align=\"absmiddle\" src=\"//img.alicdn.com/imgextra/i3/2187138540/TB2XzfismBYBeNjy0FeXXbnmFXa-2187138540.jpg\" size=\"790x292\"><img src=\"//img.alicdn.com/imgextra/i2/2187138540/O1CN01DxpRIJ2CxMu1bxd5T_!!2187138540.jpg\" align=\"absmiddle\"><img src=\"//img.alicdn.com/imgextra/i2/2187138540/O1CN01H22IVN2CxMu0eJ7Us_!!2187138540.jpg\" align=\"absmiddle\"><img align=\"absmiddle\" src=\"//img.alicdn.com/imgextra/i4/2187138540/O1CN01MU9bs12CxMtJ4JVWi_!!2187138540.jpg\" size=\"790x634\"><img align=\"absmiddle\" src=\"//img.alicdn.com/imgextra/i2/2187138540/TB2xaiJfBjTBKNjSZFDXXbVgVXa-2187138540.jpg\" size=\"790x1140\"><img align=\"absmiddle\" src=\"//img.alicdn.com/imgextra/i2/2187138540/TB20.xnsvImBKNjSZFlXXc43FXa_!!2187138540.jpg\" size=\"790x820\"><img align=\"absmiddle\" src=\"//img.alicdn.com/imgextra/i3/2187138540/TB2Cp_iivuSBuNkHFqDXXXfhVXa-2187138540.jpg\" size=\"790x1248\"><img align=\"absmiddle\" src=\"//img.alicdn.com/imgextra/i2/2187138540/TB2tqDVqMaTBuNjSszfXXXgfpXa-2187138540.jpg\" size=\"790x1236\"><img align=\"absmiddle\" src=\"//img.alicdn.com/imgextra/i3/2187138540/TB2FFhAfqAoBKNjSZSyXXaHAVXa-2187138540.jpg\" size=\"790x955\"><img align=\"absmiddle\" src=\"//img.alicdn.com/imgextra/i2/2187138540/TB2Bmm5frZnBKNjSZFrXXaRLFXa-2187138540.jpg\" size=\"790x822\"><img align=\"absmiddle\" src=\"//img.alicdn.com/imgextra/i1/2187138540/TB2PzmWiviSBuNkSnhJXXbDcpXa-2187138540.jpg\" size=\"790x850\"><img align=\"absmiddle\" src=\"//img.alicdn.com/imgextra/i2/2187138540/TB2XMNnfsUrBKNjSZPxXXX00pXa-2187138540.jpg\" size=\"790x896\"><img align=\"absmiddle\" src=\"//img.alicdn.com/imgextra/i1/2187138540/TB2NV9Dq_tYBeNjy1XdXXXXyVXa-2187138540.jpg\" size=\"790x856\"><img align=\"absmiddle\" src=\"//img.alicdn.com/imgextra/i3/2187138540/TB2ts26qFGWBuNjy0FbXXb4sXXa-2187138540.jpg\" size=\"790x976\"><img align=\"absmiddle\" src=\"//img.alicdn.com/imgextra/i4/2187138540/TB25w79qNGYBuNjy0FnXXX5lpXa-2187138540.jpg\" size=\"790x760\"><img align=\"absmiddle\" src=\"//img.alicdn.com/imgextra/i3/2187138540/TB2bVtAfqAoBKNjSZSyXXaHAVXa-2187138540.jpg\" size=\"790x768\" /></p>",
		"itemProperties": [{
			"name": "是否为特殊用途化妆品",
			"value": "否 "
		}, {
			"name": "功效",
			"value": "其他 "
		}, {
			"name": "品牌",
			"value": "滋韵堂 "
		}, {
			"name": "型号",
			"value": "zt-50贴 "
		}],
		"anchors": []
	}
}
*/

type MtopTaobaoDetailGetDesc struct {
	Data *GetDescData `json:"data"`
}
type GetDescData struct {
	PcDescContent string `json:"pcDescContent"`
}
