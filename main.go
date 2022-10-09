package main

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	web.Run()
}
func Test() {
	fmt.Println("我是来自github的远程代码")
}
