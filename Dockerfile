FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum .env ./

RUN go mod download

COPY ./api ./api
COPY ./internal ./internal
COPY ./www ./www

RUN go build -o bin ./api

# entrypoint for the binary will be /app/bin
ENTRYPOINT [ "/app/bin" ]
