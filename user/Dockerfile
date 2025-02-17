FROM golang:1.23-alpine AS build

WORKDIR /usr/app

COPY . .

RUN go build -o /usr/app/user /usr/app/cmd/main

FROM alpine:latest AS production

WORKDIR /usr/app

COPY --from=build /usr/app .

CMD ["./user"]