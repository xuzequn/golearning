package main

import "errors"

// 基础配置类型
type Config struct {
	Endpoint string
}

// 配置操作接口，包含操作集
type ConfigProvider interface {
	GetServiceConfig(serviceName string) (*Config, error)
}

// 包含配置的配置器
type InMemoryConfigProvider struct {
	cfg map[string]*Config
}

func NewInMemoryConfigProvider() *InMemoryConfigProvider {
	return &InMemoryConfigProvider{
		cfg: make(map[string]*Config, 4),
	}
}

var ErrServiceNotFound = errors.New("service not found")

func (i *InMemoryConfigProvider) GetServiceConfig(serviceName string) (*Config, error) {
	cfg, ok := i.cfg[serviceName]
	if !ok {
		return nil, ErrServiceNotFound
	}
	return cfg, nil
}

// var CfgMap map[string]*Config
// func init() {
// 	CfgMap = make(map[string]*Config, 4)
// 	CfgMap["hello"] = &Config{
// 		Endpoint: "http://127.0.0.1:8080/",
// 	}
// }

// type Config struct {
// 	Endpoint string
// }
