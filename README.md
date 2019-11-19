# Grpc.io


Commands: 

`wget https://github.com/protocolbuffers/protobuf/releases/download/v3.10.1/protoc-3.10.1-linux-x86_64.zip`

`unzip -j protoc-3.10.1-linux-x86_64.zip  bin/protoc -d bin/`

`export GOBIN=`pwd`/bin`

`go install github.com/golang/protobuf/...`

`protoc --go_out=plugins=grpc,paths=source_relative:. ./api/api.proto`
