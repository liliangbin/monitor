package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"monitor/config"
)

func init() {
	fmt.Println("欢迎使用后台管理系统，author ===> liliangbin")
	err:=config.ParseConf("config.json")
	checkError(err)
}

func main() {

	fmt.Println("欢迎使用")

	router := NewRouter()
	n := negroni.Classic()
	n.Use(negroni.NewLogger())
	n.UseHandler(router)
	n.Run(":8083")
}
