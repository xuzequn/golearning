package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Path: ", r.URL.Path)
	fmt.Fprintf(w, "Hello, %s", r.URL.Path[1:])
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.Path[1:])
	fmt.Println(r.URL.Path)
}

func hander(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	serviceName := r.Header.Get("sparrow-service")
	methodName := r.Header.Get("sparrow-service-method")

	filterIvk := &filterInvoker{
		Invoker: &httpInvoker{},
		filters: []Filter{logFilter},
	}
	output, _ := filterIvk.Invoke(&Invocation{
		MethodName:  methodName,
		ServiceName: serviceName,
		Input:       data,
	})
	fmt.Fprintf(w, "%s", string(output))

}

func main() {
	AddService(&userService{})
	AddService(&helloService{})

	// go func() {
	// 	listenSignal()
	// }()

	http.HandleFunc("/", hander)
	// if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// func listenSignal(){
// 	signals := make(chan os.Signal, 1)

// 	signal.Notify(signal, sysSignals..)

// 	select{
// 	case <- signals:
// 		forceShutdownIfneed()
// 		shutdown()
// 		os.Exit(0)
// 	}
// }
