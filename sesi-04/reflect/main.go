package main

import (
	"fmt"
	"reflect"
	"strings"
)

type student struct {
	Name  string
	Grade int
}

func main () {
	var number = 27
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("tipe variebel :", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variebel", reflectValue.Int())
	}
	fmt.Println(strings.Repeat("-", 25))

	fmt.Println("Tipe variable :", reflectValue.Interface())

	nilai := reflectValue.Interface().(int)

	_ = nilai
	fmt.Println(strings.Repeat("-", 25))

	var s1 = &student{Name: "Sultan", Grade: 2}
	fmt.Println("nama :", s1.Name)

	var reflectValue1 = reflect.ValueOf(s1)
	var method = reflectValue1.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("Eka"),
	})

	fmt.Println("nama :", s1.Name)
	fmt.Println(strings.Repeat("-", 25))

}

func (s *student) Setname(name string) {
	s.Name = name
}