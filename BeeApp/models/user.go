package models

import (
    "github.com/beego/beego/v2/client/orm"
    _ "github.com/go-sql-driver/mysql" // import your database driver
)

type User struct {
    Id       int
    Username string `orm:"unique"`
    Password string
    IsAdmin  bool
}

func init() {
    orm.RegisterModel(new(User))
}
