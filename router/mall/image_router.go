package mall

import (
	"github.com/gin-gonic/gin"
)

type ImageHttpRouter struct {
}

func (m *ImageHttpRouter) InitImageRouter(Router *gin.RouterGroup) {
	//mallCarouselRouter := Router.Group("v1")
	//{
	//	mallCarouselRouter.Static("/static/", "./static-files/img") // 读取图片
	//}
}
