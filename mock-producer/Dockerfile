FROM golang:1.18
COPY . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/producer
EXPOSE 2137
CMD /app/producer
