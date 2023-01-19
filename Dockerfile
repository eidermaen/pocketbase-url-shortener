FROM golang:1.19-buster

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY main.go ./
RUN go build -o start

EXPOSE 8080

CMD ["./start", "serve", "--http", "0.0.0.0:8080"]
