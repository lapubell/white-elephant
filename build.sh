#! /bin/bash

killall white-elephant
CGO_ENABLED=0 go build -o dist/white-elephant
dist/white-elephant &
