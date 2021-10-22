############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
ARG DOCKER_TAG=0.0.0
# checkout the project 
WORKDIR /builder
COPY . .
RUN ls
# Fetch dependencies.
# Using go get.
RUN go mod download
# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /cosmos-cashd -ldflags="-s -w -extldflags \"-static\" -X main.Version=$DOCKER_TAG" cmd/cosmos-cashd/main.go
############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /cosmos-cashd /

# P2P port
EXPOSE 26656/tcp
# RPC port
EXPOSE 26657/tcp
# Listen addres
EXPOSE 26658/tcp
# GRPC
EXPOSE 9090/tcp
# GRPC Web
EXPOSE 9091/tcp

# Run the node.
ENTRYPOINT [ "/cosmos-cashd" ]
CMD ["start", "--home", "/data"]
