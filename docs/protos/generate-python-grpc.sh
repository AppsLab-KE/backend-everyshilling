#!/usr/bin/env bash

#Generate db proto to sdk/python-proto-gen
python -m grpc_tools.protoc --proto_path=./otp --python_out=./../../sdk/python-proto-gen --pyi_out=./../../sdk/python-proto-gen --grpc_python_out=./../../sdk/python-proto-gen otp/server.proto