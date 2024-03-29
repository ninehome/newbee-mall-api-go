package manage

import (
	"github.com/gin-gonic/gin"
	v1 "main.go/api/v1"
)

type ManageGoodsInfoRouter struct {
}

func (m *ManageGoodsInfoRouter) InitManageGoodsInfoRouter(Router *gin.RouterGroup) {
	mallGoodsInfoRouter := Router.Group("v1")
	var mallGoodsInfoApi = v1.ApiGroupApp.ManageApiGroup.ManageGoodsInfoApi
	{
		mallGoodsInfoRouter.POST("goods", mallGoodsInfoApi.CreateGoodsInfo)                    // 新建MallGoodsInfo
		mallGoodsInfoRouter.DELETE("deleteMallGoodsInfo", mallGoodsInfoApi.DeleteGoodsInfo)    // 删除MallGoodsInfo
		mallGoodsInfoRouter.PUT("goods/status/:status", mallGoodsInfoApi.ChangeGoodsInfoByIds) // 上下架
		//mallGoodsInfoRouter.PUT("goods", mallGoodsInfoApi.UpdateGoodsInfo)                     // 更新MallGoodsInfo
		mallGoodsInfoRouter.POST("goods/update", mallGoodsInfoApi.UpdateGoodsInfo)       // 更新MallGoodsInfo
		mallGoodsInfoRouter.POST("goods/countdown", mallGoodsInfoApi.CountdownGoodsInfo) // 商品倒计时
		mallGoodsInfoRouter.POST("goods/cancel", mallGoodsInfoApi.CountdownCancel)       // 商品倒计时删除
		mallGoodsInfoRouter.GET("goods/:id", mallGoodsInfoApi.FindGoodsInfo)             // 根据ID获取MallGoodsInfo
		mallGoodsInfoRouter.GET("goods/list", mallGoodsInfoApi.GetGoodsInfoList)         // 获取MallGoodsInfo列表
		mallGoodsInfoRouter.GET("goods/list/oder", mallGoodsInfoApi.GetGoodsInfoListV2)  // 获取MallGoodsInfo列表 按价格排序
	}
}
