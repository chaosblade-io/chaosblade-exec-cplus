#!/bin/bash

# Install gdb
if ! [ -x "$(command -v gdb)" ]; then
  yes | yum install gdb
  echo 'Pass: gdb has been installed.'
else
  echo 'Pass: gdb has been installed.'
fi

# Install expect
if ! [ -x "$(command -v expect)" ]; then
  yes | yum install expect
  echo 'Pass: expect has been installed.'
  exit 1
else
  echo 'Pass: expect already has been installed.'
  exit 1
fi