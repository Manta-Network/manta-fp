FROM golang:1.23 as builder

WORKDIR /app/manta-fp

COPY . .

RUN make build

FROM debian:bullseye-slim

WORKDIR /app/manta-fp

RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates net-tools \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* /var/cache/apt/archives/*

COPY --from=builder /app/manta-fp/build/eotsd .
COPY --from=builder /app/manta-fp/build/bfpd .
COPY --from=builder /app/manta-fp/build/sfpd .
