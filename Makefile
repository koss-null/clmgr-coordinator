all:
	go build -I ./cmd/vkr_clmgr/

proto:
	./protobuf/compile-proto.sh

compose:
	docker-compose up

compose-start: compose
	docker-compose start

clean-proto:
	rm -rf ./protobuf/compiled/*

clean: clean-proto
	rm -rf ./build/*
	mkdir ./build/docker
