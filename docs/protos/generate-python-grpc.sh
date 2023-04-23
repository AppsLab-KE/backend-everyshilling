#!/usr/bin/env bash


#Generate otp proto to sdk/python-proto-gen/everyshillingsproto
python -m grpc_tools.protoc --proto_path=./otp --python_out=./../../sdk/python-proto-gen/everyshillingsproto/otp  --grpc_python_out=./../../sdk/python-proto-gen/everyshillingsproto/otp otp/otpserver.proto
python -m grpc_tools.protoc --proto_path=./otp --python_out=./../../sdk/python-proto-gen/everyshillingsproto/otp --grpc_python_out=./../../sdk/python-proto-gen/everyshillingsproto/otp otp/otp.proto

# Generate db proto
python -m grpc_tools.protoc --proto_path=./db --python_out=./../../sdk/python-proto-gen/everyshillingsproto/db --pyi_out=./../../sdk/python-proto-gen/everyshillingsproto/db --grpc_python_out=./../../sdk/python-proto-gen/everyshillingsproto/db db/dbserver.proto
python -m grpc_tools.protoc --proto_path=./db --python_out=./../../sdk/python-proto-gen/everyshillingsproto/db --pyi_out=./../../sdk/python-proto-gen/everyshillingsproto/db --grpc_python_out=./../../sdk/python-proto-gen/everyshillingsproto/db db/user.proto