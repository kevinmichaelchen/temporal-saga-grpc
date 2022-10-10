.PHONY: all
all:
	$(MAKE) gen-proto

.PHONY: gen-proto
gen-proto:
	docker run --rm --volume "$(shell pwd):/workspace" --workdir /workspace bufbuild/buf mod update idl
	docker run --rm --volume "$(shell pwd):/workspace" --workdir /workspace bufbuild/buf lint
	docker run --rm --volume "$(shell pwd):/workspace" --workdir /workspace bufbuild/buf format -w
	docker run --rm --volume "$(shell pwd):/workspace" --workdir /workspace bufbuild/buf generate
	# docker run --rm --volume "$(shell pwd):/workspace" --workdir /workspace bufbuild/buf breaking --against 'https://github.com/kevinmichaelchen/temporal-saga-grpc.git#branch=main'
