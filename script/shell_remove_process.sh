#!/bin/bash

pkill -f gdb

if [[ -n "$1" ]]; then
    kill -9 "$1"
fi