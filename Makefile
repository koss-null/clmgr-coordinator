all: build
	./build/clmgr-coordinator

build:
	/bin/bash -c "GOOS=linux go build -o ./build/clmgr-coordinator -i ./"

build-local:
	/bin/bash -c "go build -o ./build/clmgr-coordinator-l -i ./"

proto:
	./protobuf/compile-proto.sh

compose:
	docker-compose build
	docker-compose start

compose-start: compose
	docker-compose start

clean-compose:
	docker-compose rm --all

clean-proto:
	rm -rf ./protobuf/compiled/*

clean: clean-proto clean-compose
	rm -rf ./build/*
	mkdir ./build/docker
