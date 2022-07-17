FROM golang:1.16.3

WORKDIR /crud

COPY . /crud

RUN go mod download

RUN go build -o main cmd/main.go

EXPOSE 50052

CMD ["./main"]