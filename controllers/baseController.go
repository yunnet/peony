package controllers

import (
	"github.com/astaxie/beego"
	"github.com/yunnet/walkdog/enums"
	"github.com/yunnet/walkdog/jwt"
	"github.com/yunnet/walkdog/models"
)

const C_SESSION = "WkUser"

type BaseController struct {
	beego.Controller
	controllerName string //当前控制名称
	actionName     string //当前action名称
	curUser        models.WkUser
}

func (c *BaseController) Prepare() {
	//从Session里获取数据 设置用户信息
	c.adapterUserInfo()
}

//适配用户信息
func (c *BaseController) adapterUserInfo() {
	a := c.GetSession(C_SESSION)
	if a != nil {
		c.curUser = a.(models.WkUser)
		c.Data[C_SESSION] = a
	}
}

//SetBackendUser2Session 获取用户信息（包括资源UrlFor）保存至Session
//被 HomeController.DoLogin 调用
func (c *BaseController) setUser2Session(userId int) error {
	m, err := models.GetWkUserById(userId)
	if err != nil {
		return err
	}
	c.SetSession(C_SESSION, *m)
	return nil
}

func (c *BaseController) checkAuthor(ignores ...string) {
	//先判断是否登录
	c.checkLogin()

	//增加对权限的控制
	//如果Action在忽略列表里，则直接通用
	for _, ignore := range ignores {
		if ignore == c.actionName {
			return
		}
	}
}

func (c *BaseController) checkLogin() {
	tokenString := c.Ctx.Request.Header.Get("token")
	//tokenString := c.Ctx.Input.Query("token")
	et := jwt.EasyToken{}
	valido, _ := et.ValidateToken(tokenString)
	if !valido {
		c.jsonResult(enums.JRCode401, "无权访问", "")
	}
	beego.Info("check login user: " + et.Username)
	return
}

func (c *BaseController) jsonResult(code enums.JsonResultCode, msg string, data interface{}) {
	res := &models.JsonResult{code, msg, data}
	c.Data["json"] = res

	c.ServeJSON()
	c.StopRun()
}
