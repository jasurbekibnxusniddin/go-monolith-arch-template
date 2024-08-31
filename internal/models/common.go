package models

type ByIdReq struct {
	Id string `json:"id"`
}

type GetListReq struct {
	Page   int32  `json:"page"`
	Limit  int32  `json:"limit"`
	Search string `json:"search"`
}

type Empty struct{}
