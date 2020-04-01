#!/bin/sh

set -e
set -x

ps aux | grep mj_lobster_go_service

FILEDIR=$(cd `dirname $0` && pwd -P)

cd $FILEDIR/../

pwd

supervisorctl restart mj_lobster_go_service
