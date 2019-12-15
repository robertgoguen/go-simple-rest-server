#!/bin/bash

if [ "$#" -ne 1 ]; then
    echo "Need to specify key to GET"
    exit 1
fi

IP=127.0.0.1
PORT=8080

# Issue a curl command to POST the key data
curl -H  "Content-Type: application/json" -X GET http://$IP:$PORT/key/${1}
