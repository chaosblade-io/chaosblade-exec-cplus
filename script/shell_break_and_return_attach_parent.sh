#!/bin/bash

expect -c "
  set timeout 60
  spawn gdb -q attach $1
  expect {
    \"(gdb)\" {send \"set follow-fork-mode $2\n\";}
    timeout {puts \"Error: timeout during gdb attach\"; exit 1}
    eof {puts \"Error: gdb process ended unexpectedly\"; exit 1}
  }
  expect {
    \"(gdb)\" {send \"set pagination off\n\";}
    timeout {puts \"Error: timeout setting pagination\"; exit 1}
  }
  expect {
    \"(gdb)\" {send \"$3\n\";}
    timeout {puts \"Error: timeout sending command 3\"; exit 1}
  }
  expect {
    \"(gdb)\" {send \"$4\n\";}
    timeout {puts \"Error: timeout sending command 4\"; exit 1}
  }
  expect {
    \"(gdb)\" {send \"b $5\n\";}
    timeout {puts \"Error: timeout setting breakpoint\"; exit 1}
  }
  expect {
    \"Breakpoint\" {
      # Wait for the gdb prompt after breakpoint is set
      expect \"(gdb)\" {send \"commands\n\";}
    }
    \"No symbol table\" {
      send \"y\n\"
      expect \"(gdb)\" {send \"commands\n\";}
    }
    \"pending\" {
      send \"y\n\"
      expect \"(gdb)\" {send \"commands\n\";}
    }
    \"(gdb)\" {
      # Breakpoint was set immediately and we're at the prompt
      send \"commands\n\"
    }
    timeout {puts \"Error: timeout setting breakpoint commands\"; exit 1}
  }
  expect {
    \">\" {send \"silent\n\";}
    timeout {puts \"Error: timeout in commands mode\"; exit 1}
  }
  expect {
    \">\" {send \"return $6\n\";}
    timeout {puts \"Error: timeout setting return command\"; exit 1}
  }
  expect {
    \">\" {send \"cont\n\";}
    timeout {puts \"Error: timeout setting cont command\"; exit 1}
  }
  expect {
    \">\" {send \"end\n\";}
    timeout {puts \"Error: timeout ending commands\"; exit 1}
  }
  expect {
    \"(gdb)\" {send \"c\n\";}
    timeout {puts \"Error: timeout continuing execution\"; exit 1}
  }

  # Wait a brief moment for the continue command to take effect
  sleep 0.5
  
  # Exit cleanly - the breakpoint is already set and will trigger when hit
  puts \"Success: Breakpoint and return injection completed\"
  exit 0
"