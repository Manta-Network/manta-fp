FROM golang:1.23 as builder

WORKDIR /app/manta-fp

COPY . .

RUN make build

FROM debian:bookworm-slim

WORKDIR /app/manta-fp

RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates net-tools curl wget \
    && wget -O /usr/lib/libwasmvm.x86_64.so https://github.com/CosmWasm/wasmvm/releases/download/v2.1.3/libwasmvm.x86_64.so \
    && chmod +x /usr/lib/libwasmvm.x86_64.so \
    && ldconfig \
    && apt-get remove -y wget \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* /var/cache/apt/archives/*

COPY --from=builder /app/manta-fp/build/eotsd .
COPY --from=builder /app/manta-fp/build/bfpd .
COPY --from=builder /app/manta-fp/build/sfpd .
