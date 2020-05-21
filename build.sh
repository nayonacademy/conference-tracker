#!/bin/sh
go build main.go
sudo service goweb restart
echo "go build done"