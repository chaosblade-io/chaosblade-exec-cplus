#!/bin/bash

expect -c "
  spawn gdb
  expect {
    \"gdb\" {send \"file $1\n\";}
  }
  expect {
    \"gdb\" {send \"set follow-fork-mode $2\n\";}
  }
  expect {
    \"gdb\" {send \"$3\n\";}
  }
  expect {
    \"gdb\" {send \"set pagination off\n\";}
  }
  expect {
    \"gdb\" {send \"b $4\n\";}
  }
  expect {
    \"gdb\" {send \"commands\n\";}
  }
  expect {
    \">\" {send \"silent\n\"}
  }
  expect {
    \">\" {send \"r $5\n\"}
  }
  expect {
    \">\" {send \"cont\n\"}
  }
  expect {
    \">\" {send \"end\n\"}
  }
  expect {
    \"gdb\" {send \"r $6\n\";}
  }

 interact
"