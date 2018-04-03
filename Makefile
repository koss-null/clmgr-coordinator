all:
	go build

proto:
	./protobuf/compile-proto.sh

test:
	docker-compose up

clean-proto:
	rm -rf ./protobuf/compiled/*

clean: clean-proto
	rm -rf ./build/*
	mkdir ./build/docker
