all: zkit

clean:
	rm -rf dist

deps:
	dep ensure

zkit: clean deps

fmt:
	goimports -w $$(go list -f {{.Dir}} ./... | grep -v /vendor/)

.PHONY: deps fmt