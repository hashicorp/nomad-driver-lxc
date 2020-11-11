PROJECT_ROOT := $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST)))))

default: build

.PHONY: clean
clean: ## Remove build artifacts
	rm -rf $(PROJECT_ROOT)/pkg

.PHONY: build
build:
	go install

.PHONY: test
test:
	go test \
		-timeout=15m \
	       ./...

.PHONY: fmt
fmt:
	@echo "==> Fixing source code with gofmt..."
	gofmt -s -w ./lxc

.PHONY: lint-deps
lint-deps: ## Install linter dependencies
	@echo "==> Updating linter dependencies..."
	GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.24.0

.PHONY: check
check: ## Lint the source code
	@echo "==> Linting source code..."
	@golangci-lint run -j 1 --timeout=10m

.PHONY: changelogfmt
changelogfmt:
	@echo "--> Making [GH-xxxx] references clickable..."
	@sed -E 's|([^\[])\[GH-([0-9]+)\]|\1[[GH-\2](https://github.com/hashicorp/nomad/issues/\2)]|g' CHANGELOG.md > changelog.tmp && mv changelog.tmp CHANGELOG.md

ALL_TARGETS += linux_amd64

# Define package targets for each of the build targets we actually have on this system
define makePackageTarget

pkg/$(1).zip: pkg/$(1)/nomad-driver-lxc
	@echo "==> Packaging for $(1)..."
	@zip -j pkg/$(1).zip pkg/$(1)/*

endef

# Reify the package targets
$(foreach t,$(ALL_TARGETS),$(eval $(call makePackageTarget,$(t))))

pkg/linux_amd64/nomad-driver-lxc:
	./scripts/build.sh

.PHONY: dev
dev: clean pkg/linux_amd64/nomad-driver-lxc

.PHONY: release
release: clean pkg/linux_amd64.zip
	@echo "==> Result:"
	@tree --dirsfirst $(PROJECT_ROOT)/pkg
