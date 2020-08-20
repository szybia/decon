.PHONY: build

WHAT := rest

build:
	# Build for linux without CGO and strip binary
	@for target in $(WHAT); do \
		CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-s -w" -o bin/$$target ./cmd/$$target; \
	done