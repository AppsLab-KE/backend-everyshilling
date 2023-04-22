#!/usr/bin/env bash


#Generate otp proto to sdk/python-proto-gen
python -m grpc_tools.protoc --proto_path=./otp --python_out=./../../sdk/python-proto-gen/otp --pyi_out=./../../sdk/python-proto-gen/otp  --grpc_python_out=./../../sdk/python-proto-gen/otp otp/server.proto
python -m grpc_tools.protoc --proto_path=./otp --python_out=./../../sdk/python-proto-gen/otp --pyi_out=./../../sdk/python-proto-gen/otp --grpc_python_out=./../../sdk/python-proto-gen/otp otp/otp.proto

# Generate db proto
python -m grpc_tools.protoc --proto_path=./db --python_out=./../../sdk/python-proto-gen/db --pyi_out=./../../sdk/python-proto-gen/db --grpc_python_out=./../../sdk/python-proto-gen/db db/server.proto
python -m grpc_tools.protoc --proto_path=./db --python_out=./../../sdk/python-proto-gen/db --pyi_out=./../../sdk/python-proto-gen/db --grpc_python_out=./../../sdk/python-proto-gen/db db/user.proto