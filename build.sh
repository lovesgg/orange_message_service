#!/bin/bash

set -e
set -x

DIR=$(cd `dirname $0` && pwd -P)
cd $DIR

export GOPROXY=https://goproxy.cn
export GOFLAGS=-ldflags=-w
export env_orange_message_service=dev

run_release() {
  export env_orange_message_service=prod
  rm -rf ./release
  mkdir release
  cp -r ./bin ./release
  cp -r ./conf ./release
  go build -o ./release/orange_message_service .

  go build -o ./release/cmder ./cmd/
}

case $1 in
    "release" )
        run_release
        ;;
    "run" )
        go run .
        ;;
    "stop" )
        echo "stop"
        ;;
    "server" )
        ./output/server
        ;;
    "restart")
        supervisorctl -c /etc/supervisord.conf reload
        ;;
    "add-pre-commit")
        cp ./bin/pre-commit .git/hooks/pre-commit
        chmod +x .git/hooks/pre-commit
        ;;
    * )
        go build -o ./output/server .
        ;;
esac
