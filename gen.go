package godot

//go:generate git pull -s subtree godot-headers master

/*
#cgo CFLAGS: -Igodot_headers -std=c11
#cgo linux LDFLAGS: -Wl,-unresolved-symbols=ignore-all
#cgo darwin LDFLAGS: -Wl,-undefined,dynamic_lookup

#include <gdnative_api_struct.gen.h>
*/
import "C"
import (
	"log"
	"unsafe"
)

var (
	API             GdnativeCoreApiStruct
	NativescriptAPI GdnativeExtNativescriptApiStruct
	initFunc        func(Handle) = func(h Handle) {}
)

type Handle unsafe.Pointer

func Init(b func(Handle)) {
	initFunc = b
}

//export godot_gdnative_init
func godot_gdnative_init(options *C.godot_gdnative_init_options) {
	opt := GdnativeInitOptions(SwigcptrGdnativeInitOptions(unsafe.Pointer(options)))
	API = opt.GetApiStruct()
	exts := API.GetExtensions()
	for _, ext := range exts {
		if Enum_SS_GDNATIVE_API_TYPES(int(ext.GetType())) == GDNATIVEEXTNATIVESCRIPT {
			NativescriptAPI = GdnativeExtNativescriptApiStruct(SwigcptrGdnativeExtNativescriptApiStruct(ext.Swigcptr()))
			break
		}
	}
}

//export godot_gdnative_terminate
func godot_gdnative_terminate(options *C.godot_gdnative_terminate_options) {
	//Swig_free(API.Swigcptr())
	//Swig_free(NativescriptAPI.Swigcptr())
	log.Println("SEE YOU SPACE COWBOY...")
}

//export godot_nativescript_init
func godot_nativescript_init(handle unsafe.Pointer) {
	initFunc(Handle(handle))
}
