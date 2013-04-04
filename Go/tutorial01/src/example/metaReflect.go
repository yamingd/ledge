package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name      string "namestr"
	Age       int    "ageint"
	firstName string "firstName"
	lastName  string "lastName"
}

func ShowTag(p interface{}) {
	t := reflect.TypeOf(p)

	switch t.Kind() {
	case reflect.Ptr:
		tag := t.Elem().Field(0).Tag.Get("namstr")
		fmt.Printf("tag is %v\n", tag)
	}
}

func SetValue(i interface{}) {
	v := reflect.ValueOf(i)
	v.Elem().Field(0).SetString("Yamingd")
	v.Elem().Field(1).SetInt(11)
}

func ListFields(i interface{}) {
	t := reflect.TypeOf(i)
	nf := t.Elem().NumField()
	for j := 0; j < nf; j++ {
		f := t.Elem().Field(j)
		fmt.Printf("%v %v %v\n", j, f.Name, f.Type.Name())
	}
}

// test int or string on I
type I interface {
	Get() int
}

func g(any interface{}) int {
	if v, ok := any.(I); ok {
		return v.Get()
	}
	return -1
}

func main() {
	p := new(Person)
	ShowTag(p)
	SetValue(p)
	fmt.Printf("Person is %v\n", p)

	ListFields(p)

	a := 100
	fmt.Printf("a is %v\n", g(a))
}
