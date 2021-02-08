FROM golang:alpine

RUN apk update
RUN apk add --no-cache ca-certificates curl make git libc-dev bash gcc linux-headers eudev-dev python3

WORKDIR /root

#COPY . .
#COPY ./build/gaiad /usr/bin/gaiad

#RUN go get github.com/go-delve/delve/cmd/dlv
RUN git clone https://github.com/cosmos/gaia.git && cd gaia
RUN go install ./cmd/gaiad
COPY ../../scripts/gaia.sh ./init.sh

ENTRYPOINT ["./init.sh"]

EXPOSE 26656 26657 1317 9090 40000

