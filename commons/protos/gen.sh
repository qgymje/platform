#! /bin/bash

rm -rf ./github.com

for d in ./*; do
    if [ -d "$d" ]; then
        protoc --go_out=plugins=grpc:. $d/*.proto
    fi
done
