build/gsman: $(shell find . -name "*.go")
	go build -o build/gsman
