gen-deps: ## Install dependencies for protoc auto gen
	sudo export GOPATH=$(go env GOPATH)
	sudp export PATH="$PATH:$(go env GOPATH)/bin"
	@echo "downloading protoc-gen tool"
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go get github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
	go install github.com/infobloxopen/protoc-gen-gorm@latest
	go get github.com/srikrsna/protoc-gen-gotag
	go install github.com/mitchellh/protoc-gen-go-json@latest
	brew install protobuf protoc-gen-go protoc-gen-go-grpc

gen-models: ## Autogen service stubs
	@echo "setting up model definition via protobuf"
	protoc -I. \
			-I=$(GOPATH)/src \
			-I=$(GOPATH)/src/github.com/infobloxopen/protoc-gen-gorm \
			-I=$(GOPATH)/src/github.com/infobloxopen/atlas-app-toolkit \
			-I=$(GOPATH)/src/github.com/lyft/protoc-gen-validate/validate/validate.proto \
			-I=$(GOPATH)/src/github.com/infobloxopen/protoc-gen-gorm/options \
			-I=$(GOPATH)/src/github.com/protobuf/src/google/protobuf/timestamp.proto \
			algolia_queue.proto email_queue.proto content_enrichment_queue.proto --go_out=:$(GOPATH)/src --go_opt=paths=import \
			--go-grpc_out=:$(GOPATH)/src --go-grpc_opt=paths=import

autogen: gen-deps gen-models