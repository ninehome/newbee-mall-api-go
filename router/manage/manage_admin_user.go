package manage

import (
	"github.com/gin-gonic/gin"
	v1 "main.go/api/v1"
	"main.go/middleware"
)

type ManageAdminUserRouter struct {
}

func (r *ManageAdminUserRouter) InitManageAdminUserRouter(Router *gin.RouterGroup) {
	mallAdminUserRouter := Router.Group("v1").Use(middleware.AdminJWTAuth())
	mallAdminUserWithoutRouter := Router.Group("v1")
	//http://localhost:8888/manage-api/v1/adminUser/login
	var mallAdminUserApi = v1.ApiGroupApp.ManageApiGroup.ManageAdminUserApi
	{
		mallAdminUserRouter.POST("createMallAdminUser", mallAdminUserApi.CreateAdminUser)              // 新建MallAdminUser
		mallAdminUserRouter.PUT("adminUser/name", mallAdminUserApi.UpdateAdminUserName)                // 更新MallAdminUser
		mallAdminUserRouter.POST("adminUser/moneyAndLevel", mallAdminUserApi.UpdateAdminMoneyAndLevel) // 更新MallAdminUser 的等级和余额
		mallAdminUserRouter.POST("adminUser/rechargeMoney", mallAdminUserApi.UpdateAdminMoney)         // 充值余额
		mallAdminUserRouter.POST("update/password", mallAdminUserApi.UpdateAdminUserPsw)               // 更新MallAdminUser密码
		mallAdminUserRouter.PUT("adminUser/password", mallAdminUserApi.UpdateAdminUserPassword)
		mallAdminUserRouter.GET("users", mallAdminUserApi.UserList) //获取 注册用户列表
		mallAdminUserRouter.PUT("users/:lockStatus", mallAdminUserApi.LockUser)
		mallAdminUserRouter.GET("adminUser/profile", mallAdminUserApi.AdminUserProfile) // 根据ID获取 admin详情
		mallAdminUserRouter.POST("user/profile", mallAdminUserApi.UserProfile)          // 根据ID获取 malluser
		mallAdminUserRouter.POST("user/user/info", mallAdminUserApi.GetUserinfo)        // 根据用户名获取 malluser
		mallAdminUserRouter.DELETE("logout", mallAdminUserApi.AdminLogout)
		mallAdminUserRouter.POST("upload/file", mallAdminUserApi.UploadFile)                                //上传图片
		mallAdminUserRouter.POST("upload/filev2", mallAdminUserApi.Upload)                                  //上传图片 七牛
		mallAdminUserRouter.POST("/users/withdrawals", mallAdminUserApi.WithdrawalHistory)                  //提款列表
		mallAdminUserRouter.POST("/users/recharges", mallAdminUserApi.RechargeHistory)                      //充值列表
		mallAdminUserRouter.POST("/user/recharge/info", mallAdminUserApi.GetURecharge)                      // 根据用户名 获取充值列表
		mallAdminUserRouter.POST("/users/withdrawals/withName", mallAdminUserApi.WithdrawalHistoryWithName) //根据用户名查询提款列表
		mallAdminUserRouter.POST("/update/withdrawal", mallAdminUserApi.UpdateWithdrawal)                   // 更新提款状态
		mallAdminUserRouter.POST("/userBank/update", mallAdminUserApi.UpdateUserBank)                       //修改银行账户
		//mallAdminUserWithoutRouter.POST("adminUser/login", mallAdminUserApi.AdminLogin) //管理员登陆
	}
	{
		mallAdminUserWithoutRouter.POST("adminUser/login", mallAdminUserApi.AdminLogin)   //管理员登陆
		mallAdminUserWithoutRouter.POST("adminUser/create", mallAdminUserApi.AdminCreate) //管理员创建
		mallAdminUserWithoutRouter.GET("/user/chatList", mallAdminUserApi.UserChatList)   //获取充值客服联系方式
		mallAdminUserWithoutRouter.POST("/chat/profile", mallAdminUserApi.ChatProfile)    //根据chatId获得
		mallAdminUserWithoutRouter.POST("/chat/update", mallAdminUserApi.UpdateMallChat)  //根据chatId修改

		mallAdminUserWithoutRouter.POST("/userBank/list", mallAdminUserApi.GetBankList) //获得用户绑定的银行账户

		mallAdminUserWithoutRouter.POST("/user/msg/creat", mallAdminUserApi.CreateUserMsg) //新增或者修改用户私信

		mallAdminUserWithoutRouter.POST("/user/msg/hide", mallAdminUserApi.HideUserMsg) //隐藏用户私信

		mallAdminUserWithoutRouter.POST("/get/user/msg", mallAdminUserApi.GetUserMsg) //获得用户私信
	}
}
