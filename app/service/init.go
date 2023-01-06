package service

import (
    "errors"
    "fmt"
    "github.com/go-xorm/xorm"
    "log"
    "app/env"
    _ "github.com/go-sql-driver/mysql"
)

var DbEngine *xorm.Engine

func init()  {
    driverName := env.Get("DB_CONNECTION")
    db_host := env.Get("DB_HOST")
    db_port := env.Get("DB_PORT")
    db_database := env.Get("DB_DATABASE")
    db_user_name := env.Get("DB_USERNAME")
    db_password := env.Get("DB_PASSWORD")

    DsName := db_user_name + ":" + db_password + "@tcp(" + db_host + ":" + db_port + ")/" + db_database + "?charset=utf8mb4"
    err := errors.New("")
    DbEngine, err = xorm.NewEngine(driverName, DsName)
    if err != nil && err.Error() != ""{
        log.Fatal(err.Error())
    }
    DbEngine.ShowSQL(true)
    DbEngine.SetMaxOpenConns(2)
    
    fmt.Println("init data base ok")
}