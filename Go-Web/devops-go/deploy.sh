#! /bin/sh

kill -9 $(pgrep webserver)
cd ~/newweb/
git pull http://xxx.git
cd webserver/
./webserver &