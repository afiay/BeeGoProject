package controllers

import (
    "BeeApp/models"
    "github.com/beego/beego/v2/client/orm"
    "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
    web.Controller
}

func (c *LoginController) Get() {
    c.TplName = "login.tpl"
}

func (c *LoginController) Post() {
    username := c.GetString("username")
    password := c.GetString("password")

    o := orm.NewOrm()
    user := models.User{Username: username}
    if err := o.Read(&user, "Username"); err == nil {
        if user.Password == password { // In real applications, use hashed passwords
            c.SetSession("user", user)
            c.Redirect("/", 302)
            return
        }
    }
    c.Data["Error"] = "Invalid username or password"
    c.TplName = "login.tpl"
}
