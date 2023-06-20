DOCKER_BUF_FLAGS = --rm --volume "$(shell pwd):/workspace" --workdir /workspace
# Buf CLI versions:
# https://hub.docker.com/r/bufbuild/buf/tags
DOCKER_BUF = bufbuild/buf:1.21.0

.PHONY: all
all:
	$(MAKE) buf-gen

.PHONY: help
help : Makefile
	@sed -n 's/^##//p' $<

## gen           : Generates code
.PHONY: gen
gen:
	task gen

## lint          : Lints code
.PHONY: lint
lint:
	task lint

## format        : Formats code
.PHONY: format
format:
	task format

.PHONY: view-docs
view-docs:
	buf mod open idl/temporalapis

.PHONY: buf-mod-update
buf-mod-update:
	@for i in $(shell fd buf.yaml | xargs dirname) ; do \
	  docker run $(DOCKER_BUF_FLAGS) $(DOCKER_BUF) mod update $$i ; \
	done

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