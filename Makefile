DOCKER_BUF_FLAGS = --rm --volume "$(shell pwd):/workspace" --workdir /workspace
# Buf CLI versions:
# https://hub.docker.com/r/bufbuild/buf/tags
DOCKER_BUF = bufbuild/buf:1.9.0

.PHONY: all
all:
	$(MAKE) gen-proto

.PHONY: buf-mod-update
buf-mod-update:
	@for i in $(shell fd buf.yaml | xargs dirname) ; do \
	  docker run $(DOCKER_BUF_FLAGS) $(DOCKER_BUF) mod update $$i ; \
	done

.PHONY: buf-gen
buf-gen:
	docker run $(DOCKER_BUF_FLAGS) $(DOCKER_BUF) generate

.PHONY: buf-lint
buf-lint:
	docker run $(DOCKER_BUF_FLAGS) $(DOCKER_BUF) lint
	docker run $(DOCKER_BUF_FLAGS) $(DOCKER_BUF) format -w
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