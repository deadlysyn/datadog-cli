ARTIFACT := dd

.PHONY: all build clean

all: build

build: clean
	@go mod tidy
	@go build -a -ldflags '-extldflags "-static"' -o $(ARTIFACT) .
	@strip $(ARTIFACT)

clean:
	@rm -f $(ARTIFACT)
