#!/usr/bin/env bash

set -eo pipefail

COSMOS_SDK_DIR=${COSMOS_SDK_DIR:-$(go list -f "{{ .Dir }}" -m github.com/cosmos/cosmos-sdk)}
GRAVITY_DEX=${GRAVITY_DEX:-$(go list -f "{{ .Dir }}" -m github.com/tendermint/liquidity)}

mkdir -p ./tmp-swagger-gen
proto_dirs=$(find ./proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do

  # generate swagger files (filter query files)
  query_file=$(find "${dir}" -maxdepth 1 \( -name 'query.proto' -o -name 'msg.proto' \))
  tx_file=$(find "${dir}" -maxdepth 1 \( -name 'tx.proto' -o -name 'msg.proto' \))
  if [[ ! -z "$query_file" ]]; then
    buf protoc  \
      -I "proto" \
      -I="$COSMOS_SDK_DIR/third_party/proto" \
      -I="$COSMOS_SDK_DIR/proto" \
      "$query_file" \
      --swagger_out=./tmp-swagger-gen \
      --swagger_opt=logtostderr=true --swagger_opt=fqn_for_swagger_name=true --swagger_opt=simple_operation_ids=true
    buf protoc  \
      -I "proto" \
      -I="$COSMOS_SDK_DIR/third_party/proto" \
      -I="$COSMOS_SDK_DIR/proto" \
      "$tx_file" \
      --swagger_out=./tmp-swagger-gen \
      --swagger_opt=logtostderr=true --swagger_opt=fqn_for_swagger_name=true --swagger_opt=simple_operation_ids=true
  fi
done

# combine swagger files
# uses nodejs package `swagger-combine`.
# all the individual swagger files need to be configured in `config.json` for merging
swagger-combine ./docs/Reference/swagger/config.json -o ./docs/Reference/swagger/swagger-ui/swagger.yaml -f yaml --continueOnConflictingPaths true --includeDefinitions true

# clean swagger files
rm -rf ./tmp-swagger-gen

