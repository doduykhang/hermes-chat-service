FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o app ./cmd/rest/main.go

RUN chmod +x /app/app

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/app /app
COPY --from=builder /app/config.json /config.json
COPY --from=builder /app/migration /migration

EXPOSE 8080

CMD [ "/app/app" ]
