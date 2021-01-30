// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"sogo/controllers"
	"sogo/controllers/oauth"
	"sogo/controllers/server"
	"sogo/controllers/wechat"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
	)

	beego.Router("/oauth/github", &oauth.GithubController{})

	beego.Router("/server/wechat", &server.WechatController{}, "*:Get")
	beego.Router("/wechat/qrcode", &wechat.QrCodeController{})
	beego.Router("/s?:key.msg", &wechat.SendController{})

	beego.Router("/", &controllers.WelcomeController{}, "*:Get")
	beego.AddNamespace(ns)
}
