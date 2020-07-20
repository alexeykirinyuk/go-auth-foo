FROM golang:latest

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go mod download
RUN go build -o main src/cmd.go

CMD ["/app/main"]