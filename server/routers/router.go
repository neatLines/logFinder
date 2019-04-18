// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/neatLines/logFinder/server/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// ns := beego.NewNamespace("/v1",
	// 	beego.NSNamespace("/hosts",
	// 		beego.NSInclude(
	// 			&controllers.HostsController{},
	// 		),
	// 	),
	// 	beego.NSNamespace("/user",
	// 		beego.NSInclude(
	// 			&controllers.UserController{},
	// 		),
	// 	),
	// )
	// beego.AddNamespace(ns)
	beego.Router("/v1/ws", &controllers.WsController{})
	beego.Router("/v1/hosts/*", &controllers.HostsController{})
	beego.Router("/v1/user/info", &controllers.UserController{}, "get:Info")
	beego.Router("/v1/login", &controllers.UserController{}, "post:Login")
	beego.Router("/v1/logout", &controllers.UserController{}, "post:Logout")
}
