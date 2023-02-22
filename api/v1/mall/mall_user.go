package mall

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common/response"
	mallReq "main.go/model/mall/request"
	"strconv"
)

type MallUserApi struct {
}

func (m *MallUserApi) UserRegister(c *gin.Context) {
	var req mallReq.RegisterUserParam
	_ = c.ShouldBindJSON(&req)

	//验证 输入是否合法
	//if err := utils.Verify(req, utils.MallUserRegisterVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	if err := mallUserService.RegisterUser(req); err != nil {
		global.GVA_LOG.Error("创建失败", zap.Error(err))
		response.FailWithMessage("Не удалось создать:"+err.Error(), c)
		return
	}
	response.OkWithMessage("Создано успешно", c)
}

// UserBankList 获取绑定银行卡列表
func (m *MallUserApi) UserBankList(c *gin.Context) {
	token := c.GetHeader("token")
	if err, userAddressList := mallUserAddressService.GetMyBankList(token); err != nil {
		global.GVA_LOG.Error("获取列bank表失败", zap.Error(err))
		response.FailWithMessage("Не удалось получить столбец банковской таблицы:"+err.Error(), c)
	} else if len(userAddressList) == 0 {
		response.OkWithData(nil, c)
	} else {
		response.OkWithData(userAddressList, c)
	}
}

// 获取联系方式
func (m *MallUserApi) UserChatList(c *gin.Context) {
	token := c.GetHeader("token")
	if err, userAddressList := mallUserAddressService.GetChatList(token); err != nil {
		global.GVA_LOG.Error("获取列bank表失败", zap.Error(err))
		response.FailWithMessage("Не удалось получить столбец банковской таблицы:"+err.Error(), c)
	} else if len(userAddressList) == 0 {
		response.OkWithData(nil, c)
	} else {
		response.OkWithData(userAddressList, c)
	}
}

func (m *MallUserApi) GetMallUserBank(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("bankId"))
	token := c.GetHeader("token")
	if err, userAddress := mallUserAddressService.GetMyBank(token, id); err != nil {
		global.GVA_LOG.Error("获取地址失败", zap.Error(err))
		response.FailWithMessage("Не удалось получить адрес:"+err.Error(), c)
	} else {
		response.OkWithData(userAddress, c)
	}
}

// UserWithdrawal 用户提款
func (m *MallUserApi) UserWithdrawal(c *gin.Context) {
	var req mallReq.WithdrawalParam
	_ = c.ShouldBindJSON(&req)
	token := c.GetHeader("token")
	if err, uw := mallUserService.UserWithdrawal(token, req); err != nil {
		global.GVA_LOG.Error("Отказ от вывода средств", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(uw, c)
	}

	//var req mallReq.RegisterUserParam
	//_ = c.ShouldBindJSON(&req)
	//if err := utils.Verify(req, utils.MallUserRegisterVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	//if err := mallUserService.RegisterUser(req); err != nil {
	//	global.GVA_LOG.Error("创建失败", zap.Error(err))
	//	response.FailWithMessage("创建失败:"+err.Error(), c)
	//}
	//response.OkWithMessage("创建成功", c)
}

func (m *MallUserApi) UserInfoUpdate(c *gin.Context) {
	var req mallReq.UpdateUserInfoParam
	token := c.GetHeader("token")
	if err := mallUserService.UpdateUserInfo(token, req); err != nil {
		global.GVA_LOG.Error("更新用户信息失败", zap.Error(err))
		response.FailWithMessage("Не удалось обновить информацию о пользователе"+err.Error(), c)
	}
	response.OkWithMessage("Успешное обновление", c)
}

func (m *MallUserApi) GetUserInfo(c *gin.Context) {
	token := c.GetHeader("token")
	if err, userDetail := mallUserService.GetUserDetail(token); err != nil {
		global.GVA_LOG.Error("未查询到记录", zap.Error(err))
		response.FailWithMessage("Записи не изучались", c)
	} else {
		response.OkWithData(userDetail, c)
	}
}

func (m *MallUserApi) GetUserInfoV2(c *gin.Context) {
	//token := c.GetHeader("token")
	userId := c.GetHeader("userId")
	if err, userDetail := mallUserService.GetUserDetailV2(userId); err != nil {
		global.GVA_LOG.Error("未查询到记录", zap.Error(err))
		response.FailWithMessage("Записи не изучались", c)
	} else {
		response.OkWithData(userDetail, c)
	}
}

func (m *MallUserApi) UserLogin(c *gin.Context) {
	var req mallReq.UserLoginParam
	_ = c.ShouldBindJSON(&req)

	reqIP := c.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	req.UserIpAddr = reqIP

	if err, _, adminToken := mallUserService.UserLogin(req); err != nil {
		response.FailWithPSW("Введен неправильный пароль и номер счета", c)
	} else {
		response.OkWithData(adminToken.Token, c)
	}
}
func (m *MallUserApi) UserLoginV2(c *gin.Context) {
	var req mallReq.UserLoginParam
	_ = c.ShouldBindJSON(&req)

	reqIP := c.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	req.UserIpAddr = reqIP

	if err, _, adminToken := mallUserService.UserLogin(req); err != nil {
		response.FailWithPSW("Введен неправильный пароль и номер счета", c)
	} else {
		response.OkWithData(adminToken, c)
	}
}

func (m *MallUserApi) UserLogout(c *gin.Context) {
	token := c.GetHeader("token")
	if err := mallUserTokenService.DeleteMallUserToken(token); err != nil {
		response.FailWithMessage("Не удалось выйти из системы", c)
	} else {
		response.OkWithMessage("Успешный выход из системы", c)
	}

}
