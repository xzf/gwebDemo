package main

func main() {

}

type Inter interface {
	A()
}

type ObjOne struct {
	Id string
}

type ObjTwo struct {
	ObjOne
	TwoStr string
}