FROM golang:latest

WORKDIR $GOPATH/src/github.com/coffee_backend
COPY . $GOPATH/src/github.com/coffee_backend
RUN go build .

EXPOSE 9000
ENTRYPOINT ["./coffee_backend"]