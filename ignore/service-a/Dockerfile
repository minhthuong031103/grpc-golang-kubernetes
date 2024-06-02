FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go build -o golang-service .

EXPOSE 8080

CMD ["./golang-service"]
