package routers

import (
    "BeeApp/controllers"
    "github.com/beego/beego/v2/server/web"
    "github.com/beego/beego/v2/server/web/context"
)

func checkLogin(ctx *context.Context) {
    user := ctx.Input.Session("user")
    if user == nil {
        ctx.Redirect(302, "/login")
    }
}

func init() {
    web.InsertFilter("/*", web.BeforeRouter, checkLogin)
    web.Router("/", &controllers.IndexController{})
    web.Router("/login", &controllers.LoginController{})
}
