package controllers

import (
	"github.com/yunnet/walkdog/enums"
	"github.com/yunnet/walkdog/jwt"
	"github.com/yunnet/walkdog/models"
	"github.com/yunnet/walkdog/utils"
	"time"
)

//主要入口
type HomeController struct {
	BaseController
}

func (this *HomeController) URLMapping() {
	this.Mapping("dologin", this.DoLogin)
}

func (this *HomeController) DoLogin() {
	username := this.Ctx.Input.Query("username")
	password := this.Ctx.Input.Query("password")
	if len(username) == 0 || len(password) == 0 {
		this.jsonResult(enums.JRCodeFailed, "用户名或密码不能为空", "")
	}

	userpwd := utils.String2md5(password)
	user, err := models.GetWkUserByName(username, userpwd)
	if user != nil && err == nil {
		if user.Status == enums.Disabled {
			this.jsonResult(enums.JRCodeFailed, "用户被禁用，请联系管理员", "")
		}

		//保存用户信息到session
		this.setUser2Session(user.Id)

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
