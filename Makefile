GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)

default: build test testacc

test: fmtcheck
	go test -v . ./provider

testacc: fmtcheck
	go test -v ./provider -run="TestAcc"

build-only:
	@go install
	@mkdir -p ~/.terraform.d/plugins/
	@cp $(GOPATH)/bin/terraform-provider-qualys ~/.terraform.d/plugins/terraform-provider-qualys
	@echo "Build succeeded"

build: fmtcheck vet testacc build-only

clean:
	rm -rf pkg/

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

download-cloudview-swagger: swagger/swagger.json
	@mkdir -p swagger
	@rm swagger/swagger.json
	@wget -q https://qualysguard.qg2.apps.qualys.com/cloudview-api/v2/api-docs -O swagger/swagger.json

build-swagger:
	@scripts/swagger.sh

.PHONY: build build-only test testacc vet fmt fmtcheck errcheck vendor-status test-compile