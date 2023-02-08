package main

import (
	"github.com/chiraponkub/DPU-SosApp-v.1.git/restapi"
	"os"
	"os/signal"
)

func main() {
	ctrl := restapi.NewController()

	err := ctrl.LoadConfigFile()
	if err != nil {
		panic("LoadConfigFile from yml file error: " + err.Error())
	}

	sign := make(chan os.Signal, 1)
	signal.Notify(sign, os.Interrupt, os.Kill)

	go restapi.NewControllerMain(ctrl)
	<-sign
}
