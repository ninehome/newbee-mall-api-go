package manage

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"main.go/global"
	"main.go/model/mall"
	"main.go/model/manage"
	manageReq "main.go/model/manage/request"
	"main.go/utils"
	"strconv"
	"strings"
	"time"
)

type ManageAdminUserService struct {
}

// CreateMallAdminUser 创建MallAdminUser记录
func (m *ManageAdminUserService) CreateMallAdminUser(mallAdminUser manage.MallAdminUser) (err error) {
	if !errors.Is(global.GVA_DB.Where("login_user_name = ?", mallAdminUser.LoginUserName).First(&manage.MallAdminUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同用户名")
	}
	err = global.GVA_DB.Create(&mallAdminUser).Error
	return err
}

// UpdateMallAdminName 更新MallAdminUser昵称
func (m *ManageAdminUserService) UpdateMallAdminName(token string, req manageReq.MallUpdateNameParam) (err error) {
	var adminUserToken manage.MallAdminUserToken
	err = global.GVA_DB.Where("token =? ", token).First(&adminUserToken).Error
	if err != nil {
		return errors.New("不存在的用户")
	}
	err = global.GVA_DB.Where("admin_user_id = ?", adminUserToken.AdminUserId).Updates(&manage.MallAdminUser{
		LoginUserName: req.LoginUserName,
		NickName:      req.NickName,
	}).Error
	return err
}

func (m *ManageAdminUserService) UpdateMallUserPsw(token string, req manageReq.MallUpdatePswParam) (err error) {
	//var adminUserToken manage.MallAdminUserToken
	//err = global.GVA_DB.Where("token =? ", token).First(&adminUserToken).Error
	//if err != nil {
	//	return errors.New("不存在的用户")
	//}
	err = global.GVA_DB.Where("user_id = ?", req.UserId).Updates(&manage.MallUser{
		PasswordMd5: req.PasswordMd5,
	}).Error
	return err
}

func (m *ManageAdminUserService) UpdateMallAdminMoneyAndLevel(token string, req manageReq.MallUpdateMoneyLevelParam) (err error) {
	var adminUserToken manage.MallAdminUserToken
	err = global.GVA_DB.Where("token =? ", token).First(&adminUserToken).Error
	if err != nil {
		return errors.New("不存在的用户")
	}
	err = global.GVA_DB.Where("user_id = ?", req.UserId).Updates(&manage.MallUser{
		UserMoney: req.UserMoney,
		UserLevel: req.UserLevel,
	}).Error
	return err
}

func (m *ManageAdminUserService) UpdateMallChat(token string, req manageReq.MallUpdateChatParam) (err error) {
	//var adminUserToken manage.MallAdminUserToken
	//err = global.GVA_DB.Where("token =? ", token).First(&adminUserToken).Error
	//if err != nil {
	//	return errors.New("不存在的用户")
	//}

	//err = global.GVA_DB.Where("chat_id = ?", req.UserId).Updates(&manage.MallUser{
	//	UserMoney: req.UserMoney,
	//	UserLevel: req.UserLevel,
	//}).Error

	fmt.Println(req.IsDeleted)
	err = global.GVA_DB.Where("chat_id = ?", req.ChatId).Updates(&mall.MallUserChat{
		ChatValue: req.ChatValue,
		IsDeleted: req.IsDeleted,
	}).Error
	return err
}

func (m *ManageAdminUserService) UpdateMallAdminPassWord(token string, req manageReq.MallUpdatePasswordParam) (err error) {
	var adminUserToken manage.MallAdminUserToken
	err = global.GVA_DB.Where("token =? ", token).First(&adminUserToken).Error
	if err != nil {
		return errors.New("用户未登录")
	}
	var adminUser manage.MallAdminUser
	err = global.GVA_DB.Where("admin_user_id =?", adminUserToken.AdminUserId).First(&adminUser).Error
	if err != nil {
		return errors.New("不存在的用户")
	}
	if adminUser.LoginPassword != req.OriginalPassword {
		return errors.New("原密码不正确")
	}
	adminUser.LoginPassword = req.NewPassword

	err = global.GVA_DB.Where("admin_user_id=?", adminUser.AdminUserId).Updates(&adminUser).Error
	return
}

// GetMallAdminUser 根据id获取MallAdminUser记录
func (m *ManageAdminUserService) GetMallAdminUser(token string) (err error, mallAdminUser manage.MallAdminUser) {
	var adminToken manage.MallAdminUserToken
	if errors.Is(global.GVA_DB.Where("token =?", token).First(&adminToken).Error, gorm.ErrRecordNotFound) {
		return errors.New("不存在的用户"), mallAdminUser
	}
	err = global.GVA_DB.Where("admin_user_id = ?", adminToken.AdminUserId).First(&mallAdminUser).Error
	return err, mallAdminUser
}

func (m *ManageAdminUserService) GetMallUser(id string) (err error, mallAdminUser manage.MallUser) {
	var adminToken manage.MallUser
	if errors.Is(global.GVA_DB.Where("user_id =? ", id).First(&adminToken).Error, gorm.ErrRecordNotFound) {
		return errors.New("不存在的用户"), adminToken
	}
	//err = global.GVA_DB.Where("admin_user_id = ?", adminToken.AdminUserId).First(&mallAdminUser).Error
	return err, adminToken
}

func (m *ManageAdminUserService) GetMallChat(id string) (err error, mallChat mall.MallUserChat) {

	if errors.Is(global.GVA_DB.Where("chat_id =? ", id).First(&mallChat).Error, gorm.ErrRecordNotFound) {
		return errors.New("查询不到这个客服方式"), mallChat
	}
	//err = global.GVA_DB.Where("admin_user_id = ?", adminToken.AdminUserId).First(&mallAdminUser).Error
	return err, mallChat
}

// AdminLogin 管理员登陆
func (m *ManageAdminUserService) AdminLogin(params manageReq.MallAdminLoginParam) (err error, ad manage.MallAdminUser, adminToken manage.MallAdminUserToken) {
	var mallAdminUser manage.MallAdminUser
	err = global.GVA_DB.Where("login_user_name=? AND login_password=?", params.UserName, params.PasswordMd5).First(&mallAdminUser).Error
	if mallAdminUser != (manage.MallAdminUser{}) {
		token := getNewToken(time.Now().UnixNano()/1e6, int(mallAdminUser.AdminUserId))
		global.GVA_DB.Where("agent_id =?", mallAdminUser.AgentId).First(&adminToken)
		nowDate := time.Now()
		// 48小时过期
		expireTime, _ := time.ParseDuration("200h")
		expireDate := nowDate.Add(expireTime)
		// 没有token新增，有token 则更新
		if adminToken == (manage.MallAdminUserToken{}) {

			adminToken.Token = token
			adminToken.UpdateTime = nowDate
			adminToken.ExpireTime = expireDate
			if params.UserName == "admin" {
				adminToken.AgentId = "8888"
			} else {
				adminToken.AgentId = params.UserName
			}

			if err = global.GVA_DB.Create(&adminToken).Error; err != nil {
				return
			}
		} else {
			adminToken.Token = token
			adminToken.UpdateTime = nowDate
			adminToken.ExpireTime = expireDate
			if params.UserName == "admin" {
				adminToken.AgentId = "8888"
			} else {
				adminToken.AgentId = params.UserName
			}

			if err = global.GVA_DB.Save(&adminToken).Error; err != nil {
				return
			}
		}
	}
	return err, mallAdminUser, adminToken

}

// GetMyAddress 联系方式
func (m *ManageAdminUserService) GetChatList(token string) (err error, userBank []mall.MallUserChat) {
	//var userToken mall.MallUserToken
	//err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	//if err != nil {
	//	return errors.New("Несуществующие потребители"), userBank
	//}
	global.GVA_DB.Find(&userBank)
	return
}

func getNewToken(timeInt int64, userId int) (token string) {
	var build strings.Builder
	build.WriteString(strconv.FormatInt(timeInt, 10))
	build.WriteString(strconv.Itoa(userId))
	build.WriteString(utils.GenValidateCode(6))
	return utils.MD5V([]byte(build.String()))
}
