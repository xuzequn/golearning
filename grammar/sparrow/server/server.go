package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	// 注册服务
	AddService(&userService{})
	AddService(&helloService{})

	go func() {
		listenSignal()
	}()

	http.HandleFunc("/", hander)
	// if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var sysSignals = []os.Signal{os.Interrupt, os.Kill, syscall.SIGKILL,
	syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGILL, syscall.SIGTRAP,
	syscall.SIGABRT, syscall.SIGTERM}

func listenSignal() {
	signals := make(chan os.Signal, 1)

	signal.Notify(signals, sysSignals...)

	select {
	case <-signals:
		forceShutdownIfneed()
		shutdown()
		os.Exit(0)
	}
}

func forceShutdownIfneed() {
	time.AfterFunc(time.Minute, func() {
		os.Exit(1)
	})
}

func shutdown() {
	// 执行各种动作
	services.Range(func(key, value interface{}) bool {
		service := value.(Service)
		go func() {
			service.ShutDown()
		}()
		return true
	})
}
