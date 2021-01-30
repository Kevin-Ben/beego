package wechat

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/golang-module/carbon"
	"sogo/models"
	"sogo/utils/idgen"
	msg "sogo/utils/wechat"
	"strconv"
)

// SendController operations for Send
type SendController struct {
	beego.Controller
}

func (c *SendController) doRun() {

	response := make(map[string]interface{})

	title := c.GetString("title")
	content := c.GetString("msg")
	key := c.GetString(":key")

	sub, err := models.GetSubscribeBySubPri(key)

	if err == orm.ErrNoRows {
		response["success"] = true
		response["errorCode"] = 404
		response["errorMessage"] = "Not Found"
		response["model"] = nil
		c.Data["json"] = response
		c.ServeJSON()
		return
	}

	response["success"] = true
	response["errorCode"] = 0
	response["errorMessage"] = "提交成功"
	object := make(map[string]interface{})

	//保存数据库
	msgId, err := saveContent(key, title, content, sub.SendTo)

	if err != nil {
		response["errorMessage"] = err.Error()
		response["model"] = nil
		c.Data["json"] = response
		c.ServeJSON()
		return
	}

	//推送微信消息
	_, warn := msg.Send(strconv.FormatInt(msgId, 10), title, sub.SendTo)

	if warn != nil {
		response["errorMessage"] = warn.Error()
		response["model"] = nil
		c.Data["json"] = response
		c.ServeJSON()
		return
	}

	object["id"] = msgId
	object["title"] = title
	response["model"] = object
	c.Data["json"] = response
	c.ServeJSON()
}

func saveContent(subPri, title, content, sendTo string) (id int64, err error) {
	m := &models.Content{
		Id:      int64(GetId()),
		SubPri:  subPri,
		Title:   title,
		SendTo:  sendTo,
		Content: content,
	}

	m.CreatedAt = carbon.Now().ToGoTime()
	m.UpdatedAt = m.CreatedAt
	_, err = models.AddContent(m)
	return m.Id, err
}

func GetId() uint64 {
	gen := idgen.NewWorker(1, 1)
	id, _ := gen.NextID()
	return id
}

// Post ...
// @Title Create
// @Description create Send
// @Param	body		body 	models.Send	true		"body for Send content"
// @Success 201 {object} models.Send
// @Failure 403 body is empty
// @router / [post]
func (c *SendController) Post() {
	c.doRun()
}

// GetOne ...
// @Title Get
// @Description get Send by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Send
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SendController) Get() {
	c.doRun()
}
