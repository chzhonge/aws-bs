FROM golang:1.16

WORKDIR /usr/src/app

COPY go.* ./

RUN go mod download

COPY . .

CMD go run main.go