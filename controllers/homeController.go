package controllers

import (
	"github.com/astaxie/beego"
	"github.com/yunnet/peony/enums"
	"github.com/yunnet/peony/jwt"
	"github.com/yunnet/peony/models"
	"github.com/yunnet/peony/utils"
	"time"
)

//主要入口
type HomeController struct {
	BaseController
}

func (this *HomeController) URLMapping() {
	this.Mapping("dologin", this.DoLogin)
	this.Mapping("config", this.GetConfig)
}

// DoLogin
// @Description
// @Param	username	query	string	true	"user name"
// @Param	password	query	string	true	"user password"
// @Success 200 {object} models.SysBackendUser
// @Failure 403
// @router /dologin [get]
func (this *HomeController) DoLogin() {
	username := this.Ctx.Input.Query("username")
	password := this.Ctx.Input.Query("password")
	if len(username) == 0 || len(password) == 0 {
		this.jsonResult(enums.JRCodeFailed, "用户名或密码不能为空", "")
	}

	userpwd := utils.String2md5(password)
	user, err := models.GetSysBackendUserByName(username, userpwd)
	if user != nil && err == nil {
		if user.Status == enums.Disabled {
			this.jsonResult(enums.JRCodeFailed, "用户被禁用，请联系管理员", "")
		}

		//保存用户信息到session
		this.setBackendUser2Session(user.Id)

		et := jwt.EasyToken{
			Username: username,
			Expires:  time.Now().Unix() + 3600,
		}
		tokenString, _ := et.GetToken()
		this.jsonResult(enums.JRCodeSucc, "登陆成功。", `{"token": "`+tokenString+`"}`)
	} else {
		this.jsonResult(enums.JRCode401, "登陆失败，账号或密码错误。 ", username)
	}
}

//GetConfig
// @Description
// @Param	key	query	string	true	"key"
// @Success 200 {object} string
// @Failure 403
// @router /config [get]
func (this *HomeController) GetConfig() {
	key := this.GetString("key")
	result := ""
	err := true
	if key == "siteApp" {
		result = beego.AppConfig.String("site.app")
		err = false
	} else if key == "siteName" {
		result = beego.AppConfig.String("site.name")
		err = false
	} else if key == "siteVersion" {
		result = beego.AppConfig.String("site.version")
		err = false
	}

	if err {
		this.jsonResult(enums.JRCodeFailed, "获取参数失败", key)
	} else {
		this.jsonResult(enums.JRCodeSucc, "", result)
	}
}
