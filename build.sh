#! /bin/bash

killall white-elephant
go build -o dist/white-elephant
dist/white-elephant &
