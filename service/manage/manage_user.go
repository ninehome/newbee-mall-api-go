package manage

import (
	"errors"
	"fmt"
	"main.go/global"
	"main.go/model/common/request"
	"main.go/model/mall"
	"main.go/model/mall/response"
	"main.go/model/manage"
	manageReq "main.go/model/manage/request"
)

type ManageUserService struct {
}

// LockUser 修改用户状态
func (m *ManageUserService) LockUser(idReq request.IdsReq, lockStatus int) (err error) {
	if lockStatus != 0 && lockStatus != 1 {
		return errors.New("操作非法！")
	}
	//更新字段为0时，不能直接UpdateColumns
	err = global.GVA_DB.Model(&manage.MallUser{}).Where("user_id in ?", idReq.Ids).Update("locked_flag", lockStatus).Error
	return err
}

// GetMallUserInfoList 分页获取商城注册用户列表
func (m *ManageUserService) GetMallUserInfoList(info manageReq.MallUserSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNumber - 1)
	// 创建db
	db := global.GVA_DB.Model(&manage.MallUser{})
	var mallUsers []manage.MallUser
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("create_time desc").Find(&mallUsers).Error
	return err, mallUsers, total
}

// 获取 用户提款 列表

func (m *ManageUserService) GetMallUserWithdrawaList(info manageReq.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNumber - 1)
	// 创建db
	db := global.GVA_DB.Model(&manage.MallUserWithdraw{})
	var mallUsers []manage.MallUserWithdraw
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Order("create_time desc").Find(&mallUsers).Error
	if err != nil {
		return
	}

	var wr []response.WithdrawResponse
	for _, value := range mallUsers {

		var bank = mall.MallUserBank{}
		err := global.GVA_DB.Model(&mall.MallUserBank{}).Where("bank_id = ?", value.BankId).First(&bank).Error
		if err != nil {
			fmt.Println("查询 不到记录")
			continue
		}

		//withdraw.MallUserBank = bank
		//withdraw.MallUserWithdraw = value

		wr = append(wr, response.WithdrawResponse{
			MallUserBank:     bank,
			MallUserWithdraw: value,
		})

		fmt.Println(2222222)

	}

	fmt.Println(wr)

	return err, wr, total
	//return err, mallUsers, total
}
