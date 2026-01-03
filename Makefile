# File: Makefile

# Run templ codegen
generate:
	templ generate

# Build Go app
build:
	go build -o bin/goth

# Run app
run:
	./bin/goth

# Regenerate templ files and run app
dev: generate build 
	./bin/goth

# Format .go files
fmt:
	go fmt ./...

# Clean build
clean:
	rm -f /bin/goth
