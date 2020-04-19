package models

import "orange_message_service/app/models"

type SendReq struct {
	SourceId int `json:"source_id"`
	MsgKey   int `json:"msg_key"`
	Body     []models.Message
}

type ServerReq struct {
	SourceId int `json:"source_id"`
	MsgKey   int `json:"msg_key"`
	Body     models.Message
}

type SendRes struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type AliSendRes struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

type SendByUsersReq struct {
	SourceId int      `json:"source_id"`
	MsgKey   int      `json:"msg_key"`
	Users    []string `json:"users"`
}

type CustomerSayReq struct {
	Text string `json:"text"`
}