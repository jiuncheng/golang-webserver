FROM golang:1.15.6-alpine3.12

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go build main.go

CMD ["./main"]