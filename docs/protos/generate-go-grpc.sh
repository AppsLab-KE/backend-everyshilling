#!/usr/bin/env bash

#Generate db proto to sdk/go-proto-gen
protoc --go_out=./../../sdk/go-proto-gen --go_opt=paths=source_relative --go-grpc_out=./../../sdk/go-proto-gen --go-grpc_opt=paths=source_relative db/server.proto

#Generate otp proto to sdk/go-proto-gen
protoc --go_out=./../../sdk/go-proto-gen --go_opt=paths=source_relative --go-grpc_out=./../../sdk/go-proto-gen --go-grpc_opt=paths=source_relative otp/server.proto
