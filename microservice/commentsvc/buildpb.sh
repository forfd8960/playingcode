#!/usr/bin/env bash

protoc  \
    --proto_path=.  \
    --go_out=plugins=grpc:$PWD \
    comment.proto
