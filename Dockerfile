FROM golang:latest

WORKDIR /go/src

COPY . .

RUN go build -o math

CMD ["./math"]