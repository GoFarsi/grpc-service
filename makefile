greeting:
	protoc --proto_path=zarf/proto/greeting/v1 --proto_path=zarf/proto/include --proto_path=zarf/proto/include/googleapis --proto_path=zarf/proto/include/gateway --go-grpc_out=. --go_out=. --grpc-gateway_out=logtostderr=true:. zarf/proto/greeting/v1/*.proto
	# swagger
	protoc --proto_path=zarf/proto/greeting/v1 --proto_path=zarf/proto/include --proto_path=zarf/proto/include/googleapis --proto_path=zarf/proto/include/gateway  --openapiv2_out=allow_merge=true:zarf/gen/api --openapiv2_opt use_go_templates=true zarf/proto/greeting/v1/*.proto
	mv github.com/GoFarsi/grpc-service/zarf/gen/greeting zarf/gen
	rm -r github.com