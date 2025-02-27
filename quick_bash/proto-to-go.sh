# /bin.bash

protoc user.proto --go-grpc_out=./

protoc user.proto --go_out=.