#!/bin/bash

if [ "$#" -ne 1 ]; then
    echo "Need to specify data to POST"
    ls key*
    exit 1
fi

IP=127.0.0.1
PORT=8080

# Issue a curl command to POST the key data
curl -sD - -X POST "http://$IP:$PORT/key" -H  "Content-Type: application/json" --data @key-${1}.json
