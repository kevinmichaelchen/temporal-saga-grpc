# The DEFAULT_GOAL variable specifies the default target that will be built
# when you run the make command without any arguments.
.DEFAULT_GOAL := help

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

.PHONY: view-docs
view-docs:
	buf mod open idl/temporalapis

.PHONY: buf-mod-update
buf-mod-update:
	@for i in $(shell fd buf.yaml | xargs dirname) ; do \
	  buf mod update $$i ; \
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