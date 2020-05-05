/*
 * Author : xzf
 * Time    : 2020-04-26 00:50:22
 * Email   : xpoony@163.com
 */

package main

import (
	"github.com/xzf/gweb"
	"fmt"
	"encoding/json"
)

type web struct {
	gweb.WebApi
}
type Test1 struct {
	Str string
	I   int64
	B   bool
	F   float64
	SS  []string
	IS  []int32
	BS  []bool
	FS  []float64
	SA  [3]string
	IA  [3]int32
	BA  [3]bool
	FA  [3]float64
}
type Test2 struct {
	Test1
	Struct Test1
}
func toJsonByte(obj interface{}) []byte {
	content, _ := json.MarshalIndent(obj, "\t", "")
	return content
}
func (obj web) Api() {
	fmt.Println(obj.GetGoRequest())
	fmt.Println("call no para method")
	obj.WriteBody([]byte("call no para method"))
}
func (obj web) ApiTest1(t Test1) {
	body := toJsonByte(t)
	fmt.Println(string(body))
	obj.WriteBody(body)
}
func (obj web) ApiTest2(t Test2) {
	body := toJsonByte(t)
	fmt.Println(string(body))
	obj.WriteBody(body)
}
func main() {
	gweb.SetDebugMode()
	gweb.NewHttpServer(":2333", &web{})

}
