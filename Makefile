.PHONY: all
all: \
	go-lint \
	go-test \
	go-mod-tidy \
	go-review \
	git-verify-submodules \
	git-verify-nodiff

.PHONY: clean
clean:
	rm -rf build

.PHONY: build
build:
	@git submodule update --init --recursive $@

include build/rules.mk
build/rules.mk: build
	@# included in submodule: build

# go-mod-tidy: update go modules
.PHONY: go-mod-tidy
go-mod-tidy:
	go mod tidy -v

# go-lint: lint Go files
.PHONY: go-lint
go-lint: $(GOLANGCI_LINT)
	# funlen: disabled
	# lll: disabled due to long go:generate directives
	# maligned: disabled to avoid re-arranging the packet struct
	$(GOLANGCI_LINT) run --enable-all --disable funlen,lll,maligned

# go-test: run Go test
.PHONY: go-test
go-test:
	go test -count=1 -race -cover ./...

.PHONY: go-review
go-review: $(GOBIN)
	$(GOBIN) -m -run github.com/einride/goreview/cmd/goreview -c 1 ./...

.PHONY: go-stringer
go-stringer: \
	pkg/dprotocol/digitalinput_string.go \
	pkg/dprotocol/eventid_string.go \
	pkg/dprotocol/fieldselector_string.go \
	pkg/dprotocol/gpsflag_string.go \
	pkg/dprotocol/output_string.go \
	pkg/dprotocol/radioaccesstechnology_string.go \
	pkg/dprotocol/stateflag_string.go

%_string.go: %.go
	go generate $<
