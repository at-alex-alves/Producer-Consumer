FROM golang:1.17.6-alpine3.15

COPY . /app

WORKDIR /app

RUN go build -o main ./src

CMD ["/app/main"]
