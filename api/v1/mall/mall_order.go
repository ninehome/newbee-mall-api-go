package mall

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common/response"
	"main.go/model/mall"
	mallReq "main.go/model/mall/request"
	"main.go/utils"
	"strconv"
)

type MallOrderApi struct {
}

func (m *MallOrderApi) SaveOrder(c *gin.Context) {
	var saveOrderParam mallReq.SaveOrderParam
	_ = c.ShouldBindJSON(&saveOrderParam)
	if err := utils.Verify(saveOrderParam, utils.SaveOrderParamVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	token := c.GetHeader("token")

	priceTotal := 0
	err, itemsForSave := mallShopCartService.GetCartItemsForSettle(token, saveOrderParam.CartItemIds)
	if len(itemsForSave) < 1 {
		response.FailWithMessage("Não há dados:"+err.Error(), c)
	} else {
		//总价
		for _, newBeeMallShoppingCartItemVO := range itemsForSave {
			priceTotal = priceTotal + newBeeMallShoppingCartItemVO.GoodsCount*newBeeMallShoppingCartItemVO.SellingPrice
		}
		if priceTotal < 1 {
			response.FailWithMessage("Anomalias de preço", c)
			return
		}
		errAddress, userAddress := mallUserAddressService.GetMallUserDefaultAddress(token) //查询用户收货地址是否存在，不存在的不能支付

		if userAddress == (mall.MallUserAddress{}) {
			response.FailWithMessage("O endereço de entrega não está preenchido:"+errAddress.Error(), c)
			return

		}

		if err, saveOrderResult := mallOrderService.SaveOrder(token, userAddress, itemsForSave); err != nil {
			global.GVA_LOG.Error("生成订单失败", zap.Error(err)) //写入日志
			response.FailWithMessage("Falha ao formar uma ordem:"+err.Error(), c)
		} else {
			response.OkWithData(saveOrderResult, c)
		}
	}
}

// 商品支付
func (m *MallOrderApi) PaySuccess(c *gin.Context) {
	orderNo := c.Query("orderNo")
	payType, _ := strconv.Atoi(c.Query("payType"))
	if err := mallOrderService.PaySuccess(orderNo, payType); err != nil {
		global.GVA_LOG.Error("订单支付失败", zap.Error(err))
		response.FailWithMessage("Falha no pagamento do pedido:"+err.Error(), c)
		return
	}
	response.OkWithMessage("Pedido pago com sucesso", c)
}

func (m *MallOrderApi) FinishOrder(c *gin.Context) {
	orderNo := c.Param("orderNo")
	token := c.GetHeader("token")
	if err := mallOrderService.FinishOrder(token, orderNo); err != nil {
		global.GVA_LOG.Error("申请回购失败", zap.Error(err))
		response.FailWithMessage("Pedido de recompra não cumprido:"+err.Error(), c)
	}
	response.OkWithMessage("Oferta de compra bem-sucedida", c)

}

func (m *MallOrderApi) CancelOrder(c *gin.Context) {
	orderNo := c.Param("orderNo")
	token := c.GetHeader("token")
	if err := mallOrderService.CancelOrder(token, orderNo); err != nil {
		global.GVA_LOG.Error("订单签收失败", zap.Error(err))
		response.FailWithMessage("Falha ao assinar o recebimento do pedido:"+err.Error(), c)
	}
	response.OkWithMessage("O pedido foi assinado com sucesso", c)

}
func (m *MallOrderApi) OrderDetailPage(c *gin.Context) {
	orderNo := c.Param("orderNo")
	token := c.GetHeader("token")
	if err, orderDetail := mallOrderService.GetOrderDetailByOrderNo(token, orderNo); err != nil {
		global.GVA_LOG.Error("查询订单详情接口", zap.Error(err))
		response.FailWithMessage("Interface para verificação dos detalhes do pedido:"+err.Error(), c)
	} else {
		response.OkWithData(orderDetail, c)
	}
}

func (m *MallOrderApi) OrderList(c *gin.Context) {
	token := c.GetHeader("token")
	pageNumber, _ := strconv.Atoi(c.Query("pageNumber"))
	status := c.Query("status")
	if pageNumber <= 0 {
		pageNumber = 1
	}
	if err, list, total := mallOrderService.MallOrderListBySearch(token, pageNumber, status); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("A solicitação falhou"+err.Error(), c)
	} else if len(list) < 1 {
		// 前端项目这里有一个取数逻辑，如果数组为空，数组需要为[] 不能是Null
		response.OkWithDetailed(response.PageResult{
			List:       make([]interface{}, 0),
			TotalCount: total,
			CurrPage:   pageNumber,
			PageSize:   80,
		}, "SUCCESS", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:       list,
			TotalCount: total,
			CurrPage:   pageNumber,
			PageSize:   80,
		}, "SUCCESS", c)
	}

}
