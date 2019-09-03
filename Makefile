export GO111MODULE = on

.PHONY: all
all: \
	circleci-config-validate \
	go-lint \
	go-test \
	go-mod-tidy \
	git-verify-submodules \
	git-verify-nodiff \
	go-review

.PHONY: clean
clean:
	rm -rf build
	rm -rf output

.PHONY: build
build:
	@git submodule update --init --recursive $@

# go-mod-tidy: update go modules
.PHONY: go-mod-tidy
go-mod-tidy:
	go mod tidy -v

include build/rules.mk
build/rules.mk: build

# go-lint: lint Go files
.PHONY: go-lint
go-lint: $(GOLANGCI_LINT)
	$(GOLANGCI_LINT) run --enable-all

# go-test: run Go test 
.PHONY: go-test
go-test:
	go test -count=1 -race -cover ./...

.PHONY: go-review
go-review: $(GOREVIEW)
	$(GOREVIEW) -c 1 ./...

# circleci-config-validate: validate CircleCI config
.PHONY: circleci-config-validate
circleci-config-validate: $(CIRCLECI)
	$(CIRCLECI) config validate

# go-run-server-local:
.PHONY: go-run-server-local
go-run-server-local:
	export GOOGLE_CLOUD_PROJECT=einride-portal && \
	go run ./cmd/aplicom-go/main.go --port 5144

#  send-local-test-data:
.PHONY: send-local-test-data
send-local-test-data:
	./cmd/scripts/send-test-data.sh

.PHONY: deploy
deploy:
	cd ./cmd/aplicom-go && GOOS=linux GOARCH=amd64 go build -o ../../output/aplicom-go
	gcloud compute scp --project "einride-portal" --zone "europe-north1-a"  ./output/aplicom-go root@aplicom-server:~
	gcloud compute ssh --project "einride-portal" --zone "europe-north1-a"  root@aplicom-server \
	--command 'systemctl stop aplicom-go && sleep 4 && cp aplicom-go /usr/local/bin && systemctl start aplicom-go'

