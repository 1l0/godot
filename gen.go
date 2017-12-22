package godot

//go:generate git pull -s subtree godot_headers master

// #cgo CFLAGS: -Igodot_headers -std=c11
// #cgo linux LDFLAGS: -Wl,-unresolved-symbols=ignore-all
// #cgo darwin LDFLAGS: -Wl,-undefined,dynamic_lookup
import "C"
