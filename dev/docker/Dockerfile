FROM golang:1.20.4-buster
SHELL ["/bin/bash", "-c"]

RUN apt-get update && apt-get install -y jq

COPY . /app
WORKDIR /app

# Set go envs
RUN export GOPATH=/go
RUN export GOBIN=/go/bin

RUN make install

# rest server
EXPOSE 1317
# tendermint p2p
EXPOSE 26656
# tendermint rpc
EXPOSE 26657
# grpc
EXPOSE 9090

ENTRYPOINT [ "noriad" ]