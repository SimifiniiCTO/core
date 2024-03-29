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
	@echo "----- running core-auth-sdk tests"
	cd ./core-auth-sdk && go test && cd ..
	@echo "----- running core-pool tests"
	cd ./core-pool && go test && cd ..
	@echo "----- running core-utilities tests"
	cd ./core-utilities && go test && cd ..
	@echo "----- running core-grpc tests"

.PHONY: add-license
add-license: ## Find all .go files not in the vendor directory and try to write a license notice.
	find . -path ./vendor -prune -o -type f -name "*.go" -print | xargs ./etc/add_license.sh
	# Check for any changes made with -G. to ignore permissions changes. Exit with a non-zero
	# exit code if there is a diff.
	git diff -G. --quiet