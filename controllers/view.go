package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/golang-module/carbon"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"sogo/models"
	"strconv"
)

// ViewController operations for View
type ViewController struct {
	beego.Controller
}

// Post ...
// @Title Create
// @Description create View
// @Param	body		body 	models.View	true		"body for View content"
// @Success 201 {object} models.View
// @Failure 403 body is empty
// @router / [post]
func (c *ViewController) Get() {
	key := c.GetString(":key")
	id, _ := strconv.ParseInt(key, 10, 64)

	content, err := models.GetContentById(id)

	if err != nil {
		c.Ctx.WriteString("未找到")
		return
	}

	input := []byte(content.Content)
	unsafe := blackfriday.MarkdownCommon(input)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)

	carbon := carbon.CreateFromGoTime(content.CreatedAt)
	now := carbon.Now()

	c.Data["data"] = string(html)
	c.Data["time"] = carbon.ToDateTimeString()
	c.Data["days"] = now.DiffInDays(carbon)

	c.TplName = "content.html"
	c.Render()
}
