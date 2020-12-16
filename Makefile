VERSION=$(shell git log -1 --pretty=format:"0.0.0 (%h %ad)" --date=short)
export GO111MODULE=on

.PHONY: setup
setup:
	@echo "==> Checking dependencies..."
	which gotest > /dev/null || go get -u github.com/rakyll/gotest:
.PHONY: test
test:
	@echo "==> Running tests..."
	gotest -v -race -count=1 ./...

install:
	go install -ldflags "-X 'github.com/styvane/dcron/cmd.Version=${VERSION}'"
	@echo "Done!"

