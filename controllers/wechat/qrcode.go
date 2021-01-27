package wechat

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/util"
)

// Operations about Users
type QrCodeController struct {
	beego.Controller
}

type JsonResponse struct {
	success string `json:"success"`
}

func (o *QrCodeController) Get() {

	wc := wechat.NewWechat()
	redis := cache.NewRedis(&cache.RedisOpts{Host: "127.0.0.1:6379", Database: 6})
	cfg := &config.Config{
		AppID:     "wxc5398ace5f1edd4d",
		AppSecret: "7d714fbe0304d3c998a3bd8fbb40437f",
		Token:     "83c3893cc9e0b27efd6641198ed63a29",
		//EncodingAESKey: "xxxx",
		Cache: redis,
	}

	account := wc.GetOfficialAccount(cfg)

	token, _ := account.GetAccessToken()

	json :=`{"expire_seconds": 604800, "action_name": "QR_SCENE", "action_info": {"scene": {"scene_str": 123}}}`

	response, err :=util.HTTPPost("https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=" + token,json)
	if err != nil {
		return
	}
	o.Ctx.WriteString(string(response))

}
