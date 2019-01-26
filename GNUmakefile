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

.PHONY: bootstrap
bootstrap: deps lint-deps # install all dependencies

.PHONY: deps
deps:  ## Install build and development dependencies
	@echo "==> Updating build dependencies..."
	go get -u github.com/kardianos/govendor
	go get -u gotest.tools/gotestsum
	command -v nomad || go get -u github.com/hashicorp/nomad

.PHONY: lint-deps
lint-deps: ## Install linter dependencies
	@echo "==> Updating linter dependencies..."
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

.PHONY: check
check: ## Lint the source code
	@echo "==> Linting source code..."
	@gometalinter \
		--deadline 10m \
		--vendor \
		--sort="path" \
		--aggregate \
		--enable-gc \
		--disable-all \
		--enable goimports \
		--enable misspell \
		--enable vet \
		--enable deadcode \
		--enable varcheck \
		--enable ineffassign \
		--enable structcheck \
		--enable unconvert \
		--enable gofmt \
		./...

.PHONY: vendorfmt
vendorfmt:
	@echo "--> Formatting vendor/vendor.json"
	test -x $(GOPATH)/bin/vendorfmt || go get -u github.com/magiconair/vendorfmt/cmd/vendorfmt
		vendorfmt
.PHONY: changelogfmt
changelogfmt:
	@echo "--> Making [GH-xxxx] references clickable..."
	@sed -E 's|([^\[])\[GH-([0-9]+)\]|\1[[GH-\2](https://github.com/hashicorp/nomad/issues/\2)]|g' CHANGELOG.md > changelog.tmp && mv changelog.tmp CHANGELOG.md

.PHONY: travis
travis: check test

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
