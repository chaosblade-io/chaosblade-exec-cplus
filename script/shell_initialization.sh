#!/bin/bash
# Copyright 2025 The ChaosBlade Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


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