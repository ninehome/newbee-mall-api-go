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
		response.FailWithMessage("Не удалось получить тележку:"+err.Error(), c)
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
		response.FailWithMessage("Не удалось добавить в корзину:"+err.Error(), c)
	}
	response.OkWithMessage("Добавить в корзину успешно", c)
}

func (m *MallShopCartApi) UpdateMallShoppingCartItem(c *gin.Context) {
	token := c.GetHeader("token")
	var req mallReq.UpdateCartItemParam
	_ = c.ShouldBindJSON(&req)
	if err := mallShopCartService.UpdateMallCartItem(token, req); err != nil {
		global.GVA_LOG.Error("修改购物车失败", zap.Error(err))
		response.FailWithMessage("Не удалось модифицировать корзину:"+err.Error(), c)
	}
	response.OkWithMessage("Успешное изменение корзины", c)
}

func (m *MallShopCartApi) DelMallShoppingCartItem(c *gin.Context) {
	token := c.GetHeader("token")
	id, _ := strconv.Atoi(c.Param("newBeeMallShoppingCartItemId"))
	if err := mallShopCartService.DeleteMallCartItem(token, id); err != nil {
		global.GVA_LOG.Error("修改购物车失败", zap.Error(err))
		response.FailWithMessage("Не удалось модифицировать корзину:"+err.Error(), c)
	}
	response.OkWithMessage("Успешное изменение корзины", c)
}

func (m *MallShopCartApi) ToSettle(c *gin.Context) {
	cartItemIdsStr := c.Query("cartItemIds")
	token := c.GetHeader("token")
	cartItemIds := utils.StrToInt(cartItemIdsStr)
	if err, cartItemRes := mallShopCartService.GetCartItemsForSettle(token, cartItemIds); err != nil {
		global.GVA_LOG.Error("获取购物明细异常：", zap.Error(err))
		response.FailWithMessage("Исключения из получения данных о покупке:"+err.Error(), c)
	} else {
		response.OkWithData(cartItemRes, c)
	}

}
