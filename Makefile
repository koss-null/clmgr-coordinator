all:
	go build

proto:
	./protobuf/compile-proto.sh

test:
	docker-compose up

clean:
	rm -rf ./build/*
	rm -rf ./protobuf/compiled/*
	mkdir ./build/docker
