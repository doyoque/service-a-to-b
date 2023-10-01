TIMESTAMP = `date +%Y%m%d%H%M%S`

build:
	@BUILD_START=$$(date +%s) \
	; go build -o ./bin ./main.go \
	&& echo "Build took $$(($$(date +%s)-BUILD_START)) seconds"

.PHONY: run
run: build
	@./bin/main
