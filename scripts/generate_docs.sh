 #!/usr/bin/env bash

if ! command -v protoc &> /dev/null
then
    echo "Missing required binary: protoc"
    exit
fi

# NODE_URL is used to set the host in the swagger.yaml file
NODE_URL="archive-lcd.noria.nextnet.zone"

set -eo pipefail

current_dir=$(pwd)

mkdir -p ./tmp-swagger-gen
mkdir -p ./tmp-swagger-gen-proto

# move the vendor folder to a temp dir so that go list works properly
if [ -d vendor ]; then
  temp_dir=$(mktemp -d)
  mv ./vendor "$temp_dir"/vendor
fi

# move the vendor folder back to ./vendor
if [ -d "$temp_dir" ]; then
  mv "$temp_dir"/vendor ./vendor
  rm -rf "$temp_dir"
fi

# generate a full set of proto files with all dependencies
buf export buf.build/cosmos/cosmos-sdk:14f154b98b9b4cf381a0878e8a9d4694 --output ./tmp-swagger-gen-proto
buf export buf.build/noria-net/token-factory:42799303e69440b9b225bf1dc8f75418 --output ./tmp-swagger-gen-proto
buf export buf.build/noria-net/noria:931e6ba8bf9a4d50ae59a711c8265799 --output ./tmp-swagger-gen-proto

# copy the required buf files to the tmp dir
cp ./proto/buf.gen.swagger.yaml ./tmp-swagger-gen-proto/
head -n 3 ./proto/buf.yaml > ./tmp-swagger-gen-proto/buf.yaml
cd tmp-swagger-gen-proto
buf mod update
cd ..

# copy the tmp dir to /tmp to avoid issues already existing proto files
cp -r tmp-swagger-gen-proto /tmp


cd /tmp/tmp-swagger-gen-proto
proto_dirs=$(find cosmos noria osmosis -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  # generate swagger files (filter query files)
  query_file=$(find "${dir}" -maxdepth 1 \( -name 'query.proto' -o -name 'service.proto' \))
  if [[ ! -z "$query_file" ]]; then
    echo "Generating swagger for $query_file"
    buf generate --template buf.gen.swagger.yaml $query_file
  fi
done

cp -r /tmp/tmp-swagger-gen $current_dir/
cd $current_dir/docs

yarn install
yarn combine

#Add public servers to spec file for Noria testnet and mainnet
yq -i '."host"="'"$NODE_URL"'"' static/swagger.yaml
yq -i '."schemes"+=["https"]' static/swagger.yaml

cd ../

# clean swagger files
rm -rf ./tmp-swagger-gen*

# generate new statik docs for go
statik -src=./docs/static -include=*.html,*.yaml