package mall

import (
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"main.go/global"
	"main.go/model/common"
	"main.go/model/mall"
	mallReq "main.go/model/mall/request"
	"time"
)

type MallUserAddressService struct {
}

// GetMyAddress 获取收货地址
func (m *MallUserAddressService) GetMyAddress(token string) (err error, userAddress []mall.MallUserAddress) {
	var userToken mall.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户"), userAddress
	}
	global.GVA_DB.Where("user_id=? and is_deleted=0", userToken.UserId).Find(&userAddress)
	return
}

// GetMyAddress 获取银行列表
func (m *MallUserAddressService) GetMyBankList(token string) (err error, userAddress []mall.MallUserAddress) {
	var userToken mall.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("不存在的用户"), userAddress
	}
	global.GVA_DB.Where("user_id=? and is_deleted=0", userToken.UserId).Find(&userAddress)
	return
}

// SaveUserAddress 保存用户地址
func (m *MallUserAddressService) SaveUserAddress(token string, req mallReq.AddAddressParam) (err error) {
	var userToken mall.MallUserToken
	if err = global.GVA_DB.Where("token =?", token).First(&userToken).Error; err != nil {
		return errors.New("不存在的用户")
	}
	var defaultAddress mall.MallUserAddress
	var newAddress mall.MallUserAddress
	fmt.Println("uid ===>>> ")
	// 是否新增了默认地址，将之前的默认地址设置为非默认

	if err = global.GVA_DB.Where("user_id=? and default_flag =1 and is_deleted = 0", userToken.UserId).First(&defaultAddress).Error; err != nil {
		//没有查到记录 ,新增记录
		copier.Copy(&newAddress, &req)
		newAddress.CreateTime = common.JSONTime{Time: time.Now()}
		newAddress.UpdateTime = common.JSONTime{Time: time.Now()}
		newAddress.UserId = userToken.UserId
		err = global.GVA_DB.Create(&newAddress).Error
		if err != nil {
			return
		}
	} else {
		//先更新 之前的记录 再新增
		global.GVA_DB.Model(&mall.MallUserAddress{}).Where("address_id =?", defaultAddress.AddressId).Update("default_flag", 0)

		copier.Copy(&newAddress, &req)
		newAddress.CreateTime = common.JSONTime{Time: time.Now()}
		newAddress.UpdateTime = common.JSONTime{Time: time.Now()}
		newAddress.UserId = userToken.UserId
		err = global.GVA_DB.Create(&newAddress).Error
		if err != nil {
			return
		}

	}

	return
}

// UpdateUserAddress 更新用户地址
func (m *MallUserAddressService) UpdateUserAddress(token string, req mallReq.UpdateAddressParam) (err error) {
	var userToken mall.MallUserToken
	if err = global.GVA_DB.Where("token =?", token).First(&userToken).Error; err != nil {
		return errors.New("不存在的用户")
	}
	var userAddress mall.MallUserAddress
	if err = global.GVA_DB.Where("address_id =? and user_id =?", req.AddressId, userToken.UserId).First(&userAddress).Error; err != nil {
		return errors.New("不存在的用户地址")
	}
	if userToken.UserId != userAddress.UserId {
		return errors.New("禁止该操作！")
	}
	if req.DefaultFlag == 1 {
		var defaultUserAddress mall.MallUserAddress
		global.GVA_DB.Where("user_id=? and default_flag =1 and is_deleted = 0", userToken.UserId).First(&defaultUserAddress)
		if defaultUserAddress != (mall.MallUserAddress{}) {
			defaultUserAddress.DefaultFlag = 0
			defaultUserAddress.UpdateTime = common.JSONTime{time.Now()}
			err = global.GVA_DB.Save(&defaultUserAddress).Error
			if err != nil {
				return
			}
		}
	}
	err = copier.Copy(&userAddress, &req)
	if err != nil {
		return
	}
	userAddress.UpdateTime = common.JSONTime{time.Now()}
	userAddress.UserId = userToken.UserId
	err = global.GVA_DB.Save(&userAddress).Error
	return
}

func (m *MallUserAddressService) GetMallUserAddressById(token string, id int) (err error, userAddress mall.MallUserAddress) {
	var userToken mall.MallUserToken
	if err = global.GVA_DB.Where("token =?", token).First(&userToken).Error; err != nil {
		return errors.New("不存在的用户"), userAddress
	}
	if err = global.GVA_DB.Where("address_id =?", id).First(&userAddress).Error; err != nil {
		return errors.New("不存在的用户地址"), userAddress
	}
	if userToken.UserId != userAddress.UserId {
		return errors.New("禁止该操作！"), userAddress
	}
	return
}

func (m *MallUserAddressService) GetMallUserDefaultAddress(token string) (err error, userAddress mall.MallUserAddress) {
	var userToken mall.MallUserToken
	if err = global.GVA_DB.Where("token =?", token).First(&userToken).Error; err != nil {
		return errors.New("不存在的用户"), userAddress
	}
	if err = global.GVA_DB.Where("user_id =? and default_flag =1 and is_deleted = 0 ", userToken.UserId).First(&userAddress).Error; err != nil {
		return errors.New("不存在默认地址失败"), userAddress
	}
	return
}

func (m *MallUserAddressService) DeleteUserAddress(token string, id int) (err error) {
	var userToken mall.MallUserToken
	if err = global.GVA_DB.Where("token =?", token).First(&userToken).Error; err != nil {
		return errors.New("不存在的用户")
	}
	fmt.Println(id)
	var userAddress mall.MallUserAddress
	if err = global.GVA_DB.Where("address_id =?", id).First(&userAddress).Error; err != nil {
		return errors.New("不存在的用户地址")
	}
	if userToken.UserId != userAddress.UserId {
		return errors.New("禁止该操作！")
	}
	err = global.GVA_DB.Delete(&userAddress).Error
	return

}
