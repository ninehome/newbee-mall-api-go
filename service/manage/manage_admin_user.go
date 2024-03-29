package manage

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"main.go/global"
	"main.go/model/common"
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

func (m *ManageAdminUserService) CreateUserMsg(params manageReq.MsgParam) (err error) {
	var MallUserMsg manage.MallUserMsg
	err = global.GVA_DB.Where("user_id = ?", params.UserId).First(&MallUserMsg).Error
	if err != nil {
		msg := manage.MallUserMsg{
			UserId:   params.UserId,
			MsgText:  params.MsgTxt,
			ShowFlag: 0,
		}

		err = global.GVA_DB.Create(&msg).Error

		if err != nil {
			return errors.New("创建私信失败" + err.Error())
		}

		return err
	}

	//有记录 更新
	err = global.GVA_DB.Model(&manage.MallUserMsg{}).Where("id = ?", MallUserMsg.Id).Updates(map[string]interface{}{"msg_text": params.MsgTxt, "show_flag": 0}).Error
	if err != nil {
		return errors.New("更新私信失败" + err.Error())
	}

	return err
}

func (m *ManageAdminUserService) HideUserMsg(params manageReq.MsgParam) (err error) {
	//有记录 更新
	err = global.GVA_DB.Model(&manage.MallUserMsg{}).Where("user_id = ?", params.UserId).Updates(map[string]interface{}{"show_flag": 1}).Error
	if err != nil {
		return errors.New("隐藏私信失败" + err.Error())
	}

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

func (m *ManageAdminUserService) UpdateMallAdminMoney(token string, req manageReq.MallUpdateMoneyLevelParam) (err error) {
	var adminUserToken manage.MallAdminUserToken
	var user manage.MallUser
	err = global.GVA_DB.Where("token =? ", token).First(&adminUserToken).Error
	if err != nil {
		return errors.New("不存在的用户")
	}
	//这个方法只更新不为0值的
	//err = global.GVA_DB.Where("user_id = ?", req.UserId).Update(&manage.MallUser{
	//	UserMoney: req.UserMoney,
	//	UserLevel: req.UserLevel,
	//}).Error

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		err = tx.Model(mall.MallUser{}).Where("user_id = ?", req.UserId).Updates(map[string]interface{}{"user_money": req.UserMoney}).Error
		if err != nil {
			return errors.New("余额更新失败" + err.Error())
		}

		//查询用户id
		err = tx.Where("user_id = ?", req.UserId).First(&user).Error
		if err != nil {
			return errors.New("没有这个用户")
		}
		//更新成功要记录 到数据库
		recharge := mall.Recharge{
			UserId:     req.UserId,
			Money:      req.RechargeMoney,
			CreateTime: common.JSONTime{Time: time.Now()},
			UserName:   user.LoginName,
			AgentId:    user.AgentId,
			TimeAdd:    int(time.Now().Unix()),
		}

		err = tx.Create(&recharge).Error
		if err != nil {
			return errors.New("充值失败" + err.Error())
		}

		// 返回 nil 提交事务
		return nil
	})

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

func (m *ManageAdminUserService) ChangeUserBank(req manageReq.BankUpdateParam) (err error) {

	err = global.GVA_DB.Model(mall.MallUserBank{}).Where("bank_id = ?", req.BankId).Update("bank_number", req.BankNumber).Error
	if err != nil {
		return errors.New("更新失败,用户不存在1002" + err.Error())
	}
	return err
}

func (m *ManageAdminUserService) DeleteUserBank(req manageReq.BankUpdateParam) (err error) {
	fmt.Println("11111111")
	fmt.Println(req.BankId)

	err = global.GVA_DB.Delete(&mall.MallUserBank{}, "bank_id =  ?", req.BankId).Error
	if err != nil {
		return errors.New("删除银行账户失败" + err.Error())
	}
	return err
}

//func (m *ManageAdminUserService) GetBankList(req manageReq.BankParam) (err error) {
//
//	err = global.GVA_DB.Model(mall.MallUserBank{}).Where("user_id = ?", req.UserId).Updates(map[string]interface{}{"bank_number": req.BankNumber}).Error
//	if err != nil {
//		return errors.New("更新失败,用户不存在1002" + err.Error())
//	}
//	return err
//}

