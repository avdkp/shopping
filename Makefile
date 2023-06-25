# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get

# Build directory
BUILDDIR = build

# Build target
build:
	mkdir -p $(BUILDDIR)
	$(GOBUILD) -o $(BUILDDIR)/myapp

# Clean target
clean:
	$(GOCLEAN)
	rm -rf $(BUILDDIR)

# Test target
test:
	$(GOTEST) ./...

# Dependency management target
deps:
	$(GOGET) github.com/some/package

# Run the application
run:
	$(GOBUILD) -o $(BUILDDIR)/myapp
	./$(BUILDDIR)/myapp

# Default target
default: build