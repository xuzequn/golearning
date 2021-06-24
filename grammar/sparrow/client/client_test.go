package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetFuncFied(t *testing.T) {

	path := `..\test\client.yaml`
	ycp, _ := NewYamlConfigProvider(path)

	_ = InitApplication(WithCfgProvider(ycp))

	helloService := &hello{}

	SetFuncField(helloService)

	res, err := helloService.SayHello(&Input{
		Name: "golang",
	})
	// if err != nil {
	// 	t.FailNow()
	// }

	assert.Nil(t, err)
	assert.Equal(t, "Hello, golang", res.Msg)

	// fmt.Print("aaaaaaaaaaa")

	// assert.Nil(t, hello{endpoint: "http://127.0.0.1:8080/"}, "bbbbccc")

}

type hello struct {
	SayHello func(in *Input) (*Output, error)
}

func (h *hello) ServiceName() string {
	return "hello"
}

type Input struct {
	Name string
}

type Output struct {
	Msg string
}

// func Test
