#!/bin/bash

export GOPATH=$(pwd)

# -----------------------------------------------

echo installing fujin...
go install fujin

echo installing admin...
go install admin

cp SERVERS config.json filter.txt randname.txt ctl.sh ./bin

cp -r ./src/admin/conf ./src/admin/static ./src/admin/views ./bin 

echo -e "\033[32mDone\033[0m"
