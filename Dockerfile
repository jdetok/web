FROM golang:1.24

WORKDIR /app

RUN go install github.com/air-verse/air@latest
COPY . .
RUN go mod download

ENTRYPOINT [ "air" ]
