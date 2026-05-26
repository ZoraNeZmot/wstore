package main

import (
	"wstore/internal/app"

	"github.com/sirupsen/logrus"
)

func main() {
	app := app.New()
	if err := app.Init(); err != nil {
		logrus.WithError(err).Errorf("Failed to init app")
	}
	if err := app.Run(); err != nil {
		logrus.WithError(err).Errorf("Failed to run app")
	}
}
