package main

import (
	os "io/ioutil"

	"github.com/astaxie/beego/config/yaml"
	"gopkg.in/yaml.v2"
)

type YamlConfigProvider struct {
	Service map[string]*Config `yaml:"service"`
}

func NewYamlConfigProvider(filepath string) (*YamlConfigProvider, error) {
	// content 是一个字节流
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	ycp := &YamlConfigProvider{}
	// contnet 数据注入到我们的结构体实例
	err = yaml.Unmarshal(content, ycp)
	return ycp, err
}
