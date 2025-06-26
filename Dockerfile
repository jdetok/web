FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum .env ./

RUN go mod download

COPY ./api ./api
COPY ./internal ./internal
COPY ./www ./www
# COPY ./www/src/*.html ./www/src/*.html
# COPY ./www/src/js ./www/src/js
# COPY ./www/src/css ./www/src/css
# COPY ./www/src/img/pets ./www/src/img/pets
# COPY ./www/src/img/bronto ./www/src/img/bronto

RUN go build -o bin ./api

# entrypoint for the binary will be /app/bin
ENTRYPOINT [ "/app/bin" ]
