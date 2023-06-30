#!/bin/sh
echo "generating ${1} protobuf codes"
if [ "${1}" = "" ]; then
  protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./api/*.proto
  rm -rf ./api/src/*.pb.go
  mv ./api/*.pb.go ./api/src/
  for f in ./api/src/*.pb.go; do
    protoc-go-inject-tag -input="$f" -XXX_skip=json,xml,yaml
  done
else
  protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt=paths=source_relative \
    ./"${1}"/api/*.proto
  rm -rf ./"${1}"/api/src/*.pb.go
  mv ./"${1}"/api/*.pb.go ./"${1}"/api/src/
  for f in ./"${1}"/api/src/*.pb.go; do
    protoc-go-inject-tag -input="$f" -XXX_skip=json,xml,yaml
  done
fi