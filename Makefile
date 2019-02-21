# Revision 04.01.2019
BASEPATH = $(shell pwd)
export PATH := $(BASEPATH)/bin:$(PATH)

# Basic go commands
GOCMD      = go
GOBUILD    = $(GOCMD) build
GOINSTALL  = $(GOCMD) install
GORUN      = $(GOCMD) run
GOCLEAN    = $(GOCMD) clean
GOTEST     = $(GOCMD) test
GOGET      = $(GOCMD) get
GOFMT      = $(GOCMD) fmt
GOGENERATE = $(GOCMD) generate
GOTYPE     = $(GOCMD)type

# Basic dep commands
DEPCMD = dep

#TOOLS
MOCKERY = mockery
PGMGR = pgmgr

# Binary output name
BINARY = webshop
BUILD_DIR = $(BASEPATH)

# all src packages without vendor and generated code
PKGS = $(shell go list ./... | grep -v /vendor)

# Colors
GREEN_COLOR   = \033[0;32m
PURPLE_COLOR  = \033[0;35m
DEFAULT_COLOR = \033[m

# Database
PGMGR_URL = postgres://zloy:zloy@localhost:5432/db_shop
PGMGR_FOLDER = $(BASEPATH)/script/migration

.PHONY: all help clean deps rmdeps test lint fmt build version migration

all: clean fmt build test

help:
	@echo 'Usage: make <TARGETS> ... <OPTIONS>'
	@echo ''
	@echo 'Available targets are:'
	@echo ''
	@echo '    help               Show this help screen.'
	@echo '    clean              Remove binary.'
	@echo '    deps               Download and install build time dependencies.'
	@echo '    rmdeps             Clear all dependencies'
	@echo '    test               Run unit tests.'
	@echo '    lint               Run all linters including vet and gosec and others'
	@echo '    fmt                Run gofmt on package sources.'
	@echo '    build              Compile packages and dependencies.'
	@echo '    version            Print Go version.'
	@echo '    migration          Migration database.'
	@echo ''
	@echo 'Targets run by default are: clean fmt build test'
	@echo ''

clean:
	@echo " [$(GREEN_COLOR)clean$(DEFAULT_COLOR)]"
	@$(GOCLEAN)
	@if [ -f $(BINARY) ] ; then rm $(BINARY) ; fi
	@rm -rf $(COVERAGE_DIR)
	@rm -rf $(PROFILING_DIR)
	@rm -rf ./bin

rmdeps:
	@echo "[$(GREEN_COLOR)dep clean$(DEFAULT_COLOR)]"
ifneq ("$(wildcard ./Gopkg.lock)","")
	@echo "rm -rf ./Gopkg.lock"
	@rm -rf ./Gopkg.lock
endif
ifneq ("$(wildcard ./vendor)","")
	@echo "rm -rf ./vendor"
	@rm -rf ./vendor
endif

deps:
	@echo "[$(GREEN_COLOR)dep ensure$(DEFAULT_COLOR)]"
	@${DEPCMD} ensure -v

test:
	@echo " [$(GREEN_COLOR)test$(DEFAULT_COLOR)]"
	@$(GOTEST) -race $(PKGS)

lint:
	@echo " [$(GREEN_COLOR)lint$(DEFAULT_COLOR)]"
	@$(GORUN) ./vendor/github.com/golangci/golangci-lint/cmd/golangci-lint/main.go run \
	--no-config --enable=errcheck --enable=gosec --enable=gocyclo --enable=nakedret \
	--enable=prealloc --enable=gofmt --enable=goimports --enable=megacheck --enable=misspell ./...

fmt:
	@echo " [$(GREEN_COLOR)format$(DEFAULT_COLOR)]"
	@$(GOFMT) $(PKGS)


build: clean
	@echo " [$(GREEN_COLOR)build$(DEFAULT_COLOR)]"
	@$(GOBUILD) -o $(BINARY)

version:
	@echo " [$(GREEN_COLOR)version$(DEFAULT_COLOR)]"
	@$(GOCMD) version
