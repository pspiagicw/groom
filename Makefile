LDFLAGS="-X main.VERSION=0.0.1"
PROJNAME=groom-make
build: # Compile a binary
	go build  -o ${PROJNAME} -ldflags ${LDFLAGS} cmd/${PROJNAME}/main.go
run: # Compile and then run the project
	go run cmd/${PROJNAME}/main.go
test: # Run tests
	go test -v ./...
format: # Format your project
	go fmt ./...
help: # Show help for each of the Makefile recipes.
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done

.PHONY: build test format help
