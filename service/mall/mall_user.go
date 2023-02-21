package mall

import (
	"errors"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"main.go/global"
	"main.go/model/common"
	"main.go/model/mall"
	mallReq "main.go/model/mall/request"
	mallRes "main.go/model/mall/response"
	"main.go/utils"
	"strconv"
	"strings"
	"time"
)

type MallUserService struct {
}

// RegisterUser 注册用户
func (m *MallUserService) RegisterUser(req mallReq.RegisterUserParam) (err error) {
	if !errors.Is(global.GVA_DB.Where("login_name =?", req.LoginName).First(&mall.MallUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("This number is already registered, please change it")
	}

	return global.GVA_DB.Create(&mall.MallUser{
		LoginName:     req.LoginName,
		PasswordMd5:   utils.MD5V([]byte(req.Password)),
		IntroduceSign: "....",
		CreateTime:    common.JSONTime{Time: time.Now()},
		AgentId:       req.AgentId,
	}).Error

}

// 提款
func (m *MallUserService) UserWithdrawal(token string, req mallReq.WithdrawalParam) (err error, userw mall.MallUserWithdraw) {
	var userToken mall.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("Non-existent users"), userw
	}

	var userInfo mall.MallUser
	err = global.GVA_DB.Where("user_id =?", userToken.UserId).First(&userInfo).Error
	if err != nil {
		return errors.New("Failed to retrieve user information"), userw
	}

	//赋值 写入数据库
	userw.WithdrawMoney = req.WithdrawMoney
	userw.CreateTime = common.JSONTime{Time: time.Now()}
	if userInfo.UserMoney < req.WithdrawMoney { //提款 大于 余额
		return errors.New("Insufficient balance"), userw
	}

	userw.UserMoney = userInfo.UserMoney - req.WithdrawMoney //余额
	userw.DealFlag = 0
	userw.BankId = req.BankId
	userw.UserId = userInfo.UserId
	userw.LoginName = userInfo.LoginName
	userw.AgentId = userInfo.AgentId

	err = global.GVA_DB.Create(&userw).Error
	if err != nil {
		return errors.New("Failed to generate a recall order " + err.Error()), userw
	}

	//userInfo.UserMoney = userw.UserMoney
	//err = global.GVA_DB.Where("user_id =?", userToken.UserId).First(&userInfo).Error

	//更新 账户余额
	err = global.GVA_DB.Model(&mall.MallUser{}).Where("user_id = ?", userInfo.UserId).Update("user_money", userw.UserMoney).Error
	if err != nil {
		return errors.New("Chargeback failed"), userw
	}

	return err, userw

}

func (m *MallUserService) UpdateUserInfo(token string, req mallReq.UpdateUserInfoParam) (err error) {
	var userToken mall.MallUserToken
	//err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	//if err != nil {
	//	return errors.New("Non-existent users")
	//}
	var userInfo mall.MallUser
	err = global.GVA_DB.Where("user_id =?", userToken.UserId).First(&userInfo).Error
	// 若密码为空字符，则表明用户不打算修改密码，使用原密码保存
	if !(req.PasswordMd5 == "") {
		userInfo.PasswordMd5 = utils.MD5V([]byte(req.PasswordMd5))
	}
	userInfo.NickName = req.NickName
	userInfo.IntroduceSign = req.IntroduceSign
	err = global.GVA_DB.Where("user_id =?", userToken.UserId).UpdateColumns(&userInfo).Error
	return
}

func (m *MallUserService) GetUserDetail(token string) (err error, userDetail mallRes.MallUserDetailResponse) {
	var userToken mall.MallUserToken
	err = global.GVA_DB.Where("token =?", token).First(&userToken).Error
	if err != nil {
		return errors.New("Non-existent users"), userDetail
	}
	var userInfo mall.MallUser
	err = global.GVA_DB.Where("user_id =?", userToken.UserId).First(&userInfo).Error
	if err != nil {
		return errors.New("Failed to retrieve user information"), userDetail
	}
	err = copier.Copy(&userDetail, &userInfo)
	return
}

func (m *MallUserService) UserLogin(params mallReq.UserLoginParam) (err error, user mall.MallUser, userToken mall.MallUserToken) {
	err = global.GVA_DB.Where("login_name=? AND password_md5=?", params.LoginName, params.PasswordMd5).First(&user).Error
	if user != (mall.MallUser{}) { //查询有这个用户
		token := getNewToken(time.Now().UnixNano()/1e6, int(user.UserId))
		global.GVA_DB.Where("user_id", user.UserId).First(&token)
		nowDate := time.Now()
		// 300小时过期
		expireTime, _ := time.ParseDuration("300h")
		expireDate := nowDate.Add(expireTime)
		// 没有token新增，有token 则更新
		if userToken == (mall.MallUserToken{}) {
			userToken.UserId = user.UserId
			userToken.Token = token
			userToken.UpdateTime = nowDate
			userToken.ExpireTime = expireDate
			userToken.AgentId = user.AgentId
			if err = global.GVA_DB.Save(&userToken).Error; err != nil {
				return
			}
		} else {
			//userToken.Token = token
			userToken.UpdateTime = nowDate
			userToken.ExpireTime = expireDate
			userToken.AgentId = user.AgentId
			if err = global.GVA_DB.Save(&userToken).Error; err != nil {
				return
			}
		}
		errS := global.GVA_DB.Where("user_id =?", user.UserId).First(&user).Error
		if errS != nil {

		}

		//更新 IP 地址
	}
	return err, user, userToken
}

func getNewToken(timeInt int64, userId int) (token string) {
	var build strings.Builder
	build.WriteString(strconv.FormatInt(timeInt, 10))
	build.WriteString(strconv.Itoa(userId))
	build.WriteString(utils.GenValidateCode(6))
	return utils.MD5V([]byte(build.String()))
}
