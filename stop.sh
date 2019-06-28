#!/bin/sh
PID=`ps aux | grep /var/local/tools/ | grep bootstrap | awk '{print $2}'`

if [ -n "$PID" ]; then
    echo "Will shutdown tools: $PID"
    kill -9 $PID
    sleep 5
else echo "No Tools Process $PID"
fi