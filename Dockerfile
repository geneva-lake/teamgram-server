FROM golang:1.20 AS builder
ENV GOMEMLIMIT=1150MiB
ENV GOMAXPROCS=1
WORKDIR /app
COPY . .
RUN ./build.sh

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/teamgramd/ /app/
RUN apk update && apk add ffmpeg && chmod +x /app/docker/entrypoint.sh
RUN apk add bash
ENTRYPOINT /app/docker/entrypoint.sh
