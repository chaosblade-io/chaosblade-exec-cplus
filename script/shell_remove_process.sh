#!/bin/bash

# 杀掉所有 gdb 进程，即使没有找到也不报错
pkill -f gdb || true

if [[ -n "$1" ]]; then
    kill -9 "$1"
fi