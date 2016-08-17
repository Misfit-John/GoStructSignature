package StructSignature

import "fmt"

func ExampleComplicateSignature() {
	ret := GetSignature(EmbbedStruct{})
	fmt.Printf("%s", ret)
}
