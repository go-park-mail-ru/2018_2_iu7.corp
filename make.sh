#!/bin/bash

for d in ./*
do
  if [ -d "${d}" ]
  then
    cd ${d}
    service_name="${d:2}"
    if [ -f "./Makefile" ]
    then
      echo -e "${service_name}:\n"
      make build && make test
      echo -e "\n"
    fi
    cd ./..
  fi
done
