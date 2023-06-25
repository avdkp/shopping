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
	$(GOBUILD) -o $(BUILDDIR)/shopping-cart

# Clean target
clean:
	$(GOCLEAN)
	rm -rf $(BUILDDIR)

# Test target
test:
	$(GOTEST) ./...

# Dependency management target
deps:
	go get ./...

# Run the application
run:
	$(GOBUILD) -o $(BUILDDIR)/shopping-cart
	./$(BUILDDIR)/shopping-cart

# Default target
default: build
