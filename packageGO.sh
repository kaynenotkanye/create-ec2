#!/bin/bash
### This is a generic script to package up go as a compiled binary
### This also zips up the binary as a deployment.zip file
### (which may then be used as a .zip for AWS Lambda, if desired)

rm ./bin/create-ec2

echo "Building binary"
GOOS=linux GOARCH=amd64 go build -o create-ec2 create-ec2.go

mv create-ec2 ./bin/

#echo "Create deployment package"
#zip deployment.zip create-ec2
