FROM golang:1.20.4-buster
SHELL ["/bin/bash", "-c"]

COPY ./build/noriad /usr/bin/noriad

RUN apt-get update && apt-get install -y jq

# rest server
EXPOSE 1317
# tendermint p2p
EXPOSE 26656
# tendermint rpc
EXPOSE 26657

ENTRYPOINT [ "/usr/bin/noriad" ]