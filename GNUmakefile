default: build

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

.PHONY: deps
deps:  ## Install build and development dependencies
	@echo "==> Updating build dependencies..."
	go get -u github.com/kardianos/govendor
	go get -u gotest.tools/gotestsum

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
