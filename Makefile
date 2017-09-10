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
GIT_IMPORT=github.com/munjeli/go-scaffold/version
GOLDFLAGS=-X $(GIT_IMPORT).GitCommit=$(GIT_COMMIT)$(GIT_DIRTY)

export GOLDFLAGS

ci: deps test

deps:
	@go get golang.org/x/tools/cmd/stringer
	@go get github.com/kardianos/govendor
	@govendor sync

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
	@echo "INFO: go-scaffold deps are managed by govendor. See CONTRIBUTING.md"

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
