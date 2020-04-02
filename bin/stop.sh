#!/bin/sh

set -e
set -x

ps aux | grep orange_message_service

supervisorctl stop orange_message_service
