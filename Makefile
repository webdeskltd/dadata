PRODUCT=dadata.v2
REPNAME=gopkg.in/webdeskltd
DIR=$(strip $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST)))))

OLDGOPATH := $(GOPATH)
GOPATH := $(DIR):$(GOPATH)
DATE=$(shell date -u +%Y%m%d.%H%M%S.%Z)
GOGENERATE=$(shell if [ -f .gogenerate ]; then cat .gogenerate; fi)
TESTPACKETS=$(shell if [ -f .testpackages ]; then cat .testpackages; fi)
BENCHPACKETS=$(shell if [ -f .benchpackages ]; then cat .benchpackages; fi)

default: link test

## Creating and linking folders to meet golang requirements for locations
link:
	@mkdir -p $(DIR)/src/${REPNAME}; cd $(DIR)/src/${REPNAME} && ln -s ../../.. ${PRODUCT} 2>/dev/null; true
	@rm $(DIR)/src/vendor 2>/dev/null; ln -s $(DIR)/vendor $(DIR)/src/vendor 2>/dev/null; true
.PHONY: link

## Dependency manager
dep: link
	@if [ ! -f "${DIR}/go.mod" ]; then GO111MODULE="on" GOPATH="$(OLDGOPATH)" go mod init "${REPNAME}/${PRODUCT}"; fi
	@GO111MODULE="on" GOPATH="$(OLDGOPATH)" go mod download
	@GO111MODULE="on" GOPATH="$(OLDGOPATH)" go get
	@GO111MODULE="on" GOPATH="$(OLDGOPATH)" go mod vendor
.PHONY: dep

## Code generation (run only during development)
# All generating files are included in a .gogenerate file
gen:
	@for PKGNAME in $(GOGENERATE); do GOPATH="$(DIR)" go generate $${PKGNAME}; done
.PHONY: gen

## Testing one or multiple packages as well as applications with reporting on the percentage of test coverage
# All testing files are included in a .testpackages file
test: link
	@echo "mode: set" > $(DIR)/coverage.log
	@for PACKET in $(TESTPACKETS); do \
		touch coverage-tmp.log; \
		GOPATH=${GOPATH} go test -v -covermode=count -coverprofile=$(DIR)/coverage-tmp.log $$PACKET; \
		if [ "$$?" -ne "0" ]; then exit $$?; fi; \
		tail -n +2 $(DIR)/coverage-tmp.log | sort -r | awk '{if($$1 != last) {print $$0;last=$$1}}' >> $(DIR)/coverage.log; \
		rm -f $(DIR)/coverage-tmp.log; true; \
	done
.PHONY: test

## Displaying in the browser coverage of tested code, on the html report (run only during development)
cover: test
	GOPATH=${GOPATH} go tool cover -html=$(DIR)/coverage.log
.PHONY: cover

## Performance testing
# All testing files are included in a .benchpackages file
bench: link
	@ulimit -n 10000
	@for PACKET in $(BENCHPACKETS); do GOPATH=${GOPATH} go test -race -bench=. -benchmem -count 2 -parallel 10 -cpu 16 -cpuprofile $(DIR)/cpu.log -memprofile $(DIR)/mem.log $$PACKET; done
.PHONY: bench

## Code quality testing
# https://github.com/alecthomas/gometalinter/
# install: curl -L https://git.io/vp6lP | sh
lint:
	if command -v "gometalinter"; then gometalinter \
	--vendor \
	--deadline=15m \
	--cyclo-over=20 \
	--line-length=120 \
	--warn-unmatched-nolint \
	--disable=aligncheck \
	--enable=test \
	--enable=goimports \
	--enable=gosimple \
	--enable=misspell \
	--enable=unused \
	--enable=megacheck \
	--skip=src/vendor \
	--linter="vet:go tool vet -printfuncs=Infof,Debugf,Warningf,Errorf:PATH:LINE:MESSAGE" \
	src/...; fi
.PHONY: lint

## Clearing all temporary files and folders
clean:
	@rm -rf ${DIR}/bin/*; true
	@rm -rf ${DIR}/pkg/*; true
	@rm -rf ${DIR}/run/*; true
	@rm -rf ${DIR}/src; true
	@rm -rf ${DIR}/*.log; true
	@rm -rf ${DIR}/*.lock; true
.PHONY: clean
