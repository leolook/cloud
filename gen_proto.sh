#!/usr/bin/env bash
if [ ! -d pb ]
then
    mkdir -p pb
fi
if [ ! -d pb ]
then
    ln -s ../../pb/ proto
fi
protoc --go_out=plugins=grpc:./pb ./proto/*.proto  ./proto/*.proto -I./proto/ -I./proto/
cd pb/;
ls *.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'
