#!/bin/bash

set -ex

vmid=$(curl 'http://localhost:8080/vm/init')
script='renderBody(readBody() + readAdditionalContext())'
curl "http://localhost:8080/vm/compile?vmid=$vmid" -d "$script"
curl "http://localhost:8080/vm/setcontext?vmid=$vmid" -d '
{
        "httpRequest": {
                "method": "GET",
                "url": "https://yammer.jp?q=query",
                "headers": {
                        "Content-Type": "application/json"
                },
                "body": "{\"body\":\"bodyval\"}"
        },
        "additionalContext": "{\"foo\":\"bar\"}"
}'
curl "http://localhost:8080/vm/run?vmid=$vmid"
