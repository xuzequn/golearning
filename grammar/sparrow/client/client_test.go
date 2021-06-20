package main

import (
	"testing"
)

func TestSetFuncFied(t *testing.T) {
	helloService := &hello{endpoint: "http://127.0.0.1:8080/"}

	SetFuncField(helloService)

	res, err := helloService.SayHello("golang")
	if err != nil {
		t.FailNow()
	}

	if res != "Hello, golang" {
		t.FailNow()
	}
	// fmt.Print("aaaaaaaaaaa")

	// assert.Nil(t, hello{endpoint: "http://127.0.0.1:8080/"}, "bbbbccc")

}
