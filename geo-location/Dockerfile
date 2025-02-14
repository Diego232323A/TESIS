FROM golang:1.20

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o geo-service

CMD ["./geo-service"]

EXPOSE 5012

