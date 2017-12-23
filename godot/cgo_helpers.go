// WARNING: This file has automatically been generated on Sun, 24 Dec 2017 06:36:13 JST.
// By https://git.io/c-for-go. DO NOT EDIT.

package godot

/*
#cgo CFLAGS: -Igodot_headers -std=c11
#cgo darwin LDFLAGS: -framework Cocoa -Wl,-undefined,dynamic_lookup
#include "gdnative_api_struct.gen.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"runtime"
	"sync"
	"unsafe"
)

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

// allocGodotGdnativeExtNativescriptApiStructMemory allocates memory for type C.godot_gdnative_ext_nativescript_api_struct in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotGdnativeExtNativescriptApiStructMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotGdnativeExtNativescriptApiStructValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotGdnativeExtNativescriptApiStructValue = unsafe.Sizeof([1]C.godot_gdnative_ext_nativescript_api_struct{})

type sliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

// allocGodotGdnativeApiStructMemory allocates memory for type C.godot_gdnative_api_struct in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotGdnativeApiStructMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotGdnativeApiStructValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotGdnativeApiStructValue = unsafe.Sizeof([1]C.godot_gdnative_api_struct{})

const sizeOfPtr = unsafe.Sizeof(&struct{}{})

// unpackSGodotGdnativeApiStruct transforms a sliced Go data structure into plain C format.
func unpackSGodotGdnativeApiStruct(x []GodotGdnativeApiStruct) (unpacked *C.godot_gdnative_api_struct, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_gdnative_api_struct) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotGdnativeApiStructMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_gdnative_api_struct)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_gdnative_api_struct)(unsafe.Pointer(h.Data))
	return
}

// packSGodotGdnativeApiStruct reads sliced Go data structure out from plain C format.
func packSGodotGdnativeApiStruct(v []GodotGdnativeApiStruct, ptr0 *C.godot_gdnative_api_struct) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotGdnativeApiStructValue]C.godot_gdnative_api_struct)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotGdnativeApiStructRef(unsafe.Pointer(&ptr1))
	}
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotGdnativeExtNativescriptApiStruct) Ref() *C.godot_gdnative_ext_nativescript_api_struct {
	if x == nil {
		return nil
	}
	return x.reff0cd6324
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotGdnativeExtNativescriptApiStruct) Free() {
	if x != nil && x.allocsf0cd6324 != nil {
		x.allocsf0cd6324.(*cgoAllocMap).Free()
		x.reff0cd6324 = nil
	}
}

// NewGodotGdnativeExtNativescriptApiStructRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotGdnativeExtNativescriptApiStructRef(ref unsafe.Pointer) *GodotGdnativeExtNativescriptApiStruct {
	if ref == nil {
		return nil
	}
	obj := new(GodotGdnativeExtNativescriptApiStruct)
	obj.reff0cd6324 = (*C.godot_gdnative_ext_nativescript_api_struct)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotGdnativeExtNativescriptApiStruct) PassRef() (*C.godot_gdnative_ext_nativescript_api_struct, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.reff0cd6324 != nil {
		return x.reff0cd6324, nil
	}
	memf0cd6324 := allocGodotGdnativeExtNativescriptApiStructMemory(1)
	reff0cd6324 := (*C.godot_gdnative_ext_nativescript_api_struct)(memf0cd6324)
	allocsf0cd6324 := new(cgoAllocMap)
	allocsf0cd6324.Add(memf0cd6324)

	var c_type_allocs *cgoAllocMap
	reff0cd6324._type, c_type_allocs = (C.uint)(x.Type), cgoAllocsUnknown
	allocsf0cd6324.Borrow(c_type_allocs)

	var cversion_allocs *cgoAllocMap
	reff0cd6324.version, cversion_allocs = x.Version.PassValue()
	allocsf0cd6324.Borrow(cversion_allocs)

	var cnext_allocs *cgoAllocMap
	reff0cd6324.next, cnext_allocs = unpackSGodotGdnativeApiStruct(x.Next)
	allocsf0cd6324.Borrow(cnext_allocs)

	var cgodot_nativescript_register_class_allocs *cgoAllocMap
	reff0cd6324.godot_nativescript_register_class, cgodot_nativescript_register_class_allocs = x.GodotNativescriptRegisterClass.PassRef()
	allocsf0cd6324.Borrow(cgodot_nativescript_register_class_allocs)

	var cgodot_nativescript_register_tool_class_allocs *cgoAllocMap
	reff0cd6324.godot_nativescript_register_tool_class, cgodot_nativescript_register_tool_class_allocs = x.GodotNativescriptRegisterToolClass.PassRef()
	allocsf0cd6324.Borrow(cgodot_nativescript_register_tool_class_allocs)

	var cgodot_nativescript_register_method_allocs *cgoAllocMap
	reff0cd6324.godot_nativescript_register_method, cgodot_nativescript_register_method_allocs = x.GodotNativescriptRegisterMethod.PassRef()
	allocsf0cd6324.Borrow(cgodot_nativescript_register_method_allocs)

	var cgodot_nativescript_register_property_allocs *cgoAllocMap
	reff0cd6324.godot_nativescript_register_property, cgodot_nativescript_register_property_allocs = x.GodotNativescriptRegisterProperty.PassRef()
	allocsf0cd6324.Borrow(cgodot_nativescript_register_property_allocs)

	var cgodot_nativescript_register_signal_allocs *cgoAllocMap
	reff0cd6324.godot_nativescript_register_signal, cgodot_nativescript_register_signal_allocs = x.GodotNativescriptRegisterSignal.PassRef()
	allocsf0cd6324.Borrow(cgodot_nativescript_register_signal_allocs)

	var cgodot_nativescript_get_userdata_allocs *cgoAllocMap
	reff0cd6324.godot_nativescript_get_userdata, cgodot_nativescript_get_userdata_allocs = x.GodotNativescriptGetUserdata.PassRef()
	allocsf0cd6324.Borrow(cgodot_nativescript_get_userdata_allocs)

	x.reff0cd6324 = reff0cd6324
	x.allocsf0cd6324 = allocsf0cd6324
	return reff0cd6324, allocsf0cd6324

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotGdnativeExtNativescriptApiStruct) PassValue() (C.godot_gdnative_ext_nativescript_api_struct, *cgoAllocMap) {
	if x.reff0cd6324 != nil {
		return *x.reff0cd6324, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotGdnativeExtNativescriptApiStruct) Deref() {
	if x.reff0cd6324 == nil {
		return
	}
	x.Type = (uint32)(x.reff0cd6324._type)
	x.Version = *NewGodotGdnativeApiVersionRef(unsafe.Pointer(&x.reff0cd6324.version))
	packSGodotGdnativeApiStruct(x.Next, x.reff0cd6324.next)
	x.GodotNativescriptRegisterClass = NewRef(unsafe.Pointer(x.reff0cd6324.godot_nativescript_register_class))
	x.GodotNativescriptRegisterToolClass = NewRef(unsafe.Pointer(x.reff0cd6324.godot_nativescript_register_tool_class))
	x.GodotNativescriptRegisterMethod = NewRef(unsafe.Pointer(x.reff0cd6324.godot_nativescript_register_method))
	x.GodotNativescriptRegisterProperty = NewRef(unsafe.Pointer(x.reff0cd6324.godot_nativescript_register_property))
	x.GodotNativescriptRegisterSignal = NewRef(unsafe.Pointer(x.reff0cd6324.godot_nativescript_register_signal))
	x.GodotNativescriptGetUserdata = NewRef(unsafe.Pointer(x.reff0cd6324.godot_nativescript_get_userdata))
}

// allocGodotGdnativeExtPluginscriptApiStructMemory allocates memory for type C.godot_gdnative_ext_pluginscript_api_struct in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotGdnativeExtPluginscriptApiStructMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotGdnativeExtPluginscriptApiStructValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotGdnativeExtPluginscriptApiStructValue = unsafe.Sizeof([1]C.godot_gdnative_ext_pluginscript_api_struct{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotGdnativeExtPluginscriptApiStruct) Ref() *C.godot_gdnative_ext_pluginscript_api_struct {
	if x == nil {
		return nil
	}
	return x.refc52e13a7
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotGdnativeExtPluginscriptApiStruct) Free() {
	if x != nil && x.allocsc52e13a7 != nil {
		x.allocsc52e13a7.(*cgoAllocMap).Free()
		x.refc52e13a7 = nil
	}
}

// NewGodotGdnativeExtPluginscriptApiStructRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotGdnativeExtPluginscriptApiStructRef(ref unsafe.Pointer) *GodotGdnativeExtPluginscriptApiStruct {
	if ref == nil {
		return nil
	}
	obj := new(GodotGdnativeExtPluginscriptApiStruct)
	obj.refc52e13a7 = (*C.godot_gdnative_ext_pluginscript_api_struct)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotGdnativeExtPluginscriptApiStruct) PassRef() (*C.godot_gdnative_ext_pluginscript_api_struct, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refc52e13a7 != nil {
		return x.refc52e13a7, nil
	}
	memc52e13a7 := allocGodotGdnativeExtPluginscriptApiStructMemory(1)
	refc52e13a7 := (*C.godot_gdnative_ext_pluginscript_api_struct)(memc52e13a7)
	allocsc52e13a7 := new(cgoAllocMap)
	allocsc52e13a7.Add(memc52e13a7)

	var c_type_allocs *cgoAllocMap
	refc52e13a7._type, c_type_allocs = (C.uint)(x.Type), cgoAllocsUnknown
	allocsc52e13a7.Borrow(c_type_allocs)

	var cversion_allocs *cgoAllocMap
	refc52e13a7.version, cversion_allocs = x.Version.PassValue()
	allocsc52e13a7.Borrow(cversion_allocs)

	var cnext_allocs *cgoAllocMap
	refc52e13a7.next, cnext_allocs = unpackSGodotGdnativeApiStruct(x.Next)
	allocsc52e13a7.Borrow(cnext_allocs)

	var cgodot_pluginscript_register_language_allocs *cgoAllocMap
	refc52e13a7.godot_pluginscript_register_language, cgodot_pluginscript_register_language_allocs = x.GodotPluginscriptRegisterLanguage.PassRef()
	allocsc52e13a7.Borrow(cgodot_pluginscript_register_language_allocs)

	x.refc52e13a7 = refc52e13a7
	x.allocsc52e13a7 = allocsc52e13a7
	return refc52e13a7, allocsc52e13a7

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotGdnativeExtPluginscriptApiStruct) PassValue() (C.godot_gdnative_ext_pluginscript_api_struct, *cgoAllocMap) {
	if x.refc52e13a7 != nil {
		return *x.refc52e13a7, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotGdnativeExtPluginscriptApiStruct) Deref() {
	if x.refc52e13a7 == nil {
		return
	}
	x.Type = (uint32)(x.refc52e13a7._type)
	x.Version = *NewGodotGdnativeApiVersionRef(unsafe.Pointer(&x.refc52e13a7.version))
	packSGodotGdnativeApiStruct(x.Next, x.refc52e13a7.next)
	x.GodotPluginscriptRegisterLanguage = NewRef(unsafe.Pointer(x.refc52e13a7.godot_pluginscript_register_language))
}

// allocGodotGdnativeExtArvrApiStructMemory allocates memory for type C.godot_gdnative_ext_arvr_api_struct in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotGdnativeExtArvrApiStructMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotGdnativeExtArvrApiStructValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotGdnativeExtArvrApiStructValue = unsafe.Sizeof([1]C.godot_gdnative_ext_arvr_api_struct{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotGdnativeExtArvrApiStruct) Ref() *C.godot_gdnative_ext_arvr_api_struct {
	if x == nil {
		return nil
	}
	return x.ref64de5bf5
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotGdnativeExtArvrApiStruct) Free() {
	if x != nil && x.allocs64de5bf5 != nil {
		x.allocs64de5bf5.(*cgoAllocMap).Free()
		x.ref64de5bf5 = nil
	}
}

// NewGodotGdnativeExtArvrApiStructRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotGdnativeExtArvrApiStructRef(ref unsafe.Pointer) *GodotGdnativeExtArvrApiStruct {
	if ref == nil {
		return nil
	}
	obj := new(GodotGdnativeExtArvrApiStruct)
	obj.ref64de5bf5 = (*C.godot_gdnative_ext_arvr_api_struct)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotGdnativeExtArvrApiStruct) PassRef() (*C.godot_gdnative_ext_arvr_api_struct, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref64de5bf5 != nil {
		return x.ref64de5bf5, nil
	}
	mem64de5bf5 := allocGodotGdnativeExtArvrApiStructMemory(1)
	ref64de5bf5 := (*C.godot_gdnative_ext_arvr_api_struct)(mem64de5bf5)
	allocs64de5bf5 := new(cgoAllocMap)
	allocs64de5bf5.Add(mem64de5bf5)

	var c_type_allocs *cgoAllocMap
	ref64de5bf5._type, c_type_allocs = (C.uint)(x.Type), cgoAllocsUnknown
	allocs64de5bf5.Borrow(c_type_allocs)

	var cversion_allocs *cgoAllocMap
	ref64de5bf5.version, cversion_allocs = x.Version.PassValue()
	allocs64de5bf5.Borrow(cversion_allocs)

	var cnext_allocs *cgoAllocMap
	ref64de5bf5.next, cnext_allocs = unpackSGodotGdnativeApiStruct(x.Next)
	allocs64de5bf5.Borrow(cnext_allocs)

	var cgodot_arvr_register_interface_allocs *cgoAllocMap
	ref64de5bf5.godot_arvr_register_interface, cgodot_arvr_register_interface_allocs = x.GodotArvrRegisterInterface.PassRef()
	allocs64de5bf5.Borrow(cgodot_arvr_register_interface_allocs)

	var cgodot_arvr_get_worldscale_allocs *cgoAllocMap
	ref64de5bf5.godot_arvr_get_worldscale, cgodot_arvr_get_worldscale_allocs = x.GodotArvrGetWorldscale.PassRef()
	allocs64de5bf5.Borrow(cgodot_arvr_get_worldscale_allocs)

	var cgodot_arvr_get_reference_frame_allocs *cgoAllocMap
	ref64de5bf5.godot_arvr_get_reference_frame, cgodot_arvr_get_reference_frame_allocs = x.GodotArvrGetReferenceFrame.PassRef()
	allocs64de5bf5.Borrow(cgodot_arvr_get_reference_frame_allocs)

	var cgodot_arvr_blit_allocs *cgoAllocMap
	ref64de5bf5.godot_arvr_blit, cgodot_arvr_blit_allocs = x.GodotArvrBlit.PassRef()
	allocs64de5bf5.Borrow(cgodot_arvr_blit_allocs)

	var cgodot_arvr_get_texid_allocs *cgoAllocMap
	ref64de5bf5.godot_arvr_get_texid, cgodot_arvr_get_texid_allocs = x.GodotArvrGetTexid.PassRef()
	allocs64de5bf5.Borrow(cgodot_arvr_get_texid_allocs)

	var cgodot_arvr_add_controller_allocs *cgoAllocMap
	ref64de5bf5.godot_arvr_add_controller, cgodot_arvr_add_controller_allocs = x.GodotArvrAddController.PassRef()
	allocs64de5bf5.Borrow(cgodot_arvr_add_controller_allocs)

	var cgodot_arvr_remove_controller_allocs *cgoAllocMap
	ref64de5bf5.godot_arvr_remove_controller, cgodot_arvr_remove_controller_allocs = x.GodotArvrRemoveController.PassRef()
	allocs64de5bf5.Borrow(cgodot_arvr_remove_controller_allocs)

	var cgodot_arvr_set_controller_transform_allocs *cgoAllocMap
	ref64de5bf5.godot_arvr_set_controller_transform, cgodot_arvr_set_controller_transform_allocs = x.GodotArvrSetControllerTransform.PassRef()
	allocs64de5bf5.Borrow(cgodot_arvr_set_controller_transform_allocs)

	var cgodot_arvr_set_controller_button_allocs *cgoAllocMap
	ref64de5bf5.godot_arvr_set_controller_button, cgodot_arvr_set_controller_button_allocs = x.GodotArvrSetControllerButton.PassRef()
	allocs64de5bf5.Borrow(cgodot_arvr_set_controller_button_allocs)

	var cgodot_arvr_set_controller_axis_allocs *cgoAllocMap
	ref64de5bf5.godot_arvr_set_controller_axis, cgodot_arvr_set_controller_axis_allocs = x.GodotArvrSetControllerAxis.PassRef()
	allocs64de5bf5.Borrow(cgodot_arvr_set_controller_axis_allocs)

	var cgodot_arvr_get_controller_rumble_allocs *cgoAllocMap
	ref64de5bf5.godot_arvr_get_controller_rumble, cgodot_arvr_get_controller_rumble_allocs = x.GodotArvrGetControllerRumble.PassRef()
	allocs64de5bf5.Borrow(cgodot_arvr_get_controller_rumble_allocs)

	x.ref64de5bf5 = ref64de5bf5
	x.allocs64de5bf5 = allocs64de5bf5
	return ref64de5bf5, allocs64de5bf5

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotGdnativeExtArvrApiStruct) PassValue() (C.godot_gdnative_ext_arvr_api_struct, *cgoAllocMap) {
	if x.ref64de5bf5 != nil {
		return *x.ref64de5bf5, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotGdnativeExtArvrApiStruct) Deref() {
	if x.ref64de5bf5 == nil {
		return
	}
	x.Type = (uint32)(x.ref64de5bf5._type)
	x.Version = *NewGodotGdnativeApiVersionRef(unsafe.Pointer(&x.ref64de5bf5.version))
	packSGodotGdnativeApiStruct(x.Next, x.ref64de5bf5.next)
	x.GodotArvrRegisterInterface = NewRef(unsafe.Pointer(x.ref64de5bf5.godot_arvr_register_interface))
	x.GodotArvrGetWorldscale = NewGodotRealRef(unsafe.Pointer(x.ref64de5bf5.godot_arvr_get_worldscale))
	x.GodotArvrGetReferenceFrame = NewGodotTransformRef(unsafe.Pointer(x.ref64de5bf5.godot_arvr_get_reference_frame))
	x.GodotArvrBlit = NewRef(unsafe.Pointer(x.ref64de5bf5.godot_arvr_blit))
	x.GodotArvrGetTexid = NewGodotIntRef(unsafe.Pointer(x.ref64de5bf5.godot_arvr_get_texid))
	x.GodotArvrAddController = NewGodotIntRef(unsafe.Pointer(x.ref64de5bf5.godot_arvr_add_controller))
	x.GodotArvrRemoveController = NewRef(unsafe.Pointer(x.ref64de5bf5.godot_arvr_remove_controller))
	x.GodotArvrSetControllerTransform = NewRef(unsafe.Pointer(x.ref64de5bf5.godot_arvr_set_controller_transform))
	x.GodotArvrSetControllerButton = NewRef(unsafe.Pointer(x.ref64de5bf5.godot_arvr_set_controller_button))
	x.GodotArvrSetControllerAxis = NewRef(unsafe.Pointer(x.ref64de5bf5.godot_arvr_set_controller_axis))
	x.GodotArvrGetControllerRumble = NewGodotRealRef(unsafe.Pointer(x.ref64de5bf5.godot_arvr_get_controller_rumble))
}

// allocGodotGdnativeCoreApiStructMemory allocates memory for type C.godot_gdnative_core_api_struct in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotGdnativeCoreApiStructMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotGdnativeCoreApiStructValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotGdnativeCoreApiStructValue = unsafe.Sizeof([1]C.godot_gdnative_core_api_struct{})

// allocPGodotGdnativeApiStructMemory allocates memory for type *C.godot_gdnative_api_struct in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPGodotGdnativeApiStructMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPGodotGdnativeApiStructValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPGodotGdnativeApiStructValue = unsafe.Sizeof([1]*C.godot_gdnative_api_struct{})

// unpackSSGodotGdnativeApiStruct transforms a sliced Go data structure into plain C format.
func unpackSSGodotGdnativeApiStruct(x [][]GodotGdnativeApiStruct) (unpacked **C.godot_gdnative_api_struct, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.godot_gdnative_api_struct) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPGodotGdnativeApiStructMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.godot_gdnative_api_struct)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocGodotGdnativeApiStructMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]C.godot_gdnative_api_struct)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			allocs1 := new(cgoAllocMap)
			v1[i1], allocs1 = x[i0][i1].PassValue()
			allocs.Borrow(allocs1)
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (*C.godot_gdnative_api_struct)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.godot_gdnative_api_struct)(unsafe.Pointer(h.Data))
	return
}

// packSSGodotGdnativeApiStruct reads sliced Go data structure out from plain C format.
func packSSGodotGdnativeApiStruct(v [][]GodotGdnativeApiStruct, ptr0 **C.godot_gdnative_api_struct) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.godot_gdnative_api_struct)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfGodotGdnativeApiStructValue]C.godot_gdnative_api_struct)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *NewGodotGdnativeApiStructRef(unsafe.Pointer(&ptr2))
		}
	}
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotGdnativeCoreApiStruct) Ref() *C.godot_gdnative_core_api_struct {
	if x == nil {
		return nil
	}
	return x.ref57717e51
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotGdnativeCoreApiStruct) Free() {
	if x != nil && x.allocs57717e51 != nil {
		x.allocs57717e51.(*cgoAllocMap).Free()
		x.ref57717e51 = nil
	}
}

// NewGodotGdnativeCoreApiStructRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotGdnativeCoreApiStructRef(ref unsafe.Pointer) *GodotGdnativeCoreApiStruct {
	if ref == nil {
		return nil
	}
	obj := new(GodotGdnativeCoreApiStruct)
	obj.ref57717e51 = (*C.godot_gdnative_core_api_struct)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotGdnativeCoreApiStruct) PassRef() (*C.godot_gdnative_core_api_struct, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref57717e51 != nil {
		return x.ref57717e51, nil
	}
	mem57717e51 := allocGodotGdnativeCoreApiStructMemory(1)
	ref57717e51 := (*C.godot_gdnative_core_api_struct)(mem57717e51)
	allocs57717e51 := new(cgoAllocMap)
	allocs57717e51.Add(mem57717e51)

	var c_type_allocs *cgoAllocMap
	ref57717e51._type, c_type_allocs = (C.uint)(x.Type), cgoAllocsUnknown
	allocs57717e51.Borrow(c_type_allocs)

	var cversion_allocs *cgoAllocMap
	ref57717e51.version, cversion_allocs = x.Version.PassValue()
	allocs57717e51.Borrow(cversion_allocs)

	var cnext_allocs *cgoAllocMap
	ref57717e51.next, cnext_allocs = unpackSGodotGdnativeApiStruct(x.Next)
	allocs57717e51.Borrow(cnext_allocs)

	var cnum_extensions_allocs *cgoAllocMap
	ref57717e51.num_extensions, cnum_extensions_allocs = (C.uint)(x.NumExtensions), cgoAllocsUnknown
	allocs57717e51.Borrow(cnum_extensions_allocs)

	var cextensions_allocs *cgoAllocMap
	ref57717e51.extensions, cextensions_allocs = unpackSSGodotGdnativeApiStruct(x.Extensions)
	allocs57717e51.Borrow(cextensions_allocs)

	var cgodot_color_new_rgba_allocs *cgoAllocMap
	ref57717e51.godot_color_new_rgba, cgodot_color_new_rgba_allocs = x.GodotColorNewRgba.PassRef()
	allocs57717e51.Borrow(cgodot_color_new_rgba_allocs)

	var cgodot_color_new_rgb_allocs *cgoAllocMap
	ref57717e51.godot_color_new_rgb, cgodot_color_new_rgb_allocs = x.GodotColorNewRgb.PassRef()
	allocs57717e51.Borrow(cgodot_color_new_rgb_allocs)

	var cgodot_color_get_r_allocs *cgoAllocMap
	ref57717e51.godot_color_get_r, cgodot_color_get_r_allocs = x.GodotColorGetR.PassRef()
	allocs57717e51.Borrow(cgodot_color_get_r_allocs)

	var cgodot_color_set_r_allocs *cgoAllocMap
	ref57717e51.godot_color_set_r, cgodot_color_set_r_allocs = x.GodotColorSetR.PassRef()
	allocs57717e51.Borrow(cgodot_color_set_r_allocs)

	var cgodot_color_get_g_allocs *cgoAllocMap
	ref57717e51.godot_color_get_g, cgodot_color_get_g_allocs = x.GodotColorGetG.PassRef()
	allocs57717e51.Borrow(cgodot_color_get_g_allocs)

	var cgodot_color_set_g_allocs *cgoAllocMap
	ref57717e51.godot_color_set_g, cgodot_color_set_g_allocs = x.GodotColorSetG.PassRef()
	allocs57717e51.Borrow(cgodot_color_set_g_allocs)

	var cgodot_color_get_b_allocs *cgoAllocMap
	ref57717e51.godot_color_get_b, cgodot_color_get_b_allocs = x.GodotColorGetB.PassRef()
	allocs57717e51.Borrow(cgodot_color_get_b_allocs)

	var cgodot_color_set_b_allocs *cgoAllocMap
	ref57717e51.godot_color_set_b, cgodot_color_set_b_allocs = x.GodotColorSetB.PassRef()
	allocs57717e51.Borrow(cgodot_color_set_b_allocs)

	var cgodot_color_get_a_allocs *cgoAllocMap
	ref57717e51.godot_color_get_a, cgodot_color_get_a_allocs = x.GodotColorGetA.PassRef()
	allocs57717e51.Borrow(cgodot_color_get_a_allocs)

	var cgodot_color_set_a_allocs *cgoAllocMap
	ref57717e51.godot_color_set_a, cgodot_color_set_a_allocs = x.GodotColorSetA.PassRef()
	allocs57717e51.Borrow(cgodot_color_set_a_allocs)

	var cgodot_color_get_h_allocs *cgoAllocMap
	ref57717e51.godot_color_get_h, cgodot_color_get_h_allocs = x.GodotColorGetH.PassRef()
	allocs57717e51.Borrow(cgodot_color_get_h_allocs)

	var cgodot_color_get_s_allocs *cgoAllocMap
	ref57717e51.godot_color_get_s, cgodot_color_get_s_allocs = x.GodotColorGetS.PassRef()
	allocs57717e51.Borrow(cgodot_color_get_s_allocs)

	var cgodot_color_get_v_allocs *cgoAllocMap
	ref57717e51.godot_color_get_v, cgodot_color_get_v_allocs = x.GodotColorGetV.PassRef()
	allocs57717e51.Borrow(cgodot_color_get_v_allocs)

	var cgodot_color_as_string_allocs *cgoAllocMap
	ref57717e51.godot_color_as_string, cgodot_color_as_string_allocs = x.GodotColorAsString.PassRef()
	allocs57717e51.Borrow(cgodot_color_as_string_allocs)

	var cgodot_color_to_rgba32_allocs *cgoAllocMap
	ref57717e51.godot_color_to_rgba32, cgodot_color_to_rgba32_allocs = x.GodotColorToRgba32.PassRef()
	allocs57717e51.Borrow(cgodot_color_to_rgba32_allocs)

	var cgodot_color_to_argb32_allocs *cgoAllocMap
	ref57717e51.godot_color_to_argb32, cgodot_color_to_argb32_allocs = x.GodotColorToArgb32.PassRef()
	allocs57717e51.Borrow(cgodot_color_to_argb32_allocs)

	var cgodot_color_gray_allocs *cgoAllocMap
	ref57717e51.godot_color_gray, cgodot_color_gray_allocs = x.GodotColorGray.PassRef()
	allocs57717e51.Borrow(cgodot_color_gray_allocs)

	var cgodot_color_inverted_allocs *cgoAllocMap
	ref57717e51.godot_color_inverted, cgodot_color_inverted_allocs = x.GodotColorInverted.PassRef()
	allocs57717e51.Borrow(cgodot_color_inverted_allocs)

	var cgodot_color_contrasted_allocs *cgoAllocMap
	ref57717e51.godot_color_contrasted, cgodot_color_contrasted_allocs = x.GodotColorContrasted.PassRef()
	allocs57717e51.Borrow(cgodot_color_contrasted_allocs)

	var cgodot_color_linear_interpolate_allocs *cgoAllocMap
	ref57717e51.godot_color_linear_interpolate, cgodot_color_linear_interpolate_allocs = x.GodotColorLinearInterpolate.PassRef()
	allocs57717e51.Borrow(cgodot_color_linear_interpolate_allocs)

	var cgodot_color_blend_allocs *cgoAllocMap
	ref57717e51.godot_color_blend, cgodot_color_blend_allocs = x.GodotColorBlend.PassRef()
	allocs57717e51.Borrow(cgodot_color_blend_allocs)

	var cgodot_color_to_html_allocs *cgoAllocMap
	ref57717e51.godot_color_to_html, cgodot_color_to_html_allocs = x.GodotColorToHtml.PassRef()
	allocs57717e51.Borrow(cgodot_color_to_html_allocs)

	var cgodot_color_operator_equal_allocs *cgoAllocMap
	ref57717e51.godot_color_operator_equal, cgodot_color_operator_equal_allocs = x.GodotColorOperatorEqual.PassRef()
	allocs57717e51.Borrow(cgodot_color_operator_equal_allocs)

	var cgodot_color_operator_less_allocs *cgoAllocMap
	ref57717e51.godot_color_operator_less, cgodot_color_operator_less_allocs = x.GodotColorOperatorLess.PassRef()
	allocs57717e51.Borrow(cgodot_color_operator_less_allocs)

	var cgodot_vector2_new_allocs *cgoAllocMap
	ref57717e51.godot_vector2_new, cgodot_vector2_new_allocs = x.GodotVector2New.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_new_allocs)

	var cgodot_vector2_as_string_allocs *cgoAllocMap
	ref57717e51.godot_vector2_as_string, cgodot_vector2_as_string_allocs = x.GodotVector2AsString.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_as_string_allocs)

	var cgodot_vector2_normalized_allocs *cgoAllocMap
	ref57717e51.godot_vector2_normalized, cgodot_vector2_normalized_allocs = x.GodotVector2Normalized.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_normalized_allocs)

	var cgodot_vector2_length_allocs *cgoAllocMap
	ref57717e51.godot_vector2_length, cgodot_vector2_length_allocs = x.GodotVector2Length.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_length_allocs)

	var cgodot_vector2_angle_allocs *cgoAllocMap
	ref57717e51.godot_vector2_angle, cgodot_vector2_angle_allocs = x.GodotVector2Angle.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_angle_allocs)

	var cgodot_vector2_length_squared_allocs *cgoAllocMap
	ref57717e51.godot_vector2_length_squared, cgodot_vector2_length_squared_allocs = x.GodotVector2LengthSquared.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_length_squared_allocs)

	var cgodot_vector2_is_normalized_allocs *cgoAllocMap
	ref57717e51.godot_vector2_is_normalized, cgodot_vector2_is_normalized_allocs = x.GodotVector2IsNormalized.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_is_normalized_allocs)

	var cgodot_vector2_distance_to_allocs *cgoAllocMap
	ref57717e51.godot_vector2_distance_to, cgodot_vector2_distance_to_allocs = x.GodotVector2DistanceTo.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_distance_to_allocs)

	var cgodot_vector2_distance_squared_to_allocs *cgoAllocMap
	ref57717e51.godot_vector2_distance_squared_to, cgodot_vector2_distance_squared_to_allocs = x.GodotVector2DistanceSquaredTo.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_distance_squared_to_allocs)

	var cgodot_vector2_angle_to_allocs *cgoAllocMap
	ref57717e51.godot_vector2_angle_to, cgodot_vector2_angle_to_allocs = x.GodotVector2AngleTo.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_angle_to_allocs)

	var cgodot_vector2_angle_to_point_allocs *cgoAllocMap
	ref57717e51.godot_vector2_angle_to_point, cgodot_vector2_angle_to_point_allocs = x.GodotVector2AngleToPoint.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_angle_to_point_allocs)

	var cgodot_vector2_linear_interpolate_allocs *cgoAllocMap
	ref57717e51.godot_vector2_linear_interpolate, cgodot_vector2_linear_interpolate_allocs = x.GodotVector2LinearInterpolate.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_linear_interpolate_allocs)

	var cgodot_vector2_cubic_interpolate_allocs *cgoAllocMap
	ref57717e51.godot_vector2_cubic_interpolate, cgodot_vector2_cubic_interpolate_allocs = x.GodotVector2CubicInterpolate.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_cubic_interpolate_allocs)

	var cgodot_vector2_rotated_allocs *cgoAllocMap
	ref57717e51.godot_vector2_rotated, cgodot_vector2_rotated_allocs = x.GodotVector2Rotated.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_rotated_allocs)

	var cgodot_vector2_tangent_allocs *cgoAllocMap
	ref57717e51.godot_vector2_tangent, cgodot_vector2_tangent_allocs = x.GodotVector2Tangent.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_tangent_allocs)

	var cgodot_vector2_floor_allocs *cgoAllocMap
	ref57717e51.godot_vector2_floor, cgodot_vector2_floor_allocs = x.GodotVector2Floor.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_floor_allocs)

	var cgodot_vector2_snapped_allocs *cgoAllocMap
	ref57717e51.godot_vector2_snapped, cgodot_vector2_snapped_allocs = x.GodotVector2Snapped.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_snapped_allocs)

	var cgodot_vector2_aspect_allocs *cgoAllocMap
	ref57717e51.godot_vector2_aspect, cgodot_vector2_aspect_allocs = x.GodotVector2Aspect.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_aspect_allocs)

	var cgodot_vector2_dot_allocs *cgoAllocMap
	ref57717e51.godot_vector2_dot, cgodot_vector2_dot_allocs = x.GodotVector2Dot.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_dot_allocs)

	var cgodot_vector2_slide_allocs *cgoAllocMap
	ref57717e51.godot_vector2_slide, cgodot_vector2_slide_allocs = x.GodotVector2Slide.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_slide_allocs)

	var cgodot_vector2_bounce_allocs *cgoAllocMap
	ref57717e51.godot_vector2_bounce, cgodot_vector2_bounce_allocs = x.GodotVector2Bounce.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_bounce_allocs)

	var cgodot_vector2_reflect_allocs *cgoAllocMap
	ref57717e51.godot_vector2_reflect, cgodot_vector2_reflect_allocs = x.GodotVector2Reflect.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_reflect_allocs)

	var cgodot_vector2_abs_allocs *cgoAllocMap
	ref57717e51.godot_vector2_abs, cgodot_vector2_abs_allocs = x.GodotVector2Abs.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_abs_allocs)

	var cgodot_vector2_clamped_allocs *cgoAllocMap
	ref57717e51.godot_vector2_clamped, cgodot_vector2_clamped_allocs = x.GodotVector2Clamped.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_clamped_allocs)

	var cgodot_vector2_operator_add_allocs *cgoAllocMap
	ref57717e51.godot_vector2_operator_add, cgodot_vector2_operator_add_allocs = x.GodotVector2OperatorAdd.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_operator_add_allocs)

	var cgodot_vector2_operator_substract_allocs *cgoAllocMap
	ref57717e51.godot_vector2_operator_substract, cgodot_vector2_operator_substract_allocs = x.GodotVector2OperatorSubstract.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_operator_substract_allocs)

	var cgodot_vector2_operator_multiply_vector_allocs *cgoAllocMap
	ref57717e51.godot_vector2_operator_multiply_vector, cgodot_vector2_operator_multiply_vector_allocs = x.GodotVector2OperatorMultiplyVector.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_operator_multiply_vector_allocs)

	var cgodot_vector2_operator_multiply_scalar_allocs *cgoAllocMap
	ref57717e51.godot_vector2_operator_multiply_scalar, cgodot_vector2_operator_multiply_scalar_allocs = x.GodotVector2OperatorMultiplyScalar.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_operator_multiply_scalar_allocs)

	var cgodot_vector2_operator_divide_vector_allocs *cgoAllocMap
	ref57717e51.godot_vector2_operator_divide_vector, cgodot_vector2_operator_divide_vector_allocs = x.GodotVector2OperatorDivideVector.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_operator_divide_vector_allocs)

	var cgodot_vector2_operator_divide_scalar_allocs *cgoAllocMap
	ref57717e51.godot_vector2_operator_divide_scalar, cgodot_vector2_operator_divide_scalar_allocs = x.GodotVector2OperatorDivideScalar.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_operator_divide_scalar_allocs)

	var cgodot_vector2_operator_equal_allocs *cgoAllocMap
	ref57717e51.godot_vector2_operator_equal, cgodot_vector2_operator_equal_allocs = x.GodotVector2OperatorEqual.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_operator_equal_allocs)

	var cgodot_vector2_operator_less_allocs *cgoAllocMap
	ref57717e51.godot_vector2_operator_less, cgodot_vector2_operator_less_allocs = x.GodotVector2OperatorLess.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_operator_less_allocs)

	var cgodot_vector2_operator_neg_allocs *cgoAllocMap
	ref57717e51.godot_vector2_operator_neg, cgodot_vector2_operator_neg_allocs = x.GodotVector2OperatorNeg.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_operator_neg_allocs)

	var cgodot_vector2_set_x_allocs *cgoAllocMap
	ref57717e51.godot_vector2_set_x, cgodot_vector2_set_x_allocs = x.GodotVector2SetX.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_set_x_allocs)

	var cgodot_vector2_set_y_allocs *cgoAllocMap
	ref57717e51.godot_vector2_set_y, cgodot_vector2_set_y_allocs = x.GodotVector2SetY.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_set_y_allocs)

	var cgodot_vector2_get_x_allocs *cgoAllocMap
	ref57717e51.godot_vector2_get_x, cgodot_vector2_get_x_allocs = x.GodotVector2GetX.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_get_x_allocs)

	var cgodot_vector2_get_y_allocs *cgoAllocMap
	ref57717e51.godot_vector2_get_y, cgodot_vector2_get_y_allocs = x.GodotVector2GetY.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_get_y_allocs)

	var cgodot_quat_new_allocs *cgoAllocMap
	ref57717e51.godot_quat_new, cgodot_quat_new_allocs = x.GodotQuatNew.PassRef()
	allocs57717e51.Borrow(cgodot_quat_new_allocs)

	var cgodot_quat_new_with_axis_angle_allocs *cgoAllocMap
	ref57717e51.godot_quat_new_with_axis_angle, cgodot_quat_new_with_axis_angle_allocs = x.GodotQuatNewWithAxisAngle.PassRef()
	allocs57717e51.Borrow(cgodot_quat_new_with_axis_angle_allocs)

	var cgodot_quat_get_x_allocs *cgoAllocMap
	ref57717e51.godot_quat_get_x, cgodot_quat_get_x_allocs = x.GodotQuatGetX.PassRef()
	allocs57717e51.Borrow(cgodot_quat_get_x_allocs)

	var cgodot_quat_set_x_allocs *cgoAllocMap
	ref57717e51.godot_quat_set_x, cgodot_quat_set_x_allocs = x.GodotQuatSetX.PassRef()
	allocs57717e51.Borrow(cgodot_quat_set_x_allocs)

	var cgodot_quat_get_y_allocs *cgoAllocMap
	ref57717e51.godot_quat_get_y, cgodot_quat_get_y_allocs = x.GodotQuatGetY.PassRef()
	allocs57717e51.Borrow(cgodot_quat_get_y_allocs)

	var cgodot_quat_set_y_allocs *cgoAllocMap
	ref57717e51.godot_quat_set_y, cgodot_quat_set_y_allocs = x.GodotQuatSetY.PassRef()
	allocs57717e51.Borrow(cgodot_quat_set_y_allocs)

	var cgodot_quat_get_z_allocs *cgoAllocMap
	ref57717e51.godot_quat_get_z, cgodot_quat_get_z_allocs = x.GodotQuatGetZ.PassRef()
	allocs57717e51.Borrow(cgodot_quat_get_z_allocs)

	var cgodot_quat_set_z_allocs *cgoAllocMap
	ref57717e51.godot_quat_set_z, cgodot_quat_set_z_allocs = x.GodotQuatSetZ.PassRef()
	allocs57717e51.Borrow(cgodot_quat_set_z_allocs)

	var cgodot_quat_get_w_allocs *cgoAllocMap
	ref57717e51.godot_quat_get_w, cgodot_quat_get_w_allocs = x.GodotQuatGetW.PassRef()
	allocs57717e51.Borrow(cgodot_quat_get_w_allocs)

	var cgodot_quat_set_w_allocs *cgoAllocMap
	ref57717e51.godot_quat_set_w, cgodot_quat_set_w_allocs = x.GodotQuatSetW.PassRef()
	allocs57717e51.Borrow(cgodot_quat_set_w_allocs)

	var cgodot_quat_as_string_allocs *cgoAllocMap
	ref57717e51.godot_quat_as_string, cgodot_quat_as_string_allocs = x.GodotQuatAsString.PassRef()
	allocs57717e51.Borrow(cgodot_quat_as_string_allocs)

	var cgodot_quat_length_allocs *cgoAllocMap
	ref57717e51.godot_quat_length, cgodot_quat_length_allocs = x.GodotQuatLength.PassRef()
	allocs57717e51.Borrow(cgodot_quat_length_allocs)

	var cgodot_quat_length_squared_allocs *cgoAllocMap
	ref57717e51.godot_quat_length_squared, cgodot_quat_length_squared_allocs = x.GodotQuatLengthSquared.PassRef()
	allocs57717e51.Borrow(cgodot_quat_length_squared_allocs)

	var cgodot_quat_normalized_allocs *cgoAllocMap
	ref57717e51.godot_quat_normalized, cgodot_quat_normalized_allocs = x.GodotQuatNormalized.PassRef()
	allocs57717e51.Borrow(cgodot_quat_normalized_allocs)

	var cgodot_quat_is_normalized_allocs *cgoAllocMap
	ref57717e51.godot_quat_is_normalized, cgodot_quat_is_normalized_allocs = x.GodotQuatIsNormalized.PassRef()
	allocs57717e51.Borrow(cgodot_quat_is_normalized_allocs)

	var cgodot_quat_inverse_allocs *cgoAllocMap
	ref57717e51.godot_quat_inverse, cgodot_quat_inverse_allocs = x.GodotQuatInverse.PassRef()
	allocs57717e51.Borrow(cgodot_quat_inverse_allocs)

	var cgodot_quat_dot_allocs *cgoAllocMap
	ref57717e51.godot_quat_dot, cgodot_quat_dot_allocs = x.GodotQuatDot.PassRef()
	allocs57717e51.Borrow(cgodot_quat_dot_allocs)

	var cgodot_quat_xform_allocs *cgoAllocMap
	ref57717e51.godot_quat_xform, cgodot_quat_xform_allocs = x.GodotQuatXform.PassRef()
	allocs57717e51.Borrow(cgodot_quat_xform_allocs)

	var cgodot_quat_slerp_allocs *cgoAllocMap
	ref57717e51.godot_quat_slerp, cgodot_quat_slerp_allocs = x.GodotQuatSlerp.PassRef()
	allocs57717e51.Borrow(cgodot_quat_slerp_allocs)

	var cgodot_quat_slerpni_allocs *cgoAllocMap
	ref57717e51.godot_quat_slerpni, cgodot_quat_slerpni_allocs = x.GodotQuatSlerpni.PassRef()
	allocs57717e51.Borrow(cgodot_quat_slerpni_allocs)

	var cgodot_quat_cubic_slerp_allocs *cgoAllocMap
	ref57717e51.godot_quat_cubic_slerp, cgodot_quat_cubic_slerp_allocs = x.GodotQuatCubicSlerp.PassRef()
	allocs57717e51.Borrow(cgodot_quat_cubic_slerp_allocs)

	var cgodot_quat_operator_multiply_allocs *cgoAllocMap
	ref57717e51.godot_quat_operator_multiply, cgodot_quat_operator_multiply_allocs = x.GodotQuatOperatorMultiply.PassRef()
	allocs57717e51.Borrow(cgodot_quat_operator_multiply_allocs)

	var cgodot_quat_operator_add_allocs *cgoAllocMap
	ref57717e51.godot_quat_operator_add, cgodot_quat_operator_add_allocs = x.GodotQuatOperatorAdd.PassRef()
	allocs57717e51.Borrow(cgodot_quat_operator_add_allocs)

	var cgodot_quat_operator_substract_allocs *cgoAllocMap
	ref57717e51.godot_quat_operator_substract, cgodot_quat_operator_substract_allocs = x.GodotQuatOperatorSubstract.PassRef()
	allocs57717e51.Borrow(cgodot_quat_operator_substract_allocs)

	var cgodot_quat_operator_divide_allocs *cgoAllocMap
	ref57717e51.godot_quat_operator_divide, cgodot_quat_operator_divide_allocs = x.GodotQuatOperatorDivide.PassRef()
	allocs57717e51.Borrow(cgodot_quat_operator_divide_allocs)

	var cgodot_quat_operator_equal_allocs *cgoAllocMap
	ref57717e51.godot_quat_operator_equal, cgodot_quat_operator_equal_allocs = x.GodotQuatOperatorEqual.PassRef()
	allocs57717e51.Borrow(cgodot_quat_operator_equal_allocs)

	var cgodot_quat_operator_neg_allocs *cgoAllocMap
	ref57717e51.godot_quat_operator_neg, cgodot_quat_operator_neg_allocs = x.GodotQuatOperatorNeg.PassRef()
	allocs57717e51.Borrow(cgodot_quat_operator_neg_allocs)

	var cgodot_basis_new_with_rows_allocs *cgoAllocMap
	ref57717e51.godot_basis_new_with_rows, cgodot_basis_new_with_rows_allocs = x.GodotBasisNewWithRows.PassRef()
	allocs57717e51.Borrow(cgodot_basis_new_with_rows_allocs)

	var cgodot_basis_new_with_axis_and_angle_allocs *cgoAllocMap
	ref57717e51.godot_basis_new_with_axis_and_angle, cgodot_basis_new_with_axis_and_angle_allocs = x.GodotBasisNewWithAxisAndAngle.PassRef()
	allocs57717e51.Borrow(cgodot_basis_new_with_axis_and_angle_allocs)

	var cgodot_basis_new_with_euler_allocs *cgoAllocMap
	ref57717e51.godot_basis_new_with_euler, cgodot_basis_new_with_euler_allocs = x.GodotBasisNewWithEuler.PassRef()
	allocs57717e51.Borrow(cgodot_basis_new_with_euler_allocs)

	var cgodot_basis_as_string_allocs *cgoAllocMap
	ref57717e51.godot_basis_as_string, cgodot_basis_as_string_allocs = x.GodotBasisAsString.PassRef()
	allocs57717e51.Borrow(cgodot_basis_as_string_allocs)

	var cgodot_basis_inverse_allocs *cgoAllocMap
	ref57717e51.godot_basis_inverse, cgodot_basis_inverse_allocs = x.GodotBasisInverse.PassRef()
	allocs57717e51.Borrow(cgodot_basis_inverse_allocs)

	var cgodot_basis_transposed_allocs *cgoAllocMap
	ref57717e51.godot_basis_transposed, cgodot_basis_transposed_allocs = x.GodotBasisTransposed.PassRef()
	allocs57717e51.Borrow(cgodot_basis_transposed_allocs)

	var cgodot_basis_orthonormalized_allocs *cgoAllocMap
	ref57717e51.godot_basis_orthonormalized, cgodot_basis_orthonormalized_allocs = x.GodotBasisOrthonormalized.PassRef()
	allocs57717e51.Borrow(cgodot_basis_orthonormalized_allocs)

	var cgodot_basis_determinant_allocs *cgoAllocMap
	ref57717e51.godot_basis_determinant, cgodot_basis_determinant_allocs = x.GodotBasisDeterminant.PassRef()
	allocs57717e51.Borrow(cgodot_basis_determinant_allocs)

	var cgodot_basis_rotated_allocs *cgoAllocMap
	ref57717e51.godot_basis_rotated, cgodot_basis_rotated_allocs = x.GodotBasisRotated.PassRef()
	allocs57717e51.Borrow(cgodot_basis_rotated_allocs)

	var cgodot_basis_scaled_allocs *cgoAllocMap
	ref57717e51.godot_basis_scaled, cgodot_basis_scaled_allocs = x.GodotBasisScaled.PassRef()
	allocs57717e51.Borrow(cgodot_basis_scaled_allocs)

	var cgodot_basis_get_scale_allocs *cgoAllocMap
	ref57717e51.godot_basis_get_scale, cgodot_basis_get_scale_allocs = x.GodotBasisGetScale.PassRef()
	allocs57717e51.Borrow(cgodot_basis_get_scale_allocs)

	var cgodot_basis_get_euler_allocs *cgoAllocMap
	ref57717e51.godot_basis_get_euler, cgodot_basis_get_euler_allocs = x.GodotBasisGetEuler.PassRef()
	allocs57717e51.Borrow(cgodot_basis_get_euler_allocs)

	var cgodot_basis_tdotx_allocs *cgoAllocMap
	ref57717e51.godot_basis_tdotx, cgodot_basis_tdotx_allocs = x.GodotBasisTdotx.PassRef()
	allocs57717e51.Borrow(cgodot_basis_tdotx_allocs)

	var cgodot_basis_tdoty_allocs *cgoAllocMap
	ref57717e51.godot_basis_tdoty, cgodot_basis_tdoty_allocs = x.GodotBasisTdoty.PassRef()
	allocs57717e51.Borrow(cgodot_basis_tdoty_allocs)

	var cgodot_basis_tdotz_allocs *cgoAllocMap
	ref57717e51.godot_basis_tdotz, cgodot_basis_tdotz_allocs = x.GodotBasisTdotz.PassRef()
	allocs57717e51.Borrow(cgodot_basis_tdotz_allocs)

	var cgodot_basis_xform_allocs *cgoAllocMap
	ref57717e51.godot_basis_xform, cgodot_basis_xform_allocs = x.GodotBasisXform.PassRef()
	allocs57717e51.Borrow(cgodot_basis_xform_allocs)

	var cgodot_basis_xform_inv_allocs *cgoAllocMap
	ref57717e51.godot_basis_xform_inv, cgodot_basis_xform_inv_allocs = x.GodotBasisXformInv.PassRef()
	allocs57717e51.Borrow(cgodot_basis_xform_inv_allocs)

	var cgodot_basis_get_orthogonal_index_allocs *cgoAllocMap
	ref57717e51.godot_basis_get_orthogonal_index, cgodot_basis_get_orthogonal_index_allocs = x.GodotBasisGetOrthogonalIndex.PassRef()
	allocs57717e51.Borrow(cgodot_basis_get_orthogonal_index_allocs)

	var cgodot_basis_new_allocs *cgoAllocMap
	ref57717e51.godot_basis_new, cgodot_basis_new_allocs = x.GodotBasisNew.PassRef()
	allocs57717e51.Borrow(cgodot_basis_new_allocs)

	var cgodot_basis_new_with_euler_quat_allocs *cgoAllocMap
	ref57717e51.godot_basis_new_with_euler_quat, cgodot_basis_new_with_euler_quat_allocs = x.GodotBasisNewWithEulerQuat.PassRef()
	allocs57717e51.Borrow(cgodot_basis_new_with_euler_quat_allocs)

	var cgodot_basis_get_elements_allocs *cgoAllocMap
	ref57717e51.godot_basis_get_elements, cgodot_basis_get_elements_allocs = x.GodotBasisGetElements.PassRef()
	allocs57717e51.Borrow(cgodot_basis_get_elements_allocs)

	var cgodot_basis_get_axis_allocs *cgoAllocMap
	ref57717e51.godot_basis_get_axis, cgodot_basis_get_axis_allocs = x.GodotBasisGetAxis.PassRef()
	allocs57717e51.Borrow(cgodot_basis_get_axis_allocs)

	var cgodot_basis_set_axis_allocs *cgoAllocMap
	ref57717e51.godot_basis_set_axis, cgodot_basis_set_axis_allocs = x.GodotBasisSetAxis.PassRef()
	allocs57717e51.Borrow(cgodot_basis_set_axis_allocs)

	var cgodot_basis_get_row_allocs *cgoAllocMap
	ref57717e51.godot_basis_get_row, cgodot_basis_get_row_allocs = x.GodotBasisGetRow.PassRef()
	allocs57717e51.Borrow(cgodot_basis_get_row_allocs)

	var cgodot_basis_set_row_allocs *cgoAllocMap
	ref57717e51.godot_basis_set_row, cgodot_basis_set_row_allocs = x.GodotBasisSetRow.PassRef()
	allocs57717e51.Borrow(cgodot_basis_set_row_allocs)

	var cgodot_basis_operator_equal_allocs *cgoAllocMap
	ref57717e51.godot_basis_operator_equal, cgodot_basis_operator_equal_allocs = x.GodotBasisOperatorEqual.PassRef()
	allocs57717e51.Borrow(cgodot_basis_operator_equal_allocs)

	var cgodot_basis_operator_add_allocs *cgoAllocMap
	ref57717e51.godot_basis_operator_add, cgodot_basis_operator_add_allocs = x.GodotBasisOperatorAdd.PassRef()
	allocs57717e51.Borrow(cgodot_basis_operator_add_allocs)

	var cgodot_basis_operator_substract_allocs *cgoAllocMap
	ref57717e51.godot_basis_operator_substract, cgodot_basis_operator_substract_allocs = x.GodotBasisOperatorSubstract.PassRef()
	allocs57717e51.Borrow(cgodot_basis_operator_substract_allocs)

	var cgodot_basis_operator_multiply_vector_allocs *cgoAllocMap
	ref57717e51.godot_basis_operator_multiply_vector, cgodot_basis_operator_multiply_vector_allocs = x.GodotBasisOperatorMultiplyVector.PassRef()
	allocs57717e51.Borrow(cgodot_basis_operator_multiply_vector_allocs)

	var cgodot_basis_operator_multiply_scalar_allocs *cgoAllocMap
	ref57717e51.godot_basis_operator_multiply_scalar, cgodot_basis_operator_multiply_scalar_allocs = x.GodotBasisOperatorMultiplyScalar.PassRef()
	allocs57717e51.Borrow(cgodot_basis_operator_multiply_scalar_allocs)

	var cgodot_vector3_new_allocs *cgoAllocMap
	ref57717e51.godot_vector3_new, cgodot_vector3_new_allocs = x.GodotVector3New.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_new_allocs)

	var cgodot_vector3_as_string_allocs *cgoAllocMap
	ref57717e51.godot_vector3_as_string, cgodot_vector3_as_string_allocs = x.GodotVector3AsString.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_as_string_allocs)

	var cgodot_vector3_min_axis_allocs *cgoAllocMap
	ref57717e51.godot_vector3_min_axis, cgodot_vector3_min_axis_allocs = x.GodotVector3MinAxis.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_min_axis_allocs)

	var cgodot_vector3_max_axis_allocs *cgoAllocMap
	ref57717e51.godot_vector3_max_axis, cgodot_vector3_max_axis_allocs = x.GodotVector3MaxAxis.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_max_axis_allocs)

	var cgodot_vector3_length_allocs *cgoAllocMap
	ref57717e51.godot_vector3_length, cgodot_vector3_length_allocs = x.GodotVector3Length.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_length_allocs)

	var cgodot_vector3_length_squared_allocs *cgoAllocMap
	ref57717e51.godot_vector3_length_squared, cgodot_vector3_length_squared_allocs = x.GodotVector3LengthSquared.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_length_squared_allocs)

	var cgodot_vector3_is_normalized_allocs *cgoAllocMap
	ref57717e51.godot_vector3_is_normalized, cgodot_vector3_is_normalized_allocs = x.GodotVector3IsNormalized.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_is_normalized_allocs)

	var cgodot_vector3_normalized_allocs *cgoAllocMap
	ref57717e51.godot_vector3_normalized, cgodot_vector3_normalized_allocs = x.GodotVector3Normalized.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_normalized_allocs)

	var cgodot_vector3_inverse_allocs *cgoAllocMap
	ref57717e51.godot_vector3_inverse, cgodot_vector3_inverse_allocs = x.GodotVector3Inverse.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_inverse_allocs)

	var cgodot_vector3_snapped_allocs *cgoAllocMap
	ref57717e51.godot_vector3_snapped, cgodot_vector3_snapped_allocs = x.GodotVector3Snapped.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_snapped_allocs)

	var cgodot_vector3_rotated_allocs *cgoAllocMap
	ref57717e51.godot_vector3_rotated, cgodot_vector3_rotated_allocs = x.GodotVector3Rotated.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_rotated_allocs)

	var cgodot_vector3_linear_interpolate_allocs *cgoAllocMap
	ref57717e51.godot_vector3_linear_interpolate, cgodot_vector3_linear_interpolate_allocs = x.GodotVector3LinearInterpolate.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_linear_interpolate_allocs)

	var cgodot_vector3_cubic_interpolate_allocs *cgoAllocMap
	ref57717e51.godot_vector3_cubic_interpolate, cgodot_vector3_cubic_interpolate_allocs = x.GodotVector3CubicInterpolate.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_cubic_interpolate_allocs)

	var cgodot_vector3_dot_allocs *cgoAllocMap
	ref57717e51.godot_vector3_dot, cgodot_vector3_dot_allocs = x.GodotVector3Dot.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_dot_allocs)

	var cgodot_vector3_cross_allocs *cgoAllocMap
	ref57717e51.godot_vector3_cross, cgodot_vector3_cross_allocs = x.GodotVector3Cross.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_cross_allocs)

	var cgodot_vector3_outer_allocs *cgoAllocMap
	ref57717e51.godot_vector3_outer, cgodot_vector3_outer_allocs = x.GodotVector3Outer.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_outer_allocs)

	var cgodot_vector3_to_diagonal_matrix_allocs *cgoAllocMap
	ref57717e51.godot_vector3_to_diagonal_matrix, cgodot_vector3_to_diagonal_matrix_allocs = x.GodotVector3ToDiagonalMatrix.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_to_diagonal_matrix_allocs)

	var cgodot_vector3_abs_allocs *cgoAllocMap
	ref57717e51.godot_vector3_abs, cgodot_vector3_abs_allocs = x.GodotVector3Abs.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_abs_allocs)

	var cgodot_vector3_floor_allocs *cgoAllocMap
	ref57717e51.godot_vector3_floor, cgodot_vector3_floor_allocs = x.GodotVector3Floor.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_floor_allocs)

	var cgodot_vector3_ceil_allocs *cgoAllocMap
	ref57717e51.godot_vector3_ceil, cgodot_vector3_ceil_allocs = x.GodotVector3Ceil.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_ceil_allocs)

	var cgodot_vector3_distance_to_allocs *cgoAllocMap
	ref57717e51.godot_vector3_distance_to, cgodot_vector3_distance_to_allocs = x.GodotVector3DistanceTo.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_distance_to_allocs)

	var cgodot_vector3_distance_squared_to_allocs *cgoAllocMap
	ref57717e51.godot_vector3_distance_squared_to, cgodot_vector3_distance_squared_to_allocs = x.GodotVector3DistanceSquaredTo.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_distance_squared_to_allocs)

	var cgodot_vector3_angle_to_allocs *cgoAllocMap
	ref57717e51.godot_vector3_angle_to, cgodot_vector3_angle_to_allocs = x.GodotVector3AngleTo.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_angle_to_allocs)

	var cgodot_vector3_slide_allocs *cgoAllocMap
	ref57717e51.godot_vector3_slide, cgodot_vector3_slide_allocs = x.GodotVector3Slide.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_slide_allocs)

	var cgodot_vector3_bounce_allocs *cgoAllocMap
	ref57717e51.godot_vector3_bounce, cgodot_vector3_bounce_allocs = x.GodotVector3Bounce.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_bounce_allocs)

	var cgodot_vector3_reflect_allocs *cgoAllocMap
	ref57717e51.godot_vector3_reflect, cgodot_vector3_reflect_allocs = x.GodotVector3Reflect.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_reflect_allocs)

	var cgodot_vector3_operator_add_allocs *cgoAllocMap
	ref57717e51.godot_vector3_operator_add, cgodot_vector3_operator_add_allocs = x.GodotVector3OperatorAdd.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_operator_add_allocs)

	var cgodot_vector3_operator_substract_allocs *cgoAllocMap
	ref57717e51.godot_vector3_operator_substract, cgodot_vector3_operator_substract_allocs = x.GodotVector3OperatorSubstract.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_operator_substract_allocs)

	var cgodot_vector3_operator_multiply_vector_allocs *cgoAllocMap
	ref57717e51.godot_vector3_operator_multiply_vector, cgodot_vector3_operator_multiply_vector_allocs = x.GodotVector3OperatorMultiplyVector.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_operator_multiply_vector_allocs)

	var cgodot_vector3_operator_multiply_scalar_allocs *cgoAllocMap
	ref57717e51.godot_vector3_operator_multiply_scalar, cgodot_vector3_operator_multiply_scalar_allocs = x.GodotVector3OperatorMultiplyScalar.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_operator_multiply_scalar_allocs)

	var cgodot_vector3_operator_divide_vector_allocs *cgoAllocMap
	ref57717e51.godot_vector3_operator_divide_vector, cgodot_vector3_operator_divide_vector_allocs = x.GodotVector3OperatorDivideVector.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_operator_divide_vector_allocs)

	var cgodot_vector3_operator_divide_scalar_allocs *cgoAllocMap
	ref57717e51.godot_vector3_operator_divide_scalar, cgodot_vector3_operator_divide_scalar_allocs = x.GodotVector3OperatorDivideScalar.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_operator_divide_scalar_allocs)

	var cgodot_vector3_operator_equal_allocs *cgoAllocMap
	ref57717e51.godot_vector3_operator_equal, cgodot_vector3_operator_equal_allocs = x.GodotVector3OperatorEqual.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_operator_equal_allocs)

	var cgodot_vector3_operator_less_allocs *cgoAllocMap
	ref57717e51.godot_vector3_operator_less, cgodot_vector3_operator_less_allocs = x.GodotVector3OperatorLess.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_operator_less_allocs)

	var cgodot_vector3_operator_neg_allocs *cgoAllocMap
	ref57717e51.godot_vector3_operator_neg, cgodot_vector3_operator_neg_allocs = x.GodotVector3OperatorNeg.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_operator_neg_allocs)

	var cgodot_vector3_set_axis_allocs *cgoAllocMap
	ref57717e51.godot_vector3_set_axis, cgodot_vector3_set_axis_allocs = x.GodotVector3SetAxis.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_set_axis_allocs)

	var cgodot_vector3_get_axis_allocs *cgoAllocMap
	ref57717e51.godot_vector3_get_axis, cgodot_vector3_get_axis_allocs = x.GodotVector3GetAxis.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_get_axis_allocs)

	var cgodot_pool_byte_array_new_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_new, cgodot_pool_byte_array_new_allocs = x.GodotPoolByteArrayNew.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_new_allocs)

	var cgodot_pool_byte_array_new_copy_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_new_copy, cgodot_pool_byte_array_new_copy_allocs = x.GodotPoolByteArrayNewCopy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_new_copy_allocs)

	var cgodot_pool_byte_array_new_with_array_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_new_with_array, cgodot_pool_byte_array_new_with_array_allocs = x.GodotPoolByteArrayNewWithArray.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_new_with_array_allocs)

	var cgodot_pool_byte_array_append_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_append, cgodot_pool_byte_array_append_allocs = x.GodotPoolByteArrayAppend.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_append_allocs)

	var cgodot_pool_byte_array_append_array_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_append_array, cgodot_pool_byte_array_append_array_allocs = x.GodotPoolByteArrayAppendArray.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_append_array_allocs)

	var cgodot_pool_byte_array_insert_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_insert, cgodot_pool_byte_array_insert_allocs = x.GodotPoolByteArrayInsert.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_insert_allocs)

	var cgodot_pool_byte_array_invert_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_invert, cgodot_pool_byte_array_invert_allocs = x.GodotPoolByteArrayInvert.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_invert_allocs)

	var cgodot_pool_byte_array_push_back_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_push_back, cgodot_pool_byte_array_push_back_allocs = x.GodotPoolByteArrayPushBack.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_push_back_allocs)

	var cgodot_pool_byte_array_remove_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_remove, cgodot_pool_byte_array_remove_allocs = x.GodotPoolByteArrayRemove.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_remove_allocs)

	var cgodot_pool_byte_array_resize_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_resize, cgodot_pool_byte_array_resize_allocs = x.GodotPoolByteArrayResize.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_resize_allocs)

	var cgodot_pool_byte_array_read_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_read, cgodot_pool_byte_array_read_allocs = x.GodotPoolByteArrayRead.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_read_allocs)

	var cgodot_pool_byte_array_write_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_write, cgodot_pool_byte_array_write_allocs = x.GodotPoolByteArrayWrite.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_write_allocs)

	var cgodot_pool_byte_array_set_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_set, cgodot_pool_byte_array_set_allocs = x.GodotPoolByteArraySet.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_set_allocs)

	var cgodot_pool_byte_array_get_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_get, cgodot_pool_byte_array_get_allocs = x.GodotPoolByteArrayGet.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_get_allocs)

	var cgodot_pool_byte_array_size_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_size, cgodot_pool_byte_array_size_allocs = x.GodotPoolByteArraySize.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_size_allocs)

	var cgodot_pool_byte_array_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_destroy, cgodot_pool_byte_array_destroy_allocs = x.GodotPoolByteArrayDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_destroy_allocs)

	var cgodot_pool_int_array_new_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_new, cgodot_pool_int_array_new_allocs = x.GodotPoolIntArrayNew.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_new_allocs)

	var cgodot_pool_int_array_new_copy_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_new_copy, cgodot_pool_int_array_new_copy_allocs = x.GodotPoolIntArrayNewCopy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_new_copy_allocs)

	var cgodot_pool_int_array_new_with_array_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_new_with_array, cgodot_pool_int_array_new_with_array_allocs = x.GodotPoolIntArrayNewWithArray.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_new_with_array_allocs)

	var cgodot_pool_int_array_append_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_append, cgodot_pool_int_array_append_allocs = x.GodotPoolIntArrayAppend.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_append_allocs)

	var cgodot_pool_int_array_append_array_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_append_array, cgodot_pool_int_array_append_array_allocs = x.GodotPoolIntArrayAppendArray.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_append_array_allocs)

	var cgodot_pool_int_array_insert_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_insert, cgodot_pool_int_array_insert_allocs = x.GodotPoolIntArrayInsert.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_insert_allocs)

	var cgodot_pool_int_array_invert_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_invert, cgodot_pool_int_array_invert_allocs = x.GodotPoolIntArrayInvert.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_invert_allocs)

	var cgodot_pool_int_array_push_back_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_push_back, cgodot_pool_int_array_push_back_allocs = x.GodotPoolIntArrayPushBack.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_push_back_allocs)

	var cgodot_pool_int_array_remove_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_remove, cgodot_pool_int_array_remove_allocs = x.GodotPoolIntArrayRemove.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_remove_allocs)

	var cgodot_pool_int_array_resize_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_resize, cgodot_pool_int_array_resize_allocs = x.GodotPoolIntArrayResize.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_resize_allocs)

	var cgodot_pool_int_array_read_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_read, cgodot_pool_int_array_read_allocs = x.GodotPoolIntArrayRead.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_read_allocs)

	var cgodot_pool_int_array_write_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_write, cgodot_pool_int_array_write_allocs = x.GodotPoolIntArrayWrite.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_write_allocs)

	var cgodot_pool_int_array_set_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_set, cgodot_pool_int_array_set_allocs = x.GodotPoolIntArraySet.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_set_allocs)

	var cgodot_pool_int_array_get_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_get, cgodot_pool_int_array_get_allocs = x.GodotPoolIntArrayGet.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_get_allocs)

	var cgodot_pool_int_array_size_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_size, cgodot_pool_int_array_size_allocs = x.GodotPoolIntArraySize.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_size_allocs)

	var cgodot_pool_int_array_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_destroy, cgodot_pool_int_array_destroy_allocs = x.GodotPoolIntArrayDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_destroy_allocs)

	var cgodot_pool_real_array_new_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_new, cgodot_pool_real_array_new_allocs = x.GodotPoolRealArrayNew.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_new_allocs)

	var cgodot_pool_real_array_new_copy_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_new_copy, cgodot_pool_real_array_new_copy_allocs = x.GodotPoolRealArrayNewCopy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_new_copy_allocs)

	var cgodot_pool_real_array_new_with_array_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_new_with_array, cgodot_pool_real_array_new_with_array_allocs = x.GodotPoolRealArrayNewWithArray.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_new_with_array_allocs)

	var cgodot_pool_real_array_append_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_append, cgodot_pool_real_array_append_allocs = x.GodotPoolRealArrayAppend.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_append_allocs)

	var cgodot_pool_real_array_append_array_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_append_array, cgodot_pool_real_array_append_array_allocs = x.GodotPoolRealArrayAppendArray.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_append_array_allocs)

	var cgodot_pool_real_array_insert_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_insert, cgodot_pool_real_array_insert_allocs = x.GodotPoolRealArrayInsert.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_insert_allocs)

	var cgodot_pool_real_array_invert_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_invert, cgodot_pool_real_array_invert_allocs = x.GodotPoolRealArrayInvert.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_invert_allocs)

	var cgodot_pool_real_array_push_back_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_push_back, cgodot_pool_real_array_push_back_allocs = x.GodotPoolRealArrayPushBack.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_push_back_allocs)

	var cgodot_pool_real_array_remove_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_remove, cgodot_pool_real_array_remove_allocs = x.GodotPoolRealArrayRemove.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_remove_allocs)

	var cgodot_pool_real_array_resize_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_resize, cgodot_pool_real_array_resize_allocs = x.GodotPoolRealArrayResize.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_resize_allocs)

	var cgodot_pool_real_array_read_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_read, cgodot_pool_real_array_read_allocs = x.GodotPoolRealArrayRead.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_read_allocs)

	var cgodot_pool_real_array_write_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_write, cgodot_pool_real_array_write_allocs = x.GodotPoolRealArrayWrite.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_write_allocs)

	var cgodot_pool_real_array_set_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_set, cgodot_pool_real_array_set_allocs = x.GodotPoolRealArraySet.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_set_allocs)

	var cgodot_pool_real_array_get_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_get, cgodot_pool_real_array_get_allocs = x.GodotPoolRealArrayGet.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_get_allocs)

	var cgodot_pool_real_array_size_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_size, cgodot_pool_real_array_size_allocs = x.GodotPoolRealArraySize.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_size_allocs)

	var cgodot_pool_real_array_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_destroy, cgodot_pool_real_array_destroy_allocs = x.GodotPoolRealArrayDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_destroy_allocs)

	var cgodot_pool_string_array_new_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_new, cgodot_pool_string_array_new_allocs = x.GodotPoolStringArrayNew.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_new_allocs)

	var cgodot_pool_string_array_new_copy_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_new_copy, cgodot_pool_string_array_new_copy_allocs = x.GodotPoolStringArrayNewCopy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_new_copy_allocs)

	var cgodot_pool_string_array_new_with_array_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_new_with_array, cgodot_pool_string_array_new_with_array_allocs = x.GodotPoolStringArrayNewWithArray.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_new_with_array_allocs)

	var cgodot_pool_string_array_append_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_append, cgodot_pool_string_array_append_allocs = x.GodotPoolStringArrayAppend.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_append_allocs)

	var cgodot_pool_string_array_append_array_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_append_array, cgodot_pool_string_array_append_array_allocs = x.GodotPoolStringArrayAppendArray.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_append_array_allocs)

	var cgodot_pool_string_array_insert_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_insert, cgodot_pool_string_array_insert_allocs = x.GodotPoolStringArrayInsert.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_insert_allocs)

	var cgodot_pool_string_array_invert_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_invert, cgodot_pool_string_array_invert_allocs = x.GodotPoolStringArrayInvert.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_invert_allocs)

	var cgodot_pool_string_array_push_back_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_push_back, cgodot_pool_string_array_push_back_allocs = x.GodotPoolStringArrayPushBack.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_push_back_allocs)

	var cgodot_pool_string_array_remove_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_remove, cgodot_pool_string_array_remove_allocs = x.GodotPoolStringArrayRemove.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_remove_allocs)

	var cgodot_pool_string_array_resize_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_resize, cgodot_pool_string_array_resize_allocs = x.GodotPoolStringArrayResize.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_resize_allocs)

	var cgodot_pool_string_array_read_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_read, cgodot_pool_string_array_read_allocs = x.GodotPoolStringArrayRead.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_read_allocs)

	var cgodot_pool_string_array_write_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_write, cgodot_pool_string_array_write_allocs = x.GodotPoolStringArrayWrite.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_write_allocs)

	var cgodot_pool_string_array_set_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_set, cgodot_pool_string_array_set_allocs = x.GodotPoolStringArraySet.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_set_allocs)

	var cgodot_pool_string_array_get_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_get, cgodot_pool_string_array_get_allocs = x.GodotPoolStringArrayGet.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_get_allocs)

	var cgodot_pool_string_array_size_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_size, cgodot_pool_string_array_size_allocs = x.GodotPoolStringArraySize.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_size_allocs)

	var cgodot_pool_string_array_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_destroy, cgodot_pool_string_array_destroy_allocs = x.GodotPoolStringArrayDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_destroy_allocs)

	var cgodot_pool_vector2_array_new_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_new, cgodot_pool_vector2_array_new_allocs = x.GodotPoolVector2ArrayNew.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_new_allocs)

	var cgodot_pool_vector2_array_new_copy_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_new_copy, cgodot_pool_vector2_array_new_copy_allocs = x.GodotPoolVector2ArrayNewCopy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_new_copy_allocs)

	var cgodot_pool_vector2_array_new_with_array_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_new_with_array, cgodot_pool_vector2_array_new_with_array_allocs = x.GodotPoolVector2ArrayNewWithArray.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_new_with_array_allocs)

	var cgodot_pool_vector2_array_append_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_append, cgodot_pool_vector2_array_append_allocs = x.GodotPoolVector2ArrayAppend.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_append_allocs)

	var cgodot_pool_vector2_array_append_array_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_append_array, cgodot_pool_vector2_array_append_array_allocs = x.GodotPoolVector2ArrayAppendArray.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_append_array_allocs)

	var cgodot_pool_vector2_array_insert_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_insert, cgodot_pool_vector2_array_insert_allocs = x.GodotPoolVector2ArrayInsert.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_insert_allocs)

	var cgodot_pool_vector2_array_invert_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_invert, cgodot_pool_vector2_array_invert_allocs = x.GodotPoolVector2ArrayInvert.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_invert_allocs)

	var cgodot_pool_vector2_array_push_back_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_push_back, cgodot_pool_vector2_array_push_back_allocs = x.GodotPoolVector2ArrayPushBack.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_push_back_allocs)

	var cgodot_pool_vector2_array_remove_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_remove, cgodot_pool_vector2_array_remove_allocs = x.GodotPoolVector2ArrayRemove.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_remove_allocs)

	var cgodot_pool_vector2_array_resize_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_resize, cgodot_pool_vector2_array_resize_allocs = x.GodotPoolVector2ArrayResize.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_resize_allocs)

	var cgodot_pool_vector2_array_read_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_read, cgodot_pool_vector2_array_read_allocs = x.GodotPoolVector2ArrayRead.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_read_allocs)

	var cgodot_pool_vector2_array_write_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_write, cgodot_pool_vector2_array_write_allocs = x.GodotPoolVector2ArrayWrite.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_write_allocs)

	var cgodot_pool_vector2_array_set_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_set, cgodot_pool_vector2_array_set_allocs = x.GodotPoolVector2ArraySet.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_set_allocs)

	var cgodot_pool_vector2_array_get_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_get, cgodot_pool_vector2_array_get_allocs = x.GodotPoolVector2ArrayGet.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_get_allocs)

	var cgodot_pool_vector2_array_size_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_size, cgodot_pool_vector2_array_size_allocs = x.GodotPoolVector2ArraySize.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_size_allocs)

	var cgodot_pool_vector2_array_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_destroy, cgodot_pool_vector2_array_destroy_allocs = x.GodotPoolVector2ArrayDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_destroy_allocs)

	var cgodot_pool_vector3_array_new_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_new, cgodot_pool_vector3_array_new_allocs = x.GodotPoolVector3ArrayNew.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_new_allocs)

	var cgodot_pool_vector3_array_new_copy_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_new_copy, cgodot_pool_vector3_array_new_copy_allocs = x.GodotPoolVector3ArrayNewCopy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_new_copy_allocs)

	var cgodot_pool_vector3_array_new_with_array_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_new_with_array, cgodot_pool_vector3_array_new_with_array_allocs = x.GodotPoolVector3ArrayNewWithArray.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_new_with_array_allocs)

	var cgodot_pool_vector3_array_append_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_append, cgodot_pool_vector3_array_append_allocs = x.GodotPoolVector3ArrayAppend.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_append_allocs)

	var cgodot_pool_vector3_array_append_array_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_append_array, cgodot_pool_vector3_array_append_array_allocs = x.GodotPoolVector3ArrayAppendArray.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_append_array_allocs)

	var cgodot_pool_vector3_array_insert_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_insert, cgodot_pool_vector3_array_insert_allocs = x.GodotPoolVector3ArrayInsert.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_insert_allocs)

	var cgodot_pool_vector3_array_invert_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_invert, cgodot_pool_vector3_array_invert_allocs = x.GodotPoolVector3ArrayInvert.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_invert_allocs)

	var cgodot_pool_vector3_array_push_back_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_push_back, cgodot_pool_vector3_array_push_back_allocs = x.GodotPoolVector3ArrayPushBack.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_push_back_allocs)

	var cgodot_pool_vector3_array_remove_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_remove, cgodot_pool_vector3_array_remove_allocs = x.GodotPoolVector3ArrayRemove.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_remove_allocs)

	var cgodot_pool_vector3_array_resize_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_resize, cgodot_pool_vector3_array_resize_allocs = x.GodotPoolVector3ArrayResize.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_resize_allocs)

	var cgodot_pool_vector3_array_read_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_read, cgodot_pool_vector3_array_read_allocs = x.GodotPoolVector3ArrayRead.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_read_allocs)

	var cgodot_pool_vector3_array_write_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_write, cgodot_pool_vector3_array_write_allocs = x.GodotPoolVector3ArrayWrite.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_write_allocs)

	var cgodot_pool_vector3_array_set_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_set, cgodot_pool_vector3_array_set_allocs = x.GodotPoolVector3ArraySet.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_set_allocs)

	var cgodot_pool_vector3_array_get_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_get, cgodot_pool_vector3_array_get_allocs = x.GodotPoolVector3ArrayGet.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_get_allocs)

	var cgodot_pool_vector3_array_size_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_size, cgodot_pool_vector3_array_size_allocs = x.GodotPoolVector3ArraySize.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_size_allocs)

	var cgodot_pool_vector3_array_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_destroy, cgodot_pool_vector3_array_destroy_allocs = x.GodotPoolVector3ArrayDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_destroy_allocs)

	var cgodot_pool_color_array_new_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_new, cgodot_pool_color_array_new_allocs = x.GodotPoolColorArrayNew.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_new_allocs)

	var cgodot_pool_color_array_new_copy_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_new_copy, cgodot_pool_color_array_new_copy_allocs = x.GodotPoolColorArrayNewCopy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_new_copy_allocs)

	var cgodot_pool_color_array_new_with_array_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_new_with_array, cgodot_pool_color_array_new_with_array_allocs = x.GodotPoolColorArrayNewWithArray.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_new_with_array_allocs)

	var cgodot_pool_color_array_append_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_append, cgodot_pool_color_array_append_allocs = x.GodotPoolColorArrayAppend.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_append_allocs)

	var cgodot_pool_color_array_append_array_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_append_array, cgodot_pool_color_array_append_array_allocs = x.GodotPoolColorArrayAppendArray.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_append_array_allocs)

	var cgodot_pool_color_array_insert_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_insert, cgodot_pool_color_array_insert_allocs = x.GodotPoolColorArrayInsert.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_insert_allocs)

	var cgodot_pool_color_array_invert_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_invert, cgodot_pool_color_array_invert_allocs = x.GodotPoolColorArrayInvert.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_invert_allocs)

	var cgodot_pool_color_array_push_back_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_push_back, cgodot_pool_color_array_push_back_allocs = x.GodotPoolColorArrayPushBack.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_push_back_allocs)

	var cgodot_pool_color_array_remove_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_remove, cgodot_pool_color_array_remove_allocs = x.GodotPoolColorArrayRemove.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_remove_allocs)

	var cgodot_pool_color_array_resize_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_resize, cgodot_pool_color_array_resize_allocs = x.GodotPoolColorArrayResize.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_resize_allocs)

	var cgodot_pool_color_array_read_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_read, cgodot_pool_color_array_read_allocs = x.GodotPoolColorArrayRead.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_read_allocs)

	var cgodot_pool_color_array_write_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_write, cgodot_pool_color_array_write_allocs = x.GodotPoolColorArrayWrite.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_write_allocs)

	var cgodot_pool_color_array_set_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_set, cgodot_pool_color_array_set_allocs = x.GodotPoolColorArraySet.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_set_allocs)

	var cgodot_pool_color_array_get_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_get, cgodot_pool_color_array_get_allocs = x.GodotPoolColorArrayGet.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_get_allocs)

	var cgodot_pool_color_array_size_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_size, cgodot_pool_color_array_size_allocs = x.GodotPoolColorArraySize.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_size_allocs)

	var cgodot_pool_color_array_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_destroy, cgodot_pool_color_array_destroy_allocs = x.GodotPoolColorArrayDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_destroy_allocs)

	var cgodot_pool_byte_array_read_access_ptr_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_read_access_ptr, cgodot_pool_byte_array_read_access_ptr_allocs = x.GodotPoolByteArrayReadAccessPtr.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_read_access_ptr_allocs)

	var cgodot_pool_byte_array_read_access_operator_assign_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_read_access_operator_assign, cgodot_pool_byte_array_read_access_operator_assign_allocs = x.GodotPoolByteArrayReadAccessOperatorAssign.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_read_access_operator_assign_allocs)

	var cgodot_pool_byte_array_read_access_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_read_access_destroy, cgodot_pool_byte_array_read_access_destroy_allocs = x.GodotPoolByteArrayReadAccessDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_read_access_destroy_allocs)

	var cgodot_pool_int_array_read_access_ptr_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_read_access_ptr, cgodot_pool_int_array_read_access_ptr_allocs = x.GodotPoolIntArrayReadAccessPtr.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_read_access_ptr_allocs)

	var cgodot_pool_int_array_read_access_operator_assign_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_read_access_operator_assign, cgodot_pool_int_array_read_access_operator_assign_allocs = x.GodotPoolIntArrayReadAccessOperatorAssign.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_read_access_operator_assign_allocs)

	var cgodot_pool_int_array_read_access_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_read_access_destroy, cgodot_pool_int_array_read_access_destroy_allocs = x.GodotPoolIntArrayReadAccessDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_read_access_destroy_allocs)

	var cgodot_pool_real_array_read_access_ptr_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_read_access_ptr, cgodot_pool_real_array_read_access_ptr_allocs = x.GodotPoolRealArrayReadAccessPtr.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_read_access_ptr_allocs)

	var cgodot_pool_real_array_read_access_operator_assign_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_read_access_operator_assign, cgodot_pool_real_array_read_access_operator_assign_allocs = x.GodotPoolRealArrayReadAccessOperatorAssign.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_read_access_operator_assign_allocs)

	var cgodot_pool_real_array_read_access_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_read_access_destroy, cgodot_pool_real_array_read_access_destroy_allocs = x.GodotPoolRealArrayReadAccessDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_read_access_destroy_allocs)

	var cgodot_pool_string_array_read_access_ptr_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_read_access_ptr, cgodot_pool_string_array_read_access_ptr_allocs = x.GodotPoolStringArrayReadAccessPtr.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_read_access_ptr_allocs)

	var cgodot_pool_string_array_read_access_operator_assign_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_read_access_operator_assign, cgodot_pool_string_array_read_access_operator_assign_allocs = x.GodotPoolStringArrayReadAccessOperatorAssign.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_read_access_operator_assign_allocs)

	var cgodot_pool_string_array_read_access_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_read_access_destroy, cgodot_pool_string_array_read_access_destroy_allocs = x.GodotPoolStringArrayReadAccessDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_read_access_destroy_allocs)

	var cgodot_pool_vector2_array_read_access_ptr_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_read_access_ptr, cgodot_pool_vector2_array_read_access_ptr_allocs = x.GodotPoolVector2ArrayReadAccessPtr.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_read_access_ptr_allocs)

	var cgodot_pool_vector2_array_read_access_operator_assign_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_read_access_operator_assign, cgodot_pool_vector2_array_read_access_operator_assign_allocs = x.GodotPoolVector2ArrayReadAccessOperatorAssign.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_read_access_operator_assign_allocs)

	var cgodot_pool_vector2_array_read_access_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_read_access_destroy, cgodot_pool_vector2_array_read_access_destroy_allocs = x.GodotPoolVector2ArrayReadAccessDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_read_access_destroy_allocs)

	var cgodot_pool_vector3_array_read_access_ptr_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_read_access_ptr, cgodot_pool_vector3_array_read_access_ptr_allocs = x.GodotPoolVector3ArrayReadAccessPtr.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_read_access_ptr_allocs)

	var cgodot_pool_vector3_array_read_access_operator_assign_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_read_access_operator_assign, cgodot_pool_vector3_array_read_access_operator_assign_allocs = x.GodotPoolVector3ArrayReadAccessOperatorAssign.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_read_access_operator_assign_allocs)

	var cgodot_pool_vector3_array_read_access_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_read_access_destroy, cgodot_pool_vector3_array_read_access_destroy_allocs = x.GodotPoolVector3ArrayReadAccessDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_read_access_destroy_allocs)

	var cgodot_pool_color_array_read_access_ptr_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_read_access_ptr, cgodot_pool_color_array_read_access_ptr_allocs = x.GodotPoolColorArrayReadAccessPtr.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_read_access_ptr_allocs)

	var cgodot_pool_color_array_read_access_operator_assign_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_read_access_operator_assign, cgodot_pool_color_array_read_access_operator_assign_allocs = x.GodotPoolColorArrayReadAccessOperatorAssign.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_read_access_operator_assign_allocs)

	var cgodot_pool_color_array_read_access_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_read_access_destroy, cgodot_pool_color_array_read_access_destroy_allocs = x.GodotPoolColorArrayReadAccessDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_read_access_destroy_allocs)

	var cgodot_pool_byte_array_write_access_ptr_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_write_access_ptr, cgodot_pool_byte_array_write_access_ptr_allocs = x.GodotPoolByteArrayWriteAccessPtr.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_write_access_ptr_allocs)

	var cgodot_pool_byte_array_write_access_operator_assign_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_write_access_operator_assign, cgodot_pool_byte_array_write_access_operator_assign_allocs = x.GodotPoolByteArrayWriteAccessOperatorAssign.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_write_access_operator_assign_allocs)

	var cgodot_pool_byte_array_write_access_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_write_access_destroy, cgodot_pool_byte_array_write_access_destroy_allocs = x.GodotPoolByteArrayWriteAccessDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_write_access_destroy_allocs)

	var cgodot_pool_int_array_write_access_ptr_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_write_access_ptr, cgodot_pool_int_array_write_access_ptr_allocs = x.GodotPoolIntArrayWriteAccessPtr.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_write_access_ptr_allocs)

	var cgodot_pool_int_array_write_access_operator_assign_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_write_access_operator_assign, cgodot_pool_int_array_write_access_operator_assign_allocs = x.GodotPoolIntArrayWriteAccessOperatorAssign.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_write_access_operator_assign_allocs)

	var cgodot_pool_int_array_write_access_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_write_access_destroy, cgodot_pool_int_array_write_access_destroy_allocs = x.GodotPoolIntArrayWriteAccessDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_write_access_destroy_allocs)

	var cgodot_pool_real_array_write_access_ptr_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_write_access_ptr, cgodot_pool_real_array_write_access_ptr_allocs = x.GodotPoolRealArrayWriteAccessPtr.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_write_access_ptr_allocs)

	var cgodot_pool_real_array_write_access_operator_assign_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_write_access_operator_assign, cgodot_pool_real_array_write_access_operator_assign_allocs = x.GodotPoolRealArrayWriteAccessOperatorAssign.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_write_access_operator_assign_allocs)

	var cgodot_pool_real_array_write_access_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_write_access_destroy, cgodot_pool_real_array_write_access_destroy_allocs = x.GodotPoolRealArrayWriteAccessDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_write_access_destroy_allocs)

	var cgodot_pool_string_array_write_access_ptr_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_write_access_ptr, cgodot_pool_string_array_write_access_ptr_allocs = x.GodotPoolStringArrayWriteAccessPtr.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_write_access_ptr_allocs)

	var cgodot_pool_string_array_write_access_operator_assign_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_write_access_operator_assign, cgodot_pool_string_array_write_access_operator_assign_allocs = x.GodotPoolStringArrayWriteAccessOperatorAssign.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_write_access_operator_assign_allocs)

	var cgodot_pool_string_array_write_access_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_write_access_destroy, cgodot_pool_string_array_write_access_destroy_allocs = x.GodotPoolStringArrayWriteAccessDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_write_access_destroy_allocs)

	var cgodot_pool_vector2_array_write_access_ptr_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_write_access_ptr, cgodot_pool_vector2_array_write_access_ptr_allocs = x.GodotPoolVector2ArrayWriteAccessPtr.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_write_access_ptr_allocs)

	var cgodot_pool_vector2_array_write_access_operator_assign_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_write_access_operator_assign, cgodot_pool_vector2_array_write_access_operator_assign_allocs = x.GodotPoolVector2ArrayWriteAccessOperatorAssign.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_write_access_operator_assign_allocs)

	var cgodot_pool_vector2_array_write_access_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_write_access_destroy, cgodot_pool_vector2_array_write_access_destroy_allocs = x.GodotPoolVector2ArrayWriteAccessDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_write_access_destroy_allocs)

	var cgodot_pool_vector3_array_write_access_ptr_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_write_access_ptr, cgodot_pool_vector3_array_write_access_ptr_allocs = x.GodotPoolVector3ArrayWriteAccessPtr.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_write_access_ptr_allocs)

	var cgodot_pool_vector3_array_write_access_operator_assign_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_write_access_operator_assign, cgodot_pool_vector3_array_write_access_operator_assign_allocs = x.GodotPoolVector3ArrayWriteAccessOperatorAssign.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_write_access_operator_assign_allocs)

	var cgodot_pool_vector3_array_write_access_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_write_access_destroy, cgodot_pool_vector3_array_write_access_destroy_allocs = x.GodotPoolVector3ArrayWriteAccessDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_write_access_destroy_allocs)

	var cgodot_pool_color_array_write_access_ptr_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_write_access_ptr, cgodot_pool_color_array_write_access_ptr_allocs = x.GodotPoolColorArrayWriteAccessPtr.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_write_access_ptr_allocs)

	var cgodot_pool_color_array_write_access_operator_assign_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_write_access_operator_assign, cgodot_pool_color_array_write_access_operator_assign_allocs = x.GodotPoolColorArrayWriteAccessOperatorAssign.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_write_access_operator_assign_allocs)

	var cgodot_pool_color_array_write_access_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_write_access_destroy, cgodot_pool_color_array_write_access_destroy_allocs = x.GodotPoolColorArrayWriteAccessDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_write_access_destroy_allocs)

	var cgodot_array_new_allocs *cgoAllocMap
	ref57717e51.godot_array_new, cgodot_array_new_allocs = x.GodotArrayNew.PassRef()
	allocs57717e51.Borrow(cgodot_array_new_allocs)

	var cgodot_array_new_copy_allocs *cgoAllocMap
	ref57717e51.godot_array_new_copy, cgodot_array_new_copy_allocs = x.GodotArrayNewCopy.PassRef()
	allocs57717e51.Borrow(cgodot_array_new_copy_allocs)

	var cgodot_array_new_pool_color_array_allocs *cgoAllocMap
	ref57717e51.godot_array_new_pool_color_array, cgodot_array_new_pool_color_array_allocs = x.GodotArrayNewPoolColorArray.PassRef()
	allocs57717e51.Borrow(cgodot_array_new_pool_color_array_allocs)

	var cgodot_array_new_pool_vector3_array_allocs *cgoAllocMap
	ref57717e51.godot_array_new_pool_vector3_array, cgodot_array_new_pool_vector3_array_allocs = x.GodotArrayNewPoolVector3Array.PassRef()
	allocs57717e51.Borrow(cgodot_array_new_pool_vector3_array_allocs)

	var cgodot_array_new_pool_vector2_array_allocs *cgoAllocMap
	ref57717e51.godot_array_new_pool_vector2_array, cgodot_array_new_pool_vector2_array_allocs = x.GodotArrayNewPoolVector2Array.PassRef()
	allocs57717e51.Borrow(cgodot_array_new_pool_vector2_array_allocs)

	var cgodot_array_new_pool_string_array_allocs *cgoAllocMap
	ref57717e51.godot_array_new_pool_string_array, cgodot_array_new_pool_string_array_allocs = x.GodotArrayNewPoolStringArray.PassRef()
	allocs57717e51.Borrow(cgodot_array_new_pool_string_array_allocs)

	var cgodot_array_new_pool_real_array_allocs *cgoAllocMap
	ref57717e51.godot_array_new_pool_real_array, cgodot_array_new_pool_real_array_allocs = x.GodotArrayNewPoolRealArray.PassRef()
	allocs57717e51.Borrow(cgodot_array_new_pool_real_array_allocs)

	var cgodot_array_new_pool_int_array_allocs *cgoAllocMap
	ref57717e51.godot_array_new_pool_int_array, cgodot_array_new_pool_int_array_allocs = x.GodotArrayNewPoolIntArray.PassRef()
	allocs57717e51.Borrow(cgodot_array_new_pool_int_array_allocs)

	var cgodot_array_new_pool_byte_array_allocs *cgoAllocMap
	ref57717e51.godot_array_new_pool_byte_array, cgodot_array_new_pool_byte_array_allocs = x.GodotArrayNewPoolByteArray.PassRef()
	allocs57717e51.Borrow(cgodot_array_new_pool_byte_array_allocs)

	var cgodot_array_set_allocs *cgoAllocMap
	ref57717e51.godot_array_set, cgodot_array_set_allocs = x.GodotArraySet.PassRef()
	allocs57717e51.Borrow(cgodot_array_set_allocs)

	var cgodot_array_get_allocs *cgoAllocMap
	ref57717e51.godot_array_get, cgodot_array_get_allocs = x.GodotArrayGet.PassRef()
	allocs57717e51.Borrow(cgodot_array_get_allocs)

	var cgodot_array_operator_index_allocs *cgoAllocMap
	ref57717e51.godot_array_operator_index, cgodot_array_operator_index_allocs = x.GodotArrayOperatorIndex.PassRef()
	allocs57717e51.Borrow(cgodot_array_operator_index_allocs)

	var cgodot_array_operator_index_const_allocs *cgoAllocMap
	ref57717e51.godot_array_operator_index_const, cgodot_array_operator_index_const_allocs = x.GodotArrayOperatorIndexConst.PassRef()
	allocs57717e51.Borrow(cgodot_array_operator_index_const_allocs)

	var cgodot_array_append_allocs *cgoAllocMap
	ref57717e51.godot_array_append, cgodot_array_append_allocs = x.GodotArrayAppend.PassRef()
	allocs57717e51.Borrow(cgodot_array_append_allocs)

	var cgodot_array_clear_allocs *cgoAllocMap
	ref57717e51.godot_array_clear, cgodot_array_clear_allocs = x.GodotArrayClear.PassRef()
	allocs57717e51.Borrow(cgodot_array_clear_allocs)

	var cgodot_array_count_allocs *cgoAllocMap
	ref57717e51.godot_array_count, cgodot_array_count_allocs = x.GodotArrayCount.PassRef()
	allocs57717e51.Borrow(cgodot_array_count_allocs)

	var cgodot_array_empty_allocs *cgoAllocMap
	ref57717e51.godot_array_empty, cgodot_array_empty_allocs = x.GodotArrayEmpty.PassRef()
	allocs57717e51.Borrow(cgodot_array_empty_allocs)

	var cgodot_array_erase_allocs *cgoAllocMap
	ref57717e51.godot_array_erase, cgodot_array_erase_allocs = x.GodotArrayErase.PassRef()
	allocs57717e51.Borrow(cgodot_array_erase_allocs)

	var cgodot_array_front_allocs *cgoAllocMap
	ref57717e51.godot_array_front, cgodot_array_front_allocs = x.GodotArrayFront.PassRef()
	allocs57717e51.Borrow(cgodot_array_front_allocs)

	var cgodot_array_back_allocs *cgoAllocMap
	ref57717e51.godot_array_back, cgodot_array_back_allocs = x.GodotArrayBack.PassRef()
	allocs57717e51.Borrow(cgodot_array_back_allocs)

	var cgodot_array_find_allocs *cgoAllocMap
	ref57717e51.godot_array_find, cgodot_array_find_allocs = x.GodotArrayFind.PassRef()
	allocs57717e51.Borrow(cgodot_array_find_allocs)

	var cgodot_array_find_last_allocs *cgoAllocMap
	ref57717e51.godot_array_find_last, cgodot_array_find_last_allocs = x.GodotArrayFindLast.PassRef()
	allocs57717e51.Borrow(cgodot_array_find_last_allocs)

	var cgodot_array_has_allocs *cgoAllocMap
	ref57717e51.godot_array_has, cgodot_array_has_allocs = x.GodotArrayHas.PassRef()
	allocs57717e51.Borrow(cgodot_array_has_allocs)

	var cgodot_array_hash_allocs *cgoAllocMap
	ref57717e51.godot_array_hash, cgodot_array_hash_allocs = x.GodotArrayHash.PassRef()
	allocs57717e51.Borrow(cgodot_array_hash_allocs)

	var cgodot_array_insert_allocs *cgoAllocMap
	ref57717e51.godot_array_insert, cgodot_array_insert_allocs = x.GodotArrayInsert.PassRef()
	allocs57717e51.Borrow(cgodot_array_insert_allocs)

	var cgodot_array_invert_allocs *cgoAllocMap
	ref57717e51.godot_array_invert, cgodot_array_invert_allocs = x.GodotArrayInvert.PassRef()
	allocs57717e51.Borrow(cgodot_array_invert_allocs)

	var cgodot_array_pop_back_allocs *cgoAllocMap
	ref57717e51.godot_array_pop_back, cgodot_array_pop_back_allocs = x.GodotArrayPopBack.PassRef()
	allocs57717e51.Borrow(cgodot_array_pop_back_allocs)

	var cgodot_array_pop_front_allocs *cgoAllocMap
	ref57717e51.godot_array_pop_front, cgodot_array_pop_front_allocs = x.GodotArrayPopFront.PassRef()
	allocs57717e51.Borrow(cgodot_array_pop_front_allocs)

	var cgodot_array_push_back_allocs *cgoAllocMap
	ref57717e51.godot_array_push_back, cgodot_array_push_back_allocs = x.GodotArrayPushBack.PassRef()
	allocs57717e51.Borrow(cgodot_array_push_back_allocs)

	var cgodot_array_push_front_allocs *cgoAllocMap
	ref57717e51.godot_array_push_front, cgodot_array_push_front_allocs = x.GodotArrayPushFront.PassRef()
	allocs57717e51.Borrow(cgodot_array_push_front_allocs)

	var cgodot_array_remove_allocs *cgoAllocMap
	ref57717e51.godot_array_remove, cgodot_array_remove_allocs = x.GodotArrayRemove.PassRef()
	allocs57717e51.Borrow(cgodot_array_remove_allocs)

	var cgodot_array_resize_allocs *cgoAllocMap
	ref57717e51.godot_array_resize, cgodot_array_resize_allocs = x.GodotArrayResize.PassRef()
	allocs57717e51.Borrow(cgodot_array_resize_allocs)

	var cgodot_array_rfind_allocs *cgoAllocMap
	ref57717e51.godot_array_rfind, cgodot_array_rfind_allocs = x.GodotArrayRfind.PassRef()
	allocs57717e51.Borrow(cgodot_array_rfind_allocs)

	var cgodot_array_size_allocs *cgoAllocMap
	ref57717e51.godot_array_size, cgodot_array_size_allocs = x.GodotArraySize.PassRef()
	allocs57717e51.Borrow(cgodot_array_size_allocs)

	var cgodot_array_sort_allocs *cgoAllocMap
	ref57717e51.godot_array_sort, cgodot_array_sort_allocs = x.GodotArraySort.PassRef()
	allocs57717e51.Borrow(cgodot_array_sort_allocs)

	var cgodot_array_sort_custom_allocs *cgoAllocMap
	ref57717e51.godot_array_sort_custom, cgodot_array_sort_custom_allocs = x.GodotArraySortCustom.PassRef()
	allocs57717e51.Borrow(cgodot_array_sort_custom_allocs)

	var cgodot_array_bsearch_allocs *cgoAllocMap
	ref57717e51.godot_array_bsearch, cgodot_array_bsearch_allocs = x.GodotArrayBsearch.PassRef()
	allocs57717e51.Borrow(cgodot_array_bsearch_allocs)

	var cgodot_array_bsearch_custom_allocs *cgoAllocMap
	ref57717e51.godot_array_bsearch_custom, cgodot_array_bsearch_custom_allocs = x.GodotArrayBsearchCustom.PassRef()
	allocs57717e51.Borrow(cgodot_array_bsearch_custom_allocs)

	var cgodot_array_destroy_allocs *cgoAllocMap
	ref57717e51.godot_array_destroy, cgodot_array_destroy_allocs = x.GodotArrayDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_array_destroy_allocs)

	var cgodot_dictionary_new_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_new, cgodot_dictionary_new_allocs = x.GodotDictionaryNew.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_new_allocs)

	var cgodot_dictionary_new_copy_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_new_copy, cgodot_dictionary_new_copy_allocs = x.GodotDictionaryNewCopy.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_new_copy_allocs)

	var cgodot_dictionary_destroy_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_destroy, cgodot_dictionary_destroy_allocs = x.GodotDictionaryDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_destroy_allocs)

	var cgodot_dictionary_size_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_size, cgodot_dictionary_size_allocs = x.GodotDictionarySize.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_size_allocs)

	var cgodot_dictionary_empty_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_empty, cgodot_dictionary_empty_allocs = x.GodotDictionaryEmpty.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_empty_allocs)

	var cgodot_dictionary_clear_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_clear, cgodot_dictionary_clear_allocs = x.GodotDictionaryClear.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_clear_allocs)

	var cgodot_dictionary_has_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_has, cgodot_dictionary_has_allocs = x.GodotDictionaryHas.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_has_allocs)

	var cgodot_dictionary_has_all_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_has_all, cgodot_dictionary_has_all_allocs = x.GodotDictionaryHasAll.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_has_all_allocs)

	var cgodot_dictionary_erase_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_erase, cgodot_dictionary_erase_allocs = x.GodotDictionaryErase.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_erase_allocs)

	var cgodot_dictionary_hash_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_hash, cgodot_dictionary_hash_allocs = x.GodotDictionaryHash.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_hash_allocs)

	var cgodot_dictionary_keys_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_keys, cgodot_dictionary_keys_allocs = x.GodotDictionaryKeys.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_keys_allocs)

	var cgodot_dictionary_values_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_values, cgodot_dictionary_values_allocs = x.GodotDictionaryValues.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_values_allocs)

	var cgodot_dictionary_get_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_get, cgodot_dictionary_get_allocs = x.GodotDictionaryGet.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_get_allocs)

	var cgodot_dictionary_set_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_set, cgodot_dictionary_set_allocs = x.GodotDictionarySet.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_set_allocs)

	var cgodot_dictionary_operator_index_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_operator_index, cgodot_dictionary_operator_index_allocs = x.GodotDictionaryOperatorIndex.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_operator_index_allocs)

	var cgodot_dictionary_operator_index_const_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_operator_index_const, cgodot_dictionary_operator_index_const_allocs = x.GodotDictionaryOperatorIndexConst.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_operator_index_const_allocs)

	var cgodot_dictionary_next_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_next, cgodot_dictionary_next_allocs = x.GodotDictionaryNext.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_next_allocs)

	var cgodot_dictionary_operator_equal_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_operator_equal, cgodot_dictionary_operator_equal_allocs = x.GodotDictionaryOperatorEqual.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_operator_equal_allocs)

	var cgodot_dictionary_to_json_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_to_json, cgodot_dictionary_to_json_allocs = x.GodotDictionaryToJson.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_to_json_allocs)

	var cgodot_node_path_new_allocs *cgoAllocMap
	ref57717e51.godot_node_path_new, cgodot_node_path_new_allocs = x.GodotNodePathNew.PassRef()
	allocs57717e51.Borrow(cgodot_node_path_new_allocs)

	var cgodot_node_path_new_copy_allocs *cgoAllocMap
	ref57717e51.godot_node_path_new_copy, cgodot_node_path_new_copy_allocs = x.GodotNodePathNewCopy.PassRef()
	allocs57717e51.Borrow(cgodot_node_path_new_copy_allocs)

	var cgodot_node_path_destroy_allocs *cgoAllocMap
	ref57717e51.godot_node_path_destroy, cgodot_node_path_destroy_allocs = x.GodotNodePathDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_node_path_destroy_allocs)

	var cgodot_node_path_as_string_allocs *cgoAllocMap
	ref57717e51.godot_node_path_as_string, cgodot_node_path_as_string_allocs = x.GodotNodePathAsString.PassRef()
	allocs57717e51.Borrow(cgodot_node_path_as_string_allocs)

	var cgodot_node_path_is_absolute_allocs *cgoAllocMap
	ref57717e51.godot_node_path_is_absolute, cgodot_node_path_is_absolute_allocs = x.GodotNodePathIsAbsolute.PassRef()
	allocs57717e51.Borrow(cgodot_node_path_is_absolute_allocs)

	var cgodot_node_path_get_name_count_allocs *cgoAllocMap
	ref57717e51.godot_node_path_get_name_count, cgodot_node_path_get_name_count_allocs = x.GodotNodePathGetNameCount.PassRef()
	allocs57717e51.Borrow(cgodot_node_path_get_name_count_allocs)

	var cgodot_node_path_get_name_allocs *cgoAllocMap
	ref57717e51.godot_node_path_get_name, cgodot_node_path_get_name_allocs = x.GodotNodePathGetName.PassRef()
	allocs57717e51.Borrow(cgodot_node_path_get_name_allocs)

	var cgodot_node_path_get_subname_count_allocs *cgoAllocMap
	ref57717e51.godot_node_path_get_subname_count, cgodot_node_path_get_subname_count_allocs = x.GodotNodePathGetSubnameCount.PassRef()
	allocs57717e51.Borrow(cgodot_node_path_get_subname_count_allocs)

	var cgodot_node_path_get_subname_allocs *cgoAllocMap
	ref57717e51.godot_node_path_get_subname, cgodot_node_path_get_subname_allocs = x.GodotNodePathGetSubname.PassRef()
	allocs57717e51.Borrow(cgodot_node_path_get_subname_allocs)

	var cgodot_node_path_get_concatenated_subnames_allocs *cgoAllocMap
	ref57717e51.godot_node_path_get_concatenated_subnames, cgodot_node_path_get_concatenated_subnames_allocs = x.GodotNodePathGetConcatenatedSubnames.PassRef()
	allocs57717e51.Borrow(cgodot_node_path_get_concatenated_subnames_allocs)

	var cgodot_node_path_is_empty_allocs *cgoAllocMap
	ref57717e51.godot_node_path_is_empty, cgodot_node_path_is_empty_allocs = x.GodotNodePathIsEmpty.PassRef()
	allocs57717e51.Borrow(cgodot_node_path_is_empty_allocs)

	var cgodot_node_path_operator_equal_allocs *cgoAllocMap
	ref57717e51.godot_node_path_operator_equal, cgodot_node_path_operator_equal_allocs = x.GodotNodePathOperatorEqual.PassRef()
	allocs57717e51.Borrow(cgodot_node_path_operator_equal_allocs)

	var cgodot_plane_new_with_reals_allocs *cgoAllocMap
	ref57717e51.godot_plane_new_with_reals, cgodot_plane_new_with_reals_allocs = x.GodotPlaneNewWithReals.PassRef()
	allocs57717e51.Borrow(cgodot_plane_new_with_reals_allocs)

	var cgodot_plane_new_with_vectors_allocs *cgoAllocMap
	ref57717e51.godot_plane_new_with_vectors, cgodot_plane_new_with_vectors_allocs = x.GodotPlaneNewWithVectors.PassRef()
	allocs57717e51.Borrow(cgodot_plane_new_with_vectors_allocs)

	var cgodot_plane_new_with_normal_allocs *cgoAllocMap
	ref57717e51.godot_plane_new_with_normal, cgodot_plane_new_with_normal_allocs = x.GodotPlaneNewWithNormal.PassRef()
	allocs57717e51.Borrow(cgodot_plane_new_with_normal_allocs)

	var cgodot_plane_as_string_allocs *cgoAllocMap
	ref57717e51.godot_plane_as_string, cgodot_plane_as_string_allocs = x.GodotPlaneAsString.PassRef()
	allocs57717e51.Borrow(cgodot_plane_as_string_allocs)

	var cgodot_plane_normalized_allocs *cgoAllocMap
	ref57717e51.godot_plane_normalized, cgodot_plane_normalized_allocs = x.GodotPlaneNormalized.PassRef()
	allocs57717e51.Borrow(cgodot_plane_normalized_allocs)

	var cgodot_plane_center_allocs *cgoAllocMap
	ref57717e51.godot_plane_center, cgodot_plane_center_allocs = x.GodotPlaneCenter.PassRef()
	allocs57717e51.Borrow(cgodot_plane_center_allocs)

	var cgodot_plane_get_any_point_allocs *cgoAllocMap
	ref57717e51.godot_plane_get_any_point, cgodot_plane_get_any_point_allocs = x.GodotPlaneGetAnyPoint.PassRef()
	allocs57717e51.Borrow(cgodot_plane_get_any_point_allocs)

	var cgodot_plane_is_point_over_allocs *cgoAllocMap
	ref57717e51.godot_plane_is_point_over, cgodot_plane_is_point_over_allocs = x.GodotPlaneIsPointOver.PassRef()
	allocs57717e51.Borrow(cgodot_plane_is_point_over_allocs)

	var cgodot_plane_distance_to_allocs *cgoAllocMap
	ref57717e51.godot_plane_distance_to, cgodot_plane_distance_to_allocs = x.GodotPlaneDistanceTo.PassRef()
	allocs57717e51.Borrow(cgodot_plane_distance_to_allocs)

	var cgodot_plane_has_point_allocs *cgoAllocMap
	ref57717e51.godot_plane_has_point, cgodot_plane_has_point_allocs = x.GodotPlaneHasPoint.PassRef()
	allocs57717e51.Borrow(cgodot_plane_has_point_allocs)

	var cgodot_plane_project_allocs *cgoAllocMap
	ref57717e51.godot_plane_project, cgodot_plane_project_allocs = x.GodotPlaneProject.PassRef()
	allocs57717e51.Borrow(cgodot_plane_project_allocs)

	var cgodot_plane_intersect_3_allocs *cgoAllocMap
	ref57717e51.godot_plane_intersect_3, cgodot_plane_intersect_3_allocs = x.GodotPlaneIntersect3.PassRef()
	allocs57717e51.Borrow(cgodot_plane_intersect_3_allocs)

	var cgodot_plane_intersects_ray_allocs *cgoAllocMap
	ref57717e51.godot_plane_intersects_ray, cgodot_plane_intersects_ray_allocs = x.GodotPlaneIntersectsRay.PassRef()
	allocs57717e51.Borrow(cgodot_plane_intersects_ray_allocs)

	var cgodot_plane_intersects_segment_allocs *cgoAllocMap
	ref57717e51.godot_plane_intersects_segment, cgodot_plane_intersects_segment_allocs = x.GodotPlaneIntersectsSegment.PassRef()
	allocs57717e51.Borrow(cgodot_plane_intersects_segment_allocs)

	var cgodot_plane_operator_neg_allocs *cgoAllocMap
	ref57717e51.godot_plane_operator_neg, cgodot_plane_operator_neg_allocs = x.GodotPlaneOperatorNeg.PassRef()
	allocs57717e51.Borrow(cgodot_plane_operator_neg_allocs)

	var cgodot_plane_operator_equal_allocs *cgoAllocMap
	ref57717e51.godot_plane_operator_equal, cgodot_plane_operator_equal_allocs = x.GodotPlaneOperatorEqual.PassRef()
	allocs57717e51.Borrow(cgodot_plane_operator_equal_allocs)

	var cgodot_plane_set_normal_allocs *cgoAllocMap
	ref57717e51.godot_plane_set_normal, cgodot_plane_set_normal_allocs = x.GodotPlaneSetNormal.PassRef()
	allocs57717e51.Borrow(cgodot_plane_set_normal_allocs)

	var cgodot_plane_get_normal_allocs *cgoAllocMap
	ref57717e51.godot_plane_get_normal, cgodot_plane_get_normal_allocs = x.GodotPlaneGetNormal.PassRef()
	allocs57717e51.Borrow(cgodot_plane_get_normal_allocs)

	var cgodot_plane_get_d_allocs *cgoAllocMap
	ref57717e51.godot_plane_get_d, cgodot_plane_get_d_allocs = x.GodotPlaneGetD.PassRef()
	allocs57717e51.Borrow(cgodot_plane_get_d_allocs)

	var cgodot_plane_set_d_allocs *cgoAllocMap
	ref57717e51.godot_plane_set_d, cgodot_plane_set_d_allocs = x.GodotPlaneSetD.PassRef()
	allocs57717e51.Borrow(cgodot_plane_set_d_allocs)

	var cgodot_rect2_new_with_position_and_size_allocs *cgoAllocMap
	ref57717e51.godot_rect2_new_with_position_and_size, cgodot_rect2_new_with_position_and_size_allocs = x.GodotRect2NewWithPositionAndSize.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_new_with_position_and_size_allocs)

	var cgodot_rect2_new_allocs *cgoAllocMap
	ref57717e51.godot_rect2_new, cgodot_rect2_new_allocs = x.GodotRect2New.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_new_allocs)

	var cgodot_rect2_as_string_allocs *cgoAllocMap
	ref57717e51.godot_rect2_as_string, cgodot_rect2_as_string_allocs = x.GodotRect2AsString.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_as_string_allocs)

	var cgodot_rect2_get_area_allocs *cgoAllocMap
	ref57717e51.godot_rect2_get_area, cgodot_rect2_get_area_allocs = x.GodotRect2GetArea.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_get_area_allocs)

	var cgodot_rect2_intersects_allocs *cgoAllocMap
	ref57717e51.godot_rect2_intersects, cgodot_rect2_intersects_allocs = x.GodotRect2Intersects.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_intersects_allocs)

	var cgodot_rect2_encloses_allocs *cgoAllocMap
	ref57717e51.godot_rect2_encloses, cgodot_rect2_encloses_allocs = x.GodotRect2Encloses.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_encloses_allocs)

	var cgodot_rect2_has_no_area_allocs *cgoAllocMap
	ref57717e51.godot_rect2_has_no_area, cgodot_rect2_has_no_area_allocs = x.GodotRect2HasNoArea.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_has_no_area_allocs)

	var cgodot_rect2_clip_allocs *cgoAllocMap
	ref57717e51.godot_rect2_clip, cgodot_rect2_clip_allocs = x.GodotRect2Clip.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_clip_allocs)

	var cgodot_rect2_merge_allocs *cgoAllocMap
	ref57717e51.godot_rect2_merge, cgodot_rect2_merge_allocs = x.GodotRect2Merge.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_merge_allocs)

	var cgodot_rect2_has_point_allocs *cgoAllocMap
	ref57717e51.godot_rect2_has_point, cgodot_rect2_has_point_allocs = x.GodotRect2HasPoint.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_has_point_allocs)

	var cgodot_rect2_grow_allocs *cgoAllocMap
	ref57717e51.godot_rect2_grow, cgodot_rect2_grow_allocs = x.GodotRect2Grow.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_grow_allocs)

	var cgodot_rect2_expand_allocs *cgoAllocMap
	ref57717e51.godot_rect2_expand, cgodot_rect2_expand_allocs = x.GodotRect2Expand.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_expand_allocs)

	var cgodot_rect2_operator_equal_allocs *cgoAllocMap
	ref57717e51.godot_rect2_operator_equal, cgodot_rect2_operator_equal_allocs = x.GodotRect2OperatorEqual.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_operator_equal_allocs)

	var cgodot_rect2_get_position_allocs *cgoAllocMap
	ref57717e51.godot_rect2_get_position, cgodot_rect2_get_position_allocs = x.GodotRect2GetPosition.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_get_position_allocs)

	var cgodot_rect2_get_size_allocs *cgoAllocMap
	ref57717e51.godot_rect2_get_size, cgodot_rect2_get_size_allocs = x.GodotRect2GetSize.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_get_size_allocs)

	var cgodot_rect2_set_position_allocs *cgoAllocMap
	ref57717e51.godot_rect2_set_position, cgodot_rect2_set_position_allocs = x.GodotRect2SetPosition.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_set_position_allocs)

	var cgodot_rect2_set_size_allocs *cgoAllocMap
	ref57717e51.godot_rect2_set_size, cgodot_rect2_set_size_allocs = x.GodotRect2SetSize.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_set_size_allocs)

	var cgodot_aabb_new_allocs *cgoAllocMap
	ref57717e51.godot_aabb_new, cgodot_aabb_new_allocs = x.GodotAabbNew.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_new_allocs)

	var cgodot_aabb_get_position_allocs *cgoAllocMap
	ref57717e51.godot_aabb_get_position, cgodot_aabb_get_position_allocs = x.GodotAabbGetPosition.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_get_position_allocs)

	var cgodot_aabb_set_position_allocs *cgoAllocMap
	ref57717e51.godot_aabb_set_position, cgodot_aabb_set_position_allocs = x.GodotAabbSetPosition.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_set_position_allocs)

	var cgodot_aabb_get_size_allocs *cgoAllocMap
	ref57717e51.godot_aabb_get_size, cgodot_aabb_get_size_allocs = x.GodotAabbGetSize.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_get_size_allocs)

	var cgodot_aabb_set_size_allocs *cgoAllocMap
	ref57717e51.godot_aabb_set_size, cgodot_aabb_set_size_allocs = x.GodotAabbSetSize.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_set_size_allocs)

	var cgodot_aabb_as_string_allocs *cgoAllocMap
	ref57717e51.godot_aabb_as_string, cgodot_aabb_as_string_allocs = x.GodotAabbAsString.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_as_string_allocs)

	var cgodot_aabb_get_area_allocs *cgoAllocMap
	ref57717e51.godot_aabb_get_area, cgodot_aabb_get_area_allocs = x.GodotAabbGetArea.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_get_area_allocs)

	var cgodot_aabb_has_no_area_allocs *cgoAllocMap
	ref57717e51.godot_aabb_has_no_area, cgodot_aabb_has_no_area_allocs = x.GodotAabbHasNoArea.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_has_no_area_allocs)

	var cgodot_aabb_has_no_surface_allocs *cgoAllocMap
	ref57717e51.godot_aabb_has_no_surface, cgodot_aabb_has_no_surface_allocs = x.GodotAabbHasNoSurface.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_has_no_surface_allocs)

	var cgodot_aabb_intersects_allocs *cgoAllocMap
	ref57717e51.godot_aabb_intersects, cgodot_aabb_intersects_allocs = x.GodotAabbIntersects.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_intersects_allocs)

	var cgodot_aabb_encloses_allocs *cgoAllocMap
	ref57717e51.godot_aabb_encloses, cgodot_aabb_encloses_allocs = x.GodotAabbEncloses.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_encloses_allocs)

	var cgodot_aabb_merge_allocs *cgoAllocMap
	ref57717e51.godot_aabb_merge, cgodot_aabb_merge_allocs = x.GodotAabbMerge.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_merge_allocs)

	var cgodot_aabb_intersection_allocs *cgoAllocMap
	ref57717e51.godot_aabb_intersection, cgodot_aabb_intersection_allocs = x.GodotAabbIntersection.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_intersection_allocs)

	var cgodot_aabb_intersects_plane_allocs *cgoAllocMap
	ref57717e51.godot_aabb_intersects_plane, cgodot_aabb_intersects_plane_allocs = x.GodotAabbIntersectsPlane.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_intersects_plane_allocs)

	var cgodot_aabb_intersects_segment_allocs *cgoAllocMap
	ref57717e51.godot_aabb_intersects_segment, cgodot_aabb_intersects_segment_allocs = x.GodotAabbIntersectsSegment.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_intersects_segment_allocs)

	var cgodot_aabb_has_point_allocs *cgoAllocMap
	ref57717e51.godot_aabb_has_point, cgodot_aabb_has_point_allocs = x.GodotAabbHasPoint.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_has_point_allocs)

	var cgodot_aabb_get_support_allocs *cgoAllocMap
	ref57717e51.godot_aabb_get_support, cgodot_aabb_get_support_allocs = x.GodotAabbGetSupport.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_get_support_allocs)

	var cgodot_aabb_get_longest_axis_allocs *cgoAllocMap
	ref57717e51.godot_aabb_get_longest_axis, cgodot_aabb_get_longest_axis_allocs = x.GodotAabbGetLongestAxis.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_get_longest_axis_allocs)

	var cgodot_aabb_get_longest_axis_index_allocs *cgoAllocMap
	ref57717e51.godot_aabb_get_longest_axis_index, cgodot_aabb_get_longest_axis_index_allocs = x.GodotAabbGetLongestAxisIndex.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_get_longest_axis_index_allocs)

	var cgodot_aabb_get_longest_axis_size_allocs *cgoAllocMap
	ref57717e51.godot_aabb_get_longest_axis_size, cgodot_aabb_get_longest_axis_size_allocs = x.GodotAabbGetLongestAxisSize.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_get_longest_axis_size_allocs)

	var cgodot_aabb_get_shortest_axis_allocs *cgoAllocMap
	ref57717e51.godot_aabb_get_shortest_axis, cgodot_aabb_get_shortest_axis_allocs = x.GodotAabbGetShortestAxis.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_get_shortest_axis_allocs)

	var cgodot_aabb_get_shortest_axis_index_allocs *cgoAllocMap
	ref57717e51.godot_aabb_get_shortest_axis_index, cgodot_aabb_get_shortest_axis_index_allocs = x.GodotAabbGetShortestAxisIndex.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_get_shortest_axis_index_allocs)

	var cgodot_aabb_get_shortest_axis_size_allocs *cgoAllocMap
	ref57717e51.godot_aabb_get_shortest_axis_size, cgodot_aabb_get_shortest_axis_size_allocs = x.GodotAabbGetShortestAxisSize.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_get_shortest_axis_size_allocs)

	var cgodot_aabb_expand_allocs *cgoAllocMap
	ref57717e51.godot_aabb_expand, cgodot_aabb_expand_allocs = x.GodotAabbExpand.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_expand_allocs)

	var cgodot_aabb_grow_allocs *cgoAllocMap
	ref57717e51.godot_aabb_grow, cgodot_aabb_grow_allocs = x.GodotAabbGrow.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_grow_allocs)

	var cgodot_aabb_get_endpoint_allocs *cgoAllocMap
	ref57717e51.godot_aabb_get_endpoint, cgodot_aabb_get_endpoint_allocs = x.GodotAabbGetEndpoint.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_get_endpoint_allocs)

	var cgodot_aabb_operator_equal_allocs *cgoAllocMap
	ref57717e51.godot_aabb_operator_equal, cgodot_aabb_operator_equal_allocs = x.GodotAabbOperatorEqual.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_operator_equal_allocs)

	var cgodot_rid_new_allocs *cgoAllocMap
	ref57717e51.godot_rid_new, cgodot_rid_new_allocs = x.GodotRidNew.PassRef()
	allocs57717e51.Borrow(cgodot_rid_new_allocs)

	var cgodot_rid_get_id_allocs *cgoAllocMap
	ref57717e51.godot_rid_get_id, cgodot_rid_get_id_allocs = x.GodotRidGetId.PassRef()
	allocs57717e51.Borrow(cgodot_rid_get_id_allocs)

	var cgodot_rid_new_with_resource_allocs *cgoAllocMap
	ref57717e51.godot_rid_new_with_resource, cgodot_rid_new_with_resource_allocs = x.GodotRidNewWithResource.PassRef()
	allocs57717e51.Borrow(cgodot_rid_new_with_resource_allocs)

	var cgodot_rid_operator_equal_allocs *cgoAllocMap
	ref57717e51.godot_rid_operator_equal, cgodot_rid_operator_equal_allocs = x.GodotRidOperatorEqual.PassRef()
	allocs57717e51.Borrow(cgodot_rid_operator_equal_allocs)

	var cgodot_rid_operator_less_allocs *cgoAllocMap
	ref57717e51.godot_rid_operator_less, cgodot_rid_operator_less_allocs = x.GodotRidOperatorLess.PassRef()
	allocs57717e51.Borrow(cgodot_rid_operator_less_allocs)

	var cgodot_transform_new_with_axis_origin_allocs *cgoAllocMap
	ref57717e51.godot_transform_new_with_axis_origin, cgodot_transform_new_with_axis_origin_allocs = x.GodotTransformNewWithAxisOrigin.PassRef()
	allocs57717e51.Borrow(cgodot_transform_new_with_axis_origin_allocs)

	var cgodot_transform_new_allocs *cgoAllocMap
	ref57717e51.godot_transform_new, cgodot_transform_new_allocs = x.GodotTransformNew.PassRef()
	allocs57717e51.Borrow(cgodot_transform_new_allocs)

	var cgodot_transform_get_basis_allocs *cgoAllocMap
	ref57717e51.godot_transform_get_basis, cgodot_transform_get_basis_allocs = x.GodotTransformGetBasis.PassRef()
	allocs57717e51.Borrow(cgodot_transform_get_basis_allocs)

	var cgodot_transform_set_basis_allocs *cgoAllocMap
	ref57717e51.godot_transform_set_basis, cgodot_transform_set_basis_allocs = x.GodotTransformSetBasis.PassRef()
	allocs57717e51.Borrow(cgodot_transform_set_basis_allocs)

	var cgodot_transform_get_origin_allocs *cgoAllocMap
	ref57717e51.godot_transform_get_origin, cgodot_transform_get_origin_allocs = x.GodotTransformGetOrigin.PassRef()
	allocs57717e51.Borrow(cgodot_transform_get_origin_allocs)

	var cgodot_transform_set_origin_allocs *cgoAllocMap
	ref57717e51.godot_transform_set_origin, cgodot_transform_set_origin_allocs = x.GodotTransformSetOrigin.PassRef()
	allocs57717e51.Borrow(cgodot_transform_set_origin_allocs)

	var cgodot_transform_as_string_allocs *cgoAllocMap
	ref57717e51.godot_transform_as_string, cgodot_transform_as_string_allocs = x.GodotTransformAsString.PassRef()
	allocs57717e51.Borrow(cgodot_transform_as_string_allocs)

	var cgodot_transform_inverse_allocs *cgoAllocMap
	ref57717e51.godot_transform_inverse, cgodot_transform_inverse_allocs = x.GodotTransformInverse.PassRef()
	allocs57717e51.Borrow(cgodot_transform_inverse_allocs)

	var cgodot_transform_affine_inverse_allocs *cgoAllocMap
	ref57717e51.godot_transform_affine_inverse, cgodot_transform_affine_inverse_allocs = x.GodotTransformAffineInverse.PassRef()
	allocs57717e51.Borrow(cgodot_transform_affine_inverse_allocs)

	var cgodot_transform_orthonormalized_allocs *cgoAllocMap
	ref57717e51.godot_transform_orthonormalized, cgodot_transform_orthonormalized_allocs = x.GodotTransformOrthonormalized.PassRef()
	allocs57717e51.Borrow(cgodot_transform_orthonormalized_allocs)

	var cgodot_transform_rotated_allocs *cgoAllocMap
	ref57717e51.godot_transform_rotated, cgodot_transform_rotated_allocs = x.GodotTransformRotated.PassRef()
	allocs57717e51.Borrow(cgodot_transform_rotated_allocs)

	var cgodot_transform_scaled_allocs *cgoAllocMap
	ref57717e51.godot_transform_scaled, cgodot_transform_scaled_allocs = x.GodotTransformScaled.PassRef()
	allocs57717e51.Borrow(cgodot_transform_scaled_allocs)

	var cgodot_transform_translated_allocs *cgoAllocMap
	ref57717e51.godot_transform_translated, cgodot_transform_translated_allocs = x.GodotTransformTranslated.PassRef()
	allocs57717e51.Borrow(cgodot_transform_translated_allocs)

	var cgodot_transform_looking_at_allocs *cgoAllocMap
	ref57717e51.godot_transform_looking_at, cgodot_transform_looking_at_allocs = x.GodotTransformLookingAt.PassRef()
	allocs57717e51.Borrow(cgodot_transform_looking_at_allocs)

	var cgodot_transform_xform_plane_allocs *cgoAllocMap
	ref57717e51.godot_transform_xform_plane, cgodot_transform_xform_plane_allocs = x.GodotTransformXformPlane.PassRef()
	allocs57717e51.Borrow(cgodot_transform_xform_plane_allocs)

	var cgodot_transform_xform_inv_plane_allocs *cgoAllocMap
	ref57717e51.godot_transform_xform_inv_plane, cgodot_transform_xform_inv_plane_allocs = x.GodotTransformXformInvPlane.PassRef()
	allocs57717e51.Borrow(cgodot_transform_xform_inv_plane_allocs)

	var cgodot_transform_new_identity_allocs *cgoAllocMap
	ref57717e51.godot_transform_new_identity, cgodot_transform_new_identity_allocs = x.GodotTransformNewIdentity.PassRef()
	allocs57717e51.Borrow(cgodot_transform_new_identity_allocs)

	var cgodot_transform_operator_equal_allocs *cgoAllocMap
	ref57717e51.godot_transform_operator_equal, cgodot_transform_operator_equal_allocs = x.GodotTransformOperatorEqual.PassRef()
	allocs57717e51.Borrow(cgodot_transform_operator_equal_allocs)

	var cgodot_transform_operator_multiply_allocs *cgoAllocMap
	ref57717e51.godot_transform_operator_multiply, cgodot_transform_operator_multiply_allocs = x.GodotTransformOperatorMultiply.PassRef()
	allocs57717e51.Borrow(cgodot_transform_operator_multiply_allocs)

	var cgodot_transform_xform_vector3_allocs *cgoAllocMap
	ref57717e51.godot_transform_xform_vector3, cgodot_transform_xform_vector3_allocs = x.GodotTransformXformVector3.PassRef()
	allocs57717e51.Borrow(cgodot_transform_xform_vector3_allocs)

	var cgodot_transform_xform_inv_vector3_allocs *cgoAllocMap
	ref57717e51.godot_transform_xform_inv_vector3, cgodot_transform_xform_inv_vector3_allocs = x.GodotTransformXformInvVector3.PassRef()
	allocs57717e51.Borrow(cgodot_transform_xform_inv_vector3_allocs)

	var cgodot_transform_xform_aabb_allocs *cgoAllocMap
	ref57717e51.godot_transform_xform_aabb, cgodot_transform_xform_aabb_allocs = x.GodotTransformXformAabb.PassRef()
	allocs57717e51.Borrow(cgodot_transform_xform_aabb_allocs)

	var cgodot_transform_xform_inv_aabb_allocs *cgoAllocMap
	ref57717e51.godot_transform_xform_inv_aabb, cgodot_transform_xform_inv_aabb_allocs = x.GodotTransformXformInvAabb.PassRef()
	allocs57717e51.Borrow(cgodot_transform_xform_inv_aabb_allocs)

	var cgodot_transform2d_new_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_new, cgodot_transform2d_new_allocs = x.GodotTransform2dNew.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_new_allocs)

	var cgodot_transform2d_new_axis_origin_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_new_axis_origin, cgodot_transform2d_new_axis_origin_allocs = x.GodotTransform2dNewAxisOrigin.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_new_axis_origin_allocs)

	var cgodot_transform2d_as_string_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_as_string, cgodot_transform2d_as_string_allocs = x.GodotTransform2dAsString.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_as_string_allocs)

	var cgodot_transform2d_inverse_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_inverse, cgodot_transform2d_inverse_allocs = x.GodotTransform2dInverse.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_inverse_allocs)

	var cgodot_transform2d_affine_inverse_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_affine_inverse, cgodot_transform2d_affine_inverse_allocs = x.GodotTransform2dAffineInverse.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_affine_inverse_allocs)

	var cgodot_transform2d_get_rotation_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_get_rotation, cgodot_transform2d_get_rotation_allocs = x.GodotTransform2dGetRotation.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_get_rotation_allocs)

	var cgodot_transform2d_get_origin_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_get_origin, cgodot_transform2d_get_origin_allocs = x.GodotTransform2dGetOrigin.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_get_origin_allocs)

	var cgodot_transform2d_get_scale_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_get_scale, cgodot_transform2d_get_scale_allocs = x.GodotTransform2dGetScale.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_get_scale_allocs)

	var cgodot_transform2d_orthonormalized_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_orthonormalized, cgodot_transform2d_orthonormalized_allocs = x.GodotTransform2dOrthonormalized.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_orthonormalized_allocs)

	var cgodot_transform2d_rotated_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_rotated, cgodot_transform2d_rotated_allocs = x.GodotTransform2dRotated.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_rotated_allocs)

	var cgodot_transform2d_scaled_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_scaled, cgodot_transform2d_scaled_allocs = x.GodotTransform2dScaled.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_scaled_allocs)

	var cgodot_transform2d_translated_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_translated, cgodot_transform2d_translated_allocs = x.GodotTransform2dTranslated.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_translated_allocs)

	var cgodot_transform2d_xform_vector2_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_xform_vector2, cgodot_transform2d_xform_vector2_allocs = x.GodotTransform2dXformVector2.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_xform_vector2_allocs)

	var cgodot_transform2d_xform_inv_vector2_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_xform_inv_vector2, cgodot_transform2d_xform_inv_vector2_allocs = x.GodotTransform2dXformInvVector2.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_xform_inv_vector2_allocs)

	var cgodot_transform2d_basis_xform_vector2_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_basis_xform_vector2, cgodot_transform2d_basis_xform_vector2_allocs = x.GodotTransform2dBasisXformVector2.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_basis_xform_vector2_allocs)

	var cgodot_transform2d_basis_xform_inv_vector2_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_basis_xform_inv_vector2, cgodot_transform2d_basis_xform_inv_vector2_allocs = x.GodotTransform2dBasisXformInvVector2.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_basis_xform_inv_vector2_allocs)

	var cgodot_transform2d_interpolate_with_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_interpolate_with, cgodot_transform2d_interpolate_with_allocs = x.GodotTransform2dInterpolateWith.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_interpolate_with_allocs)

	var cgodot_transform2d_operator_equal_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_operator_equal, cgodot_transform2d_operator_equal_allocs = x.GodotTransform2dOperatorEqual.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_operator_equal_allocs)

	var cgodot_transform2d_operator_multiply_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_operator_multiply, cgodot_transform2d_operator_multiply_allocs = x.GodotTransform2dOperatorMultiply.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_operator_multiply_allocs)

	var cgodot_transform2d_new_identity_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_new_identity, cgodot_transform2d_new_identity_allocs = x.GodotTransform2dNewIdentity.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_new_identity_allocs)

	var cgodot_transform2d_xform_rect2_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_xform_rect2, cgodot_transform2d_xform_rect2_allocs = x.GodotTransform2dXformRect2.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_xform_rect2_allocs)

	var cgodot_transform2d_xform_inv_rect2_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_xform_inv_rect2, cgodot_transform2d_xform_inv_rect2_allocs = x.GodotTransform2dXformInvRect2.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_xform_inv_rect2_allocs)

	var cgodot_variant_get_type_allocs *cgoAllocMap
	ref57717e51.godot_variant_get_type, cgodot_variant_get_type_allocs = x.GodotVariantGetType.PassRef()
	allocs57717e51.Borrow(cgodot_variant_get_type_allocs)

	var cgodot_variant_new_copy_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_copy, cgodot_variant_new_copy_allocs = x.GodotVariantNewCopy.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_copy_allocs)

	var cgodot_variant_new_nil_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_nil, cgodot_variant_new_nil_allocs = x.GodotVariantNewNil.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_nil_allocs)

	var cgodot_variant_new_bool_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_bool, cgodot_variant_new_bool_allocs = x.GodotVariantNewBool.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_bool_allocs)

	var cgodot_variant_new_uint_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_uint, cgodot_variant_new_uint_allocs = x.GodotVariantNewUint.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_uint_allocs)

	var cgodot_variant_new_int_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_int, cgodot_variant_new_int_allocs = x.GodotVariantNewInt.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_int_allocs)

	var cgodot_variant_new_real_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_real, cgodot_variant_new_real_allocs = x.GodotVariantNewReal.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_real_allocs)

	var cgodot_variant_new_string_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_string, cgodot_variant_new_string_allocs = x.GodotVariantNewString.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_string_allocs)

	var cgodot_variant_new_vector2_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_vector2, cgodot_variant_new_vector2_allocs = x.GodotVariantNewVector2.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_vector2_allocs)

	var cgodot_variant_new_rect2_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_rect2, cgodot_variant_new_rect2_allocs = x.GodotVariantNewRect2.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_rect2_allocs)

	var cgodot_variant_new_vector3_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_vector3, cgodot_variant_new_vector3_allocs = x.GodotVariantNewVector3.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_vector3_allocs)

	var cgodot_variant_new_transform2d_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_transform2d, cgodot_variant_new_transform2d_allocs = x.GodotVariantNewTransform2d.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_transform2d_allocs)

	var cgodot_variant_new_plane_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_plane, cgodot_variant_new_plane_allocs = x.GodotVariantNewPlane.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_plane_allocs)

	var cgodot_variant_new_quat_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_quat, cgodot_variant_new_quat_allocs = x.GodotVariantNewQuat.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_quat_allocs)

	var cgodot_variant_new_aabb_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_aabb, cgodot_variant_new_aabb_allocs = x.GodotVariantNewAabb.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_aabb_allocs)

	var cgodot_variant_new_basis_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_basis, cgodot_variant_new_basis_allocs = x.GodotVariantNewBasis.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_basis_allocs)

	var cgodot_variant_new_transform_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_transform, cgodot_variant_new_transform_allocs = x.GodotVariantNewTransform.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_transform_allocs)

	var cgodot_variant_new_color_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_color, cgodot_variant_new_color_allocs = x.GodotVariantNewColor.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_color_allocs)

	var cgodot_variant_new_node_path_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_node_path, cgodot_variant_new_node_path_allocs = x.GodotVariantNewNodePath.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_node_path_allocs)

	var cgodot_variant_new_rid_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_rid, cgodot_variant_new_rid_allocs = x.GodotVariantNewRid.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_rid_allocs)

	var cgodot_variant_new_object_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_object, cgodot_variant_new_object_allocs = x.GodotVariantNewObject.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_object_allocs)

	var cgodot_variant_new_dictionary_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_dictionary, cgodot_variant_new_dictionary_allocs = x.GodotVariantNewDictionary.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_dictionary_allocs)

	var cgodot_variant_new_array_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_array, cgodot_variant_new_array_allocs = x.GodotVariantNewArray.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_array_allocs)

	var cgodot_variant_new_pool_byte_array_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_pool_byte_array, cgodot_variant_new_pool_byte_array_allocs = x.GodotVariantNewPoolByteArray.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_pool_byte_array_allocs)

	var cgodot_variant_new_pool_int_array_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_pool_int_array, cgodot_variant_new_pool_int_array_allocs = x.GodotVariantNewPoolIntArray.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_pool_int_array_allocs)

	var cgodot_variant_new_pool_real_array_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_pool_real_array, cgodot_variant_new_pool_real_array_allocs = x.GodotVariantNewPoolRealArray.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_pool_real_array_allocs)

	var cgodot_variant_new_pool_string_array_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_pool_string_array, cgodot_variant_new_pool_string_array_allocs = x.GodotVariantNewPoolStringArray.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_pool_string_array_allocs)

	var cgodot_variant_new_pool_vector2_array_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_pool_vector2_array, cgodot_variant_new_pool_vector2_array_allocs = x.GodotVariantNewPoolVector2Array.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_pool_vector2_array_allocs)

	var cgodot_variant_new_pool_vector3_array_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_pool_vector3_array, cgodot_variant_new_pool_vector3_array_allocs = x.GodotVariantNewPoolVector3Array.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_pool_vector3_array_allocs)

	var cgodot_variant_new_pool_color_array_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_pool_color_array, cgodot_variant_new_pool_color_array_allocs = x.GodotVariantNewPoolColorArray.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_pool_color_array_allocs)

	var cgodot_variant_as_bool_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_bool, cgodot_variant_as_bool_allocs = x.GodotVariantAsBool.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_bool_allocs)

	var cgodot_variant_as_uint_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_uint, cgodot_variant_as_uint_allocs = x.GodotVariantAsUint.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_uint_allocs)

	var cgodot_variant_as_int_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_int, cgodot_variant_as_int_allocs = x.GodotVariantAsInt.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_int_allocs)

	var cgodot_variant_as_real_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_real, cgodot_variant_as_real_allocs = x.GodotVariantAsReal.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_real_allocs)

	var cgodot_variant_as_string_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_string, cgodot_variant_as_string_allocs = x.GodotVariantAsString.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_string_allocs)

	var cgodot_variant_as_vector2_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_vector2, cgodot_variant_as_vector2_allocs = x.GodotVariantAsVector2.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_vector2_allocs)

	var cgodot_variant_as_rect2_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_rect2, cgodot_variant_as_rect2_allocs = x.GodotVariantAsRect2.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_rect2_allocs)

	var cgodot_variant_as_vector3_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_vector3, cgodot_variant_as_vector3_allocs = x.GodotVariantAsVector3.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_vector3_allocs)

	var cgodot_variant_as_transform2d_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_transform2d, cgodot_variant_as_transform2d_allocs = x.GodotVariantAsTransform2d.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_transform2d_allocs)

	var cgodot_variant_as_plane_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_plane, cgodot_variant_as_plane_allocs = x.GodotVariantAsPlane.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_plane_allocs)

	var cgodot_variant_as_quat_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_quat, cgodot_variant_as_quat_allocs = x.GodotVariantAsQuat.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_quat_allocs)

	var cgodot_variant_as_aabb_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_aabb, cgodot_variant_as_aabb_allocs = x.GodotVariantAsAabb.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_aabb_allocs)

	var cgodot_variant_as_basis_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_basis, cgodot_variant_as_basis_allocs = x.GodotVariantAsBasis.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_basis_allocs)

	var cgodot_variant_as_transform_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_transform, cgodot_variant_as_transform_allocs = x.GodotVariantAsTransform.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_transform_allocs)

	var cgodot_variant_as_color_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_color, cgodot_variant_as_color_allocs = x.GodotVariantAsColor.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_color_allocs)

	var cgodot_variant_as_node_path_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_node_path, cgodot_variant_as_node_path_allocs = x.GodotVariantAsNodePath.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_node_path_allocs)

	var cgodot_variant_as_rid_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_rid, cgodot_variant_as_rid_allocs = x.GodotVariantAsRid.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_rid_allocs)

	var cgodot_variant_as_object_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_object, cgodot_variant_as_object_allocs = x.GodotVariantAsObject.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_object_allocs)

	var cgodot_variant_as_dictionary_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_dictionary, cgodot_variant_as_dictionary_allocs = x.GodotVariantAsDictionary.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_dictionary_allocs)

	var cgodot_variant_as_array_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_array, cgodot_variant_as_array_allocs = x.GodotVariantAsArray.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_array_allocs)

	var cgodot_variant_as_pool_byte_array_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_pool_byte_array, cgodot_variant_as_pool_byte_array_allocs = x.GodotVariantAsPoolByteArray.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_pool_byte_array_allocs)

	var cgodot_variant_as_pool_int_array_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_pool_int_array, cgodot_variant_as_pool_int_array_allocs = x.GodotVariantAsPoolIntArray.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_pool_int_array_allocs)

	var cgodot_variant_as_pool_real_array_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_pool_real_array, cgodot_variant_as_pool_real_array_allocs = x.GodotVariantAsPoolRealArray.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_pool_real_array_allocs)

	var cgodot_variant_as_pool_string_array_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_pool_string_array, cgodot_variant_as_pool_string_array_allocs = x.GodotVariantAsPoolStringArray.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_pool_string_array_allocs)

	var cgodot_variant_as_pool_vector2_array_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_pool_vector2_array, cgodot_variant_as_pool_vector2_array_allocs = x.GodotVariantAsPoolVector2Array.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_pool_vector2_array_allocs)

	var cgodot_variant_as_pool_vector3_array_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_pool_vector3_array, cgodot_variant_as_pool_vector3_array_allocs = x.GodotVariantAsPoolVector3Array.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_pool_vector3_array_allocs)

	var cgodot_variant_as_pool_color_array_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_pool_color_array, cgodot_variant_as_pool_color_array_allocs = x.GodotVariantAsPoolColorArray.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_pool_color_array_allocs)

	var cgodot_variant_call_allocs *cgoAllocMap
	ref57717e51.godot_variant_call, cgodot_variant_call_allocs = x.GodotVariantCall.PassRef()
	allocs57717e51.Borrow(cgodot_variant_call_allocs)

	var cgodot_variant_has_method_allocs *cgoAllocMap
	ref57717e51.godot_variant_has_method, cgodot_variant_has_method_allocs = x.GodotVariantHasMethod.PassRef()
	allocs57717e51.Borrow(cgodot_variant_has_method_allocs)

	var cgodot_variant_operator_equal_allocs *cgoAllocMap
	ref57717e51.godot_variant_operator_equal, cgodot_variant_operator_equal_allocs = x.GodotVariantOperatorEqual.PassRef()
	allocs57717e51.Borrow(cgodot_variant_operator_equal_allocs)

	var cgodot_variant_operator_less_allocs *cgoAllocMap
	ref57717e51.godot_variant_operator_less, cgodot_variant_operator_less_allocs = x.GodotVariantOperatorLess.PassRef()
	allocs57717e51.Borrow(cgodot_variant_operator_less_allocs)

	var cgodot_variant_hash_compare_allocs *cgoAllocMap
	ref57717e51.godot_variant_hash_compare, cgodot_variant_hash_compare_allocs = x.GodotVariantHashCompare.PassRef()
	allocs57717e51.Borrow(cgodot_variant_hash_compare_allocs)

	var cgodot_variant_booleanize_allocs *cgoAllocMap
	ref57717e51.godot_variant_booleanize, cgodot_variant_booleanize_allocs = x.GodotVariantBooleanize.PassRef()
	allocs57717e51.Borrow(cgodot_variant_booleanize_allocs)

	var cgodot_variant_destroy_allocs *cgoAllocMap
	ref57717e51.godot_variant_destroy, cgodot_variant_destroy_allocs = x.GodotVariantDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_variant_destroy_allocs)

	var cgodot_string_new_allocs *cgoAllocMap
	ref57717e51.godot_string_new, cgodot_string_new_allocs = x.GodotStringNew.PassRef()
	allocs57717e51.Borrow(cgodot_string_new_allocs)

	var cgodot_string_new_copy_allocs *cgoAllocMap
	ref57717e51.godot_string_new_copy, cgodot_string_new_copy_allocs = x.GodotStringNewCopy.PassRef()
	allocs57717e51.Borrow(cgodot_string_new_copy_allocs)

	var cgodot_string_new_data_allocs *cgoAllocMap
	ref57717e51.godot_string_new_data, cgodot_string_new_data_allocs = x.GodotStringNewData.PassRef()
	allocs57717e51.Borrow(cgodot_string_new_data_allocs)

	var cgodot_string_new_unicode_data_allocs *cgoAllocMap
	ref57717e51.godot_string_new_unicode_data, cgodot_string_new_unicode_data_allocs = x.GodotStringNewUnicodeData.PassRef()
	allocs57717e51.Borrow(cgodot_string_new_unicode_data_allocs)

	var cgodot_string_get_data_allocs *cgoAllocMap
	ref57717e51.godot_string_get_data, cgodot_string_get_data_allocs = x.GodotStringGetData.PassRef()
	allocs57717e51.Borrow(cgodot_string_get_data_allocs)

	var cgodot_string_operator_index_allocs *cgoAllocMap
	ref57717e51.godot_string_operator_index, cgodot_string_operator_index_allocs = x.GodotStringOperatorIndex.PassRef()
	allocs57717e51.Borrow(cgodot_string_operator_index_allocs)

	var cgodot_string_operator_index_const_allocs *cgoAllocMap
	ref57717e51.godot_string_operator_index_const, cgodot_string_operator_index_const_allocs = x.GodotStringOperatorIndexConst.PassRef()
	allocs57717e51.Borrow(cgodot_string_operator_index_const_allocs)

	var cgodot_string_unicode_str_allocs *cgoAllocMap
	ref57717e51.godot_string_unicode_str, cgodot_string_unicode_str_allocs = x.GodotStringUnicodeStr.PassRef()
	allocs57717e51.Borrow(cgodot_string_unicode_str_allocs)

	var cgodot_string_operator_equal_allocs *cgoAllocMap
	ref57717e51.godot_string_operator_equal, cgodot_string_operator_equal_allocs = x.GodotStringOperatorEqual.PassRef()
	allocs57717e51.Borrow(cgodot_string_operator_equal_allocs)

	var cgodot_string_operator_less_allocs *cgoAllocMap
	ref57717e51.godot_string_operator_less, cgodot_string_operator_less_allocs = x.GodotStringOperatorLess.PassRef()
	allocs57717e51.Borrow(cgodot_string_operator_less_allocs)

	var cgodot_string_operator_plus_allocs *cgoAllocMap
	ref57717e51.godot_string_operator_plus, cgodot_string_operator_plus_allocs = x.GodotStringOperatorPlus.PassRef()
	allocs57717e51.Borrow(cgodot_string_operator_plus_allocs)

	var cgodot_string_length_allocs *cgoAllocMap
	ref57717e51.godot_string_length, cgodot_string_length_allocs = x.GodotStringLength.PassRef()
	allocs57717e51.Borrow(cgodot_string_length_allocs)

	var cgodot_string_begins_with_allocs *cgoAllocMap
	ref57717e51.godot_string_begins_with, cgodot_string_begins_with_allocs = x.GodotStringBeginsWith.PassRef()
	allocs57717e51.Borrow(cgodot_string_begins_with_allocs)

	var cgodot_string_begins_with_char_array_allocs *cgoAllocMap
	ref57717e51.godot_string_begins_with_char_array, cgodot_string_begins_with_char_array_allocs = x.GodotStringBeginsWithCharArray.PassRef()
	allocs57717e51.Borrow(cgodot_string_begins_with_char_array_allocs)

	var cgodot_string_bigrams_allocs *cgoAllocMap
	ref57717e51.godot_string_bigrams, cgodot_string_bigrams_allocs = x.GodotStringBigrams.PassRef()
	allocs57717e51.Borrow(cgodot_string_bigrams_allocs)

	var cgodot_string_chr_allocs *cgoAllocMap
	ref57717e51.godot_string_chr, cgodot_string_chr_allocs = x.GodotStringChr.PassRef()
	allocs57717e51.Borrow(cgodot_string_chr_allocs)

	var cgodot_string_ends_with_allocs *cgoAllocMap
	ref57717e51.godot_string_ends_with, cgodot_string_ends_with_allocs = x.GodotStringEndsWith.PassRef()
	allocs57717e51.Borrow(cgodot_string_ends_with_allocs)

	var cgodot_string_find_allocs *cgoAllocMap
	ref57717e51.godot_string_find, cgodot_string_find_allocs = x.GodotStringFind.PassRef()
	allocs57717e51.Borrow(cgodot_string_find_allocs)

	var cgodot_string_find_from_allocs *cgoAllocMap
	ref57717e51.godot_string_find_from, cgodot_string_find_from_allocs = x.GodotStringFindFrom.PassRef()
	allocs57717e51.Borrow(cgodot_string_find_from_allocs)

	var cgodot_string_findmk_allocs *cgoAllocMap
	ref57717e51.godot_string_findmk, cgodot_string_findmk_allocs = x.GodotStringFindmk.PassRef()
	allocs57717e51.Borrow(cgodot_string_findmk_allocs)

	var cgodot_string_findmk_from_allocs *cgoAllocMap
	ref57717e51.godot_string_findmk_from, cgodot_string_findmk_from_allocs = x.GodotStringFindmkFrom.PassRef()
	allocs57717e51.Borrow(cgodot_string_findmk_from_allocs)

	var cgodot_string_findmk_from_in_place_allocs *cgoAllocMap
	ref57717e51.godot_string_findmk_from_in_place, cgodot_string_findmk_from_in_place_allocs = x.GodotStringFindmkFromInPlace.PassRef()
	allocs57717e51.Borrow(cgodot_string_findmk_from_in_place_allocs)

	var cgodot_string_findn_allocs *cgoAllocMap
	ref57717e51.godot_string_findn, cgodot_string_findn_allocs = x.GodotStringFindn.PassRef()
	allocs57717e51.Borrow(cgodot_string_findn_allocs)

	var cgodot_string_findn_from_allocs *cgoAllocMap
	ref57717e51.godot_string_findn_from, cgodot_string_findn_from_allocs = x.GodotStringFindnFrom.PassRef()
	allocs57717e51.Borrow(cgodot_string_findn_from_allocs)

	var cgodot_string_find_last_allocs *cgoAllocMap
	ref57717e51.godot_string_find_last, cgodot_string_find_last_allocs = x.GodotStringFindLast.PassRef()
	allocs57717e51.Borrow(cgodot_string_find_last_allocs)

	var cgodot_string_format_allocs *cgoAllocMap
	ref57717e51.godot_string_format, cgodot_string_format_allocs = x.GodotStringFormat.PassRef()
	allocs57717e51.Borrow(cgodot_string_format_allocs)

	var cgodot_string_format_with_custom_placeholder_allocs *cgoAllocMap
	ref57717e51.godot_string_format_with_custom_placeholder, cgodot_string_format_with_custom_placeholder_allocs = x.GodotStringFormatWithCustomPlaceholder.PassRef()
	allocs57717e51.Borrow(cgodot_string_format_with_custom_placeholder_allocs)

	var cgodot_string_hex_encode_buffer_allocs *cgoAllocMap
	ref57717e51.godot_string_hex_encode_buffer, cgodot_string_hex_encode_buffer_allocs = x.GodotStringHexEncodeBuffer.PassRef()
	allocs57717e51.Borrow(cgodot_string_hex_encode_buffer_allocs)

	var cgodot_string_hex_to_int_allocs *cgoAllocMap
	ref57717e51.godot_string_hex_to_int, cgodot_string_hex_to_int_allocs = x.GodotStringHexToInt.PassRef()
	allocs57717e51.Borrow(cgodot_string_hex_to_int_allocs)

	var cgodot_string_hex_to_int_without_prefix_allocs *cgoAllocMap
	ref57717e51.godot_string_hex_to_int_without_prefix, cgodot_string_hex_to_int_without_prefix_allocs = x.GodotStringHexToIntWithoutPrefix.PassRef()
	allocs57717e51.Borrow(cgodot_string_hex_to_int_without_prefix_allocs)

	var cgodot_string_insert_allocs *cgoAllocMap
	ref57717e51.godot_string_insert, cgodot_string_insert_allocs = x.GodotStringInsert.PassRef()
	allocs57717e51.Borrow(cgodot_string_insert_allocs)

	var cgodot_string_is_numeric_allocs *cgoAllocMap
	ref57717e51.godot_string_is_numeric, cgodot_string_is_numeric_allocs = x.GodotStringIsNumeric.PassRef()
	allocs57717e51.Borrow(cgodot_string_is_numeric_allocs)

	var cgodot_string_is_subsequence_of_allocs *cgoAllocMap
	ref57717e51.godot_string_is_subsequence_of, cgodot_string_is_subsequence_of_allocs = x.GodotStringIsSubsequenceOf.PassRef()
	allocs57717e51.Borrow(cgodot_string_is_subsequence_of_allocs)

	var cgodot_string_is_subsequence_ofi_allocs *cgoAllocMap
	ref57717e51.godot_string_is_subsequence_ofi, cgodot_string_is_subsequence_ofi_allocs = x.GodotStringIsSubsequenceOfi.PassRef()
	allocs57717e51.Borrow(cgodot_string_is_subsequence_ofi_allocs)

	var cgodot_string_lpad_allocs *cgoAllocMap
	ref57717e51.godot_string_lpad, cgodot_string_lpad_allocs = x.GodotStringLpad.PassRef()
	allocs57717e51.Borrow(cgodot_string_lpad_allocs)

	var cgodot_string_lpad_with_custom_character_allocs *cgoAllocMap
	ref57717e51.godot_string_lpad_with_custom_character, cgodot_string_lpad_with_custom_character_allocs = x.GodotStringLpadWithCustomCharacter.PassRef()
	allocs57717e51.Borrow(cgodot_string_lpad_with_custom_character_allocs)

	var cgodot_string_match_allocs *cgoAllocMap
	ref57717e51.godot_string_match, cgodot_string_match_allocs = x.GodotStringMatch.PassRef()
	allocs57717e51.Borrow(cgodot_string_match_allocs)

	var cgodot_string_matchn_allocs *cgoAllocMap
	ref57717e51.godot_string_matchn, cgodot_string_matchn_allocs = x.GodotStringMatchn.PassRef()
	allocs57717e51.Borrow(cgodot_string_matchn_allocs)

	var cgodot_string_md5_allocs *cgoAllocMap
	ref57717e51.godot_string_md5, cgodot_string_md5_allocs = x.GodotStringMd5.PassRef()
	allocs57717e51.Borrow(cgodot_string_md5_allocs)

	var cgodot_string_num_allocs *cgoAllocMap
	ref57717e51.godot_string_num, cgodot_string_num_allocs = x.GodotStringNum.PassRef()
	allocs57717e51.Borrow(cgodot_string_num_allocs)

	var cgodot_string_num_int64_allocs *cgoAllocMap
	ref57717e51.godot_string_num_int64, cgodot_string_num_int64_allocs = x.GodotStringNumInt64.PassRef()
	allocs57717e51.Borrow(cgodot_string_num_int64_allocs)

	var cgodot_string_num_int64_capitalized_allocs *cgoAllocMap
	ref57717e51.godot_string_num_int64_capitalized, cgodot_string_num_int64_capitalized_allocs = x.GodotStringNumInt64Capitalized.PassRef()
	allocs57717e51.Borrow(cgodot_string_num_int64_capitalized_allocs)

	var cgodot_string_num_real_allocs *cgoAllocMap
	ref57717e51.godot_string_num_real, cgodot_string_num_real_allocs = x.GodotStringNumReal.PassRef()
	allocs57717e51.Borrow(cgodot_string_num_real_allocs)

	var cgodot_string_num_scientific_allocs *cgoAllocMap
	ref57717e51.godot_string_num_scientific, cgodot_string_num_scientific_allocs = x.GodotStringNumScientific.PassRef()
	allocs57717e51.Borrow(cgodot_string_num_scientific_allocs)

	var cgodot_string_num_with_decimals_allocs *cgoAllocMap
	ref57717e51.godot_string_num_with_decimals, cgodot_string_num_with_decimals_allocs = x.GodotStringNumWithDecimals.PassRef()
	allocs57717e51.Borrow(cgodot_string_num_with_decimals_allocs)

	var cgodot_string_pad_decimals_allocs *cgoAllocMap
	ref57717e51.godot_string_pad_decimals, cgodot_string_pad_decimals_allocs = x.GodotStringPadDecimals.PassRef()
	allocs57717e51.Borrow(cgodot_string_pad_decimals_allocs)

	var cgodot_string_pad_zeros_allocs *cgoAllocMap
	ref57717e51.godot_string_pad_zeros, cgodot_string_pad_zeros_allocs = x.GodotStringPadZeros.PassRef()
	allocs57717e51.Borrow(cgodot_string_pad_zeros_allocs)

	var cgodot_string_replace_first_allocs *cgoAllocMap
	ref57717e51.godot_string_replace_first, cgodot_string_replace_first_allocs = x.GodotStringReplaceFirst.PassRef()
	allocs57717e51.Borrow(cgodot_string_replace_first_allocs)

	var cgodot_string_replace_allocs *cgoAllocMap
	ref57717e51.godot_string_replace, cgodot_string_replace_allocs = x.GodotStringReplace.PassRef()
	allocs57717e51.Borrow(cgodot_string_replace_allocs)

	var cgodot_string_replacen_allocs *cgoAllocMap
	ref57717e51.godot_string_replacen, cgodot_string_replacen_allocs = x.GodotStringReplacen.PassRef()
	allocs57717e51.Borrow(cgodot_string_replacen_allocs)

	var cgodot_string_rfind_allocs *cgoAllocMap
	ref57717e51.godot_string_rfind, cgodot_string_rfind_allocs = x.GodotStringRfind.PassRef()
	allocs57717e51.Borrow(cgodot_string_rfind_allocs)

	var cgodot_string_rfindn_allocs *cgoAllocMap
	ref57717e51.godot_string_rfindn, cgodot_string_rfindn_allocs = x.GodotStringRfindn.PassRef()
	allocs57717e51.Borrow(cgodot_string_rfindn_allocs)

	var cgodot_string_rfind_from_allocs *cgoAllocMap
	ref57717e51.godot_string_rfind_from, cgodot_string_rfind_from_allocs = x.GodotStringRfindFrom.PassRef()
	allocs57717e51.Borrow(cgodot_string_rfind_from_allocs)

	var cgodot_string_rfindn_from_allocs *cgoAllocMap
	ref57717e51.godot_string_rfindn_from, cgodot_string_rfindn_from_allocs = x.GodotStringRfindnFrom.PassRef()
	allocs57717e51.Borrow(cgodot_string_rfindn_from_allocs)

	var cgodot_string_rpad_allocs *cgoAllocMap
	ref57717e51.godot_string_rpad, cgodot_string_rpad_allocs = x.GodotStringRpad.PassRef()
	allocs57717e51.Borrow(cgodot_string_rpad_allocs)

	var cgodot_string_rpad_with_custom_character_allocs *cgoAllocMap
	ref57717e51.godot_string_rpad_with_custom_character, cgodot_string_rpad_with_custom_character_allocs = x.GodotStringRpadWithCustomCharacter.PassRef()
	allocs57717e51.Borrow(cgodot_string_rpad_with_custom_character_allocs)

	var cgodot_string_similarity_allocs *cgoAllocMap
	ref57717e51.godot_string_similarity, cgodot_string_similarity_allocs = x.GodotStringSimilarity.PassRef()
	allocs57717e51.Borrow(cgodot_string_similarity_allocs)

	var cgodot_string_sprintf_allocs *cgoAllocMap
	ref57717e51.godot_string_sprintf, cgodot_string_sprintf_allocs = x.GodotStringSprintf.PassRef()
	allocs57717e51.Borrow(cgodot_string_sprintf_allocs)

	var cgodot_string_substr_allocs *cgoAllocMap
	ref57717e51.godot_string_substr, cgodot_string_substr_allocs = x.GodotStringSubstr.PassRef()
	allocs57717e51.Borrow(cgodot_string_substr_allocs)

	var cgodot_string_to_double_allocs *cgoAllocMap
	ref57717e51.godot_string_to_double, cgodot_string_to_double_allocs = x.GodotStringToDouble.PassRef()
	allocs57717e51.Borrow(cgodot_string_to_double_allocs)

	var cgodot_string_to_float_allocs *cgoAllocMap
	ref57717e51.godot_string_to_float, cgodot_string_to_float_allocs = x.GodotStringToFloat.PassRef()
	allocs57717e51.Borrow(cgodot_string_to_float_allocs)

	var cgodot_string_to_int_allocs *cgoAllocMap
	ref57717e51.godot_string_to_int, cgodot_string_to_int_allocs = x.GodotStringToInt.PassRef()
	allocs57717e51.Borrow(cgodot_string_to_int_allocs)

	var cgodot_string_camelcase_to_underscore_allocs *cgoAllocMap
	ref57717e51.godot_string_camelcase_to_underscore, cgodot_string_camelcase_to_underscore_allocs = x.GodotStringCamelcaseToUnderscore.PassRef()
	allocs57717e51.Borrow(cgodot_string_camelcase_to_underscore_allocs)

	var cgodot_string_camelcase_to_underscore_lowercased_allocs *cgoAllocMap
	ref57717e51.godot_string_camelcase_to_underscore_lowercased, cgodot_string_camelcase_to_underscore_lowercased_allocs = x.GodotStringCamelcaseToUnderscoreLowercased.PassRef()
	allocs57717e51.Borrow(cgodot_string_camelcase_to_underscore_lowercased_allocs)

	var cgodot_string_capitalize_allocs *cgoAllocMap
	ref57717e51.godot_string_capitalize, cgodot_string_capitalize_allocs = x.GodotStringCapitalize.PassRef()
	allocs57717e51.Borrow(cgodot_string_capitalize_allocs)

	var cgodot_string_char_to_double_allocs *cgoAllocMap
	ref57717e51.godot_string_char_to_double, cgodot_string_char_to_double_allocs = x.GodotStringCharToDouble.PassRef()
	allocs57717e51.Borrow(cgodot_string_char_to_double_allocs)

	var cgodot_string_char_to_int_allocs *cgoAllocMap
	ref57717e51.godot_string_char_to_int, cgodot_string_char_to_int_allocs = x.GodotStringCharToInt.PassRef()
	allocs57717e51.Borrow(cgodot_string_char_to_int_allocs)

	var cgodot_string_wchar_to_int_allocs *cgoAllocMap
	ref57717e51.godot_string_wchar_to_int, cgodot_string_wchar_to_int_allocs = x.GodotStringWcharToInt.PassRef()
	allocs57717e51.Borrow(cgodot_string_wchar_to_int_allocs)

	var cgodot_string_char_to_int_with_len_allocs *cgoAllocMap
	ref57717e51.godot_string_char_to_int_with_len, cgodot_string_char_to_int_with_len_allocs = x.GodotStringCharToIntWithLen.PassRef()
	allocs57717e51.Borrow(cgodot_string_char_to_int_with_len_allocs)

	var cgodot_string_char_to_int64_with_len_allocs *cgoAllocMap
	ref57717e51.godot_string_char_to_int64_with_len, cgodot_string_char_to_int64_with_len_allocs = x.GodotStringCharToInt64WithLen.PassRef()
	allocs57717e51.Borrow(cgodot_string_char_to_int64_with_len_allocs)

	var cgodot_string_hex_to_int64_allocs *cgoAllocMap
	ref57717e51.godot_string_hex_to_int64, cgodot_string_hex_to_int64_allocs = x.GodotStringHexToInt64.PassRef()
	allocs57717e51.Borrow(cgodot_string_hex_to_int64_allocs)

	var cgodot_string_hex_to_int64_with_prefix_allocs *cgoAllocMap
	ref57717e51.godot_string_hex_to_int64_with_prefix, cgodot_string_hex_to_int64_with_prefix_allocs = x.GodotStringHexToInt64WithPrefix.PassRef()
	allocs57717e51.Borrow(cgodot_string_hex_to_int64_with_prefix_allocs)

	var cgodot_string_to_int64_allocs *cgoAllocMap
	ref57717e51.godot_string_to_int64, cgodot_string_to_int64_allocs = x.GodotStringToInt64.PassRef()
	allocs57717e51.Borrow(cgodot_string_to_int64_allocs)

	var cgodot_string_unicode_char_to_double_allocs *cgoAllocMap
	ref57717e51.godot_string_unicode_char_to_double, cgodot_string_unicode_char_to_double_allocs = x.GodotStringUnicodeCharToDouble.PassRef()
	allocs57717e51.Borrow(cgodot_string_unicode_char_to_double_allocs)

	var cgodot_string_get_slice_count_allocs *cgoAllocMap
	ref57717e51.godot_string_get_slice_count, cgodot_string_get_slice_count_allocs = x.GodotStringGetSliceCount.PassRef()
	allocs57717e51.Borrow(cgodot_string_get_slice_count_allocs)

	var cgodot_string_get_slice_allocs *cgoAllocMap
	ref57717e51.godot_string_get_slice, cgodot_string_get_slice_allocs = x.GodotStringGetSlice.PassRef()
	allocs57717e51.Borrow(cgodot_string_get_slice_allocs)

	var cgodot_string_get_slicec_allocs *cgoAllocMap
	ref57717e51.godot_string_get_slicec, cgodot_string_get_slicec_allocs = x.GodotStringGetSlicec.PassRef()
	allocs57717e51.Borrow(cgodot_string_get_slicec_allocs)

	var cgodot_string_split_allocs *cgoAllocMap
	ref57717e51.godot_string_split, cgodot_string_split_allocs = x.GodotStringSplit.PassRef()
	allocs57717e51.Borrow(cgodot_string_split_allocs)

	var cgodot_string_split_allow_empty_allocs *cgoAllocMap
	ref57717e51.godot_string_split_allow_empty, cgodot_string_split_allow_empty_allocs = x.GodotStringSplitAllowEmpty.PassRef()
	allocs57717e51.Borrow(cgodot_string_split_allow_empty_allocs)

	var cgodot_string_split_floats_allocs *cgoAllocMap
	ref57717e51.godot_string_split_floats, cgodot_string_split_floats_allocs = x.GodotStringSplitFloats.PassRef()
	allocs57717e51.Borrow(cgodot_string_split_floats_allocs)

	var cgodot_string_split_floats_allows_empty_allocs *cgoAllocMap
	ref57717e51.godot_string_split_floats_allows_empty, cgodot_string_split_floats_allows_empty_allocs = x.GodotStringSplitFloatsAllowsEmpty.PassRef()
	allocs57717e51.Borrow(cgodot_string_split_floats_allows_empty_allocs)

	var cgodot_string_split_floats_mk_allocs *cgoAllocMap
	ref57717e51.godot_string_split_floats_mk, cgodot_string_split_floats_mk_allocs = x.GodotStringSplitFloatsMk.PassRef()
	allocs57717e51.Borrow(cgodot_string_split_floats_mk_allocs)

	var cgodot_string_split_floats_mk_allows_empty_allocs *cgoAllocMap
	ref57717e51.godot_string_split_floats_mk_allows_empty, cgodot_string_split_floats_mk_allows_empty_allocs = x.GodotStringSplitFloatsMkAllowsEmpty.PassRef()
	allocs57717e51.Borrow(cgodot_string_split_floats_mk_allows_empty_allocs)

	var cgodot_string_split_ints_allocs *cgoAllocMap
	ref57717e51.godot_string_split_ints, cgodot_string_split_ints_allocs = x.GodotStringSplitInts.PassRef()
	allocs57717e51.Borrow(cgodot_string_split_ints_allocs)

	var cgodot_string_split_ints_allows_empty_allocs *cgoAllocMap
	ref57717e51.godot_string_split_ints_allows_empty, cgodot_string_split_ints_allows_empty_allocs = x.GodotStringSplitIntsAllowsEmpty.PassRef()
	allocs57717e51.Borrow(cgodot_string_split_ints_allows_empty_allocs)

	var cgodot_string_split_ints_mk_allocs *cgoAllocMap
	ref57717e51.godot_string_split_ints_mk, cgodot_string_split_ints_mk_allocs = x.GodotStringSplitIntsMk.PassRef()
	allocs57717e51.Borrow(cgodot_string_split_ints_mk_allocs)

	var cgodot_string_split_ints_mk_allows_empty_allocs *cgoAllocMap
	ref57717e51.godot_string_split_ints_mk_allows_empty, cgodot_string_split_ints_mk_allows_empty_allocs = x.GodotStringSplitIntsMkAllowsEmpty.PassRef()
	allocs57717e51.Borrow(cgodot_string_split_ints_mk_allows_empty_allocs)

	var cgodot_string_split_spaces_allocs *cgoAllocMap
	ref57717e51.godot_string_split_spaces, cgodot_string_split_spaces_allocs = x.GodotStringSplitSpaces.PassRef()
	allocs57717e51.Borrow(cgodot_string_split_spaces_allocs)

	var cgodot_string_char_lowercase_allocs *cgoAllocMap
	ref57717e51.godot_string_char_lowercase, cgodot_string_char_lowercase_allocs = x.GodotStringCharLowercase.PassRef()
	allocs57717e51.Borrow(cgodot_string_char_lowercase_allocs)

	var cgodot_string_char_uppercase_allocs *cgoAllocMap
	ref57717e51.godot_string_char_uppercase, cgodot_string_char_uppercase_allocs = x.GodotStringCharUppercase.PassRef()
	allocs57717e51.Borrow(cgodot_string_char_uppercase_allocs)

	var cgodot_string_to_lower_allocs *cgoAllocMap
	ref57717e51.godot_string_to_lower, cgodot_string_to_lower_allocs = x.GodotStringToLower.PassRef()
	allocs57717e51.Borrow(cgodot_string_to_lower_allocs)

	var cgodot_string_to_upper_allocs *cgoAllocMap
	ref57717e51.godot_string_to_upper, cgodot_string_to_upper_allocs = x.GodotStringToUpper.PassRef()
	allocs57717e51.Borrow(cgodot_string_to_upper_allocs)

	var cgodot_string_get_basename_allocs *cgoAllocMap
	ref57717e51.godot_string_get_basename, cgodot_string_get_basename_allocs = x.GodotStringGetBasename.PassRef()
	allocs57717e51.Borrow(cgodot_string_get_basename_allocs)

	var cgodot_string_get_extension_allocs *cgoAllocMap
	ref57717e51.godot_string_get_extension, cgodot_string_get_extension_allocs = x.GodotStringGetExtension.PassRef()
	allocs57717e51.Borrow(cgodot_string_get_extension_allocs)

	var cgodot_string_left_allocs *cgoAllocMap
	ref57717e51.godot_string_left, cgodot_string_left_allocs = x.GodotStringLeft.PassRef()
	allocs57717e51.Borrow(cgodot_string_left_allocs)

	var cgodot_string_ord_at_allocs *cgoAllocMap
	ref57717e51.godot_string_ord_at, cgodot_string_ord_at_allocs = x.GodotStringOrdAt.PassRef()
	allocs57717e51.Borrow(cgodot_string_ord_at_allocs)

	var cgodot_string_plus_file_allocs *cgoAllocMap
	ref57717e51.godot_string_plus_file, cgodot_string_plus_file_allocs = x.GodotStringPlusFile.PassRef()
	allocs57717e51.Borrow(cgodot_string_plus_file_allocs)

	var cgodot_string_right_allocs *cgoAllocMap
	ref57717e51.godot_string_right, cgodot_string_right_allocs = x.GodotStringRight.PassRef()
	allocs57717e51.Borrow(cgodot_string_right_allocs)

	var cgodot_string_strip_edges_allocs *cgoAllocMap
	ref57717e51.godot_string_strip_edges, cgodot_string_strip_edges_allocs = x.GodotStringStripEdges.PassRef()
	allocs57717e51.Borrow(cgodot_string_strip_edges_allocs)

	var cgodot_string_strip_escapes_allocs *cgoAllocMap
	ref57717e51.godot_string_strip_escapes, cgodot_string_strip_escapes_allocs = x.GodotStringStripEscapes.PassRef()
	allocs57717e51.Borrow(cgodot_string_strip_escapes_allocs)

	var cgodot_string_erase_allocs *cgoAllocMap
	ref57717e51.godot_string_erase, cgodot_string_erase_allocs = x.GodotStringErase.PassRef()
	allocs57717e51.Borrow(cgodot_string_erase_allocs)

	var cgodot_string_ascii_allocs *cgoAllocMap
	ref57717e51.godot_string_ascii, cgodot_string_ascii_allocs = x.GodotStringAscii.PassRef()
	allocs57717e51.Borrow(cgodot_string_ascii_allocs)

	var cgodot_string_ascii_extended_allocs *cgoAllocMap
	ref57717e51.godot_string_ascii_extended, cgodot_string_ascii_extended_allocs = x.GodotStringAsciiExtended.PassRef()
	allocs57717e51.Borrow(cgodot_string_ascii_extended_allocs)

	var cgodot_string_utf8_allocs *cgoAllocMap
	ref57717e51.godot_string_utf8, cgodot_string_utf8_allocs = x.GodotStringUtf8.PassRef()
	allocs57717e51.Borrow(cgodot_string_utf8_allocs)

	var cgodot_string_parse_utf8_allocs *cgoAllocMap
	ref57717e51.godot_string_parse_utf8, cgodot_string_parse_utf8_allocs = x.GodotStringParseUtf8.PassRef()
	allocs57717e51.Borrow(cgodot_string_parse_utf8_allocs)

	var cgodot_string_parse_utf8_with_len_allocs *cgoAllocMap
	ref57717e51.godot_string_parse_utf8_with_len, cgodot_string_parse_utf8_with_len_allocs = x.GodotStringParseUtf8WithLen.PassRef()
	allocs57717e51.Borrow(cgodot_string_parse_utf8_with_len_allocs)

	var cgodot_string_chars_to_utf8_allocs *cgoAllocMap
	ref57717e51.godot_string_chars_to_utf8, cgodot_string_chars_to_utf8_allocs = x.GodotStringCharsToUtf8.PassRef()
	allocs57717e51.Borrow(cgodot_string_chars_to_utf8_allocs)

	var cgodot_string_chars_to_utf8_with_len_allocs *cgoAllocMap
	ref57717e51.godot_string_chars_to_utf8_with_len, cgodot_string_chars_to_utf8_with_len_allocs = x.GodotStringCharsToUtf8WithLen.PassRef()
	allocs57717e51.Borrow(cgodot_string_chars_to_utf8_with_len_allocs)

	var cgodot_string_hash_allocs *cgoAllocMap
	ref57717e51.godot_string_hash, cgodot_string_hash_allocs = x.GodotStringHash.PassRef()
	allocs57717e51.Borrow(cgodot_string_hash_allocs)

	var cgodot_string_hash64_allocs *cgoAllocMap
	ref57717e51.godot_string_hash64, cgodot_string_hash64_allocs = x.GodotStringHash64.PassRef()
	allocs57717e51.Borrow(cgodot_string_hash64_allocs)

	var cgodot_string_hash_chars_allocs *cgoAllocMap
	ref57717e51.godot_string_hash_chars, cgodot_string_hash_chars_allocs = x.GodotStringHashChars.PassRef()
	allocs57717e51.Borrow(cgodot_string_hash_chars_allocs)

	var cgodot_string_hash_chars_with_len_allocs *cgoAllocMap
	ref57717e51.godot_string_hash_chars_with_len, cgodot_string_hash_chars_with_len_allocs = x.GodotStringHashCharsWithLen.PassRef()
	allocs57717e51.Borrow(cgodot_string_hash_chars_with_len_allocs)

	var cgodot_string_hash_utf8_chars_allocs *cgoAllocMap
	ref57717e51.godot_string_hash_utf8_chars, cgodot_string_hash_utf8_chars_allocs = x.GodotStringHashUtf8Chars.PassRef()
	allocs57717e51.Borrow(cgodot_string_hash_utf8_chars_allocs)

	var cgodot_string_hash_utf8_chars_with_len_allocs *cgoAllocMap
	ref57717e51.godot_string_hash_utf8_chars_with_len, cgodot_string_hash_utf8_chars_with_len_allocs = x.GodotStringHashUtf8CharsWithLen.PassRef()
	allocs57717e51.Borrow(cgodot_string_hash_utf8_chars_with_len_allocs)

	var cgodot_string_md5_buffer_allocs *cgoAllocMap
	ref57717e51.godot_string_md5_buffer, cgodot_string_md5_buffer_allocs = x.GodotStringMd5Buffer.PassRef()
	allocs57717e51.Borrow(cgodot_string_md5_buffer_allocs)

	var cgodot_string_md5_text_allocs *cgoAllocMap
	ref57717e51.godot_string_md5_text, cgodot_string_md5_text_allocs = x.GodotStringMd5Text.PassRef()
	allocs57717e51.Borrow(cgodot_string_md5_text_allocs)

	var cgodot_string_sha256_buffer_allocs *cgoAllocMap
	ref57717e51.godot_string_sha256_buffer, cgodot_string_sha256_buffer_allocs = x.GodotStringSha256Buffer.PassRef()
	allocs57717e51.Borrow(cgodot_string_sha256_buffer_allocs)

	var cgodot_string_sha256_text_allocs *cgoAllocMap
	ref57717e51.godot_string_sha256_text, cgodot_string_sha256_text_allocs = x.GodotStringSha256Text.PassRef()
	allocs57717e51.Borrow(cgodot_string_sha256_text_allocs)

	var cgodot_string_empty_allocs *cgoAllocMap
	ref57717e51.godot_string_empty, cgodot_string_empty_allocs = x.GodotStringEmpty.PassRef()
	allocs57717e51.Borrow(cgodot_string_empty_allocs)

	var cgodot_string_get_base_dir_allocs *cgoAllocMap
	ref57717e51.godot_string_get_base_dir, cgodot_string_get_base_dir_allocs = x.GodotStringGetBaseDir.PassRef()
	allocs57717e51.Borrow(cgodot_string_get_base_dir_allocs)

	var cgodot_string_get_file_allocs *cgoAllocMap
	ref57717e51.godot_string_get_file, cgodot_string_get_file_allocs = x.GodotStringGetFile.PassRef()
	allocs57717e51.Borrow(cgodot_string_get_file_allocs)

	var cgodot_string_humanize_size_allocs *cgoAllocMap
	ref57717e51.godot_string_humanize_size, cgodot_string_humanize_size_allocs = x.GodotStringHumanizeSize.PassRef()
	allocs57717e51.Borrow(cgodot_string_humanize_size_allocs)

	var cgodot_string_is_abs_path_allocs *cgoAllocMap
	ref57717e51.godot_string_is_abs_path, cgodot_string_is_abs_path_allocs = x.GodotStringIsAbsPath.PassRef()
	allocs57717e51.Borrow(cgodot_string_is_abs_path_allocs)

	var cgodot_string_is_rel_path_allocs *cgoAllocMap
	ref57717e51.godot_string_is_rel_path, cgodot_string_is_rel_path_allocs = x.GodotStringIsRelPath.PassRef()
	allocs57717e51.Borrow(cgodot_string_is_rel_path_allocs)

	var cgodot_string_is_resource_file_allocs *cgoAllocMap
	ref57717e51.godot_string_is_resource_file, cgodot_string_is_resource_file_allocs = x.GodotStringIsResourceFile.PassRef()
	allocs57717e51.Borrow(cgodot_string_is_resource_file_allocs)

	var cgodot_string_path_to_allocs *cgoAllocMap
	ref57717e51.godot_string_path_to, cgodot_string_path_to_allocs = x.GodotStringPathTo.PassRef()
	allocs57717e51.Borrow(cgodot_string_path_to_allocs)

	var cgodot_string_path_to_file_allocs *cgoAllocMap
	ref57717e51.godot_string_path_to_file, cgodot_string_path_to_file_allocs = x.GodotStringPathToFile.PassRef()
	allocs57717e51.Borrow(cgodot_string_path_to_file_allocs)

	var cgodot_string_simplify_path_allocs *cgoAllocMap
	ref57717e51.godot_string_simplify_path, cgodot_string_simplify_path_allocs = x.GodotStringSimplifyPath.PassRef()
	allocs57717e51.Borrow(cgodot_string_simplify_path_allocs)

	var cgodot_string_c_escape_allocs *cgoAllocMap
	ref57717e51.godot_string_c_escape, cgodot_string_c_escape_allocs = x.GodotStringCEscape.PassRef()
	allocs57717e51.Borrow(cgodot_string_c_escape_allocs)

	var cgodot_string_c_escape_multiline_allocs *cgoAllocMap
	ref57717e51.godot_string_c_escape_multiline, cgodot_string_c_escape_multiline_allocs = x.GodotStringCEscapeMultiline.PassRef()
	allocs57717e51.Borrow(cgodot_string_c_escape_multiline_allocs)

	var cgodot_string_c_unescape_allocs *cgoAllocMap
	ref57717e51.godot_string_c_unescape, cgodot_string_c_unescape_allocs = x.GodotStringCUnescape.PassRef()
	allocs57717e51.Borrow(cgodot_string_c_unescape_allocs)

	var cgodot_string_http_escape_allocs *cgoAllocMap
	ref57717e51.godot_string_http_escape, cgodot_string_http_escape_allocs = x.GodotStringHttpEscape.PassRef()
	allocs57717e51.Borrow(cgodot_string_http_escape_allocs)

	var cgodot_string_http_unescape_allocs *cgoAllocMap
	ref57717e51.godot_string_http_unescape, cgodot_string_http_unescape_allocs = x.GodotStringHttpUnescape.PassRef()
	allocs57717e51.Borrow(cgodot_string_http_unescape_allocs)

	var cgodot_string_json_escape_allocs *cgoAllocMap
	ref57717e51.godot_string_json_escape, cgodot_string_json_escape_allocs = x.GodotStringJsonEscape.PassRef()
	allocs57717e51.Borrow(cgodot_string_json_escape_allocs)

	var cgodot_string_word_wrap_allocs *cgoAllocMap
	ref57717e51.godot_string_word_wrap, cgodot_string_word_wrap_allocs = x.GodotStringWordWrap.PassRef()
	allocs57717e51.Borrow(cgodot_string_word_wrap_allocs)

	var cgodot_string_xml_escape_allocs *cgoAllocMap
	ref57717e51.godot_string_xml_escape, cgodot_string_xml_escape_allocs = x.GodotStringXmlEscape.PassRef()
	allocs57717e51.Borrow(cgodot_string_xml_escape_allocs)

	var cgodot_string_xml_escape_with_quotes_allocs *cgoAllocMap
	ref57717e51.godot_string_xml_escape_with_quotes, cgodot_string_xml_escape_with_quotes_allocs = x.GodotStringXmlEscapeWithQuotes.PassRef()
	allocs57717e51.Borrow(cgodot_string_xml_escape_with_quotes_allocs)

	var cgodot_string_xml_unescape_allocs *cgoAllocMap
	ref57717e51.godot_string_xml_unescape, cgodot_string_xml_unescape_allocs = x.GodotStringXmlUnescape.PassRef()
	allocs57717e51.Borrow(cgodot_string_xml_unescape_allocs)

	var cgodot_string_percent_decode_allocs *cgoAllocMap
	ref57717e51.godot_string_percent_decode, cgodot_string_percent_decode_allocs = x.GodotStringPercentDecode.PassRef()
	allocs57717e51.Borrow(cgodot_string_percent_decode_allocs)

	var cgodot_string_percent_encode_allocs *cgoAllocMap
	ref57717e51.godot_string_percent_encode, cgodot_string_percent_encode_allocs = x.GodotStringPercentEncode.PassRef()
	allocs57717e51.Borrow(cgodot_string_percent_encode_allocs)

	var cgodot_string_is_valid_float_allocs *cgoAllocMap
	ref57717e51.godot_string_is_valid_float, cgodot_string_is_valid_float_allocs = x.GodotStringIsValidFloat.PassRef()
	allocs57717e51.Borrow(cgodot_string_is_valid_float_allocs)

	var cgodot_string_is_valid_hex_number_allocs *cgoAllocMap
	ref57717e51.godot_string_is_valid_hex_number, cgodot_string_is_valid_hex_number_allocs = x.GodotStringIsValidHexNumber.PassRef()
	allocs57717e51.Borrow(cgodot_string_is_valid_hex_number_allocs)

	var cgodot_string_is_valid_html_color_allocs *cgoAllocMap
	ref57717e51.godot_string_is_valid_html_color, cgodot_string_is_valid_html_color_allocs = x.GodotStringIsValidHtmlColor.PassRef()
	allocs57717e51.Borrow(cgodot_string_is_valid_html_color_allocs)

	var cgodot_string_is_valid_identifier_allocs *cgoAllocMap
	ref57717e51.godot_string_is_valid_identifier, cgodot_string_is_valid_identifier_allocs = x.GodotStringIsValidIdentifier.PassRef()
	allocs57717e51.Borrow(cgodot_string_is_valid_identifier_allocs)

	var cgodot_string_is_valid_integer_allocs *cgoAllocMap
	ref57717e51.godot_string_is_valid_integer, cgodot_string_is_valid_integer_allocs = x.GodotStringIsValidInteger.PassRef()
	allocs57717e51.Borrow(cgodot_string_is_valid_integer_allocs)

	var cgodot_string_is_valid_ip_address_allocs *cgoAllocMap
	ref57717e51.godot_string_is_valid_ip_address, cgodot_string_is_valid_ip_address_allocs = x.GodotStringIsValidIpAddress.PassRef()
	allocs57717e51.Borrow(cgodot_string_is_valid_ip_address_allocs)

	var cgodot_string_destroy_allocs *cgoAllocMap
	ref57717e51.godot_string_destroy, cgodot_string_destroy_allocs = x.GodotStringDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_string_destroy_allocs)

	var cgodot_string_name_new_allocs *cgoAllocMap
	ref57717e51.godot_string_name_new, cgodot_string_name_new_allocs = x.GodotStringNameNew.PassRef()
	allocs57717e51.Borrow(cgodot_string_name_new_allocs)

	var cgodot_string_name_new_data_allocs *cgoAllocMap
	ref57717e51.godot_string_name_new_data, cgodot_string_name_new_data_allocs = x.GodotStringNameNewData.PassRef()
	allocs57717e51.Borrow(cgodot_string_name_new_data_allocs)

	var cgodot_string_name_get_name_allocs *cgoAllocMap
	ref57717e51.godot_string_name_get_name, cgodot_string_name_get_name_allocs = x.GodotStringNameGetName.PassRef()
	allocs57717e51.Borrow(cgodot_string_name_get_name_allocs)

	var cgodot_string_name_get_hash_allocs *cgoAllocMap
	ref57717e51.godot_string_name_get_hash, cgodot_string_name_get_hash_allocs = x.GodotStringNameGetHash.PassRef()
	allocs57717e51.Borrow(cgodot_string_name_get_hash_allocs)

	var cgodot_string_name_get_data_unique_pointer_allocs *cgoAllocMap
	ref57717e51.godot_string_name_get_data_unique_pointer, cgodot_string_name_get_data_unique_pointer_allocs = x.GodotStringNameGetDataUniquePointer.PassRef()
	allocs57717e51.Borrow(cgodot_string_name_get_data_unique_pointer_allocs)

	var cgodot_string_name_operator_equal_allocs *cgoAllocMap
	ref57717e51.godot_string_name_operator_equal, cgodot_string_name_operator_equal_allocs = x.GodotStringNameOperatorEqual.PassRef()
	allocs57717e51.Borrow(cgodot_string_name_operator_equal_allocs)

	var cgodot_string_name_operator_less_allocs *cgoAllocMap
	ref57717e51.godot_string_name_operator_less, cgodot_string_name_operator_less_allocs = x.GodotStringNameOperatorLess.PassRef()
	allocs57717e51.Borrow(cgodot_string_name_operator_less_allocs)

	var cgodot_string_name_destroy_allocs *cgoAllocMap
	ref57717e51.godot_string_name_destroy, cgodot_string_name_destroy_allocs = x.GodotStringNameDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_string_name_destroy_allocs)

	var cgodot_object_destroy_allocs *cgoAllocMap
	ref57717e51.godot_object_destroy, cgodot_object_destroy_allocs = x.GodotObjectDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_object_destroy_allocs)

	var cgodot_global_get_singleton_allocs *cgoAllocMap
	ref57717e51.godot_global_get_singleton, cgodot_global_get_singleton_allocs = x.GodotGlobalGetSingleton.PassRef()
	allocs57717e51.Borrow(cgodot_global_get_singleton_allocs)

	var cgodot_method_bind_get_method_allocs *cgoAllocMap
	ref57717e51.godot_method_bind_get_method, cgodot_method_bind_get_method_allocs = x.GodotMethodBindGetMethod.PassRef()
	allocs57717e51.Borrow(cgodot_method_bind_get_method_allocs)

	var cgodot_method_bind_ptrcall_allocs *cgoAllocMap
	ref57717e51.godot_method_bind_ptrcall, cgodot_method_bind_ptrcall_allocs = x.GodotMethodBindPtrcall.PassRef()
	allocs57717e51.Borrow(cgodot_method_bind_ptrcall_allocs)

	var cgodot_method_bind_call_allocs *cgoAllocMap
	ref57717e51.godot_method_bind_call, cgodot_method_bind_call_allocs = x.GodotMethodBindCall.PassRef()
	allocs57717e51.Borrow(cgodot_method_bind_call_allocs)

	var cgodot_get_class_constructor_allocs *cgoAllocMap
	ref57717e51.godot_get_class_constructor, cgodot_get_class_constructor_allocs = x.GodotGetClassConstructor.PassValue()
	allocs57717e51.Borrow(cgodot_get_class_constructor_allocs)

	var cgodot_register_native_call_type_allocs *cgoAllocMap
	ref57717e51.godot_register_native_call_type, cgodot_register_native_call_type_allocs = x.GodotRegisterNativeCallType.PassRef()
	allocs57717e51.Borrow(cgodot_register_native_call_type_allocs)

	var cgodot_alloc_allocs *cgoAllocMap
	ref57717e51.godot_alloc, cgodot_alloc_allocs = x.GodotAlloc.PassRef()
	allocs57717e51.Borrow(cgodot_alloc_allocs)

	var cgodot_realloc_allocs *cgoAllocMap
	ref57717e51.godot_realloc, cgodot_realloc_allocs = x.GodotRealloc.PassRef()
	allocs57717e51.Borrow(cgodot_realloc_allocs)

	var cgodot_free_allocs *cgoAllocMap
	ref57717e51.godot_free, cgodot_free_allocs = x.GodotFree.PassRef()
	allocs57717e51.Borrow(cgodot_free_allocs)

	var cgodot_print_error_allocs *cgoAllocMap
	ref57717e51.godot_print_error, cgodot_print_error_allocs = x.GodotPrintError.PassRef()
	allocs57717e51.Borrow(cgodot_print_error_allocs)

	var cgodot_print_warning_allocs *cgoAllocMap
	ref57717e51.godot_print_warning, cgodot_print_warning_allocs = x.GodotPrintWarning.PassRef()
	allocs57717e51.Borrow(cgodot_print_warning_allocs)

	var cgodot_print_allocs *cgoAllocMap
	ref57717e51.godot_print, cgodot_print_allocs = x.GodotPrint.PassRef()
	allocs57717e51.Borrow(cgodot_print_allocs)

	x.ref57717e51 = ref57717e51
	x.allocs57717e51 = allocs57717e51
	return ref57717e51, allocs57717e51

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotGdnativeCoreApiStruct) PassValue() (C.godot_gdnative_core_api_struct, *cgoAllocMap) {
	if x.ref57717e51 != nil {
		return *x.ref57717e51, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotGdnativeCoreApiStruct) Deref() {
	if x.ref57717e51 == nil {
		return
	}
	x.Type = (uint32)(x.ref57717e51._type)
	x.Version = *NewGodotGdnativeApiVersionRef(unsafe.Pointer(&x.ref57717e51.version))
	packSGodotGdnativeApiStruct(x.Next, x.ref57717e51.next)
	x.NumExtensions = (uint32)(x.ref57717e51.num_extensions)
	packSSGodotGdnativeApiStruct(x.Extensions, x.ref57717e51.extensions)
	x.GodotColorNewRgba = NewRef(unsafe.Pointer(x.ref57717e51.godot_color_new_rgba))
	x.GodotColorNewRgb = NewRef(unsafe.Pointer(x.ref57717e51.godot_color_new_rgb))
	x.GodotColorGetR = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_color_get_r))
	x.GodotColorSetR = NewRef(unsafe.Pointer(x.ref57717e51.godot_color_set_r))
	x.GodotColorGetG = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_color_get_g))
	x.GodotColorSetG = NewRef(unsafe.Pointer(x.ref57717e51.godot_color_set_g))
	x.GodotColorGetB = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_color_get_b))
	x.GodotColorSetB = NewRef(unsafe.Pointer(x.ref57717e51.godot_color_set_b))
	x.GodotColorGetA = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_color_get_a))
	x.GodotColorSetA = NewRef(unsafe.Pointer(x.ref57717e51.godot_color_set_a))
	x.GodotColorGetH = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_color_get_h))
	x.GodotColorGetS = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_color_get_s))
	x.GodotColorGetV = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_color_get_v))
	x.GodotColorAsString = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_color_as_string))
	x.GodotColorToRgba32 = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_color_to_rgba32))
	x.GodotColorToArgb32 = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_color_to_argb32))
	x.GodotColorGray = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_color_gray))
	x.GodotColorInverted = NewGodotColorRef(unsafe.Pointer(x.ref57717e51.godot_color_inverted))
	x.GodotColorContrasted = NewGodotColorRef(unsafe.Pointer(x.ref57717e51.godot_color_contrasted))
	x.GodotColorLinearInterpolate = NewGodotColorRef(unsafe.Pointer(x.ref57717e51.godot_color_linear_interpolate))
	x.GodotColorBlend = NewGodotColorRef(unsafe.Pointer(x.ref57717e51.godot_color_blend))
	x.GodotColorToHtml = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_color_to_html))
	x.GodotColorOperatorEqual = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_color_operator_equal))
	x.GodotColorOperatorLess = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_color_operator_less))
	x.GodotVector2New = NewRef(unsafe.Pointer(x.ref57717e51.godot_vector2_new))
	x.GodotVector2AsString = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_vector2_as_string))
	x.GodotVector2Normalized = NewGodotVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_normalized))
	x.GodotVector2Length = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_vector2_length))
	x.GodotVector2Angle = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_vector2_angle))
	x.GodotVector2LengthSquared = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_vector2_length_squared))
	x.GodotVector2IsNormalized = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_vector2_is_normalized))
	x.GodotVector2DistanceTo = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_vector2_distance_to))
	x.GodotVector2DistanceSquaredTo = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_vector2_distance_squared_to))
	x.GodotVector2AngleTo = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_vector2_angle_to))
	x.GodotVector2AngleToPoint = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_vector2_angle_to_point))
	x.GodotVector2LinearInterpolate = NewGodotVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_linear_interpolate))
	x.GodotVector2CubicInterpolate = NewGodotVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_cubic_interpolate))
	x.GodotVector2Rotated = NewGodotVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_rotated))
	x.GodotVector2Tangent = NewGodotVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_tangent))
	x.GodotVector2Floor = NewGodotVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_floor))
	x.GodotVector2Snapped = NewGodotVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_snapped))
	x.GodotVector2Aspect = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_vector2_aspect))
	x.GodotVector2Dot = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_vector2_dot))
	x.GodotVector2Slide = NewGodotVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_slide))
	x.GodotVector2Bounce = NewGodotVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_bounce))
	x.GodotVector2Reflect = NewGodotVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_reflect))
	x.GodotVector2Abs = NewGodotVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_abs))
	x.GodotVector2Clamped = NewGodotVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_clamped))
	x.GodotVector2OperatorAdd = NewGodotVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_operator_add))
	x.GodotVector2OperatorSubstract = NewGodotVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_operator_substract))
	x.GodotVector2OperatorMultiplyVector = NewGodotVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_operator_multiply_vector))
	x.GodotVector2OperatorMultiplyScalar = NewGodotVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_operator_multiply_scalar))
	x.GodotVector2OperatorDivideVector = NewGodotVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_operator_divide_vector))
	x.GodotVector2OperatorDivideScalar = NewGodotVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_operator_divide_scalar))
	x.GodotVector2OperatorEqual = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_vector2_operator_equal))
	x.GodotVector2OperatorLess = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_vector2_operator_less))
	x.GodotVector2OperatorNeg = NewGodotVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_operator_neg))
	x.GodotVector2SetX = NewRef(unsafe.Pointer(x.ref57717e51.godot_vector2_set_x))
	x.GodotVector2SetY = NewRef(unsafe.Pointer(x.ref57717e51.godot_vector2_set_y))
	x.GodotVector2GetX = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_vector2_get_x))
	x.GodotVector2GetY = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_vector2_get_y))
	x.GodotQuatNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_quat_new))
	x.GodotQuatNewWithAxisAngle = NewRef(unsafe.Pointer(x.ref57717e51.godot_quat_new_with_axis_angle))
	x.GodotQuatGetX = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_quat_get_x))
	x.GodotQuatSetX = NewRef(unsafe.Pointer(x.ref57717e51.godot_quat_set_x))
	x.GodotQuatGetY = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_quat_get_y))
	x.GodotQuatSetY = NewRef(unsafe.Pointer(x.ref57717e51.godot_quat_set_y))
	x.GodotQuatGetZ = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_quat_get_z))
	x.GodotQuatSetZ = NewRef(unsafe.Pointer(x.ref57717e51.godot_quat_set_z))
	x.GodotQuatGetW = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_quat_get_w))
	x.GodotQuatSetW = NewRef(unsafe.Pointer(x.ref57717e51.godot_quat_set_w))
	x.GodotQuatAsString = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_quat_as_string))
	x.GodotQuatLength = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_quat_length))
	x.GodotQuatLengthSquared = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_quat_length_squared))
	x.GodotQuatNormalized = NewGodotQuatRef(unsafe.Pointer(x.ref57717e51.godot_quat_normalized))
	x.GodotQuatIsNormalized = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_quat_is_normalized))
	x.GodotQuatInverse = NewGodotQuatRef(unsafe.Pointer(x.ref57717e51.godot_quat_inverse))
	x.GodotQuatDot = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_quat_dot))
	x.GodotQuatXform = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_quat_xform))
	x.GodotQuatSlerp = NewGodotQuatRef(unsafe.Pointer(x.ref57717e51.godot_quat_slerp))
	x.GodotQuatSlerpni = NewGodotQuatRef(unsafe.Pointer(x.ref57717e51.godot_quat_slerpni))
	x.GodotQuatCubicSlerp = NewGodotQuatRef(unsafe.Pointer(x.ref57717e51.godot_quat_cubic_slerp))
	x.GodotQuatOperatorMultiply = NewGodotQuatRef(unsafe.Pointer(x.ref57717e51.godot_quat_operator_multiply))
	x.GodotQuatOperatorAdd = NewGodotQuatRef(unsafe.Pointer(x.ref57717e51.godot_quat_operator_add))
	x.GodotQuatOperatorSubstract = NewGodotQuatRef(unsafe.Pointer(x.ref57717e51.godot_quat_operator_substract))
	x.GodotQuatOperatorDivide = NewGodotQuatRef(unsafe.Pointer(x.ref57717e51.godot_quat_operator_divide))
	x.GodotQuatOperatorEqual = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_quat_operator_equal))
	x.GodotQuatOperatorNeg = NewGodotQuatRef(unsafe.Pointer(x.ref57717e51.godot_quat_operator_neg))
	x.GodotBasisNewWithRows = NewRef(unsafe.Pointer(x.ref57717e51.godot_basis_new_with_rows))
	x.GodotBasisNewWithAxisAndAngle = NewRef(unsafe.Pointer(x.ref57717e51.godot_basis_new_with_axis_and_angle))
	x.GodotBasisNewWithEuler = NewRef(unsafe.Pointer(x.ref57717e51.godot_basis_new_with_euler))
	x.GodotBasisAsString = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_basis_as_string))
	x.GodotBasisInverse = NewGodotBasisRef(unsafe.Pointer(x.ref57717e51.godot_basis_inverse))
	x.GodotBasisTransposed = NewGodotBasisRef(unsafe.Pointer(x.ref57717e51.godot_basis_transposed))
	x.GodotBasisOrthonormalized = NewGodotBasisRef(unsafe.Pointer(x.ref57717e51.godot_basis_orthonormalized))
	x.GodotBasisDeterminant = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_basis_determinant))
	x.GodotBasisRotated = NewGodotBasisRef(unsafe.Pointer(x.ref57717e51.godot_basis_rotated))
	x.GodotBasisScaled = NewGodotBasisRef(unsafe.Pointer(x.ref57717e51.godot_basis_scaled))
	x.GodotBasisGetScale = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_basis_get_scale))
	x.GodotBasisGetEuler = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_basis_get_euler))
	x.GodotBasisTdotx = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_basis_tdotx))
	x.GodotBasisTdoty = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_basis_tdoty))
	x.GodotBasisTdotz = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_basis_tdotz))
	x.GodotBasisXform = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_basis_xform))
	x.GodotBasisXformInv = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_basis_xform_inv))
	x.GodotBasisGetOrthogonalIndex = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_basis_get_orthogonal_index))
	x.GodotBasisNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_basis_new))
	x.GodotBasisNewWithEulerQuat = NewRef(unsafe.Pointer(x.ref57717e51.godot_basis_new_with_euler_quat))
	x.GodotBasisGetElements = NewRef(unsafe.Pointer(x.ref57717e51.godot_basis_get_elements))
	x.GodotBasisGetAxis = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_basis_get_axis))
	x.GodotBasisSetAxis = NewRef(unsafe.Pointer(x.ref57717e51.godot_basis_set_axis))
	x.GodotBasisGetRow = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_basis_get_row))
	x.GodotBasisSetRow = NewRef(unsafe.Pointer(x.ref57717e51.godot_basis_set_row))
	x.GodotBasisOperatorEqual = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_basis_operator_equal))
	x.GodotBasisOperatorAdd = NewGodotBasisRef(unsafe.Pointer(x.ref57717e51.godot_basis_operator_add))
	x.GodotBasisOperatorSubstract = NewGodotBasisRef(unsafe.Pointer(x.ref57717e51.godot_basis_operator_substract))
	x.GodotBasisOperatorMultiplyVector = NewGodotBasisRef(unsafe.Pointer(x.ref57717e51.godot_basis_operator_multiply_vector))
	x.GodotBasisOperatorMultiplyScalar = NewGodotBasisRef(unsafe.Pointer(x.ref57717e51.godot_basis_operator_multiply_scalar))
	x.GodotVector3New = NewRef(unsafe.Pointer(x.ref57717e51.godot_vector3_new))
	x.GodotVector3AsString = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_vector3_as_string))
	x.GodotVector3MinAxis = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_vector3_min_axis))
	x.GodotVector3MaxAxis = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_vector3_max_axis))
	x.GodotVector3Length = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_vector3_length))
	x.GodotVector3LengthSquared = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_vector3_length_squared))
	x.GodotVector3IsNormalized = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_vector3_is_normalized))
	x.GodotVector3Normalized = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_normalized))
	x.GodotVector3Inverse = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_inverse))
	x.GodotVector3Snapped = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_snapped))
	x.GodotVector3Rotated = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_rotated))
	x.GodotVector3LinearInterpolate = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_linear_interpolate))
	x.GodotVector3CubicInterpolate = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_cubic_interpolate))
	x.GodotVector3Dot = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_vector3_dot))
	x.GodotVector3Cross = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_cross))
	x.GodotVector3Outer = NewGodotBasisRef(unsafe.Pointer(x.ref57717e51.godot_vector3_outer))
	x.GodotVector3ToDiagonalMatrix = NewGodotBasisRef(unsafe.Pointer(x.ref57717e51.godot_vector3_to_diagonal_matrix))
	x.GodotVector3Abs = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_abs))
	x.GodotVector3Floor = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_floor))
	x.GodotVector3Ceil = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_ceil))
	x.GodotVector3DistanceTo = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_vector3_distance_to))
	x.GodotVector3DistanceSquaredTo = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_vector3_distance_squared_to))
	x.GodotVector3AngleTo = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_vector3_angle_to))
	x.GodotVector3Slide = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_slide))
	x.GodotVector3Bounce = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_bounce))
	x.GodotVector3Reflect = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_reflect))
	x.GodotVector3OperatorAdd = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_operator_add))
	x.GodotVector3OperatorSubstract = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_operator_substract))
	x.GodotVector3OperatorMultiplyVector = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_operator_multiply_vector))
	x.GodotVector3OperatorMultiplyScalar = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_operator_multiply_scalar))
	x.GodotVector3OperatorDivideVector = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_operator_divide_vector))
	x.GodotVector3OperatorDivideScalar = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_operator_divide_scalar))
	x.GodotVector3OperatorEqual = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_vector3_operator_equal))
	x.GodotVector3OperatorLess = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_vector3_operator_less))
	x.GodotVector3OperatorNeg = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_operator_neg))
	x.GodotVector3SetAxis = NewRef(unsafe.Pointer(x.ref57717e51.godot_vector3_set_axis))
	x.GodotVector3GetAxis = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_vector3_get_axis))
	x.GodotPoolByteArrayNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_new))
	x.GodotPoolByteArrayNewCopy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_new_copy))
	x.GodotPoolByteArrayNewWithArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_new_with_array))
	x.GodotPoolByteArrayAppend = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_append))
	x.GodotPoolByteArrayAppendArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_append_array))
	x.GodotPoolByteArrayInsert = NewGodotErrorRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_insert))
	x.GodotPoolByteArrayInvert = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_invert))
	x.GodotPoolByteArrayPushBack = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_push_back))
	x.GodotPoolByteArrayRemove = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_remove))
	x.GodotPoolByteArrayResize = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_resize))
	x.GodotPoolByteArrayRead = NewGodotPoolByteArrayReadAccessRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_read))
	x.GodotPoolByteArrayWrite = NewGodotPoolByteArrayWriteAccessRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_write))
	x.GodotPoolByteArraySet = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_set))
	x.GodotPoolByteArrayGet = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_get))
	x.GodotPoolByteArraySize = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_size))
	x.GodotPoolByteArrayDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_destroy))
	x.GodotPoolIntArrayNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_new))
	x.GodotPoolIntArrayNewCopy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_new_copy))
	x.GodotPoolIntArrayNewWithArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_new_with_array))
	x.GodotPoolIntArrayAppend = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_append))
	x.GodotPoolIntArrayAppendArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_append_array))
	x.GodotPoolIntArrayInsert = NewGodotErrorRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_insert))
	x.GodotPoolIntArrayInvert = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_invert))
	x.GodotPoolIntArrayPushBack = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_push_back))
	x.GodotPoolIntArrayRemove = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_remove))
	x.GodotPoolIntArrayResize = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_resize))
	x.GodotPoolIntArrayRead = NewGodotPoolIntArrayReadAccessRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_read))
	x.GodotPoolIntArrayWrite = NewGodotPoolIntArrayWriteAccessRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_write))
	x.GodotPoolIntArraySet = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_set))
	x.GodotPoolIntArrayGet = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_get))
	x.GodotPoolIntArraySize = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_size))
	x.GodotPoolIntArrayDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_destroy))
	x.GodotPoolRealArrayNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_new))
	x.GodotPoolRealArrayNewCopy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_new_copy))
	x.GodotPoolRealArrayNewWithArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_new_with_array))
	x.GodotPoolRealArrayAppend = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_append))
	x.GodotPoolRealArrayAppendArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_append_array))
	x.GodotPoolRealArrayInsert = NewGodotErrorRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_insert))
	x.GodotPoolRealArrayInvert = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_invert))
	x.GodotPoolRealArrayPushBack = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_push_back))
	x.GodotPoolRealArrayRemove = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_remove))
	x.GodotPoolRealArrayResize = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_resize))
	x.GodotPoolRealArrayRead = NewGodotPoolRealArrayReadAccessRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_read))
	x.GodotPoolRealArrayWrite = NewGodotPoolRealArrayWriteAccessRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_write))
	x.GodotPoolRealArraySet = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_set))
	x.GodotPoolRealArrayGet = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_get))
	x.GodotPoolRealArraySize = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_size))
	x.GodotPoolRealArrayDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_destroy))
	x.GodotPoolStringArrayNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_new))
	x.GodotPoolStringArrayNewCopy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_new_copy))
	x.GodotPoolStringArrayNewWithArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_new_with_array))
	x.GodotPoolStringArrayAppend = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_append))
	x.GodotPoolStringArrayAppendArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_append_array))
	x.GodotPoolStringArrayInsert = NewGodotErrorRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_insert))
	x.GodotPoolStringArrayInvert = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_invert))
	x.GodotPoolStringArrayPushBack = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_push_back))
	x.GodotPoolStringArrayRemove = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_remove))
	x.GodotPoolStringArrayResize = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_resize))
	x.GodotPoolStringArrayRead = NewGodotPoolStringArrayReadAccessRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_read))
	x.GodotPoolStringArrayWrite = NewGodotPoolStringArrayWriteAccessRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_write))
	x.GodotPoolStringArraySet = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_set))
	x.GodotPoolStringArrayGet = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_get))
	x.GodotPoolStringArraySize = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_size))
	x.GodotPoolStringArrayDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_destroy))
	x.GodotPoolVector2ArrayNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_new))
	x.GodotPoolVector2ArrayNewCopy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_new_copy))
	x.GodotPoolVector2ArrayNewWithArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_new_with_array))
	x.GodotPoolVector2ArrayAppend = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_append))
	x.GodotPoolVector2ArrayAppendArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_append_array))
	x.GodotPoolVector2ArrayInsert = NewGodotErrorRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_insert))
	x.GodotPoolVector2ArrayInvert = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_invert))
	x.GodotPoolVector2ArrayPushBack = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_push_back))
	x.GodotPoolVector2ArrayRemove = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_remove))
	x.GodotPoolVector2ArrayResize = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_resize))
	x.GodotPoolVector2ArrayRead = NewGodotPoolVector2ArrayReadAccessRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_read))
	x.GodotPoolVector2ArrayWrite = NewGodotPoolVector2ArrayWriteAccessRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_write))
	x.GodotPoolVector2ArraySet = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_set))
	x.GodotPoolVector2ArrayGet = NewGodotVector2Ref(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_get))
	x.GodotPoolVector2ArraySize = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_size))
	x.GodotPoolVector2ArrayDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_destroy))
	x.GodotPoolVector3ArrayNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_new))
	x.GodotPoolVector3ArrayNewCopy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_new_copy))
	x.GodotPoolVector3ArrayNewWithArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_new_with_array))
	x.GodotPoolVector3ArrayAppend = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_append))
	x.GodotPoolVector3ArrayAppendArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_append_array))
	x.GodotPoolVector3ArrayInsert = NewGodotErrorRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_insert))
	x.GodotPoolVector3ArrayInvert = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_invert))
	x.GodotPoolVector3ArrayPushBack = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_push_back))
	x.GodotPoolVector3ArrayRemove = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_remove))
	x.GodotPoolVector3ArrayResize = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_resize))
	x.GodotPoolVector3ArrayRead = NewGodotPoolVector3ArrayReadAccessRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_read))
	x.GodotPoolVector3ArrayWrite = NewGodotPoolVector3ArrayWriteAccessRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_write))
	x.GodotPoolVector3ArraySet = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_set))
	x.GodotPoolVector3ArrayGet = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_get))
	x.GodotPoolVector3ArraySize = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_size))
	x.GodotPoolVector3ArrayDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_destroy))
	x.GodotPoolColorArrayNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_new))
	x.GodotPoolColorArrayNewCopy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_new_copy))
	x.GodotPoolColorArrayNewWithArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_new_with_array))
	x.GodotPoolColorArrayAppend = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_append))
	x.GodotPoolColorArrayAppendArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_append_array))
	x.GodotPoolColorArrayInsert = NewGodotErrorRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_insert))
	x.GodotPoolColorArrayInvert = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_invert))
	x.GodotPoolColorArrayPushBack = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_push_back))
	x.GodotPoolColorArrayRemove = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_remove))
	x.GodotPoolColorArrayResize = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_resize))
	x.GodotPoolColorArrayRead = NewGodotPoolColorArrayReadAccessRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_read))
	x.GodotPoolColorArrayWrite = NewGodotPoolColorArrayWriteAccessRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_write))
	x.GodotPoolColorArraySet = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_set))
	x.GodotPoolColorArrayGet = NewGodotColorRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_get))
	x.GodotPoolColorArraySize = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_size))
	x.GodotPoolColorArrayDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_destroy))
	x.GodotPoolByteArrayReadAccessPtr = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_read_access_ptr))
	x.GodotPoolByteArrayReadAccessOperatorAssign = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_read_access_operator_assign))
	x.GodotPoolByteArrayReadAccessDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_read_access_destroy))
	x.GodotPoolIntArrayReadAccessPtr = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_read_access_ptr))
	x.GodotPoolIntArrayReadAccessOperatorAssign = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_read_access_operator_assign))
	x.GodotPoolIntArrayReadAccessDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_read_access_destroy))
	x.GodotPoolRealArrayReadAccessPtr = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_read_access_ptr))
	x.GodotPoolRealArrayReadAccessOperatorAssign = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_read_access_operator_assign))
	x.GodotPoolRealArrayReadAccessDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_read_access_destroy))
	x.GodotPoolStringArrayReadAccessPtr = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_read_access_ptr))
	x.GodotPoolStringArrayReadAccessOperatorAssign = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_read_access_operator_assign))
	x.GodotPoolStringArrayReadAccessDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_read_access_destroy))
	x.GodotPoolVector2ArrayReadAccessPtr = NewGodotVector2Ref(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_read_access_ptr))
	x.GodotPoolVector2ArrayReadAccessOperatorAssign = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_read_access_operator_assign))
	x.GodotPoolVector2ArrayReadAccessDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_read_access_destroy))
	x.GodotPoolVector3ArrayReadAccessPtr = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_read_access_ptr))
	x.GodotPoolVector3ArrayReadAccessOperatorAssign = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_read_access_operator_assign))
	x.GodotPoolVector3ArrayReadAccessDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_read_access_destroy))
	x.GodotPoolColorArrayReadAccessPtr = NewGodotColorRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_read_access_ptr))
	x.GodotPoolColorArrayReadAccessOperatorAssign = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_read_access_operator_assign))
	x.GodotPoolColorArrayReadAccessDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_read_access_destroy))
	x.GodotPoolByteArrayWriteAccessPtr = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_write_access_ptr))
	x.GodotPoolByteArrayWriteAccessOperatorAssign = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_write_access_operator_assign))
	x.GodotPoolByteArrayWriteAccessDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_write_access_destroy))
	x.GodotPoolIntArrayWriteAccessPtr = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_write_access_ptr))
	x.GodotPoolIntArrayWriteAccessOperatorAssign = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_write_access_operator_assign))
	x.GodotPoolIntArrayWriteAccessDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_write_access_destroy))
	x.GodotPoolRealArrayWriteAccessPtr = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_write_access_ptr))
	x.GodotPoolRealArrayWriteAccessOperatorAssign = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_write_access_operator_assign))
	x.GodotPoolRealArrayWriteAccessDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_write_access_destroy))
	x.GodotPoolStringArrayWriteAccessPtr = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_write_access_ptr))
	x.GodotPoolStringArrayWriteAccessOperatorAssign = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_write_access_operator_assign))
	x.GodotPoolStringArrayWriteAccessDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_write_access_destroy))
	x.GodotPoolVector2ArrayWriteAccessPtr = NewGodotVector2Ref(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_write_access_ptr))
	x.GodotPoolVector2ArrayWriteAccessOperatorAssign = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_write_access_operator_assign))
	x.GodotPoolVector2ArrayWriteAccessDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_write_access_destroy))
	x.GodotPoolVector3ArrayWriteAccessPtr = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_write_access_ptr))
	x.GodotPoolVector3ArrayWriteAccessOperatorAssign = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_write_access_operator_assign))
	x.GodotPoolVector3ArrayWriteAccessDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_write_access_destroy))
	x.GodotPoolColorArrayWriteAccessPtr = NewGodotColorRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_write_access_ptr))
	x.GodotPoolColorArrayWriteAccessOperatorAssign = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_write_access_operator_assign))
	x.GodotPoolColorArrayWriteAccessDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_write_access_destroy))
	x.GodotArrayNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_new))
	x.GodotArrayNewCopy = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_new_copy))
	x.GodotArrayNewPoolColorArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_new_pool_color_array))
	x.GodotArrayNewPoolVector3Array = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_new_pool_vector3_array))
	x.GodotArrayNewPoolVector2Array = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_new_pool_vector2_array))
	x.GodotArrayNewPoolStringArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_new_pool_string_array))
	x.GodotArrayNewPoolRealArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_new_pool_real_array))
	x.GodotArrayNewPoolIntArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_new_pool_int_array))
	x.GodotArrayNewPoolByteArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_new_pool_byte_array))
	x.GodotArraySet = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_set))
	x.GodotArrayGet = NewGodotVariantRef(unsafe.Pointer(x.ref57717e51.godot_array_get))
	x.GodotArrayOperatorIndex = NewGodotVariantRef(unsafe.Pointer(x.ref57717e51.godot_array_operator_index))
	x.GodotArrayOperatorIndexConst = NewGodotVariantRef(unsafe.Pointer(x.ref57717e51.godot_array_operator_index_const))
	x.GodotArrayAppend = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_append))
	x.GodotArrayClear = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_clear))
	x.GodotArrayCount = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_array_count))
	x.GodotArrayEmpty = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_array_empty))
	x.GodotArrayErase = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_erase))
	x.GodotArrayFront = NewGodotVariantRef(unsafe.Pointer(x.ref57717e51.godot_array_front))
	x.GodotArrayBack = NewGodotVariantRef(unsafe.Pointer(x.ref57717e51.godot_array_back))
	x.GodotArrayFind = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_array_find))
	x.GodotArrayFindLast = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_array_find_last))
	x.GodotArrayHas = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_array_has))
	x.GodotArrayHash = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_array_hash))
	x.GodotArrayInsert = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_insert))
	x.GodotArrayInvert = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_invert))
	x.GodotArrayPopBack = NewGodotVariantRef(unsafe.Pointer(x.ref57717e51.godot_array_pop_back))
	x.GodotArrayPopFront = NewGodotVariantRef(unsafe.Pointer(x.ref57717e51.godot_array_pop_front))
	x.GodotArrayPushBack = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_push_back))
	x.GodotArrayPushFront = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_push_front))
	x.GodotArrayRemove = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_remove))
	x.GodotArrayResize = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_resize))
	x.GodotArrayRfind = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_array_rfind))
	x.GodotArraySize = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_array_size))
	x.GodotArraySort = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_sort))
	x.GodotArraySortCustom = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_sort_custom))
	x.GodotArrayBsearch = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_array_bsearch))
	x.GodotArrayBsearchCustom = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_array_bsearch_custom))
	x.GodotArrayDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_destroy))
	x.GodotDictionaryNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_new))
	x.GodotDictionaryNewCopy = NewRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_new_copy))
	x.GodotDictionaryDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_destroy))
	x.GodotDictionarySize = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_size))
	x.GodotDictionaryEmpty = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_empty))
	x.GodotDictionaryClear = NewRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_clear))
	x.GodotDictionaryHas = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_has))
	x.GodotDictionaryHasAll = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_has_all))
	x.GodotDictionaryErase = NewRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_erase))
	x.GodotDictionaryHash = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_hash))
	x.GodotDictionaryKeys = NewGodotArrayRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_keys))
	x.GodotDictionaryValues = NewGodotArrayRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_values))
	x.GodotDictionaryGet = NewGodotVariantRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_get))
	x.GodotDictionarySet = NewRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_set))
	x.GodotDictionaryOperatorIndex = NewGodotVariantRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_operator_index))
	x.GodotDictionaryOperatorIndexConst = NewGodotVariantRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_operator_index_const))
	x.GodotDictionaryNext = NewGodotVariantRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_next))
	x.GodotDictionaryOperatorEqual = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_operator_equal))
	x.GodotDictionaryToJson = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_to_json))
	x.GodotNodePathNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_node_path_new))
	x.GodotNodePathNewCopy = NewRef(unsafe.Pointer(x.ref57717e51.godot_node_path_new_copy))
	x.GodotNodePathDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_node_path_destroy))
	x.GodotNodePathAsString = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_node_path_as_string))
	x.GodotNodePathIsAbsolute = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_node_path_is_absolute))
	x.GodotNodePathGetNameCount = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_node_path_get_name_count))
	x.GodotNodePathGetName = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_node_path_get_name))
	x.GodotNodePathGetSubnameCount = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_node_path_get_subname_count))
	x.GodotNodePathGetSubname = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_node_path_get_subname))
	x.GodotNodePathGetConcatenatedSubnames = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_node_path_get_concatenated_subnames))
	x.GodotNodePathIsEmpty = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_node_path_is_empty))
	x.GodotNodePathOperatorEqual = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_node_path_operator_equal))
	x.GodotPlaneNewWithReals = NewRef(unsafe.Pointer(x.ref57717e51.godot_plane_new_with_reals))
	x.GodotPlaneNewWithVectors = NewRef(unsafe.Pointer(x.ref57717e51.godot_plane_new_with_vectors))
	x.GodotPlaneNewWithNormal = NewRef(unsafe.Pointer(x.ref57717e51.godot_plane_new_with_normal))
	x.GodotPlaneAsString = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_plane_as_string))
	x.GodotPlaneNormalized = NewGodotPlaneRef(unsafe.Pointer(x.ref57717e51.godot_plane_normalized))
	x.GodotPlaneCenter = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_plane_center))
	x.GodotPlaneGetAnyPoint = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_plane_get_any_point))
	x.GodotPlaneIsPointOver = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_plane_is_point_over))
	x.GodotPlaneDistanceTo = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_plane_distance_to))
	x.GodotPlaneHasPoint = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_plane_has_point))
	x.GodotPlaneProject = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_plane_project))
	x.GodotPlaneIntersect3 = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_plane_intersect_3))
	x.GodotPlaneIntersectsRay = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_plane_intersects_ray))
	x.GodotPlaneIntersectsSegment = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_plane_intersects_segment))
	x.GodotPlaneOperatorNeg = NewGodotPlaneRef(unsafe.Pointer(x.ref57717e51.godot_plane_operator_neg))
	x.GodotPlaneOperatorEqual = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_plane_operator_equal))
	x.GodotPlaneSetNormal = NewRef(unsafe.Pointer(x.ref57717e51.godot_plane_set_normal))
	x.GodotPlaneGetNormal = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_plane_get_normal))
	x.GodotPlaneGetD = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_plane_get_d))
	x.GodotPlaneSetD = NewRef(unsafe.Pointer(x.ref57717e51.godot_plane_set_d))
	x.GodotRect2NewWithPositionAndSize = NewRef(unsafe.Pointer(x.ref57717e51.godot_rect2_new_with_position_and_size))
	x.GodotRect2New = NewRef(unsafe.Pointer(x.ref57717e51.godot_rect2_new))
	x.GodotRect2AsString = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_rect2_as_string))
	x.GodotRect2GetArea = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_rect2_get_area))
	x.GodotRect2Intersects = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_rect2_intersects))
	x.GodotRect2Encloses = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_rect2_encloses))
	x.GodotRect2HasNoArea = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_rect2_has_no_area))
	x.GodotRect2Clip = NewGodotRect2Ref(unsafe.Pointer(x.ref57717e51.godot_rect2_clip))
	x.GodotRect2Merge = NewGodotRect2Ref(unsafe.Pointer(x.ref57717e51.godot_rect2_merge))
	x.GodotRect2HasPoint = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_rect2_has_point))
	x.GodotRect2Grow = NewGodotRect2Ref(unsafe.Pointer(x.ref57717e51.godot_rect2_grow))
	x.GodotRect2Expand = NewGodotRect2Ref(unsafe.Pointer(x.ref57717e51.godot_rect2_expand))
	x.GodotRect2OperatorEqual = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_rect2_operator_equal))
	x.GodotRect2GetPosition = NewGodotVector2Ref(unsafe.Pointer(x.ref57717e51.godot_rect2_get_position))
	x.GodotRect2GetSize = NewGodotVector2Ref(unsafe.Pointer(x.ref57717e51.godot_rect2_get_size))
	x.GodotRect2SetPosition = NewRef(unsafe.Pointer(x.ref57717e51.godot_rect2_set_position))
	x.GodotRect2SetSize = NewRef(unsafe.Pointer(x.ref57717e51.godot_rect2_set_size))
	x.GodotAabbNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_aabb_new))
	x.GodotAabbGetPosition = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_aabb_get_position))
	x.GodotAabbSetPosition = NewRef(unsafe.Pointer(x.ref57717e51.godot_aabb_set_position))
	x.GodotAabbGetSize = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_aabb_get_size))
	x.GodotAabbSetSize = NewRef(unsafe.Pointer(x.ref57717e51.godot_aabb_set_size))
	x.GodotAabbAsString = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_aabb_as_string))
	x.GodotAabbGetArea = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_aabb_get_area))
	x.GodotAabbHasNoArea = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_aabb_has_no_area))
	x.GodotAabbHasNoSurface = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_aabb_has_no_surface))
	x.GodotAabbIntersects = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_aabb_intersects))
	x.GodotAabbEncloses = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_aabb_encloses))
	x.GodotAabbMerge = NewGodotAabbRef(unsafe.Pointer(x.ref57717e51.godot_aabb_merge))
	x.GodotAabbIntersection = NewGodotAabbRef(unsafe.Pointer(x.ref57717e51.godot_aabb_intersection))
	x.GodotAabbIntersectsPlane = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_aabb_intersects_plane))
	x.GodotAabbIntersectsSegment = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_aabb_intersects_segment))
	x.GodotAabbHasPoint = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_aabb_has_point))
	x.GodotAabbGetSupport = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_aabb_get_support))
	x.GodotAabbGetLongestAxis = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_aabb_get_longest_axis))
	x.GodotAabbGetLongestAxisIndex = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_aabb_get_longest_axis_index))
	x.GodotAabbGetLongestAxisSize = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_aabb_get_longest_axis_size))
	x.GodotAabbGetShortestAxis = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_aabb_get_shortest_axis))
	x.GodotAabbGetShortestAxisIndex = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_aabb_get_shortest_axis_index))
	x.GodotAabbGetShortestAxisSize = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_aabb_get_shortest_axis_size))
	x.GodotAabbExpand = NewGodotAabbRef(unsafe.Pointer(x.ref57717e51.godot_aabb_expand))
	x.GodotAabbGrow = NewGodotAabbRef(unsafe.Pointer(x.ref57717e51.godot_aabb_grow))
	x.GodotAabbGetEndpoint = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_aabb_get_endpoint))
	x.GodotAabbOperatorEqual = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_aabb_operator_equal))
	x.GodotRidNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_rid_new))
	x.GodotRidGetId = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_rid_get_id))
	x.GodotRidNewWithResource = NewRef(unsafe.Pointer(x.ref57717e51.godot_rid_new_with_resource))
	x.GodotRidOperatorEqual = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_rid_operator_equal))
	x.GodotRidOperatorLess = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_rid_operator_less))
	x.GodotTransformNewWithAxisOrigin = NewRef(unsafe.Pointer(x.ref57717e51.godot_transform_new_with_axis_origin))
	x.GodotTransformNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_transform_new))
	x.GodotTransformGetBasis = NewGodotBasisRef(unsafe.Pointer(x.ref57717e51.godot_transform_get_basis))
	x.GodotTransformSetBasis = NewRef(unsafe.Pointer(x.ref57717e51.godot_transform_set_basis))
	x.GodotTransformGetOrigin = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_transform_get_origin))
	x.GodotTransformSetOrigin = NewRef(unsafe.Pointer(x.ref57717e51.godot_transform_set_origin))
	x.GodotTransformAsString = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_transform_as_string))
	x.GodotTransformInverse = NewGodotTransformRef(unsafe.Pointer(x.ref57717e51.godot_transform_inverse))
	x.GodotTransformAffineInverse = NewGodotTransformRef(unsafe.Pointer(x.ref57717e51.godot_transform_affine_inverse))
	x.GodotTransformOrthonormalized = NewGodotTransformRef(unsafe.Pointer(x.ref57717e51.godot_transform_orthonormalized))
	x.GodotTransformRotated = NewGodotTransformRef(unsafe.Pointer(x.ref57717e51.godot_transform_rotated))
	x.GodotTransformScaled = NewGodotTransformRef(unsafe.Pointer(x.ref57717e51.godot_transform_scaled))
	x.GodotTransformTranslated = NewGodotTransformRef(unsafe.Pointer(x.ref57717e51.godot_transform_translated))
	x.GodotTransformLookingAt = NewGodotTransformRef(unsafe.Pointer(x.ref57717e51.godot_transform_looking_at))
	x.GodotTransformXformPlane = NewGodotPlaneRef(unsafe.Pointer(x.ref57717e51.godot_transform_xform_plane))
	x.GodotTransformXformInvPlane = NewGodotPlaneRef(unsafe.Pointer(x.ref57717e51.godot_transform_xform_inv_plane))
	x.GodotTransformNewIdentity = NewRef(unsafe.Pointer(x.ref57717e51.godot_transform_new_identity))
	x.GodotTransformOperatorEqual = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_transform_operator_equal))
	x.GodotTransformOperatorMultiply = NewGodotTransformRef(unsafe.Pointer(x.ref57717e51.godot_transform_operator_multiply))
	x.GodotTransformXformVector3 = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_transform_xform_vector3))
	x.GodotTransformXformInvVector3 = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_transform_xform_inv_vector3))
	x.GodotTransformXformAabb = NewGodotAabbRef(unsafe.Pointer(x.ref57717e51.godot_transform_xform_aabb))
	x.GodotTransformXformInvAabb = NewGodotAabbRef(unsafe.Pointer(x.ref57717e51.godot_transform_xform_inv_aabb))
	x.GodotTransform2dNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_transform2d_new))
	x.GodotTransform2dNewAxisOrigin = NewRef(unsafe.Pointer(x.ref57717e51.godot_transform2d_new_axis_origin))
	x.GodotTransform2dAsString = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_transform2d_as_string))
	x.GodotTransform2dInverse = NewGodotTransform2dRef(unsafe.Pointer(x.ref57717e51.godot_transform2d_inverse))
	x.GodotTransform2dAffineInverse = NewGodotTransform2dRef(unsafe.Pointer(x.ref57717e51.godot_transform2d_affine_inverse))
	x.GodotTransform2dGetRotation = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_transform2d_get_rotation))
	x.GodotTransform2dGetOrigin = NewGodotVector2Ref(unsafe.Pointer(x.ref57717e51.godot_transform2d_get_origin))
	x.GodotTransform2dGetScale = NewGodotVector2Ref(unsafe.Pointer(x.ref57717e51.godot_transform2d_get_scale))
	x.GodotTransform2dOrthonormalized = NewGodotTransform2dRef(unsafe.Pointer(x.ref57717e51.godot_transform2d_orthonormalized))
	x.GodotTransform2dRotated = NewGodotTransform2dRef(unsafe.Pointer(x.ref57717e51.godot_transform2d_rotated))
	x.GodotTransform2dScaled = NewGodotTransform2dRef(unsafe.Pointer(x.ref57717e51.godot_transform2d_scaled))
	x.GodotTransform2dTranslated = NewGodotTransform2dRef(unsafe.Pointer(x.ref57717e51.godot_transform2d_translated))
	x.GodotTransform2dXformVector2 = NewGodotVector2Ref(unsafe.Pointer(x.ref57717e51.godot_transform2d_xform_vector2))
	x.GodotTransform2dXformInvVector2 = NewGodotVector2Ref(unsafe.Pointer(x.ref57717e51.godot_transform2d_xform_inv_vector2))
	x.GodotTransform2dBasisXformVector2 = NewGodotVector2Ref(unsafe.Pointer(x.ref57717e51.godot_transform2d_basis_xform_vector2))
	x.GodotTransform2dBasisXformInvVector2 = NewGodotVector2Ref(unsafe.Pointer(x.ref57717e51.godot_transform2d_basis_xform_inv_vector2))
	x.GodotTransform2dInterpolateWith = NewGodotTransform2dRef(unsafe.Pointer(x.ref57717e51.godot_transform2d_interpolate_with))
	x.GodotTransform2dOperatorEqual = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_transform2d_operator_equal))
	x.GodotTransform2dOperatorMultiply = NewGodotTransform2dRef(unsafe.Pointer(x.ref57717e51.godot_transform2d_operator_multiply))
	x.GodotTransform2dNewIdentity = NewRef(unsafe.Pointer(x.ref57717e51.godot_transform2d_new_identity))
	x.GodotTransform2dXformRect2 = NewGodotRect2Ref(unsafe.Pointer(x.ref57717e51.godot_transform2d_xform_rect2))
	x.GodotTransform2dXformInvRect2 = NewGodotRect2Ref(unsafe.Pointer(x.ref57717e51.godot_transform2d_xform_inv_rect2))
	x.GodotVariantGetType = NewGodotVariantTypeRef(unsafe.Pointer(x.ref57717e51.godot_variant_get_type))
	x.GodotVariantNewCopy = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_copy))
	x.GodotVariantNewNil = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_nil))
	x.GodotVariantNewBool = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_bool))
	x.GodotVariantNewUint = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_uint))
	x.GodotVariantNewInt = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_int))
	x.GodotVariantNewReal = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_real))
	x.GodotVariantNewString = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_string))
	x.GodotVariantNewVector2 = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_vector2))
	x.GodotVariantNewRect2 = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_rect2))
	x.GodotVariantNewVector3 = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_vector3))
	x.GodotVariantNewTransform2d = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_transform2d))
	x.GodotVariantNewPlane = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_plane))
	x.GodotVariantNewQuat = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_quat))
	x.GodotVariantNewAabb = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_aabb))
	x.GodotVariantNewBasis = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_basis))
	x.GodotVariantNewTransform = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_transform))
	x.GodotVariantNewColor = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_color))
	x.GodotVariantNewNodePath = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_node_path))
	x.GodotVariantNewRid = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_rid))
	x.GodotVariantNewObject = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_object))
	x.GodotVariantNewDictionary = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_dictionary))
	x.GodotVariantNewArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_array))
	x.GodotVariantNewPoolByteArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_pool_byte_array))
	x.GodotVariantNewPoolIntArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_pool_int_array))
	x.GodotVariantNewPoolRealArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_pool_real_array))
	x.GodotVariantNewPoolStringArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_pool_string_array))
	x.GodotVariantNewPoolVector2Array = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_pool_vector2_array))
	x.GodotVariantNewPoolVector3Array = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_pool_vector3_array))
	x.GodotVariantNewPoolColorArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_pool_color_array))
	x.GodotVariantAsBool = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_bool))
	x.GodotVariantAsUint = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_uint))
	x.GodotVariantAsInt = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_int))
	x.GodotVariantAsReal = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_real))
	x.GodotVariantAsString = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_string))
	x.GodotVariantAsVector2 = NewGodotVector2Ref(unsafe.Pointer(x.ref57717e51.godot_variant_as_vector2))
	x.GodotVariantAsRect2 = NewGodotRect2Ref(unsafe.Pointer(x.ref57717e51.godot_variant_as_rect2))
	x.GodotVariantAsVector3 = NewGodotVector3Ref(unsafe.Pointer(x.ref57717e51.godot_variant_as_vector3))
	x.GodotVariantAsTransform2d = NewGodotTransform2dRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_transform2d))
	x.GodotVariantAsPlane = NewGodotPlaneRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_plane))
	x.GodotVariantAsQuat = NewGodotQuatRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_quat))
	x.GodotVariantAsAabb = NewGodotAabbRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_aabb))
	x.GodotVariantAsBasis = NewGodotBasisRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_basis))
	x.GodotVariantAsTransform = NewGodotTransformRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_transform))
	x.GodotVariantAsColor = NewGodotColorRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_color))
	x.GodotVariantAsNodePath = NewGodotNodePathRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_node_path))
	x.GodotVariantAsRid = NewGodotRidRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_rid))
	x.GodotVariantAsObject = NewGodotObjectRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_object))
	x.GodotVariantAsDictionary = NewGodotDictionaryRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_dictionary))
	x.GodotVariantAsArray = NewGodotArrayRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_array))
	x.GodotVariantAsPoolByteArray = NewGodotPoolByteArrayRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_pool_byte_array))
	x.GodotVariantAsPoolIntArray = NewGodotPoolIntArrayRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_pool_int_array))
	x.GodotVariantAsPoolRealArray = NewGodotPoolRealArrayRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_pool_real_array))
	x.GodotVariantAsPoolStringArray = NewGodotPoolStringArrayRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_pool_string_array))
	x.GodotVariantAsPoolVector2Array = NewGodotPoolVector2ArrayRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_pool_vector2_array))
	x.GodotVariantAsPoolVector3Array = NewGodotPoolVector3ArrayRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_pool_vector3_array))
	x.GodotVariantAsPoolColorArray = NewGodotPoolColorArrayRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_pool_color_array))
	x.GodotVariantCall = NewGodotVariantRef(unsafe.Pointer(x.ref57717e51.godot_variant_call))
	x.GodotVariantHasMethod = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_variant_has_method))
	x.GodotVariantOperatorEqual = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_variant_operator_equal))
	x.GodotVariantOperatorLess = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_variant_operator_less))
	x.GodotVariantHashCompare = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_variant_hash_compare))
	x.GodotVariantBooleanize = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_variant_booleanize))
	x.GodotVariantDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_destroy))
	x.GodotStringNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_new))
	x.GodotStringNewCopy = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_new_copy))
	x.GodotStringNewData = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_new_data))
	x.GodotStringNewUnicodeData = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_new_unicode_data))
	x.GodotStringGetData = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_get_data))
	x.GodotStringOperatorIndex = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_operator_index))
	x.GodotStringOperatorIndexConst = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_operator_index_const))
	x.GodotStringUnicodeStr = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_unicode_str))
	x.GodotStringOperatorEqual = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_operator_equal))
	x.GodotStringOperatorLess = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_operator_less))
	x.GodotStringOperatorPlus = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_operator_plus))
	x.GodotStringLength = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_string_length))
	x.GodotStringBeginsWith = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_begins_with))
	x.GodotStringBeginsWithCharArray = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_begins_with_char_array))
	x.GodotStringBigrams = NewGodotArrayRef(unsafe.Pointer(x.ref57717e51.godot_string_bigrams))
	x.GodotStringChr = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_chr))
	x.GodotStringEndsWith = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_ends_with))
	x.GodotStringFind = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_string_find))
	x.GodotStringFindFrom = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_string_find_from))
	x.GodotStringFindmk = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_string_findmk))
	x.GodotStringFindmkFrom = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_string_findmk_from))
	x.GodotStringFindmkFromInPlace = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_string_findmk_from_in_place))
	x.GodotStringFindn = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_string_findn))
	x.GodotStringFindnFrom = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_string_findn_from))
	x.GodotStringFindLast = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_string_find_last))
	x.GodotStringFormat = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_format))
	x.GodotStringFormatWithCustomPlaceholder = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_format_with_custom_placeholder))
	x.GodotStringHexEncodeBuffer = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_hex_encode_buffer))
	x.GodotStringHexToInt = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_string_hex_to_int))
	x.GodotStringHexToIntWithoutPrefix = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_string_hex_to_int_without_prefix))
	x.GodotStringInsert = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_insert))
	x.GodotStringIsNumeric = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_is_numeric))
	x.GodotStringIsSubsequenceOf = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_is_subsequence_of))
	x.GodotStringIsSubsequenceOfi = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_is_subsequence_ofi))
	x.GodotStringLpad = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_lpad))
	x.GodotStringLpadWithCustomCharacter = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_lpad_with_custom_character))
	x.GodotStringMatch = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_match))
	x.GodotStringMatchn = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_matchn))
	x.GodotStringMd5 = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_md5))
	x.GodotStringNum = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_num))
	x.GodotStringNumInt64 = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_num_int64))
	x.GodotStringNumInt64Capitalized = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_num_int64_capitalized))
	x.GodotStringNumReal = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_num_real))
	x.GodotStringNumScientific = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_num_scientific))
	x.GodotStringNumWithDecimals = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_num_with_decimals))
	x.GodotStringPadDecimals = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_pad_decimals))
	x.GodotStringPadZeros = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_pad_zeros))
	x.GodotStringReplaceFirst = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_replace_first))
	x.GodotStringReplace = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_replace))
	x.GodotStringReplacen = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_replacen))
	x.GodotStringRfind = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_string_rfind))
	x.GodotStringRfindn = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_string_rfindn))
	x.GodotStringRfindFrom = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_string_rfind_from))
	x.GodotStringRfindnFrom = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_string_rfindn_from))
	x.GodotStringRpad = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_rpad))
	x.GodotStringRpadWithCustomCharacter = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_rpad_with_custom_character))
	x.GodotStringSimilarity = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_string_similarity))
	x.GodotStringSprintf = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_sprintf))
	x.GodotStringSubstr = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_substr))
	x.GodotStringToDouble = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_to_double))
	x.GodotStringToFloat = NewGodotRealRef(unsafe.Pointer(x.ref57717e51.godot_string_to_float))
	x.GodotStringToInt = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_string_to_int))
	x.GodotStringCamelcaseToUnderscore = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_camelcase_to_underscore))
	x.GodotStringCamelcaseToUnderscoreLowercased = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_camelcase_to_underscore_lowercased))
	x.GodotStringCapitalize = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_capitalize))
	x.GodotStringCharToDouble = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_char_to_double))
	x.GodotStringCharToInt = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_string_char_to_int))
	x.GodotStringWcharToInt = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_wchar_to_int))
	x.GodotStringCharToIntWithLen = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_string_char_to_int_with_len))
	x.GodotStringCharToInt64WithLen = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_char_to_int64_with_len))
	x.GodotStringHexToInt64 = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_hex_to_int64))
	x.GodotStringHexToInt64WithPrefix = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_hex_to_int64_with_prefix))
	x.GodotStringToInt64 = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_to_int64))
	x.GodotStringUnicodeCharToDouble = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_unicode_char_to_double))
	x.GodotStringGetSliceCount = NewGodotIntRef(unsafe.Pointer(x.ref57717e51.godot_string_get_slice_count))
	x.GodotStringGetSlice = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_get_slice))
	x.GodotStringGetSlicec = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_get_slicec))
	x.GodotStringSplit = NewGodotArrayRef(unsafe.Pointer(x.ref57717e51.godot_string_split))
	x.GodotStringSplitAllowEmpty = NewGodotArrayRef(unsafe.Pointer(x.ref57717e51.godot_string_split_allow_empty))
	x.GodotStringSplitFloats = NewGodotArrayRef(unsafe.Pointer(x.ref57717e51.godot_string_split_floats))
	x.GodotStringSplitFloatsAllowsEmpty = NewGodotArrayRef(unsafe.Pointer(x.ref57717e51.godot_string_split_floats_allows_empty))
	x.GodotStringSplitFloatsMk = NewGodotArrayRef(unsafe.Pointer(x.ref57717e51.godot_string_split_floats_mk))
	x.GodotStringSplitFloatsMkAllowsEmpty = NewGodotArrayRef(unsafe.Pointer(x.ref57717e51.godot_string_split_floats_mk_allows_empty))
	x.GodotStringSplitInts = NewGodotArrayRef(unsafe.Pointer(x.ref57717e51.godot_string_split_ints))
	x.GodotStringSplitIntsAllowsEmpty = NewGodotArrayRef(unsafe.Pointer(x.ref57717e51.godot_string_split_ints_allows_empty))
	x.GodotStringSplitIntsMk = NewGodotArrayRef(unsafe.Pointer(x.ref57717e51.godot_string_split_ints_mk))
	x.GodotStringSplitIntsMkAllowsEmpty = NewGodotArrayRef(unsafe.Pointer(x.ref57717e51.godot_string_split_ints_mk_allows_empty))
	x.GodotStringSplitSpaces = NewGodotArrayRef(unsafe.Pointer(x.ref57717e51.godot_string_split_spaces))
	x.GodotStringCharLowercase = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_char_lowercase))
	x.GodotStringCharUppercase = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_char_uppercase))
	x.GodotStringToLower = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_to_lower))
	x.GodotStringToUpper = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_to_upper))
	x.GodotStringGetBasename = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_get_basename))
	x.GodotStringGetExtension = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_get_extension))
	x.GodotStringLeft = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_left))
	x.GodotStringOrdAt = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_ord_at))
	x.GodotStringPlusFile = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_plus_file))
	x.GodotStringRight = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_right))
	x.GodotStringStripEdges = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_strip_edges))
	x.GodotStringStripEscapes = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_strip_escapes))
	x.GodotStringErase = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_erase))
	x.GodotStringAscii = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_ascii))
	x.GodotStringAsciiExtended = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_ascii_extended))
	x.GodotStringUtf8 = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_utf8))
	x.GodotStringParseUtf8 = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_parse_utf8))
	x.GodotStringParseUtf8WithLen = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_parse_utf8_with_len))
	x.GodotStringCharsToUtf8 = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_chars_to_utf8))
	x.GodotStringCharsToUtf8WithLen = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_chars_to_utf8_with_len))
	x.GodotStringHash = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_hash))
	x.GodotStringHash64 = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_hash64))
	x.GodotStringHashChars = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_hash_chars))
	x.GodotStringHashCharsWithLen = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_hash_chars_with_len))
	x.GodotStringHashUtf8Chars = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_hash_utf8_chars))
	x.GodotStringHashUtf8CharsWithLen = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_hash_utf8_chars_with_len))
	x.GodotStringMd5Buffer = NewGodotPoolByteArrayRef(unsafe.Pointer(x.ref57717e51.godot_string_md5_buffer))
	x.GodotStringMd5Text = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_md5_text))
	x.GodotStringSha256Buffer = NewGodotPoolByteArrayRef(unsafe.Pointer(x.ref57717e51.godot_string_sha256_buffer))
	x.GodotStringSha256Text = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_sha256_text))
	x.GodotStringEmpty = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_empty))
	x.GodotStringGetBaseDir = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_get_base_dir))
	x.GodotStringGetFile = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_get_file))
	x.GodotStringHumanizeSize = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_humanize_size))
	x.GodotStringIsAbsPath = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_is_abs_path))
	x.GodotStringIsRelPath = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_is_rel_path))
	x.GodotStringIsResourceFile = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_is_resource_file))
	x.GodotStringPathTo = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_path_to))
	x.GodotStringPathToFile = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_path_to_file))
	x.GodotStringSimplifyPath = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_simplify_path))
	x.GodotStringCEscape = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_c_escape))
	x.GodotStringCEscapeMultiline = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_c_escape_multiline))
	x.GodotStringCUnescape = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_c_unescape))
	x.GodotStringHttpEscape = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_http_escape))
	x.GodotStringHttpUnescape = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_http_unescape))
	x.GodotStringJsonEscape = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_json_escape))
	x.GodotStringWordWrap = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_word_wrap))
	x.GodotStringXmlEscape = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_xml_escape))
	x.GodotStringXmlEscapeWithQuotes = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_xml_escape_with_quotes))
	x.GodotStringXmlUnescape = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_xml_unescape))
	x.GodotStringPercentDecode = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_percent_decode))
	x.GodotStringPercentEncode = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_percent_encode))
	x.GodotStringIsValidFloat = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_is_valid_float))
	x.GodotStringIsValidHexNumber = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_is_valid_hex_number))
	x.GodotStringIsValidHtmlColor = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_is_valid_html_color))
	x.GodotStringIsValidIdentifier = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_is_valid_identifier))
	x.GodotStringIsValidInteger = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_is_valid_integer))
	x.GodotStringIsValidIpAddress = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_is_valid_ip_address))
	x.GodotStringDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_destroy))
	x.GodotStringNameNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_name_new))
	x.GodotStringNameNewData = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_name_new_data))
	x.GodotStringNameGetName = NewGodotStringRef(unsafe.Pointer(x.ref57717e51.godot_string_name_get_name))
	x.GodotStringNameGetHash = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_name_get_hash))
	x.GodotStringNameGetDataUniquePointer = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_name_get_data_unique_pointer))
	x.GodotStringNameOperatorEqual = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_name_operator_equal))
	x.GodotStringNameOperatorLess = NewGodotBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_name_operator_less))
	x.GodotStringNameDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_name_destroy))
	x.GodotObjectDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_object_destroy))
	x.GodotGlobalGetSingleton = NewGodotObjectRef(unsafe.Pointer(x.ref57717e51.godot_global_get_singleton))
	x.GodotMethodBindGetMethod = NewGodotMethodBindRef(unsafe.Pointer(x.ref57717e51.godot_method_bind_get_method))
	x.GodotMethodBindPtrcall = NewRef(unsafe.Pointer(x.ref57717e51.godot_method_bind_ptrcall))
	x.GodotMethodBindCall = NewGodotVariantRef(unsafe.Pointer(x.ref57717e51.godot_method_bind_call))
	x.GodotGetClassConstructor = *NewGodotClassConstructorRef(unsafe.Pointer(&x.ref57717e51.godot_get_class_constructor))
	x.GodotRegisterNativeCallType = NewRef(unsafe.Pointer(x.ref57717e51.godot_register_native_call_type))
	x.GodotAlloc = NewRef(unsafe.Pointer(x.ref57717e51.godot_alloc))
	x.GodotRealloc = NewRef(unsafe.Pointer(x.ref57717e51.godot_realloc))
	x.GodotFree = NewRef(unsafe.Pointer(x.ref57717e51.godot_free))
	x.GodotPrintError = NewRef(unsafe.Pointer(x.ref57717e51.godot_print_error))
	x.GodotPrintWarning = NewRef(unsafe.Pointer(x.ref57717e51.godot_print_warning))
	x.GodotPrint = NewRef(unsafe.Pointer(x.ref57717e51.godot_print))
}

// allocGodotMethodBindMemory allocates memory for type C.godot_method_bind in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotMethodBindMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotMethodBindValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotMethodBindValue = unsafe.Sizeof([1]C.godot_method_bind{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotMethodBind) Ref() *C.godot_method_bind {
	if x == nil {
		return nil
	}
	return x.ref3a05c0bc
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotMethodBind) Free() {
	if x != nil && x.allocs3a05c0bc != nil {
		x.allocs3a05c0bc.(*cgoAllocMap).Free()
		x.ref3a05c0bc = nil
	}
}

// NewGodotMethodBindRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotMethodBindRef(ref unsafe.Pointer) *GodotMethodBind {
	if ref == nil {
		return nil
	}
	obj := new(GodotMethodBind)
	obj.ref3a05c0bc = (*C.godot_method_bind)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotMethodBind) PassRef() (*C.godot_method_bind, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref3a05c0bc != nil {
		return x.ref3a05c0bc, nil
	}
	mem3a05c0bc := allocGodotMethodBindMemory(1)
	ref3a05c0bc := (*C.godot_method_bind)(mem3a05c0bc)
	allocs3a05c0bc := new(cgoAllocMap)
	allocs3a05c0bc.Add(mem3a05c0bc)

	var c_dont_touch_that_allocs *cgoAllocMap
	ref3a05c0bc._dont_touch_that, c_dont_touch_that_allocs = *(*[1]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocs3a05c0bc.Borrow(c_dont_touch_that_allocs)

	x.ref3a05c0bc = ref3a05c0bc
	x.allocs3a05c0bc = allocs3a05c0bc
	return ref3a05c0bc, allocs3a05c0bc

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotMethodBind) PassValue() (C.godot_method_bind, *cgoAllocMap) {
	if x.ref3a05c0bc != nil {
		return *x.ref3a05c0bc, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotMethodBind) Deref() {
	if x.ref3a05c0bc == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.ref3a05c0bc._dont_touch_that))
}

// allocGodotGdnativeApiVersionMemory allocates memory for type C.godot_gdnative_api_version in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotGdnativeApiVersionMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotGdnativeApiVersionValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotGdnativeApiVersionValue = unsafe.Sizeof([1]C.godot_gdnative_api_version{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotGdnativeApiVersion) Ref() *C.godot_gdnative_api_version {
	if x == nil {
		return nil
	}
	return x.ref5eed2c27
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotGdnativeApiVersion) Free() {
	if x != nil && x.allocs5eed2c27 != nil {
		x.allocs5eed2c27.(*cgoAllocMap).Free()
		x.ref5eed2c27 = nil
	}
}

// NewGodotGdnativeApiVersionRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotGdnativeApiVersionRef(ref unsafe.Pointer) *GodotGdnativeApiVersion {
	if ref == nil {
		return nil
	}
	obj := new(GodotGdnativeApiVersion)
	obj.ref5eed2c27 = (*C.godot_gdnative_api_version)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotGdnativeApiVersion) PassRef() (*C.godot_gdnative_api_version, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref5eed2c27 != nil {
		return x.ref5eed2c27, nil
	}
	mem5eed2c27 := allocGodotGdnativeApiVersionMemory(1)
	ref5eed2c27 := (*C.godot_gdnative_api_version)(mem5eed2c27)
	allocs5eed2c27 := new(cgoAllocMap)
	allocs5eed2c27.Add(mem5eed2c27)

	var cmajor_allocs *cgoAllocMap
	ref5eed2c27.major, cmajor_allocs = (C.uint)(x.Major), cgoAllocsUnknown
	allocs5eed2c27.Borrow(cmajor_allocs)

	var cminor_allocs *cgoAllocMap
	ref5eed2c27.minor, cminor_allocs = (C.uint)(x.Minor), cgoAllocsUnknown
	allocs5eed2c27.Borrow(cminor_allocs)

	x.ref5eed2c27 = ref5eed2c27
	x.allocs5eed2c27 = allocs5eed2c27
	return ref5eed2c27, allocs5eed2c27

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotGdnativeApiVersion) PassValue() (C.godot_gdnative_api_version, *cgoAllocMap) {
	if x.ref5eed2c27 != nil {
		return *x.ref5eed2c27, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotGdnativeApiVersion) Deref() {
	if x.ref5eed2c27 == nil {
		return
	}
	x.Major = (uint32)(x.ref5eed2c27.major)
	x.Minor = (uint32)(x.ref5eed2c27.minor)
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotGdnativeApiStruct) Ref() *C.godot_gdnative_api_struct {
	if x == nil {
		return nil
	}
	return x.ref45f52b65
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotGdnativeApiStruct) Free() {
	if x != nil && x.allocs45f52b65 != nil {
		x.allocs45f52b65.(*cgoAllocMap).Free()
		x.ref45f52b65 = nil
	}
}

// NewGodotGdnativeApiStructRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotGdnativeApiStructRef(ref unsafe.Pointer) *GodotGdnativeApiStruct {
	if ref == nil {
		return nil
	}
	obj := new(GodotGdnativeApiStruct)
	obj.ref45f52b65 = (*C.godot_gdnative_api_struct)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotGdnativeApiStruct) PassRef() (*C.godot_gdnative_api_struct, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref45f52b65 != nil {
		return x.ref45f52b65, nil
	}
	mem45f52b65 := allocGodotGdnativeApiStructMemory(1)
	ref45f52b65 := (*C.godot_gdnative_api_struct)(mem45f52b65)
	allocs45f52b65 := new(cgoAllocMap)
	allocs45f52b65.Add(mem45f52b65)

	var c_type_allocs *cgoAllocMap
	ref45f52b65._type, c_type_allocs = (C.uint)(x.Type), cgoAllocsUnknown
	allocs45f52b65.Borrow(c_type_allocs)

	var cversion_allocs *cgoAllocMap
	ref45f52b65.version, cversion_allocs = x.Version.PassValue()
	allocs45f52b65.Borrow(cversion_allocs)

	var cnext_allocs *cgoAllocMap
	ref45f52b65.next, cnext_allocs = unpackSGodotGdnativeApiStruct(x.Next)
	allocs45f52b65.Borrow(cnext_allocs)

	x.ref45f52b65 = ref45f52b65
	x.allocs45f52b65 = allocs45f52b65
	return ref45f52b65, allocs45f52b65

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotGdnativeApiStruct) PassValue() (C.godot_gdnative_api_struct, *cgoAllocMap) {
	if x.ref45f52b65 != nil {
		return *x.ref45f52b65, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotGdnativeApiStruct) Deref() {
	if x.ref45f52b65 == nil {
		return
	}
	x.Type = (uint32)(x.ref45f52b65._type)
	x.Version = *NewGodotGdnativeApiVersionRef(unsafe.Pointer(&x.ref45f52b65.version))
	packSGodotGdnativeApiStruct(x.Next, x.ref45f52b65.next)
}

// allocGodotGdnativeInitOptionsMemory allocates memory for type C.godot_gdnative_init_options in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotGdnativeInitOptionsMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotGdnativeInitOptionsValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotGdnativeInitOptionsValue = unsafe.Sizeof([1]C.godot_gdnative_init_options{})

// allocStructGodotGdnativeCoreApiStructMemory allocates memory for type C.struct_godot_gdnative_core_api_struct in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStructGodotGdnativeCoreApiStructMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStructGodotGdnativeCoreApiStructValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStructGodotGdnativeCoreApiStructValue = unsafe.Sizeof([1]C.struct_godot_gdnative_core_api_struct{})

// unpackSGodotGdnativeCoreApiStruct transforms a sliced Go data structure into plain C format.
func unpackSGodotGdnativeCoreApiStruct(x []GodotGdnativeCoreApiStruct) (unpacked *C.struct_godot_gdnative_core_api_struct, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.struct_godot_gdnative_core_api_struct) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocStructGodotGdnativeCoreApiStructMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.struct_godot_gdnative_core_api_struct)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.struct_godot_gdnative_core_api_struct)(unsafe.Pointer(h.Data))
	return
}

// allocGodotStringMemory allocates memory for type C.godot_string in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotStringMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotStringValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotStringValue = unsafe.Sizeof([1]C.godot_string{})

// unpackSGodotString transforms a sliced Go data structure into plain C format.
func unpackSGodotString(x []GodotString) (unpacked *C.godot_string, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_string) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotStringMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_string)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_string)(unsafe.Pointer(h.Data))
	return
}

// packSGodotGdnativeCoreApiStruct reads sliced Go data structure out from plain C format.
func packSGodotGdnativeCoreApiStruct(v []GodotGdnativeCoreApiStruct, ptr0 *C.struct_godot_gdnative_core_api_struct) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfStructGodotGdnativeCoreApiStructValue]C.struct_godot_gdnative_core_api_struct)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotGdnativeCoreApiStructRef(unsafe.Pointer(&ptr1))
	}
}

// packSGodotString reads sliced Go data structure out from plain C format.
func packSGodotString(v []GodotString, ptr0 *C.godot_string) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotStringValue]C.godot_string)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotStringRef(unsafe.Pointer(&ptr1))
	}
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotGdnativeInitOptions) Ref() *C.godot_gdnative_init_options {
	if x == nil {
		return nil
	}
	return x.reff9d34929
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotGdnativeInitOptions) Free() {
	if x != nil && x.allocsf9d34929 != nil {
		x.allocsf9d34929.(*cgoAllocMap).Free()
		x.reff9d34929 = nil
	}
}

// NewGodotGdnativeInitOptionsRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotGdnativeInitOptionsRef(ref unsafe.Pointer) *GodotGdnativeInitOptions {
	if ref == nil {
		return nil
	}
	obj := new(GodotGdnativeInitOptions)
	obj.reff9d34929 = (*C.godot_gdnative_init_options)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotGdnativeInitOptions) PassRef() (*C.godot_gdnative_init_options, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.reff9d34929 != nil {
		return x.reff9d34929, nil
	}
	memf9d34929 := allocGodotGdnativeInitOptionsMemory(1)
	reff9d34929 := (*C.godot_gdnative_init_options)(memf9d34929)
	allocsf9d34929 := new(cgoAllocMap)
	allocsf9d34929.Add(memf9d34929)

	var cin_editor_allocs *cgoAllocMap
	reff9d34929.in_editor, cin_editor_allocs = (C.godot_bool)(x.InEditor), cgoAllocsUnknown
	allocsf9d34929.Borrow(cin_editor_allocs)

	var ccore_api_hash_allocs *cgoAllocMap
	reff9d34929.core_api_hash, ccore_api_hash_allocs = (C.uint64_t)(x.CoreApiHash), cgoAllocsUnknown
	allocsf9d34929.Borrow(ccore_api_hash_allocs)

	var ceditor_api_hash_allocs *cgoAllocMap
	reff9d34929.editor_api_hash, ceditor_api_hash_allocs = (C.uint64_t)(x.EditorApiHash), cgoAllocsUnknown
	allocsf9d34929.Borrow(ceditor_api_hash_allocs)

	var cno_api_hash_allocs *cgoAllocMap
	reff9d34929.no_api_hash, cno_api_hash_allocs = (C.uint64_t)(x.NoApiHash), cgoAllocsUnknown
	allocsf9d34929.Borrow(cno_api_hash_allocs)

	var creport_version_mismatch_allocs *cgoAllocMap
	reff9d34929.report_version_mismatch, creport_version_mismatch_allocs = x.ReportVersionMismatch.PassRef()
	allocsf9d34929.Borrow(creport_version_mismatch_allocs)

	var creport_loading_error_allocs *cgoAllocMap
	reff9d34929.report_loading_error, creport_loading_error_allocs = x.ReportLoadingError.PassRef()
	allocsf9d34929.Borrow(creport_loading_error_allocs)

	var cgd_native_library_allocs *cgoAllocMap
	reff9d34929.gd_native_library, cgd_native_library_allocs = *(*C.godot_object)(unsafe.Pointer(&x.GdNativeLibrary)), cgoAllocsUnknown
	allocsf9d34929.Borrow(cgd_native_library_allocs)

	var capi_struct_allocs *cgoAllocMap
	reff9d34929.api_struct, capi_struct_allocs = unpackSGodotGdnativeCoreApiStruct(x.ApiStruct)
	allocsf9d34929.Borrow(capi_struct_allocs)

	var cactive_library_path_allocs *cgoAllocMap
	reff9d34929.active_library_path, cactive_library_path_allocs = unpackSGodotString(x.ActiveLibraryPath)
	allocsf9d34929.Borrow(cactive_library_path_allocs)

	x.reff9d34929 = reff9d34929
	x.allocsf9d34929 = allocsf9d34929
	return reff9d34929, allocsf9d34929

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotGdnativeInitOptions) PassValue() (C.godot_gdnative_init_options, *cgoAllocMap) {
	if x.reff9d34929 != nil {
		return *x.reff9d34929, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotGdnativeInitOptions) Deref() {
	if x.reff9d34929 == nil {
		return
	}
	x.InEditor = (GodotBool)(x.reff9d34929.in_editor)
	x.CoreApiHash = (uint64)(x.reff9d34929.core_api_hash)
	x.EditorApiHash = (uint64)(x.reff9d34929.editor_api_hash)
	x.NoApiHash = (uint64)(x.reff9d34929.no_api_hash)
	x.ReportVersionMismatch = NewRef(unsafe.Pointer(x.reff9d34929.report_version_mismatch))
	x.ReportLoadingError = NewRef(unsafe.Pointer(x.reff9d34929.report_loading_error))
	x.GdNativeLibrary = (*GodotObject)(unsafe.Pointer(x.reff9d34929.gd_native_library))
	packSGodotGdnativeCoreApiStruct(x.ApiStruct, x.reff9d34929.api_struct)
	packSGodotString(x.ActiveLibraryPath, x.reff9d34929.active_library_path)
}

// allocGodotGdnativeTerminateOptionsMemory allocates memory for type C.godot_gdnative_terminate_options in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotGdnativeTerminateOptionsMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotGdnativeTerminateOptionsValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotGdnativeTerminateOptionsValue = unsafe.Sizeof([1]C.godot_gdnative_terminate_options{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotGdnativeTerminateOptions) Ref() *C.godot_gdnative_terminate_options {
	if x == nil {
		return nil
	}
	return x.ref80c63fdc
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotGdnativeTerminateOptions) Free() {
	if x != nil && x.allocs80c63fdc != nil {
		x.allocs80c63fdc.(*cgoAllocMap).Free()
		x.ref80c63fdc = nil
	}
}

// NewGodotGdnativeTerminateOptionsRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotGdnativeTerminateOptionsRef(ref unsafe.Pointer) *GodotGdnativeTerminateOptions {
	if ref == nil {
		return nil
	}
	obj := new(GodotGdnativeTerminateOptions)
	obj.ref80c63fdc = (*C.godot_gdnative_terminate_options)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotGdnativeTerminateOptions) PassRef() (*C.godot_gdnative_terminate_options, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref80c63fdc != nil {
		return x.ref80c63fdc, nil
	}
	mem80c63fdc := allocGodotGdnativeTerminateOptionsMemory(1)
	ref80c63fdc := (*C.godot_gdnative_terminate_options)(mem80c63fdc)
	allocs80c63fdc := new(cgoAllocMap)
	allocs80c63fdc.Add(mem80c63fdc)

	var cin_editor_allocs *cgoAllocMap
	ref80c63fdc.in_editor, cin_editor_allocs = (C.godot_bool)(x.InEditor), cgoAllocsUnknown
	allocs80c63fdc.Borrow(cin_editor_allocs)

	x.ref80c63fdc = ref80c63fdc
	x.allocs80c63fdc = allocs80c63fdc
	return ref80c63fdc, allocs80c63fdc

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotGdnativeTerminateOptions) PassValue() (C.godot_gdnative_terminate_options, *cgoAllocMap) {
	if x.ref80c63fdc != nil {
		return *x.ref80c63fdc, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotGdnativeTerminateOptions) Deref() {
	if x.ref80c63fdc == nil {
		return
	}
	x.InEditor = (GodotBool)(x.ref80c63fdc.in_editor)
}

func (x GodotClassConstructor) PassRef() (ref *C.godot_class_constructor, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if godotClassConstructorC099ECDEFunc == nil {
		godotClassConstructorC099ECDEFunc = x
	}
	return (*C.godot_class_constructor)(C.godot_class_constructor_c099ecde), nil
}

func (x GodotClassConstructor) PassValue() (ref C.godot_class_constructor, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if godotClassConstructorC099ECDEFunc == nil {
		godotClassConstructorC099ECDEFunc = x
	}
	return (C.godot_class_constructor)(C.godot_class_constructor_c099ecde), nil
}

func NewGodotClassConstructorRef(ref unsafe.Pointer) *GodotClassConstructor {
	return (*GodotClassConstructor)(ref)
}

//export godotClassConstructorC099ECDE
func godotClassConstructorC099ECDE() C.godot_object {
	if godotClassConstructorC099ECDEFunc != nil {
		retc099ecde := godotClassConstructorC099ECDEFunc()
		ret, _ := (C.godot_object)(unsafe.Pointer(retc099ecde)), cgoAllocsUnknown
		return ret
	}
	panic("callback func has not been set (race?)")
}

var godotClassConstructorC099ECDEFunc GodotClassConstructor

// packSGodotGdnativeInitOptions reads sliced Go data structure out from plain C format.
func packSGodotGdnativeInitOptions(v []GodotGdnativeInitOptions, ptr0 *C.godot_gdnative_init_options) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotGdnativeInitOptionsValue]C.godot_gdnative_init_options)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotGdnativeInitOptionsRef(unsafe.Pointer(&ptr1))
	}
}

func (x GodotGdnativeInitFn) PassRef() (ref *C.godot_gdnative_init_fn, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if godotGdnativeInitFnF0F9A077Func == nil {
		godotGdnativeInitFnF0F9A077Func = x
	}
	return (*C.godot_gdnative_init_fn)(C.godot_gdnative_init_fn_f0f9a077), nil
}

func (x GodotGdnativeInitFn) PassValue() (ref C.godot_gdnative_init_fn, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if godotGdnativeInitFnF0F9A077Func == nil {
		godotGdnativeInitFnF0F9A077Func = x
	}
	return (C.godot_gdnative_init_fn)(C.godot_gdnative_init_fn_f0f9a077), nil
}

func NewGodotGdnativeInitFnRef(ref unsafe.Pointer) *GodotGdnativeInitFn {
	return (*GodotGdnativeInitFn)(ref)
}

//export godotGdnativeInitFnF0F9A077
func godotGdnativeInitFnF0F9A077(carg0 *C.godot_gdnative_init_options) {
	if godotGdnativeInitFnF0F9A077Func != nil {
		var arg0f0f9a077 []GodotGdnativeInitOptions
		packSGodotGdnativeInitOptions(arg0f0f9a077, carg0)
		godotGdnativeInitFnF0F9A077Func(arg0f0f9a077)
		return
	}
	panic("callback func has not been set (race?)")
}

var godotGdnativeInitFnF0F9A077Func GodotGdnativeInitFn

// packSGodotGdnativeTerminateOptions reads sliced Go data structure out from plain C format.
func packSGodotGdnativeTerminateOptions(v []GodotGdnativeTerminateOptions, ptr0 *C.godot_gdnative_terminate_options) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotGdnativeTerminateOptionsValue]C.godot_gdnative_terminate_options)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotGdnativeTerminateOptionsRef(unsafe.Pointer(&ptr1))
	}
}

func (x GodotGdnativeTerminateFn) PassRef() (ref *C.godot_gdnative_terminate_fn, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if godotGdnativeTerminateFn7F5590C3Func == nil {
		godotGdnativeTerminateFn7F5590C3Func = x
	}
	return (*C.godot_gdnative_terminate_fn)(C.godot_gdnative_terminate_fn_7f5590c3), nil
}

func (x GodotGdnativeTerminateFn) PassValue() (ref C.godot_gdnative_terminate_fn, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if godotGdnativeTerminateFn7F5590C3Func == nil {
		godotGdnativeTerminateFn7F5590C3Func = x
	}
	return (C.godot_gdnative_terminate_fn)(C.godot_gdnative_terminate_fn_7f5590c3), nil
}

func NewGodotGdnativeTerminateFnRef(ref unsafe.Pointer) *GodotGdnativeTerminateFn {
	return (*GodotGdnativeTerminateFn)(ref)
}

//export godotGdnativeTerminateFn7F5590C3
func godotGdnativeTerminateFn7F5590C3(carg0 *C.godot_gdnative_terminate_options) {
	if godotGdnativeTerminateFn7F5590C3Func != nil {
		var arg07f5590c3 []GodotGdnativeTerminateOptions
		packSGodotGdnativeTerminateOptions(arg07f5590c3, carg0)
		godotGdnativeTerminateFn7F5590C3Func(arg07f5590c3)
		return
	}
	panic("callback func has not been set (race?)")
}

var godotGdnativeTerminateFn7F5590C3Func GodotGdnativeTerminateFn

// packSGodotArray reads sliced Go data structure out from plain C format.
func packSGodotArray(v []GodotArray, ptr0 *C.godot_array) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotArrayValue]C.godot_array)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotArrayRef(unsafe.Pointer(&ptr1))
	}
}

func (x GodotGdnativeProcedureFn) PassRef() (ref *C.godot_gdnative_procedure_fn, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if godotGdnativeProcedureFn5D45412Func == nil {
		godotGdnativeProcedureFn5D45412Func = x
	}
	return (*C.godot_gdnative_procedure_fn)(C.godot_gdnative_procedure_fn_5d45412), nil
}

func (x GodotGdnativeProcedureFn) PassValue() (ref C.godot_gdnative_procedure_fn, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if godotGdnativeProcedureFn5D45412Func == nil {
		godotGdnativeProcedureFn5D45412Func = x
	}
	return (C.godot_gdnative_procedure_fn)(C.godot_gdnative_procedure_fn_5d45412), nil
}

func NewGodotGdnativeProcedureFnRef(ref unsafe.Pointer) *GodotGdnativeProcedureFn {
	return (*GodotGdnativeProcedureFn)(ref)
}

//export godotGdnativeProcedureFn5D45412
func godotGdnativeProcedureFn5D45412(carg0 *C.godot_array) C.godot_variant {
	if godotGdnativeProcedureFn5D45412Func != nil {
		var arg05d45412 []GodotArray
		packSGodotArray(arg05d45412, carg0)
		ret5d45412 := godotGdnativeProcedureFn5D45412Func(arg05d45412)
		ret, _ := ret5d45412.PassValue()
		return ret
	}
	panic("callback func has not been set (race?)")
}

var godotGdnativeProcedureFn5D45412Func GodotGdnativeProcedureFn

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotString) Ref() *C.godot_string {
	if x == nil {
		return nil
	}
	return x.ref6d1ede57
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotString) Free() {
	if x != nil && x.allocs6d1ede57 != nil {
		x.allocs6d1ede57.(*cgoAllocMap).Free()
		x.ref6d1ede57 = nil
	}
}

// NewGodotStringRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotStringRef(ref unsafe.Pointer) *GodotString {
	if ref == nil {
		return nil
	}
	obj := new(GodotString)
	obj.ref6d1ede57 = (*C.godot_string)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotString) PassRef() (*C.godot_string, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref6d1ede57 != nil {
		return x.ref6d1ede57, nil
	}
	mem6d1ede57 := allocGodotStringMemory(1)
	ref6d1ede57 := (*C.godot_string)(mem6d1ede57)
	allocs6d1ede57 := new(cgoAllocMap)
	allocs6d1ede57.Add(mem6d1ede57)

	var c_dont_touch_that_allocs *cgoAllocMap
	ref6d1ede57._dont_touch_that, c_dont_touch_that_allocs = *(*[8]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocs6d1ede57.Borrow(c_dont_touch_that_allocs)

	x.ref6d1ede57 = ref6d1ede57
	x.allocs6d1ede57 = allocs6d1ede57
	return ref6d1ede57, allocs6d1ede57

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotString) PassValue() (C.godot_string, *cgoAllocMap) {
	if x.ref6d1ede57 != nil {
		return *x.ref6d1ede57, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotString) Deref() {
	if x.ref6d1ede57 == nil {
		return
	}
	x.DontTouchThat = *(*[8]byte)(unsafe.Pointer(&x.ref6d1ede57._dont_touch_that))
}

// allocGodotArrayMemory allocates memory for type C.godot_array in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotArrayMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotArrayValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotArrayValue = unsafe.Sizeof([1]C.godot_array{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotArray) Ref() *C.godot_array {
	if x == nil {
		return nil
	}
	return x.refb81158a5
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotArray) Free() {
	if x != nil && x.allocsb81158a5 != nil {
		x.allocsb81158a5.(*cgoAllocMap).Free()
		x.refb81158a5 = nil
	}
}

// NewGodotArrayRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotArrayRef(ref unsafe.Pointer) *GodotArray {
	if ref == nil {
		return nil
	}
	obj := new(GodotArray)
	obj.refb81158a5 = (*C.godot_array)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotArray) PassRef() (*C.godot_array, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refb81158a5 != nil {
		return x.refb81158a5, nil
	}
	memb81158a5 := allocGodotArrayMemory(1)
	refb81158a5 := (*C.godot_array)(memb81158a5)
	allocsb81158a5 := new(cgoAllocMap)
	allocsb81158a5.Add(memb81158a5)

	var c_dont_touch_that_allocs *cgoAllocMap
	refb81158a5._dont_touch_that, c_dont_touch_that_allocs = *(*[8]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocsb81158a5.Borrow(c_dont_touch_that_allocs)

	x.refb81158a5 = refb81158a5
	x.allocsb81158a5 = allocsb81158a5
	return refb81158a5, allocsb81158a5

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotArray) PassValue() (C.godot_array, *cgoAllocMap) {
	if x.refb81158a5 != nil {
		return *x.refb81158a5, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotArray) Deref() {
	if x.refb81158a5 == nil {
		return
	}
	x.DontTouchThat = *(*[8]byte)(unsafe.Pointer(&x.refb81158a5._dont_touch_that))
}

// allocGodotPoolArrayReadAccessMemory allocates memory for type C.godot_pool_array_read_access in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotPoolArrayReadAccessMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotPoolArrayReadAccessValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotPoolArrayReadAccessValue = unsafe.Sizeof([1]C.godot_pool_array_read_access{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotPoolArrayReadAccess) Ref() *C.godot_pool_array_read_access {
	if x == nil {
		return nil
	}
	return x.ref172179be
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotPoolArrayReadAccess) Free() {
	if x != nil && x.allocs172179be != nil {
		x.allocs172179be.(*cgoAllocMap).Free()
		x.ref172179be = nil
	}
}

// NewGodotPoolArrayReadAccessRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotPoolArrayReadAccessRef(ref unsafe.Pointer) *GodotPoolArrayReadAccess {
	if ref == nil {
		return nil
	}
	obj := new(GodotPoolArrayReadAccess)
	obj.ref172179be = (*C.godot_pool_array_read_access)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotPoolArrayReadAccess) PassRef() (*C.godot_pool_array_read_access, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref172179be != nil {
		return x.ref172179be, nil
	}
	mem172179be := allocGodotPoolArrayReadAccessMemory(1)
	ref172179be := (*C.godot_pool_array_read_access)(mem172179be)
	allocs172179be := new(cgoAllocMap)
	allocs172179be.Add(mem172179be)

	var c_dont_touch_that_allocs *cgoAllocMap
	ref172179be._dont_touch_that, c_dont_touch_that_allocs = *(*[1]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocs172179be.Borrow(c_dont_touch_that_allocs)

	x.ref172179be = ref172179be
	x.allocs172179be = allocs172179be
	return ref172179be, allocs172179be

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotPoolArrayReadAccess) PassValue() (C.godot_pool_array_read_access, *cgoAllocMap) {
	if x.ref172179be != nil {
		return *x.ref172179be, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotPoolArrayReadAccess) Deref() {
	if x.ref172179be == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.ref172179be._dont_touch_that))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotPoolByteArrayReadAccess) Ref() *C.godot_pool_array_read_access {
	if x == nil {
		return nil
	}
	return x.ref172179be
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotPoolByteArrayReadAccess) Free() {
	if x != nil && x.allocs172179be != nil {
		x.allocs172179be.(*cgoAllocMap).Free()
		x.ref172179be = nil
	}
}

// NewGodotPoolByteArrayReadAccessRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotPoolByteArrayReadAccessRef(ref unsafe.Pointer) *GodotPoolByteArrayReadAccess {
	if ref == nil {
		return nil
	}
	obj := new(GodotPoolByteArrayReadAccess)
	obj.ref172179be = (*C.godot_pool_array_read_access)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotPoolByteArrayReadAccess) PassRef() (*C.godot_pool_array_read_access, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref172179be != nil {
		return x.ref172179be, nil
	}
	mem172179be := allocGodotPoolArrayReadAccessMemory(1)
	ref172179be := (*C.godot_pool_array_read_access)(mem172179be)
	allocs172179be := new(cgoAllocMap)
	allocs172179be.Add(mem172179be)

	var c_dont_touch_that_allocs *cgoAllocMap
	ref172179be._dont_touch_that, c_dont_touch_that_allocs = *(*[1]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocs172179be.Borrow(c_dont_touch_that_allocs)

	x.ref172179be = ref172179be
	x.allocs172179be = allocs172179be
	return ref172179be, allocs172179be

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotPoolByteArrayReadAccess) PassValue() (C.godot_pool_array_read_access, *cgoAllocMap) {
	if x.ref172179be != nil {
		return *x.ref172179be, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotPoolByteArrayReadAccess) Deref() {
	if x.ref172179be == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.ref172179be._dont_touch_that))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotPoolIntArrayReadAccess) Ref() *C.godot_pool_array_read_access {
	if x == nil {
		return nil
	}
	return x.ref172179be
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotPoolIntArrayReadAccess) Free() {
	if x != nil && x.allocs172179be != nil {
		x.allocs172179be.(*cgoAllocMap).Free()
		x.ref172179be = nil
	}
}

// NewGodotPoolIntArrayReadAccessRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotPoolIntArrayReadAccessRef(ref unsafe.Pointer) *GodotPoolIntArrayReadAccess {
	if ref == nil {
		return nil
	}
	obj := new(GodotPoolIntArrayReadAccess)
	obj.ref172179be = (*C.godot_pool_array_read_access)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotPoolIntArrayReadAccess) PassRef() (*C.godot_pool_array_read_access, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref172179be != nil {
		return x.ref172179be, nil
	}
	mem172179be := allocGodotPoolArrayReadAccessMemory(1)
	ref172179be := (*C.godot_pool_array_read_access)(mem172179be)
	allocs172179be := new(cgoAllocMap)
	allocs172179be.Add(mem172179be)

	var c_dont_touch_that_allocs *cgoAllocMap
	ref172179be._dont_touch_that, c_dont_touch_that_allocs = *(*[1]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocs172179be.Borrow(c_dont_touch_that_allocs)

	x.ref172179be = ref172179be
	x.allocs172179be = allocs172179be
	return ref172179be, allocs172179be

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotPoolIntArrayReadAccess) PassValue() (C.godot_pool_array_read_access, *cgoAllocMap) {
	if x.ref172179be != nil {
		return *x.ref172179be, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotPoolIntArrayReadAccess) Deref() {
	if x.ref172179be == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.ref172179be._dont_touch_that))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotPoolRealArrayReadAccess) Ref() *C.godot_pool_array_read_access {
	if x == nil {
		return nil
	}
	return x.ref172179be
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotPoolRealArrayReadAccess) Free() {
	if x != nil && x.allocs172179be != nil {
		x.allocs172179be.(*cgoAllocMap).Free()
		x.ref172179be = nil
	}
}

// NewGodotPoolRealArrayReadAccessRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotPoolRealArrayReadAccessRef(ref unsafe.Pointer) *GodotPoolRealArrayReadAccess {
	if ref == nil {
		return nil
	}
	obj := new(GodotPoolRealArrayReadAccess)
	obj.ref172179be = (*C.godot_pool_array_read_access)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotPoolRealArrayReadAccess) PassRef() (*C.godot_pool_array_read_access, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref172179be != nil {
		return x.ref172179be, nil
	}
	mem172179be := allocGodotPoolArrayReadAccessMemory(1)
	ref172179be := (*C.godot_pool_array_read_access)(mem172179be)
	allocs172179be := new(cgoAllocMap)
	allocs172179be.Add(mem172179be)

	var c_dont_touch_that_allocs *cgoAllocMap
	ref172179be._dont_touch_that, c_dont_touch_that_allocs = *(*[1]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocs172179be.Borrow(c_dont_touch_that_allocs)

	x.ref172179be = ref172179be
	x.allocs172179be = allocs172179be
	return ref172179be, allocs172179be

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotPoolRealArrayReadAccess) PassValue() (C.godot_pool_array_read_access, *cgoAllocMap) {
	if x.ref172179be != nil {
		return *x.ref172179be, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotPoolRealArrayReadAccess) Deref() {
	if x.ref172179be == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.ref172179be._dont_touch_that))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotPoolStringArrayReadAccess) Ref() *C.godot_pool_array_read_access {
	if x == nil {
		return nil
	}
	return x.ref172179be
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotPoolStringArrayReadAccess) Free() {
	if x != nil && x.allocs172179be != nil {
		x.allocs172179be.(*cgoAllocMap).Free()
		x.ref172179be = nil
	}
}

// NewGodotPoolStringArrayReadAccessRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotPoolStringArrayReadAccessRef(ref unsafe.Pointer) *GodotPoolStringArrayReadAccess {
	if ref == nil {
		return nil
	}
	obj := new(GodotPoolStringArrayReadAccess)
	obj.ref172179be = (*C.godot_pool_array_read_access)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotPoolStringArrayReadAccess) PassRef() (*C.godot_pool_array_read_access, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref172179be != nil {
		return x.ref172179be, nil
	}
	mem172179be := allocGodotPoolArrayReadAccessMemory(1)
	ref172179be := (*C.godot_pool_array_read_access)(mem172179be)
	allocs172179be := new(cgoAllocMap)
	allocs172179be.Add(mem172179be)

	var c_dont_touch_that_allocs *cgoAllocMap
	ref172179be._dont_touch_that, c_dont_touch_that_allocs = *(*[1]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocs172179be.Borrow(c_dont_touch_that_allocs)

	x.ref172179be = ref172179be
	x.allocs172179be = allocs172179be
	return ref172179be, allocs172179be

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotPoolStringArrayReadAccess) PassValue() (C.godot_pool_array_read_access, *cgoAllocMap) {
	if x.ref172179be != nil {
		return *x.ref172179be, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotPoolStringArrayReadAccess) Deref() {
	if x.ref172179be == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.ref172179be._dont_touch_that))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotPoolVector2ArrayReadAccess) Ref() *C.godot_pool_array_read_access {
	if x == nil {
		return nil
	}
	return x.ref172179be
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotPoolVector2ArrayReadAccess) Free() {
	if x != nil && x.allocs172179be != nil {
		x.allocs172179be.(*cgoAllocMap).Free()
		x.ref172179be = nil
	}
}

// NewGodotPoolVector2ArrayReadAccessRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotPoolVector2ArrayReadAccessRef(ref unsafe.Pointer) *GodotPoolVector2ArrayReadAccess {
	if ref == nil {
		return nil
	}
	obj := new(GodotPoolVector2ArrayReadAccess)
	obj.ref172179be = (*C.godot_pool_array_read_access)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotPoolVector2ArrayReadAccess) PassRef() (*C.godot_pool_array_read_access, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref172179be != nil {
		return x.ref172179be, nil
	}
	mem172179be := allocGodotPoolArrayReadAccessMemory(1)
	ref172179be := (*C.godot_pool_array_read_access)(mem172179be)
	allocs172179be := new(cgoAllocMap)
	allocs172179be.Add(mem172179be)

	var c_dont_touch_that_allocs *cgoAllocMap
	ref172179be._dont_touch_that, c_dont_touch_that_allocs = *(*[1]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocs172179be.Borrow(c_dont_touch_that_allocs)

	x.ref172179be = ref172179be
	x.allocs172179be = allocs172179be
	return ref172179be, allocs172179be

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotPoolVector2ArrayReadAccess) PassValue() (C.godot_pool_array_read_access, *cgoAllocMap) {
	if x.ref172179be != nil {
		return *x.ref172179be, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotPoolVector2ArrayReadAccess) Deref() {
	if x.ref172179be == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.ref172179be._dont_touch_that))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotPoolVector3ArrayReadAccess) Ref() *C.godot_pool_array_read_access {
	if x == nil {
		return nil
	}
	return x.ref172179be
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotPoolVector3ArrayReadAccess) Free() {
	if x != nil && x.allocs172179be != nil {
		x.allocs172179be.(*cgoAllocMap).Free()
		x.ref172179be = nil
	}
}

// NewGodotPoolVector3ArrayReadAccessRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotPoolVector3ArrayReadAccessRef(ref unsafe.Pointer) *GodotPoolVector3ArrayReadAccess {
	if ref == nil {
		return nil
	}
	obj := new(GodotPoolVector3ArrayReadAccess)
	obj.ref172179be = (*C.godot_pool_array_read_access)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotPoolVector3ArrayReadAccess) PassRef() (*C.godot_pool_array_read_access, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref172179be != nil {
		return x.ref172179be, nil
	}
	mem172179be := allocGodotPoolArrayReadAccessMemory(1)
	ref172179be := (*C.godot_pool_array_read_access)(mem172179be)
	allocs172179be := new(cgoAllocMap)
	allocs172179be.Add(mem172179be)

	var c_dont_touch_that_allocs *cgoAllocMap
	ref172179be._dont_touch_that, c_dont_touch_that_allocs = *(*[1]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocs172179be.Borrow(c_dont_touch_that_allocs)

	x.ref172179be = ref172179be
	x.allocs172179be = allocs172179be
	return ref172179be, allocs172179be

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotPoolVector3ArrayReadAccess) PassValue() (C.godot_pool_array_read_access, *cgoAllocMap) {
	if x.ref172179be != nil {
		return *x.ref172179be, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotPoolVector3ArrayReadAccess) Deref() {
	if x.ref172179be == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.ref172179be._dont_touch_that))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotPoolColorArrayReadAccess) Ref() *C.godot_pool_array_read_access {
	if x == nil {
		return nil
	}
	return x.ref172179be
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotPoolColorArrayReadAccess) Free() {
	if x != nil && x.allocs172179be != nil {
		x.allocs172179be.(*cgoAllocMap).Free()
		x.ref172179be = nil
	}
}

// NewGodotPoolColorArrayReadAccessRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotPoolColorArrayReadAccessRef(ref unsafe.Pointer) *GodotPoolColorArrayReadAccess {
	if ref == nil {
		return nil
	}
	obj := new(GodotPoolColorArrayReadAccess)
	obj.ref172179be = (*C.godot_pool_array_read_access)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotPoolColorArrayReadAccess) PassRef() (*C.godot_pool_array_read_access, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref172179be != nil {
		return x.ref172179be, nil
	}
	mem172179be := allocGodotPoolArrayReadAccessMemory(1)
	ref172179be := (*C.godot_pool_array_read_access)(mem172179be)
	allocs172179be := new(cgoAllocMap)
	allocs172179be.Add(mem172179be)

	var c_dont_touch_that_allocs *cgoAllocMap
	ref172179be._dont_touch_that, c_dont_touch_that_allocs = *(*[1]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocs172179be.Borrow(c_dont_touch_that_allocs)

	x.ref172179be = ref172179be
	x.allocs172179be = allocs172179be
	return ref172179be, allocs172179be

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotPoolColorArrayReadAccess) PassValue() (C.godot_pool_array_read_access, *cgoAllocMap) {
	if x.ref172179be != nil {
		return *x.ref172179be, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotPoolColorArrayReadAccess) Deref() {
	if x.ref172179be == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.ref172179be._dont_touch_that))
}

// allocGodotPoolArrayWriteAccessMemory allocates memory for type C.godot_pool_array_write_access in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotPoolArrayWriteAccessMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotPoolArrayWriteAccessValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotPoolArrayWriteAccessValue = unsafe.Sizeof([1]C.godot_pool_array_write_access{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotPoolArrayWriteAccess) Ref() *C.godot_pool_array_write_access {
	if x == nil {
		return nil
	}
	return x.refbe0648b
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotPoolArrayWriteAccess) Free() {
	if x != nil && x.allocsbe0648b != nil {
		x.allocsbe0648b.(*cgoAllocMap).Free()
		x.refbe0648b = nil
	}
}

// NewGodotPoolArrayWriteAccessRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotPoolArrayWriteAccessRef(ref unsafe.Pointer) *GodotPoolArrayWriteAccess {
	if ref == nil {
		return nil
	}
	obj := new(GodotPoolArrayWriteAccess)
	obj.refbe0648b = (*C.godot_pool_array_write_access)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotPoolArrayWriteAccess) PassRef() (*C.godot_pool_array_write_access, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refbe0648b != nil {
		return x.refbe0648b, nil
	}
	membe0648b := allocGodotPoolArrayWriteAccessMemory(1)
	refbe0648b := (*C.godot_pool_array_write_access)(membe0648b)
	allocsbe0648b := new(cgoAllocMap)
	allocsbe0648b.Add(membe0648b)

	var c_dont_touch_that_allocs *cgoAllocMap
	refbe0648b._dont_touch_that, c_dont_touch_that_allocs = *(*[1]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocsbe0648b.Borrow(c_dont_touch_that_allocs)

	x.refbe0648b = refbe0648b
	x.allocsbe0648b = allocsbe0648b
	return refbe0648b, allocsbe0648b

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotPoolArrayWriteAccess) PassValue() (C.godot_pool_array_write_access, *cgoAllocMap) {
	if x.refbe0648b != nil {
		return *x.refbe0648b, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotPoolArrayWriteAccess) Deref() {
	if x.refbe0648b == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.refbe0648b._dont_touch_that))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotPoolByteArrayWriteAccess) Ref() *C.godot_pool_array_write_access {
	if x == nil {
		return nil
	}
	return x.refbe0648b
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotPoolByteArrayWriteAccess) Free() {
	if x != nil && x.allocsbe0648b != nil {
		x.allocsbe0648b.(*cgoAllocMap).Free()
		x.refbe0648b = nil
	}
}

// NewGodotPoolByteArrayWriteAccessRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotPoolByteArrayWriteAccessRef(ref unsafe.Pointer) *GodotPoolByteArrayWriteAccess {
	if ref == nil {
		return nil
	}
	obj := new(GodotPoolByteArrayWriteAccess)
	obj.refbe0648b = (*C.godot_pool_array_write_access)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotPoolByteArrayWriteAccess) PassRef() (*C.godot_pool_array_write_access, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refbe0648b != nil {
		return x.refbe0648b, nil
	}
	membe0648b := allocGodotPoolArrayWriteAccessMemory(1)
	refbe0648b := (*C.godot_pool_array_write_access)(membe0648b)
	allocsbe0648b := new(cgoAllocMap)
	allocsbe0648b.Add(membe0648b)

	var c_dont_touch_that_allocs *cgoAllocMap
	refbe0648b._dont_touch_that, c_dont_touch_that_allocs = *(*[1]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocsbe0648b.Borrow(c_dont_touch_that_allocs)

	x.refbe0648b = refbe0648b
	x.allocsbe0648b = allocsbe0648b
	return refbe0648b, allocsbe0648b

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotPoolByteArrayWriteAccess) PassValue() (C.godot_pool_array_write_access, *cgoAllocMap) {
	if x.refbe0648b != nil {
		return *x.refbe0648b, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotPoolByteArrayWriteAccess) Deref() {
	if x.refbe0648b == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.refbe0648b._dont_touch_that))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotPoolIntArrayWriteAccess) Ref() *C.godot_pool_array_write_access {
	if x == nil {
		return nil
	}
	return x.refbe0648b
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotPoolIntArrayWriteAccess) Free() {
	if x != nil && x.allocsbe0648b != nil {
		x.allocsbe0648b.(*cgoAllocMap).Free()
		x.refbe0648b = nil
	}
}

// NewGodotPoolIntArrayWriteAccessRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotPoolIntArrayWriteAccessRef(ref unsafe.Pointer) *GodotPoolIntArrayWriteAccess {
	if ref == nil {
		return nil
	}
	obj := new(GodotPoolIntArrayWriteAccess)
	obj.refbe0648b = (*C.godot_pool_array_write_access)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotPoolIntArrayWriteAccess) PassRef() (*C.godot_pool_array_write_access, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refbe0648b != nil {
		return x.refbe0648b, nil
	}
	membe0648b := allocGodotPoolArrayWriteAccessMemory(1)
	refbe0648b := (*C.godot_pool_array_write_access)(membe0648b)
	allocsbe0648b := new(cgoAllocMap)
	allocsbe0648b.Add(membe0648b)

	var c_dont_touch_that_allocs *cgoAllocMap
	refbe0648b._dont_touch_that, c_dont_touch_that_allocs = *(*[1]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocsbe0648b.Borrow(c_dont_touch_that_allocs)

	x.refbe0648b = refbe0648b
	x.allocsbe0648b = allocsbe0648b
	return refbe0648b, allocsbe0648b

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotPoolIntArrayWriteAccess) PassValue() (C.godot_pool_array_write_access, *cgoAllocMap) {
	if x.refbe0648b != nil {
		return *x.refbe0648b, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotPoolIntArrayWriteAccess) Deref() {
	if x.refbe0648b == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.refbe0648b._dont_touch_that))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotPoolRealArrayWriteAccess) Ref() *C.godot_pool_array_write_access {
	if x == nil {
		return nil
	}
	return x.refbe0648b
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotPoolRealArrayWriteAccess) Free() {
	if x != nil && x.allocsbe0648b != nil {
		x.allocsbe0648b.(*cgoAllocMap).Free()
		x.refbe0648b = nil
	}
}

// NewGodotPoolRealArrayWriteAccessRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotPoolRealArrayWriteAccessRef(ref unsafe.Pointer) *GodotPoolRealArrayWriteAccess {
	if ref == nil {
		return nil
	}
	obj := new(GodotPoolRealArrayWriteAccess)
	obj.refbe0648b = (*C.godot_pool_array_write_access)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotPoolRealArrayWriteAccess) PassRef() (*C.godot_pool_array_write_access, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refbe0648b != nil {
		return x.refbe0648b, nil
	}
	membe0648b := allocGodotPoolArrayWriteAccessMemory(1)
	refbe0648b := (*C.godot_pool_array_write_access)(membe0648b)
	allocsbe0648b := new(cgoAllocMap)
	allocsbe0648b.Add(membe0648b)

	var c_dont_touch_that_allocs *cgoAllocMap
	refbe0648b._dont_touch_that, c_dont_touch_that_allocs = *(*[1]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocsbe0648b.Borrow(c_dont_touch_that_allocs)

	x.refbe0648b = refbe0648b
	x.allocsbe0648b = allocsbe0648b
	return refbe0648b, allocsbe0648b

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotPoolRealArrayWriteAccess) PassValue() (C.godot_pool_array_write_access, *cgoAllocMap) {
	if x.refbe0648b != nil {
		return *x.refbe0648b, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotPoolRealArrayWriteAccess) Deref() {
	if x.refbe0648b == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.refbe0648b._dont_touch_that))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotPoolStringArrayWriteAccess) Ref() *C.godot_pool_array_write_access {
	if x == nil {
		return nil
	}
	return x.refbe0648b
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotPoolStringArrayWriteAccess) Free() {
	if x != nil && x.allocsbe0648b != nil {
		x.allocsbe0648b.(*cgoAllocMap).Free()
		x.refbe0648b = nil
	}
}

// NewGodotPoolStringArrayWriteAccessRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotPoolStringArrayWriteAccessRef(ref unsafe.Pointer) *GodotPoolStringArrayWriteAccess {
	if ref == nil {
		return nil
	}
	obj := new(GodotPoolStringArrayWriteAccess)
	obj.refbe0648b = (*C.godot_pool_array_write_access)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotPoolStringArrayWriteAccess) PassRef() (*C.godot_pool_array_write_access, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refbe0648b != nil {
		return x.refbe0648b, nil
	}
	membe0648b := allocGodotPoolArrayWriteAccessMemory(1)
	refbe0648b := (*C.godot_pool_array_write_access)(membe0648b)
	allocsbe0648b := new(cgoAllocMap)
	allocsbe0648b.Add(membe0648b)

	var c_dont_touch_that_allocs *cgoAllocMap
	refbe0648b._dont_touch_that, c_dont_touch_that_allocs = *(*[1]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocsbe0648b.Borrow(c_dont_touch_that_allocs)

	x.refbe0648b = refbe0648b
	x.allocsbe0648b = allocsbe0648b
	return refbe0648b, allocsbe0648b

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotPoolStringArrayWriteAccess) PassValue() (C.godot_pool_array_write_access, *cgoAllocMap) {
	if x.refbe0648b != nil {
		return *x.refbe0648b, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotPoolStringArrayWriteAccess) Deref() {
	if x.refbe0648b == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.refbe0648b._dont_touch_that))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotPoolVector2ArrayWriteAccess) Ref() *C.godot_pool_array_write_access {
	if x == nil {
		return nil
	}
	return x.refbe0648b
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotPoolVector2ArrayWriteAccess) Free() {
	if x != nil && x.allocsbe0648b != nil {
		x.allocsbe0648b.(*cgoAllocMap).Free()
		x.refbe0648b = nil
	}
}

// NewGodotPoolVector2ArrayWriteAccessRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotPoolVector2ArrayWriteAccessRef(ref unsafe.Pointer) *GodotPoolVector2ArrayWriteAccess {
	if ref == nil {
		return nil
	}
	obj := new(GodotPoolVector2ArrayWriteAccess)
	obj.refbe0648b = (*C.godot_pool_array_write_access)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotPoolVector2ArrayWriteAccess) PassRef() (*C.godot_pool_array_write_access, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refbe0648b != nil {
		return x.refbe0648b, nil
	}
	membe0648b := allocGodotPoolArrayWriteAccessMemory(1)
	refbe0648b := (*C.godot_pool_array_write_access)(membe0648b)
	allocsbe0648b := new(cgoAllocMap)
	allocsbe0648b.Add(membe0648b)

	var c_dont_touch_that_allocs *cgoAllocMap
	refbe0648b._dont_touch_that, c_dont_touch_that_allocs = *(*[1]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocsbe0648b.Borrow(c_dont_touch_that_allocs)

	x.refbe0648b = refbe0648b
	x.allocsbe0648b = allocsbe0648b
	return refbe0648b, allocsbe0648b

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotPoolVector2ArrayWriteAccess) PassValue() (C.godot_pool_array_write_access, *cgoAllocMap) {
	if x.refbe0648b != nil {
		return *x.refbe0648b, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotPoolVector2ArrayWriteAccess) Deref() {
	if x.refbe0648b == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.refbe0648b._dont_touch_that))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotPoolVector3ArrayWriteAccess) Ref() *C.godot_pool_array_write_access {
	if x == nil {
		return nil
	}
	return x.refbe0648b
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotPoolVector3ArrayWriteAccess) Free() {
	if x != nil && x.allocsbe0648b != nil {
		x.allocsbe0648b.(*cgoAllocMap).Free()
		x.refbe0648b = nil
	}
}

// NewGodotPoolVector3ArrayWriteAccessRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotPoolVector3ArrayWriteAccessRef(ref unsafe.Pointer) *GodotPoolVector3ArrayWriteAccess {
	if ref == nil {
		return nil
	}
	obj := new(GodotPoolVector3ArrayWriteAccess)
	obj.refbe0648b = (*C.godot_pool_array_write_access)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotPoolVector3ArrayWriteAccess) PassRef() (*C.godot_pool_array_write_access, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refbe0648b != nil {
		return x.refbe0648b, nil
	}
	membe0648b := allocGodotPoolArrayWriteAccessMemory(1)
	refbe0648b := (*C.godot_pool_array_write_access)(membe0648b)
	allocsbe0648b := new(cgoAllocMap)
	allocsbe0648b.Add(membe0648b)

	var c_dont_touch_that_allocs *cgoAllocMap
	refbe0648b._dont_touch_that, c_dont_touch_that_allocs = *(*[1]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocsbe0648b.Borrow(c_dont_touch_that_allocs)

	x.refbe0648b = refbe0648b
	x.allocsbe0648b = allocsbe0648b
	return refbe0648b, allocsbe0648b

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotPoolVector3ArrayWriteAccess) PassValue() (C.godot_pool_array_write_access, *cgoAllocMap) {
	if x.refbe0648b != nil {
		return *x.refbe0648b, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotPoolVector3ArrayWriteAccess) Deref() {
	if x.refbe0648b == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.refbe0648b._dont_touch_that))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotPoolColorArrayWriteAccess) Ref() *C.godot_pool_array_write_access {
	if x == nil {
		return nil
	}
	return x.refbe0648b
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotPoolColorArrayWriteAccess) Free() {
	if x != nil && x.allocsbe0648b != nil {
		x.allocsbe0648b.(*cgoAllocMap).Free()
		x.refbe0648b = nil
	}
}

// NewGodotPoolColorArrayWriteAccessRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotPoolColorArrayWriteAccessRef(ref unsafe.Pointer) *GodotPoolColorArrayWriteAccess {
	if ref == nil {
		return nil
	}
	obj := new(GodotPoolColorArrayWriteAccess)
	obj.refbe0648b = (*C.godot_pool_array_write_access)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotPoolColorArrayWriteAccess) PassRef() (*C.godot_pool_array_write_access, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refbe0648b != nil {
		return x.refbe0648b, nil
	}
	membe0648b := allocGodotPoolArrayWriteAccessMemory(1)
	refbe0648b := (*C.godot_pool_array_write_access)(membe0648b)
	allocsbe0648b := new(cgoAllocMap)
	allocsbe0648b.Add(membe0648b)

	var c_dont_touch_that_allocs *cgoAllocMap
	refbe0648b._dont_touch_that, c_dont_touch_that_allocs = *(*[1]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocsbe0648b.Borrow(c_dont_touch_that_allocs)

	x.refbe0648b = refbe0648b
	x.allocsbe0648b = allocsbe0648b
	return refbe0648b, allocsbe0648b

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotPoolColorArrayWriteAccess) PassValue() (C.godot_pool_array_write_access, *cgoAllocMap) {
	if x.refbe0648b != nil {
		return *x.refbe0648b, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotPoolColorArrayWriteAccess) Deref() {
	if x.refbe0648b == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.refbe0648b._dont_touch_that))
}

// allocGodotPoolByteArrayMemory allocates memory for type C.godot_pool_byte_array in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotPoolByteArrayMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotPoolByteArrayValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotPoolByteArrayValue = unsafe.Sizeof([1]C.godot_pool_byte_array{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotPoolByteArray) Ref() *C.godot_pool_byte_array {
	if x == nil {
		return nil
	}
	return x.refbf60ed2
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotPoolByteArray) Free() {
	if x != nil && x.allocsbf60ed2 != nil {
		x.allocsbf60ed2.(*cgoAllocMap).Free()
		x.refbf60ed2 = nil
	}
}

// NewGodotPoolByteArrayRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotPoolByteArrayRef(ref unsafe.Pointer) *GodotPoolByteArray {
	if ref == nil {
		return nil
	}
	obj := new(GodotPoolByteArray)
	obj.refbf60ed2 = (*C.godot_pool_byte_array)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotPoolByteArray) PassRef() (*C.godot_pool_byte_array, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refbf60ed2 != nil {
		return x.refbf60ed2, nil
	}
	membf60ed2 := allocGodotPoolByteArrayMemory(1)
	refbf60ed2 := (*C.godot_pool_byte_array)(membf60ed2)
	allocsbf60ed2 := new(cgoAllocMap)
	allocsbf60ed2.Add(membf60ed2)

	var c_dont_touch_that_allocs *cgoAllocMap
	refbf60ed2._dont_touch_that, c_dont_touch_that_allocs = *(*[8]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocsbf60ed2.Borrow(c_dont_touch_that_allocs)

	x.refbf60ed2 = refbf60ed2
	x.allocsbf60ed2 = allocsbf60ed2
	return refbf60ed2, allocsbf60ed2

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotPoolByteArray) PassValue() (C.godot_pool_byte_array, *cgoAllocMap) {
	if x.refbf60ed2 != nil {
		return *x.refbf60ed2, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotPoolByteArray) Deref() {
	if x.refbf60ed2 == nil {
		return
	}
	x.DontTouchThat = *(*[8]byte)(unsafe.Pointer(&x.refbf60ed2._dont_touch_that))
}

// allocGodotPoolIntArrayMemory allocates memory for type C.godot_pool_int_array in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotPoolIntArrayMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotPoolIntArrayValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotPoolIntArrayValue = unsafe.Sizeof([1]C.godot_pool_int_array{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotPoolIntArray) Ref() *C.godot_pool_int_array {
	if x == nil {
		return nil
	}
	return x.ref5d4f26e6
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotPoolIntArray) Free() {
	if x != nil && x.allocs5d4f26e6 != nil {
		x.allocs5d4f26e6.(*cgoAllocMap).Free()
		x.ref5d4f26e6 = nil
	}
}

// NewGodotPoolIntArrayRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotPoolIntArrayRef(ref unsafe.Pointer) *GodotPoolIntArray {
	if ref == nil {
		return nil
	}
	obj := new(GodotPoolIntArray)
	obj.ref5d4f26e6 = (*C.godot_pool_int_array)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotPoolIntArray) PassRef() (*C.godot_pool_int_array, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref5d4f26e6 != nil {
		return x.ref5d4f26e6, nil
	}
	mem5d4f26e6 := allocGodotPoolIntArrayMemory(1)
	ref5d4f26e6 := (*C.godot_pool_int_array)(mem5d4f26e6)
	allocs5d4f26e6 := new(cgoAllocMap)
	allocs5d4f26e6.Add(mem5d4f26e6)

	var c_dont_touch_that_allocs *cgoAllocMap
	ref5d4f26e6._dont_touch_that, c_dont_touch_that_allocs = *(*[8]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocs5d4f26e6.Borrow(c_dont_touch_that_allocs)

	x.ref5d4f26e6 = ref5d4f26e6
	x.allocs5d4f26e6 = allocs5d4f26e6
	return ref5d4f26e6, allocs5d4f26e6

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotPoolIntArray) PassValue() (C.godot_pool_int_array, *cgoAllocMap) {
	if x.ref5d4f26e6 != nil {
		return *x.ref5d4f26e6, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotPoolIntArray) Deref() {
	if x.ref5d4f26e6 == nil {
		return
	}
	x.DontTouchThat = *(*[8]byte)(unsafe.Pointer(&x.ref5d4f26e6._dont_touch_that))
}

// allocGodotPoolRealArrayMemory allocates memory for type C.godot_pool_real_array in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotPoolRealArrayMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotPoolRealArrayValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotPoolRealArrayValue = unsafe.Sizeof([1]C.godot_pool_real_array{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotPoolRealArray) Ref() *C.godot_pool_real_array {
	if x == nil {
		return nil
	}
	return x.refc76f44c3
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotPoolRealArray) Free() {
	if x != nil && x.allocsc76f44c3 != nil {
		x.allocsc76f44c3.(*cgoAllocMap).Free()
		x.refc76f44c3 = nil
	}
}

// NewGodotPoolRealArrayRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotPoolRealArrayRef(ref unsafe.Pointer) *GodotPoolRealArray {
	if ref == nil {
		return nil
	}
	obj := new(GodotPoolRealArray)
	obj.refc76f44c3 = (*C.godot_pool_real_array)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotPoolRealArray) PassRef() (*C.godot_pool_real_array, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refc76f44c3 != nil {
		return x.refc76f44c3, nil
	}
	memc76f44c3 := allocGodotPoolRealArrayMemory(1)
	refc76f44c3 := (*C.godot_pool_real_array)(memc76f44c3)
	allocsc76f44c3 := new(cgoAllocMap)
	allocsc76f44c3.Add(memc76f44c3)

	var c_dont_touch_that_allocs *cgoAllocMap
	refc76f44c3._dont_touch_that, c_dont_touch_that_allocs = *(*[8]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocsc76f44c3.Borrow(c_dont_touch_that_allocs)

	x.refc76f44c3 = refc76f44c3
	x.allocsc76f44c3 = allocsc76f44c3
	return refc76f44c3, allocsc76f44c3

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotPoolRealArray) PassValue() (C.godot_pool_real_array, *cgoAllocMap) {
	if x.refc76f44c3 != nil {
		return *x.refc76f44c3, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotPoolRealArray) Deref() {
	if x.refc76f44c3 == nil {
		return
	}
	x.DontTouchThat = *(*[8]byte)(unsafe.Pointer(&x.refc76f44c3._dont_touch_that))
}

// allocGodotPoolStringArrayMemory allocates memory for type C.godot_pool_string_array in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotPoolStringArrayMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotPoolStringArrayValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotPoolStringArrayValue = unsafe.Sizeof([1]C.godot_pool_string_array{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotPoolStringArray) Ref() *C.godot_pool_string_array {
	if x == nil {
		return nil
	}
	return x.reff6fe5d9
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotPoolStringArray) Free() {
	if x != nil && x.allocsf6fe5d9 != nil {
		x.allocsf6fe5d9.(*cgoAllocMap).Free()
		x.reff6fe5d9 = nil
	}
}

// NewGodotPoolStringArrayRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotPoolStringArrayRef(ref unsafe.Pointer) *GodotPoolStringArray {
	if ref == nil {
		return nil
	}
	obj := new(GodotPoolStringArray)
	obj.reff6fe5d9 = (*C.godot_pool_string_array)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotPoolStringArray) PassRef() (*C.godot_pool_string_array, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.reff6fe5d9 != nil {
		return x.reff6fe5d9, nil
	}
	memf6fe5d9 := allocGodotPoolStringArrayMemory(1)
	reff6fe5d9 := (*C.godot_pool_string_array)(memf6fe5d9)
	allocsf6fe5d9 := new(cgoAllocMap)
	allocsf6fe5d9.Add(memf6fe5d9)

	var c_dont_touch_that_allocs *cgoAllocMap
	reff6fe5d9._dont_touch_that, c_dont_touch_that_allocs = *(*[8]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocsf6fe5d9.Borrow(c_dont_touch_that_allocs)

	x.reff6fe5d9 = reff6fe5d9
	x.allocsf6fe5d9 = allocsf6fe5d9
	return reff6fe5d9, allocsf6fe5d9

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotPoolStringArray) PassValue() (C.godot_pool_string_array, *cgoAllocMap) {
	if x.reff6fe5d9 != nil {
		return *x.reff6fe5d9, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotPoolStringArray) Deref() {
	if x.reff6fe5d9 == nil {
		return
	}
	x.DontTouchThat = *(*[8]byte)(unsafe.Pointer(&x.reff6fe5d9._dont_touch_that))
}

// allocGodotPoolVector2ArrayMemory allocates memory for type C.godot_pool_vector2_array in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotPoolVector2ArrayMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotPoolVector2ArrayValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotPoolVector2ArrayValue = unsafe.Sizeof([1]C.godot_pool_vector2_array{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotPoolVector2Array) Ref() *C.godot_pool_vector2_array {
	if x == nil {
		return nil
	}
	return x.ref7f6b2885
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotPoolVector2Array) Free() {
	if x != nil && x.allocs7f6b2885 != nil {
		x.allocs7f6b2885.(*cgoAllocMap).Free()
		x.ref7f6b2885 = nil
	}
}

// NewGodotPoolVector2ArrayRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotPoolVector2ArrayRef(ref unsafe.Pointer) *GodotPoolVector2Array {
	if ref == nil {
		return nil
	}
	obj := new(GodotPoolVector2Array)
	obj.ref7f6b2885 = (*C.godot_pool_vector2_array)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotPoolVector2Array) PassRef() (*C.godot_pool_vector2_array, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref7f6b2885 != nil {
		return x.ref7f6b2885, nil
	}
	mem7f6b2885 := allocGodotPoolVector2ArrayMemory(1)
	ref7f6b2885 := (*C.godot_pool_vector2_array)(mem7f6b2885)
	allocs7f6b2885 := new(cgoAllocMap)
	allocs7f6b2885.Add(mem7f6b2885)

	var c_dont_touch_that_allocs *cgoAllocMap
	ref7f6b2885._dont_touch_that, c_dont_touch_that_allocs = *(*[8]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocs7f6b2885.Borrow(c_dont_touch_that_allocs)

	x.ref7f6b2885 = ref7f6b2885
	x.allocs7f6b2885 = allocs7f6b2885
	return ref7f6b2885, allocs7f6b2885

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotPoolVector2Array) PassValue() (C.godot_pool_vector2_array, *cgoAllocMap) {
	if x.ref7f6b2885 != nil {
		return *x.ref7f6b2885, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotPoolVector2Array) Deref() {
	if x.ref7f6b2885 == nil {
		return
	}
	x.DontTouchThat = *(*[8]byte)(unsafe.Pointer(&x.ref7f6b2885._dont_touch_that))
}

// allocGodotPoolVector3ArrayMemory allocates memory for type C.godot_pool_vector3_array in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotPoolVector3ArrayMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotPoolVector3ArrayValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotPoolVector3ArrayValue = unsafe.Sizeof([1]C.godot_pool_vector3_array{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotPoolVector3Array) Ref() *C.godot_pool_vector3_array {
	if x == nil {
		return nil
	}
	return x.refd91c2331
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotPoolVector3Array) Free() {
	if x != nil && x.allocsd91c2331 != nil {
		x.allocsd91c2331.(*cgoAllocMap).Free()
		x.refd91c2331 = nil
	}
}

// NewGodotPoolVector3ArrayRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotPoolVector3ArrayRef(ref unsafe.Pointer) *GodotPoolVector3Array {
	if ref == nil {
		return nil
	}
	obj := new(GodotPoolVector3Array)
	obj.refd91c2331 = (*C.godot_pool_vector3_array)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotPoolVector3Array) PassRef() (*C.godot_pool_vector3_array, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refd91c2331 != nil {
		return x.refd91c2331, nil
	}
	memd91c2331 := allocGodotPoolVector3ArrayMemory(1)
	refd91c2331 := (*C.godot_pool_vector3_array)(memd91c2331)
	allocsd91c2331 := new(cgoAllocMap)
	allocsd91c2331.Add(memd91c2331)

	var c_dont_touch_that_allocs *cgoAllocMap
	refd91c2331._dont_touch_that, c_dont_touch_that_allocs = *(*[8]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocsd91c2331.Borrow(c_dont_touch_that_allocs)

	x.refd91c2331 = refd91c2331
	x.allocsd91c2331 = allocsd91c2331
	return refd91c2331, allocsd91c2331

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotPoolVector3Array) PassValue() (C.godot_pool_vector3_array, *cgoAllocMap) {
	if x.refd91c2331 != nil {
		return *x.refd91c2331, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotPoolVector3Array) Deref() {
	if x.refd91c2331 == nil {
		return
	}
	x.DontTouchThat = *(*[8]byte)(unsafe.Pointer(&x.refd91c2331._dont_touch_that))
}

// allocGodotPoolColorArrayMemory allocates memory for type C.godot_pool_color_array in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotPoolColorArrayMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotPoolColorArrayValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotPoolColorArrayValue = unsafe.Sizeof([1]C.godot_pool_color_array{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotPoolColorArray) Ref() *C.godot_pool_color_array {
	if x == nil {
		return nil
	}
	return x.refd5cae78e
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotPoolColorArray) Free() {
	if x != nil && x.allocsd5cae78e != nil {
		x.allocsd5cae78e.(*cgoAllocMap).Free()
		x.refd5cae78e = nil
	}
}

// NewGodotPoolColorArrayRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotPoolColorArrayRef(ref unsafe.Pointer) *GodotPoolColorArray {
	if ref == nil {
		return nil
	}
	obj := new(GodotPoolColorArray)
	obj.refd5cae78e = (*C.godot_pool_color_array)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotPoolColorArray) PassRef() (*C.godot_pool_color_array, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refd5cae78e != nil {
		return x.refd5cae78e, nil
	}
	memd5cae78e := allocGodotPoolColorArrayMemory(1)
	refd5cae78e := (*C.godot_pool_color_array)(memd5cae78e)
	allocsd5cae78e := new(cgoAllocMap)
	allocsd5cae78e.Add(memd5cae78e)

	var c_dont_touch_that_allocs *cgoAllocMap
	refd5cae78e._dont_touch_that, c_dont_touch_that_allocs = *(*[8]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocsd5cae78e.Borrow(c_dont_touch_that_allocs)

	x.refd5cae78e = refd5cae78e
	x.allocsd5cae78e = allocsd5cae78e
	return refd5cae78e, allocsd5cae78e

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotPoolColorArray) PassValue() (C.godot_pool_color_array, *cgoAllocMap) {
	if x.refd5cae78e != nil {
		return *x.refd5cae78e, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotPoolColorArray) Deref() {
	if x.refd5cae78e == nil {
		return
	}
	x.DontTouchThat = *(*[8]byte)(unsafe.Pointer(&x.refd5cae78e._dont_touch_that))
}

// allocGodotColorMemory allocates memory for type C.godot_color in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotColorMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotColorValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotColorValue = unsafe.Sizeof([1]C.godot_color{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotColor) Ref() *C.godot_color {
	if x == nil {
		return nil
	}
	return x.ref7f4bfefb
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotColor) Free() {
	if x != nil && x.allocs7f4bfefb != nil {
		x.allocs7f4bfefb.(*cgoAllocMap).Free()
		x.ref7f4bfefb = nil
	}
}

// NewGodotColorRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotColorRef(ref unsafe.Pointer) *GodotColor {
	if ref == nil {
		return nil
	}
	obj := new(GodotColor)
	obj.ref7f4bfefb = (*C.godot_color)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotColor) PassRef() (*C.godot_color, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref7f4bfefb != nil {
		return x.ref7f4bfefb, nil
	}
	mem7f4bfefb := allocGodotColorMemory(1)
	ref7f4bfefb := (*C.godot_color)(mem7f4bfefb)
	allocs7f4bfefb := new(cgoAllocMap)
	allocs7f4bfefb.Add(mem7f4bfefb)

	var c_dont_touch_that_allocs *cgoAllocMap
	ref7f4bfefb._dont_touch_that, c_dont_touch_that_allocs = *(*[16]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocs7f4bfefb.Borrow(c_dont_touch_that_allocs)

	x.ref7f4bfefb = ref7f4bfefb
	x.allocs7f4bfefb = allocs7f4bfefb
	return ref7f4bfefb, allocs7f4bfefb

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotColor) PassValue() (C.godot_color, *cgoAllocMap) {
	if x.ref7f4bfefb != nil {
		return *x.ref7f4bfefb, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotColor) Deref() {
	if x.ref7f4bfefb == nil {
		return
	}
	x.DontTouchThat = *(*[16]byte)(unsafe.Pointer(&x.ref7f4bfefb._dont_touch_that))
}

// allocGodotVector2Memory allocates memory for type C.godot_vector2 in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotVector2Memory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotVector2Value))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotVector2Value = unsafe.Sizeof([1]C.godot_vector2{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotVector2) Ref() *C.godot_vector2 {
	if x == nil {
		return nil
	}
	return x.refbc81274e
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotVector2) Free() {
	if x != nil && x.allocsbc81274e != nil {
		x.allocsbc81274e.(*cgoAllocMap).Free()
		x.refbc81274e = nil
	}
}

// NewGodotVector2Ref creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotVector2Ref(ref unsafe.Pointer) *GodotVector2 {
	if ref == nil {
		return nil
	}
	obj := new(GodotVector2)
	obj.refbc81274e = (*C.godot_vector2)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotVector2) PassRef() (*C.godot_vector2, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refbc81274e != nil {
		return x.refbc81274e, nil
	}
	membc81274e := allocGodotVector2Memory(1)
	refbc81274e := (*C.godot_vector2)(membc81274e)
	allocsbc81274e := new(cgoAllocMap)
	allocsbc81274e.Add(membc81274e)

	var c_dont_touch_that_allocs *cgoAllocMap
	refbc81274e._dont_touch_that, c_dont_touch_that_allocs = *(*[8]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocsbc81274e.Borrow(c_dont_touch_that_allocs)

	x.refbc81274e = refbc81274e
	x.allocsbc81274e = allocsbc81274e
	return refbc81274e, allocsbc81274e

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotVector2) PassValue() (C.godot_vector2, *cgoAllocMap) {
	if x.refbc81274e != nil {
		return *x.refbc81274e, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotVector2) Deref() {
	if x.refbc81274e == nil {
		return
	}
	x.DontTouchThat = *(*[8]byte)(unsafe.Pointer(&x.refbc81274e._dont_touch_that))
}

// allocGodotVector3Memory allocates memory for type C.godot_vector3 in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotVector3Memory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotVector3Value))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotVector3Value = unsafe.Sizeof([1]C.godot_vector3{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotVector3) Ref() *C.godot_vector3 {
	if x == nil {
		return nil
	}
	return x.refcb8617d8
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotVector3) Free() {
	if x != nil && x.allocscb8617d8 != nil {
		x.allocscb8617d8.(*cgoAllocMap).Free()
		x.refcb8617d8 = nil
	}
}

// NewGodotVector3Ref creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotVector3Ref(ref unsafe.Pointer) *GodotVector3 {
	if ref == nil {
		return nil
	}
	obj := new(GodotVector3)
	obj.refcb8617d8 = (*C.godot_vector3)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotVector3) PassRef() (*C.godot_vector3, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refcb8617d8 != nil {
		return x.refcb8617d8, nil
	}
	memcb8617d8 := allocGodotVector3Memory(1)
	refcb8617d8 := (*C.godot_vector3)(memcb8617d8)
	allocscb8617d8 := new(cgoAllocMap)
	allocscb8617d8.Add(memcb8617d8)

	var c_dont_touch_that_allocs *cgoAllocMap
	refcb8617d8._dont_touch_that, c_dont_touch_that_allocs = *(*[12]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocscb8617d8.Borrow(c_dont_touch_that_allocs)

	x.refcb8617d8 = refcb8617d8
	x.allocscb8617d8 = allocscb8617d8
	return refcb8617d8, allocscb8617d8

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotVector3) PassValue() (C.godot_vector3, *cgoAllocMap) {
	if x.refcb8617d8 != nil {
		return *x.refcb8617d8, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotVector3) Deref() {
	if x.refcb8617d8 == nil {
		return
	}
	x.DontTouchThat = *(*[12]byte)(unsafe.Pointer(&x.refcb8617d8._dont_touch_that))
}

// allocGodotBasisMemory allocates memory for type C.godot_basis in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotBasisMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotBasisValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotBasisValue = unsafe.Sizeof([1]C.godot_basis{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotBasis) Ref() *C.godot_basis {
	if x == nil {
		return nil
	}
	return x.ref94d3d325
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotBasis) Free() {
	if x != nil && x.allocs94d3d325 != nil {
		x.allocs94d3d325.(*cgoAllocMap).Free()
		x.ref94d3d325 = nil
	}
}

// NewGodotBasisRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotBasisRef(ref unsafe.Pointer) *GodotBasis {
	if ref == nil {
		return nil
	}
	obj := new(GodotBasis)
	obj.ref94d3d325 = (*C.godot_basis)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotBasis) PassRef() (*C.godot_basis, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref94d3d325 != nil {
		return x.ref94d3d325, nil
	}
	mem94d3d325 := allocGodotBasisMemory(1)
	ref94d3d325 := (*C.godot_basis)(mem94d3d325)
	allocs94d3d325 := new(cgoAllocMap)
	allocs94d3d325.Add(mem94d3d325)

	var c_dont_touch_that_allocs *cgoAllocMap
	ref94d3d325._dont_touch_that, c_dont_touch_that_allocs = *(*[36]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocs94d3d325.Borrow(c_dont_touch_that_allocs)

	x.ref94d3d325 = ref94d3d325
	x.allocs94d3d325 = allocs94d3d325
	return ref94d3d325, allocs94d3d325

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotBasis) PassValue() (C.godot_basis, *cgoAllocMap) {
	if x.ref94d3d325 != nil {
		return *x.ref94d3d325, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotBasis) Deref() {
	if x.ref94d3d325 == nil {
		return
	}
	x.DontTouchThat = *(*[36]byte)(unsafe.Pointer(&x.ref94d3d325._dont_touch_that))
}

// allocGodotQuatMemory allocates memory for type C.godot_quat in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotQuatMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotQuatValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotQuatValue = unsafe.Sizeof([1]C.godot_quat{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotQuat) Ref() *C.godot_quat {
	if x == nil {
		return nil
	}
	return x.reffaf33e0b
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotQuat) Free() {
	if x != nil && x.allocsfaf33e0b != nil {
		x.allocsfaf33e0b.(*cgoAllocMap).Free()
		x.reffaf33e0b = nil
	}
}

// NewGodotQuatRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotQuatRef(ref unsafe.Pointer) *GodotQuat {
	if ref == nil {
		return nil
	}
	obj := new(GodotQuat)
	obj.reffaf33e0b = (*C.godot_quat)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotQuat) PassRef() (*C.godot_quat, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.reffaf33e0b != nil {
		return x.reffaf33e0b, nil
	}
	memfaf33e0b := allocGodotQuatMemory(1)
	reffaf33e0b := (*C.godot_quat)(memfaf33e0b)
	allocsfaf33e0b := new(cgoAllocMap)
	allocsfaf33e0b.Add(memfaf33e0b)

	var c_dont_touch_that_allocs *cgoAllocMap
	reffaf33e0b._dont_touch_that, c_dont_touch_that_allocs = *(*[16]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocsfaf33e0b.Borrow(c_dont_touch_that_allocs)

	x.reffaf33e0b = reffaf33e0b
	x.allocsfaf33e0b = allocsfaf33e0b
	return reffaf33e0b, allocsfaf33e0b

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotQuat) PassValue() (C.godot_quat, *cgoAllocMap) {
	if x.reffaf33e0b != nil {
		return *x.reffaf33e0b, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotQuat) Deref() {
	if x.reffaf33e0b == nil {
		return
	}
	x.DontTouchThat = *(*[16]byte)(unsafe.Pointer(&x.reffaf33e0b._dont_touch_that))
}

// allocGodotVariantMemory allocates memory for type C.godot_variant in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotVariantMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotVariantValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotVariantValue = unsafe.Sizeof([1]C.godot_variant{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotVariant) Ref() *C.godot_variant {
	if x == nil {
		return nil
	}
	return x.refabb5c0da
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotVariant) Free() {
	if x != nil && x.allocsabb5c0da != nil {
		x.allocsabb5c0da.(*cgoAllocMap).Free()
		x.refabb5c0da = nil
	}
}

// NewGodotVariantRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotVariantRef(ref unsafe.Pointer) *GodotVariant {
	if ref == nil {
		return nil
	}
	obj := new(GodotVariant)
	obj.refabb5c0da = (*C.godot_variant)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotVariant) PassRef() (*C.godot_variant, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refabb5c0da != nil {
		return x.refabb5c0da, nil
	}
	memabb5c0da := allocGodotVariantMemory(1)
	refabb5c0da := (*C.godot_variant)(memabb5c0da)
	allocsabb5c0da := new(cgoAllocMap)
	allocsabb5c0da.Add(memabb5c0da)

	var c_dont_touch_that_allocs *cgoAllocMap
	refabb5c0da._dont_touch_that, c_dont_touch_that_allocs = *(*[24]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocsabb5c0da.Borrow(c_dont_touch_that_allocs)

	x.refabb5c0da = refabb5c0da
	x.allocsabb5c0da = allocsabb5c0da
	return refabb5c0da, allocsabb5c0da

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotVariant) PassValue() (C.godot_variant, *cgoAllocMap) {
	if x.refabb5c0da != nil {
		return *x.refabb5c0da, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotVariant) Deref() {
	if x.refabb5c0da == nil {
		return
	}
	x.DontTouchThat = *(*[24]byte)(unsafe.Pointer(&x.refabb5c0da._dont_touch_that))
}

// allocGodotVariantCallErrorMemory allocates memory for type C.godot_variant_call_error in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotVariantCallErrorMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotVariantCallErrorValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotVariantCallErrorValue = unsafe.Sizeof([1]C.godot_variant_call_error{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotVariantCallError) Ref() *C.godot_variant_call_error {
	if x == nil {
		return nil
	}
	return x.ref3ce71027
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotVariantCallError) Free() {
	if x != nil && x.allocs3ce71027 != nil {
		x.allocs3ce71027.(*cgoAllocMap).Free()
		x.ref3ce71027 = nil
	}
}

// NewGodotVariantCallErrorRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotVariantCallErrorRef(ref unsafe.Pointer) *GodotVariantCallError {
	if ref == nil {
		return nil
	}
	obj := new(GodotVariantCallError)
	obj.ref3ce71027 = (*C.godot_variant_call_error)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotVariantCallError) PassRef() (*C.godot_variant_call_error, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref3ce71027 != nil {
		return x.ref3ce71027, nil
	}
	mem3ce71027 := allocGodotVariantCallErrorMemory(1)
	ref3ce71027 := (*C.godot_variant_call_error)(mem3ce71027)
	allocs3ce71027 := new(cgoAllocMap)
	allocs3ce71027.Add(mem3ce71027)

	var cerror_allocs *cgoAllocMap
	ref3ce71027.error, cerror_allocs = (C.godot_variant_call_error_error)(x.Error), cgoAllocsUnknown
	allocs3ce71027.Borrow(cerror_allocs)

	var cargument_allocs *cgoAllocMap
	ref3ce71027.argument, cargument_allocs = (C.int)(x.Argument), cgoAllocsUnknown
	allocs3ce71027.Borrow(cargument_allocs)

	var cexpected_allocs *cgoAllocMap
	ref3ce71027.expected, cexpected_allocs = (C.godot_variant_type)(x.Expected), cgoAllocsUnknown
	allocs3ce71027.Borrow(cexpected_allocs)

	x.ref3ce71027 = ref3ce71027
	x.allocs3ce71027 = allocs3ce71027
	return ref3ce71027, allocs3ce71027

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotVariantCallError) PassValue() (C.godot_variant_call_error, *cgoAllocMap) {
	if x.ref3ce71027 != nil {
		return *x.ref3ce71027, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotVariantCallError) Deref() {
	if x.ref3ce71027 == nil {
		return
	}
	x.Error = (GodotVariantCallErrorError)(x.ref3ce71027.error)
	x.Argument = (int32)(x.ref3ce71027.argument)
	x.Expected = (GodotVariantType)(x.ref3ce71027.expected)
}

// allocGodotAabbMemory allocates memory for type C.godot_aabb in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotAabbMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotAabbValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotAabbValue = unsafe.Sizeof([1]C.godot_aabb{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotAabb) Ref() *C.godot_aabb {
	if x == nil {
		return nil
	}
	return x.ref6e3c84aa
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotAabb) Free() {
	if x != nil && x.allocs6e3c84aa != nil {
		x.allocs6e3c84aa.(*cgoAllocMap).Free()
		x.ref6e3c84aa = nil
	}
}

// NewGodotAabbRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotAabbRef(ref unsafe.Pointer) *GodotAabb {
	if ref == nil {
		return nil
	}
	obj := new(GodotAabb)
	obj.ref6e3c84aa = (*C.godot_aabb)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotAabb) PassRef() (*C.godot_aabb, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref6e3c84aa != nil {
		return x.ref6e3c84aa, nil
	}
	mem6e3c84aa := allocGodotAabbMemory(1)
	ref6e3c84aa := (*C.godot_aabb)(mem6e3c84aa)
	allocs6e3c84aa := new(cgoAllocMap)
	allocs6e3c84aa.Add(mem6e3c84aa)

	var c_dont_touch_that_allocs *cgoAllocMap
	ref6e3c84aa._dont_touch_that, c_dont_touch_that_allocs = *(*[24]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocs6e3c84aa.Borrow(c_dont_touch_that_allocs)

	x.ref6e3c84aa = ref6e3c84aa
	x.allocs6e3c84aa = allocs6e3c84aa
	return ref6e3c84aa, allocs6e3c84aa

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotAabb) PassValue() (C.godot_aabb, *cgoAllocMap) {
	if x.ref6e3c84aa != nil {
		return *x.ref6e3c84aa, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotAabb) Deref() {
	if x.ref6e3c84aa == nil {
		return
	}
	x.DontTouchThat = *(*[24]byte)(unsafe.Pointer(&x.ref6e3c84aa._dont_touch_that))
}

// allocGodotPlaneMemory allocates memory for type C.godot_plane in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotPlaneMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotPlaneValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotPlaneValue = unsafe.Sizeof([1]C.godot_plane{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotPlane) Ref() *C.godot_plane {
	if x == nil {
		return nil
	}
	return x.refd8ae9b92
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotPlane) Free() {
	if x != nil && x.allocsd8ae9b92 != nil {
		x.allocsd8ae9b92.(*cgoAllocMap).Free()
		x.refd8ae9b92 = nil
	}
}

// NewGodotPlaneRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotPlaneRef(ref unsafe.Pointer) *GodotPlane {
	if ref == nil {
		return nil
	}
	obj := new(GodotPlane)
	obj.refd8ae9b92 = (*C.godot_plane)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotPlane) PassRef() (*C.godot_plane, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refd8ae9b92 != nil {
		return x.refd8ae9b92, nil
	}
	memd8ae9b92 := allocGodotPlaneMemory(1)
	refd8ae9b92 := (*C.godot_plane)(memd8ae9b92)
	allocsd8ae9b92 := new(cgoAllocMap)
	allocsd8ae9b92.Add(memd8ae9b92)

	var c_dont_touch_that_allocs *cgoAllocMap
	refd8ae9b92._dont_touch_that, c_dont_touch_that_allocs = *(*[16]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocsd8ae9b92.Borrow(c_dont_touch_that_allocs)

	x.refd8ae9b92 = refd8ae9b92
	x.allocsd8ae9b92 = allocsd8ae9b92
	return refd8ae9b92, allocsd8ae9b92

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotPlane) PassValue() (C.godot_plane, *cgoAllocMap) {
	if x.refd8ae9b92 != nil {
		return *x.refd8ae9b92, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotPlane) Deref() {
	if x.refd8ae9b92 == nil {
		return
	}
	x.DontTouchThat = *(*[16]byte)(unsafe.Pointer(&x.refd8ae9b92._dont_touch_that))
}

// allocGodotDictionaryMemory allocates memory for type C.godot_dictionary in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotDictionaryMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotDictionaryValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotDictionaryValue = unsafe.Sizeof([1]C.godot_dictionary{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotDictionary) Ref() *C.godot_dictionary {
	if x == nil {
		return nil
	}
	return x.refb267a9b9
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotDictionary) Free() {
	if x != nil && x.allocsb267a9b9 != nil {
		x.allocsb267a9b9.(*cgoAllocMap).Free()
		x.refb267a9b9 = nil
	}
}

// NewGodotDictionaryRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotDictionaryRef(ref unsafe.Pointer) *GodotDictionary {
	if ref == nil {
		return nil
	}
	obj := new(GodotDictionary)
	obj.refb267a9b9 = (*C.godot_dictionary)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotDictionary) PassRef() (*C.godot_dictionary, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refb267a9b9 != nil {
		return x.refb267a9b9, nil
	}
	memb267a9b9 := allocGodotDictionaryMemory(1)
	refb267a9b9 := (*C.godot_dictionary)(memb267a9b9)
	allocsb267a9b9 := new(cgoAllocMap)
	allocsb267a9b9.Add(memb267a9b9)

	var c_dont_touch_that_allocs *cgoAllocMap
	refb267a9b9._dont_touch_that, c_dont_touch_that_allocs = *(*[8]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocsb267a9b9.Borrow(c_dont_touch_that_allocs)

	x.refb267a9b9 = refb267a9b9
	x.allocsb267a9b9 = allocsb267a9b9
	return refb267a9b9, allocsb267a9b9

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotDictionary) PassValue() (C.godot_dictionary, *cgoAllocMap) {
	if x.refb267a9b9 != nil {
		return *x.refb267a9b9, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotDictionary) Deref() {
	if x.refb267a9b9 == nil {
		return
	}
	x.DontTouchThat = *(*[8]byte)(unsafe.Pointer(&x.refb267a9b9._dont_touch_that))
}

// allocGodotNodePathMemory allocates memory for type C.godot_node_path in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotNodePathMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotNodePathValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotNodePathValue = unsafe.Sizeof([1]C.godot_node_path{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotNodePath) Ref() *C.godot_node_path {
	if x == nil {
		return nil
	}
	return x.ref6c34dff3
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotNodePath) Free() {
	if x != nil && x.allocs6c34dff3 != nil {
		x.allocs6c34dff3.(*cgoAllocMap).Free()
		x.ref6c34dff3 = nil
	}
}

// NewGodotNodePathRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotNodePathRef(ref unsafe.Pointer) *GodotNodePath {
	if ref == nil {
		return nil
	}
	obj := new(GodotNodePath)
	obj.ref6c34dff3 = (*C.godot_node_path)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotNodePath) PassRef() (*C.godot_node_path, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref6c34dff3 != nil {
		return x.ref6c34dff3, nil
	}
	mem6c34dff3 := allocGodotNodePathMemory(1)
	ref6c34dff3 := (*C.godot_node_path)(mem6c34dff3)
	allocs6c34dff3 := new(cgoAllocMap)
	allocs6c34dff3.Add(mem6c34dff3)

	var c_dont_touch_that_allocs *cgoAllocMap
	ref6c34dff3._dont_touch_that, c_dont_touch_that_allocs = *(*[8]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocs6c34dff3.Borrow(c_dont_touch_that_allocs)

	x.ref6c34dff3 = ref6c34dff3
	x.allocs6c34dff3 = allocs6c34dff3
	return ref6c34dff3, allocs6c34dff3

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotNodePath) PassValue() (C.godot_node_path, *cgoAllocMap) {
	if x.ref6c34dff3 != nil {
		return *x.ref6c34dff3, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotNodePath) Deref() {
	if x.ref6c34dff3 == nil {
		return
	}
	x.DontTouchThat = *(*[8]byte)(unsafe.Pointer(&x.ref6c34dff3._dont_touch_that))
}

// allocGodotRect2Memory allocates memory for type C.godot_rect2 in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotRect2Memory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotRect2Value))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotRect2Value = unsafe.Sizeof([1]C.godot_rect2{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotRect2) Ref() *C.godot_rect2 {
	if x == nil {
		return nil
	}
	return x.ref99c06d9a
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotRect2) Free() {
	if x != nil && x.allocs99c06d9a != nil {
		x.allocs99c06d9a.(*cgoAllocMap).Free()
		x.ref99c06d9a = nil
	}
}

// NewGodotRect2Ref creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotRect2Ref(ref unsafe.Pointer) *GodotRect2 {
	if ref == nil {
		return nil
	}
	obj := new(GodotRect2)
	obj.ref99c06d9a = (*C.godot_rect2)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotRect2) PassRef() (*C.godot_rect2, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref99c06d9a != nil {
		return x.ref99c06d9a, nil
	}
	mem99c06d9a := allocGodotRect2Memory(1)
	ref99c06d9a := (*C.godot_rect2)(mem99c06d9a)
	allocs99c06d9a := new(cgoAllocMap)
	allocs99c06d9a.Add(mem99c06d9a)

	var c_dont_touch_that_allocs *cgoAllocMap
	ref99c06d9a._dont_touch_that, c_dont_touch_that_allocs = *(*[16]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocs99c06d9a.Borrow(c_dont_touch_that_allocs)

	x.ref99c06d9a = ref99c06d9a
	x.allocs99c06d9a = allocs99c06d9a
	return ref99c06d9a, allocs99c06d9a

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotRect2) PassValue() (C.godot_rect2, *cgoAllocMap) {
	if x.ref99c06d9a != nil {
		return *x.ref99c06d9a, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotRect2) Deref() {
	if x.ref99c06d9a == nil {
		return
	}
	x.DontTouchThat = *(*[16]byte)(unsafe.Pointer(&x.ref99c06d9a._dont_touch_that))
}

// allocGodotRidMemory allocates memory for type C.godot_rid in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotRidMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotRidValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotRidValue = unsafe.Sizeof([1]C.godot_rid{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotRid) Ref() *C.godot_rid {
	if x == nil {
		return nil
	}
	return x.ref67320fc7
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotRid) Free() {
	if x != nil && x.allocs67320fc7 != nil {
		x.allocs67320fc7.(*cgoAllocMap).Free()
		x.ref67320fc7 = nil
	}
}

// NewGodotRidRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotRidRef(ref unsafe.Pointer) *GodotRid {
	if ref == nil {
		return nil
	}
	obj := new(GodotRid)
	obj.ref67320fc7 = (*C.godot_rid)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotRid) PassRef() (*C.godot_rid, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref67320fc7 != nil {
		return x.ref67320fc7, nil
	}
	mem67320fc7 := allocGodotRidMemory(1)
	ref67320fc7 := (*C.godot_rid)(mem67320fc7)
	allocs67320fc7 := new(cgoAllocMap)
	allocs67320fc7.Add(mem67320fc7)

	var c_dont_touch_that_allocs *cgoAllocMap
	ref67320fc7._dont_touch_that, c_dont_touch_that_allocs = *(*[8]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocs67320fc7.Borrow(c_dont_touch_that_allocs)

	x.ref67320fc7 = ref67320fc7
	x.allocs67320fc7 = allocs67320fc7
	return ref67320fc7, allocs67320fc7

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotRid) PassValue() (C.godot_rid, *cgoAllocMap) {
	if x.ref67320fc7 != nil {
		return *x.ref67320fc7, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotRid) Deref() {
	if x.ref67320fc7 == nil {
		return
	}
	x.DontTouchThat = *(*[8]byte)(unsafe.Pointer(&x.ref67320fc7._dont_touch_that))
}

// allocGodotTransformMemory allocates memory for type C.godot_transform in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotTransformMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotTransformValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotTransformValue = unsafe.Sizeof([1]C.godot_transform{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotTransform) Ref() *C.godot_transform {
	if x == nil {
		return nil
	}
	return x.refd77658c7
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotTransform) Free() {
	if x != nil && x.allocsd77658c7 != nil {
		x.allocsd77658c7.(*cgoAllocMap).Free()
		x.refd77658c7 = nil
	}
}

// NewGodotTransformRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotTransformRef(ref unsafe.Pointer) *GodotTransform {
	if ref == nil {
		return nil
	}
	obj := new(GodotTransform)
	obj.refd77658c7 = (*C.godot_transform)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotTransform) PassRef() (*C.godot_transform, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refd77658c7 != nil {
		return x.refd77658c7, nil
	}
	memd77658c7 := allocGodotTransformMemory(1)
	refd77658c7 := (*C.godot_transform)(memd77658c7)
	allocsd77658c7 := new(cgoAllocMap)
	allocsd77658c7.Add(memd77658c7)

	var c_dont_touch_that_allocs *cgoAllocMap
	refd77658c7._dont_touch_that, c_dont_touch_that_allocs = *(*[48]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocsd77658c7.Borrow(c_dont_touch_that_allocs)

	x.refd77658c7 = refd77658c7
	x.allocsd77658c7 = allocsd77658c7
	return refd77658c7, allocsd77658c7

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotTransform) PassValue() (C.godot_transform, *cgoAllocMap) {
	if x.refd77658c7 != nil {
		return *x.refd77658c7, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotTransform) Deref() {
	if x.refd77658c7 == nil {
		return
	}
	x.DontTouchThat = *(*[48]byte)(unsafe.Pointer(&x.refd77658c7._dont_touch_that))
}

// allocGodotTransform2dMemory allocates memory for type C.godot_transform2d in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotTransform2dMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotTransform2dValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotTransform2dValue = unsafe.Sizeof([1]C.godot_transform2d{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotTransform2d) Ref() *C.godot_transform2d {
	if x == nil {
		return nil
	}
	return x.ref77dacf6
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotTransform2d) Free() {
	if x != nil && x.allocs77dacf6 != nil {
		x.allocs77dacf6.(*cgoAllocMap).Free()
		x.ref77dacf6 = nil
	}
}

// NewGodotTransform2dRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotTransform2dRef(ref unsafe.Pointer) *GodotTransform2d {
	if ref == nil {
		return nil
	}
	obj := new(GodotTransform2d)
	obj.ref77dacf6 = (*C.godot_transform2d)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotTransform2d) PassRef() (*C.godot_transform2d, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref77dacf6 != nil {
		return x.ref77dacf6, nil
	}
	mem77dacf6 := allocGodotTransform2dMemory(1)
	ref77dacf6 := (*C.godot_transform2d)(mem77dacf6)
	allocs77dacf6 := new(cgoAllocMap)
	allocs77dacf6.Add(mem77dacf6)

	var c_dont_touch_that_allocs *cgoAllocMap
	ref77dacf6._dont_touch_that, c_dont_touch_that_allocs = *(*[24]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocs77dacf6.Borrow(c_dont_touch_that_allocs)

	x.ref77dacf6 = ref77dacf6
	x.allocs77dacf6 = allocs77dacf6
	return ref77dacf6, allocs77dacf6

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotTransform2d) PassValue() (C.godot_transform2d, *cgoAllocMap) {
	if x.ref77dacf6 != nil {
		return *x.ref77dacf6, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotTransform2d) Deref() {
	if x.ref77dacf6 == nil {
		return
	}
	x.DontTouchThat = *(*[24]byte)(unsafe.Pointer(&x.ref77dacf6._dont_touch_that))
}

// allocGodotStringNameMemory allocates memory for type C.godot_string_name in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotStringNameMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotStringNameValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotStringNameValue = unsafe.Sizeof([1]C.godot_string_name{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotStringName) Ref() *C.godot_string_name {
	if x == nil {
		return nil
	}
	return x.ref895548fc
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotStringName) Free() {
	if x != nil && x.allocs895548fc != nil {
		x.allocs895548fc.(*cgoAllocMap).Free()
		x.ref895548fc = nil
	}
}

// NewGodotStringNameRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotStringNameRef(ref unsafe.Pointer) *GodotStringName {
	if ref == nil {
		return nil
	}
	obj := new(GodotStringName)
	obj.ref895548fc = (*C.godot_string_name)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotStringName) PassRef() (*C.godot_string_name, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref895548fc != nil {
		return x.ref895548fc, nil
	}
	mem895548fc := allocGodotStringNameMemory(1)
	ref895548fc := (*C.godot_string_name)(mem895548fc)
	allocs895548fc := new(cgoAllocMap)
	allocs895548fc.Add(mem895548fc)

	var c_dont_touch_that_allocs *cgoAllocMap
	ref895548fc._dont_touch_that, c_dont_touch_that_allocs = *(*[8]C.uint8_t)(unsafe.Pointer(&x.DontTouchThat)), cgoAllocsUnknown
	allocs895548fc.Borrow(c_dont_touch_that_allocs)

	x.ref895548fc = ref895548fc
	x.allocs895548fc = allocs895548fc
	return ref895548fc, allocs895548fc

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotStringName) PassValue() (C.godot_string_name, *cgoAllocMap) {
	if x.ref895548fc != nil {
		return *x.ref895548fc, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotStringName) Deref() {
	if x.ref895548fc == nil {
		return
	}
	x.DontTouchThat = *(*[8]byte)(unsafe.Pointer(&x.ref895548fc._dont_touch_that))
}

// allocGodotArvrInterfaceGdnativeMemory allocates memory for type C.godot_arvr_interface_gdnative in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotArvrInterfaceGdnativeMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotArvrInterfaceGdnativeValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotArvrInterfaceGdnativeValue = unsafe.Sizeof([1]C.godot_arvr_interface_gdnative{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotArvrInterfaceGdnative) Ref() *C.godot_arvr_interface_gdnative {
	if x == nil {
		return nil
	}
	return x.reff96a0b88
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotArvrInterfaceGdnative) Free() {
	if x != nil && x.allocsf96a0b88 != nil {
		x.allocsf96a0b88.(*cgoAllocMap).Free()
		x.reff96a0b88 = nil
	}
}

// NewGodotArvrInterfaceGdnativeRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotArvrInterfaceGdnativeRef(ref unsafe.Pointer) *GodotArvrInterfaceGdnative {
	if ref == nil {
		return nil
	}
	obj := new(GodotArvrInterfaceGdnative)
	obj.reff96a0b88 = (*C.godot_arvr_interface_gdnative)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotArvrInterfaceGdnative) PassRef() (*C.godot_arvr_interface_gdnative, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.reff96a0b88 != nil {
		return x.reff96a0b88, nil
	}
	memf96a0b88 := allocGodotArvrInterfaceGdnativeMemory(1)
	reff96a0b88 := (*C.godot_arvr_interface_gdnative)(memf96a0b88)
	allocsf96a0b88 := new(cgoAllocMap)
	allocsf96a0b88.Add(memf96a0b88)

	var cconstructor_allocs *cgoAllocMap
	reff96a0b88.constructor, cconstructor_allocs = x.Constructor.PassRef()
	allocsf96a0b88.Borrow(cconstructor_allocs)

	var cdestructor_allocs *cgoAllocMap
	reff96a0b88.destructor, cdestructor_allocs = x.Destructor.PassRef()
	allocsf96a0b88.Borrow(cdestructor_allocs)

	var cget_name_allocs *cgoAllocMap
	reff96a0b88.get_name, cget_name_allocs = x.GetName.PassRef()
	allocsf96a0b88.Borrow(cget_name_allocs)

	var cget_capabilities_allocs *cgoAllocMap
	reff96a0b88.get_capabilities, cget_capabilities_allocs = x.GetCapabilities.PassRef()
	allocsf96a0b88.Borrow(cget_capabilities_allocs)

	var cget_anchor_detection_is_enabled_allocs *cgoAllocMap
	reff96a0b88.get_anchor_detection_is_enabled, cget_anchor_detection_is_enabled_allocs = x.GetAnchorDetectionIsEnabled.PassRef()
	allocsf96a0b88.Borrow(cget_anchor_detection_is_enabled_allocs)

	var cset_anchor_detection_is_enabled_allocs *cgoAllocMap
	reff96a0b88.set_anchor_detection_is_enabled, cset_anchor_detection_is_enabled_allocs = x.SetAnchorDetectionIsEnabled.PassRef()
	allocsf96a0b88.Borrow(cset_anchor_detection_is_enabled_allocs)

	var cis_stereo_allocs *cgoAllocMap
	reff96a0b88.is_stereo, cis_stereo_allocs = x.IsStereo.PassRef()
	allocsf96a0b88.Borrow(cis_stereo_allocs)

	var cis_initialized_allocs *cgoAllocMap
	reff96a0b88.is_initialized, cis_initialized_allocs = x.IsInitialized.PassRef()
	allocsf96a0b88.Borrow(cis_initialized_allocs)

	var cinitialize_allocs *cgoAllocMap
	reff96a0b88.initialize, cinitialize_allocs = x.Initialize.PassRef()
	allocsf96a0b88.Borrow(cinitialize_allocs)

	var cuninitialize_allocs *cgoAllocMap
	reff96a0b88.uninitialize, cuninitialize_allocs = x.Uninitialize.PassRef()
	allocsf96a0b88.Borrow(cuninitialize_allocs)

	var cget_render_targetsize_allocs *cgoAllocMap
	reff96a0b88.get_render_targetsize, cget_render_targetsize_allocs = x.GetRenderTargetsize.PassRef()
	allocsf96a0b88.Borrow(cget_render_targetsize_allocs)

	var cget_transform_for_eye_allocs *cgoAllocMap
	reff96a0b88.get_transform_for_eye, cget_transform_for_eye_allocs = x.GetTransformForEye.PassRef()
	allocsf96a0b88.Borrow(cget_transform_for_eye_allocs)

	var cfill_projection_for_eye_allocs *cgoAllocMap
	reff96a0b88.fill_projection_for_eye, cfill_projection_for_eye_allocs = x.FillProjectionForEye.PassRef()
	allocsf96a0b88.Borrow(cfill_projection_for_eye_allocs)

	var ccommit_for_eye_allocs *cgoAllocMap
	reff96a0b88.commit_for_eye, ccommit_for_eye_allocs = x.CommitForEye.PassRef()
	allocsf96a0b88.Borrow(ccommit_for_eye_allocs)

	var cprocess_allocs *cgoAllocMap
	reff96a0b88.process, cprocess_allocs = x.Process.PassRef()
	allocsf96a0b88.Borrow(cprocess_allocs)

	x.reff96a0b88 = reff96a0b88
	x.allocsf96a0b88 = allocsf96a0b88
	return reff96a0b88, allocsf96a0b88

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotArvrInterfaceGdnative) PassValue() (C.godot_arvr_interface_gdnative, *cgoAllocMap) {
	if x.reff96a0b88 != nil {
		return *x.reff96a0b88, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotArvrInterfaceGdnative) Deref() {
	if x.reff96a0b88 == nil {
		return
	}
	x.Constructor = NewRef(unsafe.Pointer(x.reff96a0b88.constructor))
	x.Destructor = NewRef(unsafe.Pointer(x.reff96a0b88.destructor))
	x.GetName = NewGodotStringRef(unsafe.Pointer(x.reff96a0b88.get_name))
	x.GetCapabilities = NewGodotIntRef(unsafe.Pointer(x.reff96a0b88.get_capabilities))
	x.GetAnchorDetectionIsEnabled = NewGodotBoolRef(unsafe.Pointer(x.reff96a0b88.get_anchor_detection_is_enabled))
	x.SetAnchorDetectionIsEnabled = NewRef(unsafe.Pointer(x.reff96a0b88.set_anchor_detection_is_enabled))
	x.IsStereo = NewGodotBoolRef(unsafe.Pointer(x.reff96a0b88.is_stereo))
	x.IsInitialized = NewGodotBoolRef(unsafe.Pointer(x.reff96a0b88.is_initialized))
	x.Initialize = NewGodotBoolRef(unsafe.Pointer(x.reff96a0b88.initialize))
	x.Uninitialize = NewRef(unsafe.Pointer(x.reff96a0b88.uninitialize))
	x.GetRenderTargetsize = NewGodotVector2Ref(unsafe.Pointer(x.reff96a0b88.get_render_targetsize))
	x.GetTransformForEye = NewGodotTransformRef(unsafe.Pointer(x.reff96a0b88.get_transform_for_eye))
	x.FillProjectionForEye = NewRef(unsafe.Pointer(x.reff96a0b88.fill_projection_for_eye))
	x.CommitForEye = NewRef(unsafe.Pointer(x.reff96a0b88.commit_for_eye))
	x.Process = NewRef(unsafe.Pointer(x.reff96a0b88.process))
}

// allocGodotPropertyAttributesMemory allocates memory for type C.godot_property_attributes in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotPropertyAttributesMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotPropertyAttributesValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotPropertyAttributesValue = unsafe.Sizeof([1]C.godot_property_attributes{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotPropertyAttributes) Ref() *C.godot_property_attributes {
	if x == nil {
		return nil
	}
	return x.ref431c473b
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotPropertyAttributes) Free() {
	if x != nil && x.allocs431c473b != nil {
		x.allocs431c473b.(*cgoAllocMap).Free()
		x.ref431c473b = nil
	}
}

// NewGodotPropertyAttributesRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotPropertyAttributesRef(ref unsafe.Pointer) *GodotPropertyAttributes {
	if ref == nil {
		return nil
	}
	obj := new(GodotPropertyAttributes)
	obj.ref431c473b = (*C.godot_property_attributes)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotPropertyAttributes) PassRef() (*C.godot_property_attributes, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref431c473b != nil {
		return x.ref431c473b, nil
	}
	mem431c473b := allocGodotPropertyAttributesMemory(1)
	ref431c473b := (*C.godot_property_attributes)(mem431c473b)
	allocs431c473b := new(cgoAllocMap)
	allocs431c473b.Add(mem431c473b)

	var crset_type_allocs *cgoAllocMap
	ref431c473b.rset_type, crset_type_allocs = (C.godot_method_rpc_mode)(x.RsetType), cgoAllocsUnknown
	allocs431c473b.Borrow(crset_type_allocs)

	var c_type_allocs *cgoAllocMap
	ref431c473b._type, c_type_allocs = (C.godot_int)(x.Type), cgoAllocsUnknown
	allocs431c473b.Borrow(c_type_allocs)

	var chint_allocs *cgoAllocMap
	ref431c473b.hint, chint_allocs = (C.godot_property_hint)(x.Hint), cgoAllocsUnknown
	allocs431c473b.Borrow(chint_allocs)

	var chint_string_allocs *cgoAllocMap
	ref431c473b.hint_string, chint_string_allocs = x.HintString.PassValue()
	allocs431c473b.Borrow(chint_string_allocs)

	var cusage_allocs *cgoAllocMap
	ref431c473b.usage, cusage_allocs = (C.godot_property_usage_flags)(x.Usage), cgoAllocsUnknown
	allocs431c473b.Borrow(cusage_allocs)

	var cdefault_value_allocs *cgoAllocMap
	ref431c473b.default_value, cdefault_value_allocs = x.DefaultValue.PassValue()
	allocs431c473b.Borrow(cdefault_value_allocs)

	x.ref431c473b = ref431c473b
	x.allocs431c473b = allocs431c473b
	return ref431c473b, allocs431c473b

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotPropertyAttributes) PassValue() (C.godot_property_attributes, *cgoAllocMap) {
	if x.ref431c473b != nil {
		return *x.ref431c473b, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotPropertyAttributes) Deref() {
	if x.ref431c473b == nil {
		return
	}
	x.RsetType = (GodotMethodRpcMode)(x.ref431c473b.rset_type)
	x.Type = (GodotInt)(x.ref431c473b._type)
	x.Hint = (GodotPropertyHint)(x.ref431c473b.hint)
	x.HintString = *NewGodotStringRef(unsafe.Pointer(&x.ref431c473b.hint_string))
	x.Usage = (GodotPropertyUsageFlags)(x.ref431c473b.usage)
	x.DefaultValue = *NewGodotVariantRef(unsafe.Pointer(&x.ref431c473b.default_value))
}

// allocGodotInstanceCreateFuncMemory allocates memory for type C.godot_instance_create_func in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotInstanceCreateFuncMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotInstanceCreateFuncValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotInstanceCreateFuncValue = unsafe.Sizeof([1]C.godot_instance_create_func{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotInstanceCreateFunc) Ref() *C.godot_instance_create_func {
	if x == nil {
		return nil
	}
	return x.ref70ecb5db
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotInstanceCreateFunc) Free() {
	if x != nil && x.allocs70ecb5db != nil {
		x.allocs70ecb5db.(*cgoAllocMap).Free()
		x.ref70ecb5db = nil
	}
}

// NewGodotInstanceCreateFuncRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotInstanceCreateFuncRef(ref unsafe.Pointer) *GodotInstanceCreateFunc {
	if ref == nil {
		return nil
	}
	obj := new(GodotInstanceCreateFunc)
	obj.ref70ecb5db = (*C.godot_instance_create_func)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotInstanceCreateFunc) PassRef() (*C.godot_instance_create_func, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref70ecb5db != nil {
		return x.ref70ecb5db, nil
	}
	mem70ecb5db := allocGodotInstanceCreateFuncMemory(1)
	ref70ecb5db := (*C.godot_instance_create_func)(mem70ecb5db)
	allocs70ecb5db := new(cgoAllocMap)
	allocs70ecb5db.Add(mem70ecb5db)

	var ccreate_func_allocs *cgoAllocMap
	ref70ecb5db.create_func, ccreate_func_allocs = x.CreateFunc.PassRef()
	allocs70ecb5db.Borrow(ccreate_func_allocs)

	var cmethod_data_allocs *cgoAllocMap
	ref70ecb5db.method_data, cmethod_data_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.MethodData)), cgoAllocsUnknown
	allocs70ecb5db.Borrow(cmethod_data_allocs)

	var cfree_func_allocs *cgoAllocMap
	ref70ecb5db.free_func, cfree_func_allocs = x.FreeFunc.PassRef()
	allocs70ecb5db.Borrow(cfree_func_allocs)

	x.ref70ecb5db = ref70ecb5db
	x.allocs70ecb5db = allocs70ecb5db
	return ref70ecb5db, allocs70ecb5db

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotInstanceCreateFunc) PassValue() (C.godot_instance_create_func, *cgoAllocMap) {
	if x.ref70ecb5db != nil {
		return *x.ref70ecb5db, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotInstanceCreateFunc) Deref() {
	if x.ref70ecb5db == nil {
		return
	}
	x.CreateFunc = NewRef(unsafe.Pointer(x.ref70ecb5db.create_func))
	x.MethodData = (unsafe.Pointer)(unsafe.Pointer(x.ref70ecb5db.method_data))
	x.FreeFunc = NewRef(unsafe.Pointer(x.ref70ecb5db.free_func))
}

// allocGodotInstanceDestroyFuncMemory allocates memory for type C.godot_instance_destroy_func in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotInstanceDestroyFuncMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotInstanceDestroyFuncValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotInstanceDestroyFuncValue = unsafe.Sizeof([1]C.godot_instance_destroy_func{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotInstanceDestroyFunc) Ref() *C.godot_instance_destroy_func {
	if x == nil {
		return nil
	}
	return x.refd0d05668
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotInstanceDestroyFunc) Free() {
	if x != nil && x.allocsd0d05668 != nil {
		x.allocsd0d05668.(*cgoAllocMap).Free()
		x.refd0d05668 = nil
	}
}

// NewGodotInstanceDestroyFuncRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotInstanceDestroyFuncRef(ref unsafe.Pointer) *GodotInstanceDestroyFunc {
	if ref == nil {
		return nil
	}
	obj := new(GodotInstanceDestroyFunc)
	obj.refd0d05668 = (*C.godot_instance_destroy_func)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotInstanceDestroyFunc) PassRef() (*C.godot_instance_destroy_func, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refd0d05668 != nil {
		return x.refd0d05668, nil
	}
	memd0d05668 := allocGodotInstanceDestroyFuncMemory(1)
	refd0d05668 := (*C.godot_instance_destroy_func)(memd0d05668)
	allocsd0d05668 := new(cgoAllocMap)
	allocsd0d05668.Add(memd0d05668)

	var cdestroy_func_allocs *cgoAllocMap
	refd0d05668.destroy_func, cdestroy_func_allocs = x.DestroyFunc.PassRef()
	allocsd0d05668.Borrow(cdestroy_func_allocs)

	var cmethod_data_allocs *cgoAllocMap
	refd0d05668.method_data, cmethod_data_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.MethodData)), cgoAllocsUnknown
	allocsd0d05668.Borrow(cmethod_data_allocs)

	var cfree_func_allocs *cgoAllocMap
	refd0d05668.free_func, cfree_func_allocs = x.FreeFunc.PassRef()
	allocsd0d05668.Borrow(cfree_func_allocs)

	x.refd0d05668 = refd0d05668
	x.allocsd0d05668 = allocsd0d05668
	return refd0d05668, allocsd0d05668

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotInstanceDestroyFunc) PassValue() (C.godot_instance_destroy_func, *cgoAllocMap) {
	if x.refd0d05668 != nil {
		return *x.refd0d05668, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotInstanceDestroyFunc) Deref() {
	if x.refd0d05668 == nil {
		return
	}
	x.DestroyFunc = NewRef(unsafe.Pointer(x.refd0d05668.destroy_func))
	x.MethodData = (unsafe.Pointer)(unsafe.Pointer(x.refd0d05668.method_data))
	x.FreeFunc = NewRef(unsafe.Pointer(x.refd0d05668.free_func))
}

// allocGodotMethodAttributesMemory allocates memory for type C.godot_method_attributes in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotMethodAttributesMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotMethodAttributesValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotMethodAttributesValue = unsafe.Sizeof([1]C.godot_method_attributes{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotMethodAttributes) Ref() *C.godot_method_attributes {
	if x == nil {
		return nil
	}
	return x.ref66a6c5c9
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotMethodAttributes) Free() {
	if x != nil && x.allocs66a6c5c9 != nil {
		x.allocs66a6c5c9.(*cgoAllocMap).Free()
		x.ref66a6c5c9 = nil
	}
}

// NewGodotMethodAttributesRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotMethodAttributesRef(ref unsafe.Pointer) *GodotMethodAttributes {
	if ref == nil {
		return nil
	}
	obj := new(GodotMethodAttributes)
	obj.ref66a6c5c9 = (*C.godot_method_attributes)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotMethodAttributes) PassRef() (*C.godot_method_attributes, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref66a6c5c9 != nil {
		return x.ref66a6c5c9, nil
	}
	mem66a6c5c9 := allocGodotMethodAttributesMemory(1)
	ref66a6c5c9 := (*C.godot_method_attributes)(mem66a6c5c9)
	allocs66a6c5c9 := new(cgoAllocMap)
	allocs66a6c5c9.Add(mem66a6c5c9)

	var crpc_type_allocs *cgoAllocMap
	ref66a6c5c9.rpc_type, crpc_type_allocs = (C.godot_method_rpc_mode)(x.RpcType), cgoAllocsUnknown
	allocs66a6c5c9.Borrow(crpc_type_allocs)

	x.ref66a6c5c9 = ref66a6c5c9
	x.allocs66a6c5c9 = allocs66a6c5c9
	return ref66a6c5c9, allocs66a6c5c9

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotMethodAttributes) PassValue() (C.godot_method_attributes, *cgoAllocMap) {
	if x.ref66a6c5c9 != nil {
		return *x.ref66a6c5c9, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotMethodAttributes) Deref() {
	if x.ref66a6c5c9 == nil {
		return
	}
	x.RpcType = (GodotMethodRpcMode)(x.ref66a6c5c9.rpc_type)
}

// allocGodotInstanceMethodMemory allocates memory for type C.godot_instance_method in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotInstanceMethodMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotInstanceMethodValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotInstanceMethodValue = unsafe.Sizeof([1]C.godot_instance_method{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotInstanceMethod) Ref() *C.godot_instance_method {
	if x == nil {
		return nil
	}
	return x.ref10e1583e
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotInstanceMethod) Free() {
	if x != nil && x.allocs10e1583e != nil {
		x.allocs10e1583e.(*cgoAllocMap).Free()
		x.ref10e1583e = nil
	}
}

// NewGodotInstanceMethodRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotInstanceMethodRef(ref unsafe.Pointer) *GodotInstanceMethod {
	if ref == nil {
		return nil
	}
	obj := new(GodotInstanceMethod)
	obj.ref10e1583e = (*C.godot_instance_method)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotInstanceMethod) PassRef() (*C.godot_instance_method, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref10e1583e != nil {
		return x.ref10e1583e, nil
	}
	mem10e1583e := allocGodotInstanceMethodMemory(1)
	ref10e1583e := (*C.godot_instance_method)(mem10e1583e)
	allocs10e1583e := new(cgoAllocMap)
	allocs10e1583e.Add(mem10e1583e)

	var cmethod_allocs *cgoAllocMap
	ref10e1583e.method, cmethod_allocs = x.Method.PassRef()
	allocs10e1583e.Borrow(cmethod_allocs)

	var cmethod_data_allocs *cgoAllocMap
	ref10e1583e.method_data, cmethod_data_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.MethodData)), cgoAllocsUnknown
	allocs10e1583e.Borrow(cmethod_data_allocs)

	var cfree_func_allocs *cgoAllocMap
	ref10e1583e.free_func, cfree_func_allocs = x.FreeFunc.PassRef()
	allocs10e1583e.Borrow(cfree_func_allocs)

	x.ref10e1583e = ref10e1583e
	x.allocs10e1583e = allocs10e1583e
	return ref10e1583e, allocs10e1583e

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotInstanceMethod) PassValue() (C.godot_instance_method, *cgoAllocMap) {
	if x.ref10e1583e != nil {
		return *x.ref10e1583e, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotInstanceMethod) Deref() {
	if x.ref10e1583e == nil {
		return
	}
	x.Method = NewGodotVariantRef(unsafe.Pointer(x.ref10e1583e.method))
	x.MethodData = (unsafe.Pointer)(unsafe.Pointer(x.ref10e1583e.method_data))
	x.FreeFunc = NewRef(unsafe.Pointer(x.ref10e1583e.free_func))
}

// allocGodotPropertySetFuncMemory allocates memory for type C.godot_property_set_func in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotPropertySetFuncMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotPropertySetFuncValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotPropertySetFuncValue = unsafe.Sizeof([1]C.godot_property_set_func{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotPropertySetFunc) Ref() *C.godot_property_set_func {
	if x == nil {
		return nil
	}
	return x.refc9844af
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotPropertySetFunc) Free() {
	if x != nil && x.allocsc9844af != nil {
		x.allocsc9844af.(*cgoAllocMap).Free()
		x.refc9844af = nil
	}
}

// NewGodotPropertySetFuncRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotPropertySetFuncRef(ref unsafe.Pointer) *GodotPropertySetFunc {
	if ref == nil {
		return nil
	}
	obj := new(GodotPropertySetFunc)
	obj.refc9844af = (*C.godot_property_set_func)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotPropertySetFunc) PassRef() (*C.godot_property_set_func, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refc9844af != nil {
		return x.refc9844af, nil
	}
	memc9844af := allocGodotPropertySetFuncMemory(1)
	refc9844af := (*C.godot_property_set_func)(memc9844af)
	allocsc9844af := new(cgoAllocMap)
	allocsc9844af.Add(memc9844af)

	var cset_func_allocs *cgoAllocMap
	refc9844af.set_func, cset_func_allocs = x.SetFunc.PassRef()
	allocsc9844af.Borrow(cset_func_allocs)

	var cmethod_data_allocs *cgoAllocMap
	refc9844af.method_data, cmethod_data_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.MethodData)), cgoAllocsUnknown
	allocsc9844af.Borrow(cmethod_data_allocs)

	var cfree_func_allocs *cgoAllocMap
	refc9844af.free_func, cfree_func_allocs = x.FreeFunc.PassRef()
	allocsc9844af.Borrow(cfree_func_allocs)

	x.refc9844af = refc9844af
	x.allocsc9844af = allocsc9844af
	return refc9844af, allocsc9844af

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotPropertySetFunc) PassValue() (C.godot_property_set_func, *cgoAllocMap) {
	if x.refc9844af != nil {
		return *x.refc9844af, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotPropertySetFunc) Deref() {
	if x.refc9844af == nil {
		return
	}
	x.SetFunc = NewRef(unsafe.Pointer(x.refc9844af.set_func))
	x.MethodData = (unsafe.Pointer)(unsafe.Pointer(x.refc9844af.method_data))
	x.FreeFunc = NewRef(unsafe.Pointer(x.refc9844af.free_func))
}

// allocGodotPropertyGetFuncMemory allocates memory for type C.godot_property_get_func in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotPropertyGetFuncMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotPropertyGetFuncValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotPropertyGetFuncValue = unsafe.Sizeof([1]C.godot_property_get_func{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotPropertyGetFunc) Ref() *C.godot_property_get_func {
	if x == nil {
		return nil
	}
	return x.reff4697b7e
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotPropertyGetFunc) Free() {
	if x != nil && x.allocsf4697b7e != nil {
		x.allocsf4697b7e.(*cgoAllocMap).Free()
		x.reff4697b7e = nil
	}
}

// NewGodotPropertyGetFuncRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotPropertyGetFuncRef(ref unsafe.Pointer) *GodotPropertyGetFunc {
	if ref == nil {
		return nil
	}
	obj := new(GodotPropertyGetFunc)
	obj.reff4697b7e = (*C.godot_property_get_func)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotPropertyGetFunc) PassRef() (*C.godot_property_get_func, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.reff4697b7e != nil {
		return x.reff4697b7e, nil
	}
	memf4697b7e := allocGodotPropertyGetFuncMemory(1)
	reff4697b7e := (*C.godot_property_get_func)(memf4697b7e)
	allocsf4697b7e := new(cgoAllocMap)
	allocsf4697b7e.Add(memf4697b7e)

	var cget_func_allocs *cgoAllocMap
	reff4697b7e.get_func, cget_func_allocs = x.GetFunc.PassRef()
	allocsf4697b7e.Borrow(cget_func_allocs)

	var cmethod_data_allocs *cgoAllocMap
	reff4697b7e.method_data, cmethod_data_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.MethodData)), cgoAllocsUnknown
	allocsf4697b7e.Borrow(cmethod_data_allocs)

	var cfree_func_allocs *cgoAllocMap
	reff4697b7e.free_func, cfree_func_allocs = x.FreeFunc.PassRef()
	allocsf4697b7e.Borrow(cfree_func_allocs)

	x.reff4697b7e = reff4697b7e
	x.allocsf4697b7e = allocsf4697b7e
	return reff4697b7e, allocsf4697b7e

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotPropertyGetFunc) PassValue() (C.godot_property_get_func, *cgoAllocMap) {
	if x.reff4697b7e != nil {
		return *x.reff4697b7e, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotPropertyGetFunc) Deref() {
	if x.reff4697b7e == nil {
		return
	}
	x.GetFunc = NewGodotVariantRef(unsafe.Pointer(x.reff4697b7e.get_func))
	x.MethodData = (unsafe.Pointer)(unsafe.Pointer(x.reff4697b7e.method_data))
	x.FreeFunc = NewRef(unsafe.Pointer(x.reff4697b7e.free_func))
}

// allocGodotSignalArgumentMemory allocates memory for type C.godot_signal_argument in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotSignalArgumentMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotSignalArgumentValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotSignalArgumentValue = unsafe.Sizeof([1]C.godot_signal_argument{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotSignalArgument) Ref() *C.godot_signal_argument {
	if x == nil {
		return nil
	}
	return x.refc21e72ac
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotSignalArgument) Free() {
	if x != nil && x.allocsc21e72ac != nil {
		x.allocsc21e72ac.(*cgoAllocMap).Free()
		x.refc21e72ac = nil
	}
}

// NewGodotSignalArgumentRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotSignalArgumentRef(ref unsafe.Pointer) *GodotSignalArgument {
	if ref == nil {
		return nil
	}
	obj := new(GodotSignalArgument)
	obj.refc21e72ac = (*C.godot_signal_argument)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotSignalArgument) PassRef() (*C.godot_signal_argument, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refc21e72ac != nil {
		return x.refc21e72ac, nil
	}
	memc21e72ac := allocGodotSignalArgumentMemory(1)
	refc21e72ac := (*C.godot_signal_argument)(memc21e72ac)
	allocsc21e72ac := new(cgoAllocMap)
	allocsc21e72ac.Add(memc21e72ac)

	var cname_allocs *cgoAllocMap
	refc21e72ac.name, cname_allocs = x.Name.PassValue()
	allocsc21e72ac.Borrow(cname_allocs)

	var c_type_allocs *cgoAllocMap
	refc21e72ac._type, c_type_allocs = (C.godot_int)(x.Type), cgoAllocsUnknown
	allocsc21e72ac.Borrow(c_type_allocs)

	var chint_allocs *cgoAllocMap
	refc21e72ac.hint, chint_allocs = (C.godot_property_hint)(x.Hint), cgoAllocsUnknown
	allocsc21e72ac.Borrow(chint_allocs)

	var chint_string_allocs *cgoAllocMap
	refc21e72ac.hint_string, chint_string_allocs = x.HintString.PassValue()
	allocsc21e72ac.Borrow(chint_string_allocs)

	var cusage_allocs *cgoAllocMap
	refc21e72ac.usage, cusage_allocs = (C.godot_property_usage_flags)(x.Usage), cgoAllocsUnknown
	allocsc21e72ac.Borrow(cusage_allocs)

	var cdefault_value_allocs *cgoAllocMap
	refc21e72ac.default_value, cdefault_value_allocs = x.DefaultValue.PassValue()
	allocsc21e72ac.Borrow(cdefault_value_allocs)

	x.refc21e72ac = refc21e72ac
	x.allocsc21e72ac = allocsc21e72ac
	return refc21e72ac, allocsc21e72ac

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotSignalArgument) PassValue() (C.godot_signal_argument, *cgoAllocMap) {
	if x.refc21e72ac != nil {
		return *x.refc21e72ac, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotSignalArgument) Deref() {
	if x.refc21e72ac == nil {
		return
	}
	x.Name = *NewGodotStringRef(unsafe.Pointer(&x.refc21e72ac.name))
	x.Type = (GodotInt)(x.refc21e72ac._type)
	x.Hint = (GodotPropertyHint)(x.refc21e72ac.hint)
	x.HintString = *NewGodotStringRef(unsafe.Pointer(&x.refc21e72ac.hint_string))
	x.Usage = (GodotPropertyUsageFlags)(x.refc21e72ac.usage)
	x.DefaultValue = *NewGodotVariantRef(unsafe.Pointer(&x.refc21e72ac.default_value))
}

// allocGodotSignalMemory allocates memory for type C.godot_signal in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotSignalMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotSignalValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotSignalValue = unsafe.Sizeof([1]C.godot_signal{})

// unpackSGodotSignalArgument transforms a sliced Go data structure into plain C format.
func unpackSGodotSignalArgument(x []GodotSignalArgument) (unpacked *C.godot_signal_argument, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_signal_argument) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotSignalArgumentMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_signal_argument)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_signal_argument)(unsafe.Pointer(h.Data))
	return
}

// unpackSGodotVariant transforms a sliced Go data structure into plain C format.
func unpackSGodotVariant(x []GodotVariant) (unpacked *C.godot_variant, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_variant) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotVariantMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_variant)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_variant)(unsafe.Pointer(h.Data))
	return
}

// packSGodotSignalArgument reads sliced Go data structure out from plain C format.
func packSGodotSignalArgument(v []GodotSignalArgument, ptr0 *C.godot_signal_argument) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotSignalArgumentValue]C.godot_signal_argument)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotSignalArgumentRef(unsafe.Pointer(&ptr1))
	}
}

// packSGodotVariant reads sliced Go data structure out from plain C format.
func packSGodotVariant(v []GodotVariant, ptr0 *C.godot_variant) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotVariantValue]C.godot_variant)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotVariantRef(unsafe.Pointer(&ptr1))
	}
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotSignal) Ref() *C.godot_signal {
	if x == nil {
		return nil
	}
	return x.ref87acf90b
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotSignal) Free() {
	if x != nil && x.allocs87acf90b != nil {
		x.allocs87acf90b.(*cgoAllocMap).Free()
		x.ref87acf90b = nil
	}
}

// NewGodotSignalRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotSignalRef(ref unsafe.Pointer) *GodotSignal {
	if ref == nil {
		return nil
	}
	obj := new(GodotSignal)
	obj.ref87acf90b = (*C.godot_signal)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotSignal) PassRef() (*C.godot_signal, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref87acf90b != nil {
		return x.ref87acf90b, nil
	}
	mem87acf90b := allocGodotSignalMemory(1)
	ref87acf90b := (*C.godot_signal)(mem87acf90b)
	allocs87acf90b := new(cgoAllocMap)
	allocs87acf90b.Add(mem87acf90b)

	var cname_allocs *cgoAllocMap
	ref87acf90b.name, cname_allocs = x.Name.PassValue()
	allocs87acf90b.Borrow(cname_allocs)

	var cnum_args_allocs *cgoAllocMap
	ref87acf90b.num_args, cnum_args_allocs = (C.int)(x.NumArgs), cgoAllocsUnknown
	allocs87acf90b.Borrow(cnum_args_allocs)

	var cargs_allocs *cgoAllocMap
	ref87acf90b.args, cargs_allocs = unpackSGodotSignalArgument(x.Args)
	allocs87acf90b.Borrow(cargs_allocs)

	var cnum_default_args_allocs *cgoAllocMap
	ref87acf90b.num_default_args, cnum_default_args_allocs = (C.int)(x.NumDefaultArgs), cgoAllocsUnknown
	allocs87acf90b.Borrow(cnum_default_args_allocs)

	var cdefault_args_allocs *cgoAllocMap
	ref87acf90b.default_args, cdefault_args_allocs = unpackSGodotVariant(x.DefaultArgs)
	allocs87acf90b.Borrow(cdefault_args_allocs)

	x.ref87acf90b = ref87acf90b
	x.allocs87acf90b = allocs87acf90b
	return ref87acf90b, allocs87acf90b

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotSignal) PassValue() (C.godot_signal, *cgoAllocMap) {
	if x.ref87acf90b != nil {
		return *x.ref87acf90b, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotSignal) Deref() {
	if x.ref87acf90b == nil {
		return
	}
	x.Name = *NewGodotStringRef(unsafe.Pointer(&x.ref87acf90b.name))
	x.NumArgs = (int32)(x.ref87acf90b.num_args)
	packSGodotSignalArgument(x.Args, x.ref87acf90b.args)
	x.NumDefaultArgs = (int32)(x.ref87acf90b.num_default_args)
	packSGodotVariant(x.DefaultArgs, x.ref87acf90b.default_args)
}

// allocGodotPluginscriptInstanceDescMemory allocates memory for type C.godot_pluginscript_instance_desc in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotPluginscriptInstanceDescMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotPluginscriptInstanceDescValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotPluginscriptInstanceDescValue = unsafe.Sizeof([1]C.godot_pluginscript_instance_desc{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotPluginscriptInstanceDesc) Ref() *C.godot_pluginscript_instance_desc {
	if x == nil {
		return nil
	}
	return x.refc0c19139
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotPluginscriptInstanceDesc) Free() {
	if x != nil && x.allocsc0c19139 != nil {
		x.allocsc0c19139.(*cgoAllocMap).Free()
		x.refc0c19139 = nil
	}
}

// NewGodotPluginscriptInstanceDescRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotPluginscriptInstanceDescRef(ref unsafe.Pointer) *GodotPluginscriptInstanceDesc {
	if ref == nil {
		return nil
	}
	obj := new(GodotPluginscriptInstanceDesc)
	obj.refc0c19139 = (*C.godot_pluginscript_instance_desc)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotPluginscriptInstanceDesc) PassRef() (*C.godot_pluginscript_instance_desc, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refc0c19139 != nil {
		return x.refc0c19139, nil
	}
	memc0c19139 := allocGodotPluginscriptInstanceDescMemory(1)
	refc0c19139 := (*C.godot_pluginscript_instance_desc)(memc0c19139)
	allocsc0c19139 := new(cgoAllocMap)
	allocsc0c19139.Add(memc0c19139)

	var cinit_allocs *cgoAllocMap
	refc0c19139.init, cinit_allocs = x.Init.PassRef()
	allocsc0c19139.Borrow(cinit_allocs)

	var cfinish_allocs *cgoAllocMap
	refc0c19139.finish, cfinish_allocs = x.Finish.PassRef()
	allocsc0c19139.Borrow(cfinish_allocs)

	var cset_prop_allocs *cgoAllocMap
	refc0c19139.set_prop, cset_prop_allocs = x.SetProp.PassRef()
	allocsc0c19139.Borrow(cset_prop_allocs)

	var cget_prop_allocs *cgoAllocMap
	refc0c19139.get_prop, cget_prop_allocs = x.GetProp.PassRef()
	allocsc0c19139.Borrow(cget_prop_allocs)

	var ccall_method_allocs *cgoAllocMap
	refc0c19139.call_method, ccall_method_allocs = x.CallMethod.PassRef()
	allocsc0c19139.Borrow(ccall_method_allocs)

	var cnotification_allocs *cgoAllocMap
	refc0c19139.notification, cnotification_allocs = x.Notification.PassRef()
	allocsc0c19139.Borrow(cnotification_allocs)

	var cget_rpc_mode_allocs *cgoAllocMap
	refc0c19139.get_rpc_mode, cget_rpc_mode_allocs = x.GetRpcMode.PassRef()
	allocsc0c19139.Borrow(cget_rpc_mode_allocs)

	var cget_rset_mode_allocs *cgoAllocMap
	refc0c19139.get_rset_mode, cget_rset_mode_allocs = x.GetRsetMode.PassRef()
	allocsc0c19139.Borrow(cget_rset_mode_allocs)

	var crefcount_incremented_allocs *cgoAllocMap
	refc0c19139.refcount_incremented, crefcount_incremented_allocs = x.RefcountIncremented.PassRef()
	allocsc0c19139.Borrow(crefcount_incremented_allocs)

	var crefcount_decremented_allocs *cgoAllocMap
	refc0c19139.refcount_decremented, crefcount_decremented_allocs = x.RefcountDecremented.PassRef()
	allocsc0c19139.Borrow(crefcount_decremented_allocs)

	x.refc0c19139 = refc0c19139
	x.allocsc0c19139 = allocsc0c19139
	return refc0c19139, allocsc0c19139

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotPluginscriptInstanceDesc) PassValue() (C.godot_pluginscript_instance_desc, *cgoAllocMap) {
	if x.refc0c19139 != nil {
		return *x.refc0c19139, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotPluginscriptInstanceDesc) Deref() {
	if x.refc0c19139 == nil {
		return
	}
	x.Init = NewGodotPluginscriptInstanceDataRef(unsafe.Pointer(x.refc0c19139.init))
	x.Finish = NewRef(unsafe.Pointer(x.refc0c19139.finish))
	x.SetProp = NewGodotBoolRef(unsafe.Pointer(x.refc0c19139.set_prop))
	x.GetProp = NewGodotBoolRef(unsafe.Pointer(x.refc0c19139.get_prop))
	x.CallMethod = NewGodotVariantRef(unsafe.Pointer(x.refc0c19139.call_method))
	x.Notification = NewRef(unsafe.Pointer(x.refc0c19139.notification))
	x.GetRpcMode = NewGodotMethodRpcModeRef(unsafe.Pointer(x.refc0c19139.get_rpc_mode))
	x.GetRsetMode = NewGodotMethodRpcModeRef(unsafe.Pointer(x.refc0c19139.get_rset_mode))
	x.RefcountIncremented = NewRef(unsafe.Pointer(x.refc0c19139.refcount_incremented))
	x.RefcountDecremented = NewRef(unsafe.Pointer(x.refc0c19139.refcount_decremented))
}

// allocGodotPluginscriptScriptManifestMemory allocates memory for type C.godot_pluginscript_script_manifest in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotPluginscriptScriptManifestMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotPluginscriptScriptManifestValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotPluginscriptScriptManifestValue = unsafe.Sizeof([1]C.godot_pluginscript_script_manifest{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotPluginscriptScriptManifest) Ref() *C.godot_pluginscript_script_manifest {
	if x == nil {
		return nil
	}
	return x.reffbf02dfd
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotPluginscriptScriptManifest) Free() {
	if x != nil && x.allocsfbf02dfd != nil {
		x.allocsfbf02dfd.(*cgoAllocMap).Free()
		x.reffbf02dfd = nil
	}
}

// NewGodotPluginscriptScriptManifestRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotPluginscriptScriptManifestRef(ref unsafe.Pointer) *GodotPluginscriptScriptManifest {
	if ref == nil {
		return nil
	}
	obj := new(GodotPluginscriptScriptManifest)
	obj.reffbf02dfd = (*C.godot_pluginscript_script_manifest)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotPluginscriptScriptManifest) PassRef() (*C.godot_pluginscript_script_manifest, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.reffbf02dfd != nil {
		return x.reffbf02dfd, nil
	}
	memfbf02dfd := allocGodotPluginscriptScriptManifestMemory(1)
	reffbf02dfd := (*C.godot_pluginscript_script_manifest)(memfbf02dfd)
	allocsfbf02dfd := new(cgoAllocMap)
	allocsfbf02dfd.Add(memfbf02dfd)

	var cdata_allocs *cgoAllocMap
	reffbf02dfd.data, cdata_allocs = *(*C.godot_pluginscript_script_data)(unsafe.Pointer(&x.Data)), cgoAllocsUnknown
	allocsfbf02dfd.Borrow(cdata_allocs)

	var cname_allocs *cgoAllocMap
	reffbf02dfd.name, cname_allocs = x.Name.PassValue()
	allocsfbf02dfd.Borrow(cname_allocs)

	var cis_tool_allocs *cgoAllocMap
	reffbf02dfd.is_tool, cis_tool_allocs = (C.godot_bool)(x.IsTool), cgoAllocsUnknown
	allocsfbf02dfd.Borrow(cis_tool_allocs)

	var cbase_allocs *cgoAllocMap
	reffbf02dfd.base, cbase_allocs = x.Base.PassValue()
	allocsfbf02dfd.Borrow(cbase_allocs)

	var cmember_lines_allocs *cgoAllocMap
	reffbf02dfd.member_lines, cmember_lines_allocs = x.MemberLines.PassValue()
	allocsfbf02dfd.Borrow(cmember_lines_allocs)

	var cmethods_allocs *cgoAllocMap
	reffbf02dfd.methods, cmethods_allocs = x.Methods.PassValue()
	allocsfbf02dfd.Borrow(cmethods_allocs)

	var csignals_allocs *cgoAllocMap
	reffbf02dfd.signals, csignals_allocs = x.Signals.PassValue()
	allocsfbf02dfd.Borrow(csignals_allocs)

	var cproperties_allocs *cgoAllocMap
	reffbf02dfd.properties, cproperties_allocs = x.Properties.PassValue()
	allocsfbf02dfd.Borrow(cproperties_allocs)

	x.reffbf02dfd = reffbf02dfd
	x.allocsfbf02dfd = allocsfbf02dfd
	return reffbf02dfd, allocsfbf02dfd

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotPluginscriptScriptManifest) PassValue() (C.godot_pluginscript_script_manifest, *cgoAllocMap) {
	if x.reffbf02dfd != nil {
		return *x.reffbf02dfd, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotPluginscriptScriptManifest) Deref() {
	if x.reffbf02dfd == nil {
		return
	}
	x.Data = (*GodotPluginscriptScriptData)(unsafe.Pointer(x.reffbf02dfd.data))
	x.Name = *NewGodotStringNameRef(unsafe.Pointer(&x.reffbf02dfd.name))
	x.IsTool = (GodotBool)(x.reffbf02dfd.is_tool)
	x.Base = *NewGodotStringNameRef(unsafe.Pointer(&x.reffbf02dfd.base))
	x.MemberLines = *NewGodotDictionaryRef(unsafe.Pointer(&x.reffbf02dfd.member_lines))
	x.Methods = *NewGodotArrayRef(unsafe.Pointer(&x.reffbf02dfd.methods))
	x.Signals = *NewGodotArrayRef(unsafe.Pointer(&x.reffbf02dfd.signals))
	x.Properties = *NewGodotArrayRef(unsafe.Pointer(&x.reffbf02dfd.properties))
}

// allocGodotPluginscriptScriptDescMemory allocates memory for type C.godot_pluginscript_script_desc in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotPluginscriptScriptDescMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotPluginscriptScriptDescValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotPluginscriptScriptDescValue = unsafe.Sizeof([1]C.godot_pluginscript_script_desc{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotPluginscriptScriptDesc) Ref() *C.godot_pluginscript_script_desc {
	if x == nil {
		return nil
	}
	return x.ref1aab3210
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotPluginscriptScriptDesc) Free() {
	if x != nil && x.allocs1aab3210 != nil {
		x.allocs1aab3210.(*cgoAllocMap).Free()
		x.ref1aab3210 = nil
	}
}

// NewGodotPluginscriptScriptDescRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotPluginscriptScriptDescRef(ref unsafe.Pointer) *GodotPluginscriptScriptDesc {
	if ref == nil {
		return nil
	}
	obj := new(GodotPluginscriptScriptDesc)
	obj.ref1aab3210 = (*C.godot_pluginscript_script_desc)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotPluginscriptScriptDesc) PassRef() (*C.godot_pluginscript_script_desc, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref1aab3210 != nil {
		return x.ref1aab3210, nil
	}
	mem1aab3210 := allocGodotPluginscriptScriptDescMemory(1)
	ref1aab3210 := (*C.godot_pluginscript_script_desc)(mem1aab3210)
	allocs1aab3210 := new(cgoAllocMap)
	allocs1aab3210.Add(mem1aab3210)

	var cinit_allocs *cgoAllocMap
	ref1aab3210.init, cinit_allocs = x.Init.PassRef()
	allocs1aab3210.Borrow(cinit_allocs)

	var cfinish_allocs *cgoAllocMap
	ref1aab3210.finish, cfinish_allocs = x.Finish.PassRef()
	allocs1aab3210.Borrow(cfinish_allocs)

	var cinstance_desc_allocs *cgoAllocMap
	ref1aab3210.instance_desc, cinstance_desc_allocs = x.InstanceDesc.PassValue()
	allocs1aab3210.Borrow(cinstance_desc_allocs)

	x.ref1aab3210 = ref1aab3210
	x.allocs1aab3210 = allocs1aab3210
	return ref1aab3210, allocs1aab3210

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotPluginscriptScriptDesc) PassValue() (C.godot_pluginscript_script_desc, *cgoAllocMap) {
	if x.ref1aab3210 != nil {
		return *x.ref1aab3210, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotPluginscriptScriptDesc) Deref() {
	if x.ref1aab3210 == nil {
		return
	}
	x.Init = NewGodotPluginscriptScriptManifestRef(unsafe.Pointer(x.ref1aab3210.init))
	x.Finish = NewRef(unsafe.Pointer(x.ref1aab3210.finish))
	x.InstanceDesc = *NewGodotPluginscriptInstanceDescRef(unsafe.Pointer(&x.ref1aab3210.instance_desc))
}

// allocGodotPluginscriptProfilingDataMemory allocates memory for type C.godot_pluginscript_profiling_data in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotPluginscriptProfilingDataMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotPluginscriptProfilingDataValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotPluginscriptProfilingDataValue = unsafe.Sizeof([1]C.godot_pluginscript_profiling_data{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotPluginscriptProfilingData) Ref() *C.godot_pluginscript_profiling_data {
	if x == nil {
		return nil
	}
	return x.ref9c004e5a
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotPluginscriptProfilingData) Free() {
	if x != nil && x.allocs9c004e5a != nil {
		x.allocs9c004e5a.(*cgoAllocMap).Free()
		x.ref9c004e5a = nil
	}
}

// NewGodotPluginscriptProfilingDataRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotPluginscriptProfilingDataRef(ref unsafe.Pointer) *GodotPluginscriptProfilingData {
	if ref == nil {
		return nil
	}
	obj := new(GodotPluginscriptProfilingData)
	obj.ref9c004e5a = (*C.godot_pluginscript_profiling_data)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotPluginscriptProfilingData) PassRef() (*C.godot_pluginscript_profiling_data, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref9c004e5a != nil {
		return x.ref9c004e5a, nil
	}
	mem9c004e5a := allocGodotPluginscriptProfilingDataMemory(1)
	ref9c004e5a := (*C.godot_pluginscript_profiling_data)(mem9c004e5a)
	allocs9c004e5a := new(cgoAllocMap)
	allocs9c004e5a.Add(mem9c004e5a)

	var csignature_allocs *cgoAllocMap
	ref9c004e5a.signature, csignature_allocs = x.Signature.PassValue()
	allocs9c004e5a.Borrow(csignature_allocs)

	var ccall_count_allocs *cgoAllocMap
	ref9c004e5a.call_count, ccall_count_allocs = (C.godot_int)(x.CallCount), cgoAllocsUnknown
	allocs9c004e5a.Borrow(ccall_count_allocs)

	var ctotal_time_allocs *cgoAllocMap
	ref9c004e5a.total_time, ctotal_time_allocs = (C.godot_int)(x.TotalTime), cgoAllocsUnknown
	allocs9c004e5a.Borrow(ctotal_time_allocs)

	var cself_time_allocs *cgoAllocMap
	ref9c004e5a.self_time, cself_time_allocs = (C.godot_int)(x.SelfTime), cgoAllocsUnknown
	allocs9c004e5a.Borrow(cself_time_allocs)

	x.ref9c004e5a = ref9c004e5a
	x.allocs9c004e5a = allocs9c004e5a
	return ref9c004e5a, allocs9c004e5a

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotPluginscriptProfilingData) PassValue() (C.godot_pluginscript_profiling_data, *cgoAllocMap) {
	if x.ref9c004e5a != nil {
		return *x.ref9c004e5a, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotPluginscriptProfilingData) Deref() {
	if x.ref9c004e5a == nil {
		return
	}
	x.Signature = *NewGodotStringNameRef(unsafe.Pointer(&x.ref9c004e5a.signature))
	x.CallCount = (GodotInt)(x.ref9c004e5a.call_count)
	x.TotalTime = (GodotInt)(x.ref9c004e5a.total_time)
	x.SelfTime = (GodotInt)(x.ref9c004e5a.self_time)
}

// allocGodotPluginscriptLanguageDescMemory allocates memory for type C.godot_pluginscript_language_desc in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotPluginscriptLanguageDescMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotPluginscriptLanguageDescValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotPluginscriptLanguageDescValue = unsafe.Sizeof([1]C.godot_pluginscript_language_desc{})

// unpackPCharString represents the data from Go string as *C.char and avoids copying.
func unpackPCharString(str string) (*C.char, *cgoAllocMap) {
	h := (*stringHeader)(unsafe.Pointer(&str))
	return (*C.char)(unsafe.Pointer(h.Data)), cgoAllocsUnknown
}

type stringHeader struct {
	Data uintptr
	Len  int
}

// allocPCharMemory allocates memory for type *C.char in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPCharMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPCharValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPCharValue = unsafe.Sizeof([1]*C.char{})

// unpackSString transforms a sliced Go data structure into plain C format.
func unpackSString(x []string) (unpacked **C.char, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.char) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPCharMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.char)(unsafe.Pointer(h0))
	for i0 := range x {
		v0[i0], _ = unpackPCharString(x[i0])
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.char)(unsafe.Pointer(h.Data))
	return
}

// packPCharString creates a Go string backed by *C.char and avoids copying.
func packPCharString(p *C.char) (raw string) {
	if p != nil && *p != 0 {
		h := (*stringHeader)(unsafe.Pointer(&raw))
		h.Data = uintptr(unsafe.Pointer(p))
		for *p != 0 {
			p = (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) // p++
		}
		h.Len = int(uintptr(unsafe.Pointer(p)) - h.Data)
	}
	return
}

// RawString reperesents a string backed by data on the C side.
type RawString string

// Copy returns a Go-managed copy of raw string.
func (raw RawString) Copy() string {
	if len(raw) == 0 {
		return ""
	}
	h := (*stringHeader)(unsafe.Pointer(&raw))
	return C.GoStringN((*C.char)(unsafe.Pointer(h.Data)), C.int(h.Len))
}

// packSString reads sliced Go data structure out from plain C format.
func packSString(v []string, ptr0 **C.char) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.char)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = packPCharString(ptr1)
	}
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GodotPluginscriptLanguageDesc) Ref() *C.godot_pluginscript_language_desc {
	if x == nil {
		return nil
	}
	return x.refdac22bbe
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GodotPluginscriptLanguageDesc) Free() {
	if x != nil && x.allocsdac22bbe != nil {
		x.allocsdac22bbe.(*cgoAllocMap).Free()
		x.refdac22bbe = nil
	}
}

// NewGodotPluginscriptLanguageDescRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGodotPluginscriptLanguageDescRef(ref unsafe.Pointer) *GodotPluginscriptLanguageDesc {
	if ref == nil {
		return nil
	}
	obj := new(GodotPluginscriptLanguageDesc)
	obj.refdac22bbe = (*C.godot_pluginscript_language_desc)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GodotPluginscriptLanguageDesc) PassRef() (*C.godot_pluginscript_language_desc, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refdac22bbe != nil {
		return x.refdac22bbe, nil
	}
	memdac22bbe := allocGodotPluginscriptLanguageDescMemory(1)
	refdac22bbe := (*C.godot_pluginscript_language_desc)(memdac22bbe)
	allocsdac22bbe := new(cgoAllocMap)
	allocsdac22bbe.Add(memdac22bbe)

	var cname_allocs *cgoAllocMap
	refdac22bbe.name, cname_allocs = unpackPCharString(x.Name)
	allocsdac22bbe.Borrow(cname_allocs)

	var c_type_allocs *cgoAllocMap
	refdac22bbe._type, c_type_allocs = unpackPCharString(x.Type)
	allocsdac22bbe.Borrow(c_type_allocs)

	var cextension_allocs *cgoAllocMap
	refdac22bbe.extension, cextension_allocs = unpackPCharString(x.Extension)
	allocsdac22bbe.Borrow(cextension_allocs)

	var crecognized_extensions_allocs *cgoAllocMap
	refdac22bbe.recognized_extensions, crecognized_extensions_allocs = unpackSString(x.RecognizedExtensions)
	allocsdac22bbe.Borrow(crecognized_extensions_allocs)

	var cinit_allocs *cgoAllocMap
	refdac22bbe.init, cinit_allocs = x.Init.PassRef()
	allocsdac22bbe.Borrow(cinit_allocs)

	var cfinish_allocs *cgoAllocMap
	refdac22bbe.finish, cfinish_allocs = x.Finish.PassRef()
	allocsdac22bbe.Borrow(cfinish_allocs)

	var creserved_words_allocs *cgoAllocMap
	refdac22bbe.reserved_words, creserved_words_allocs = unpackSString(x.ReservedWords)
	allocsdac22bbe.Borrow(creserved_words_allocs)

	var ccomment_delimiters_allocs *cgoAllocMap
	refdac22bbe.comment_delimiters, ccomment_delimiters_allocs = unpackSString(x.CommentDelimiters)
	allocsdac22bbe.Borrow(ccomment_delimiters_allocs)

	var cstring_delimiters_allocs *cgoAllocMap
	refdac22bbe.string_delimiters, cstring_delimiters_allocs = unpackSString(x.StringDelimiters)
	allocsdac22bbe.Borrow(cstring_delimiters_allocs)

	var chas_named_classes_allocs *cgoAllocMap
	refdac22bbe.has_named_classes, chas_named_classes_allocs = (C.godot_bool)(x.HasNamedClasses), cgoAllocsUnknown
	allocsdac22bbe.Borrow(chas_named_classes_allocs)

	var csupports_builtin_mode_allocs *cgoAllocMap
	refdac22bbe.supports_builtin_mode, csupports_builtin_mode_allocs = (C.godot_bool)(x.SupportsBuiltinMode), cgoAllocsUnknown
	allocsdac22bbe.Borrow(csupports_builtin_mode_allocs)

	var cget_template_source_code_allocs *cgoAllocMap
	refdac22bbe.get_template_source_code, cget_template_source_code_allocs = x.GetTemplateSourceCode.PassRef()
	allocsdac22bbe.Borrow(cget_template_source_code_allocs)

	var cvalidate_allocs *cgoAllocMap
	refdac22bbe.validate, cvalidate_allocs = x.Validate.PassRef()
	allocsdac22bbe.Borrow(cvalidate_allocs)

	var cfind_function_allocs *cgoAllocMap
	refdac22bbe.find_function, cfind_function_allocs = x.FindFunction.PassRef()
	allocsdac22bbe.Borrow(cfind_function_allocs)

	var cmake_function_allocs *cgoAllocMap
	refdac22bbe.make_function, cmake_function_allocs = x.MakeFunction.PassRef()
	allocsdac22bbe.Borrow(cmake_function_allocs)

	var ccomplete_code_allocs *cgoAllocMap
	refdac22bbe.complete_code, ccomplete_code_allocs = x.CompleteCode.PassRef()
	allocsdac22bbe.Borrow(ccomplete_code_allocs)

	var cauto_indent_code_allocs *cgoAllocMap
	refdac22bbe.auto_indent_code, cauto_indent_code_allocs = x.AutoIndentCode.PassRef()
	allocsdac22bbe.Borrow(cauto_indent_code_allocs)

	var cadd_global_constant_allocs *cgoAllocMap
	refdac22bbe.add_global_constant, cadd_global_constant_allocs = x.AddGlobalConstant.PassRef()
	allocsdac22bbe.Borrow(cadd_global_constant_allocs)

	var cdebug_get_error_allocs *cgoAllocMap
	refdac22bbe.debug_get_error, cdebug_get_error_allocs = x.DebugGetError.PassRef()
	allocsdac22bbe.Borrow(cdebug_get_error_allocs)

	var cdebug_get_stack_level_count_allocs *cgoAllocMap
	refdac22bbe.debug_get_stack_level_count, cdebug_get_stack_level_count_allocs = x.DebugGetStackLevelCount.PassRef()
	allocsdac22bbe.Borrow(cdebug_get_stack_level_count_allocs)

	var cdebug_get_stack_level_line_allocs *cgoAllocMap
	refdac22bbe.debug_get_stack_level_line, cdebug_get_stack_level_line_allocs = x.DebugGetStackLevelLine.PassRef()
	allocsdac22bbe.Borrow(cdebug_get_stack_level_line_allocs)

	var cdebug_get_stack_level_function_allocs *cgoAllocMap
	refdac22bbe.debug_get_stack_level_function, cdebug_get_stack_level_function_allocs = x.DebugGetStackLevelFunction.PassRef()
	allocsdac22bbe.Borrow(cdebug_get_stack_level_function_allocs)

	var cdebug_get_stack_level_source_allocs *cgoAllocMap
	refdac22bbe.debug_get_stack_level_source, cdebug_get_stack_level_source_allocs = x.DebugGetStackLevelSource.PassRef()
	allocsdac22bbe.Borrow(cdebug_get_stack_level_source_allocs)

	var cdebug_get_stack_level_locals_allocs *cgoAllocMap
	refdac22bbe.debug_get_stack_level_locals, cdebug_get_stack_level_locals_allocs = x.DebugGetStackLevelLocals.PassRef()
	allocsdac22bbe.Borrow(cdebug_get_stack_level_locals_allocs)

	var cdebug_get_stack_level_members_allocs *cgoAllocMap
	refdac22bbe.debug_get_stack_level_members, cdebug_get_stack_level_members_allocs = x.DebugGetStackLevelMembers.PassRef()
	allocsdac22bbe.Borrow(cdebug_get_stack_level_members_allocs)

	var cdebug_get_globals_allocs *cgoAllocMap
	refdac22bbe.debug_get_globals, cdebug_get_globals_allocs = x.DebugGetGlobals.PassRef()
	allocsdac22bbe.Borrow(cdebug_get_globals_allocs)

	var cdebug_parse_stack_level_expression_allocs *cgoAllocMap
	refdac22bbe.debug_parse_stack_level_expression, cdebug_parse_stack_level_expression_allocs = x.DebugParseStackLevelExpression.PassRef()
	allocsdac22bbe.Borrow(cdebug_parse_stack_level_expression_allocs)

	var cget_public_functions_allocs *cgoAllocMap
	refdac22bbe.get_public_functions, cget_public_functions_allocs = x.GetPublicFunctions.PassRef()
	allocsdac22bbe.Borrow(cget_public_functions_allocs)

	var cget_public_constants_allocs *cgoAllocMap
	refdac22bbe.get_public_constants, cget_public_constants_allocs = x.GetPublicConstants.PassRef()
	allocsdac22bbe.Borrow(cget_public_constants_allocs)

	var cprofiling_start_allocs *cgoAllocMap
	refdac22bbe.profiling_start, cprofiling_start_allocs = x.ProfilingStart.PassRef()
	allocsdac22bbe.Borrow(cprofiling_start_allocs)

	var cprofiling_stop_allocs *cgoAllocMap
	refdac22bbe.profiling_stop, cprofiling_stop_allocs = x.ProfilingStop.PassRef()
	allocsdac22bbe.Borrow(cprofiling_stop_allocs)

	var cprofiling_get_accumulated_data_allocs *cgoAllocMap
	refdac22bbe.profiling_get_accumulated_data, cprofiling_get_accumulated_data_allocs = x.ProfilingGetAccumulatedData.PassRef()
	allocsdac22bbe.Borrow(cprofiling_get_accumulated_data_allocs)

	var cprofiling_get_frame_data_allocs *cgoAllocMap
	refdac22bbe.profiling_get_frame_data, cprofiling_get_frame_data_allocs = x.ProfilingGetFrameData.PassRef()
	allocsdac22bbe.Borrow(cprofiling_get_frame_data_allocs)

	var cprofiling_frame_allocs *cgoAllocMap
	refdac22bbe.profiling_frame, cprofiling_frame_allocs = x.ProfilingFrame.PassRef()
	allocsdac22bbe.Borrow(cprofiling_frame_allocs)

	var cscript_desc_allocs *cgoAllocMap
	refdac22bbe.script_desc, cscript_desc_allocs = x.ScriptDesc.PassValue()
	allocsdac22bbe.Borrow(cscript_desc_allocs)

	x.refdac22bbe = refdac22bbe
	x.allocsdac22bbe = allocsdac22bbe
	return refdac22bbe, allocsdac22bbe

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GodotPluginscriptLanguageDesc) PassValue() (C.godot_pluginscript_language_desc, *cgoAllocMap) {
	if x.refdac22bbe != nil {
		return *x.refdac22bbe, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GodotPluginscriptLanguageDesc) Deref() {
	if x.refdac22bbe == nil {
		return
	}
	x.Name = packPCharString(x.refdac22bbe.name)
	x.Type = packPCharString(x.refdac22bbe._type)
	x.Extension = packPCharString(x.refdac22bbe.extension)
	packSString(x.RecognizedExtensions, x.refdac22bbe.recognized_extensions)
	x.Init = NewGodotPluginscriptLanguageDataRef(unsafe.Pointer(x.refdac22bbe.init))
	x.Finish = NewRef(unsafe.Pointer(x.refdac22bbe.finish))
	packSString(x.ReservedWords, x.refdac22bbe.reserved_words)
	packSString(x.CommentDelimiters, x.refdac22bbe.comment_delimiters)
	packSString(x.StringDelimiters, x.refdac22bbe.string_delimiters)
	x.HasNamedClasses = (GodotBool)(x.refdac22bbe.has_named_classes)
	x.SupportsBuiltinMode = (GodotBool)(x.refdac22bbe.supports_builtin_mode)
	x.GetTemplateSourceCode = NewGodotStringRef(unsafe.Pointer(x.refdac22bbe.get_template_source_code))
	x.Validate = NewGodotBoolRef(unsafe.Pointer(x.refdac22bbe.validate))
	x.FindFunction = NewRef(unsafe.Pointer(x.refdac22bbe.find_function))
	x.MakeFunction = NewGodotStringRef(unsafe.Pointer(x.refdac22bbe.make_function))
	x.CompleteCode = NewGodotErrorRef(unsafe.Pointer(x.refdac22bbe.complete_code))
	x.AutoIndentCode = NewRef(unsafe.Pointer(x.refdac22bbe.auto_indent_code))
	x.AddGlobalConstant = NewRef(unsafe.Pointer(x.refdac22bbe.add_global_constant))
	x.DebugGetError = NewGodotStringRef(unsafe.Pointer(x.refdac22bbe.debug_get_error))
	x.DebugGetStackLevelCount = NewRef(unsafe.Pointer(x.refdac22bbe.debug_get_stack_level_count))
	x.DebugGetStackLevelLine = NewRef(unsafe.Pointer(x.refdac22bbe.debug_get_stack_level_line))
	x.DebugGetStackLevelFunction = NewGodotStringRef(unsafe.Pointer(x.refdac22bbe.debug_get_stack_level_function))
	x.DebugGetStackLevelSource = NewGodotStringRef(unsafe.Pointer(x.refdac22bbe.debug_get_stack_level_source))
	x.DebugGetStackLevelLocals = NewRef(unsafe.Pointer(x.refdac22bbe.debug_get_stack_level_locals))
	x.DebugGetStackLevelMembers = NewRef(unsafe.Pointer(x.refdac22bbe.debug_get_stack_level_members))
	x.DebugGetGlobals = NewRef(unsafe.Pointer(x.refdac22bbe.debug_get_globals))
	x.DebugParseStackLevelExpression = NewGodotStringRef(unsafe.Pointer(x.refdac22bbe.debug_parse_stack_level_expression))
	x.GetPublicFunctions = NewRef(unsafe.Pointer(x.refdac22bbe.get_public_functions))
	x.GetPublicConstants = NewRef(unsafe.Pointer(x.refdac22bbe.get_public_constants))
	x.ProfilingStart = NewRef(unsafe.Pointer(x.refdac22bbe.profiling_start))
	x.ProfilingStop = NewRef(unsafe.Pointer(x.refdac22bbe.profiling_stop))
	x.ProfilingGetAccumulatedData = NewRef(unsafe.Pointer(x.refdac22bbe.profiling_get_accumulated_data))
	x.ProfilingGetFrameData = NewRef(unsafe.Pointer(x.refdac22bbe.profiling_get_frame_data))
	x.ProfilingFrame = NewRef(unsafe.Pointer(x.refdac22bbe.profiling_frame))
	x.ScriptDesc = *NewGodotPluginscriptScriptDescRef(unsafe.Pointer(&x.refdac22bbe.script_desc))
}

// unpackArgSGodotMethodBind transforms a sliced Go data structure into plain C format.
func unpackArgSGodotMethodBind(x []GodotMethodBind) (unpacked *C.godot_method_bind, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_method_bind) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotMethodBindMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_method_bind)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_method_bind)(unsafe.Pointer(h.Data))
	return
}

// packSGodotMethodBind reads sliced Go data structure out from plain C format.
func packSGodotMethodBind(v []GodotMethodBind, ptr0 *C.godot_method_bind) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotMethodBindValue]C.godot_method_bind)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotMethodBindRef(unsafe.Pointer(&ptr1))
	}
}

// allocPGodotVariantMemory allocates memory for type *C.godot_variant in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPGodotVariantMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPGodotVariantValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPGodotVariantValue = unsafe.Sizeof([1]*C.godot_variant{})

// unpackArgSSGodotVariant transforms a sliced Go data structure into plain C format.
func unpackArgSSGodotVariant(x [][]GodotVariant) (unpacked **C.godot_variant, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.godot_variant) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPGodotVariantMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.godot_variant)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocGodotVariantMemory(len1)
		allocs.Add(mem1)
		h1 := &sliceHeader{
			Data: uintptr(mem1),
			Cap:  len1,
			Len:  len1,
		}
		v1 := *(*[]C.godot_variant)(unsafe.Pointer(h1))
		for i1 := range x[i0] {
			allocs1 := new(cgoAllocMap)
			v1[i1], allocs1 = x[i0][i1].PassValue()
			allocs.Borrow(allocs1)
		}
		h := (*sliceHeader)(unsafe.Pointer(&v1))
		v0[i0] = (*C.godot_variant)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.godot_variant)(unsafe.Pointer(h.Data))
	return
}

// packSSGodotVariant reads sliced Go data structure out from plain C format.
func packSSGodotVariant(v [][]GodotVariant, ptr0 **C.godot_variant) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.godot_variant)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfGodotVariantValue]C.godot_variant)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *NewGodotVariantRef(unsafe.Pointer(&ptr2))
		}
	}
}

// unpackArgSGodotVariantCallError transforms a sliced Go data structure into plain C format.
func unpackArgSGodotVariantCallError(x []GodotVariantCallError) (unpacked *C.godot_variant_call_error, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_variant_call_error) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotVariantCallErrorMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_variant_call_error)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_variant_call_error)(unsafe.Pointer(h.Data))
	return
}

// packSGodotVariantCallError reads sliced Go data structure out from plain C format.
func packSGodotVariantCallError(v []GodotVariantCallError, ptr0 *C.godot_variant_call_error) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotVariantCallErrorValue]C.godot_variant_call_error)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotVariantCallErrorRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSGodotString transforms a sliced Go data structure into plain C format.
func unpackArgSGodotString(x []GodotString) (unpacked *C.godot_string, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_string) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotStringMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_string)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_string)(unsafe.Pointer(h.Data))
	return
}

// unpackArgSGodotArray transforms a sliced Go data structure into plain C format.
func unpackArgSGodotArray(x []GodotArray) (unpacked *C.godot_array, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_array) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotArrayMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_array)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_array)(unsafe.Pointer(h.Data))
	return
}

// unpackArgSGodotVariant transforms a sliced Go data structure into plain C format.
func unpackArgSGodotVariant(x []GodotVariant) (unpacked *C.godot_variant, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_variant) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotVariantMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_variant)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_variant)(unsafe.Pointer(h.Data))
	return
}

// unpackPUint8TString represents the data from Go string as *C.uint8_t and avoids copying.
func unpackPUint8TString(str string) (*C.uint8_t, *cgoAllocMap) {
	h := (*stringHeader)(unsafe.Pointer(&str))
	return (*C.uint8_t)(unsafe.Pointer(h.Data)), cgoAllocsUnknown
}

// allocPWcharTMemory allocates memory for type *C.wchar_t in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPWcharTMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPWcharTValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPWcharTValue = unsafe.Sizeof([1]*C.wchar_t{})

// unpackArgSSInt32 transforms a sliced Go data structure into plain C format.
func unpackArgSSInt32(x [][]int32) (unpacked **C.wchar_t, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.wchar_t) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPWcharTMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.wchar_t)(unsafe.Pointer(h0))
	for i0 := range x {
		h := (*sliceHeader)(unsafe.Pointer(&x[i0]))
		v0[i0] = (*C.wchar_t)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.wchar_t)(unsafe.Pointer(h.Data))
	return
}

// packSSInt32 reads sliced Go data structure out from plain C format.
func packSSInt32(v [][]int32, ptr0 **C.wchar_t) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.wchar_t)(unsafe.Pointer(ptr0)))[i0]
		hxfc4425b := (*sliceHeader)(unsafe.Pointer(&v[i0]))
		hxfc4425b.Data = uintptr(unsafe.Pointer(ptr1))
		hxfc4425b.Cap = 0x7fffffff
		// hxfc4425b.Len = ?
	}
}

// unpackArgSGodotPoolColorArray transforms a sliced Go data structure into plain C format.
func unpackArgSGodotPoolColorArray(x []GodotPoolColorArray) (unpacked *C.godot_pool_color_array, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_color_array) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotPoolColorArrayMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_pool_color_array)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_pool_color_array)(unsafe.Pointer(h.Data))
	return
}

// packSGodotPoolColorArray reads sliced Go data structure out from plain C format.
func packSGodotPoolColorArray(v []GodotPoolColorArray, ptr0 *C.godot_pool_color_array) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotPoolColorArrayValue]C.godot_pool_color_array)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotPoolColorArrayRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSGodotPoolVector3Array transforms a sliced Go data structure into plain C format.
func unpackArgSGodotPoolVector3Array(x []GodotPoolVector3Array) (unpacked *C.godot_pool_vector3_array, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_vector3_array) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotPoolVector3ArrayMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_pool_vector3_array)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_pool_vector3_array)(unsafe.Pointer(h.Data))
	return
}

// packSGodotPoolVector3Array reads sliced Go data structure out from plain C format.
func packSGodotPoolVector3Array(v []GodotPoolVector3Array, ptr0 *C.godot_pool_vector3_array) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotPoolVector3ArrayValue]C.godot_pool_vector3_array)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotPoolVector3ArrayRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSGodotPoolVector2Array transforms a sliced Go data structure into plain C format.
func unpackArgSGodotPoolVector2Array(x []GodotPoolVector2Array) (unpacked *C.godot_pool_vector2_array, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_vector2_array) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotPoolVector2ArrayMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_pool_vector2_array)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_pool_vector2_array)(unsafe.Pointer(h.Data))
	return
}

// packSGodotPoolVector2Array reads sliced Go data structure out from plain C format.
func packSGodotPoolVector2Array(v []GodotPoolVector2Array, ptr0 *C.godot_pool_vector2_array) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotPoolVector2ArrayValue]C.godot_pool_vector2_array)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotPoolVector2ArrayRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSGodotPoolStringArray transforms a sliced Go data structure into plain C format.
func unpackArgSGodotPoolStringArray(x []GodotPoolStringArray) (unpacked *C.godot_pool_string_array, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_string_array) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotPoolStringArrayMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_pool_string_array)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_pool_string_array)(unsafe.Pointer(h.Data))
	return
}

// packSGodotPoolStringArray reads sliced Go data structure out from plain C format.
func packSGodotPoolStringArray(v []GodotPoolStringArray, ptr0 *C.godot_pool_string_array) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotPoolStringArrayValue]C.godot_pool_string_array)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotPoolStringArrayRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSGodotPoolRealArray transforms a sliced Go data structure into plain C format.
func unpackArgSGodotPoolRealArray(x []GodotPoolRealArray) (unpacked *C.godot_pool_real_array, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_real_array) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotPoolRealArrayMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_pool_real_array)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_pool_real_array)(unsafe.Pointer(h.Data))
	return
}

// packSGodotPoolRealArray reads sliced Go data structure out from plain C format.
func packSGodotPoolRealArray(v []GodotPoolRealArray, ptr0 *C.godot_pool_real_array) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotPoolRealArrayValue]C.godot_pool_real_array)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotPoolRealArrayRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSGodotPoolIntArray transforms a sliced Go data structure into plain C format.
func unpackArgSGodotPoolIntArray(x []GodotPoolIntArray) (unpacked *C.godot_pool_int_array, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_int_array) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotPoolIntArrayMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_pool_int_array)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_pool_int_array)(unsafe.Pointer(h.Data))
	return
}

// packSGodotPoolIntArray reads sliced Go data structure out from plain C format.
func packSGodotPoolIntArray(v []GodotPoolIntArray, ptr0 *C.godot_pool_int_array) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotPoolIntArrayValue]C.godot_pool_int_array)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotPoolIntArrayRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSGodotPoolByteArray transforms a sliced Go data structure into plain C format.
func unpackArgSGodotPoolByteArray(x []GodotPoolByteArray) (unpacked *C.godot_pool_byte_array, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_byte_array) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotPoolByteArrayMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_pool_byte_array)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_pool_byte_array)(unsafe.Pointer(h.Data))
	return
}

// packSGodotPoolByteArray reads sliced Go data structure out from plain C format.
func packSGodotPoolByteArray(v []GodotPoolByteArray, ptr0 *C.godot_pool_byte_array) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotPoolByteArrayValue]C.godot_pool_byte_array)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotPoolByteArrayRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSGodotVector2 transforms a sliced Go data structure into plain C format.
func unpackArgSGodotVector2(x []GodotVector2) (unpacked *C.godot_vector2, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_vector2) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotVector2Memory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_vector2)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_vector2)(unsafe.Pointer(h.Data))
	return
}

// packSGodotVector2 reads sliced Go data structure out from plain C format.
func packSGodotVector2(v []GodotVector2, ptr0 *C.godot_vector2) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotVector2Value]C.godot_vector2)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotVector2Ref(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSGodotVector3 transforms a sliced Go data structure into plain C format.
func unpackArgSGodotVector3(x []GodotVector3) (unpacked *C.godot_vector3, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_vector3) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotVector3Memory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_vector3)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_vector3)(unsafe.Pointer(h.Data))
	return
}

// packSGodotVector3 reads sliced Go data structure out from plain C format.
func packSGodotVector3(v []GodotVector3, ptr0 *C.godot_vector3) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotVector3Value]C.godot_vector3)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotVector3Ref(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSGodotColor transforms a sliced Go data structure into plain C format.
func unpackArgSGodotColor(x []GodotColor) (unpacked *C.godot_color, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_color) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotColorMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_color)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_color)(unsafe.Pointer(h.Data))
	return
}

// packSGodotColor reads sliced Go data structure out from plain C format.
func packSGodotColor(v []GodotColor, ptr0 *C.godot_color) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotColorValue]C.godot_color)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotColorRef(unsafe.Pointer(&ptr1))
	}
}

// allocGodotPoolByteArrayReadAccessMemory allocates memory for type C.godot_pool_byte_array_read_access in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotPoolByteArrayReadAccessMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotPoolByteArrayReadAccessValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotPoolByteArrayReadAccessValue = unsafe.Sizeof([1]C.godot_pool_byte_array_read_access{})

// unpackArgSGodotPoolByteArrayReadAccess transforms a sliced Go data structure into plain C format.
func unpackArgSGodotPoolByteArrayReadAccess(x []GodotPoolByteArrayReadAccess) (unpacked *C.godot_pool_byte_array_read_access, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_byte_array_read_access) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotPoolByteArrayReadAccessMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_pool_byte_array_read_access)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_pool_byte_array_read_access)(unsafe.Pointer(h.Data))
	return
}

// packSGodotPoolByteArrayReadAccess reads sliced Go data structure out from plain C format.
func packSGodotPoolByteArrayReadAccess(v []GodotPoolByteArrayReadAccess, ptr0 *C.godot_pool_byte_array_read_access) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotPoolByteArrayReadAccessValue]C.godot_pool_byte_array_read_access)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotPoolByteArrayReadAccessRef(unsafe.Pointer(&ptr1))
	}
}

// packPUint8TString creates a Go string backed by *C.uint8_t and avoids copying.
func packPUint8TString(p *C.uint8_t) (raw string) {
	if p != nil && *p != 0 {
		h := (*stringHeader)(unsafe.Pointer(&raw))
		h.Data = uintptr(unsafe.Pointer(p))
		for *p != 0 {
			p = (*C.uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) // p++
		}
		h.Len = int(uintptr(unsafe.Pointer(p)) - h.Data)
	}
	return
}

// allocGodotPoolIntArrayReadAccessMemory allocates memory for type C.godot_pool_int_array_read_access in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotPoolIntArrayReadAccessMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotPoolIntArrayReadAccessValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotPoolIntArrayReadAccessValue = unsafe.Sizeof([1]C.godot_pool_int_array_read_access{})

// unpackArgSGodotPoolIntArrayReadAccess transforms a sliced Go data structure into plain C format.
func unpackArgSGodotPoolIntArrayReadAccess(x []GodotPoolIntArrayReadAccess) (unpacked *C.godot_pool_int_array_read_access, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_int_array_read_access) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotPoolIntArrayReadAccessMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_pool_int_array_read_access)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_pool_int_array_read_access)(unsafe.Pointer(h.Data))
	return
}

// packSGodotPoolIntArrayReadAccess reads sliced Go data structure out from plain C format.
func packSGodotPoolIntArrayReadAccess(v []GodotPoolIntArrayReadAccess, ptr0 *C.godot_pool_int_array_read_access) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotPoolIntArrayReadAccessValue]C.godot_pool_int_array_read_access)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotPoolIntArrayReadAccessRef(unsafe.Pointer(&ptr1))
	}
}

// allocGodotPoolRealArrayReadAccessMemory allocates memory for type C.godot_pool_real_array_read_access in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotPoolRealArrayReadAccessMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotPoolRealArrayReadAccessValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotPoolRealArrayReadAccessValue = unsafe.Sizeof([1]C.godot_pool_real_array_read_access{})

// unpackArgSGodotPoolRealArrayReadAccess transforms a sliced Go data structure into plain C format.
func unpackArgSGodotPoolRealArrayReadAccess(x []GodotPoolRealArrayReadAccess) (unpacked *C.godot_pool_real_array_read_access, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_real_array_read_access) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotPoolRealArrayReadAccessMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_pool_real_array_read_access)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_pool_real_array_read_access)(unsafe.Pointer(h.Data))
	return
}

// packSGodotPoolRealArrayReadAccess reads sliced Go data structure out from plain C format.
func packSGodotPoolRealArrayReadAccess(v []GodotPoolRealArrayReadAccess, ptr0 *C.godot_pool_real_array_read_access) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotPoolRealArrayReadAccessValue]C.godot_pool_real_array_read_access)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotPoolRealArrayReadAccessRef(unsafe.Pointer(&ptr1))
	}
}

// allocGodotPoolStringArrayReadAccessMemory allocates memory for type C.godot_pool_string_array_read_access in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotPoolStringArrayReadAccessMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotPoolStringArrayReadAccessValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotPoolStringArrayReadAccessValue = unsafe.Sizeof([1]C.godot_pool_string_array_read_access{})

// unpackArgSGodotPoolStringArrayReadAccess transforms a sliced Go data structure into plain C format.
func unpackArgSGodotPoolStringArrayReadAccess(x []GodotPoolStringArrayReadAccess) (unpacked *C.godot_pool_string_array_read_access, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_string_array_read_access) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotPoolStringArrayReadAccessMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_pool_string_array_read_access)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_pool_string_array_read_access)(unsafe.Pointer(h.Data))
	return
}

// packSGodotPoolStringArrayReadAccess reads sliced Go data structure out from plain C format.
func packSGodotPoolStringArrayReadAccess(v []GodotPoolStringArrayReadAccess, ptr0 *C.godot_pool_string_array_read_access) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotPoolStringArrayReadAccessValue]C.godot_pool_string_array_read_access)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotPoolStringArrayReadAccessRef(unsafe.Pointer(&ptr1))
	}
}

// allocGodotPoolVector2ArrayReadAccessMemory allocates memory for type C.godot_pool_vector2_array_read_access in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotPoolVector2ArrayReadAccessMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotPoolVector2ArrayReadAccessValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotPoolVector2ArrayReadAccessValue = unsafe.Sizeof([1]C.godot_pool_vector2_array_read_access{})

// unpackArgSGodotPoolVector2ArrayReadAccess transforms a sliced Go data structure into plain C format.
func unpackArgSGodotPoolVector2ArrayReadAccess(x []GodotPoolVector2ArrayReadAccess) (unpacked *C.godot_pool_vector2_array_read_access, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_vector2_array_read_access) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotPoolVector2ArrayReadAccessMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_pool_vector2_array_read_access)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_pool_vector2_array_read_access)(unsafe.Pointer(h.Data))
	return
}

// packSGodotPoolVector2ArrayReadAccess reads sliced Go data structure out from plain C format.
func packSGodotPoolVector2ArrayReadAccess(v []GodotPoolVector2ArrayReadAccess, ptr0 *C.godot_pool_vector2_array_read_access) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotPoolVector2ArrayReadAccessValue]C.godot_pool_vector2_array_read_access)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotPoolVector2ArrayReadAccessRef(unsafe.Pointer(&ptr1))
	}
}

// allocGodotPoolVector3ArrayReadAccessMemory allocates memory for type C.godot_pool_vector3_array_read_access in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotPoolVector3ArrayReadAccessMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotPoolVector3ArrayReadAccessValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotPoolVector3ArrayReadAccessValue = unsafe.Sizeof([1]C.godot_pool_vector3_array_read_access{})

// unpackArgSGodotPoolVector3ArrayReadAccess transforms a sliced Go data structure into plain C format.
func unpackArgSGodotPoolVector3ArrayReadAccess(x []GodotPoolVector3ArrayReadAccess) (unpacked *C.godot_pool_vector3_array_read_access, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_vector3_array_read_access) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotPoolVector3ArrayReadAccessMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_pool_vector3_array_read_access)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_pool_vector3_array_read_access)(unsafe.Pointer(h.Data))
	return
}

// packSGodotPoolVector3ArrayReadAccess reads sliced Go data structure out from plain C format.
func packSGodotPoolVector3ArrayReadAccess(v []GodotPoolVector3ArrayReadAccess, ptr0 *C.godot_pool_vector3_array_read_access) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotPoolVector3ArrayReadAccessValue]C.godot_pool_vector3_array_read_access)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotPoolVector3ArrayReadAccessRef(unsafe.Pointer(&ptr1))
	}
}

// allocGodotPoolColorArrayReadAccessMemory allocates memory for type C.godot_pool_color_array_read_access in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotPoolColorArrayReadAccessMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotPoolColorArrayReadAccessValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotPoolColorArrayReadAccessValue = unsafe.Sizeof([1]C.godot_pool_color_array_read_access{})

// unpackArgSGodotPoolColorArrayReadAccess transforms a sliced Go data structure into plain C format.
func unpackArgSGodotPoolColorArrayReadAccess(x []GodotPoolColorArrayReadAccess) (unpacked *C.godot_pool_color_array_read_access, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_color_array_read_access) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotPoolColorArrayReadAccessMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_pool_color_array_read_access)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_pool_color_array_read_access)(unsafe.Pointer(h.Data))
	return
}

// packSGodotPoolColorArrayReadAccess reads sliced Go data structure out from plain C format.
func packSGodotPoolColorArrayReadAccess(v []GodotPoolColorArrayReadAccess, ptr0 *C.godot_pool_color_array_read_access) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotPoolColorArrayReadAccessValue]C.godot_pool_color_array_read_access)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotPoolColorArrayReadAccessRef(unsafe.Pointer(&ptr1))
	}
}

// allocGodotPoolByteArrayWriteAccessMemory allocates memory for type C.godot_pool_byte_array_write_access in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotPoolByteArrayWriteAccessMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotPoolByteArrayWriteAccessValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotPoolByteArrayWriteAccessValue = unsafe.Sizeof([1]C.godot_pool_byte_array_write_access{})

// unpackArgSGodotPoolByteArrayWriteAccess transforms a sliced Go data structure into plain C format.
func unpackArgSGodotPoolByteArrayWriteAccess(x []GodotPoolByteArrayWriteAccess) (unpacked *C.godot_pool_byte_array_write_access, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_byte_array_write_access) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotPoolByteArrayWriteAccessMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_pool_byte_array_write_access)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_pool_byte_array_write_access)(unsafe.Pointer(h.Data))
	return
}

// packSGodotPoolByteArrayWriteAccess reads sliced Go data structure out from plain C format.
func packSGodotPoolByteArrayWriteAccess(v []GodotPoolByteArrayWriteAccess, ptr0 *C.godot_pool_byte_array_write_access) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotPoolByteArrayWriteAccessValue]C.godot_pool_byte_array_write_access)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotPoolByteArrayWriteAccessRef(unsafe.Pointer(&ptr1))
	}
}

// allocGodotPoolIntArrayWriteAccessMemory allocates memory for type C.godot_pool_int_array_write_access in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotPoolIntArrayWriteAccessMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotPoolIntArrayWriteAccessValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotPoolIntArrayWriteAccessValue = unsafe.Sizeof([1]C.godot_pool_int_array_write_access{})

// unpackArgSGodotPoolIntArrayWriteAccess transforms a sliced Go data structure into plain C format.
func unpackArgSGodotPoolIntArrayWriteAccess(x []GodotPoolIntArrayWriteAccess) (unpacked *C.godot_pool_int_array_write_access, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_int_array_write_access) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotPoolIntArrayWriteAccessMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_pool_int_array_write_access)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_pool_int_array_write_access)(unsafe.Pointer(h.Data))
	return
}

// packSGodotPoolIntArrayWriteAccess reads sliced Go data structure out from plain C format.
func packSGodotPoolIntArrayWriteAccess(v []GodotPoolIntArrayWriteAccess, ptr0 *C.godot_pool_int_array_write_access) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotPoolIntArrayWriteAccessValue]C.godot_pool_int_array_write_access)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotPoolIntArrayWriteAccessRef(unsafe.Pointer(&ptr1))
	}
}

// allocGodotPoolRealArrayWriteAccessMemory allocates memory for type C.godot_pool_real_array_write_access in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotPoolRealArrayWriteAccessMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotPoolRealArrayWriteAccessValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotPoolRealArrayWriteAccessValue = unsafe.Sizeof([1]C.godot_pool_real_array_write_access{})

// unpackArgSGodotPoolRealArrayWriteAccess transforms a sliced Go data structure into plain C format.
func unpackArgSGodotPoolRealArrayWriteAccess(x []GodotPoolRealArrayWriteAccess) (unpacked *C.godot_pool_real_array_write_access, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_real_array_write_access) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotPoolRealArrayWriteAccessMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_pool_real_array_write_access)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_pool_real_array_write_access)(unsafe.Pointer(h.Data))
	return
}

// packSGodotPoolRealArrayWriteAccess reads sliced Go data structure out from plain C format.
func packSGodotPoolRealArrayWriteAccess(v []GodotPoolRealArrayWriteAccess, ptr0 *C.godot_pool_real_array_write_access) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotPoolRealArrayWriteAccessValue]C.godot_pool_real_array_write_access)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotPoolRealArrayWriteAccessRef(unsafe.Pointer(&ptr1))
	}
}

// allocGodotPoolStringArrayWriteAccessMemory allocates memory for type C.godot_pool_string_array_write_access in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotPoolStringArrayWriteAccessMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotPoolStringArrayWriteAccessValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotPoolStringArrayWriteAccessValue = unsafe.Sizeof([1]C.godot_pool_string_array_write_access{})

// unpackArgSGodotPoolStringArrayWriteAccess transforms a sliced Go data structure into plain C format.
func unpackArgSGodotPoolStringArrayWriteAccess(x []GodotPoolStringArrayWriteAccess) (unpacked *C.godot_pool_string_array_write_access, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_string_array_write_access) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotPoolStringArrayWriteAccessMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_pool_string_array_write_access)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_pool_string_array_write_access)(unsafe.Pointer(h.Data))
	return
}

// packSGodotPoolStringArrayWriteAccess reads sliced Go data structure out from plain C format.
func packSGodotPoolStringArrayWriteAccess(v []GodotPoolStringArrayWriteAccess, ptr0 *C.godot_pool_string_array_write_access) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotPoolStringArrayWriteAccessValue]C.godot_pool_string_array_write_access)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotPoolStringArrayWriteAccessRef(unsafe.Pointer(&ptr1))
	}
}

// allocGodotPoolVector2ArrayWriteAccessMemory allocates memory for type C.godot_pool_vector2_array_write_access in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotPoolVector2ArrayWriteAccessMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotPoolVector2ArrayWriteAccessValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotPoolVector2ArrayWriteAccessValue = unsafe.Sizeof([1]C.godot_pool_vector2_array_write_access{})

// unpackArgSGodotPoolVector2ArrayWriteAccess transforms a sliced Go data structure into plain C format.
func unpackArgSGodotPoolVector2ArrayWriteAccess(x []GodotPoolVector2ArrayWriteAccess) (unpacked *C.godot_pool_vector2_array_write_access, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_vector2_array_write_access) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotPoolVector2ArrayWriteAccessMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_pool_vector2_array_write_access)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_pool_vector2_array_write_access)(unsafe.Pointer(h.Data))
	return
}

// packSGodotPoolVector2ArrayWriteAccess reads sliced Go data structure out from plain C format.
func packSGodotPoolVector2ArrayWriteAccess(v []GodotPoolVector2ArrayWriteAccess, ptr0 *C.godot_pool_vector2_array_write_access) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotPoolVector2ArrayWriteAccessValue]C.godot_pool_vector2_array_write_access)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotPoolVector2ArrayWriteAccessRef(unsafe.Pointer(&ptr1))
	}
}

// allocGodotPoolVector3ArrayWriteAccessMemory allocates memory for type C.godot_pool_vector3_array_write_access in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotPoolVector3ArrayWriteAccessMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotPoolVector3ArrayWriteAccessValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotPoolVector3ArrayWriteAccessValue = unsafe.Sizeof([1]C.godot_pool_vector3_array_write_access{})

// unpackArgSGodotPoolVector3ArrayWriteAccess transforms a sliced Go data structure into plain C format.
func unpackArgSGodotPoolVector3ArrayWriteAccess(x []GodotPoolVector3ArrayWriteAccess) (unpacked *C.godot_pool_vector3_array_write_access, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_vector3_array_write_access) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotPoolVector3ArrayWriteAccessMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_pool_vector3_array_write_access)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_pool_vector3_array_write_access)(unsafe.Pointer(h.Data))
	return
}

// packSGodotPoolVector3ArrayWriteAccess reads sliced Go data structure out from plain C format.
func packSGodotPoolVector3ArrayWriteAccess(v []GodotPoolVector3ArrayWriteAccess, ptr0 *C.godot_pool_vector3_array_write_access) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotPoolVector3ArrayWriteAccessValue]C.godot_pool_vector3_array_write_access)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotPoolVector3ArrayWriteAccessRef(unsafe.Pointer(&ptr1))
	}
}

// allocGodotPoolColorArrayWriteAccessMemory allocates memory for type C.godot_pool_color_array_write_access in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGodotPoolColorArrayWriteAccessMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGodotPoolColorArrayWriteAccessValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGodotPoolColorArrayWriteAccessValue = unsafe.Sizeof([1]C.godot_pool_color_array_write_access{})

// unpackArgSGodotPoolColorArrayWriteAccess transforms a sliced Go data structure into plain C format.
func unpackArgSGodotPoolColorArrayWriteAccess(x []GodotPoolColorArrayWriteAccess) (unpacked *C.godot_pool_color_array_write_access, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_color_array_write_access) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotPoolColorArrayWriteAccessMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_pool_color_array_write_access)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_pool_color_array_write_access)(unsafe.Pointer(h.Data))
	return
}

// packSGodotPoolColorArrayWriteAccess reads sliced Go data structure out from plain C format.
func packSGodotPoolColorArrayWriteAccess(v []GodotPoolColorArrayWriteAccess, ptr0 *C.godot_pool_color_array_write_access) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotPoolColorArrayWriteAccessValue]C.godot_pool_color_array_write_access)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotPoolColorArrayWriteAccessRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSGodotBasis transforms a sliced Go data structure into plain C format.
func unpackArgSGodotBasis(x []GodotBasis) (unpacked *C.godot_basis, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_basis) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotBasisMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_basis)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_basis)(unsafe.Pointer(h.Data))
	return
}

// packSGodotBasis reads sliced Go data structure out from plain C format.
func packSGodotBasis(v []GodotBasis, ptr0 *C.godot_basis) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotBasisValue]C.godot_basis)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotBasisRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSGodotQuat transforms a sliced Go data structure into plain C format.
func unpackArgSGodotQuat(x []GodotQuat) (unpacked *C.godot_quat, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_quat) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotQuatMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_quat)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_quat)(unsafe.Pointer(h.Data))
	return
}

// packSGodotQuat reads sliced Go data structure out from plain C format.
func packSGodotQuat(v []GodotQuat, ptr0 *C.godot_quat) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotQuatValue]C.godot_quat)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotQuatRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSGodotRect2 transforms a sliced Go data structure into plain C format.
func unpackArgSGodotRect2(x []GodotRect2) (unpacked *C.godot_rect2, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_rect2) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotRect2Memory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_rect2)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_rect2)(unsafe.Pointer(h.Data))
	return
}

// packSGodotRect2 reads sliced Go data structure out from plain C format.
func packSGodotRect2(v []GodotRect2, ptr0 *C.godot_rect2) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotRect2Value]C.godot_rect2)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotRect2Ref(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSGodotTransform2d transforms a sliced Go data structure into plain C format.
func unpackArgSGodotTransform2d(x []GodotTransform2d) (unpacked *C.godot_transform2d, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_transform2d) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotTransform2dMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_transform2d)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_transform2d)(unsafe.Pointer(h.Data))
	return
}

// packSGodotTransform2d reads sliced Go data structure out from plain C format.
func packSGodotTransform2d(v []GodotTransform2d, ptr0 *C.godot_transform2d) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotTransform2dValue]C.godot_transform2d)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotTransform2dRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSGodotPlane transforms a sliced Go data structure into plain C format.
func unpackArgSGodotPlane(x []GodotPlane) (unpacked *C.godot_plane, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_plane) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotPlaneMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_plane)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_plane)(unsafe.Pointer(h.Data))
	return
}

// packSGodotPlane reads sliced Go data structure out from plain C format.
func packSGodotPlane(v []GodotPlane, ptr0 *C.godot_plane) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotPlaneValue]C.godot_plane)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotPlaneRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSGodotAabb transforms a sliced Go data structure into plain C format.
func unpackArgSGodotAabb(x []GodotAabb) (unpacked *C.godot_aabb, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_aabb) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotAabbMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_aabb)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_aabb)(unsafe.Pointer(h.Data))
	return
}

// packSGodotAabb reads sliced Go data structure out from plain C format.
func packSGodotAabb(v []GodotAabb, ptr0 *C.godot_aabb) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotAabbValue]C.godot_aabb)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotAabbRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSGodotTransform transforms a sliced Go data structure into plain C format.
func unpackArgSGodotTransform(x []GodotTransform) (unpacked *C.godot_transform, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_transform) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotTransformMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_transform)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_transform)(unsafe.Pointer(h.Data))
	return
}

// packSGodotTransform reads sliced Go data structure out from plain C format.
func packSGodotTransform(v []GodotTransform, ptr0 *C.godot_transform) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotTransformValue]C.godot_transform)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotTransformRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSGodotNodePath transforms a sliced Go data structure into plain C format.
func unpackArgSGodotNodePath(x []GodotNodePath) (unpacked *C.godot_node_path, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_node_path) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotNodePathMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_node_path)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_node_path)(unsafe.Pointer(h.Data))
	return
}

// packSGodotNodePath reads sliced Go data structure out from plain C format.
func packSGodotNodePath(v []GodotNodePath, ptr0 *C.godot_node_path) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotNodePathValue]C.godot_node_path)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotNodePathRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSGodotRid transforms a sliced Go data structure into plain C format.
func unpackArgSGodotRid(x []GodotRid) (unpacked *C.godot_rid, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_rid) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotRidMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_rid)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_rid)(unsafe.Pointer(h.Data))
	return
}

// packSGodotRid reads sliced Go data structure out from plain C format.
func packSGodotRid(v []GodotRid, ptr0 *C.godot_rid) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotRidValue]C.godot_rid)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotRidRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSGodotDictionary transforms a sliced Go data structure into plain C format.
func unpackArgSGodotDictionary(x []GodotDictionary) (unpacked *C.godot_dictionary, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_dictionary) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotDictionaryMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_dictionary)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_dictionary)(unsafe.Pointer(h.Data))
	return
}

// packSGodotDictionary reads sliced Go data structure out from plain C format.
func packSGodotDictionary(v []GodotDictionary, ptr0 *C.godot_dictionary) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotDictionaryValue]C.godot_dictionary)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotDictionaryRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSGodotStringName transforms a sliced Go data structure into plain C format.
func unpackArgSGodotStringName(x []GodotStringName) (unpacked *C.godot_string_name, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_string_name) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotStringNameMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_string_name)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_string_name)(unsafe.Pointer(h.Data))
	return
}

// packSGodotStringName reads sliced Go data structure out from plain C format.
func packSGodotStringName(v []GodotStringName, ptr0 *C.godot_string_name) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotStringNameValue]C.godot_string_name)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotStringNameRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSGodotArvrInterfaceGdnative transforms a sliced Go data structure into plain C format.
func unpackArgSGodotArvrInterfaceGdnative(x []GodotArvrInterfaceGdnative) (unpacked *C.godot_arvr_interface_gdnative, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_arvr_interface_gdnative) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotArvrInterfaceGdnativeMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_arvr_interface_gdnative)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_arvr_interface_gdnative)(unsafe.Pointer(h.Data))
	return
}

// packSGodotArvrInterfaceGdnative reads sliced Go data structure out from plain C format.
func packSGodotArvrInterfaceGdnative(v []GodotArvrInterfaceGdnative, ptr0 *C.godot_arvr_interface_gdnative) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotArvrInterfaceGdnativeValue]C.godot_arvr_interface_gdnative)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotArvrInterfaceGdnativeRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSGodotPropertyAttributes transforms a sliced Go data structure into plain C format.
func unpackArgSGodotPropertyAttributes(x []GodotPropertyAttributes) (unpacked *C.godot_property_attributes, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_property_attributes) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotPropertyAttributesMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_property_attributes)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_property_attributes)(unsafe.Pointer(h.Data))
	return
}

// packSGodotPropertyAttributes reads sliced Go data structure out from plain C format.
func packSGodotPropertyAttributes(v []GodotPropertyAttributes, ptr0 *C.godot_property_attributes) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotPropertyAttributesValue]C.godot_property_attributes)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotPropertyAttributesRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSGodotSignal transforms a sliced Go data structure into plain C format.
func unpackArgSGodotSignal(x []GodotSignal) (unpacked *C.godot_signal, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_signal) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotSignalMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_signal)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_signal)(unsafe.Pointer(h.Data))
	return
}

// packSGodotSignal reads sliced Go data structure out from plain C format.
func packSGodotSignal(v []GodotSignal, ptr0 *C.godot_signal) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotSignalValue]C.godot_signal)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotSignalRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSGodotPluginscriptLanguageDesc transforms a sliced Go data structure into plain C format.
func unpackArgSGodotPluginscriptLanguageDesc(x []GodotPluginscriptLanguageDesc) (unpacked *C.godot_pluginscript_language_desc, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pluginscript_language_desc) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGodotPluginscriptLanguageDescMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]C.godot_pluginscript_language_desc)(unsafe.Pointer(h0))
	for i0 := range x {
		allocs0 := new(cgoAllocMap)
		v0[i0], allocs0 = x[i0].PassValue()
		allocs.Borrow(allocs0)
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (*C.godot_pluginscript_language_desc)(unsafe.Pointer(h.Data))
	return
}

// packSGodotPluginscriptLanguageDesc reads sliced Go data structure out from plain C format.
func packSGodotPluginscriptLanguageDesc(v []GodotPluginscriptLanguageDesc, ptr0 *C.godot_pluginscript_language_desc) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGodotPluginscriptLanguageDescValue]C.godot_pluginscript_language_desc)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGodotPluginscriptLanguageDescRef(unsafe.Pointer(&ptr1))
	}
}
