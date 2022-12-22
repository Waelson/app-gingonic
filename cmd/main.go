package main

import (
	"github.com/Waelson/internal/app"
)

func main() {
	configs := getConfigs()

	application := app.NewApplication(configs...)
	error := application.Start()

	if error != nil {
		panic(error)
	}

}

func getConfigs() []app.Option {
	configs := make([]app.Option, 0, 0)
	configs = append(configs, app.WithPort(8080))
	return configs
}
