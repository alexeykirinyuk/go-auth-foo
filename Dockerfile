FROM golang:latest

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go mod download
RUN go build -o go-auth-foo-cmd main/cmd.go main/migrations.go main/routes.go

CMD ["/app/go-auth-foo-cmd"]