package main

import "sync"

var App *application

var appOnce sync.Once

type application struct{}

func InitApplication(opts ...AppOption) error {
	var err error
	appOnce.Do(func() {
		App = &application{
			CfgProvider: NewInMemoryConfigProvider(),
		}
	})
}
