# Builder
FROM golang:1.15.3-alpine3.12 as builder

RUN apk update && \
    apk upgrade && \
    apk --update add git make

WORKDIR /cmd

COPY . .

RUN make build

# Distribution
FROM alpine:latest

RUN apk update && \
    apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app

WORKDIR /app

EXPOSE 8080

COPY --from=builder /cmd/ppserver /app

CMD /app/ppserver