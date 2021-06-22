package main

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Endpoint string
}

type YamlConfigProvider struct {
	// 作为名字到配置的映射
	Services map[string]*Config `yaml:"service"`
}

var ErrServiceNotFound = errors.New("service not found")

func (y *YamlConfigProvider) GetServiceConfig(serviceName string) (*Config, error) {
	cfg, ok := y.Services[serviceName]
	if !ok {
		return nil, ErrServiceNotFound
	}
	return cfg, nil
}

func main() {
	content, err := os.ReadFile("/Users/xuzequn/OperatingPlatform/Learning/go进阶训练营/golearning/grammar/sparrow/test/client.yaml")
	if err != nil {
		fmt.Printf("%+v", err)
	}
	ycp := &YamlConfigProvider{}
	text := string(content)
	fmt.Println(text)
	err = yaml.Unmarshal(content, ycp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ycp)

	cfg, err := ycp.GetServiceConfig("hello")
	fmt.Println(cfg)
	if err != nil {
		fmt.Println(err)
	}

}
