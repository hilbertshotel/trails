#!/bin/bash

DIR="$HOME/src/trails"

# upgrade & update
sudo apt -y update && apt upgrade

# install git
sudo apt -y install git

# install go
GOTAR="go1.18.1.linux-amd64.tar.gz"
wget https://dl.google.com/go/$GOTAR
tar -C $HOME -xzf $GOTAR
rm -r $GOTAR

# clone & build trails
mkdir $HOME/src
cd $HOME/src
git clone https://github.com/hilbertshotel/trails.git
go build

# setup logs
mkdir $DIR/logs
touch $DIR/logs/trails.log

# install mongo
wget -qO - https://www.mongodb.org/static/pgp/server-5.0.asc | sudo apt-key add -
echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu focal/mongodb-org/5.0 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-5.0.list
sudo apt-get update
sudo apt-get install -y mongodb-org

# enable mongo service
sudo cp -r $DIR/setup/mongod.service /lib/systemd/system/mongod.service
sudo systemctl enable mongod.service
sudo systemctl start mongod.service

# enable trails service
sudo cp -r $DIR/setup/trails.service /lib/systemd/system/trails.service
sudo systemctl enable trails.service
sudo systemctl start trails.service

# handle nginx
apt -y install nginx
sudo cp -r $DIR/setup/trails /etc/nginx/sites-available/trails
sudo ln -s /etc/nginx/sites-available/trails /etc/nginx/sites-enabled/trails
sudo systemctl restart nginx.service