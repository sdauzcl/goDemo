package main

import (
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"goDemo/test"
)

func main() {
	test.Say("dsds")
	web.Run()
}
