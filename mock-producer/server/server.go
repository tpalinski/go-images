package server

import (
	"mock-producer/rabbit"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func StartWebServer(message_content []byte) {
        http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		num_string := r.URL.Query().Get("n")
		amount, err := strconv.Atoi(num_string);
		if err != nil {
			amount = 1;
		}
		log.Infof("Sending %d message(s)", amount)
		go handleSend(amount, message_content)
	});
	http.ListenAndServe(":2137", nil)
}

func handleSend(amount int, data []byte) {
	for i:=0; i<amount; i++ {
		rabbit.SendMessage(data);
	}
}
