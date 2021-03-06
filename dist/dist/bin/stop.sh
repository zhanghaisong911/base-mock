#!/bin/bash

#set -e

BIN_DIR=$(dirname $0)
BIN_DIR=$(cd $BIN_DIR; pwd)
ROOT_DIR=$(cd $BIN_DIR/..; pwd)
cd $ROOT_DIR

ps aux | grep 'bin/be-record-run.sh' | grep -v grep | awk '{ print $2 }' | xargs kill -9
ps aux | grep './be-record -log_dir=../log -config_file=../config/config.json' | grep -v grep | awk '{ print $2 }' | xargs kill -TERM
