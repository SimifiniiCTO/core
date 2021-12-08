.PHONY: proto
proto: contracts
	@echo "setting up service schema definition via protobuf"
	protoc -I. \
			-I$(GOPATH)/src \
			-I=$(GOPATH)/src/github.com/infobloxopen/protoc-gen-gorm \
			-I=$(GOPATH)/src/github.com/infobloxopen/atlas-app-toolkit \
			-I=$(GOPATH)/src/github.com/lyft/protoc-gen-validate/validate/validate.proto \
			-I=$(GOPATH)/src/github.com/infobloxopen/protoc-gen-gorm/options \
			-I=$(GOPATH)/src/github.com/protobuf/src/google/protobuf/timestamp.proto \
			--proto_path=${GOPATH}/src/github.com/gogo/protobuf/protobuf \
            --govalidators_out=./schema/models/ \
			--go_out=plugins=grpc:./schema/models/  --gorm_out="engine=postgres:./schema/models/" ./schema/proto/models/*.proto
.PHONY: contracts
contracts:
		@echo "setting up service level contracts with the other microservices via proto definitions"
		protoc -I. \
    			-I=$(GOPATH)/src \
    			--go_out=plugins=grpc:./schema/models/ ./schema/proto/contracts/*.proto

run-tests:
	@echo "running core-library tests"
	@echo "----- running core-metrics tests"
	cd ./core-metrics && go test && cd ..
	@echo "running core-library tests"
	@echo "----- running core-logging tests"
	cd ./core-logging && go test && cd ..
	@echo "----- running core-auth-sdk tests"
	cd ./core-auth-sdk && go test && cd ..
	@echo "----- running core-middleware tests"
	cd ./core-middleware/server && go test && cd ../..
	@echo "----- running core-pool tests"
	cd ./core-pool && go test && cd ..
	@echo "----- running core-tlsCert tests"
	cd ./core-tlsCert && go test && cd ..
	@echo "----- running core-tracing tests"
	cd ./core-tracing/datadog && go test && cd ../..
	cd ./core-tracing/jaeger && go test && cd ../..
	@echo "----- running core-utilities tests"
	cd ./core-utilities && go test && cd ..
	@echo "----- running core-grpc tests"

