// @APIVersion 1.0.0
// @Title Walkdog Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact yunnet@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/yunnet/walkdog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/wk_user_status",
			beego.NSInclude(
				&controllers.WkUserStatusController{},
			),
		),

		beego.NSNamespace("/wk_user",
			beego.NSInclude(
				&controllers.WkUserController{},
			),
		),

		beego.NSNamespace("/wk_stores",
			beego.NSInclude(
				&controllers.WkStoresController{},
			),
		),

		beego.NSNamespace("/home",
			beego.NSInclude(
				&controllers.HomeController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
