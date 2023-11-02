all: generate flat errors-gen

generate:
	echo "start proto generation"; \
	rm -rf codegen; \
   	buf generate; \
   	echo "finish proto generation";

flat:
	echo "start flat codegen folder"; \
	root=./codegen; \
	if [ -d "$root" ]; then \
		for dir in `ls $$root`; do \
			folder=`echo $(basename $${dir}) | cut -d/ -f2` ; \
			mv $$root/$$folder/proto/* $$root/$$folder/; \
			rm -r $$root/$$folder/proto; \
			echo "Done $$folder"; \
		done; \
	fi; \
   	echo "finish flat codegen folder";
