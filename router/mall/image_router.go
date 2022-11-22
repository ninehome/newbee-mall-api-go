package mall

import (
	"github.com/gin-gonic/gin"
)

type ImageHttpRouter struct {
}

func (m *ImageHttpRouter) InitImageRouter(Router *gin.RouterGroup) {
	mallCarouselRouter := Router.Group("v1")
	{
		mallCarouselRouter.Static("/img", "./static-files/newbee-mall.png") // 读取图片
	}
}
