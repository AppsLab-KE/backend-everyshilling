#!/usr/bin/env bash

#Generate db proto to sdk/go-proto-gen
protoc --proto_path=./db --go_out=./../../sdk/go-proto-gen/db --go_opt=paths=source_relative --go-grpc_out=./../../sdk/go-proto-gen/db --go-grpc_opt=paths=source_relative db/dbserver.proto
protoc --proto_path=./db --go_out=./../../sdk/go-proto-gen/db --go_opt=paths=source_relative --go-grpc_out=./../../sdk/go-proto-gen/db --go-grpc_opt=paths=source_relative db/user.proto
protoc --proto_path=./db --go_out=./../../sdk/go-proto-gen/db --go_opt=paths=source_relative --go-grpc_out=./../../sdk/go-proto-gen/db --go-grpc_opt=paths=source_relative db/rate.proto
protoc --proto_path=./db --go_out=./../../sdk/go-proto-gen/db --go_opt=paths=source_relative --go-grpc_out=./../../sdk/go-proto-gen/db --go-grpc_opt=paths=source_relative db/account.proto
protoc --proto_path=./db --go_out=./../../sdk/go-proto-gen/db --go_opt=paths=source_relative --go-grpc_out=./../../sdk/go-proto-gen/db --go-grpc_opt=paths=source_relative db/trade.proto
protoc --proto_path=./db --go_out=./../../sdk/go-proto-gen/db --go_opt=paths=source_relative --go-grpc_out=./../../sdk/go-proto-gen/db --go-grpc_opt=paths=source_relative db/transactions.proto

# Generate exchange proto to sdk/go-proto-gen
protoc --proto_path=./exchange --go_out=./../../sdk/go-proto-gen/exchange --go_opt=paths=source_relative --go-grpc_out=./../../sdk/go-proto-gen/exchange --go-grpc_opt=paths=source_relative exchange/exchangeserver.proto
protoc --proto_path=./exchange --go_out=./../../sdk/go-proto-gen/exchange --go_opt=paths=source_relative --go-grpc_out=./../../sdk/go-proto-gen/exchange --go-grpc_opt=paths=source_relative exchange/rate.proto
protoc --proto_path=./exchange --go_out=./../../sdk/go-proto-gen/exchange --go_opt=paths=source_relative --go-grpc_out=./../../sdk/go-proto-gen/exchange --go-grpc_opt=paths=source_relative exchange/account.proto
protoc --proto_path=./exchange --go_out=./../../sdk/go-proto-gen/exchange --go_opt=paths=source_relative --go-grpc_out=./../../sdk/go-proto-gen/exchange --go-grpc_opt=paths=source_relative exchange/trade.proto
protoc --proto_path=./exchange --go_out=./../../sdk/go-proto-gen/exchange --go_opt=paths=source_relative --go-grpc_out=./../../sdk/go-proto-gen/exchange --go-grpc_opt=paths=source_relative exchange/transactions.proto

#Generate otp proto to sdk/go-proto-gen
protoc  --proto_path=./otp --go_out=./../../sdk/go-proto-gen/otp --go_opt=paths=source_relative --go-grpc_out=./../../sdk/go-proto-gen/otp --go-grpc_opt=paths=source_relative otp/otpserver.proto
protoc  --proto_path=./otp --go_out=./../../sdk/go-proto-gen/otp --go_opt=paths=source_relative --go-grpc_out=./../../sdk/go-proto-gen/otp --go-grpc_opt=paths=source_relative otp/otp.proto