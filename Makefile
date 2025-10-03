.PHONY: prepare test bench

prepare:
	/bin/bash prepare.sh $(filter-out $@,$(MAKECMDGOALS))

test:
	go test ./...

bench:
	go test -bench=. ./...

%:
	@:
