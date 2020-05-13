FROM golang:1.14-alpine3.11 AS builder

COPY ${PWD} /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -o /app/appbin *.go

FROM alpine:3.11

RUN apk --update add ca-certificates && \
    rm -rf /var/cache/apk/*

RUN adduser -D appuser
USER appuser

COPY --from=builder /app /home/appuser/app

WORKDIR /home/appuser/app

EXPOSE 8080

CMD ["./appbin"]