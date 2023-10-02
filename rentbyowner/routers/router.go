package routers

import (
	"rentbyowner/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/refine", &controllers.RefineSearch{}, "get:GetRefineSearch")
	beego.Router("/property", &controllers.PropertyDetail{}, "get:GetPropertyDetail")
	beego.Router("/property/:countrycode/:idname", &controllers.PropertyInfoDetail{}, "get:GetPropertyDetails")
}
