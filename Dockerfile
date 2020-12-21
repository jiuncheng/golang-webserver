FROM golang:1.15.6-alpine3.12

WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .

RUN go build -o main .
CMD ["main"]