package wechat

import (
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/silenceper/wechat/v2/util"
	log "github.com/sirupsen/logrus"
	"sogo/utils/wechat"
	"time"
)

// Operations about Users
type QrCodeController struct {
	beego.Controller
}

type JsonResponse struct {
	success string `json:"success"`
}

func (o *QrCodeController) Get() {

	var err error

	log.SetLevel(log.WarnLevel)

	account := wechat.GetAccount()

	token, _ := account.GetAccessToken()

	json := `{"expire_seconds": 604800, "action_name": "QR_SCENE", "action_info": {"scene": {"scene_str": 123}}}`

	response, err := util.HTTPPost("https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token="+token, json)
	if err != nil {
		return
	}

	HashMap, _ := JsonToMap(string(response))

	cache := wechat.GetRedisCache()

	HashMap["userId"] = "10203040"

	timeOut := time.Duration(3600) * time.Second

	err = cache.Set(HashMap["ticket"].(string), HashMap, timeOut)

	if err != nil {
		fmt.Println(err.Error())
	}

	o.Ctx.WriteString(string(response))

}

func JsonToMap(jsonStr string) (m map[string]interface{}, err error) {
	a := make(map[string]interface{})
	unmarsha1Err := json.Unmarshal([]byte(jsonStr), &a)
	if unmarsha1Err != nil {
		return nil, unmarsha1Err
	}
	return a, nil
}
