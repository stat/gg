GO       := go
GO_BUILD := $(GO) build
GO_CLEAN := $(GO) clean
GO_TEST  := $(GO) clean -testcache; $(GO) test

SRC      := $(shell find . -path "./\.*" -prune -type f -name '*.go')

TARGET   := main
BUILDS   := ./build
TESTS    := ./...

BUILD    := `git rev-parse HEAD`
MAJOR    := $(shell cat .version | cut -d. -f1)
MINOR    := $(shell cat .version | cut -d. -f2)
PATCH    := $(shell cat .version | cut -d. -f3)

GOFLAGS  := -v
GOTAGS   :=
LDFLAGS  := -X=$(TARGET)/pkg/version.Build=$(BUILD) \
            -X=$(TARGET)/pkg/version.Major=$(MAJOR) \
            -X=$(TARGET)/pkg/version.Minor=$(MINOR) \
            -X=$(TARGET)/pkg/version.Patch=$(PATCH) \
            -X=$(TARGET)/pkg/version.String=${MAJOR}.${MINOR}.${PATCH}

.PHONY: test
test: ## run tests
	@$(GO_TEST) -ldflags="$(LDFLAGS)" -tags="${GOTAGS}" -v $(TESTS)
