IMPORT_PATH=github.com/hdtradeservices/api-proto/go
OUTPUT_DIR=go
GO_BUILD_DIR=${BUILD_DIR}/github.com/hdtradeservices/api-proto/go
OPEN_API_V2=openapiv2
DOCS_BUILD_DIR=docs

PROTO_SRC=src

all: build

build: tools clean
	mkdir -p $(OUTPUT_DIR) $(OPEN_API_V2) $(DOCS_BUILD_DIR)
	cd $(PROTO_SRC); \
		protoc **/*.proto \
		-I ../googleapis/ \
		-I ../src/ \
		--go_out=../$(OUTPUT_DIR) \
		--grpc-gateway_out=../$(OUTPUT_DIR) \
		--openapiv2_out ../openapiv2 \
		--openapiv2_opt logtostderr=true \
		--go_opt=paths=source_relative \
		--go-grpc_out=require_unimplemented_servers=false,paths=source_relative:../$(OUTPUT_DIR) \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--doc_out=:../$(DOCS_BUILD_DIR) \
		--doc_opt=markdown,docs.md;\
	rm -rf ../googleapis;

check-update:
	@git diff --exit-code
	@git status --porcelain
	@test -z "$(shell git status --porcelain)"

clean:
	find ${OUTPUT_DIR} -type f -name '*.pb.go' -delete
	rm -rf ${DOCS_BUILD_DIR}
	rm -rf ${OPEN_API_V2}

format:
	go fmt go/*/*_test.go

test: tools
	go test ./go/...

tools: 
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.6.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.6.0
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26.0
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
	go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest
	git clone --depth 1 https://github.com/googleapis/googleapis.git

lint:
	protolint src

lint-fix:
	protolint lint -fix src

.PHONY: all build check-update clean format test tools lint lint-fix 
