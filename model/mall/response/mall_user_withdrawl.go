package response

import (
	"main.go/model/mall"
	"main.go/model/manage"
)

type WithdrawResponse struct {
	MallUserWithdraw manage.MallUserWithdraw `json:"withdraw" `
	MallUserBank     mall.MallUserBank       `json:"bank" `
}
