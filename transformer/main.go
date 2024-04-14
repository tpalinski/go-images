package main

import (
	"transformer/rabbit"
	"transformer/utils"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Starting image transformer")
	err := rabbit.InitRabbitConnection(5, 10, &rabbit.DefaultHandler{});
	utils.PanicOnError(err, "Something went horribly wrong")
}
