package main

import (
    "BeeApp/models"
    _ "BeeApp/routers"
    "fmt"
    "github.com/beego/beego/v2/client/orm"
    "github.com/beego/beego/v2/server/web"
    _ "github.com/mattn/go-sqlite3"
    "log"
)

func init() {
    // Register the SQLite3 driver and database
    orm.RegisterDriver("sqlite3", orm.DRSqlite)

    dbURL, err := web.AppConfig.String("db_url")
    if err != nil {
        log.Fatalf("Failed to get database URL: %v", err)
    }

    orm.RegisterDataBase("default", "sqlite3", dbURL)

    // Run automatic migrations
    orm.RunSyncdb("default", false, true)
}

func createAdminUser() {
    o := orm.NewOrm()
    user := models.User{
        Username: "admin",
        Password: "admin123", // Make sure to hash passwords in real applications
        IsAdmin:  true,
    }
    if _, err := o.Insert(&user); err != nil {
        fmt.Println("Admin user already exists.")
    } else {
        fmt.Println("Admin user created.")
    }
}

func main() {
    createAdminUser()
    web.Run()
}
