#!/bin/bash

export GOPATH=$(pwd)

# -----------------------------------------------

echo installing fujin...
go install fujin

cp SERVERS www.esiyou.com.key www.esiyou.com.pem config.json ctl.sh ./bin

echo -e "\033[32mDone\033[0m"
