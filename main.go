/*
 * Author : xzf
 * Time    : 2020-04-26 00:50:22
 * Email   : xpoony@163.com
 */

package main

import (
	"fmt"
	"github.com/xzf/gweb"
	"net/http"
	"time"
	"encoding/json"
	"io/ioutil"
)

type web struct {
	gweb.WebApi
}

//func(a web)FuncA(){
//	fmt.Println(111)
//}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path
		var result = `path : ` + path + `
get query : `
		query, _ := json.MarshalIndent(request.URL.Query(), "", "\t")
		result += string(query) + `
body : `
		body, _ := ioutil.ReadAll(request.Body)
		result += string(body) + `
post query : `
		fmt.Println("", )
		err := request.ParseForm()
		if err == nil {
			postQuery, _ := json.MarshalIndent(request.PostForm, "", "\t")
			result += string(postQuery)
		}
		writer.Write([]byte(result))
	})
	http.ListenAndServe(":2333", nil)
	for {
		time.Sleep(time.Second)
	}
	//w := web{}
	//gweb.NewHttpServer(":2333",w)
	//fmt.Println("------------")
	//gweb.ParseWebApiObj(w)
	//gweb.NewHttpServer("", &w)
	//fmt.Println("------------")
	//w := web{}
	//
	//t := reflect.ValueOf(w)
	//for i := 0; i < t.NumMethod(); i++ {
	//	method := t.Method(i)
	//	in:=[]reflect.Value{}
	//	for ii := 0; ii < method.Type().NumIn(); ii++ {
	//		tt := method.Type().In(ii)
	//		switch tt.Kind() {
	//		case reflect.Struct:
	//			for iti:=0;iti<tt.NumField();iti++{
	//				itf:=tt.Field(iti)
	//				fmt.Println(tt.Name(),itf.Name,itf.Type)
	//
	//			}
	//		default:
	//			fmt.Println("unsupport method para type",tt.Kind().String())
	//			return
	//		}
	//		//p:=reflect.New(tt)
	//
	//		in = append(in, )
	//		fmt.Println(tt.Name())
	//	}
	//	//method.Call([]reflect.Value{
	//	//
	//	//})
	//}
}

type AReq struct {
	As string
	Ai int
}

func (w web) A(req AReq) {
	fmt.Println(req)
}

type BReq struct {
	Bs string
	Bi int
}

func (w web) B(req BReq) {
	fmt.Println(req)
}
