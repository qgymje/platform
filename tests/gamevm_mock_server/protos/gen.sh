#! /bin/bash

protoc2 --go_out=. *.proto
protoc2 --cpp_out=. *.proto
