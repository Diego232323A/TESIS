FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o update-review ./src

EXPOSE 5006

CMD ["./update-review"]
