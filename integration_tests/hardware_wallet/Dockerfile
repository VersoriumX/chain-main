FROM rust:latest as BUILDER
LABEL maintainer="chain@crypto.org"

RUN apt-get update && \
    apt-get -y install protobuf-compiler && \
    cd /tmp && \
    git clone https://github.com/crypto-com/ledger-rs.git && \
    cd ledger-rs && \
    git checkout -b crypto abba8c6cb31dc81b89e24a0132be101432b994b5 && \
    cd examples/zemu-grpc-server && \
    cargo build --release



