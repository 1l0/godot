package godot

//go:generate git pull -s subtree godot-headers master
//go:generate swig -go -cgo -intgosize 32 -Igodot_headers godot.swig_

// #cgo CFLAGS: -Igodot_headers -std=c11
// #cgo linux LDFLAGS: -Wl,-unresolved-symbols=ignore-all
// #cgo darwin LDFLAGS: -Wl,-undefined,dynamic_lookup
import "C"
