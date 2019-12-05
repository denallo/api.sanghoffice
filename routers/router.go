// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"api.sanghoffice/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/kuties",
			beego.NSInclude(
				&controllers.KutiController{},
			),
		),
		beego.NSNamespace("/residents",
			beego.NSInclude(
				&controllers.ResidentCtrl{},
			),
		),
		beego.NSNamespace("/resiStatus",
			beego.NSInclude(
				&controllers.ResiStatusCtrl{},
			),
		),
		beego.NSNamespace("/items",
			beego.NSInclude(
				&controllers.ItemController{},
			),
		),
		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UsersCtrl{},
			),
		),
	)
	beego.AddNamespace(ns)
}
