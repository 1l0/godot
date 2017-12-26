all: build run

build: 
	go build -buildmode=c-shared -o go.dylib main.go
run:
	godot main.tscn
clean:
	rm go.dylib go.h