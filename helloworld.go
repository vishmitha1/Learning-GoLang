package main

import (
	"fmt"
	"reflect"
)

func main()  {
	fmt.Println("HelloWolrd");
	var x1  = "String1";
	var x2 string = "String2";
	var x3 = 3;
	
	fmt.Println(x1,x2,reflect.TypeOf(x3));



}