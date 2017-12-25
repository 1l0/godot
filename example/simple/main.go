package main

import (
	"C"
	"log"

	"github.com/1l0/godot"
)

func main() {}

var (
	api godot.GdnativeCoreApiStruct
)

func init() {
	log.Println("init()")
}

//export godot_gdnative_init
func godot_gdnative_init(options uintptr) {
	log.Println("godot_gdnative_init()")

}

//export godot_gdnative_terminate
func godot_gdnative_terminate(options uintptr) {

}

//export godot_nativescript_init
func godot_nativescript_init(handle uintptr) {

}