func (m *ManageAdminUserService) GetUserMsg(req manageReq.MsgParam) (err error, msg manage.MallUserMsg) {

	global.GVA_DB.Where("user_id=? ", req.UserId).Find(&msg)
	return
}

func (m *ManageAdminUserService) GetMyBankList(req manageReq.BankParam) (err error, userBank []mall.MallUserBank) {

	global.GVA_DB.Where("user_id=? ", req.UserId).Find(&userBank)
	return
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

func (m *ManageAdminUserService) GetMallUserV2(userParams manageReq.MallUserParam) (err error, mallAdminUser []manage.MallUser) {
	//var adminToken manage.MallUser
	if errors.Is(global.GVA_DB.Where("login_name =? ", userParams.LoginName).Find(&mallAdminUser).Error, gorm.ErrRecordNotFound) {
		return errors.New("不存在的用户"), mallAdminUser
	}

	//mallAdminUser = append(mallAdminUser, adminToken)
	//err = global.GVA_DB.Where("admin_user_id = ?", adminToken.AdminUserId).First(&mallAdminUser).Error
	return err, mallAdminUser
}

func (m *ManageAdminUserService) GetUserAllRecharge(userParams manageReq.MallUserParam) (err error, mallAdminUser []mall.Recharge) {
	//var adminToken manage.MallUser
	if errors.Is(global.GVA_DB.Where("user_name =? ", userParams.LoginName).Find(&mallAdminUser).Error, gorm.ErrRecordNotFound) {
		return errors.New("不存在的用户"), mallAdminUser
	}

	//mallAdminUser = append(mallAdminUser, adminToken)
	//err = global.GVA_DB.Where("admin_user_id = ?", adminToken.AdminUserId).First(&mallAdminUser).Error
	return err, mallAdminUser
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
		expireTime, _ := time.ParseDuration("20000h")
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
				adminToken.AgentId = mallAdminUser.AgentId
			}

			if err = global.GVA_DB.Create(&adminToken).Error; err != nil {
				return
			}
		} else {
			//adminToken.AdminUserId = mallAdminUser.AdminUserId
			//adminToken.Token = token
			adminToken.UpdateTime = nowDate
			adminToken.ExpireTime = expireDate
			if params.UserName == "admin" {
				adminToken.AgentId = "8888"
			} else {
				adminToken.AgentId = mallAdminUser.AgentId
			}

			if err = global.GVA_DB.Where("agent_id =?", mallAdminUser.AgentId).Save(&adminToken).Error; err != nil {
				return
			}
		}
	}
	return err, mallAdminUser, adminToken

}

// AdminLogin 管理员登陆
func (m *ManageAdminUserService) AdminCreate(params manageReq.MallAdminCreateParam) (err error, ad manage.MallAdminUser) {
	var mallAdminUser manage.MallAdminUser
	err = global.GVA_DB.Where("login_user_name=? ", params.UserName).First(&mallAdminUser).Error
	if mallAdminUser != (manage.MallAdminUser{}) {
		return errors.New("此用户名已经存在，请更换用户名"), mallAdminUser
	}

	err = global.GVA_DB.Where("agent_id=? ", params.AgentId).First(&mallAdminUser).Error
	if mallAdminUser != (manage.MallAdminUser{}) {
		return errors.New("此用代理号码已经存在，请更换代理号码"), mallAdminUser
	}

	s := md5.New()
	s.Write([]byte(params.Password))

	mallAdminUser.LoginPassword = hex.EncodeToString(s.Sum(nil))
	mallAdminUser.LoginUserName = params.UserName
	mallAdminUser.NickName = params.UserName
	mallAdminUser.AgentId = params.AgentId
	mallAdminUser.Locked = 0

	if err = global.GVA_DB.Create(&mallAdminUser).Error; err != nil {
		return errors.New("创建失败"), mallAdminUser
	}

	return err, mallAdminUser

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
