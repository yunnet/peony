package models

import "github.com/yunnet/walkdog/enums"

type JsonResult struct {
	Code enums.JsonResultCode `json:"code"`
	Msg  string               `json:"msg"`
	Data interface{}          `json:"data"`
}

type JsonPaging struct {
	Pagings Paging      `json:"paging"`
	Item    interface{} `json:"item"`
}

type Paging struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
	Total  int64 `json:"total"`
}
