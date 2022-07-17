FROM golang:1.16.3

WORKDIR /gateway

COPY . /gateway

RUN go mod download

RUN go build -o main cmd/main.go

EXPOSE 3000

CMD ["./main"]