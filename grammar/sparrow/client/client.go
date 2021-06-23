package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

// var Cfmap map[string]*Config

// func init() {
// 	CfgMap = make(map[string]*Config, 4)
// 	CfgMap["hello"] = &Config{
// 		Endpoint:=
// 	}
// }

func main() {
	// h := &hello{
	// 	endpoint: "http://127.0.0.1:8080/",
	// }
	// msg, err := h.SayHello("golang")

	// if err != nil {
	// 	fmt.Printf("%+v", err)
	// }

	// fmt.Println(msg)

	// // PrintFuncName(h)
	// SetFuncField(h)
	// h.FuncField("golang")

	defer func() {
		if data := recover(); data != nil {
			str := data.(string)
			fmt.Println(str)
		}
	}()

	fcg, err := NewYamlConfigProvider("your path")
	if err != nil {
		panic("初始化配置失败")
	}
	err = InitApplication(WithCfgProvider(fcg))
	if err != nil {
		// 不可挽回的错误，直接崩掉
		panic("初始化应用失败")
	}

}

type HelloService interface {
	SayHello(name string) (string, error)
}

type hello struct {
	endpoint  string
	FuncField func(name string) (string, error) //通过反射可以修改
	// GetUser(req *UserReq) (*User, error)
}

// 通过反射改不了, 这是实现interface的方法
func (h hello) SayHello(name string) (string, error) {
	// panic ("inplement me")
	client := http.Client{}
	resp, err := client.Get(fmt.Sprintf(h.endpoint + name))
	if err != nil {
		fmt.Printf("%+v", err)
		return "", nil
	}
	date, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%+v", err)
		return "", err
	}
	return string(date), nil
}

// 通过反射获取原方法信息
func PrintFuncName(val interface{}) {
	// 反射 reflection
	t := reflect.TypeOf(val)
	t.NumMethod()
	t.NumField()
	v := reflect.ValueOf(val)
	for i := 0; i < t.NumField(); i++ {
		// fmt.Println(t.Method(i).Name)
		fmt.Println(t.Field(i).Name)
		field := t.Field(i)
		fieldValue := v.Field(i)
		fmt.Println(field.Name)
		fmt.Println(fieldValue.CanSet())
	}

}

// 通过反射篡改原方法
// func SetFuncField(val interface{}) {
// 	// 反射 reflection.Valueof
// 	v := reflect.ValueOf(val) //这是指针的反射
// 	ele := v.Elem()           // 拿到了指针指向的结构体
// 	t := ele.Type()           // 拿到了指针指向的结构体的类型信息

// 	num := t.NumField()
// 	for i := 0; i < num; i++ {
// 		f := ele.Field(i)
// 		if f.CanSet() {
// 			fn := func(args []reflect.Value) (results []reflect.Value) {
// 				name := args[0].Interface().(string)
// 				fmt.Printf("这是一个篡改的方法")
// 				client := http.Client{}
// 				serviceName := val.(Service).ServiceName()

// 				endpoint := CfgMap[serviceName].Endpoint

// 				resp, err := client.Get(endpoint + name)
// 				if err != nil {
// 					fmt.Printf("%+v", err)
// 					return []reflect.Value{reflect.ValueOf(""), reflect.ValueOf(err)}
// 				}
// 				data, err := ioutil.ReadAll(resp.Body)
// 				if err != nil {
// 					fmt.Printf("%+v", err)
// 					return []reflect.Value{reflect.ValueOf(""), reflect.ValueOf(err)}
// 				}
// 				fmt.Println(string(data))
// 				return []reflect.Value{reflect.ValueOf(string(data)), reflect.Zero(reflect.TypeOf(new(error)).Elem())}

// 			}
// 			f.Set(reflect.MakeFunc(f.Type(), fn))
// 		}

// 	}
// }

func SetFuncField(val Service) {
	// 反射 reflection.Valueof
	v := reflect.ValueOf(val) //这是指针的反射
	ele := v.Elem()           // 拿到了指针指向的结构体
	t := ele.Type()           // 拿到了指针指向的结构体的类型信息

	num := t.NumField()
	for i := 0; i < num; i++ {
		field := t.Field(i)
		f := ele.Field(i)
		if f.CanSet() {
			fn := func(args []reflect.Value) (results []reflect.Value) {
				in := args[0].Interface()
				out := reflect.New(field.Type.Out(0).Elem()).Interface()
				inData, err := json.Marshal(in)

				if err != nil {
					return []reflect.Value{reflect.ValueOf(out), reflect.ValueOf(err)}
				}
				fmt.Printf("这是一个篡改的方法")
				client := http.Client{}

				name := val.ServiceName()

				cfg, err := App.CfgProvider.GetServiceConfig(name)

				if err != nil {
					return []reflect.Value{reflect.ValueOf(out), reflect.ValueOf(err)}
				}
				req, err := http.NewRequest("POST", cfg.Endpoint, bytes.NewReader(inData))

				if err != nil {
					return []reflect.Value{reflect.ValueOf(out), reflect.ValueOf(err)}
				}

				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("sparrow-service", name)
				req.Header.Set("sparrow-service-method", field.Name)

				resp, err := client.Do(req)

				if err != nil {
					fmt.Printf("%+v", err)
					return []reflect.Value{reflect.ValueOf(""), reflect.ValueOf(err)}
				}
				data, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					return []reflect.Value{reflect.ValueOf(out), reflect.ValueOf(err)}
				}
				err = json.Unmarshal(data, out)
				if err != nil {
					return []reflect.Value{reflect.ValueOf(out), reflect.ValueOf(err)}
				}
				fmt.Println(string(data))
				return []reflect.Value{reflect.ValueOf(out), reflect.Zero(reflect.TypeOf(new(error)).Elem())}

			}
			f.Set(reflect.MakeFunc(f.Type(), fn))
		}
	}
}

type Service interface {
	ServiceName() string
}

func (h *hello) ServiceName() string {
	return "hello"

}

var ErrorServiceNotFound = errors.New("service not found") // sentiel error 预定义错误
