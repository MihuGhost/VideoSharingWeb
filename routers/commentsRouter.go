package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {
	/*
		UserController
	*/
	beego.GlobalControllerRouter["VideoSharingWeb/controllers:UserController"] = append(beego.GlobalControllerRouter["VideoSharingWeb/controllers:UserController"],
		beego.ControllerComments{
			Method:           "LoginDo",
			Router:           `/login/do`,
			AllowHTTPMethods: []string{"*"},
			Params:           nil,
			MethodParams:     param.Make(),
		})

	beego.GlobalControllerRouter["VideoSharingWeb/controllers:UserController"] = append(beego.GlobalControllerRouter["VideoSharingWeb/controllers:UserController"],
		beego.ControllerComments{
			Method:           "SaveRegister",
			Router:           `/register/save`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil,
			MethodParams:     param.Make(),
		})

	beego.GlobalControllerRouter["VideoSharingWeb/controllers:UserController"] = append(beego.GlobalControllerRouter["VideoSharingWeb/controllers:UserController"],
		beego.ControllerComments{
			Method:           "SendMessageDo",
			Router:           `/send/message`,
			AllowHTTPMethods: []string{"*"},
			Params:           nil,
			MethodParams:     param.Make(),
		})

	/*
		VideoController
	*/
	beego.GlobalControllerRouter["VideoSharingWeb/controllers:VideoController"] = append(beego.GlobalControllerRouter["VideoSharingWeb/controllers:VideoController"],
		beego.ControllerComments{
			Method:           "ChannelAdvert",
			Router:           `/channel/advert`,
			AllowHTTPMethods: []string{"*"},
			Params:           nil,
			MethodParams:     param.Make(),
		})

	beego.GlobalControllerRouter["VideoSharingWeb/controllers:VideoController"] = append(beego.GlobalControllerRouter["VideoSharingWeb/controllers:VideoController"],
		beego.ControllerComments{
			Method:           "ChannelHotList",
			Router:           `/channel/hot`,
			AllowHTTPMethods: []string{"*"},
			Params:           nil,
			MethodParams:     param.Make(),
		})

	beego.GlobalControllerRouter["VideoSharingWeb/controllers:VideoController"] = append(beego.GlobalControllerRouter["VideoSharingWeb/controllers:VideoController"],
		beego.ControllerComments{
			Method:           "ChannelRecommendRegionList",
			Router:           `/channel/recommend/region`,
			AllowHTTPMethods: []string{"*"},
			Params:           nil,
			MethodParams:     param.Make(),
		})

	beego.GlobalControllerRouter["VideoSharingWeb/controllers:VideoController"] = append(beego.GlobalControllerRouter["VideoSharingWeb/controllers:VideoController"],
		beego.ControllerComments{
			Method:           "GetChannelRecomendTypeList",
			Router:           `/channel/recommend/type`,
			AllowHTTPMethods: []string{"*"},
			Params:           nil,
			MethodParams:     param.Make(),
		})

	beego.GlobalControllerRouter["VideoSharingWeb/controllers:VideoController"] = append(beego.GlobalControllerRouter["VideoSharingWeb/controllers:VideoController"],
		beego.ControllerComments{
			Method:           "ChannelVideo",
			Router:           `/channel/video`,
			AllowHTTPMethods: []string{"*"},
			Params:           nil,
			MethodParams:     param.Make(),
		})

	beego.GlobalControllerRouter["VideoSharingWeb/controllers:VideoController"] = append(beego.GlobalControllerRouter["VideoSharingWeb/controllers:VideoController"],
		beego.ControllerComments{
			Method:           "UserVideo",
			Router:           `/user/video`,
			AllowHTTPMethods: []string{"*"},
			Params:           nil,
			MethodParams:     param.Make(),
		})

	beego.GlobalControllerRouter["VideoSharingWeb/controllers:VideoController"] = append(beego.GlobalControllerRouter["VideoSharingWeb/controllers:VideoController"],
		beego.ControllerComments{
			Method:           "VideoEpisodesList",
			Router:           `/video/episodes/list`,
			AllowHTTPMethods: []string{"*"},
			Params:           nil,
			MethodParams:     param.Make(),
		})

	beego.GlobalControllerRouter["VideoSharingWeb/controllers:VideoController"] = append(beego.GlobalControllerRouter["VideoSharingWeb/controllers:VideoController"],
		beego.ControllerComments{
			Method:           "VideoInfo",
			Router:           `/video/info`,
			AllowHTTPMethods: []string{"*"},
			Params:           nil,
			MethodParams:     param.Make(),
		})

	beego.GlobalControllerRouter["VideoSharingWeb/controllers:VideoController"] = append(beego.GlobalControllerRouter["VideoSharingWeb/controllers:VideoController"],
		beego.ControllerComments{
			Method:           "VideoSave",
			Router:           `/video/save`,
			AllowHTTPMethods: []string{"*"},
			Params:           nil,
			MethodParams:     param.Make(),
		})

	beego.GlobalControllerRouter["VideoSharingWeb/controllers:VideoController"] = append(beego.GlobalControllerRouter["VideoSharingWeb/controllers:VideoController"],
		beego.ControllerComments{
			Method:           "SaveAll",
			Router:           `/video/save/all`,
			AllowHTTPMethods: []string{"*"},
			Params:           nil,
			MethodParams:     param.Make(),
		})

	beego.GlobalControllerRouter["VideoSharingWeb/controllers:VideoController"] = append(beego.GlobalControllerRouter["VideoSharingWeb/controllers:VideoController"],
		beego.ControllerComments{
			Method:           "Search",
			Router:           `/video/search`,
			AllowHTTPMethods: []string{"*"},
			Params:           nil,
			MethodParams:     param.Make(),
		})
}
