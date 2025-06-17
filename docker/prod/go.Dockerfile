FROM golang:1.24.3

WORKDIR /go-api

COPY . .
RUN go mod download

RUN go build -o api ./api

CMD ["./api"]