package main
import "fmt"

func main()  {
	var arr1 = [5]int {1,2};
	arr1[0] = 11;
	fmt.Printf("arr1: %v\n", arr1);
	//array without lenth
	arr2 := [...]int{}
	fmt.Printf("arr1: %v\n", arr2);

	//inisilize only the 2nd and 3rd element of an array
	arr3 := [5]int{1:10,2:20}
	fmt.Printf("arr3: %v\n", arr3);

	//slice
	var slice1 = []int{};
	slice1 = append(slice1, 1,2,34,5,6);
	for x := range slice1 {
		fmt.Printf("x: %v\n", x)
	}

	for index,value  := range slice1 {
		fmt.Printf("index: %v\n", index);
		fmt.Printf("value: %v\n", value);
	}
}