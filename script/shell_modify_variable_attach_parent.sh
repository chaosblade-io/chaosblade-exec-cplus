#!/bin/bash

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
    \">\" {send \"set $6 = $7\n\"}
  }
  expect {
    \">\" {send \"cont\n\"}
  }
  expect {
    \">\" {send \"end\n\"}
  }
  expect {
    \"gdb\" {send \"c\n\";}
  }

 interact
"