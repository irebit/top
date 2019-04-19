package api

type TaobaoTbkScInvitecodeGet struct {
	TbkScInvitecodeGetResponse *TbkScInvitecodeGetResponse `json:"tbk_sc_invitecode_get_response"`
}

type TbkScInvitecodeGetResponse struct {
	Data *InviteCode `json:"data"`
}

type InviteCode struct {
	InviteCode string `json:"inviter_code"`
}
