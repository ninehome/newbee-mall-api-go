package mall

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common/response"
	mallReq "main.go/model/mall/request"
	manageReq "main.go/model/manage/request"
	"strconv"
	"sync"
)

type MallUserApi struct {
}

var once sync.Once

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
		response.FailWithMessage("Falha ao criar:"+err.Error(), c)
		return
	}

	response.OkWithMessage("Criado com sucesso", c)
}

// UserBankList 获取绑定银行卡列表
func (m *MallUserApi) UserBankList(c *gin.Context) {
	token := c.GetHeader("token")
	if err, userAddressList := mallUserAddressService.GetMyBankList(token); err != nil {
		global.GVA_LOG.Error("获取列bank表失败", zap.Error(err))
		response.FailWithMessage("Falha ao recuperar uma coluna da tabela do banco:"+err.Error(), c)
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
		response.FailWithMessage("Falha ao recuperar uma coluna da tabela do banco:"+err.Error(), c)
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
		response.FailWithMessage("Não foi possível obter um endereço:"+err.Error(), c)
	} else {
		response.OkWithData(userAddress, c)
	}
}

func (m *MallUserApi) GetUserMsg(c *gin.Context) {
	var req manageReq.MsgParam
	_ = c.ShouldBindJSON(&req)
	if err, userMsg := mallUserService.GetUserMsg(req); err != nil {
		global.GVA_LOG.Error("此用户无私信", zap.Error(err))
		response.FailWithMessage("此用户无私信:"+err.Error(), c)
	} else {
		response.OkWithData(userMsg, c)
	}
}

// UserWithdrawal 用户提款
func (m *MallUserApi) UserWithdrawal(c *gin.Context) {
	var req mallReq.WithdrawalParam
	_ = c.ShouldBindJSON(&req)
	token := c.GetHeader("token")
	if err, uw := mallUserService.UserWithdrawal(token, req); err != nil {
		global.GVA_LOG.Error("Cancelamento de retirada", zap.Error(err))
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
		response.FailWithMessage("Falha ao atualizar as informações do usuário"+err.Error(), c)
	}
	response.OkWithMessage("Renovação bem-sucedida", c)
}

func (m *MallUserApi) GetUserInfo(c *gin.Context) {
	token := c.GetHeader("token")
	if err, userDetail := mallUserService.GetUserDetail(token); err != nil {
		global.GVA_LOG.Error("未查询到记录", zap.Error(err))
		response.FailWithMessage("Os registros não foram examinados", c)
	} else {
		response.OkWithData(userDetail, c)
	}
}

func (m *MallUserApi) GetUserInfoV2(c *gin.Context) {
	var req mallReq.UserInfoParam
	_ = c.ShouldBindJSON(&req)
	if err, userDetail := mallUserService.GetUserDetailV2(req.UserId); err != nil {
		global.GVA_LOG.Error("未查询到记录", zap.Error(err))
		response.UnLogin("Os registros não foram examinados122"+err.Error(), c)
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
		response.FailWithPSW("Senha e número de conta incorretos inseridos", c)
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
		response.FailWithPSW("Senha e número de conta incorretos inseridos", c)
	} else {
		response.OkWithData(adminToken, c)
	}
}

func (m *MallUserApi) UserLogout(c *gin.Context) {
	token := c.GetHeader("token")
	if err := mallUserTokenService.DeleteMallUserToken(token); err != nil {
		response.FailWithMessage("Falha ao fazer logout", c)
	} else {
		response.OkWithMessage("Logout bem-sucedido", c)
	}

}
