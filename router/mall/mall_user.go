package mall

import (
	"github.com/gin-gonic/gin"
	v1 "main.go/api/v1"
	"main.go/middleware"
)

type MallUserRouter struct {
}

func (m *MallUserRouter) InitMallUserRouter(Router *gin.RouterGroup) {
	mallUserRouter := Router.Group("v1").Use(middleware.UserJWTAuth())
	userRouter := Router.Group("v1")
	var mallUserApi = v1.ApiGroupApp.MallApiGroup.MallUserApi
	{
		mallUserRouter.PUT("/user/info", mallUserApi.UserInfoUpdate)        //修改用户信息
		mallUserRouter.GET("/user/info", mallUserApi.GetUserInfo)           //获取用户信息
		mallUserRouter.POST("/user/info/V2", mallUserApi.GetUserInfoV2)     //获取用户信息
		mallUserRouter.POST("/user/logout", mallUserApi.UserLogout)         //登出
		mallUserRouter.POST("/user/withdrawal", mallUserApi.UserWithdrawal) //提款
		mallUserRouter.POST("/user/bankList", mallUserApi.UserBankList)     //银行卡列表
		mallUserRouter.GET("/bank/:bankId", mallUserApi.GetMallUserBank)    //获取 单个银行账户

	}
	{
		userRouter.POST("/user/register", mallUserApi.UserRegister) //用户注册
		userRouter.POST("/user/login", mallUserApi.UserLogin)       //登陆
		userRouter.POST("/user/login/v2", mallUserApi.UserLoginV2)  //登陆
		userRouter.POST("/user/chatList", mallUserApi.UserChatList) //获取充值客服联系方式
	}

}
