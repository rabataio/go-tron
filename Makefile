VERSION ?= 'dev'
GIT_COMMIT ?= $(shell git rev-parse HEAD || echo 'unknown')
CURRENT_TIME = $(shell LC_TIME=C.UTF-8 date -u)
TRON_PROTO_DIR = proto

lint:
	golangci-lint run

format:
	golangci-lint fmt

test:
	go test -race ./...

generate_client:
	protoc -I $(TRON_PROTO_DIR)/tron  \
	-I $(TRON_PROTO_DIR)/third_party \
	--go_out paths=source_relative:$(TRON_PROTO_DIR) \
	--go-grpc_out paths=source_relative:$(TRON_PROTO_DIR) \
	@$(TRON_PROTO_DIR)/proto_go_package_mapping.txt \
	$(TRON_PROTO_DIR)/tron/core/*.proto \
	$(TRON_PROTO_DIR)/tron/core/contract/*.proto \
	$(TRON_PROTO_DIR)/tron/api/*.proto
	mv $(TRON_PROTO_DIR)/core/contract/*.pb.go $(TRON_PROTO_DIR)/core/
	rm -rf $(TRON_PROTO_DIR)/core/contract
