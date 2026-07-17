package main

import (
	_ "wstore/cmd/wstore/docs" // http-swagger middleware
	"wstore/internal/app"
	_ "wstore/internal/handlers/help/response"

	"github.com/sirupsen/logrus"
)

// @title						WStore API
// @version					1.0
// @description				WStore API. Response model is struct{Success bool, Data any, Error *ResponseError}. Only success flag is obligatory.
// @contact.name				API Supports
// @contact.email				kureichukns@gmail.com
//
// @license.name				Apache 2.0
// @license.url				api/license.txt
//
// @schemes					https http
// @BasePath					/api/
//
// @securityDefinitions.apikey	Bearer
// @in							header
// @name						Authorization
// @description				Type "Bearer " followed by your JWT token.
func main() {
	app := app.New()
	if err := app.Init(); err != nil {
		logrus.WithError(err).Errorf("Failed to init app")
	}
	if err := app.Run(); err != nil {
		logrus.WithError(err).Errorf("Failed to run app")
	}
}
