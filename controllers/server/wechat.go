package server

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"sogo/models"
	utils "sogo/utils/wechat"
	"strconv"
)

// Operations about Users
type WechatController struct {
	beego.Controller
}

func (o *WechatController) Get() {

	officialAccount := utils.GetAccount()
	// 传入request和responseWriter
	server := officialAccount.GetServer(o.Ctx.Request, o.Ctx.ResponseWriter)

	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {

		if msg.Event == "SCAN" {
			redisCache := utils.GetRedisCache()
			HashMap := redisCache.Get(msg.Ticket)
			if HashMap == nil {
				text := message.NewText("二维码已过期，请重新生成")
				return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
			}

			//更新数据库
			userId := HashMap.(map[string]interface{})["userId"].(string)
			objecId, _ := strconv.ParseInt(userId, 10, 64)
			sendTo := string(msg.FromUserName)

			_, err := models.ReadOrCreateWechat(sendTo, objecId)
			if err == nil {
				text := message.NewText(err.Error())
				return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
			}

			//更新订阅
			subPri := Get16MD5Encode(userId)
			_, err = models.ReadOrCreateSubscribe(subPri, sendTo)
			if err == nil {
				text := message.NewText(err.Error())
				return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
			}

			text := message.NewText("操作成功")
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
		}

		//回复消息：演示回复用户发送的消息
		text := message.NewText("感谢您的关注与支持")
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}

	//发送回复的消息
	server.Send()
}

func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

//返回一个16位md5加密后的字符串
func Get16MD5Encode(data string) string {
	return GetMD5Encode(data)[8:24]
}
