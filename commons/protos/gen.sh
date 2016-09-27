#! /bin/bash

for d in ./*; do
    if [ -d "$d" ] 
    then

        if [ $d == "./gamevm" ]
        then 
            protoc2 --go_out=. $d/*.proto
        else
        protoc --go_out=plugins=grpc:. $d/*.proto
        fi

    fi
done
