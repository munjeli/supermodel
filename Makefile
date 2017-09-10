TEST?=$(shell go list ./... | grep -v vendor)
VET?=$(shell ls -d */ | grep -v vendor | grep -v website)
# Get the current full sha from git
GITSHA:=$(shell git rev-parse HEAD)
# Get the current local branch name from git (if we can, this may be blank)
GITBRANCH:=$(shell git symbolic-ref --short HEAD 2>/dev/null)
GOFMT_FILES?=$$(find . -not -path "./vendor/*" -name "*.go")
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

# Get the git commit
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
GIT_COMMIT=$(shell git rev-parse --short HEAD)
GIT_IMPORT=github.com/munjeli/treecli/version
GOLDFLAGS=-X $(GIT_IMPORT).GitCommit=$(GIT_COMMIT)$(GIT_DIRTY)

export GOLDFLAGS

default: deps generate test dev

ci: deps test

release: deps test releasebin package ## Build a release build

bin: deps ## Build debug/test build
	@go get github.com/mitchellh/gox
	@echo "WARN: 'make bin' is for debug / test builds only. Use 'make release' for release builds."
	@sh -c "$(CURDIR)/scripts/build.sh"

releasebin: deps
	@go get github.com/mitchellh/gox
	@grep 'const VersionPrerelease = "dev"' version.go > /dev/null ; if [ $$? -eq 0 ]; then \
		echo "ERROR: You must remove prerelease tags from version.go prior to release."; \
		exit 1; \
	fi
	@sh -c "$(CURDIR)/scripts/build.sh"

package:
	$(if $(VERSION),,@echo 'VERSION= needed to release; Use make package skip compilation'; exit 1)
	@sh -c "$(CURDIR)/scripts/dist.sh $(VERSION)"

deps:
	@go get golang.org/x/tools/cmd/stringer
	@go get github.com/kardianos/govendor
	@govendor sync

dev: deps ## Build and install a development build
	@grep 'const VersionPrerelease = ""' version.go > /dev/null ; if [ $$? -eq 0 ]; then \
		echo "ERROR: You must add prerelease tags to version/version.go prior to making a dev build."; \
		exit 1; \
	fi
	@mkdir -p pkg/$(GOOS)_$(GOARCH)
	@go install -ldflags '$(GOLDFLAGS)'
	@cp $(GOPATH)/bin/treecli bin
	@cp $(GOPATH)/bin/treecli pkg/$(GOOS)_$(GOARCH)

fmt: ## Format Go code
	@gofmt -w -s $(GOFMT_FILES)

fmt-check: ## Check go code formatting
	$(CURDIR)/scripts/gofmtcheck.sh $(GOFMT_FILES)

# generate runs `go generate` to build the dynamically generated
# source files.
generate: deps ## Generate dynamically generated code
	go generate .

test: deps fmt-check ## Run unit tests
	@go test $(TEST) $(TESTARGS) -timeout=2m
	@go tool vet $(VET)  ; if [ $$? -eq 1 ]; then \
		echo "ERROR: Vet found problems in the code."; \
		exit 1; \
	fi

updatedeps:
	@echo "INFO: treecli deps are managed by govendor. See README.md"

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
