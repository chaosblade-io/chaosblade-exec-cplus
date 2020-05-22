#!/bin/bash

ps -ef|grep $1|grep -v 'grep'|grep -v 'initParams'|grep -v 'shell_check_process_id'| awk '{print $2}'