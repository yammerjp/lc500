#!/bin/bash

set -e

export vmid="$(grpcurl -plaintext  localhost:8080 v1.Worker.InitVM | jq -r '.vmid')"

echo "VMID: $vmid"

export script='
renderBody("hello, grpc for lc500!\n" + readAdditionalContext());
'

export context='
{
  "httpRequestMethod": "GET",
  "httpRequestUrl": "https://yammer.jp?q=query",
  "httpRequestBody": "",
  "httpRequestHeaders": {
    "X-Additional-Header": "lc500"
  },
  "additionalContext": "o-ma-ke"
}'

grpcurl -plaintext -d "$(echo '{}' | jq -n '{"vmid":env.vmid, "script": env.script}')" localhost:8080 v1.Worker.Compile
grpcurl -plaintext -d "$(echo "$context" | jq '.vmid=env.vmid')" localhost:8080 v1.Worker.SetContext
grpcurl -plaintext -d "{\"vmid\":\"$vmid\"}" localhost:8080 v1.Worker.Run
