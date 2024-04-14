package main

import (
	"os"
	"transformer/rabbit"
	"transformer/utils"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Starting image transformer")
	os.Mkdir("data", 0777);
	log.Info("Created output directory")
	err := rabbit.InitRabbitConnection(5, 10, &rabbit.DefaultHandler{});
	utils.PanicOnError(err, "Something went horribly wrong")
}
