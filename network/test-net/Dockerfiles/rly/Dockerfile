FROM golang:alpine

RUN apk update
RUN apk add --no-cache ca-certificates curl make git libc-dev bash gcc linux-headers eudev-dev python3

WORKDIR /root

#COPY . .

#RUN go get github.com/go-delve/delve/cmd/dlv
RUN git clone https://github.com/relayer/relayer.git && cd relayer
RUN go install .
COPY ../../rly-config ./config

ENTRYPOINT ["./config/init.sh"]

