# Go Images
This project simulates some kind of data procesing pipeline using amqp and go. It consists of a mocked consumer, which sends image data which is then transformed by the transformer service.

## Running the application
The easiest way to run the app is via docker:  

  `docker compose up`

To send messages, simply send a GET request (for instance, via your browser) to `localhost:2137/send?n=$amount_of_messages`
