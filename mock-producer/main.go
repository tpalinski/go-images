package main

import (
	imgloader "mock-producer/img_loader"
	"mock-producer/utils"

	log "github.com/sirupsen/logrus"
)

const IMAGE_PATH = "./data/cool-dog.jpg"

func main() {
	log.Info("Starting producer")
	_, err := imgloader.LoadImage(IMAGE_PATH);
	utils.PanicOnError(err, "Error while reading image from disk");
}
