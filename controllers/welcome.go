package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

// WelcomeController operations for Welcome
type WelcomeController struct {
	beego.Controller
}

type ErrorJson struct {
	success      bool
	errorCode    int64
	errorMessage string
	model        interface{}
}

func (c *WelcomeController) Get() {

	response := make(map[string]interface{})
	response["success"]	= true
	response["errorCode"]	= 0
	response["errorMessage"]	= nil
	response["model"]	= nil

	c.Data["json"] = response

	c.ServeJSON()
}
