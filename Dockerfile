
FROM golang:1.16.15-alpine3.15 AS builder

WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz




FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate

COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration

# EXPOSE 8080
# ENV PORT=8080
# ENV HOST=0.0.0.0

# CMD [ "./main" ]
CMD [ "/app/main" ]
# ENTRYPOINT [ "/app/start.sh" ]

