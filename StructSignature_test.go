package StructSignature

import "testing"
import "reflect"
import "fmt"

type TestStruct struct {
	i  int
	s  string
	f  float32
	ff float64
}

type EmbbedStruct struct {
	t TestStruct
	i int
	s []int
	m map[int]string
}

func TestInterfaceName(t *testing.T) {
	ret := searchInterface(reflect.TypeOf(TestStruct{}))
	if "StructSignature.TestStruct:i:2s:24f:13ff:14" != ret {
		t.Errorf("wrong signature string %s", ret)
	}
	t.Logf("%s", ret)
}

func TestGetSignature(t *testing.T) {
	ret := GetSignature(TestStruct{})
	if "19e61b60e670dfe0f5913ac4c5875590" != ret {
		t.Errorf("Wrong signature %s", ret)
	}
	t.Logf("%s", ret)
}

func TestComplicateStruct(t *testing.T) {
	ret := searchInterface(reflect.TypeOf(EmbbedStruct{}))
	if "StructSignature.EmbbedStruct:t:StructSignature.TestStruct:i:2s:24f:13ff:14i:2s:23:2m:2_24" != ret {
		t.Errorf("wrong signature ret %s", ret)
	}
	t.Logf("%s", ret)
}

func TestComplicateSignature(t *testing.T) {
	ret := GetSignature(EmbbedStruct{})
	if "9c2943a7d6d4f8d64128698493d4e99a" != ret {
		t.Errorf("Wrong signature %s", ret)
	}
	t.Logf("%s", ret)
}

func ExampleGetSignature() {

	fmt.Printf("%s", GetSignature(EmbbedStruct{}))
	// Output: 9c2943a7d6d4f8d64128698493d4e99a
}
