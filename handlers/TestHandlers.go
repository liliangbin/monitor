/**
  * @author liliangbin  dumpling1520@gmail.com
  * @date 2019/4/12  15:10
  *
  **/

package handlers

import (
	"net/http"
	"fmt"
)

func TestInfo(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "hello world")
}
