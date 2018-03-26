#!/bin/bash

# call from vkr_clmgr directory

protoc --go_out=./protobuf/compiled `find ./protobuf | grep .*.proto$`