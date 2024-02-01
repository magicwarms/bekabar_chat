FROM golang:1.21.1-alpine AS build-stage

LABEL MAINTAINER="Andhana Utama <andhanautama@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o bekabar-chat .

FROM alpine:latest AS final

WORKDIR /app

COPY --from=build-stage /app/bekabar-chat .

COPY .env ./

EXPOSE 9005

CMD ["./bekabar-chat"]

