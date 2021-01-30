package wechat

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"sogo/models"
	"sogo/utils/idgen"
	msg "sogo/utils/wechat"
)

// SendController operations for Send
type SendController struct {
	beego.Controller
}

func (c *SendController) doRun() {

	response := make(map[string]interface{})

	title := c.GetString("title")
	//content := c.GetString("msg")

	sub, err := models.GetSubscribeBySubPri(c.GetString(":key"))

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

	_, warn := msg.Send(title, sub.SendTo)

	if warn != nil {
		response["errorMessage"] = warn.Error()
		response["model"] = nil
		c.Data["json"] = response
		c.ServeJSON()
		return
	}

	gen := idgen.NewWorker(1, 1)
	object["id"], _ = gen.NextID()
	object["title"] = title
	response["model"] = object
	c.Data["json"] = response
	c.ServeJSON()
}

func saveContent(title string ,content string, sendTo string)  {

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
