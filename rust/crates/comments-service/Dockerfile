FROM lukemathwalker/cargo-chef:latest-rust-1.83.0 AS chef
WORKDIR /app
RUN apt update && apt upgrade -y --no-install-recommends &&\
    apt install -y --no-install-recommends lld clang openssl libssl-dev cmake

#--------------------------------------------------------------------

FROM chef AS planner
COPY . .
RUN cargo chef prepare --recipe-path recipe.json --bin comments-service

#--------------------------------------------------------------------

FROM chef AS builder
COPY --from=planner /app/recipe.json recipe.json
RUN cargo chef cook --release --recipe-path recipe.json
COPY . .
RUN cargo build --release --workspace --bin comments-service

#--------------------------------------------------------------------

FROM debian:bookworm-slim AS runtime
WORKDIR /app
RUN apt-get update -y &&\
    apt-get install -y --no-install-recommends ca-certificates openssl &&\
    apt-get autoremove -y &&\
    apt-get clean -y &&\
    rm -rf /var/lib/apt/lists/*
COPY --from=builder /app/target/release/comments-service comments-service
COPY crates/comments-service/migrations migrations

EXPOSE 4000

ENTRYPOINT [ "./comments-service --port=4000" ]
