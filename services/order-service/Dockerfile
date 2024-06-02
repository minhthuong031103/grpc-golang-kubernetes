FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go build -o order-service .

EXPOSE 50051

CMD ["./order-service"]
