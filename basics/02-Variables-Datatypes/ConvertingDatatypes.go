package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// Convert String to Float
	s := "3.1415926535"
	f, _ := strconv.ParseFloat(s, 8)
	fmt.Println(f)

	// Convert String to Boolean
	str := "true"
	b, _ := strconv.ParseBool(str)
	fmt.Println(b)

	// Float to String
	var flo float64 = 3.1415926535
	var strFlo string = strconv.FormatFloat(flo, 'E', -1, 32)
	fmt.Println(reflect.TypeOf(strFlo))

	// Converting float to int
	var f32 float32 = 3.1415926535
	fmt.Println(reflect.TypeOf(f32))

	i32 := int32(f32)
	fmt.Println(reflect.TypeOf(i32))
}
