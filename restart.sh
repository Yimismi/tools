#!/bin/sh
sh /var/local/tools/stop.sh
cd /var/local/tools
/var/local/tools/bootstrap > /dev/null &
cd -