#!/bin/sh -e
cd $HOME
wget https://dl.google.com/go/go1.13.5.linux-amd64.tar.gz
sudo tar -xvf go1.13.5.linux-amd64.tar.gz
sudo mv go /usr/local
sudo mkdir go
cd go
sudo mkdir pkg
sudo mkdir bin
sudo mkdir src
cd $HOME
sudo echo "export GOROOT=/usr/local/go" >> .bashrc
sudo echo "export GOPATH=$HOME/go" >> .bashrc
sudo echo "export PATH=$GOPATH/bin:$GOROOT/bin:$PATH" >> .bashrc
source ~/.profile
go version
go env
