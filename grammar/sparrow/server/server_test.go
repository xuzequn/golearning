package main

import (
	"fmt"
)

// func Testzuhe(t *testing.T) {
// 	c := &Concrete{
// 		Base: Base{
// 			Name: "Tome",
// 		},
// 	}
// 	// c.Name
// 	// c.Base.Name
// 	// c.Hello()
// 	// c.Base.Hello()

// }

type Base struct {
	// Name string
}

type Concrete struct {
	Base
	Age int
}

type Concrete1 struct {
	*Base
}

func (b Base) SayHello() {
	fmt.Printf("Base say hello " + b.Name())
}

func (b Base) Name() string {
	return "Base"
}

func (c Concrete1) Name() string {
	return "Concrete"

}

// func main() {
// 	base := Base{}
// 	c := Concrete1{
// 		Base: &base,
// 	}
// 	c.SayHello()
// }
