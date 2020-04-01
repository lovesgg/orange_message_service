#!/bin/sh

set -e
set -x

ps aux | grep mj_lobster_go_service

supervisorctl stop mj_lobster_go_service
