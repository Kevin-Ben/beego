package wechat

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/golang-module/carbon"
	"github.com/silenceper/wechat/v2/officialaccount/message"
)

func Send(title string, sendTo string) (msgID int64, err error) {

	account := GetAccount()

	template := message.NewTemplate(account.GetContext())

	var data = map[string]*message.TemplateDataItem{}

	data["title"] = &message.TemplateDataItem{
		Value: title + "\n",
	}

	data["content"] = &message.TemplateDataItem{
		Value: "见后文，消息来自哔哩推\n",
	}

	datetime := carbon.Now().ToDateTimeString()

	data["time"] = &message.TemplateDataItem{
		Value: datetime + "\n",
	}

	data["remark"] = &message.TemplateDataItem{
		Value: "点下方[详情]查看推送内容，本卡片内容为您自行编辑，并调用接口推送，并非营销内容，退订请访问blit.run或直接取消订阅。",
	}

	TemplateID, _ := beego.AppConfig.String("TemplateID")
	WebSite, _ := beego.AppConfig.String("WebSite")

	msg := message.TemplateMessage{
		ToUser:     sendTo,
		TemplateID: TemplateID,
		URL:        WebSite + "message/xxx.html",
		Data:       data,
	}

	return template.Send(&msg)
}
