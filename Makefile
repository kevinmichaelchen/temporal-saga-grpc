# The DEFAULT_GOAL variable specifies the default target that will be built
# when you run the make command without any arguments.
.DEFAULT_GOAL := all

RED = \033[0;31m
YELLOW = \033[1;33m
BLUE = \033[0;34m
GREEN = \033[0;32m
NC = \033[0m

.PHONY: check_pkgx
check_pkgx:
	@echo "$(YELLOW)Checking if pkgx is installed...$(NC)"
	@if ! command -v pkgx > /dev/null; then \
		echo "$(RED)Error: pkgx is not installed. Please install it and try again.$(NC)"; \
		echo "$(RED)Consult the docs here: https://docs.pkgx.sh/$(NC)"; \
		echo "$(BLUE)The easiest way to install is with:$(NC)"; \
		echo "$(BLUE)curl -fsS https://pkgx.sh | sh$(NC)"; \
		exit 1; \
	fi
	@echo "pkgx is installed."
	pkgx --sync

.PHONY: all
all: check_pkgx stop
	pkgx temporal@latest server start-dev &
	pkgx eb -x 20 -- lsof -i :7233
	pkgx task@latest run:all

.PHONY: stop
stop:
	pkgx killport@latest 7233
	sleep 1
	pkgx docker compose stop
