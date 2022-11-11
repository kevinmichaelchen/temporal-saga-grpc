.PHONY: all
all:
	$(MAKE) gen-proto

.PHONY: gen-proto
gen-proto:
	docker run --rm --volume "$(shell pwd):/workspace" --workdir /workspace bufbuild/buf generate

.PHONY: buf-lint
buf-lint:
	docker run --rm --volume "$(shell pwd):/workspace" --workdir /workspace bufbuild/buf lint
	docker run --rm --volume "$(shell pwd):/workspace" --workdir /workspace bufbuild/buf format -w
	# docker run --rm --volume "$(shell pwd):/workspace" --workdir /workspace bufbuild/buf breaking --against 'https://github.com/kevinmichaelchen/temporal-saga-grpc.git#branch=main'

.PHONY: buf-push
buf-push:
	@for i in $(shell fd buf.yaml | xargs dirname) ; do \
	  echo $$i ; \
	  pushd . ; \
	  cd $$i ; \
	  pwd ; \
	  buf push ; \
	  popd ; \
	  echo "" ; \
	done