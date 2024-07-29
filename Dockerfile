FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN apt-get update

RUN go build -o main ./cmd/app/main.go

CMD ["./main"]