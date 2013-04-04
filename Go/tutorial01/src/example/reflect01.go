package main

import (
	"fmt"
	"reflect"
)

type MethodFunc interface{}

type Person struct {
	Name      string "namestr"
	Age       int    "ageint"
	firstName string "firstName"
	lastName  string "lastName"
}

type Worker struct {
	*Person
}

func (p *Person) Add() {
	fmt.Println("Person.Add")
}
func (p *Person) Get(name string) (person *Person) {
	return
}
func (p *Person) Save() bool {
	fmt.Println("Person.Save")
	return false
}
func (p *Person) Update(name string, args ...string) bool {
	fmt.Println("Person.Update")
	return false
}

func InspectMethodInfo(m reflect.Method) {
	fmt.Println("--------------------------Method Info-------------------------------")
	fmt.Printf("Name:%v\n", m.Name)
	fmt.Printf("PkgPath:%v\n", m.PkgPath)
	fmt.Printf("Func:%v\n", m.Func)
	fmt.Printf("Index:%v\n", m.Index)
	fmt.Printf("Type:%v\n", m.Type)
}

func InspectFieldInfo(f reflect.StructField) {
	fmt.Println("--------------------------StructField Info-------------------------------")
	fmt.Printf("Name:%v\n", f.Name)
	fmt.Printf("PkgPath:%v\n", f.PkgPath)
	fmt.Printf("Index:%v\n", f.Index)
	fmt.Printf("Type:%v\n", f.Type)
	fmt.Printf("Anonymous :%v\n", f.Anonymous)
	fmt.Printf("IsPtr :%v\n", f.Type.Kind() == reflect.Ptr)
}

func InspectFuncInfo(m MethodFunc) {
	fmt.Println("========================Func Info==================================")
	fmt.Printf("method:%v\n", m)
	methodType := reflect.TypeOf(m)
	fmt.Printf("typeOf: %v \n", methodType)
	fmt.Printf("name: %v\n", methodType.Name())
	// reflect.Func
	fmt.Printf("kind: %v\n", methodType.Kind())
	// number of returned Args
	numOut := methodType.NumOut()
	fmt.Printf("numOut: %v\n", numOut)
	// number of passed in Args
	numIn := methodType.NumIn()
	fmt.Printf("numIn: %v\n", numIn)
	// is there ... args
	fmt.Printf("isVariadic: %v\n", methodType.IsVariadic())

	for i := 0; i < numIn; i++ {
		fmt.Printf("In: %v, type=%v, name=%v\n", i, methodType.In(i), methodType.In(i).Name())
	}

	for i := 0; i < numOut; i++ {
		fmt.Printf("Out: %v, type=%v, name=%v\n", i, methodType.Out(i), methodType.Out(i).Name())
	}
}

func InspectPerson(p *Person) {
	fmt.Println("===================Inspect Person=====================")
	t := reflect.TypeOf(p)
	tm := t.Elem()

	// type is *Person, elem is Person
	fmt.Printf("name:%v, type:%v, kind:%v, elem:%v\n", t.Name(), t, t.Kind(), tm)

	numMethod := t.NumMethod()

	fmt.Printf("====================Inspect Method Info, numMethod = %v\n", numMethod)
	for i := 0; i < numMethod; i++ {
		InspectMethodInfo(t.Method(i))
	}

	numField := tm.NumField()
	fmt.Printf("====================Inspect Field info, numField = %v\n", numField)
	for i := 0; i < numField; i++ {
		InspectFieldInfo(tm.Field(i))
	}
}

func InspectWorker(p *Worker) {
	fmt.Println("===================Inspect Worker=====================")
	t := reflect.TypeOf(p)
	tm := t.Elem()

	// type is *Person, elem is Person
	fmt.Printf("name:%v, type:%v, kind:%v, elem:%v\n", t.Name(), t, t.Kind(), tm)

	numMethod := t.NumMethod()

	fmt.Printf("====================Inspect Method Info, numMethod = %v\n", numMethod)
	for i := 0; i < numMethod; i++ {
		InspectMethodInfo(t.Method(i))
	}

	numField := tm.NumField()
	fmt.Printf("====================Inspect Field info, numField = %v\n", numField)
	for i := 0; i < numField; i++ {
		InspectFieldInfo(tm.Field(i))
	}
}

func init() {
	fmt.Println("init...")
}

func main() {
	fmt.Println("main...")

	p := &Person{}
	InspectPerson(p)

	w := &Worker{}
	InspectWorker(w)

	//InspectFuncInfo((*Person).Add) // pass in method instance
	//InspectFuncInfo((*Person).Get)
	//InspectFuncInfo((*Person).Save)
	//InspectFuncInfo((*Person).Update)
}
