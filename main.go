package main

import (
	"bytes"
	"fmt"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func exec_shell(s string) (string, error) {

	fmt.Println(s)
	cmd := exec.Command("bash", "-C", s)

	var out bytes.Buffer
	cmd.Stdout = &out

	fmt.Println(cmd.Args)
	fmt.Println(cmd.Path)
	// 字符串的编码方式
	err := cmd.Run()
	checkError(err)
	str := ConvertToString(out.String(), "gbk", "utf-8")

	return str, err
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	/*	info, err := exec_shell("ls")

		checkError(err)
		fmt.Println(info)*/

	cmd := exec.Command("dir")
	stdout, err := cmd.StdoutPipe() //指向cmd命令的stdout，然后就可以从stdout读出信息
	cmd.Start()
	content, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content))
	fmt.Println(string(content))

	if name, err := os.Hostname(); err == nil {
		fmt.Println(name + "=====>  hostname")

		fmt.Println("runtime====> corekernel number", runtime.GOMAXPROCS(0))
		fmt.Println(runtime.NumCPU())
		fmt.Println()
	}

}
