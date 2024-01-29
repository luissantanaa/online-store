FROM golang:1.21.3

WORKDIR /usr/src/app
COPY . .
WORKDIR /usr/src/app/client-service/
RUN go mod tidy
WORKDIR /usr/src/app/store-service/
RUN go mod tidy
WORKDIR /usr/src/app/order-service/
RUN go mod tidy