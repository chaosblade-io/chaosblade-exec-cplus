#!/bin/bash

# Parameters: $1=PID, $2=forkMode, $3=unused, $4=unused, $5=breakLine, $6=delayDuration, $7=initParams

# Get the executable path for the process
EXEC_PATH=$(readlink -f /proc/$1/exe 2>/dev/null)
if [ -z "$EXEC_PATH" ] || [ ! -f "$EXEC_PATH" ]; then
    echo "Error: Cannot find executable for process $1"
    exit 1
fi

expect -c "
  spawn gdb -q \"$EXEC_PATH\" $1
  expect {
    \"gdb\" {send \"set follow-fork-mode $2\n\";}
  }
  expect {
    \"gdb\" {send \"set pagination off\n\";}
  }
  expect {
    \"gdb\" {send \"attach $1\n\";}
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
    \"*\" {
      if {\[string match \"*No symbol table*\" \$expect_out(buffer)\]} {
        send \"y\n\"
      } else {
        send \"commands\n\"
      }
    }
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
    \"gdb\" {send \"c\n\";}
  }

interact
"