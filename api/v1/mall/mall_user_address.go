package mall

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common/response"
	mallReq "main.go/model/mall/request"
	"strconv"
)

type MallUserAddressApi struct {
}

func (m *MallUserAddressApi) AddressList(c *gin.Context) {
	token := c.GetHeader("token")
	if err, userAddressList := mallUserAddressService.GetMyAddress(token); err != nil {
		global.GVA_LOG.Error("获取地址失败", zap.Error(err))
		response.FailWithMessage("Não foi possível obter um endereço:"+err.Error(), c)
	} else if len(userAddressList) == 0 {
		response.OkWithData(nil, c)
	} else {
		response.OkWithData(userAddressList, c)
	}
}

func (m *MallUserAddressApi) SaveUserAddress(c *gin.Context) {
	var req mallReq.AddAddressParam
	_ = c.ShouldBindJSON(&req)
	byte, _ := json.Marshal(req)
	fmt.Println(string(byte))
	token := c.GetHeader("token")
	err := mallUserAddressService.SaveUserAddress(token, req)
	if err != nil {
		global.GVA_LOG.Error("创建失败", zap.Error(err))
		response.FailWithMessage("Falha ao criar:"+err.Error(), c)
		return
	}
	response.OkWithMessage("Criado com sucesso", c)

}

func (m *MallUserAddressApi) SaveUserBank(c *gin.Context) {
	var req mallReq.BankParam
	_ = c.ShouldBindJSON(&req)
	byte, _ := json.Marshal(req)
	fmt.Println(string(byte))
	token := c.GetHeader("token")
	err := mallUserAddressService.SaveUserBank(token, req)
	if err != nil {
		global.GVA_LOG.Error("添加银行账户失败", zap.Error(err))
		response.FailWithMessage("Falha ao adicionar uma conta bancária:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)

}

func (m *MallUserAddressApi) UpdateMallUserAddress(c *gin.Context) {
	var req mallReq.UpdateAddressParam
	_ = c.ShouldBindJSON(&req)
	token := c.GetHeader("token")
	err := mallUserAddressService.UpdateUserAddress(token, req)
	if err != nil {
		global.GVA_LOG.Error("更新用户地址失败", zap.Error(err))
		response.FailWithMessage("Falha ao atualizar o endereço do usuário:"+err.Error(), c)
	}
	response.OkWithMessage("Atualização bem-sucedida do endereço do usuário", c)
}

func (m *MallUserAddressApi) GetMallUserAddress(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("addressId"))
	token := c.GetHeader("token")
	if err, userAddress := mallUserAddressService.GetMallUserAddressById(token, id); err != nil {
		global.GVA_LOG.Error("获取地址失败", zap.Error(err))
		response.FailWithMessage("Não foi possível obter um endereço:"+err.Error(), c)
	} else {
		response.OkWithData(userAddress, c)
	}
}

func (m *MallUserAddressApi) GetMallUserDefaultAddress(c *gin.Context) {
	token := c.GetHeader("token")
	if err, userAddress := mallUserAddressService.GetMallUserDefaultAddress(token); err != nil {
		global.GVA_LOG.Error("获取地址失败", zap.Error(err))
		response.FailWithMessage("e conseguiu obter um endereço:"+err.Error(), c)
	} else {
		response.OkWithData(userAddress, c)
	}
}

func (m *MallUserAddressApi) DeleteUserAddress(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("addressId"))
	token := c.GetHeader("token")
	err := mallUserAddressService.DeleteUserAddress(token, id)
	if err != nil {
		global.GVA_LOG.Error("删除用户地址失败", zap.Error(err))
		response.FailWithMessage("Falha ao excluir o endereço do usuário:"+err.Error(), c)
	}
	response.OkWithMessage("A exclusão do endereço do usuário foi bem-sucedida", c)

}

func (m *MallUserAddressApi) DeleteUserBank(c *gin.Context) {
	id := c.Param("bankId")

	token := c.GetHeader("token")
	err := mallUserAddressService.DeleteUserBank(token, id)
	if err != nil {
		global.GVA_LOG.Error("删除用户银行账户", zap.Error(err))
		response.FailWithMessage("Excluir a conta bancária do usuário:"+err.Error(), c)
	}
	response.OkWithMessage("A exclusão da conta bancária do usuário foi bem-sucedida", c)

}
