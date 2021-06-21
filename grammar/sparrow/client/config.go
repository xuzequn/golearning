package main

var CfgMap map[string]*Config

func init() {
	CfgMap = make(map[string]*Config, 4)
	CfgMap["hello"] = &Config{
		Endpoint: "http://127.0.0.1:8080/",
	}
}

type Config struct {
	Endpoint string
}
