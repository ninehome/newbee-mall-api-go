package manage

import (
	"github.com/gin-gonic/gin"
	v1 "main.go/api/v1"
	"main.go/middleware"
)

type ManageOrderRouter struct {
}

func (r *ManageOrderRouter) InitManageOrderRouter(Router *gin.RouterGroup) {
	mallOrderRouter := Router.Group("v1").Use(middleware.AdminJWTAuth())
	//mallAdminUserWithoutRouter := Router.Group("v1") //不鉴权
	var mallOrderApi = v1.ApiGroupApp.ManageApiGroup.ManageOrderApi
	{
		mallOrderRouter.PUT("orders/checkDone", mallOrderApi.CheckDoneOrder)   // 发货
		mallOrderRouter.PUT("orders/checkOut", mallOrderApi.CheckOutOrder)     // 出库
		mallOrderRouter.PUT("orders/close", mallOrderApi.CloseOrder)           // 出库
		mallOrderRouter.GET("orders/:orderId", mallOrderApi.FindMallOrder)     // 根据ID获取MallOrder
		mallOrderRouter.GET("orders", mallOrderApi.GetMallOrderList)           // 获取MallOrder列表
		mallOrderRouter.POST("orders/v2", mallOrderApi.GetMallBuyBackListV2)   // 获取MallOrder列表
		mallOrderRouter.GET("orders/v2", mallOrderApi.GetMallBuyBackListV2)    // 获取MallOrder列表
		mallOrderRouter.GET("orders/buyback", mallOrderApi.GetMallBuyBackList) // 获取 回购的列表
		mallOrderRouter.POST("orders/back", mallOrderApi.ChangeOrderStatus)    //订单回购接口
		mallOrderRouter.POST("orders/name", mallOrderApi.GetMallUserAllOrder)  // 根据传入的用户名 获得此用户所有的订单列表
		mallOrderRouter.GET("orders/name", mallOrderApi.GetMallUserAllOrder)   // 根据传入的用户名 获得此用户所有的订单列表
	}

}
