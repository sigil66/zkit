all: zkit

clean:
	rm -rf dist
	go mod tidy

zkit: clean

fmt:
	goimports -w $$(go list -f {{.Dir}} ./... | grep -v /vendor/)

.PHONY: deps fmt
