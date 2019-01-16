if [ -z "${PROTOBUF_HOME}" ]; then
    GOOGLE_PROTOBUF_INCLUDE="$(dirname $(dirname $(which protoc)))/include"
else
    GOOGLE_PROTOBUF_INCLUDE="${PROTOBUF_HOME}/include"
fi
VERSION=$1
ENTITY=$2
protoc -I${GOOGLE_PROTOBUF_INCLUDE} \
  -I./pb/${VERSION}/${ENTITY} \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:./pb/${VERSION}/${ENTITY} \
  --grpc-gateway_out=logtostderr=true:./pb/${VERSION}/${ENTITY} \
  --swagger_out=logtostderr=true:./pb/${VERSION}/${ENTITY} \
  ./pb/${VERSION}/${ENTITY}/*.proto
