#!/bin/sh
rm -rf ./target
go build bootstrap.go
go build main.go
mkdir target
mkdir target/tools
mkdir target/tools/bin
mv bootstrap target/tools
cd target/tools
cp  -r ../../config ./
cp  -r ../../web ./
cp ../../stop.sh ./bin
cp ../../restart.sh ./bin
cp ../../Dockerfile .
cd -