protoc --go_out=. --go-grpc_out=. proto/product/product.proto
rm -rf ../../shared_proto/product
mkdir ../../shared_proto/product
cp -r proto/product/* ../../shared_proto/product/