GO=go
NAME := flaver
VERSION := 1.0.0
DIST := $(NAME)-$(VERSION)

all: test build

setup: update_version
	git submodule update --init

update_version:
	@for i in README.md ; do\
	    sed -e 's!Version-[0-9.]*-blue!Version-${VERSION}-blue!g' -e 's!tag/v[0-9.]*!tag/v${VERSION}!g' $$i > a ; mv a $$i; \
	done
	@sed 's/const VERSION = .*/const VERSION = "${VERSION}"/g' version.go > a; mv a version.go
	@sed 's/ARG version=".*"/ARG version="${VERSION}"/g' Dockerfile > a ; mv a Dockerfile
	@echo "Replace version to \"${VERSION}\""

test: setup
	$(GO) test -covermode=count -coverprofile=coverage.out $$(go list ./... | grep -v vendor)

# refer from https://pod.hatenablog.com/entry/2017/06/13/150342
define _createDist
	mkdir -p dist/$(1)_$(2)/$(DIST)
	GOOS=$1 GOARCH=$2 go build -o dist/$(1)_$(2)/$(DIST)/$(NAME)$(3) cmd/$(NAME)/main.go
	cp -r README.md LICENSE dist/$(1)_$(2)/$(DIST)
	tar cfz dist/$(DIST)_$(1)_$(2).tar.gz -C dist/$(1)_$(2) $(DIST)
endef

dist: build
	@$(call _createDist,darwin,amd64,)
	@$(call _createDist,darwin,arm64,)
	@$(call _createDist,windows,amd64,.exe)
	@$(call _createDist,windows,386,.exe)
	@$(call _createDist,linux,amd64,)
	@$(call _createDist,linux,386,)

build: setup
	$(GO) build -o $(NAME) -v cmd/main.go

clean:
	$(GO) clean
	rm -rf $(NAME)
