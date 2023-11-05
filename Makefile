DEST=./codegen/go
DEST2= ./codegen/swagger

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
    --openapiv2_out ${DEST2} \
    --openapiv2_opt use_go_templates=true \
	./proto/noolingo/user.proto 

	@echo "user - done"

#deck
	@protoc -I ./proto -I . -I ./proto/third_party --go_out ${DEST} \
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

# # flat:
# 	echo "start flat codegen folder"; \
# 	root=./codegen; \
# 	if [ -d "$root" ]; then \
# 		for dir in `ls $$root`; do \
# 			folder=`echo $(basename $${dir}) | cut -d/ -f2` ; \
# 			mv $$root/$$folder/proto/* $$root/$$folder/; \
# 			rm -r $$root/$$folder/proto; \
# 			echo "Done $$folder"; \
# 		done; \
# 	fi; \
#    	echo "finish flat codegen folder";

# protoc -I ./proto -I . -I ./proto/third_party --go_out ./gen/go/ --go_opt paths=source_relative     --go-grpc_out ./gen/go/ --go-grpc_opt paths=source_relative ./proto/noolingo/user.proto

##
