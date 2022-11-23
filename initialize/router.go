package initialize

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"main.go/global"
	"main.go/middleware"
	"main.go/router"
	"net/http"
)

func Routers() *gin.Engine {
	var Router = gin.Default()

	//这里读取了静态文件
	Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址

	//Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.GVA_LOG.Info("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors()) // 如需跨域可以打开
	global.GVA_LOG.Info("use middleware cors")
	// 方便统一添加路由组前缀 多服务器上线使用
	//商城后管路由
	manageRouter := router.RouterGroupApp.Manage
	ManageGroup := Router.Group("manage-api")

	//
	//https://wizardforcel.gitbooks.io/go42/content/content/42_42_gin.html
	//读取本地图片
	//http://localhost:8888/getImage?imageName=./001.png
	Router.GET("/getImage", func(c *gin.Context) {
		getImage(c)
	}) // 读取图片

	PublicGroup := Router.Group("")

	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	{
		//商城后管路由初始化
		manageRouter.InitManageAdminUserRouter(ManageGroup)
		manageRouter.InitManageGoodsCategoryRouter(ManageGroup)
		manageRouter.InitManageGoodsInfoRouter(ManageGroup)
		manageRouter.InitManageCarouselRouter(ManageGroup)
		manageRouter.InitManageIndexConfigRouter(ManageGroup)
		manageRouter.InitManageOrderRouter(ManageGroup)
	}
	//商城前端路由
	mallRouter := router.RouterGroupApp.Mall
	MallGroup := Router.Group("api")
	{
		// 商城前端路由
		mallRouter.InitMallCarouselIndexRouter(MallGroup)
		mallRouter.InitMallGoodsInfoIndexRouter(MallGroup)
		mallRouter.InitMallGoodsCategoryIndexRouter(MallGroup)
		mallRouter.InitMallUserRouter(MallGroup)
		mallRouter.InitMallUserAddressRouter(MallGroup)
		mallRouter.InitMallShopCartRouter(MallGroup)
		mallRouter.InitMallOrderRouter(MallGroup)
	}
	global.GVA_LOG.Info("router register success")
	return Router
}

func getImage(c *gin.Context) {
	imageName := c.Query("imageName")
	file, _ := ioutil.ReadFile(imageName)
	c.Writer.WriteString(string(file))
}
