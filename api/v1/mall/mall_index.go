package mall

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common/enum"
	"main.go/model/common/response"
)

type MallIndexApi struct {
}

// MallIndexInfo 加载首页信息
func (m *MallIndexApi) MallIndexInfo(c *gin.Context) {
	err, _, mallCarouseInfo := mallCarouselService.GetCarouselsForIndex(16)
	if err != nil {
		global.GVA_LOG.Error("轮播图获取失败"+err.Error(), zap.Error(err))
		response.FailWithMessage("Не удалось получить вращающееся изображение", c)
	}
	err, hotGoodses := mallIndexConfigService.GetConfigGoodsForIndex(enum.IndexGoodsHot.Code(), 36)
	if err != nil {
		global.GVA_LOG.Error("热销商品获取失败"+err.Error(), zap.Error(err))
		response.FailWithMessage("Неудача при приобретении горячего продукта", c)
	}
	err, newGoodses := mallIndexConfigService.GetConfigGoodsForIndex(enum.IndexGoodsNew.Code(), 36)
	if err != nil {
		global.GVA_LOG.Error("新品获取失败"+err.Error(), zap.Error(err))
		response.FailWithMessage("Неудачное приобретение новой продукции", c)
	}
	err, recommendGoodses := mallIndexConfigService.GetConfigGoodsForIndex(enum.IndexGoodsRecommond.Code(), 30)
	if err != nil {
		global.GVA_LOG.Error("推荐商品获取失败"+err.Error(), zap.Error(err))
		response.FailWithMessage("Не удалось получить рекомендованные продукты", c)
	}
	indexResult := make(map[string]interface{})
	indexResult["carousels"] = mallCarouseInfo
	indexResult["hotGoodses"] = hotGoodses
	indexResult["newGoodses"] = newGoodses
	indexResult["recommendGoodses"] = recommendGoodses
	response.OkWithData(indexResult, c)

}
