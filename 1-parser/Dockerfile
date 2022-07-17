FROM golang:1.16.3

WORKDIR /parser

COPY . /parser

RUN go mod download

RUN go build -o main cmd/main.go

EXPOSE 50051

CMD ["./main"]