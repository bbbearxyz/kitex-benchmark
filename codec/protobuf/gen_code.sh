# grpc-gen
rm -rf ./grpc_gen && mkdir ./grpc_gen
protoc --go_out=./grpc_gen --go-grpc_out=./grpc_gen --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative ./echo-grpc.proto

# kitex-gen
rm -rf ./kitex_gen && mkdir ./kitex_gen
kitex -type protobuf -module github.com/bbbearxyz/kitex-benchmark ./echo-kitex.proto

## dubbo-gen
#rm -rf ./dubbo_gen && mkdir ./dubbo_gen
#protoc --go_out=./dubbo_gen --go-triple_out=./dubbo_gen --go_opt=paths=source_relative --go-triple_opt=paths=source_relative ./echo-dubbo.proto
#
## tars-gen
#rm -rf ./tars_gen && mkdir ./tars_gen
#protoc --go_out=plugins=tarsrpc:./tars_gen --go_opt=paths=source_relative ./echo-tars.proto