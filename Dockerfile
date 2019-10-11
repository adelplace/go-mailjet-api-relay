FROM golang:alpine3.10

RUN apk add git

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build cmd/web/*

EXPOSE 8080

CMD ["./application"]
