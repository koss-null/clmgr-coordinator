all:
	go build -o ./build/clmgr-coordinator -i ./

proto:
	./protobuf/compile-proto.sh

compose:
	docker-compose up

compose-start: compose
	docker-compose start

clean-compose:
	docker-compose rm --all

clean-proto:
	rm -rf ./protobuf/compiled/*

clean: clean-proto clean-compose
	rm -rf ./build/*
	mkdir ./build/docker
