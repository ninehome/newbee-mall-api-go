package request

import (
	"main.go/model/common/request"
	"main.go/model/manage"
)

type MallUserSearch struct {
	manage.MallUser
	request.PageInfo
}

//type WithdrawalSearch struct {
//	manage.MallUserWithdraw
//}

type PageInfo struct {
	PageNumber int `json:"pageNumber"` // 页码
	PageSize   int `json:"pageSize"`   // 每页大小
}
