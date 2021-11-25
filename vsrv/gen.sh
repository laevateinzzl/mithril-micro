#!/bin/bash
#function gen {
#  SVCS=$1
#  PROTO_PATH=./${SVCS}/pb
#  GO_OUT_PATH=./${SVCS}/api/gen/v1
#  mkdir -p ${GO_OUT_PATH}
#
#  protoc -I=$PROTO_PATH --go_out=paths=source_relative:$GO_OUT_PATH ${SVCS}.proto
#  protoc -I=$PROTO_PATH --go-grpc_out=paths=source_relative:$GO_OUT_PATH ${SVCS}.proto
#  protoc -I=$PROTO_PATH --grpc-gateway_out=paths=source_relative,grpc_api_configuration=$PROTO_PATH/${SVCS}.yaml:$GO_OUT_PATH ${SVCS}.proto
#}
#gen video_srv

PROTO_PATH=./pb
GO_OUT_PATH=./api/gen

mkdir -p ${GO_OUT_PATH}
protoc -I=$PROTO_PATH --go_out=paths=source_relative:$GO_OUT_PATH video_srv.proto
protoc -I=$PROTO_PATH --go-grpc_out=paths=source_relative:$GO_OUT_PATH video_srv.proto
protoc -I=$PROTO_PATH --grpc-gateway_out=paths=source_relative,grpc_api_configuration=$PROTO_PATH/video_srv.yaml:$GO_OUT_PATH video_srv.proto