 #!/usr/bin/env bash

set -eo pipefail

mkdir -p ./tmp-swagger-gen

# move the vendor folder to a temp dir so that go list works properly
if [ -d vendor ]; then
  temp_dir=$(mktemp -d)
  mv ./vendor "$temp_dir"/vendor
fi

# Get the path of the cosmos-sdk repo from go/pkg/mod
cosmos_sdk_dir=$(go list -f '{{ .Dir }}' -m github.com/cosmos/cosmos-sdk)

# Get the path of the tokenfactory repo from go/pkg/mod
tokenfactory_dir=$(go list -f '{{ .Dir }}' -m github.com/CosmWasm/token-factory)

# move the vendor folder back to ./vendor
if [ -d $temp_dir ]; then
  mv "$temp_dir"/vendor ./vendor
  rm -rf "$temp_dir"
fi

proto_dirs=$(find ./proto "$cosmos_sdk_dir"/proto "$tokenfactory_dir"/proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  # generate swagger files (filter query files)
  query_file=$(find "${dir}" -maxdepth 1 \( -name 'query.proto' -o -name 'service.proto' \))
  if [[ ! -z "$query_file" ]]; then
    protoc \
      -I "proto" \
      -I "$cosmos_sdk_dir/third_party/proto" \
      -I "$cosmos_sdk_dir/proto" \
      -I "$tokenfactory_dir/proto" \
      "$query_file" \
      --swagger_out ./tmp-swagger-gen \
      --swagger_opt logtostderr=true \
      --swagger_opt fqn_for_swagger_name=true \
      --swagger_opt simple_operation_ids=true
  fi
done

cd ./docs
yarn install
yarn combine

#Add public servers to spec file for Noria testnet and mainnet
yq -i '."host"="archive-lcd.noria.nextnet.zone"' static/swagger.yaml
yq -i '."schemes"+=["https"]' static/swagger.yaml

cd ../

# clean swagger files
rm -rf ./tmp-swagger-gen

# generate new statik docs for go
statik -src=./docs/static -include=*.html,*.yaml