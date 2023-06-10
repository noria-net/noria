# syntax=docker/dockerfile:1.3

# Use the official Golang base image as a starting point
FROM golang:1.20

# Set the working directory
WORKDIR /app

# Install necessary dependencies
RUN apt-get update && apt-get install -y build-essential jq

# # Clone the Noria repository
RUN git clone https://github.com/noria-net/noria.git /app

# # Build and install the Noria node binary
RUN make build && make install

# Copy the entrypoint script
COPY ./entrypoint.sh /app/entrypoint.sh

# Install cosmovisor
RUN export GOPATH=/go
RUN export GOBIN=/go/bin
RUN go install cosmossdk.io/tools/cosmovisor/cmd/cosmovisor@latest

EXPOSE 26656 26657 1317 9090

# Set the entrypoint
ENTRYPOINT ["/bin/bash", "/app/entrypoint.sh"]
