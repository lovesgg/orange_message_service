#!/bin/sh

set -e
set -x

ps aux | grep orange_message_service

FILEDIR=$(cd `dirname $0` && pwd -P)

cd $FILEDIR/../

pwd

supervisorctl restart orange_message_service
