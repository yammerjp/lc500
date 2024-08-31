FROM golang:1.23 AS builder

COPY . /app
WORKDIR /app
RUN make build

FROM debian:bookworm-slim
RUN apt update  -qqy && \
    apt install -qqy  ca-certificates && \
    update-ca-certificates \
    && apt clean \
    && rm -rf /var/lib/apt/lists/*
COPY --from=builder /app/lc500 /bin/lc500

CMD ["/bin/lc500"]
