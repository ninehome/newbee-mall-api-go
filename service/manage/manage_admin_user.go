package manage

import (
	"errors"
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
	//这个方法只更新不为0值的
	//err = global.GVA_DB.Where("user_id = ?", req.UserId).Update(&manage.MallUser{
	//	UserMoney: req.UserMoney,
	//	UserLevel: req.UserLevel,
	//}).Error

	err = global.GVA_DB.Model(mall.MallUser{}).Where("user_id = ?", req.UserId).Updates(map[string]interface{}{"user_money": req.UserMoney, "user_level": req.UserLevel}).Error
	if err != nil {
		return errors.New("余额更新失败" + err.Error())
	}
	return err
}

// 更新提款状态
func (m *ManageAdminUserService) UpdateWithdrawal(token string, req manageReq.MallUpdateWithdrawalParam) (err error) {
	var userWithdraw mall.MallUserWithdraw
	var user mall.MallUser
	err = global.GVA_DB.Where("withdraw_id =? ", req.WithdrawId).First(&userWithdraw).Error
	if err != nil {
		return errors.New("不存在的提款订单：999" + err.Error())
	}

	//处理提款逻辑 如果是确认出款，不需要变化用户余额， 如果是驳回提款，需要退回用户余额
	if req.DealFlag == 2 {
		//驳回提款 需要把余额退回用户账户
		err = global.GVA_DB.Model(mall.MallUser{}).Where("user_id =? ", userWithdraw.UserId).First(&user).Error
		if err != nil {
			return errors.New("用户不存在：1000  userWithdraw.UserId= " + strconv.Itoa(userWithdraw.UserId) + err.Error())
		}

		var totalMoney = user.UserMoney + userWithdraw.WithdrawMoney
		err = global.GVA_DB.Model(mall.MallUser{}).Where("user_id = ?", userWithdraw.UserId).Updates(map[string]interface{}{"user_money": totalMoney}).Error
		if err != nil {
			return errors.New("更新余额失败：1001" + err.Error())
		}

	}

	err = global.GVA_DB.Model(mall.MallUserWithdraw{}).Where("withdraw_id = ?", req.WithdrawId).Updates(map[string]interface{}{"deal_flag": req.DealFlag}).Error
	if err != nil {
		return errors.New("更新失败,用户不存在1002" + err.Error())
	}
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

	//fmt.Println(req.IsDeleted)
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
			adminToken.AdminUserId = mallAdminUser.AdminUserId
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
			adminToken.AdminUserId = mallAdminUser.AdminUserId
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
