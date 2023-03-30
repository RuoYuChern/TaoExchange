@ECHO OFF

SET context=Mprotos/taoContext.proto=./grpc/taoContext
SET exchange=Mprotos/taoExchange.proto=./grpc/taoExchange
SET protoc-pluging=protoc-gen-go="C:\Users\Administrator\go\bin\protoc-gen-go.exe" 
SET common=--go_out=. --go-grpc_out=. --proto_path=./protos
SET go_opt=--go_opt=%context% --go_opt=%exchange%  
SET file_lists=protos/taoContext.proto protos/taoExchange.proto 

echo "%common% %go_opt% --plugin=%protoc-pluging% ./protos/*.proto"
protoc %common% %go_opt% --plugin=%protoc-pluging% %file_lists%