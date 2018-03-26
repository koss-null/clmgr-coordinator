#!/bin/bash

# call from vkr_clmgr directory

protoc --go_out=./protobuffer/compiled ./protobuffer/*.proto