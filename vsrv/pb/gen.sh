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

PROTO_PATH=./
GO_OUT_PATH=./

# protoc -I=$PROTO_PATH -I=$GOPATH/src -I=$GOPATH/src/mithril-micro/shared/third_party_pb/ --gogo_out=plugins=grpc,paths=source_relative,--grpc-gateway_out=allow_patch_feature=false,paths=source_relative:$GO_OUT_PATH video_srv.proto
truss video_srv.proto --svcout ..
protoc -I=$PROTO_PATH -I=$GOPATH/src -I=$GOPATH/src/mithril-micro/shared/third_party_pb/ --go_out=paths=source_relative:$GO_OUT_PATH video_srv.proto
protoc -I=$PROTO_PATH -I=$GOPATH/src -I=$GOPATH/src/mithril-micro/shared/third_party_pb/ --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:$GO_OUT_PATH video_srv.proto
protoc -I=$PROTO_PATH -I=$GOPATH/src -I=$GOPATH/src/mithril-micro/shared/third_party_pb/ --grpc-gateway_out=allow_patch_feature=true,paths=source_relative:$GO_OUT_PATH video_srv.proto
# protoc -I=$PROTO_PATH -I=$GOPATH/src -I=$GOPATH/src/mithril-micro/shared/third_party_pb/ --grpc-gateway_out=paths=source_relative:$GO_OUT_PATH video_srv.proto
