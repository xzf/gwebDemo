/*
 * Author : xzf
 * Time    : 2020-04-26 00:50:22
 * Email   : xpoony@163.com
 */

package main

import (
	"encoding/json"
	"fmt"
	"github.com/xzf/gweb"
	"net/http"
	"reflect"
)

type web struct {
	gweb.WebApi
}

//func(a web)FuncA(){
//	fmt.Println(111)
//}
type SampleRequestParameter struct {
	QueryFormMap map[string][]string `json:",omitempty"`
	PostFormMap  map[string][]string `json:",omitempty"`
	Body         []byte              `json:",omitempty"`
	Path         string              `json:",omitempty"`
	Request      *http.Request       `json:"-"`
	// cookie
	// header
	// session
}

func (para SampleRequestParameter) ToJsonByteSlice() (result []byte) {
	var err error
	result, err = json.Marshal(para)
	if err != nil {
		return
	}
	return
}

func main() {
	slice:=[]int{1,2,3}
	t:=reflect.TypeOf(slice)
	//v:=reflect.ValueOf(slice)
	fmt.Println(t,t.Elem())

	slice1:=[4]int{}
	t=reflect.TypeOf(slice1)
	fmt.Println(t.Kind())
	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	var para SampleRequestParameter
	//	para.Path = request.URL.Path
	//	queryMap := request.URL.Query()
	//	if len(queryMap) != 0 {
	//		para.QueryFormMap = queryMap
	//	}
	//	body, err := ioutil.ReadAll(request.Body)
	//	if err == nil && len(body) != 0 {
	//		para.Body = body
	//	}
	//	err = request.ParseForm()
	//	postForm := request.PostForm
	//	if err == nil && len(postForm) != 0 {
	//		para.PostFormMap = postForm
	//	}
	//	_, _ = writer.Write(para.ToJsonByteSlice())
	//})
	//http.ListenAndServe(":2333", nil)
	//for {
	//	time.Sleep(time.Second)
	//}

	/*
	 */

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
