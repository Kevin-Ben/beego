package ding

import (
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	"io/ioutil"
)

// Operations about Single
type SingleController struct {
	beego.Controller
}

type DingSingle struct {
	Token  string `json:"token"`
	Secret string `json:"secret"`
}

// @Title Create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (c *SingleController) Post() {

	var object DingSingle

	body := c.Ctx.Request.Body

	bytes, _ := ioutil.ReadAll(body)

	json.Unmarshal(bytes, &object)

	c.Data["json"] = object

	c.ServeJSON()

}
