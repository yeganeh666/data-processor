#!/bin/sh
echo "running ${1} ${2}"
cd ./"${1}" && go run ./cmd "${2}"