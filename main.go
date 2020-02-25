package main

import (
	"fmt"
	_ "myphotos/routers"
	"time"

	models "myphotos/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/session"
	_ "github.com/go-sql-driver/mysql"
)

var (
	mysqlCon = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s",
		beego.AppConfig.String("mysqluser"),
		beego.AppConfig.String("mysqlpass"),
		beego.AppConfig.String("mysqlhost"),
		beego.AppConfig.String("mysqlport"),
		beego.AppConfig.String("mysqldb"),
		beego.AppConfig.String("mysqlcharset"),
	)
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", mysqlCon)
	orm.RegisterModel(new(models.User))
	// orm.RegisterModel(new(models.UserAuth))
	orm.DefaultTimeLoc = time.UTC

	logs.SetLogger(logs.AdapterFile, `{"filename":"myphotos.log"}`)
}

func main() {
	sessionconf := &session.ManagerConfig{
		CookieName: "begoosessionID",
		Gclifetime: 3600,
	}
	beego.GlobalSessions, _ = session.NewManager("memory", sessionconf)
	go beego.GlobalSessions.GC()

	beego.Run()
}
