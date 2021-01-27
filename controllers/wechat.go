package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
)

// Operations about Users
type WechatController struct {
	beego.Controller
}

func (o *WechatController) Get() {
	wc := wechat.NewWechat()
	redis := cache.NewRedis(&cache.RedisOpts{Host: "127.0.0.1:6379", Database: 6})
	cfg := &config.Config{
		AppID:     "wxc5398ace5f1edd4d",
		AppSecret: "7d714fbe0304d3c998a3bd8fbb40437f",
		Token:     "83c3893cc9e0b27efd6641198ed63a29",
		//EncodingAESKey: "xxxx",
		Cache: redis,
	}
	officialAccount := wc.GetOfficialAccount(cfg)
	// 传入request和responseWriter
	server := officialAccount.GetServer(o.Ctx.Request, o.Ctx.ResponseWriter)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {

		//回复消息：演示回复用户发送的消息
		text := message.NewText(msg.Content)
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
