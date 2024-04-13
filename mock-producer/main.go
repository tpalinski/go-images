package main

import (
	imgloader "mock-producer/img_loader"
	"mock-producer/rabbit"
	"mock-producer/utils"

	log "github.com/sirupsen/logrus"
)

const IMAGE_PATH = "./data/cool-dog.jpg"
const RETRIES = 10;
const TIMEOUT = 5;

func main() {
	log.Info("Starting producer")
	img, err := imgloader.LoadImage(IMAGE_PATH);
	utils.PanicOnError(err, "Error while reading image from disk");
	deser, err := imgloader.DeserializeImage(img, "cool-dog");
	utils.PanicOnError(err, "Error while creating protobuf message")
	log.Info("Image loaded.")
	rabbit.InitRabbitConnection(RETRIES, TIMEOUT)
	log.Info("Starting sending messages")
	rabbit.SendMessage(deser);
	rabbit.CloseConnection()
}
