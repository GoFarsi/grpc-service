greeting:
	protoc --proto_path=zarf/proto/greeting/v1 --go-grpc_out=. --go_out=. --grpc-gateway_out=logtostderr=true:. --validate_out=lang=go:.  zarf/proto/greeting/v1/*.proto
	# swagger
	protoc --proto_path=zarf/proto/greeting/v1 --openapiv2_out=allow_merge=true:zarf/gen/api --openapiv2_opt use_go_templates=true zarf/proto/greeting/v1/*.proto
	mv github.com/GoFarsi/grpc-service/zarf/gen/greeting zarf/gen
	rm -r github.com