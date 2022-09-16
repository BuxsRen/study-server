protoc --go_out=./generate --go_opt=paths=source_relative --go-grpc_out=./generate --go-grpc_opt=paths=source_relative *.proto

:: --go_out=. 生成rpc
:: --go-grpc_out=. 生成grpc
:: --go_opt=paths=source_relative 使用相对路径

:: 将生成的protobuf复制到 /bootstrap/grpc/proto/ 目录下的对应目录中