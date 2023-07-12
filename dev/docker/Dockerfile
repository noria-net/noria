FROM golang:1.20.4-buster
SHELL ["/bin/bash", "-c"]

RUN apt-get update && apt-get install -y jq

# When testing noria only

# COPY . /app

# or, instead

# When testing noria with local modules (run from parent directory)
# Copy the noria repo in the docker container

COPY ./noria /app

# Copy the dev modules in the docker container

COPY ./alliance /alliance

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