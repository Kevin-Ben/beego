package ding

import (
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"io/ioutil"
)

type BatchController struct {
	beego.Controller
}

type DingPostItem struct {
	Token  string `json:"token"`
	Secret string `json:"secret"`
}

type DingPost struct {
	Name  string         `json:"name"`
	Items []DingPostItem `json:"items"`
}

func (c *BatchController) Batch() {

	var object DingPost
	//var body io.Reader

	body := c.Ctx.Request.Body

	bytes, _ := ioutil.ReadAll(body)

	json.Unmarshal(bytes, &object)

	for _, v := range object.Items {
		fmt.Printf("%v -- %v\n", v.Token, v.Secret)
	}

}
