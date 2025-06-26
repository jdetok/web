FROM golang:1.24

WORKDIR /app

# RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum .env ./

RUN go mod download

COPY . .
ENTRYPOINT [ "air" ]
# COPY ./api ./api
# COPY ./internal ./internal
# COPY ./www ./www

# ENTRYPOINT [ "/bin/sh" ]

# uncomment once ready for prod
# RUN go build -o bin ./api

# ENTRYPOINT [ "/app/bin" ]
