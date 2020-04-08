package controllers

import (
	"github.com/astaxie/beego"
	"github.com/yunnet/peony/enums"
	"github.com/yunnet/peony/jwt"
	"github.com/yunnet/peony/models"
)

const C_SESSION = "SysBackendUser"

type BaseController struct {
	beego.Controller
	curUser models.SysBackendUser
}

func (this *BaseController) Prepare() {
	//从Session里获取数据 设置用户信息
	this.adapterUserInfo()
}

//适配用户信息
func (this *BaseController) adapterUserInfo() {
	a := this.GetSession(C_SESSION)
	if a != nil {
		this.curUser = a.(models.SysBackendUser)
		this.Data[C_SESSION] = a
	}
}

//SetBackendUser2Session 获取用户信息（包括资源UrlFor）保存至Session
//被 HomeController.DoLogin 调用
func (this *BaseController) setBackendUser2Session(userId int) error {
	m, err := models.GetSysBackendUserById(userId)
	if err != nil {
		return err
	}
	this.SetSession(C_SESSION, *m)
	return nil
}

func (this *BaseController) checkAuthor(ignores ...string) {
	//先判断是否登录
	this.checkLogin()

	//增加对权限的控制
}

func (this *BaseController) checkLogin() {
	tokenString := this.Ctx.Request.Header.Get("token")
	//tokenString := this.Ctx.Input.Query("token")
	et := jwt.EasyToken{}
	valido, _ := et.ValidateToken(tokenString)
	if !valido {
		this.jsonResult(enums.JRCode401, "无权访问", "")
	}
	beego.Info("check login user: " + et.Username)
	return
}

func (this *BaseController) jsonResult(code enums.JsonResultCode, msg string, data interface{}) {
	res := &models.JsonResult{code, msg, data}
	this.Data["json"] = res

	this.ServeJSON()
	this.StopRun()
}
