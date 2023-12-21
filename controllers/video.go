package controllers

import (
	"VideoSharingWeb/models"
	"github.com/astaxie/beego"
)

type VideoController struct {
	beego.Controller
}

// @router /channel/advert [*]
func (this *VideoController) ChannelAdvert() {
	//todo GetInt
	channelId, _ := this.GetInt("channelId")

	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
	}
	num, videos, err := models.GetChannelAdvert(channelId)

	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "请求数据失败，请稍后重试")
		this.ServeJSON()
	}
}

// @router /channel/hot [*]
func (this *VideoController) ChannelHotList() {
	channelId, _ := this.GetInt("channelId")

	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
	}
	num, videos, err := models.GetChannelHotList(channelId)

	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
		this.ServeJSON()
	}
}

// @router /channel/recommend/type [*]
func (this *VideoController) GetChannelRecomendTypeList() {
	channelId, _ := this.GetInt("channelId")
	typeId, _ := this.GetInt("typeId")

	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
	}
	if typeId == 0 {
		this.Data["json"] = ReturnError(4002, "必须指定频道类型")
		this.ServeJSON()
	}

	num, videos, err := models.GetChannelRecommendTypeList(channelId, typeId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "没有相关内容")
		this.ServeJSON()
	}

}
