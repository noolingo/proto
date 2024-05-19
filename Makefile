DEST=./codegen/go
DEST2= ./codegen/swagger
ERRORS_GENERATED := ./codegen/go/apierrors/apierrors_gen.go
ERRORS_SOURCE    := ./errors/errors.json

all: generate errors-gen

generate:
	@rm -rf codegen; 
	@mkdir -p ${DEST}/noolingo 
	@mkdir -p  ${DEST}/common
	@mkdir -p ${DEST2}/noolingo 
	@mkdir -p  ${DEST2}/common
	@echo "start proto generation"; 

#common
	@protoc -I ./proto -I . -I ./proto/third_party --go_out ${DEST} \
	--go_opt paths=source_relative --go-grpc_out ${DEST} \
	--go-grpc_opt paths=source_relative \
	--grpc-gateway_out ${DEST} \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
	--openapiv2_out ${DEST2} \
    --openapiv2_opt use_go_templates=true \
	./proto/common/common.proto
	@echo "common - done"

#user
	@protoc -I ./proto -I . -I ./proto/third_party --go_out ${DEST} \
	--go_opt paths=source_relative --go-grpc_out ${DEST} \
	--go-grpc_opt paths=source_relative \
	--grpc-gateway_out ${DEST} \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
	--experimental_allow_proto3_optional=true \
    --openapiv2_out ${DEST2} \
    --openapiv2_opt use_go_templates=true \
	./proto/noolingo/user.proto 
	@echo "user - done"

#deck
	@protoc -I  ./proto -I . -I ./proto/third_party --go_out ${DEST} \
	--go_opt paths=source_relative --go-grpc_out ${DEST} \
	--go-grpc_opt paths=source_relative \
	--grpc-gateway_out ${DEST} \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
	--experimental_allow_proto3_optional=true \
	--openapiv2_out ${DEST2} \
    --openapiv2_opt use_go_templates=true \
	./proto/noolingo/deck.proto 
	@echo "deck - done"

#statistic
	@protoc -I ./proto -I . -I ./proto/third_party --go_out ${DEST} \
	--go_opt paths=source_relative --go-grpc_out ${DEST} \
	--go-grpc_opt paths=source_relative \
	--grpc-gateway_out ${DEST} \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
	--experimental_allow_proto3_optional=true \
	--openapiv2_out ${DEST2} \
    --openapiv2_opt use_go_templates=true \
	./proto/noolingo/statistic.proto 
	@echo "statistic - done"

#cards
	@protoc -I ./proto -I . -I ./proto/third_party --go_out ${DEST} \
	--go_opt paths=source_relative --go-grpc_out ${DEST} \
	--go-grpc_opt paths=source_relative \
	--grpc-gateway_out ${DEST} \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
	--experimental_allow_proto3_optional=true \
	--openapiv2_out ${DEST2} \
    --openapiv2_opt use_go_templates=true \
	./proto/noolingo/cards.proto 
	@echo "cards - done"

#info
	@protoc -I ./proto -I . -I ./proto/third_party --go_out ${DEST} \
	--go_opt paths=source_relative --go-grpc_out ${DEST} \
	--go-grpc_opt paths=source_relative \
	--grpc-gateway_out ${DEST} \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
	--openapiv2_out ${DEST2} \
    --openapiv2_opt use_go_templates=true \
	./proto/common/info.proto
	@echo "info - done"

	@echo "finish proto generation";


gen:
	protoc --proto_path=./proto \
	-I ./proto -I . -I ./proto/third_party\
    --descriptor_set_out=user.protoset \
    --include_imports ./proto/noolingo/user.proto

	protoc --proto_path=./proto \
	-I ./proto -I . -I ./proto/third_party\
    --descriptor_set_out=cards.protoset \
    --include_imports ./proto/noolingo/cards.proto 

	protoc --proto_path=./proto \
	-I ./proto -I . -I ./proto/third_party\
    --descriptor_set_out=deck.protoset \
    --include_imports ./proto/noolingo/deck.proto 

errors-gen:
	go run cmd/err-gen/gen.go
	gofmt -w $(ERRORS_GENERATED)

# refact:
# 	go mod edit -module github.com/noolingo/proto
# 	-- rename all imported module
# 	find . -type f -name '*.go' \
#   	-exec sed -i -e 's,github.com/MelnikovNA/noolingoproto,github.com/noolingo/proto,g' {} \;