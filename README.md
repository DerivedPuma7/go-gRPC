# go-gRPC

## plugins necessários:

### compilador do protocol buffer:
https://grpc.io/docs/protoc-installation/  
apt install -y protobuf-compiler

## para entender os contratos protoc e transformar no arquivo go:
protoc --go_out=. --go-grpc_out=. proto/course_category.proto