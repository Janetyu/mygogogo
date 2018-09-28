package simple_factory

import (
	"testing"
)

// TestType1 test get police_1 with factory
func TestType1(t *testing.T) {
	newcall := NewCall(1)
	res := newcall.Call("xiaoming")
	if res != "Hi xiaoming , it's time to work" {
		t.Fatal("Type1 test fail")
	}
}

// TestType2 test get police_2 with factory
func TestType2(t *testing.T) {
	newcall := NewCall(2)
	res := newcall.Call("xiaohong")
	if res != "Hello xiaohong , it's time to work" {
		t.Fatal("Type2 test fail")
	}
}
