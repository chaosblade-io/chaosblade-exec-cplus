#!/bin/bash

# Parameters: $1=PID, $2=forkMode, $3=unused, $4=unused, $5=breakLine, $6=delayDuration, $7=initParams

# Get the executable path for the process
EXEC_PATH=$(readlink -f /proc/$1/exe 2>/dev/null)
if [ -z "$EXEC_PATH" ] || [ ! -f "$EXEC_PATH" ]; then
    echo "Error: Cannot find executable for process $1"
    exit 1
fi

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
    \"(gdb)\" {send \"b $5\n\";}
    timeout {puts \"Error: timeout setting breakpoint\"; exit 1}
  }
  expect {
    \"Breakpoint\" {
      expect {
        \"(gdb)\" {send \"commands\n\";}
        timeout {puts \"Error: timeout after breakpoint set\"; exit 1}
      }
    }
    \"No symbol table\" {
      send \"y\n\"
      expect {
        \"(gdb)\" {send \"commands\n\";}
        timeout {puts \"Error: timeout after answering y\"; exit 1}
      }
    }
    \"pending\" {
      send \"y\n\"
      expect {
        \"(gdb)\" {send \"commands\n\";}
        timeout {puts \"Error: timeout after pending breakpoint\"; exit 1}
      }
    }
    timeout {puts \"Error: timeout setting breakpoint commands\"; exit 1}
  }
  expect {
    \">\" {send \"silent\n\";}
    timeout {puts \"Error: timeout in commands mode\"; exit 1}
  }
  expect {
    \">\" {send \"shell sleep $6\n\";}
    timeout {puts \"Error: timeout setting sleep command\"; exit 1}
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

interact
"