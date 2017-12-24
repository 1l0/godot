all:
	c-for-go -out .. godot.yml

clean:
	rm -f cgo_helpers.go cgo_helpers.h cgo_helpers.c
	rm -f doc.go types.go const.go
	rm -f godot.go

test:
	go build
