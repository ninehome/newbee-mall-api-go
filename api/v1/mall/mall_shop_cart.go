package mall

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common/response"
	mallReq "main.go/model/mall/request"
	"main.go/utils"
	"strconv"
)

type MallShopCartApi struct {
}

func (m *MallShopCartApi) CartItemList(c *gin.Context) {
	token := c.GetHeader("token")
	if err, shopCartItem := mallShopCartService.GetMyShoppingCartItems(token); err != nil {
		global.GVA_LOG.Error("获取购物车失败", zap.Error(err))
		response.FailWithMessage("Não foi possível obter um carrinho:"+err.Error(), c)
	} else {
		response.OkWithData(shopCartItem, c)
	}
}

func (m *MallShopCartApi) SaveMallShoppingCartItem(c *gin.Context) {
	token := c.GetHeader("token")
	var req mallReq.SaveCartItemParam
	_ = c.ShouldBindJSON(&req)
	if err := mallShopCartService.SaveMallCartItem(token, req); err != nil {
		global.GVA_LOG.Error("添加购物车失败", zap.Error(err))
		response.FailWithMessage("Não foi possível adicionar à cesta:"+err.Error(), c)
	}
	response.OkWithMessage("Adicionar à cesta com sucesso", c)
}

func (m *MallShopCartApi) UpdateMallShoppingCartItem(c *gin.Context) {
	token := c.GetHeader("token")
	var req mallReq.UpdateCartItemParam
	_ = c.ShouldBindJSON(&req)
	if err := mallShopCartService.UpdateMallCartItem(token, req); err != nil {
		global.GVA_LOG.Error("修改购物车失败", zap.Error(err))
		response.FailWithMessage("Falha ao modificar o carrinho de compras:"+err.Error(), c)
	}
	response.OkWithMessage("Mudança de cesta bem-sucedida", c)
}

func (m *MallShopCartApi) DelMallShoppingCartItem(c *gin.Context) {
	token := c.GetHeader("token")
	id, _ := strconv.Atoi(c.Param("newBeeMallShoppingCartItemId"))
	if err := mallShopCartService.DeleteMallCartItem(token, id); err != nil {
		global.GVA_LOG.Error("修改购物车失败", zap.Error(err))
		response.FailWithMessage("Falha ao modificar o carrinho de compras:"+err.Error(), c)
	}
	response.OkWithMessage("Mudança de cesta bem-sucedida", c)
}

func (m *MallShopCartApi) ToSettle(c *gin.Context) {
	cartItemIdsStr := c.Query("cartItemIds")
	token := c.GetHeader("token")
	cartItemIds := utils.StrToInt(cartItemIdsStr)
	if err, cartItemRes := mallShopCartService.GetCartItemsForSettle(token, cartItemIds); err != nil {
		global.GVA_LOG.Error("获取购物明细异常：", zap.Error(err))
		response.FailWithMessage("Exceções à obtenção de dados de compra:"+err.Error(), c)
	} else {
		response.OkWithData(cartItemRes, c)
	}

}
