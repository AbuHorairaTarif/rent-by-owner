package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

type RefineSearch struct {
	beego.Controller
	Location     string
	Checkin_date string
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.vip"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *RefineSearch) GetRefineSearch() {
	c.Data["refine"] = &RefineSearch{

		Location:     "Dhaka,Bangladesh",
		Checkin_date: "02-10-2023",
	}

	c.TplName = "refine.tpl"
}
