FROM golang:1.24.3

RUN apt-get update && apt-get install -y nano

WORKDIR /go-api

COPY . .
RUN go mod download

RUN go build -o web ./api/*.go

CMD ["./web"]