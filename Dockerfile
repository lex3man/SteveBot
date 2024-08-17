FROM golang:1.22-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main cmd/app/main.go

CMD ["./main"]
