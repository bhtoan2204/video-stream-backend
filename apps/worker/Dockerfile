FROM golang:1.23-alpine:latest AS build

WORKDIR /usr/app

COPY . .

RUN go build -o /usr/app/worker /usr/app/cmd/main

FROM alpine:latest AS production

RUN apk add ffmpeg

WORKDIR /usr/app

COPY --from=build /usr/app .

CMD ["./worker"]