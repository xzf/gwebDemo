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
	"runtime"
	"strings"
	"time"
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
	fmt.Println("call no para method")
}

func (obj web) ApiTest1(t Test1) {
	fmt.Println(string(toJsonByte(t)))
}
func (obj web) ApiTest2(t Test2) {
	fmt.Println(string(toJsonByte(t)))
}

func getGoroutineId() string {
	var buf [128]byte
	runtime.Stack(buf[:], false)
	allInfo := string(buf[10:])
	n := strings.Index(allInfo, " ")
	return allInfo[:n]
}
func main() {
	//gweb.SetDebugMode()
	//gweb.NewHttpServer(":2333", web{})

	go func() {
		for{
			fmt.Println("------------ go 1",getGoroutineId())
			time.Sleep(time.Second)
		}
	}()
	go func() {
		for{
			fmt.Println("------------ go 2",getGoroutineId())
			time.Sleep(time.Second)
		}
	}()
	for{
		fmt.Println("------------ go 4",getGoroutineId())
		time.Sleep(time.Second)
	}
}

/*
	//t:=reflect.TypeOf(Test1{})
	//v := reflect.New(t).Elem()
	//v.Field(5).Set(reflect.ValueOf([]int64{1,2,3}))
	//s := v.Interface()
	//fmt.Println(s)
*/

/*

t:=reflect.TypeOf(Test1{})
	v := reflect.New(t).Elem()
	s := v.Addr().Interface()

	var setFunc func(v reflect.Value)
	setFunc= func(v reflect.Value) {
		for i:=0;i<v.NumField();i++{
			fieldValue:=v.Field(i)
			fieldType:=fieldValue.Type()
			switch fieldType.Kind() {
			case reflect.Int64,reflect.Int:
				v.Field(i).SetInt(2333)
			case reflect.String:
				v.Field(i).SetString("str 2333333")
			case reflect.Bool:
				v.Field(i).SetBool(true)
			case reflect.Float64:
				v.Field(i).SetFloat(2.33)
			//case reflect.Slice:
			//	switch fieldType.Elem().Kind() {
			//	case reflect.Int:
			//		v.Field(i).Set(reflect.ValueOf([]int{2,3,3}))
			//	case reflect.String:
			//		v.Field(i).Set(reflect.ValueOf([]string{"2","3","3"}))
			//	case reflect.Bool:
			//		v.Field(i).Set(reflect.ValueOf([]int{2,3,3}))
			//	case reflect.Float64:
			//	case reflect.Slice:
			//
			//	case reflect.Array:
			//
			//		//case reflect.Struct:
			//		//	v:= reflect.New(fieldType).Elem()
			//		//	setFunc(v)
			//	}
			//case reflect.Array:

			//case reflect.Struct:
			//	v:= reflect.New(fieldType).Elem()
			//	setFunc(v)
			}
		}
	}
	setFunc(v )
	v.Field(4).SetString("2333")


	s = v.Addr().Interface()

	fmt.Println(s)
*/

//slice:=[]int{1,2,3}
//t:=reflect.TypeOf(slice)
////v:=reflect.ValueOf(slice)
//fmt.Println(t,t.Elem())
//
//slice1:=[4]int{}
//t=reflect.TypeOf(slice1)
//fmt.Println(t.Kind())

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
/*
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
*/

/*

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

func testFunc(){
	fmt.Println("call testFunc")
}
*/
