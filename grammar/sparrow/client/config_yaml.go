package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type YamlConfigProvider struct {
	// 作为名字到配置的映射
	Services map[string]*Config `yaml:"service"`
}

func NewYamlConfigProvider(filepath string) (*YamlConfigProvider, error) {
	// content 是一个字节流
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	ycp := &YamlConfigProvider{}
	// contnet 数据注入到我们的结构体实例, content 是一个字节切片，
	err = yaml.Unmarshal(content, ycp)
	return ycp, err
}

// 拿出YamlConfigProvider的Service参数下 map的存放的某个serviceName下的配置类型是config
func (y *YamlConfigProvider) GetServiceConfig(serviceName string) (*Config, error) {
	cfg, ok := y.Services[serviceName]
	if !ok {
		return nil, ErrServiceNotFound
	}
	return cfg, nil
}
