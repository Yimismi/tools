#!/bin/sh
sh /var/local/tools/bin/stop.sh
nohup /var/local/tools/bootstrap > /dev/null 2>&1 &

PID=`ps aux | grep /var/local/tools/ | grep bootstrap | awk '{print $2}'`

if [ -n "$PID" ]; then
    echo "start server succeed. pid:$PID"
else echo "start server failed"
fi