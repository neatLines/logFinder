package controllers

import (
	"encoding/json"

	"github.com/neatLines/logFinder/server/models"

	"github.com/astaxie/beego"
)

// Operations about object
type HostsController struct {
	beego.Controller
}

// @Title Create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (h *HostsController) Post() {
	var hs models.Host
	json.Unmarshal(h.Ctx.Input.RequestBody, &hs)
	models.AddOne(hs)
	h.Data["json"] = "register success"
	h.ServeJSON()
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (h *HostsController) Get() {
	hosts := models.GetAll()
	h.Data["json"] = hosts
	h.ServeJSON()
}
