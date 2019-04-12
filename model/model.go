/**
  * @author liliangbin  dumpling1520@gmail.com
  * @date 2019/4/12  15:05
  *
  **/

package model

import "net/http"

type Router struct {
	Name       string `json:"name"`
	Method     string `json:"method"`
	Pattern    string `json:"pattern"`
	HandleFunc http.HandlerFunc
}

type Routes []Router

type Rii struct {
	Name   string `json:"name"`
	Data   string `json:"data"`
	Indexx string `json:"indexx"`
}
