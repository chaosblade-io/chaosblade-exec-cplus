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


expect -c "
  spawn gdb -q attach $1
  expect {
    \"gdb\" {send \"set follow-fork-mode $2\n\";}
  }
  expect {
    \"gdb\" {send \"set pagination off\n\";}
  }
  expect {
    \"gdb\" {send \"$3\n\";}
  }
  expect {
    \"gdb\" {send \"$4\n\";}
  }
  expect {
    \"gdb\" {send \"b $5\n\";}
  }
  expect {
    \"gdb\" {send \"commands\n\";}
  }
  expect {
    \">\" {send \"silent\n\"}
  }
  expect {
    \">\" {send \"shell sleep $6\n\"}
  }
  expect {
    \">\" {send \"cont\n\"}
  }
  expect {
    \">\" {send \"end\n\"}
  }
  expect {
    \"gdb\" {send \"r $7\n\";}
  }
  expect {
    \"beginning\" {send \"y\n\";}
  }

 interact
"