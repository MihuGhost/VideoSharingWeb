package controllers

import (
	"VideoSharingWeb/models"
	"github.com/astaxie/beego"
	"regexp"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) SaveRegister() {
	var (
		mobile   string
		password string
		err      error
	)
	mobile = this.GetString("mobile")
	password = this.GetString("password")
	if mobile == "" {
		this.Data["json"] = ReturnError(4001, "手机号不能为空")
		this.ServeJSON()
	}
	isorno, _ := regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, mobile)
	if !isorno {
		this.Data["json"] = ReturnError(4002, "请输入正确的手机号")
		this.ServeJSON()
	}
	if password == "" {
		this.Data["json"] = ReturnError(4003, "密码不能为空")
		this.ServeJSON()
	}

	status := models.IsUserMobile(mobile)
	if status {
		this.Data["json"] = ReturnError(4005, "此手机号已经注册")
		this.ServeJSON()
	} else {
		err = models.UserSave(mobile, MD5V(password))
		if err == nil {
			this.Data["json"] = ReturnSuccess(0, "注册成功", nil, 0)
			this.ServeJSON()
		} else {
			this.Data["json"] = ReturnError(5000, err)
			this.ServeJSON()
		}
	}
}
