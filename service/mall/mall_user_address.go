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
		return errors.New("Несуществующие пользователи"), userAddress
	}
	global.GVA_DB.Where("user_id=? and is_deleted=0", userToken.UserId).Find(&userAddress)
	return
}

// GetMyAddress 获取银行列表
func (m *MallUserAddressService) GetMyBankList(token string) (err error, userBank []mall.MallUserBank) {
	var userToken mall.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("Несуществующие потребители"), userBank
	}
	global.GVA_DB.Where("user_id=? ", userToken.UserId).Find(&userBank)
	return
}

// GetMyAddress 联系方式
func (m *MallUserAddressService) GetChatList(token string) (err error, userBank []mall.MallUserChat) {
	//var userToken mall.MallUserToken
	//err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	//if err != nil {
	//	return errors.New("Несуществующие потребители"), userBank
	//}
	global.GVA_DB.Where("is_delete = 0").Find(&userBank)
	return
}

// 保存 银行账户
func (m *MallUserAddressService) SaveUserBank(token string, req mallReq.BankParam) (err error) {
	var userToken mall.MallUserToken
	if err = global.GVA_DB.Where("token =?", token).First(&userToken).Error; err != nil {
		return errors.New("Несуществующие пользователи")
	}
	var defaultAddress mall.MallUserBank
	var newAddress mall.MallUserBank
	// 是否新增了默认地址，将之前的默认地址设置为非默认

	if err = global.GVA_DB.Where("user_id=? and default =1 and is_deleted = 0", userToken.UserId).First(&defaultAddress).Error; err != nil {
		//没有查到记录 ,新增记录
		copier.Copy(&newAddress, &req)
		//newAddress.CreateTime = common.JSONTime{Time: time.Now()}
		//newAddress.UpdateTime = common.JSONTime{Time: time.Now()}
		newAddress.UserId = userToken.UserId
		err = global.GVA_DB.Create(&newAddress).Error
		if err != nil {
			return
		}
	} else {
		//先更新 之前的记录 再新增
		global.GVA_DB.Model(&mall.MallUserAddress{}).Where("bank_id =?", defaultAddress.BankId).Update("default", 0)

		copier.Copy(&newAddress, &req)
		//newAddress.CreateTime = common.JSONTime{Time: time.Now()}
		//newAddress.UpdateTime = common.JSONTime{Time: time.Now()}
		newAddress.UserId = userToken.UserId
		err = global.GVA_DB.Create(&newAddress).Error
		if err != nil {
			return
		}

	}

	return
}

// SaveUserAddress 保存用户地址
func (m *MallUserAddressService) SaveUserAddress(token string, req mallReq.AddAddressParam) (err error) {
	var userToken mall.MallUserToken
	if err = global.GVA_DB.Where("token =?", token).First(&userToken).Error; err != nil {
		return errors.New("Несуществующие пользователи")
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
		return errors.New("Несуществующие пользователи")
	}
	var userAddress mall.MallUserAddress
	if err = global.GVA_DB.Where("address_id =? and user_id =?", req.AddressId, userToken.UserId).First(&userAddress).Error; err != nil {
		return errors.New("Несуществующий адрес пользователя")
	}
	if userToken.UserId != userAddress.UserId {
		return errors.New("Отключить эту операцию！")
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
		return errors.New("Несуществующие пользователи"), userAddress
	}
	if err = global.GVA_DB.Where("address_id =?", id).First(&userAddress).Error; err != nil {
		return errors.New("Несуществующий адрес пользователя"), userAddress
	}
	if userToken.UserId != userAddress.UserId {
		return errors.New("Отключить эту операцию！"), userAddress
	}
	return
}

// 获取单个银行账户
func (m *MallUserAddressService) GetMyBank(token string, id int) (err error, userAddress mall.MallUserBank) {
	var userToken mall.MallUserToken
	if err = global.GVA_DB.Where("token =?", token).First(&userToken).Error; err != nil {
		return errors.New("Несуществующие пользователи"), userAddress
	}
	if err = global.GVA_DB.Where("bank_id =?", id).First(&userAddress).Error; err != nil {
		return errors.New("Отсутствие банковского счета" + err.Error()), userAddress
	}
	if userToken.UserId != userAddress.UserId {
		return errors.New("Отключить эту операцию！"), userAddress
	}
	return
}

func (m *MallUserAddressService) GetMallUserDefaultAddress(token string) (err error, userAddress mall.MallUserAddress) {
	var userToken mall.MallUserToken
	if err = global.GVA_DB.Where("token =?", token).First(&userToken).Error; err != nil {
		return errors.New("Несуществующие пользователи"), userAddress
	}
	if err = global.GVA_DB.Where("user_id =? and default_flag =1 and is_deleted = 0 ", userToken.UserId).First(&userAddress).Error; err != nil {
		return errors.New("Адрес по умолчанию отсутствует сбой"), userAddress
	}
	return
}

func (m *MallUserAddressService) DeleteUserAddress(token string, id int) (err error) {
	var userToken mall.MallUserToken
	if err = global.GVA_DB.Where("token =?", token).First(&userToken).Error; err != nil {
		return errors.New("Несуществующие пользователи")
	}
	fmt.Println(id)
	var userAddress mall.MallUserAddress
	if err = global.GVA_DB.Where("address_id =?", id).First(&userAddress).Error; err != nil {
		return errors.New("Несуществующий адрес пользователя")
	}
	if userToken.UserId != userAddress.UserId {
		return errors.New("Отключить эту операцию！")
	}
	err = global.GVA_DB.Delete(&userAddress).Error
	return

}

func (m *MallUserAddressService) DeleteUserBank(token string, id string) (err error) {
	var userToken mall.MallUserToken
	if err = global.GVA_DB.Where("token =?", token).First(&userToken).Error; err != nil {
		return errors.New("Несуществующие пользователи")
	}
	fmt.Println(id)
	var userBank mall.MallUserBank
	if err = global.GVA_DB.Where("bank_id =?", id).First(&userBank).Error; err != nil {
		return errors.New("Несуществующие банковские счета")
	}
	if userToken.UserId != userBank.UserId {
		return errors.New("Отключить эту операцию！")
	}
	err = global.GVA_DB.Delete(&userBank).Error
	if err != nil {
		return errors.New("Не удалось удалить" + err.Error())
	}
	return

}
