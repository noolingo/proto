DEST=./codegen/go

generate:
	rm -rf codegen; 
	mkdir -p ${DEST}/noolingo 
	mkdir -p  ${DEST}/common
	@echo "start proto generation"; 

	protoc -I ./proto -I . -I ./proto/third_party --go_out ${DEST} \
	--go_opt paths=source_relative --go-grpc_out ${DEST} \
	--go-grpc_opt paths=source_relative \
	--grpc-gateway_out ${DEST} \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
	./proto/noolingo/user.proto 

	protoc -I ./proto -I . -I ./proto/third_party --go_out ${DEST} \
	--go_opt paths=source_relative --go-grpc_out ${DEST} \
	--go-grpc_opt paths=source_relative \
	--grpc-gateway_out ${DEST} \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
	./proto/common/common.proto

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
