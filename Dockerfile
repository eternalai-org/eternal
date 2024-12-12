FROM  golang:1.21-alpine3.17 AS build

RUN apk update && apk add gcc musl-dev gcompat libc-dev linux-headers
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=arm64 go build -ldflags "-linkmode external -extldflags -static" -o workersv

RUN chmod +x /app/workersv

FROM  alpine:3.17 AS export-stage

RUN  apk update && apk add --no-cache ca-certificates && update-ca-certificates

COPY --from=build /app/workersv .
