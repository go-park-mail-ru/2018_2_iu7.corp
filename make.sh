#!/bin/bash


# Install protobuf
curl -L https://github.com/google/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip -o /tmp/protoc.zip
unzip /tmp/protoc.zip -d /tmp/protoc
PATH="/tmp/protoc/bin:${PATH}"

go get -v github.com/micro/protoc-gen-micro
go get -v github.com/golang/protobuf/protoc-gen-go


# Install easyjson
go get -v github.com/mailru/easyjson/...


# Make services
for d in ./*
do
  if [ -d "${d}" ]
  then
    cd ${d}
    service_name="${d:2}"
    if [ -f "./Makefile" ]
    then
      echo -e "${service_name}:\n"
      make build && make test && make image
      echo -e "\n"
    fi
    cd ./..
  fi
done
