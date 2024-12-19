FROM golang:1.23-alpine AS builder
WORKDIR /go/src/app

Run apk update && apk upgrade && \
  apk --update add git gcc make libc-dev openssh linux-headers

COPY . .
RUN go mod tidy
RUN go install github.com/swaggo/swag/cmd/swag@v1.16.3
RUN swag init --parseDependency --parseInternal -g cmd/main.go
RUN go build -o build/main cmd/*.go

FROM alpine:latest AS release
WORKDIR /app

Run apk update && apk upgrade && \
  apk --update add git gcc make libc-dev openssh linux-headers && \
  apk --update --no-cache add chromium-chromedriver chromium-swiftshader chromium

COPY --from=builder /go/src/app/build/main main

ENTRYPOINT [ "./main", "-env=production" ]
