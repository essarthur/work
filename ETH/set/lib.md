#!/bin/bash

clear
echo "Start process"
echo

# Settings 
# export GOPATH=$HOME/mainproject
export GOPATH=$PWD
export GOROOT=$HOME/go
export PATH=$PATH:$GOROOT/bin

echo Start....
# Base library for load
# Warning: Very long time is loading !

# go get -u -v github.com/ovcharovvladimir/Prysm
# go get -u -v github.com/ovcharovvladimir/essentiaHybrid
# go get -u -v github.com/fjl/memsize/memsizeui
# go get -u -v github.com/gogo/protobuf/proto
# go get -u -v github.com/golang/protobuf/ptypes
# go get -u -v github.com/golang/protobuf/ptypes/empty
# go get -u -v github.com/golang/protobuf/ptypes/timestamp
# go get -u -v github.com/libp2p/go-floodsub
# go get -u -v github.com/libp2p/go-libp2p
# go get -u -v github.com/libp2p/go-libp2p-crypto
# go get -u -v github.com/libp2p/go-libp2p-host
# go get -u -v github.com/libp2p/go-libp2p-peer
# go get -u -v github.com/libp2p/go-libp2p-peerstore
# go get -u -v github.com/libp2p/go-libp2p/p2p/discovery
# go get -u -v github.com/multiformats/go-multiaddr
# go get -u -v github.com/sirupsen/logrus
# go get -u -v github.com/urfave/cli
# go get -u -v github.com/x-cray/logrus-prefixed-formatter
# go get -u -v golang.org/x/crypto/blake2b
# go get -u -v golang.org/x/net/context
# go get -u -v google.golang.org/grpc
# go get -u -v google.golang.org/grpc/credentials
# go get -u -v google.golang.org/genproto/googleapis/rpc/status
echo Finish ....    

start Buildin Process...
go build -o node
./node

