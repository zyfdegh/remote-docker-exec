#!/bin/sh

go build -o webconsole main.go

docker build -t webconsole .
