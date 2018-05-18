all: build
	/bin/bash -c "./build/clmgr-coordinator"

build:
	/bin/bash -c "GOOS=linux go build -o ./build/clmgr-coordinator -i ./"

build-local:
	/bin/bash -c "go build -o ./build/clmgr-coordinator-l -i ./"

proto:
	./protobuf/compile-proto.sh

clean-proto:
	rm -rf ./protobuf/compiled/*

clean: clean-proto
	rm -rf ./build/*
	mkdir ./build/docker
