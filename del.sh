
# First get the running pod name and the ip address it is running at

if [ "$#" -ne 1 ]; then
    echo "Need to specify key to DELETE"
    exit 1
fi

IP=127.0.0.1
PORT=8080

# Issue a curl command to POST the binding data
curl -H "Content-Type: application/json" -X DELETE http://$IP:$PORT/keys/${1}
