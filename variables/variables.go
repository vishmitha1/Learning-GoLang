package main

import (
	"fmt";
	"reflect"
)

func main (){
	// var x1 = 2+5;
	x2 := 15; //use to declare variable
	var x3 int = 12; 
	const c1 = "Cosnatant string"
	
	fmt.Println(reflect.TypeOf(x2));
	fmt.Printf("%T \n",x2);
	fmt.Println(x3,c1);
}
