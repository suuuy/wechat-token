#!/usr/bin/env bash

set -eo pipefail

echo "Begin build docker image."

./build.sh

docker build . -t wechat-token:latest
