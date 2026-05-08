FROM golang:1.26-alpine AS builder
WORKDIR /build

RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /out/triangled ./cmd/triangled

FROM alpine:3.20
WORKDIR /app

RUN apk add --no-cache \
    chromium \
    nss \
    freetype \
    harfbuzz \
    ca-certificates \
    ttf-freefont

COPY --from=builder /out/triangled /app/triangled

ENV CHROME_BIN=/usr/bin/chromium-browser \
    CHROME_PATH=/usr/bin/chromium-browser \
    DATA_DIR=/data \
    SYMBOLS=BTCUSDT,ETHUSDT

VOLUME ["/data"]

ENTRYPOINT ["/app/triangled"]
