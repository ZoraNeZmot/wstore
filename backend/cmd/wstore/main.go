package main

import (
	_ "wstore/cmd/wstore/docs" // http-swagger middleware
	"wstore/internal/app"

	"github.com/sirupsen/logrus"
)

//	@title			WStore API
//	@version		1.0
//	@description	WStore API

//	@contact.name	API Supports
//	@contact.email	kureichukns@gmail.com

//	@license.name	Apache 2.0
//	@license.url	license.txt

//	@host	localhost:8080

func main() {
	app := app.New()
	if err := app.Init(); err != nil {
		logrus.WithError(err).Errorf("Failed to init app")
	}
	if err := app.Run(); err != nil {
		logrus.WithError(err).Errorf("Failed to run app")
	}
}
