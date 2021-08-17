GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

default: build test

test: fmtcheck
	go test -v . ./provider

build-only:
	@go install
	@mkdir -p ~/.terraform.d/plugins/
	@cp $(GOPATH)/bin/terraform-provider-qualys ~/.terraform.d/plugins/terraform-provider-qualys
	@echo "Build succeeded"

build: fmtcheck vet build-only

fmt:
	gofmt -w $(GOFMT_FILES)

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

.PHONY:test vet fmt fmtcheck
