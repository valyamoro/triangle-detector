FROM golang:1.26-alpine AS builder
WORKDIR /build

RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /build/triangled ./cmd/triangled

FROM alpine:3.20
WORKDIR /app

RUN apk add --no-cache \
    chromium \
    nss \
    freetype \
    harfbuzz \
    ca-certificates \
    ttf-freefont \
 && mkdir -p /data

COPY --from=builder /build/triangled /usr/local/bin/triangled

ENV CHROME_BIN=chromium \
    CHROME_PATH=chromium \
    DATA_DIR=/data/triangled \
    SYMBOLS=BTCUSDT,ETHUSDT

VOLUME ["/data"]

ENTRYPOINT ["/usr/local/bin/triangled"]
