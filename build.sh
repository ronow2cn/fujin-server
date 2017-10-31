#!/bin/bash

export GOPATH=$(pwd)

# -----------------------------------------------

echo installing fujin...
go install fujin

cp SERVERS config.json randname.txt ctl.sh ./bin

echo -e "\033[32mDone\033[0m"
