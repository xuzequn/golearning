package main

import "sync"

// 一个对象定义，有一个配置文件操作接口
type application struct {
	CfgProvider ConfigProvider
}

// App 必须要显示使用 InitApplication 来创建
var App *application
var appOnce sync.Once

//  返回方法的定义，定义了一个类型， 是一个方法的类型
type AppOption func(app *application) error

var _ AppOption = App1

func App1(app *application) error {
	return nil
}

// 初始化应用，接收配置，初始化应用
func InitApplication(opts ...AppOption) error {
	var err error
	appOnce.Do(func() {
		App = &application{
			CfgProvider: NewInMemoryConfigProvider(),
		}
		// 初始化多个app操作
		for _, opt := range opts {
			err = opt(App)
			if err != nil {
				return
			}
		}
	})
	return err
}

// 返回值为一个方法， 生成初始化应用的配置，配置组装
func WithCfgProvider(cfg ConfigProvider) AppOption {
	// 返回值的实现，一个application对象
	return func(app *application) error {
		app.CfgProvider = cfg
		return nil
	}
}
