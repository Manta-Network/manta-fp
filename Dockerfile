FROM golang:1.23-alpine3.21 as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git libc6-compat build-base

WORKDIR /app/manta-fp

COPY . .

RUN make build

FROM alpine:3.21

WORKDIR /app/manta-fp

RUN apk add --no-cache ca-certificates busybox-extras

COPY --from=builder /app/manta-fp/build/eotsd .
COPY --from=builder /app/manta-fp/build/bfpd .
COPY --from=builder /app/manta-fp/build/sfpd .
