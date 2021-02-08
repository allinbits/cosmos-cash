FROM golang:alpine

RUN apk update
RUN apk add --no-cache ca-certificates curl make git libc-dev bash gcc linux-headers eudev-dev python3

WORKDIR /root

COPY . .

#RUN go get github.com/go-delve/delve/cmd/dlv
RUN go install ./cmd/cosmos-cashd

ENTRYPOINT ["./network/test-net/scripts/cash.sh"]

EXPOSE 26656 26657 1317 9090 40000

