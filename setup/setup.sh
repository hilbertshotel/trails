#!/bin/bash

# upgrade & update
sudo apt -y update
sudo apt -y upgrade

# install go
wget https://dl.google.com/go/go1.19.3.linux-amd64.tar.gz
tar -C . -xzf go1.19.3.linux-amd64.tar.gz
rm -r go1.19.3.linux-amd64.tar.gz
echo "export PATH=$PATH:~/go/bin" >> $HOME/.profile
source $HOME/.profile

# build trails
cd $HOME/src/trails
go build

# setup logs
mkdir $HOME/src/trails/logs
touch $HOME/src/trails/logs/trails.log

# install postgres
sudo apt -y install postgresql

# enable trails service
sudo cp -r $HOME/src/trails/setup/trails.service /lib/systemd/system/trails.service
sudo systemctl enable trails.service
sudo systemctl start trails.service
