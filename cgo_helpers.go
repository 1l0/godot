// WARNING: This file has automatically been generated on Sun, 24 Dec 2017 07:11:06 JST.
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

// allocGdnativeExtNativescriptApiStructMemory allocates memory for type C.godot_gdnative_ext_nativescript_api_struct in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGdnativeExtNativescriptApiStructMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGdnativeExtNativescriptApiStructValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGdnativeExtNativescriptApiStructValue = unsafe.Sizeof([1]C.godot_gdnative_ext_nativescript_api_struct{})

type sliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

// allocGdnativeApiStructMemory allocates memory for type C.godot_gdnative_api_struct in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGdnativeApiStructMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGdnativeApiStructValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGdnativeApiStructValue = unsafe.Sizeof([1]C.godot_gdnative_api_struct{})

const sizeOfPtr = unsafe.Sizeof(&struct{}{})

// unpackSGdnativeApiStruct transforms a sliced Go data structure into plain C format.
func unpackSGdnativeApiStruct(x []GdnativeApiStruct) (unpacked *C.godot_gdnative_api_struct, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_gdnative_api_struct) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocGdnativeApiStructMemory(len0)
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

// packSGdnativeApiStruct reads sliced Go data structure out from plain C format.
func packSGdnativeApiStruct(v []GdnativeApiStruct, ptr0 *C.godot_gdnative_api_struct) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGdnativeApiStructValue]C.godot_gdnative_api_struct)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGdnativeApiStructRef(unsafe.Pointer(&ptr1))
	}
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GdnativeExtNativescriptApiStruct) Ref() *C.godot_gdnative_ext_nativescript_api_struct {
	if x == nil {
		return nil
	}
	return x.reff0cd6324
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GdnativeExtNativescriptApiStruct) Free() {
	if x != nil && x.allocsf0cd6324 != nil {
		x.allocsf0cd6324.(*cgoAllocMap).Free()
		x.reff0cd6324 = nil
	}
}

// NewGdnativeExtNativescriptApiStructRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGdnativeExtNativescriptApiStructRef(ref unsafe.Pointer) *GdnativeExtNativescriptApiStruct {
	if ref == nil {
		return nil
	}
	obj := new(GdnativeExtNativescriptApiStruct)
	obj.reff0cd6324 = (*C.godot_gdnative_ext_nativescript_api_struct)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GdnativeExtNativescriptApiStruct) PassRef() (*C.godot_gdnative_ext_nativescript_api_struct, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.reff0cd6324 != nil {
		return x.reff0cd6324, nil
	}
	memf0cd6324 := allocGdnativeExtNativescriptApiStructMemory(1)
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
	reff0cd6324.next, cnext_allocs = unpackSGdnativeApiStruct(x.Next)
	allocsf0cd6324.Borrow(cnext_allocs)

	var cgodot_nativescript_register_class_allocs *cgoAllocMap
	reff0cd6324.godot_nativescript_register_class, cgodot_nativescript_register_class_allocs = x.NativescriptRegisterClass.PassRef()
	allocsf0cd6324.Borrow(cgodot_nativescript_register_class_allocs)

	var cgodot_nativescript_register_tool_class_allocs *cgoAllocMap
	reff0cd6324.godot_nativescript_register_tool_class, cgodot_nativescript_register_tool_class_allocs = x.NativescriptRegisterToolClass.PassRef()
	allocsf0cd6324.Borrow(cgodot_nativescript_register_tool_class_allocs)

	var cgodot_nativescript_register_method_allocs *cgoAllocMap
	reff0cd6324.godot_nativescript_register_method, cgodot_nativescript_register_method_allocs = x.NativescriptRegisterMethod.PassRef()
	allocsf0cd6324.Borrow(cgodot_nativescript_register_method_allocs)

	var cgodot_nativescript_register_property_allocs *cgoAllocMap
	reff0cd6324.godot_nativescript_register_property, cgodot_nativescript_register_property_allocs = x.NativescriptRegisterProperty.PassRef()
	allocsf0cd6324.Borrow(cgodot_nativescript_register_property_allocs)

	var cgodot_nativescript_register_signal_allocs *cgoAllocMap
	reff0cd6324.godot_nativescript_register_signal, cgodot_nativescript_register_signal_allocs = x.NativescriptRegisterSignal.PassRef()
	allocsf0cd6324.Borrow(cgodot_nativescript_register_signal_allocs)

	var cgodot_nativescript_get_userdata_allocs *cgoAllocMap
	reff0cd6324.godot_nativescript_get_userdata, cgodot_nativescript_get_userdata_allocs = x.NativescriptGetUserdata.PassRef()
	allocsf0cd6324.Borrow(cgodot_nativescript_get_userdata_allocs)

	x.reff0cd6324 = reff0cd6324
	x.allocsf0cd6324 = allocsf0cd6324
	return reff0cd6324, allocsf0cd6324

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GdnativeExtNativescriptApiStruct) PassValue() (C.godot_gdnative_ext_nativescript_api_struct, *cgoAllocMap) {
	if x.reff0cd6324 != nil {
		return *x.reff0cd6324, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GdnativeExtNativescriptApiStruct) Deref() {
	if x.reff0cd6324 == nil {
		return
	}
	x.Type = (uint32)(x.reff0cd6324._type)
	x.Version = *NewGdnativeApiVersionRef(unsafe.Pointer(&x.reff0cd6324.version))
	packSGdnativeApiStruct(x.Next, x.reff0cd6324.next)
	x.NativescriptRegisterClass = NewRef(unsafe.Pointer(x.reff0cd6324.godot_nativescript_register_class))
	x.NativescriptRegisterToolClass = NewRef(unsafe.Pointer(x.reff0cd6324.godot_nativescript_register_tool_class))
	x.NativescriptRegisterMethod = NewRef(unsafe.Pointer(x.reff0cd6324.godot_nativescript_register_method))
	x.NativescriptRegisterProperty = NewRef(unsafe.Pointer(x.reff0cd6324.godot_nativescript_register_property))
	x.NativescriptRegisterSignal = NewRef(unsafe.Pointer(x.reff0cd6324.godot_nativescript_register_signal))
	x.NativescriptGetUserdata = NewRef(unsafe.Pointer(x.reff0cd6324.godot_nativescript_get_userdata))
}

// allocGdnativeExtPluginscriptApiStructMemory allocates memory for type C.godot_gdnative_ext_pluginscript_api_struct in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGdnativeExtPluginscriptApiStructMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGdnativeExtPluginscriptApiStructValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGdnativeExtPluginscriptApiStructValue = unsafe.Sizeof([1]C.godot_gdnative_ext_pluginscript_api_struct{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GdnativeExtPluginscriptApiStruct) Ref() *C.godot_gdnative_ext_pluginscript_api_struct {
	if x == nil {
		return nil
	}
	return x.refc52e13a7
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GdnativeExtPluginscriptApiStruct) Free() {
	if x != nil && x.allocsc52e13a7 != nil {
		x.allocsc52e13a7.(*cgoAllocMap).Free()
		x.refc52e13a7 = nil
	}
}

// NewGdnativeExtPluginscriptApiStructRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGdnativeExtPluginscriptApiStructRef(ref unsafe.Pointer) *GdnativeExtPluginscriptApiStruct {
	if ref == nil {
		return nil
	}
	obj := new(GdnativeExtPluginscriptApiStruct)
	obj.refc52e13a7 = (*C.godot_gdnative_ext_pluginscript_api_struct)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GdnativeExtPluginscriptApiStruct) PassRef() (*C.godot_gdnative_ext_pluginscript_api_struct, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refc52e13a7 != nil {
		return x.refc52e13a7, nil
	}
	memc52e13a7 := allocGdnativeExtPluginscriptApiStructMemory(1)
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
	refc52e13a7.next, cnext_allocs = unpackSGdnativeApiStruct(x.Next)
	allocsc52e13a7.Borrow(cnext_allocs)

	var cgodot_pluginscript_register_language_allocs *cgoAllocMap
	refc52e13a7.godot_pluginscript_register_language, cgodot_pluginscript_register_language_allocs = x.PluginscriptRegisterLanguage.PassRef()
	allocsc52e13a7.Borrow(cgodot_pluginscript_register_language_allocs)

	x.refc52e13a7 = refc52e13a7
	x.allocsc52e13a7 = allocsc52e13a7
	return refc52e13a7, allocsc52e13a7

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GdnativeExtPluginscriptApiStruct) PassValue() (C.godot_gdnative_ext_pluginscript_api_struct, *cgoAllocMap) {
	if x.refc52e13a7 != nil {
		return *x.refc52e13a7, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GdnativeExtPluginscriptApiStruct) Deref() {
	if x.refc52e13a7 == nil {
		return
	}
	x.Type = (uint32)(x.refc52e13a7._type)
	x.Version = *NewGdnativeApiVersionRef(unsafe.Pointer(&x.refc52e13a7.version))
	packSGdnativeApiStruct(x.Next, x.refc52e13a7.next)
	x.PluginscriptRegisterLanguage = NewRef(unsafe.Pointer(x.refc52e13a7.godot_pluginscript_register_language))
}

// allocGdnativeExtArvrApiStructMemory allocates memory for type C.godot_gdnative_ext_arvr_api_struct in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGdnativeExtArvrApiStructMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGdnativeExtArvrApiStructValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGdnativeExtArvrApiStructValue = unsafe.Sizeof([1]C.godot_gdnative_ext_arvr_api_struct{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GdnativeExtArvrApiStruct) Ref() *C.godot_gdnative_ext_arvr_api_struct {
	if x == nil {
		return nil
	}
	return x.ref64de5bf5
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GdnativeExtArvrApiStruct) Free() {
	if x != nil && x.allocs64de5bf5 != nil {
		x.allocs64de5bf5.(*cgoAllocMap).Free()
		x.ref64de5bf5 = nil
	}
}

// NewGdnativeExtArvrApiStructRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGdnativeExtArvrApiStructRef(ref unsafe.Pointer) *GdnativeExtArvrApiStruct {
	if ref == nil {
		return nil
	}
	obj := new(GdnativeExtArvrApiStruct)
	obj.ref64de5bf5 = (*C.godot_gdnative_ext_arvr_api_struct)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GdnativeExtArvrApiStruct) PassRef() (*C.godot_gdnative_ext_arvr_api_struct, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref64de5bf5 != nil {
		return x.ref64de5bf5, nil
	}
	mem64de5bf5 := allocGdnativeExtArvrApiStructMemory(1)
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
	ref64de5bf5.next, cnext_allocs = unpackSGdnativeApiStruct(x.Next)
	allocs64de5bf5.Borrow(cnext_allocs)

	var cgodot_arvr_register_interface_allocs *cgoAllocMap
	ref64de5bf5.godot_arvr_register_interface, cgodot_arvr_register_interface_allocs = x.ArvrRegisterInterface.PassRef()
	allocs64de5bf5.Borrow(cgodot_arvr_register_interface_allocs)

	var cgodot_arvr_get_worldscale_allocs *cgoAllocMap
	ref64de5bf5.godot_arvr_get_worldscale, cgodot_arvr_get_worldscale_allocs = x.ArvrGetWorldscale.PassRef()
	allocs64de5bf5.Borrow(cgodot_arvr_get_worldscale_allocs)

	var cgodot_arvr_get_reference_frame_allocs *cgoAllocMap
	ref64de5bf5.godot_arvr_get_reference_frame, cgodot_arvr_get_reference_frame_allocs = x.ArvrGetReferenceFrame.PassRef()
	allocs64de5bf5.Borrow(cgodot_arvr_get_reference_frame_allocs)

	var cgodot_arvr_blit_allocs *cgoAllocMap
	ref64de5bf5.godot_arvr_blit, cgodot_arvr_blit_allocs = x.ArvrBlit.PassRef()
	allocs64de5bf5.Borrow(cgodot_arvr_blit_allocs)

	var cgodot_arvr_get_texid_allocs *cgoAllocMap
	ref64de5bf5.godot_arvr_get_texid, cgodot_arvr_get_texid_allocs = x.ArvrGetTexid.PassRef()
	allocs64de5bf5.Borrow(cgodot_arvr_get_texid_allocs)

	var cgodot_arvr_add_controller_allocs *cgoAllocMap
	ref64de5bf5.godot_arvr_add_controller, cgodot_arvr_add_controller_allocs = x.ArvrAddController.PassRef()
	allocs64de5bf5.Borrow(cgodot_arvr_add_controller_allocs)

	var cgodot_arvr_remove_controller_allocs *cgoAllocMap
	ref64de5bf5.godot_arvr_remove_controller, cgodot_arvr_remove_controller_allocs = x.ArvrRemoveController.PassRef()
	allocs64de5bf5.Borrow(cgodot_arvr_remove_controller_allocs)

	var cgodot_arvr_set_controller_transform_allocs *cgoAllocMap
	ref64de5bf5.godot_arvr_set_controller_transform, cgodot_arvr_set_controller_transform_allocs = x.ArvrSetControllerTransform.PassRef()
	allocs64de5bf5.Borrow(cgodot_arvr_set_controller_transform_allocs)

	var cgodot_arvr_set_controller_button_allocs *cgoAllocMap
	ref64de5bf5.godot_arvr_set_controller_button, cgodot_arvr_set_controller_button_allocs = x.ArvrSetControllerButton.PassRef()
	allocs64de5bf5.Borrow(cgodot_arvr_set_controller_button_allocs)

	var cgodot_arvr_set_controller_axis_allocs *cgoAllocMap
	ref64de5bf5.godot_arvr_set_controller_axis, cgodot_arvr_set_controller_axis_allocs = x.ArvrSetControllerAxis.PassRef()
	allocs64de5bf5.Borrow(cgodot_arvr_set_controller_axis_allocs)

	var cgodot_arvr_get_controller_rumble_allocs *cgoAllocMap
	ref64de5bf5.godot_arvr_get_controller_rumble, cgodot_arvr_get_controller_rumble_allocs = x.ArvrGetControllerRumble.PassRef()
	allocs64de5bf5.Borrow(cgodot_arvr_get_controller_rumble_allocs)

	x.ref64de5bf5 = ref64de5bf5
	x.allocs64de5bf5 = allocs64de5bf5
	return ref64de5bf5, allocs64de5bf5

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GdnativeExtArvrApiStruct) PassValue() (C.godot_gdnative_ext_arvr_api_struct, *cgoAllocMap) {
	if x.ref64de5bf5 != nil {
		return *x.ref64de5bf5, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GdnativeExtArvrApiStruct) Deref() {
	if x.ref64de5bf5 == nil {
		return
	}
	x.Type = (uint32)(x.ref64de5bf5._type)
	x.Version = *NewGdnativeApiVersionRef(unsafe.Pointer(&x.ref64de5bf5.version))
	packSGdnativeApiStruct(x.Next, x.ref64de5bf5.next)
	x.ArvrRegisterInterface = NewRef(unsafe.Pointer(x.ref64de5bf5.godot_arvr_register_interface))
	x.ArvrGetWorldscale = NewRealRef(unsafe.Pointer(x.ref64de5bf5.godot_arvr_get_worldscale))
	x.ArvrGetReferenceFrame = NewTransformRef(unsafe.Pointer(x.ref64de5bf5.godot_arvr_get_reference_frame))
	x.ArvrBlit = NewRef(unsafe.Pointer(x.ref64de5bf5.godot_arvr_blit))
	x.ArvrGetTexid = NewIntRef(unsafe.Pointer(x.ref64de5bf5.godot_arvr_get_texid))
	x.ArvrAddController = NewIntRef(unsafe.Pointer(x.ref64de5bf5.godot_arvr_add_controller))
	x.ArvrRemoveController = NewRef(unsafe.Pointer(x.ref64de5bf5.godot_arvr_remove_controller))
	x.ArvrSetControllerTransform = NewRef(unsafe.Pointer(x.ref64de5bf5.godot_arvr_set_controller_transform))
	x.ArvrSetControllerButton = NewRef(unsafe.Pointer(x.ref64de5bf5.godot_arvr_set_controller_button))
	x.ArvrSetControllerAxis = NewRef(unsafe.Pointer(x.ref64de5bf5.godot_arvr_set_controller_axis))
	x.ArvrGetControllerRumble = NewRealRef(unsafe.Pointer(x.ref64de5bf5.godot_arvr_get_controller_rumble))
}

// allocGdnativeCoreApiStructMemory allocates memory for type C.godot_gdnative_core_api_struct in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGdnativeCoreApiStructMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGdnativeCoreApiStructValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGdnativeCoreApiStructValue = unsafe.Sizeof([1]C.godot_gdnative_core_api_struct{})

// allocPGdnativeApiStructMemory allocates memory for type *C.godot_gdnative_api_struct in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPGdnativeApiStructMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPGdnativeApiStructValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPGdnativeApiStructValue = unsafe.Sizeof([1]*C.godot_gdnative_api_struct{})

// unpackSSGdnativeApiStruct transforms a sliced Go data structure into plain C format.
func unpackSSGdnativeApiStruct(x [][]GdnativeApiStruct) (unpacked **C.godot_gdnative_api_struct, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.godot_gdnative_api_struct) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPGdnativeApiStructMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.godot_gdnative_api_struct)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocGdnativeApiStructMemory(len1)
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

// packSSGdnativeApiStruct reads sliced Go data structure out from plain C format.
func packSSGdnativeApiStruct(v [][]GdnativeApiStruct, ptr0 **C.godot_gdnative_api_struct) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.godot_gdnative_api_struct)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfGdnativeApiStructValue]C.godot_gdnative_api_struct)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *NewGdnativeApiStructRef(unsafe.Pointer(&ptr2))
		}
	}
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GdnativeCoreApiStruct) Ref() *C.godot_gdnative_core_api_struct {
	if x == nil {
		return nil
	}
	return x.ref57717e51
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GdnativeCoreApiStruct) Free() {
	if x != nil && x.allocs57717e51 != nil {
		x.allocs57717e51.(*cgoAllocMap).Free()
		x.ref57717e51 = nil
	}
}

// NewGdnativeCoreApiStructRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGdnativeCoreApiStructRef(ref unsafe.Pointer) *GdnativeCoreApiStruct {
	if ref == nil {
		return nil
	}
	obj := new(GdnativeCoreApiStruct)
	obj.ref57717e51 = (*C.godot_gdnative_core_api_struct)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GdnativeCoreApiStruct) PassRef() (*C.godot_gdnative_core_api_struct, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref57717e51 != nil {
		return x.ref57717e51, nil
	}
	mem57717e51 := allocGdnativeCoreApiStructMemory(1)
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
	ref57717e51.next, cnext_allocs = unpackSGdnativeApiStruct(x.Next)
	allocs57717e51.Borrow(cnext_allocs)

	var cnum_extensions_allocs *cgoAllocMap
	ref57717e51.num_extensions, cnum_extensions_allocs = (C.uint)(x.NumExtensions), cgoAllocsUnknown
	allocs57717e51.Borrow(cnum_extensions_allocs)

	var cextensions_allocs *cgoAllocMap
	ref57717e51.extensions, cextensions_allocs = unpackSSGdnativeApiStruct(x.Extensions)
	allocs57717e51.Borrow(cextensions_allocs)

	var cgodot_color_new_rgba_allocs *cgoAllocMap
	ref57717e51.godot_color_new_rgba, cgodot_color_new_rgba_allocs = x.ColorNewRgba.PassRef()
	allocs57717e51.Borrow(cgodot_color_new_rgba_allocs)

	var cgodot_color_new_rgb_allocs *cgoAllocMap
	ref57717e51.godot_color_new_rgb, cgodot_color_new_rgb_allocs = x.ColorNewRgb.PassRef()
	allocs57717e51.Borrow(cgodot_color_new_rgb_allocs)

	var cgodot_color_get_r_allocs *cgoAllocMap
	ref57717e51.godot_color_get_r, cgodot_color_get_r_allocs = x.ColorGetR.PassRef()
	allocs57717e51.Borrow(cgodot_color_get_r_allocs)

	var cgodot_color_set_r_allocs *cgoAllocMap
	ref57717e51.godot_color_set_r, cgodot_color_set_r_allocs = x.ColorSetR.PassRef()
	allocs57717e51.Borrow(cgodot_color_set_r_allocs)

	var cgodot_color_get_g_allocs *cgoAllocMap
	ref57717e51.godot_color_get_g, cgodot_color_get_g_allocs = x.ColorGetG.PassRef()
	allocs57717e51.Borrow(cgodot_color_get_g_allocs)

	var cgodot_color_set_g_allocs *cgoAllocMap
	ref57717e51.godot_color_set_g, cgodot_color_set_g_allocs = x.ColorSetG.PassRef()
	allocs57717e51.Borrow(cgodot_color_set_g_allocs)

	var cgodot_color_get_b_allocs *cgoAllocMap
	ref57717e51.godot_color_get_b, cgodot_color_get_b_allocs = x.ColorGetB.PassRef()
	allocs57717e51.Borrow(cgodot_color_get_b_allocs)

	var cgodot_color_set_b_allocs *cgoAllocMap
	ref57717e51.godot_color_set_b, cgodot_color_set_b_allocs = x.ColorSetB.PassRef()
	allocs57717e51.Borrow(cgodot_color_set_b_allocs)

	var cgodot_color_get_a_allocs *cgoAllocMap
	ref57717e51.godot_color_get_a, cgodot_color_get_a_allocs = x.ColorGetA.PassRef()
	allocs57717e51.Borrow(cgodot_color_get_a_allocs)

	var cgodot_color_set_a_allocs *cgoAllocMap
	ref57717e51.godot_color_set_a, cgodot_color_set_a_allocs = x.ColorSetA.PassRef()
	allocs57717e51.Borrow(cgodot_color_set_a_allocs)

	var cgodot_color_get_h_allocs *cgoAllocMap
	ref57717e51.godot_color_get_h, cgodot_color_get_h_allocs = x.ColorGetH.PassRef()
	allocs57717e51.Borrow(cgodot_color_get_h_allocs)

	var cgodot_color_get_s_allocs *cgoAllocMap
	ref57717e51.godot_color_get_s, cgodot_color_get_s_allocs = x.ColorGetS.PassRef()
	allocs57717e51.Borrow(cgodot_color_get_s_allocs)

	var cgodot_color_get_v_allocs *cgoAllocMap
	ref57717e51.godot_color_get_v, cgodot_color_get_v_allocs = x.ColorGetV.PassRef()
	allocs57717e51.Borrow(cgodot_color_get_v_allocs)

	var cgodot_color_as_string_allocs *cgoAllocMap
	ref57717e51.godot_color_as_string, cgodot_color_as_string_allocs = x.ColorAsString.PassRef()
	allocs57717e51.Borrow(cgodot_color_as_string_allocs)

	var cgodot_color_to_rgba32_allocs *cgoAllocMap
	ref57717e51.godot_color_to_rgba32, cgodot_color_to_rgba32_allocs = x.ColorToRgba32.PassRef()
	allocs57717e51.Borrow(cgodot_color_to_rgba32_allocs)

	var cgodot_color_to_argb32_allocs *cgoAllocMap
	ref57717e51.godot_color_to_argb32, cgodot_color_to_argb32_allocs = x.ColorToArgb32.PassRef()
	allocs57717e51.Borrow(cgodot_color_to_argb32_allocs)

	var cgodot_color_gray_allocs *cgoAllocMap
	ref57717e51.godot_color_gray, cgodot_color_gray_allocs = x.ColorGray.PassRef()
	allocs57717e51.Borrow(cgodot_color_gray_allocs)

	var cgodot_color_inverted_allocs *cgoAllocMap
	ref57717e51.godot_color_inverted, cgodot_color_inverted_allocs = x.ColorInverted.PassRef()
	allocs57717e51.Borrow(cgodot_color_inverted_allocs)

	var cgodot_color_contrasted_allocs *cgoAllocMap
	ref57717e51.godot_color_contrasted, cgodot_color_contrasted_allocs = x.ColorContrasted.PassRef()
	allocs57717e51.Borrow(cgodot_color_contrasted_allocs)

	var cgodot_color_linear_interpolate_allocs *cgoAllocMap
	ref57717e51.godot_color_linear_interpolate, cgodot_color_linear_interpolate_allocs = x.ColorLinearInterpolate.PassRef()
	allocs57717e51.Borrow(cgodot_color_linear_interpolate_allocs)

	var cgodot_color_blend_allocs *cgoAllocMap
	ref57717e51.godot_color_blend, cgodot_color_blend_allocs = x.ColorBlend.PassRef()
	allocs57717e51.Borrow(cgodot_color_blend_allocs)

	var cgodot_color_to_html_allocs *cgoAllocMap
	ref57717e51.godot_color_to_html, cgodot_color_to_html_allocs = x.ColorToHtml.PassRef()
	allocs57717e51.Borrow(cgodot_color_to_html_allocs)

	var cgodot_color_operator_equal_allocs *cgoAllocMap
	ref57717e51.godot_color_operator_equal, cgodot_color_operator_equal_allocs = x.ColorOperatorEqual.PassRef()
	allocs57717e51.Borrow(cgodot_color_operator_equal_allocs)

	var cgodot_color_operator_less_allocs *cgoAllocMap
	ref57717e51.godot_color_operator_less, cgodot_color_operator_less_allocs = x.ColorOperatorLess.PassRef()
	allocs57717e51.Borrow(cgodot_color_operator_less_allocs)

	var cgodot_vector2_new_allocs *cgoAllocMap
	ref57717e51.godot_vector2_new, cgodot_vector2_new_allocs = x.Vector2New.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_new_allocs)

	var cgodot_vector2_as_string_allocs *cgoAllocMap
	ref57717e51.godot_vector2_as_string, cgodot_vector2_as_string_allocs = x.Vector2AsString.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_as_string_allocs)

	var cgodot_vector2_normalized_allocs *cgoAllocMap
	ref57717e51.godot_vector2_normalized, cgodot_vector2_normalized_allocs = x.Vector2Normalized.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_normalized_allocs)

	var cgodot_vector2_length_allocs *cgoAllocMap
	ref57717e51.godot_vector2_length, cgodot_vector2_length_allocs = x.Vector2Length.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_length_allocs)

	var cgodot_vector2_angle_allocs *cgoAllocMap
	ref57717e51.godot_vector2_angle, cgodot_vector2_angle_allocs = x.Vector2Angle.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_angle_allocs)

	var cgodot_vector2_length_squared_allocs *cgoAllocMap
	ref57717e51.godot_vector2_length_squared, cgodot_vector2_length_squared_allocs = x.Vector2LengthSquared.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_length_squared_allocs)

	var cgodot_vector2_is_normalized_allocs *cgoAllocMap
	ref57717e51.godot_vector2_is_normalized, cgodot_vector2_is_normalized_allocs = x.Vector2IsNormalized.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_is_normalized_allocs)

	var cgodot_vector2_distance_to_allocs *cgoAllocMap
	ref57717e51.godot_vector2_distance_to, cgodot_vector2_distance_to_allocs = x.Vector2DistanceTo.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_distance_to_allocs)

	var cgodot_vector2_distance_squared_to_allocs *cgoAllocMap
	ref57717e51.godot_vector2_distance_squared_to, cgodot_vector2_distance_squared_to_allocs = x.Vector2DistanceSquaredTo.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_distance_squared_to_allocs)

	var cgodot_vector2_angle_to_allocs *cgoAllocMap
	ref57717e51.godot_vector2_angle_to, cgodot_vector2_angle_to_allocs = x.Vector2AngleTo.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_angle_to_allocs)

	var cgodot_vector2_angle_to_point_allocs *cgoAllocMap
	ref57717e51.godot_vector2_angle_to_point, cgodot_vector2_angle_to_point_allocs = x.Vector2AngleToPoint.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_angle_to_point_allocs)

	var cgodot_vector2_linear_interpolate_allocs *cgoAllocMap
	ref57717e51.godot_vector2_linear_interpolate, cgodot_vector2_linear_interpolate_allocs = x.Vector2LinearInterpolate.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_linear_interpolate_allocs)

	var cgodot_vector2_cubic_interpolate_allocs *cgoAllocMap
	ref57717e51.godot_vector2_cubic_interpolate, cgodot_vector2_cubic_interpolate_allocs = x.Vector2CubicInterpolate.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_cubic_interpolate_allocs)

	var cgodot_vector2_rotated_allocs *cgoAllocMap
	ref57717e51.godot_vector2_rotated, cgodot_vector2_rotated_allocs = x.Vector2Rotated.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_rotated_allocs)

	var cgodot_vector2_tangent_allocs *cgoAllocMap
	ref57717e51.godot_vector2_tangent, cgodot_vector2_tangent_allocs = x.Vector2Tangent.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_tangent_allocs)

	var cgodot_vector2_floor_allocs *cgoAllocMap
	ref57717e51.godot_vector2_floor, cgodot_vector2_floor_allocs = x.Vector2Floor.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_floor_allocs)

	var cgodot_vector2_snapped_allocs *cgoAllocMap
	ref57717e51.godot_vector2_snapped, cgodot_vector2_snapped_allocs = x.Vector2Snapped.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_snapped_allocs)

	var cgodot_vector2_aspect_allocs *cgoAllocMap
	ref57717e51.godot_vector2_aspect, cgodot_vector2_aspect_allocs = x.Vector2Aspect.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_aspect_allocs)

	var cgodot_vector2_dot_allocs *cgoAllocMap
	ref57717e51.godot_vector2_dot, cgodot_vector2_dot_allocs = x.Vector2Dot.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_dot_allocs)

	var cgodot_vector2_slide_allocs *cgoAllocMap
	ref57717e51.godot_vector2_slide, cgodot_vector2_slide_allocs = x.Vector2Slide.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_slide_allocs)

	var cgodot_vector2_bounce_allocs *cgoAllocMap
	ref57717e51.godot_vector2_bounce, cgodot_vector2_bounce_allocs = x.Vector2Bounce.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_bounce_allocs)

	var cgodot_vector2_reflect_allocs *cgoAllocMap
	ref57717e51.godot_vector2_reflect, cgodot_vector2_reflect_allocs = x.Vector2Reflect.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_reflect_allocs)

	var cgodot_vector2_abs_allocs *cgoAllocMap
	ref57717e51.godot_vector2_abs, cgodot_vector2_abs_allocs = x.Vector2Abs.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_abs_allocs)

	var cgodot_vector2_clamped_allocs *cgoAllocMap
	ref57717e51.godot_vector2_clamped, cgodot_vector2_clamped_allocs = x.Vector2Clamped.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_clamped_allocs)

	var cgodot_vector2_operator_add_allocs *cgoAllocMap
	ref57717e51.godot_vector2_operator_add, cgodot_vector2_operator_add_allocs = x.Vector2OperatorAdd.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_operator_add_allocs)

	var cgodot_vector2_operator_substract_allocs *cgoAllocMap
	ref57717e51.godot_vector2_operator_substract, cgodot_vector2_operator_substract_allocs = x.Vector2OperatorSubstract.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_operator_substract_allocs)

	var cgodot_vector2_operator_multiply_vector_allocs *cgoAllocMap
	ref57717e51.godot_vector2_operator_multiply_vector, cgodot_vector2_operator_multiply_vector_allocs = x.Vector2OperatorMultiplyVector.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_operator_multiply_vector_allocs)

	var cgodot_vector2_operator_multiply_scalar_allocs *cgoAllocMap
	ref57717e51.godot_vector2_operator_multiply_scalar, cgodot_vector2_operator_multiply_scalar_allocs = x.Vector2OperatorMultiplyScalar.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_operator_multiply_scalar_allocs)

	var cgodot_vector2_operator_divide_vector_allocs *cgoAllocMap
	ref57717e51.godot_vector2_operator_divide_vector, cgodot_vector2_operator_divide_vector_allocs = x.Vector2OperatorDivideVector.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_operator_divide_vector_allocs)

	var cgodot_vector2_operator_divide_scalar_allocs *cgoAllocMap
	ref57717e51.godot_vector2_operator_divide_scalar, cgodot_vector2_operator_divide_scalar_allocs = x.Vector2OperatorDivideScalar.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_operator_divide_scalar_allocs)

	var cgodot_vector2_operator_equal_allocs *cgoAllocMap
	ref57717e51.godot_vector2_operator_equal, cgodot_vector2_operator_equal_allocs = x.Vector2OperatorEqual.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_operator_equal_allocs)

	var cgodot_vector2_operator_less_allocs *cgoAllocMap
	ref57717e51.godot_vector2_operator_less, cgodot_vector2_operator_less_allocs = x.Vector2OperatorLess.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_operator_less_allocs)

	var cgodot_vector2_operator_neg_allocs *cgoAllocMap
	ref57717e51.godot_vector2_operator_neg, cgodot_vector2_operator_neg_allocs = x.Vector2OperatorNeg.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_operator_neg_allocs)

	var cgodot_vector2_set_x_allocs *cgoAllocMap
	ref57717e51.godot_vector2_set_x, cgodot_vector2_set_x_allocs = x.Vector2SetX.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_set_x_allocs)

	var cgodot_vector2_set_y_allocs *cgoAllocMap
	ref57717e51.godot_vector2_set_y, cgodot_vector2_set_y_allocs = x.Vector2SetY.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_set_y_allocs)

	var cgodot_vector2_get_x_allocs *cgoAllocMap
	ref57717e51.godot_vector2_get_x, cgodot_vector2_get_x_allocs = x.Vector2GetX.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_get_x_allocs)

	var cgodot_vector2_get_y_allocs *cgoAllocMap
	ref57717e51.godot_vector2_get_y, cgodot_vector2_get_y_allocs = x.Vector2GetY.PassRef()
	allocs57717e51.Borrow(cgodot_vector2_get_y_allocs)

	var cgodot_quat_new_allocs *cgoAllocMap
	ref57717e51.godot_quat_new, cgodot_quat_new_allocs = x.QuatNew.PassRef()
	allocs57717e51.Borrow(cgodot_quat_new_allocs)

	var cgodot_quat_new_with_axis_angle_allocs *cgoAllocMap
	ref57717e51.godot_quat_new_with_axis_angle, cgodot_quat_new_with_axis_angle_allocs = x.QuatNewWithAxisAngle.PassRef()
	allocs57717e51.Borrow(cgodot_quat_new_with_axis_angle_allocs)

	var cgodot_quat_get_x_allocs *cgoAllocMap
	ref57717e51.godot_quat_get_x, cgodot_quat_get_x_allocs = x.QuatGetX.PassRef()
	allocs57717e51.Borrow(cgodot_quat_get_x_allocs)

	var cgodot_quat_set_x_allocs *cgoAllocMap
	ref57717e51.godot_quat_set_x, cgodot_quat_set_x_allocs = x.QuatSetX.PassRef()
	allocs57717e51.Borrow(cgodot_quat_set_x_allocs)

	var cgodot_quat_get_y_allocs *cgoAllocMap
	ref57717e51.godot_quat_get_y, cgodot_quat_get_y_allocs = x.QuatGetY.PassRef()
	allocs57717e51.Borrow(cgodot_quat_get_y_allocs)

	var cgodot_quat_set_y_allocs *cgoAllocMap
	ref57717e51.godot_quat_set_y, cgodot_quat_set_y_allocs = x.QuatSetY.PassRef()
	allocs57717e51.Borrow(cgodot_quat_set_y_allocs)

	var cgodot_quat_get_z_allocs *cgoAllocMap
	ref57717e51.godot_quat_get_z, cgodot_quat_get_z_allocs = x.QuatGetZ.PassRef()
	allocs57717e51.Borrow(cgodot_quat_get_z_allocs)

	var cgodot_quat_set_z_allocs *cgoAllocMap
	ref57717e51.godot_quat_set_z, cgodot_quat_set_z_allocs = x.QuatSetZ.PassRef()
	allocs57717e51.Borrow(cgodot_quat_set_z_allocs)

	var cgodot_quat_get_w_allocs *cgoAllocMap
	ref57717e51.godot_quat_get_w, cgodot_quat_get_w_allocs = x.QuatGetW.PassRef()
	allocs57717e51.Borrow(cgodot_quat_get_w_allocs)

	var cgodot_quat_set_w_allocs *cgoAllocMap
	ref57717e51.godot_quat_set_w, cgodot_quat_set_w_allocs = x.QuatSetW.PassRef()
	allocs57717e51.Borrow(cgodot_quat_set_w_allocs)

	var cgodot_quat_as_string_allocs *cgoAllocMap
	ref57717e51.godot_quat_as_string, cgodot_quat_as_string_allocs = x.QuatAsString.PassRef()
	allocs57717e51.Borrow(cgodot_quat_as_string_allocs)

	var cgodot_quat_length_allocs *cgoAllocMap
	ref57717e51.godot_quat_length, cgodot_quat_length_allocs = x.QuatLength.PassRef()
	allocs57717e51.Borrow(cgodot_quat_length_allocs)

	var cgodot_quat_length_squared_allocs *cgoAllocMap
	ref57717e51.godot_quat_length_squared, cgodot_quat_length_squared_allocs = x.QuatLengthSquared.PassRef()
	allocs57717e51.Borrow(cgodot_quat_length_squared_allocs)

	var cgodot_quat_normalized_allocs *cgoAllocMap
	ref57717e51.godot_quat_normalized, cgodot_quat_normalized_allocs = x.QuatNormalized.PassRef()
	allocs57717e51.Borrow(cgodot_quat_normalized_allocs)

	var cgodot_quat_is_normalized_allocs *cgoAllocMap
	ref57717e51.godot_quat_is_normalized, cgodot_quat_is_normalized_allocs = x.QuatIsNormalized.PassRef()
	allocs57717e51.Borrow(cgodot_quat_is_normalized_allocs)

	var cgodot_quat_inverse_allocs *cgoAllocMap
	ref57717e51.godot_quat_inverse, cgodot_quat_inverse_allocs = x.QuatInverse.PassRef()
	allocs57717e51.Borrow(cgodot_quat_inverse_allocs)

	var cgodot_quat_dot_allocs *cgoAllocMap
	ref57717e51.godot_quat_dot, cgodot_quat_dot_allocs = x.QuatDot.PassRef()
	allocs57717e51.Borrow(cgodot_quat_dot_allocs)

	var cgodot_quat_xform_allocs *cgoAllocMap
	ref57717e51.godot_quat_xform, cgodot_quat_xform_allocs = x.QuatXform.PassRef()
	allocs57717e51.Borrow(cgodot_quat_xform_allocs)

	var cgodot_quat_slerp_allocs *cgoAllocMap
	ref57717e51.godot_quat_slerp, cgodot_quat_slerp_allocs = x.QuatSlerp.PassRef()
	allocs57717e51.Borrow(cgodot_quat_slerp_allocs)

	var cgodot_quat_slerpni_allocs *cgoAllocMap
	ref57717e51.godot_quat_slerpni, cgodot_quat_slerpni_allocs = x.QuatSlerpni.PassRef()
	allocs57717e51.Borrow(cgodot_quat_slerpni_allocs)

	var cgodot_quat_cubic_slerp_allocs *cgoAllocMap
	ref57717e51.godot_quat_cubic_slerp, cgodot_quat_cubic_slerp_allocs = x.QuatCubicSlerp.PassRef()
	allocs57717e51.Borrow(cgodot_quat_cubic_slerp_allocs)

	var cgodot_quat_operator_multiply_allocs *cgoAllocMap
	ref57717e51.godot_quat_operator_multiply, cgodot_quat_operator_multiply_allocs = x.QuatOperatorMultiply.PassRef()
	allocs57717e51.Borrow(cgodot_quat_operator_multiply_allocs)

	var cgodot_quat_operator_add_allocs *cgoAllocMap
	ref57717e51.godot_quat_operator_add, cgodot_quat_operator_add_allocs = x.QuatOperatorAdd.PassRef()
	allocs57717e51.Borrow(cgodot_quat_operator_add_allocs)

	var cgodot_quat_operator_substract_allocs *cgoAllocMap
	ref57717e51.godot_quat_operator_substract, cgodot_quat_operator_substract_allocs = x.QuatOperatorSubstract.PassRef()
	allocs57717e51.Borrow(cgodot_quat_operator_substract_allocs)

	var cgodot_quat_operator_divide_allocs *cgoAllocMap
	ref57717e51.godot_quat_operator_divide, cgodot_quat_operator_divide_allocs = x.QuatOperatorDivide.PassRef()
	allocs57717e51.Borrow(cgodot_quat_operator_divide_allocs)

	var cgodot_quat_operator_equal_allocs *cgoAllocMap
	ref57717e51.godot_quat_operator_equal, cgodot_quat_operator_equal_allocs = x.QuatOperatorEqual.PassRef()
	allocs57717e51.Borrow(cgodot_quat_operator_equal_allocs)

	var cgodot_quat_operator_neg_allocs *cgoAllocMap
	ref57717e51.godot_quat_operator_neg, cgodot_quat_operator_neg_allocs = x.QuatOperatorNeg.PassRef()
	allocs57717e51.Borrow(cgodot_quat_operator_neg_allocs)

	var cgodot_basis_new_with_rows_allocs *cgoAllocMap
	ref57717e51.godot_basis_new_with_rows, cgodot_basis_new_with_rows_allocs = x.BasisNewWithRows.PassRef()
	allocs57717e51.Borrow(cgodot_basis_new_with_rows_allocs)

	var cgodot_basis_new_with_axis_and_angle_allocs *cgoAllocMap
	ref57717e51.godot_basis_new_with_axis_and_angle, cgodot_basis_new_with_axis_and_angle_allocs = x.BasisNewWithAxisAndAngle.PassRef()
	allocs57717e51.Borrow(cgodot_basis_new_with_axis_and_angle_allocs)

	var cgodot_basis_new_with_euler_allocs *cgoAllocMap
	ref57717e51.godot_basis_new_with_euler, cgodot_basis_new_with_euler_allocs = x.BasisNewWithEuler.PassRef()
	allocs57717e51.Borrow(cgodot_basis_new_with_euler_allocs)

	var cgodot_basis_as_string_allocs *cgoAllocMap
	ref57717e51.godot_basis_as_string, cgodot_basis_as_string_allocs = x.BasisAsString.PassRef()
	allocs57717e51.Borrow(cgodot_basis_as_string_allocs)

	var cgodot_basis_inverse_allocs *cgoAllocMap
	ref57717e51.godot_basis_inverse, cgodot_basis_inverse_allocs = x.BasisInverse.PassRef()
	allocs57717e51.Borrow(cgodot_basis_inverse_allocs)

	var cgodot_basis_transposed_allocs *cgoAllocMap
	ref57717e51.godot_basis_transposed, cgodot_basis_transposed_allocs = x.BasisTransposed.PassRef()
	allocs57717e51.Borrow(cgodot_basis_transposed_allocs)

	var cgodot_basis_orthonormalized_allocs *cgoAllocMap
	ref57717e51.godot_basis_orthonormalized, cgodot_basis_orthonormalized_allocs = x.BasisOrthonormalized.PassRef()
	allocs57717e51.Borrow(cgodot_basis_orthonormalized_allocs)

	var cgodot_basis_determinant_allocs *cgoAllocMap
	ref57717e51.godot_basis_determinant, cgodot_basis_determinant_allocs = x.BasisDeterminant.PassRef()
	allocs57717e51.Borrow(cgodot_basis_determinant_allocs)

	var cgodot_basis_rotated_allocs *cgoAllocMap
	ref57717e51.godot_basis_rotated, cgodot_basis_rotated_allocs = x.BasisRotated.PassRef()
	allocs57717e51.Borrow(cgodot_basis_rotated_allocs)

	var cgodot_basis_scaled_allocs *cgoAllocMap
	ref57717e51.godot_basis_scaled, cgodot_basis_scaled_allocs = x.BasisScaled.PassRef()
	allocs57717e51.Borrow(cgodot_basis_scaled_allocs)

	var cgodot_basis_get_scale_allocs *cgoAllocMap
	ref57717e51.godot_basis_get_scale, cgodot_basis_get_scale_allocs = x.BasisGetScale.PassRef()
	allocs57717e51.Borrow(cgodot_basis_get_scale_allocs)

	var cgodot_basis_get_euler_allocs *cgoAllocMap
	ref57717e51.godot_basis_get_euler, cgodot_basis_get_euler_allocs = x.BasisGetEuler.PassRef()
	allocs57717e51.Borrow(cgodot_basis_get_euler_allocs)

	var cgodot_basis_tdotx_allocs *cgoAllocMap
	ref57717e51.godot_basis_tdotx, cgodot_basis_tdotx_allocs = x.BasisTdotx.PassRef()
	allocs57717e51.Borrow(cgodot_basis_tdotx_allocs)

	var cgodot_basis_tdoty_allocs *cgoAllocMap
	ref57717e51.godot_basis_tdoty, cgodot_basis_tdoty_allocs = x.BasisTdoty.PassRef()
	allocs57717e51.Borrow(cgodot_basis_tdoty_allocs)

	var cgodot_basis_tdotz_allocs *cgoAllocMap
	ref57717e51.godot_basis_tdotz, cgodot_basis_tdotz_allocs = x.BasisTdotz.PassRef()
	allocs57717e51.Borrow(cgodot_basis_tdotz_allocs)

	var cgodot_basis_xform_allocs *cgoAllocMap
	ref57717e51.godot_basis_xform, cgodot_basis_xform_allocs = x.BasisXform.PassRef()
	allocs57717e51.Borrow(cgodot_basis_xform_allocs)

	var cgodot_basis_xform_inv_allocs *cgoAllocMap
	ref57717e51.godot_basis_xform_inv, cgodot_basis_xform_inv_allocs = x.BasisXformInv.PassRef()
	allocs57717e51.Borrow(cgodot_basis_xform_inv_allocs)

	var cgodot_basis_get_orthogonal_index_allocs *cgoAllocMap
	ref57717e51.godot_basis_get_orthogonal_index, cgodot_basis_get_orthogonal_index_allocs = x.BasisGetOrthogonalIndex.PassRef()
	allocs57717e51.Borrow(cgodot_basis_get_orthogonal_index_allocs)

	var cgodot_basis_new_allocs *cgoAllocMap
	ref57717e51.godot_basis_new, cgodot_basis_new_allocs = x.BasisNew.PassRef()
	allocs57717e51.Borrow(cgodot_basis_new_allocs)

	var cgodot_basis_new_with_euler_quat_allocs *cgoAllocMap
	ref57717e51.godot_basis_new_with_euler_quat, cgodot_basis_new_with_euler_quat_allocs = x.BasisNewWithEulerQuat.PassRef()
	allocs57717e51.Borrow(cgodot_basis_new_with_euler_quat_allocs)

	var cgodot_basis_get_elements_allocs *cgoAllocMap
	ref57717e51.godot_basis_get_elements, cgodot_basis_get_elements_allocs = x.BasisGetElements.PassRef()
	allocs57717e51.Borrow(cgodot_basis_get_elements_allocs)

	var cgodot_basis_get_axis_allocs *cgoAllocMap
	ref57717e51.godot_basis_get_axis, cgodot_basis_get_axis_allocs = x.BasisGetAxis.PassRef()
	allocs57717e51.Borrow(cgodot_basis_get_axis_allocs)

	var cgodot_basis_set_axis_allocs *cgoAllocMap
	ref57717e51.godot_basis_set_axis, cgodot_basis_set_axis_allocs = x.BasisSetAxis.PassRef()
	allocs57717e51.Borrow(cgodot_basis_set_axis_allocs)

	var cgodot_basis_get_row_allocs *cgoAllocMap
	ref57717e51.godot_basis_get_row, cgodot_basis_get_row_allocs = x.BasisGetRow.PassRef()
	allocs57717e51.Borrow(cgodot_basis_get_row_allocs)

	var cgodot_basis_set_row_allocs *cgoAllocMap
	ref57717e51.godot_basis_set_row, cgodot_basis_set_row_allocs = x.BasisSetRow.PassRef()
	allocs57717e51.Borrow(cgodot_basis_set_row_allocs)

	var cgodot_basis_operator_equal_allocs *cgoAllocMap
	ref57717e51.godot_basis_operator_equal, cgodot_basis_operator_equal_allocs = x.BasisOperatorEqual.PassRef()
	allocs57717e51.Borrow(cgodot_basis_operator_equal_allocs)

	var cgodot_basis_operator_add_allocs *cgoAllocMap
	ref57717e51.godot_basis_operator_add, cgodot_basis_operator_add_allocs = x.BasisOperatorAdd.PassRef()
	allocs57717e51.Borrow(cgodot_basis_operator_add_allocs)

	var cgodot_basis_operator_substract_allocs *cgoAllocMap
	ref57717e51.godot_basis_operator_substract, cgodot_basis_operator_substract_allocs = x.BasisOperatorSubstract.PassRef()
	allocs57717e51.Borrow(cgodot_basis_operator_substract_allocs)

	var cgodot_basis_operator_multiply_vector_allocs *cgoAllocMap
	ref57717e51.godot_basis_operator_multiply_vector, cgodot_basis_operator_multiply_vector_allocs = x.BasisOperatorMultiplyVector.PassRef()
	allocs57717e51.Borrow(cgodot_basis_operator_multiply_vector_allocs)

	var cgodot_basis_operator_multiply_scalar_allocs *cgoAllocMap
	ref57717e51.godot_basis_operator_multiply_scalar, cgodot_basis_operator_multiply_scalar_allocs = x.BasisOperatorMultiplyScalar.PassRef()
	allocs57717e51.Borrow(cgodot_basis_operator_multiply_scalar_allocs)

	var cgodot_vector3_new_allocs *cgoAllocMap
	ref57717e51.godot_vector3_new, cgodot_vector3_new_allocs = x.Vector3New.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_new_allocs)

	var cgodot_vector3_as_string_allocs *cgoAllocMap
	ref57717e51.godot_vector3_as_string, cgodot_vector3_as_string_allocs = x.Vector3AsString.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_as_string_allocs)

	var cgodot_vector3_min_axis_allocs *cgoAllocMap
	ref57717e51.godot_vector3_min_axis, cgodot_vector3_min_axis_allocs = x.Vector3MinAxis.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_min_axis_allocs)

	var cgodot_vector3_max_axis_allocs *cgoAllocMap
	ref57717e51.godot_vector3_max_axis, cgodot_vector3_max_axis_allocs = x.Vector3MaxAxis.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_max_axis_allocs)

	var cgodot_vector3_length_allocs *cgoAllocMap
	ref57717e51.godot_vector3_length, cgodot_vector3_length_allocs = x.Vector3Length.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_length_allocs)

	var cgodot_vector3_length_squared_allocs *cgoAllocMap
	ref57717e51.godot_vector3_length_squared, cgodot_vector3_length_squared_allocs = x.Vector3LengthSquared.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_length_squared_allocs)

	var cgodot_vector3_is_normalized_allocs *cgoAllocMap
	ref57717e51.godot_vector3_is_normalized, cgodot_vector3_is_normalized_allocs = x.Vector3IsNormalized.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_is_normalized_allocs)

	var cgodot_vector3_normalized_allocs *cgoAllocMap
	ref57717e51.godot_vector3_normalized, cgodot_vector3_normalized_allocs = x.Vector3Normalized.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_normalized_allocs)

	var cgodot_vector3_inverse_allocs *cgoAllocMap
	ref57717e51.godot_vector3_inverse, cgodot_vector3_inverse_allocs = x.Vector3Inverse.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_inverse_allocs)

	var cgodot_vector3_snapped_allocs *cgoAllocMap
	ref57717e51.godot_vector3_snapped, cgodot_vector3_snapped_allocs = x.Vector3Snapped.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_snapped_allocs)

	var cgodot_vector3_rotated_allocs *cgoAllocMap
	ref57717e51.godot_vector3_rotated, cgodot_vector3_rotated_allocs = x.Vector3Rotated.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_rotated_allocs)

	var cgodot_vector3_linear_interpolate_allocs *cgoAllocMap
	ref57717e51.godot_vector3_linear_interpolate, cgodot_vector3_linear_interpolate_allocs = x.Vector3LinearInterpolate.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_linear_interpolate_allocs)

	var cgodot_vector3_cubic_interpolate_allocs *cgoAllocMap
	ref57717e51.godot_vector3_cubic_interpolate, cgodot_vector3_cubic_interpolate_allocs = x.Vector3CubicInterpolate.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_cubic_interpolate_allocs)

	var cgodot_vector3_dot_allocs *cgoAllocMap
	ref57717e51.godot_vector3_dot, cgodot_vector3_dot_allocs = x.Vector3Dot.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_dot_allocs)

	var cgodot_vector3_cross_allocs *cgoAllocMap
	ref57717e51.godot_vector3_cross, cgodot_vector3_cross_allocs = x.Vector3Cross.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_cross_allocs)

	var cgodot_vector3_outer_allocs *cgoAllocMap
	ref57717e51.godot_vector3_outer, cgodot_vector3_outer_allocs = x.Vector3Outer.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_outer_allocs)

	var cgodot_vector3_to_diagonal_matrix_allocs *cgoAllocMap
	ref57717e51.godot_vector3_to_diagonal_matrix, cgodot_vector3_to_diagonal_matrix_allocs = x.Vector3ToDiagonalMatrix.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_to_diagonal_matrix_allocs)

	var cgodot_vector3_abs_allocs *cgoAllocMap
	ref57717e51.godot_vector3_abs, cgodot_vector3_abs_allocs = x.Vector3Abs.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_abs_allocs)

	var cgodot_vector3_floor_allocs *cgoAllocMap
	ref57717e51.godot_vector3_floor, cgodot_vector3_floor_allocs = x.Vector3Floor.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_floor_allocs)

	var cgodot_vector3_ceil_allocs *cgoAllocMap
	ref57717e51.godot_vector3_ceil, cgodot_vector3_ceil_allocs = x.Vector3Ceil.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_ceil_allocs)

	var cgodot_vector3_distance_to_allocs *cgoAllocMap
	ref57717e51.godot_vector3_distance_to, cgodot_vector3_distance_to_allocs = x.Vector3DistanceTo.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_distance_to_allocs)

	var cgodot_vector3_distance_squared_to_allocs *cgoAllocMap
	ref57717e51.godot_vector3_distance_squared_to, cgodot_vector3_distance_squared_to_allocs = x.Vector3DistanceSquaredTo.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_distance_squared_to_allocs)

	var cgodot_vector3_angle_to_allocs *cgoAllocMap
	ref57717e51.godot_vector3_angle_to, cgodot_vector3_angle_to_allocs = x.Vector3AngleTo.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_angle_to_allocs)

	var cgodot_vector3_slide_allocs *cgoAllocMap
	ref57717e51.godot_vector3_slide, cgodot_vector3_slide_allocs = x.Vector3Slide.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_slide_allocs)

	var cgodot_vector3_bounce_allocs *cgoAllocMap
	ref57717e51.godot_vector3_bounce, cgodot_vector3_bounce_allocs = x.Vector3Bounce.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_bounce_allocs)

	var cgodot_vector3_reflect_allocs *cgoAllocMap
	ref57717e51.godot_vector3_reflect, cgodot_vector3_reflect_allocs = x.Vector3Reflect.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_reflect_allocs)

	var cgodot_vector3_operator_add_allocs *cgoAllocMap
	ref57717e51.godot_vector3_operator_add, cgodot_vector3_operator_add_allocs = x.Vector3OperatorAdd.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_operator_add_allocs)

	var cgodot_vector3_operator_substract_allocs *cgoAllocMap
	ref57717e51.godot_vector3_operator_substract, cgodot_vector3_operator_substract_allocs = x.Vector3OperatorSubstract.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_operator_substract_allocs)

	var cgodot_vector3_operator_multiply_vector_allocs *cgoAllocMap
	ref57717e51.godot_vector3_operator_multiply_vector, cgodot_vector3_operator_multiply_vector_allocs = x.Vector3OperatorMultiplyVector.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_operator_multiply_vector_allocs)

	var cgodot_vector3_operator_multiply_scalar_allocs *cgoAllocMap
	ref57717e51.godot_vector3_operator_multiply_scalar, cgodot_vector3_operator_multiply_scalar_allocs = x.Vector3OperatorMultiplyScalar.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_operator_multiply_scalar_allocs)

	var cgodot_vector3_operator_divide_vector_allocs *cgoAllocMap
	ref57717e51.godot_vector3_operator_divide_vector, cgodot_vector3_operator_divide_vector_allocs = x.Vector3OperatorDivideVector.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_operator_divide_vector_allocs)

	var cgodot_vector3_operator_divide_scalar_allocs *cgoAllocMap
	ref57717e51.godot_vector3_operator_divide_scalar, cgodot_vector3_operator_divide_scalar_allocs = x.Vector3OperatorDivideScalar.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_operator_divide_scalar_allocs)

	var cgodot_vector3_operator_equal_allocs *cgoAllocMap
	ref57717e51.godot_vector3_operator_equal, cgodot_vector3_operator_equal_allocs = x.Vector3OperatorEqual.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_operator_equal_allocs)

	var cgodot_vector3_operator_less_allocs *cgoAllocMap
	ref57717e51.godot_vector3_operator_less, cgodot_vector3_operator_less_allocs = x.Vector3OperatorLess.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_operator_less_allocs)

	var cgodot_vector3_operator_neg_allocs *cgoAllocMap
	ref57717e51.godot_vector3_operator_neg, cgodot_vector3_operator_neg_allocs = x.Vector3OperatorNeg.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_operator_neg_allocs)

	var cgodot_vector3_set_axis_allocs *cgoAllocMap
	ref57717e51.godot_vector3_set_axis, cgodot_vector3_set_axis_allocs = x.Vector3SetAxis.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_set_axis_allocs)

	var cgodot_vector3_get_axis_allocs *cgoAllocMap
	ref57717e51.godot_vector3_get_axis, cgodot_vector3_get_axis_allocs = x.Vector3GetAxis.PassRef()
	allocs57717e51.Borrow(cgodot_vector3_get_axis_allocs)

	var cgodot_pool_byte_array_new_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_new, cgodot_pool_byte_array_new_allocs = x.PoolByteArrayNew.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_new_allocs)

	var cgodot_pool_byte_array_new_copy_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_new_copy, cgodot_pool_byte_array_new_copy_allocs = x.PoolByteArrayNewCopy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_new_copy_allocs)

	var cgodot_pool_byte_array_new_with_array_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_new_with_array, cgodot_pool_byte_array_new_with_array_allocs = x.PoolByteArrayNewWithArray.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_new_with_array_allocs)

	var cgodot_pool_byte_array_append_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_append, cgodot_pool_byte_array_append_allocs = x.PoolByteArrayAppend.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_append_allocs)

	var cgodot_pool_byte_array_append_array_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_append_array, cgodot_pool_byte_array_append_array_allocs = x.PoolByteArrayAppendArray.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_append_array_allocs)

	var cgodot_pool_byte_array_insert_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_insert, cgodot_pool_byte_array_insert_allocs = x.PoolByteArrayInsert.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_insert_allocs)

	var cgodot_pool_byte_array_invert_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_invert, cgodot_pool_byte_array_invert_allocs = x.PoolByteArrayInvert.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_invert_allocs)

	var cgodot_pool_byte_array_push_back_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_push_back, cgodot_pool_byte_array_push_back_allocs = x.PoolByteArrayPushBack.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_push_back_allocs)

	var cgodot_pool_byte_array_remove_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_remove, cgodot_pool_byte_array_remove_allocs = x.PoolByteArrayRemove.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_remove_allocs)

	var cgodot_pool_byte_array_resize_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_resize, cgodot_pool_byte_array_resize_allocs = x.PoolByteArrayResize.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_resize_allocs)

	var cgodot_pool_byte_array_read_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_read, cgodot_pool_byte_array_read_allocs = x.PoolByteArrayRead.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_read_allocs)

	var cgodot_pool_byte_array_write_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_write, cgodot_pool_byte_array_write_allocs = x.PoolByteArrayWrite.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_write_allocs)

	var cgodot_pool_byte_array_set_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_set, cgodot_pool_byte_array_set_allocs = x.PoolByteArraySet.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_set_allocs)

	var cgodot_pool_byte_array_get_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_get, cgodot_pool_byte_array_get_allocs = x.PoolByteArrayGet.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_get_allocs)

	var cgodot_pool_byte_array_size_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_size, cgodot_pool_byte_array_size_allocs = x.PoolByteArraySize.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_size_allocs)

	var cgodot_pool_byte_array_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_destroy, cgodot_pool_byte_array_destroy_allocs = x.PoolByteArrayDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_destroy_allocs)

	var cgodot_pool_int_array_new_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_new, cgodot_pool_int_array_new_allocs = x.PoolIntArrayNew.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_new_allocs)

	var cgodot_pool_int_array_new_copy_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_new_copy, cgodot_pool_int_array_new_copy_allocs = x.PoolIntArrayNewCopy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_new_copy_allocs)

	var cgodot_pool_int_array_new_with_array_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_new_with_array, cgodot_pool_int_array_new_with_array_allocs = x.PoolIntArrayNewWithArray.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_new_with_array_allocs)

	var cgodot_pool_int_array_append_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_append, cgodot_pool_int_array_append_allocs = x.PoolIntArrayAppend.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_append_allocs)

	var cgodot_pool_int_array_append_array_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_append_array, cgodot_pool_int_array_append_array_allocs = x.PoolIntArrayAppendArray.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_append_array_allocs)

	var cgodot_pool_int_array_insert_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_insert, cgodot_pool_int_array_insert_allocs = x.PoolIntArrayInsert.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_insert_allocs)

	var cgodot_pool_int_array_invert_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_invert, cgodot_pool_int_array_invert_allocs = x.PoolIntArrayInvert.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_invert_allocs)

	var cgodot_pool_int_array_push_back_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_push_back, cgodot_pool_int_array_push_back_allocs = x.PoolIntArrayPushBack.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_push_back_allocs)

	var cgodot_pool_int_array_remove_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_remove, cgodot_pool_int_array_remove_allocs = x.PoolIntArrayRemove.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_remove_allocs)

	var cgodot_pool_int_array_resize_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_resize, cgodot_pool_int_array_resize_allocs = x.PoolIntArrayResize.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_resize_allocs)

	var cgodot_pool_int_array_read_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_read, cgodot_pool_int_array_read_allocs = x.PoolIntArrayRead.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_read_allocs)

	var cgodot_pool_int_array_write_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_write, cgodot_pool_int_array_write_allocs = x.PoolIntArrayWrite.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_write_allocs)

	var cgodot_pool_int_array_set_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_set, cgodot_pool_int_array_set_allocs = x.PoolIntArraySet.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_set_allocs)

	var cgodot_pool_int_array_get_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_get, cgodot_pool_int_array_get_allocs = x.PoolIntArrayGet.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_get_allocs)

	var cgodot_pool_int_array_size_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_size, cgodot_pool_int_array_size_allocs = x.PoolIntArraySize.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_size_allocs)

	var cgodot_pool_int_array_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_destroy, cgodot_pool_int_array_destroy_allocs = x.PoolIntArrayDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_destroy_allocs)

	var cgodot_pool_real_array_new_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_new, cgodot_pool_real_array_new_allocs = x.PoolRealArrayNew.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_new_allocs)

	var cgodot_pool_real_array_new_copy_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_new_copy, cgodot_pool_real_array_new_copy_allocs = x.PoolRealArrayNewCopy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_new_copy_allocs)

	var cgodot_pool_real_array_new_with_array_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_new_with_array, cgodot_pool_real_array_new_with_array_allocs = x.PoolRealArrayNewWithArray.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_new_with_array_allocs)

	var cgodot_pool_real_array_append_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_append, cgodot_pool_real_array_append_allocs = x.PoolRealArrayAppend.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_append_allocs)

	var cgodot_pool_real_array_append_array_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_append_array, cgodot_pool_real_array_append_array_allocs = x.PoolRealArrayAppendArray.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_append_array_allocs)

	var cgodot_pool_real_array_insert_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_insert, cgodot_pool_real_array_insert_allocs = x.PoolRealArrayInsert.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_insert_allocs)

	var cgodot_pool_real_array_invert_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_invert, cgodot_pool_real_array_invert_allocs = x.PoolRealArrayInvert.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_invert_allocs)

	var cgodot_pool_real_array_push_back_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_push_back, cgodot_pool_real_array_push_back_allocs = x.PoolRealArrayPushBack.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_push_back_allocs)

	var cgodot_pool_real_array_remove_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_remove, cgodot_pool_real_array_remove_allocs = x.PoolRealArrayRemove.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_remove_allocs)

	var cgodot_pool_real_array_resize_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_resize, cgodot_pool_real_array_resize_allocs = x.PoolRealArrayResize.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_resize_allocs)

	var cgodot_pool_real_array_read_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_read, cgodot_pool_real_array_read_allocs = x.PoolRealArrayRead.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_read_allocs)

	var cgodot_pool_real_array_write_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_write, cgodot_pool_real_array_write_allocs = x.PoolRealArrayWrite.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_write_allocs)

	var cgodot_pool_real_array_set_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_set, cgodot_pool_real_array_set_allocs = x.PoolRealArraySet.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_set_allocs)

	var cgodot_pool_real_array_get_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_get, cgodot_pool_real_array_get_allocs = x.PoolRealArrayGet.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_get_allocs)

	var cgodot_pool_real_array_size_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_size, cgodot_pool_real_array_size_allocs = x.PoolRealArraySize.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_size_allocs)

	var cgodot_pool_real_array_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_destroy, cgodot_pool_real_array_destroy_allocs = x.PoolRealArrayDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_destroy_allocs)

	var cgodot_pool_string_array_new_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_new, cgodot_pool_string_array_new_allocs = x.PoolStringArrayNew.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_new_allocs)

	var cgodot_pool_string_array_new_copy_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_new_copy, cgodot_pool_string_array_new_copy_allocs = x.PoolStringArrayNewCopy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_new_copy_allocs)

	var cgodot_pool_string_array_new_with_array_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_new_with_array, cgodot_pool_string_array_new_with_array_allocs = x.PoolStringArrayNewWithArray.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_new_with_array_allocs)

	var cgodot_pool_string_array_append_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_append, cgodot_pool_string_array_append_allocs = x.PoolStringArrayAppend.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_append_allocs)

	var cgodot_pool_string_array_append_array_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_append_array, cgodot_pool_string_array_append_array_allocs = x.PoolStringArrayAppendArray.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_append_array_allocs)

	var cgodot_pool_string_array_insert_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_insert, cgodot_pool_string_array_insert_allocs = x.PoolStringArrayInsert.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_insert_allocs)

	var cgodot_pool_string_array_invert_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_invert, cgodot_pool_string_array_invert_allocs = x.PoolStringArrayInvert.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_invert_allocs)

	var cgodot_pool_string_array_push_back_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_push_back, cgodot_pool_string_array_push_back_allocs = x.PoolStringArrayPushBack.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_push_back_allocs)

	var cgodot_pool_string_array_remove_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_remove, cgodot_pool_string_array_remove_allocs = x.PoolStringArrayRemove.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_remove_allocs)

	var cgodot_pool_string_array_resize_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_resize, cgodot_pool_string_array_resize_allocs = x.PoolStringArrayResize.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_resize_allocs)

	var cgodot_pool_string_array_read_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_read, cgodot_pool_string_array_read_allocs = x.PoolStringArrayRead.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_read_allocs)

	var cgodot_pool_string_array_write_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_write, cgodot_pool_string_array_write_allocs = x.PoolStringArrayWrite.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_write_allocs)

	var cgodot_pool_string_array_set_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_set, cgodot_pool_string_array_set_allocs = x.PoolStringArraySet.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_set_allocs)

	var cgodot_pool_string_array_get_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_get, cgodot_pool_string_array_get_allocs = x.PoolStringArrayGet.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_get_allocs)

	var cgodot_pool_string_array_size_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_size, cgodot_pool_string_array_size_allocs = x.PoolStringArraySize.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_size_allocs)

	var cgodot_pool_string_array_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_destroy, cgodot_pool_string_array_destroy_allocs = x.PoolStringArrayDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_destroy_allocs)

	var cgodot_pool_vector2_array_new_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_new, cgodot_pool_vector2_array_new_allocs = x.PoolVector2ArrayNew.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_new_allocs)

	var cgodot_pool_vector2_array_new_copy_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_new_copy, cgodot_pool_vector2_array_new_copy_allocs = x.PoolVector2ArrayNewCopy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_new_copy_allocs)

	var cgodot_pool_vector2_array_new_with_array_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_new_with_array, cgodot_pool_vector2_array_new_with_array_allocs = x.PoolVector2ArrayNewWithArray.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_new_with_array_allocs)

	var cgodot_pool_vector2_array_append_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_append, cgodot_pool_vector2_array_append_allocs = x.PoolVector2ArrayAppend.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_append_allocs)

	var cgodot_pool_vector2_array_append_array_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_append_array, cgodot_pool_vector2_array_append_array_allocs = x.PoolVector2ArrayAppendArray.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_append_array_allocs)

	var cgodot_pool_vector2_array_insert_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_insert, cgodot_pool_vector2_array_insert_allocs = x.PoolVector2ArrayInsert.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_insert_allocs)

	var cgodot_pool_vector2_array_invert_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_invert, cgodot_pool_vector2_array_invert_allocs = x.PoolVector2ArrayInvert.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_invert_allocs)

	var cgodot_pool_vector2_array_push_back_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_push_back, cgodot_pool_vector2_array_push_back_allocs = x.PoolVector2ArrayPushBack.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_push_back_allocs)

	var cgodot_pool_vector2_array_remove_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_remove, cgodot_pool_vector2_array_remove_allocs = x.PoolVector2ArrayRemove.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_remove_allocs)

	var cgodot_pool_vector2_array_resize_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_resize, cgodot_pool_vector2_array_resize_allocs = x.PoolVector2ArrayResize.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_resize_allocs)

	var cgodot_pool_vector2_array_read_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_read, cgodot_pool_vector2_array_read_allocs = x.PoolVector2ArrayRead.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_read_allocs)

	var cgodot_pool_vector2_array_write_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_write, cgodot_pool_vector2_array_write_allocs = x.PoolVector2ArrayWrite.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_write_allocs)

	var cgodot_pool_vector2_array_set_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_set, cgodot_pool_vector2_array_set_allocs = x.PoolVector2ArraySet.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_set_allocs)

	var cgodot_pool_vector2_array_get_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_get, cgodot_pool_vector2_array_get_allocs = x.PoolVector2ArrayGet.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_get_allocs)

	var cgodot_pool_vector2_array_size_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_size, cgodot_pool_vector2_array_size_allocs = x.PoolVector2ArraySize.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_size_allocs)

	var cgodot_pool_vector2_array_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_destroy, cgodot_pool_vector2_array_destroy_allocs = x.PoolVector2ArrayDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_destroy_allocs)

	var cgodot_pool_vector3_array_new_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_new, cgodot_pool_vector3_array_new_allocs = x.PoolVector3ArrayNew.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_new_allocs)

	var cgodot_pool_vector3_array_new_copy_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_new_copy, cgodot_pool_vector3_array_new_copy_allocs = x.PoolVector3ArrayNewCopy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_new_copy_allocs)

	var cgodot_pool_vector3_array_new_with_array_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_new_with_array, cgodot_pool_vector3_array_new_with_array_allocs = x.PoolVector3ArrayNewWithArray.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_new_with_array_allocs)

	var cgodot_pool_vector3_array_append_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_append, cgodot_pool_vector3_array_append_allocs = x.PoolVector3ArrayAppend.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_append_allocs)

	var cgodot_pool_vector3_array_append_array_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_append_array, cgodot_pool_vector3_array_append_array_allocs = x.PoolVector3ArrayAppendArray.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_append_array_allocs)

	var cgodot_pool_vector3_array_insert_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_insert, cgodot_pool_vector3_array_insert_allocs = x.PoolVector3ArrayInsert.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_insert_allocs)

	var cgodot_pool_vector3_array_invert_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_invert, cgodot_pool_vector3_array_invert_allocs = x.PoolVector3ArrayInvert.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_invert_allocs)

	var cgodot_pool_vector3_array_push_back_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_push_back, cgodot_pool_vector3_array_push_back_allocs = x.PoolVector3ArrayPushBack.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_push_back_allocs)

	var cgodot_pool_vector3_array_remove_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_remove, cgodot_pool_vector3_array_remove_allocs = x.PoolVector3ArrayRemove.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_remove_allocs)

	var cgodot_pool_vector3_array_resize_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_resize, cgodot_pool_vector3_array_resize_allocs = x.PoolVector3ArrayResize.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_resize_allocs)

	var cgodot_pool_vector3_array_read_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_read, cgodot_pool_vector3_array_read_allocs = x.PoolVector3ArrayRead.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_read_allocs)

	var cgodot_pool_vector3_array_write_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_write, cgodot_pool_vector3_array_write_allocs = x.PoolVector3ArrayWrite.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_write_allocs)

	var cgodot_pool_vector3_array_set_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_set, cgodot_pool_vector3_array_set_allocs = x.PoolVector3ArraySet.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_set_allocs)

	var cgodot_pool_vector3_array_get_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_get, cgodot_pool_vector3_array_get_allocs = x.PoolVector3ArrayGet.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_get_allocs)

	var cgodot_pool_vector3_array_size_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_size, cgodot_pool_vector3_array_size_allocs = x.PoolVector3ArraySize.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_size_allocs)

	var cgodot_pool_vector3_array_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_destroy, cgodot_pool_vector3_array_destroy_allocs = x.PoolVector3ArrayDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_destroy_allocs)

	var cgodot_pool_color_array_new_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_new, cgodot_pool_color_array_new_allocs = x.PoolColorArrayNew.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_new_allocs)

	var cgodot_pool_color_array_new_copy_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_new_copy, cgodot_pool_color_array_new_copy_allocs = x.PoolColorArrayNewCopy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_new_copy_allocs)

	var cgodot_pool_color_array_new_with_array_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_new_with_array, cgodot_pool_color_array_new_with_array_allocs = x.PoolColorArrayNewWithArray.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_new_with_array_allocs)

	var cgodot_pool_color_array_append_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_append, cgodot_pool_color_array_append_allocs = x.PoolColorArrayAppend.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_append_allocs)

	var cgodot_pool_color_array_append_array_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_append_array, cgodot_pool_color_array_append_array_allocs = x.PoolColorArrayAppendArray.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_append_array_allocs)

	var cgodot_pool_color_array_insert_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_insert, cgodot_pool_color_array_insert_allocs = x.PoolColorArrayInsert.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_insert_allocs)

	var cgodot_pool_color_array_invert_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_invert, cgodot_pool_color_array_invert_allocs = x.PoolColorArrayInvert.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_invert_allocs)

	var cgodot_pool_color_array_push_back_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_push_back, cgodot_pool_color_array_push_back_allocs = x.PoolColorArrayPushBack.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_push_back_allocs)

	var cgodot_pool_color_array_remove_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_remove, cgodot_pool_color_array_remove_allocs = x.PoolColorArrayRemove.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_remove_allocs)

	var cgodot_pool_color_array_resize_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_resize, cgodot_pool_color_array_resize_allocs = x.PoolColorArrayResize.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_resize_allocs)

	var cgodot_pool_color_array_read_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_read, cgodot_pool_color_array_read_allocs = x.PoolColorArrayRead.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_read_allocs)

	var cgodot_pool_color_array_write_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_write, cgodot_pool_color_array_write_allocs = x.PoolColorArrayWrite.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_write_allocs)

	var cgodot_pool_color_array_set_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_set, cgodot_pool_color_array_set_allocs = x.PoolColorArraySet.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_set_allocs)

	var cgodot_pool_color_array_get_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_get, cgodot_pool_color_array_get_allocs = x.PoolColorArrayGet.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_get_allocs)

	var cgodot_pool_color_array_size_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_size, cgodot_pool_color_array_size_allocs = x.PoolColorArraySize.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_size_allocs)

	var cgodot_pool_color_array_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_destroy, cgodot_pool_color_array_destroy_allocs = x.PoolColorArrayDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_destroy_allocs)

	var cgodot_pool_byte_array_read_access_ptr_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_read_access_ptr, cgodot_pool_byte_array_read_access_ptr_allocs = x.PoolByteArrayReadAccessPtr.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_read_access_ptr_allocs)

	var cgodot_pool_byte_array_read_access_operator_assign_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_read_access_operator_assign, cgodot_pool_byte_array_read_access_operator_assign_allocs = x.PoolByteArrayReadAccessOperatorAssign.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_read_access_operator_assign_allocs)

	var cgodot_pool_byte_array_read_access_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_read_access_destroy, cgodot_pool_byte_array_read_access_destroy_allocs = x.PoolByteArrayReadAccessDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_read_access_destroy_allocs)

	var cgodot_pool_int_array_read_access_ptr_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_read_access_ptr, cgodot_pool_int_array_read_access_ptr_allocs = x.PoolIntArrayReadAccessPtr.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_read_access_ptr_allocs)

	var cgodot_pool_int_array_read_access_operator_assign_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_read_access_operator_assign, cgodot_pool_int_array_read_access_operator_assign_allocs = x.PoolIntArrayReadAccessOperatorAssign.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_read_access_operator_assign_allocs)

	var cgodot_pool_int_array_read_access_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_read_access_destroy, cgodot_pool_int_array_read_access_destroy_allocs = x.PoolIntArrayReadAccessDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_read_access_destroy_allocs)

	var cgodot_pool_real_array_read_access_ptr_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_read_access_ptr, cgodot_pool_real_array_read_access_ptr_allocs = x.PoolRealArrayReadAccessPtr.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_read_access_ptr_allocs)

	var cgodot_pool_real_array_read_access_operator_assign_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_read_access_operator_assign, cgodot_pool_real_array_read_access_operator_assign_allocs = x.PoolRealArrayReadAccessOperatorAssign.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_read_access_operator_assign_allocs)

	var cgodot_pool_real_array_read_access_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_read_access_destroy, cgodot_pool_real_array_read_access_destroy_allocs = x.PoolRealArrayReadAccessDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_read_access_destroy_allocs)

	var cgodot_pool_string_array_read_access_ptr_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_read_access_ptr, cgodot_pool_string_array_read_access_ptr_allocs = x.PoolStringArrayReadAccessPtr.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_read_access_ptr_allocs)

	var cgodot_pool_string_array_read_access_operator_assign_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_read_access_operator_assign, cgodot_pool_string_array_read_access_operator_assign_allocs = x.PoolStringArrayReadAccessOperatorAssign.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_read_access_operator_assign_allocs)

	var cgodot_pool_string_array_read_access_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_read_access_destroy, cgodot_pool_string_array_read_access_destroy_allocs = x.PoolStringArrayReadAccessDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_read_access_destroy_allocs)

	var cgodot_pool_vector2_array_read_access_ptr_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_read_access_ptr, cgodot_pool_vector2_array_read_access_ptr_allocs = x.PoolVector2ArrayReadAccessPtr.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_read_access_ptr_allocs)

	var cgodot_pool_vector2_array_read_access_operator_assign_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_read_access_operator_assign, cgodot_pool_vector2_array_read_access_operator_assign_allocs = x.PoolVector2ArrayReadAccessOperatorAssign.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_read_access_operator_assign_allocs)

	var cgodot_pool_vector2_array_read_access_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_read_access_destroy, cgodot_pool_vector2_array_read_access_destroy_allocs = x.PoolVector2ArrayReadAccessDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_read_access_destroy_allocs)

	var cgodot_pool_vector3_array_read_access_ptr_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_read_access_ptr, cgodot_pool_vector3_array_read_access_ptr_allocs = x.PoolVector3ArrayReadAccessPtr.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_read_access_ptr_allocs)

	var cgodot_pool_vector3_array_read_access_operator_assign_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_read_access_operator_assign, cgodot_pool_vector3_array_read_access_operator_assign_allocs = x.PoolVector3ArrayReadAccessOperatorAssign.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_read_access_operator_assign_allocs)

	var cgodot_pool_vector3_array_read_access_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_read_access_destroy, cgodot_pool_vector3_array_read_access_destroy_allocs = x.PoolVector3ArrayReadAccessDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_read_access_destroy_allocs)

	var cgodot_pool_color_array_read_access_ptr_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_read_access_ptr, cgodot_pool_color_array_read_access_ptr_allocs = x.PoolColorArrayReadAccessPtr.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_read_access_ptr_allocs)

	var cgodot_pool_color_array_read_access_operator_assign_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_read_access_operator_assign, cgodot_pool_color_array_read_access_operator_assign_allocs = x.PoolColorArrayReadAccessOperatorAssign.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_read_access_operator_assign_allocs)

	var cgodot_pool_color_array_read_access_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_read_access_destroy, cgodot_pool_color_array_read_access_destroy_allocs = x.PoolColorArrayReadAccessDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_read_access_destroy_allocs)

	var cgodot_pool_byte_array_write_access_ptr_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_write_access_ptr, cgodot_pool_byte_array_write_access_ptr_allocs = x.PoolByteArrayWriteAccessPtr.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_write_access_ptr_allocs)

	var cgodot_pool_byte_array_write_access_operator_assign_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_write_access_operator_assign, cgodot_pool_byte_array_write_access_operator_assign_allocs = x.PoolByteArrayWriteAccessOperatorAssign.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_write_access_operator_assign_allocs)

	var cgodot_pool_byte_array_write_access_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_byte_array_write_access_destroy, cgodot_pool_byte_array_write_access_destroy_allocs = x.PoolByteArrayWriteAccessDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_byte_array_write_access_destroy_allocs)

	var cgodot_pool_int_array_write_access_ptr_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_write_access_ptr, cgodot_pool_int_array_write_access_ptr_allocs = x.PoolIntArrayWriteAccessPtr.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_write_access_ptr_allocs)

	var cgodot_pool_int_array_write_access_operator_assign_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_write_access_operator_assign, cgodot_pool_int_array_write_access_operator_assign_allocs = x.PoolIntArrayWriteAccessOperatorAssign.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_write_access_operator_assign_allocs)

	var cgodot_pool_int_array_write_access_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_int_array_write_access_destroy, cgodot_pool_int_array_write_access_destroy_allocs = x.PoolIntArrayWriteAccessDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_int_array_write_access_destroy_allocs)

	var cgodot_pool_real_array_write_access_ptr_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_write_access_ptr, cgodot_pool_real_array_write_access_ptr_allocs = x.PoolRealArrayWriteAccessPtr.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_write_access_ptr_allocs)

	var cgodot_pool_real_array_write_access_operator_assign_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_write_access_operator_assign, cgodot_pool_real_array_write_access_operator_assign_allocs = x.PoolRealArrayWriteAccessOperatorAssign.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_write_access_operator_assign_allocs)

	var cgodot_pool_real_array_write_access_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_real_array_write_access_destroy, cgodot_pool_real_array_write_access_destroy_allocs = x.PoolRealArrayWriteAccessDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_real_array_write_access_destroy_allocs)

	var cgodot_pool_string_array_write_access_ptr_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_write_access_ptr, cgodot_pool_string_array_write_access_ptr_allocs = x.PoolStringArrayWriteAccessPtr.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_write_access_ptr_allocs)

	var cgodot_pool_string_array_write_access_operator_assign_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_write_access_operator_assign, cgodot_pool_string_array_write_access_operator_assign_allocs = x.PoolStringArrayWriteAccessOperatorAssign.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_write_access_operator_assign_allocs)

	var cgodot_pool_string_array_write_access_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_string_array_write_access_destroy, cgodot_pool_string_array_write_access_destroy_allocs = x.PoolStringArrayWriteAccessDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_string_array_write_access_destroy_allocs)

	var cgodot_pool_vector2_array_write_access_ptr_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_write_access_ptr, cgodot_pool_vector2_array_write_access_ptr_allocs = x.PoolVector2ArrayWriteAccessPtr.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_write_access_ptr_allocs)

	var cgodot_pool_vector2_array_write_access_operator_assign_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_write_access_operator_assign, cgodot_pool_vector2_array_write_access_operator_assign_allocs = x.PoolVector2ArrayWriteAccessOperatorAssign.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_write_access_operator_assign_allocs)

	var cgodot_pool_vector2_array_write_access_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector2_array_write_access_destroy, cgodot_pool_vector2_array_write_access_destroy_allocs = x.PoolVector2ArrayWriteAccessDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector2_array_write_access_destroy_allocs)

	var cgodot_pool_vector3_array_write_access_ptr_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_write_access_ptr, cgodot_pool_vector3_array_write_access_ptr_allocs = x.PoolVector3ArrayWriteAccessPtr.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_write_access_ptr_allocs)

	var cgodot_pool_vector3_array_write_access_operator_assign_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_write_access_operator_assign, cgodot_pool_vector3_array_write_access_operator_assign_allocs = x.PoolVector3ArrayWriteAccessOperatorAssign.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_write_access_operator_assign_allocs)

	var cgodot_pool_vector3_array_write_access_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_vector3_array_write_access_destroy, cgodot_pool_vector3_array_write_access_destroy_allocs = x.PoolVector3ArrayWriteAccessDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_vector3_array_write_access_destroy_allocs)

	var cgodot_pool_color_array_write_access_ptr_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_write_access_ptr, cgodot_pool_color_array_write_access_ptr_allocs = x.PoolColorArrayWriteAccessPtr.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_write_access_ptr_allocs)

	var cgodot_pool_color_array_write_access_operator_assign_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_write_access_operator_assign, cgodot_pool_color_array_write_access_operator_assign_allocs = x.PoolColorArrayWriteAccessOperatorAssign.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_write_access_operator_assign_allocs)

	var cgodot_pool_color_array_write_access_destroy_allocs *cgoAllocMap
	ref57717e51.godot_pool_color_array_write_access_destroy, cgodot_pool_color_array_write_access_destroy_allocs = x.PoolColorArrayWriteAccessDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_pool_color_array_write_access_destroy_allocs)

	var cgodot_array_new_allocs *cgoAllocMap
	ref57717e51.godot_array_new, cgodot_array_new_allocs = x.ArrayNew.PassRef()
	allocs57717e51.Borrow(cgodot_array_new_allocs)

	var cgodot_array_new_copy_allocs *cgoAllocMap
	ref57717e51.godot_array_new_copy, cgodot_array_new_copy_allocs = x.ArrayNewCopy.PassRef()
	allocs57717e51.Borrow(cgodot_array_new_copy_allocs)

	var cgodot_array_new_pool_color_array_allocs *cgoAllocMap
	ref57717e51.godot_array_new_pool_color_array, cgodot_array_new_pool_color_array_allocs = x.ArrayNewPoolColorArray.PassRef()
	allocs57717e51.Borrow(cgodot_array_new_pool_color_array_allocs)

	var cgodot_array_new_pool_vector3_array_allocs *cgoAllocMap
	ref57717e51.godot_array_new_pool_vector3_array, cgodot_array_new_pool_vector3_array_allocs = x.ArrayNewPoolVector3Array.PassRef()
	allocs57717e51.Borrow(cgodot_array_new_pool_vector3_array_allocs)

	var cgodot_array_new_pool_vector2_array_allocs *cgoAllocMap
	ref57717e51.godot_array_new_pool_vector2_array, cgodot_array_new_pool_vector2_array_allocs = x.ArrayNewPoolVector2Array.PassRef()
	allocs57717e51.Borrow(cgodot_array_new_pool_vector2_array_allocs)

	var cgodot_array_new_pool_string_array_allocs *cgoAllocMap
	ref57717e51.godot_array_new_pool_string_array, cgodot_array_new_pool_string_array_allocs = x.ArrayNewPoolStringArray.PassRef()
	allocs57717e51.Borrow(cgodot_array_new_pool_string_array_allocs)

	var cgodot_array_new_pool_real_array_allocs *cgoAllocMap
	ref57717e51.godot_array_new_pool_real_array, cgodot_array_new_pool_real_array_allocs = x.ArrayNewPoolRealArray.PassRef()
	allocs57717e51.Borrow(cgodot_array_new_pool_real_array_allocs)

	var cgodot_array_new_pool_int_array_allocs *cgoAllocMap
	ref57717e51.godot_array_new_pool_int_array, cgodot_array_new_pool_int_array_allocs = x.ArrayNewPoolIntArray.PassRef()
	allocs57717e51.Borrow(cgodot_array_new_pool_int_array_allocs)

	var cgodot_array_new_pool_byte_array_allocs *cgoAllocMap
	ref57717e51.godot_array_new_pool_byte_array, cgodot_array_new_pool_byte_array_allocs = x.ArrayNewPoolByteArray.PassRef()
	allocs57717e51.Borrow(cgodot_array_new_pool_byte_array_allocs)

	var cgodot_array_set_allocs *cgoAllocMap
	ref57717e51.godot_array_set, cgodot_array_set_allocs = x.ArraySet.PassRef()
	allocs57717e51.Borrow(cgodot_array_set_allocs)

	var cgodot_array_get_allocs *cgoAllocMap
	ref57717e51.godot_array_get, cgodot_array_get_allocs = x.ArrayGet.PassRef()
	allocs57717e51.Borrow(cgodot_array_get_allocs)

	var cgodot_array_operator_index_allocs *cgoAllocMap
	ref57717e51.godot_array_operator_index, cgodot_array_operator_index_allocs = x.ArrayOperatorIndex.PassRef()
	allocs57717e51.Borrow(cgodot_array_operator_index_allocs)

	var cgodot_array_operator_index_const_allocs *cgoAllocMap
	ref57717e51.godot_array_operator_index_const, cgodot_array_operator_index_const_allocs = x.ArrayOperatorIndexConst.PassRef()
	allocs57717e51.Borrow(cgodot_array_operator_index_const_allocs)

	var cgodot_array_append_allocs *cgoAllocMap
	ref57717e51.godot_array_append, cgodot_array_append_allocs = x.ArrayAppend.PassRef()
	allocs57717e51.Borrow(cgodot_array_append_allocs)

	var cgodot_array_clear_allocs *cgoAllocMap
	ref57717e51.godot_array_clear, cgodot_array_clear_allocs = x.ArrayClear.PassRef()
	allocs57717e51.Borrow(cgodot_array_clear_allocs)

	var cgodot_array_count_allocs *cgoAllocMap
	ref57717e51.godot_array_count, cgodot_array_count_allocs = x.ArrayCount.PassRef()
	allocs57717e51.Borrow(cgodot_array_count_allocs)

	var cgodot_array_empty_allocs *cgoAllocMap
	ref57717e51.godot_array_empty, cgodot_array_empty_allocs = x.ArrayEmpty.PassRef()
	allocs57717e51.Borrow(cgodot_array_empty_allocs)

	var cgodot_array_erase_allocs *cgoAllocMap
	ref57717e51.godot_array_erase, cgodot_array_erase_allocs = x.ArrayErase.PassRef()
	allocs57717e51.Borrow(cgodot_array_erase_allocs)

	var cgodot_array_front_allocs *cgoAllocMap
	ref57717e51.godot_array_front, cgodot_array_front_allocs = x.ArrayFront.PassRef()
	allocs57717e51.Borrow(cgodot_array_front_allocs)

	var cgodot_array_back_allocs *cgoAllocMap
	ref57717e51.godot_array_back, cgodot_array_back_allocs = x.ArrayBack.PassRef()
	allocs57717e51.Borrow(cgodot_array_back_allocs)

	var cgodot_array_find_allocs *cgoAllocMap
	ref57717e51.godot_array_find, cgodot_array_find_allocs = x.ArrayFind.PassRef()
	allocs57717e51.Borrow(cgodot_array_find_allocs)

	var cgodot_array_find_last_allocs *cgoAllocMap
	ref57717e51.godot_array_find_last, cgodot_array_find_last_allocs = x.ArrayFindLast.PassRef()
	allocs57717e51.Borrow(cgodot_array_find_last_allocs)

	var cgodot_array_has_allocs *cgoAllocMap
	ref57717e51.godot_array_has, cgodot_array_has_allocs = x.ArrayHas.PassRef()
	allocs57717e51.Borrow(cgodot_array_has_allocs)

	var cgodot_array_hash_allocs *cgoAllocMap
	ref57717e51.godot_array_hash, cgodot_array_hash_allocs = x.ArrayHash.PassRef()
	allocs57717e51.Borrow(cgodot_array_hash_allocs)

	var cgodot_array_insert_allocs *cgoAllocMap
	ref57717e51.godot_array_insert, cgodot_array_insert_allocs = x.ArrayInsert.PassRef()
	allocs57717e51.Borrow(cgodot_array_insert_allocs)

	var cgodot_array_invert_allocs *cgoAllocMap
	ref57717e51.godot_array_invert, cgodot_array_invert_allocs = x.ArrayInvert.PassRef()
	allocs57717e51.Borrow(cgodot_array_invert_allocs)

	var cgodot_array_pop_back_allocs *cgoAllocMap
	ref57717e51.godot_array_pop_back, cgodot_array_pop_back_allocs = x.ArrayPopBack.PassRef()
	allocs57717e51.Borrow(cgodot_array_pop_back_allocs)

	var cgodot_array_pop_front_allocs *cgoAllocMap
	ref57717e51.godot_array_pop_front, cgodot_array_pop_front_allocs = x.ArrayPopFront.PassRef()
	allocs57717e51.Borrow(cgodot_array_pop_front_allocs)

	var cgodot_array_push_back_allocs *cgoAllocMap
	ref57717e51.godot_array_push_back, cgodot_array_push_back_allocs = x.ArrayPushBack.PassRef()
	allocs57717e51.Borrow(cgodot_array_push_back_allocs)

	var cgodot_array_push_front_allocs *cgoAllocMap
	ref57717e51.godot_array_push_front, cgodot_array_push_front_allocs = x.ArrayPushFront.PassRef()
	allocs57717e51.Borrow(cgodot_array_push_front_allocs)

	var cgodot_array_remove_allocs *cgoAllocMap
	ref57717e51.godot_array_remove, cgodot_array_remove_allocs = x.ArrayRemove.PassRef()
	allocs57717e51.Borrow(cgodot_array_remove_allocs)

	var cgodot_array_resize_allocs *cgoAllocMap
	ref57717e51.godot_array_resize, cgodot_array_resize_allocs = x.ArrayResize.PassRef()
	allocs57717e51.Borrow(cgodot_array_resize_allocs)

	var cgodot_array_rfind_allocs *cgoAllocMap
	ref57717e51.godot_array_rfind, cgodot_array_rfind_allocs = x.ArrayRfind.PassRef()
	allocs57717e51.Borrow(cgodot_array_rfind_allocs)

	var cgodot_array_size_allocs *cgoAllocMap
	ref57717e51.godot_array_size, cgodot_array_size_allocs = x.ArraySize.PassRef()
	allocs57717e51.Borrow(cgodot_array_size_allocs)

	var cgodot_array_sort_allocs *cgoAllocMap
	ref57717e51.godot_array_sort, cgodot_array_sort_allocs = x.ArraySort.PassRef()
	allocs57717e51.Borrow(cgodot_array_sort_allocs)

	var cgodot_array_sort_custom_allocs *cgoAllocMap
	ref57717e51.godot_array_sort_custom, cgodot_array_sort_custom_allocs = x.ArraySortCustom.PassRef()
	allocs57717e51.Borrow(cgodot_array_sort_custom_allocs)

	var cgodot_array_bsearch_allocs *cgoAllocMap
	ref57717e51.godot_array_bsearch, cgodot_array_bsearch_allocs = x.ArrayBsearch.PassRef()
	allocs57717e51.Borrow(cgodot_array_bsearch_allocs)

	var cgodot_array_bsearch_custom_allocs *cgoAllocMap
	ref57717e51.godot_array_bsearch_custom, cgodot_array_bsearch_custom_allocs = x.ArrayBsearchCustom.PassRef()
	allocs57717e51.Borrow(cgodot_array_bsearch_custom_allocs)

	var cgodot_array_destroy_allocs *cgoAllocMap
	ref57717e51.godot_array_destroy, cgodot_array_destroy_allocs = x.ArrayDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_array_destroy_allocs)

	var cgodot_dictionary_new_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_new, cgodot_dictionary_new_allocs = x.DictionaryNew.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_new_allocs)

	var cgodot_dictionary_new_copy_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_new_copy, cgodot_dictionary_new_copy_allocs = x.DictionaryNewCopy.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_new_copy_allocs)

	var cgodot_dictionary_destroy_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_destroy, cgodot_dictionary_destroy_allocs = x.DictionaryDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_destroy_allocs)

	var cgodot_dictionary_size_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_size, cgodot_dictionary_size_allocs = x.DictionarySize.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_size_allocs)

	var cgodot_dictionary_empty_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_empty, cgodot_dictionary_empty_allocs = x.DictionaryEmpty.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_empty_allocs)

	var cgodot_dictionary_clear_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_clear, cgodot_dictionary_clear_allocs = x.DictionaryClear.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_clear_allocs)

	var cgodot_dictionary_has_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_has, cgodot_dictionary_has_allocs = x.DictionaryHas.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_has_allocs)

	var cgodot_dictionary_has_all_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_has_all, cgodot_dictionary_has_all_allocs = x.DictionaryHasAll.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_has_all_allocs)

	var cgodot_dictionary_erase_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_erase, cgodot_dictionary_erase_allocs = x.DictionaryErase.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_erase_allocs)

	var cgodot_dictionary_hash_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_hash, cgodot_dictionary_hash_allocs = x.DictionaryHash.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_hash_allocs)

	var cgodot_dictionary_keys_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_keys, cgodot_dictionary_keys_allocs = x.DictionaryKeys.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_keys_allocs)

	var cgodot_dictionary_values_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_values, cgodot_dictionary_values_allocs = x.DictionaryValues.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_values_allocs)

	var cgodot_dictionary_get_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_get, cgodot_dictionary_get_allocs = x.DictionaryGet.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_get_allocs)

	var cgodot_dictionary_set_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_set, cgodot_dictionary_set_allocs = x.DictionarySet.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_set_allocs)

	var cgodot_dictionary_operator_index_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_operator_index, cgodot_dictionary_operator_index_allocs = x.DictionaryOperatorIndex.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_operator_index_allocs)

	var cgodot_dictionary_operator_index_const_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_operator_index_const, cgodot_dictionary_operator_index_const_allocs = x.DictionaryOperatorIndexConst.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_operator_index_const_allocs)

	var cgodot_dictionary_next_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_next, cgodot_dictionary_next_allocs = x.DictionaryNext.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_next_allocs)

	var cgodot_dictionary_operator_equal_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_operator_equal, cgodot_dictionary_operator_equal_allocs = x.DictionaryOperatorEqual.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_operator_equal_allocs)

	var cgodot_dictionary_to_json_allocs *cgoAllocMap
	ref57717e51.godot_dictionary_to_json, cgodot_dictionary_to_json_allocs = x.DictionaryToJson.PassRef()
	allocs57717e51.Borrow(cgodot_dictionary_to_json_allocs)

	var cgodot_node_path_new_allocs *cgoAllocMap
	ref57717e51.godot_node_path_new, cgodot_node_path_new_allocs = x.NodePathNew.PassRef()
	allocs57717e51.Borrow(cgodot_node_path_new_allocs)

	var cgodot_node_path_new_copy_allocs *cgoAllocMap
	ref57717e51.godot_node_path_new_copy, cgodot_node_path_new_copy_allocs = x.NodePathNewCopy.PassRef()
	allocs57717e51.Borrow(cgodot_node_path_new_copy_allocs)

	var cgodot_node_path_destroy_allocs *cgoAllocMap
	ref57717e51.godot_node_path_destroy, cgodot_node_path_destroy_allocs = x.NodePathDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_node_path_destroy_allocs)

	var cgodot_node_path_as_string_allocs *cgoAllocMap
	ref57717e51.godot_node_path_as_string, cgodot_node_path_as_string_allocs = x.NodePathAsString.PassRef()
	allocs57717e51.Borrow(cgodot_node_path_as_string_allocs)

	var cgodot_node_path_is_absolute_allocs *cgoAllocMap
	ref57717e51.godot_node_path_is_absolute, cgodot_node_path_is_absolute_allocs = x.NodePathIsAbsolute.PassRef()
	allocs57717e51.Borrow(cgodot_node_path_is_absolute_allocs)

	var cgodot_node_path_get_name_count_allocs *cgoAllocMap
	ref57717e51.godot_node_path_get_name_count, cgodot_node_path_get_name_count_allocs = x.NodePathGetNameCount.PassRef()
	allocs57717e51.Borrow(cgodot_node_path_get_name_count_allocs)

	var cgodot_node_path_get_name_allocs *cgoAllocMap
	ref57717e51.godot_node_path_get_name, cgodot_node_path_get_name_allocs = x.NodePathGetName.PassRef()
	allocs57717e51.Borrow(cgodot_node_path_get_name_allocs)

	var cgodot_node_path_get_subname_count_allocs *cgoAllocMap
	ref57717e51.godot_node_path_get_subname_count, cgodot_node_path_get_subname_count_allocs = x.NodePathGetSubnameCount.PassRef()
	allocs57717e51.Borrow(cgodot_node_path_get_subname_count_allocs)

	var cgodot_node_path_get_subname_allocs *cgoAllocMap
	ref57717e51.godot_node_path_get_subname, cgodot_node_path_get_subname_allocs = x.NodePathGetSubname.PassRef()
	allocs57717e51.Borrow(cgodot_node_path_get_subname_allocs)

	var cgodot_node_path_get_concatenated_subnames_allocs *cgoAllocMap
	ref57717e51.godot_node_path_get_concatenated_subnames, cgodot_node_path_get_concatenated_subnames_allocs = x.NodePathGetConcatenatedSubnames.PassRef()
	allocs57717e51.Borrow(cgodot_node_path_get_concatenated_subnames_allocs)

	var cgodot_node_path_is_empty_allocs *cgoAllocMap
	ref57717e51.godot_node_path_is_empty, cgodot_node_path_is_empty_allocs = x.NodePathIsEmpty.PassRef()
	allocs57717e51.Borrow(cgodot_node_path_is_empty_allocs)

	var cgodot_node_path_operator_equal_allocs *cgoAllocMap
	ref57717e51.godot_node_path_operator_equal, cgodot_node_path_operator_equal_allocs = x.NodePathOperatorEqual.PassRef()
	allocs57717e51.Borrow(cgodot_node_path_operator_equal_allocs)

	var cgodot_plane_new_with_reals_allocs *cgoAllocMap
	ref57717e51.godot_plane_new_with_reals, cgodot_plane_new_with_reals_allocs = x.PlaneNewWithReals.PassRef()
	allocs57717e51.Borrow(cgodot_plane_new_with_reals_allocs)

	var cgodot_plane_new_with_vectors_allocs *cgoAllocMap
	ref57717e51.godot_plane_new_with_vectors, cgodot_plane_new_with_vectors_allocs = x.PlaneNewWithVectors.PassRef()
	allocs57717e51.Borrow(cgodot_plane_new_with_vectors_allocs)

	var cgodot_plane_new_with_normal_allocs *cgoAllocMap
	ref57717e51.godot_plane_new_with_normal, cgodot_plane_new_with_normal_allocs = x.PlaneNewWithNormal.PassRef()
	allocs57717e51.Borrow(cgodot_plane_new_with_normal_allocs)

	var cgodot_plane_as_string_allocs *cgoAllocMap
	ref57717e51.godot_plane_as_string, cgodot_plane_as_string_allocs = x.PlaneAsString.PassRef()
	allocs57717e51.Borrow(cgodot_plane_as_string_allocs)

	var cgodot_plane_normalized_allocs *cgoAllocMap
	ref57717e51.godot_plane_normalized, cgodot_plane_normalized_allocs = x.PlaneNormalized.PassRef()
	allocs57717e51.Borrow(cgodot_plane_normalized_allocs)

	var cgodot_plane_center_allocs *cgoAllocMap
	ref57717e51.godot_plane_center, cgodot_plane_center_allocs = x.PlaneCenter.PassRef()
	allocs57717e51.Borrow(cgodot_plane_center_allocs)

	var cgodot_plane_get_any_point_allocs *cgoAllocMap
	ref57717e51.godot_plane_get_any_point, cgodot_plane_get_any_point_allocs = x.PlaneGetAnyPoint.PassRef()
	allocs57717e51.Borrow(cgodot_plane_get_any_point_allocs)

	var cgodot_plane_is_point_over_allocs *cgoAllocMap
	ref57717e51.godot_plane_is_point_over, cgodot_plane_is_point_over_allocs = x.PlaneIsPointOver.PassRef()
	allocs57717e51.Borrow(cgodot_plane_is_point_over_allocs)

	var cgodot_plane_distance_to_allocs *cgoAllocMap
	ref57717e51.godot_plane_distance_to, cgodot_plane_distance_to_allocs = x.PlaneDistanceTo.PassRef()
	allocs57717e51.Borrow(cgodot_plane_distance_to_allocs)

	var cgodot_plane_has_point_allocs *cgoAllocMap
	ref57717e51.godot_plane_has_point, cgodot_plane_has_point_allocs = x.PlaneHasPoint.PassRef()
	allocs57717e51.Borrow(cgodot_plane_has_point_allocs)

	var cgodot_plane_project_allocs *cgoAllocMap
	ref57717e51.godot_plane_project, cgodot_plane_project_allocs = x.PlaneProject.PassRef()
	allocs57717e51.Borrow(cgodot_plane_project_allocs)

	var cgodot_plane_intersect_3_allocs *cgoAllocMap
	ref57717e51.godot_plane_intersect_3, cgodot_plane_intersect_3_allocs = x.PlaneIntersect3.PassRef()
	allocs57717e51.Borrow(cgodot_plane_intersect_3_allocs)

	var cgodot_plane_intersects_ray_allocs *cgoAllocMap
	ref57717e51.godot_plane_intersects_ray, cgodot_plane_intersects_ray_allocs = x.PlaneIntersectsRay.PassRef()
	allocs57717e51.Borrow(cgodot_plane_intersects_ray_allocs)

	var cgodot_plane_intersects_segment_allocs *cgoAllocMap
	ref57717e51.godot_plane_intersects_segment, cgodot_plane_intersects_segment_allocs = x.PlaneIntersectsSegment.PassRef()
	allocs57717e51.Borrow(cgodot_plane_intersects_segment_allocs)

	var cgodot_plane_operator_neg_allocs *cgoAllocMap
	ref57717e51.godot_plane_operator_neg, cgodot_plane_operator_neg_allocs = x.PlaneOperatorNeg.PassRef()
	allocs57717e51.Borrow(cgodot_plane_operator_neg_allocs)

	var cgodot_plane_operator_equal_allocs *cgoAllocMap
	ref57717e51.godot_plane_operator_equal, cgodot_plane_operator_equal_allocs = x.PlaneOperatorEqual.PassRef()
	allocs57717e51.Borrow(cgodot_plane_operator_equal_allocs)

	var cgodot_plane_set_normal_allocs *cgoAllocMap
	ref57717e51.godot_plane_set_normal, cgodot_plane_set_normal_allocs = x.PlaneSetNormal.PassRef()
	allocs57717e51.Borrow(cgodot_plane_set_normal_allocs)

	var cgodot_plane_get_normal_allocs *cgoAllocMap
	ref57717e51.godot_plane_get_normal, cgodot_plane_get_normal_allocs = x.PlaneGetNormal.PassRef()
	allocs57717e51.Borrow(cgodot_plane_get_normal_allocs)

	var cgodot_plane_get_d_allocs *cgoAllocMap
	ref57717e51.godot_plane_get_d, cgodot_plane_get_d_allocs = x.PlaneGetD.PassRef()
	allocs57717e51.Borrow(cgodot_plane_get_d_allocs)

	var cgodot_plane_set_d_allocs *cgoAllocMap
	ref57717e51.godot_plane_set_d, cgodot_plane_set_d_allocs = x.PlaneSetD.PassRef()
	allocs57717e51.Borrow(cgodot_plane_set_d_allocs)

	var cgodot_rect2_new_with_position_and_size_allocs *cgoAllocMap
	ref57717e51.godot_rect2_new_with_position_and_size, cgodot_rect2_new_with_position_and_size_allocs = x.Rect2NewWithPositionAndSize.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_new_with_position_and_size_allocs)

	var cgodot_rect2_new_allocs *cgoAllocMap
	ref57717e51.godot_rect2_new, cgodot_rect2_new_allocs = x.Rect2New.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_new_allocs)

	var cgodot_rect2_as_string_allocs *cgoAllocMap
	ref57717e51.godot_rect2_as_string, cgodot_rect2_as_string_allocs = x.Rect2AsString.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_as_string_allocs)

	var cgodot_rect2_get_area_allocs *cgoAllocMap
	ref57717e51.godot_rect2_get_area, cgodot_rect2_get_area_allocs = x.Rect2GetArea.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_get_area_allocs)

	var cgodot_rect2_intersects_allocs *cgoAllocMap
	ref57717e51.godot_rect2_intersects, cgodot_rect2_intersects_allocs = x.Rect2Intersects.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_intersects_allocs)

	var cgodot_rect2_encloses_allocs *cgoAllocMap
	ref57717e51.godot_rect2_encloses, cgodot_rect2_encloses_allocs = x.Rect2Encloses.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_encloses_allocs)

	var cgodot_rect2_has_no_area_allocs *cgoAllocMap
	ref57717e51.godot_rect2_has_no_area, cgodot_rect2_has_no_area_allocs = x.Rect2HasNoArea.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_has_no_area_allocs)

	var cgodot_rect2_clip_allocs *cgoAllocMap
	ref57717e51.godot_rect2_clip, cgodot_rect2_clip_allocs = x.Rect2Clip.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_clip_allocs)

	var cgodot_rect2_merge_allocs *cgoAllocMap
	ref57717e51.godot_rect2_merge, cgodot_rect2_merge_allocs = x.Rect2Merge.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_merge_allocs)

	var cgodot_rect2_has_point_allocs *cgoAllocMap
	ref57717e51.godot_rect2_has_point, cgodot_rect2_has_point_allocs = x.Rect2HasPoint.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_has_point_allocs)

	var cgodot_rect2_grow_allocs *cgoAllocMap
	ref57717e51.godot_rect2_grow, cgodot_rect2_grow_allocs = x.Rect2Grow.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_grow_allocs)

	var cgodot_rect2_expand_allocs *cgoAllocMap
	ref57717e51.godot_rect2_expand, cgodot_rect2_expand_allocs = x.Rect2Expand.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_expand_allocs)

	var cgodot_rect2_operator_equal_allocs *cgoAllocMap
	ref57717e51.godot_rect2_operator_equal, cgodot_rect2_operator_equal_allocs = x.Rect2OperatorEqual.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_operator_equal_allocs)

	var cgodot_rect2_get_position_allocs *cgoAllocMap
	ref57717e51.godot_rect2_get_position, cgodot_rect2_get_position_allocs = x.Rect2GetPosition.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_get_position_allocs)

	var cgodot_rect2_get_size_allocs *cgoAllocMap
	ref57717e51.godot_rect2_get_size, cgodot_rect2_get_size_allocs = x.Rect2GetSize.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_get_size_allocs)

	var cgodot_rect2_set_position_allocs *cgoAllocMap
	ref57717e51.godot_rect2_set_position, cgodot_rect2_set_position_allocs = x.Rect2SetPosition.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_set_position_allocs)

	var cgodot_rect2_set_size_allocs *cgoAllocMap
	ref57717e51.godot_rect2_set_size, cgodot_rect2_set_size_allocs = x.Rect2SetSize.PassRef()
	allocs57717e51.Borrow(cgodot_rect2_set_size_allocs)

	var cgodot_aabb_new_allocs *cgoAllocMap
	ref57717e51.godot_aabb_new, cgodot_aabb_new_allocs = x.AabbNew.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_new_allocs)

	var cgodot_aabb_get_position_allocs *cgoAllocMap
	ref57717e51.godot_aabb_get_position, cgodot_aabb_get_position_allocs = x.AabbGetPosition.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_get_position_allocs)

	var cgodot_aabb_set_position_allocs *cgoAllocMap
	ref57717e51.godot_aabb_set_position, cgodot_aabb_set_position_allocs = x.AabbSetPosition.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_set_position_allocs)

	var cgodot_aabb_get_size_allocs *cgoAllocMap
	ref57717e51.godot_aabb_get_size, cgodot_aabb_get_size_allocs = x.AabbGetSize.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_get_size_allocs)

	var cgodot_aabb_set_size_allocs *cgoAllocMap
	ref57717e51.godot_aabb_set_size, cgodot_aabb_set_size_allocs = x.AabbSetSize.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_set_size_allocs)

	var cgodot_aabb_as_string_allocs *cgoAllocMap
	ref57717e51.godot_aabb_as_string, cgodot_aabb_as_string_allocs = x.AabbAsString.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_as_string_allocs)

	var cgodot_aabb_get_area_allocs *cgoAllocMap
	ref57717e51.godot_aabb_get_area, cgodot_aabb_get_area_allocs = x.AabbGetArea.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_get_area_allocs)

	var cgodot_aabb_has_no_area_allocs *cgoAllocMap
	ref57717e51.godot_aabb_has_no_area, cgodot_aabb_has_no_area_allocs = x.AabbHasNoArea.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_has_no_area_allocs)

	var cgodot_aabb_has_no_surface_allocs *cgoAllocMap
	ref57717e51.godot_aabb_has_no_surface, cgodot_aabb_has_no_surface_allocs = x.AabbHasNoSurface.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_has_no_surface_allocs)

	var cgodot_aabb_intersects_allocs *cgoAllocMap
	ref57717e51.godot_aabb_intersects, cgodot_aabb_intersects_allocs = x.AabbIntersects.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_intersects_allocs)

	var cgodot_aabb_encloses_allocs *cgoAllocMap
	ref57717e51.godot_aabb_encloses, cgodot_aabb_encloses_allocs = x.AabbEncloses.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_encloses_allocs)

	var cgodot_aabb_merge_allocs *cgoAllocMap
	ref57717e51.godot_aabb_merge, cgodot_aabb_merge_allocs = x.AabbMerge.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_merge_allocs)

	var cgodot_aabb_intersection_allocs *cgoAllocMap
	ref57717e51.godot_aabb_intersection, cgodot_aabb_intersection_allocs = x.AabbIntersection.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_intersection_allocs)

	var cgodot_aabb_intersects_plane_allocs *cgoAllocMap
	ref57717e51.godot_aabb_intersects_plane, cgodot_aabb_intersects_plane_allocs = x.AabbIntersectsPlane.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_intersects_plane_allocs)

	var cgodot_aabb_intersects_segment_allocs *cgoAllocMap
	ref57717e51.godot_aabb_intersects_segment, cgodot_aabb_intersects_segment_allocs = x.AabbIntersectsSegment.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_intersects_segment_allocs)

	var cgodot_aabb_has_point_allocs *cgoAllocMap
	ref57717e51.godot_aabb_has_point, cgodot_aabb_has_point_allocs = x.AabbHasPoint.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_has_point_allocs)

	var cgodot_aabb_get_support_allocs *cgoAllocMap
	ref57717e51.godot_aabb_get_support, cgodot_aabb_get_support_allocs = x.AabbGetSupport.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_get_support_allocs)

	var cgodot_aabb_get_longest_axis_allocs *cgoAllocMap
	ref57717e51.godot_aabb_get_longest_axis, cgodot_aabb_get_longest_axis_allocs = x.AabbGetLongestAxis.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_get_longest_axis_allocs)

	var cgodot_aabb_get_longest_axis_index_allocs *cgoAllocMap
	ref57717e51.godot_aabb_get_longest_axis_index, cgodot_aabb_get_longest_axis_index_allocs = x.AabbGetLongestAxisIndex.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_get_longest_axis_index_allocs)

	var cgodot_aabb_get_longest_axis_size_allocs *cgoAllocMap
	ref57717e51.godot_aabb_get_longest_axis_size, cgodot_aabb_get_longest_axis_size_allocs = x.AabbGetLongestAxisSize.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_get_longest_axis_size_allocs)

	var cgodot_aabb_get_shortest_axis_allocs *cgoAllocMap
	ref57717e51.godot_aabb_get_shortest_axis, cgodot_aabb_get_shortest_axis_allocs = x.AabbGetShortestAxis.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_get_shortest_axis_allocs)

	var cgodot_aabb_get_shortest_axis_index_allocs *cgoAllocMap
	ref57717e51.godot_aabb_get_shortest_axis_index, cgodot_aabb_get_shortest_axis_index_allocs = x.AabbGetShortestAxisIndex.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_get_shortest_axis_index_allocs)

	var cgodot_aabb_get_shortest_axis_size_allocs *cgoAllocMap
	ref57717e51.godot_aabb_get_shortest_axis_size, cgodot_aabb_get_shortest_axis_size_allocs = x.AabbGetShortestAxisSize.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_get_shortest_axis_size_allocs)

	var cgodot_aabb_expand_allocs *cgoAllocMap
	ref57717e51.godot_aabb_expand, cgodot_aabb_expand_allocs = x.AabbExpand.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_expand_allocs)

	var cgodot_aabb_grow_allocs *cgoAllocMap
	ref57717e51.godot_aabb_grow, cgodot_aabb_grow_allocs = x.AabbGrow.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_grow_allocs)

	var cgodot_aabb_get_endpoint_allocs *cgoAllocMap
	ref57717e51.godot_aabb_get_endpoint, cgodot_aabb_get_endpoint_allocs = x.AabbGetEndpoint.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_get_endpoint_allocs)

	var cgodot_aabb_operator_equal_allocs *cgoAllocMap
	ref57717e51.godot_aabb_operator_equal, cgodot_aabb_operator_equal_allocs = x.AabbOperatorEqual.PassRef()
	allocs57717e51.Borrow(cgodot_aabb_operator_equal_allocs)

	var cgodot_rid_new_allocs *cgoAllocMap
	ref57717e51.godot_rid_new, cgodot_rid_new_allocs = x.RidNew.PassRef()
	allocs57717e51.Borrow(cgodot_rid_new_allocs)

	var cgodot_rid_get_id_allocs *cgoAllocMap
	ref57717e51.godot_rid_get_id, cgodot_rid_get_id_allocs = x.RidGetId.PassRef()
	allocs57717e51.Borrow(cgodot_rid_get_id_allocs)

	var cgodot_rid_new_with_resource_allocs *cgoAllocMap
	ref57717e51.godot_rid_new_with_resource, cgodot_rid_new_with_resource_allocs = x.RidNewWithResource.PassRef()
	allocs57717e51.Borrow(cgodot_rid_new_with_resource_allocs)

	var cgodot_rid_operator_equal_allocs *cgoAllocMap
	ref57717e51.godot_rid_operator_equal, cgodot_rid_operator_equal_allocs = x.RidOperatorEqual.PassRef()
	allocs57717e51.Borrow(cgodot_rid_operator_equal_allocs)

	var cgodot_rid_operator_less_allocs *cgoAllocMap
	ref57717e51.godot_rid_operator_less, cgodot_rid_operator_less_allocs = x.RidOperatorLess.PassRef()
	allocs57717e51.Borrow(cgodot_rid_operator_less_allocs)

	var cgodot_transform_new_with_axis_origin_allocs *cgoAllocMap
	ref57717e51.godot_transform_new_with_axis_origin, cgodot_transform_new_with_axis_origin_allocs = x.TransformNewWithAxisOrigin.PassRef()
	allocs57717e51.Borrow(cgodot_transform_new_with_axis_origin_allocs)

	var cgodot_transform_new_allocs *cgoAllocMap
	ref57717e51.godot_transform_new, cgodot_transform_new_allocs = x.TransformNew.PassRef()
	allocs57717e51.Borrow(cgodot_transform_new_allocs)

	var cgodot_transform_get_basis_allocs *cgoAllocMap
	ref57717e51.godot_transform_get_basis, cgodot_transform_get_basis_allocs = x.TransformGetBasis.PassRef()
	allocs57717e51.Borrow(cgodot_transform_get_basis_allocs)

	var cgodot_transform_set_basis_allocs *cgoAllocMap
	ref57717e51.godot_transform_set_basis, cgodot_transform_set_basis_allocs = x.TransformSetBasis.PassRef()
	allocs57717e51.Borrow(cgodot_transform_set_basis_allocs)

	var cgodot_transform_get_origin_allocs *cgoAllocMap
	ref57717e51.godot_transform_get_origin, cgodot_transform_get_origin_allocs = x.TransformGetOrigin.PassRef()
	allocs57717e51.Borrow(cgodot_transform_get_origin_allocs)

	var cgodot_transform_set_origin_allocs *cgoAllocMap
	ref57717e51.godot_transform_set_origin, cgodot_transform_set_origin_allocs = x.TransformSetOrigin.PassRef()
	allocs57717e51.Borrow(cgodot_transform_set_origin_allocs)

	var cgodot_transform_as_string_allocs *cgoAllocMap
	ref57717e51.godot_transform_as_string, cgodot_transform_as_string_allocs = x.TransformAsString.PassRef()
	allocs57717e51.Borrow(cgodot_transform_as_string_allocs)

	var cgodot_transform_inverse_allocs *cgoAllocMap
	ref57717e51.godot_transform_inverse, cgodot_transform_inverse_allocs = x.TransformInverse.PassRef()
	allocs57717e51.Borrow(cgodot_transform_inverse_allocs)

	var cgodot_transform_affine_inverse_allocs *cgoAllocMap
	ref57717e51.godot_transform_affine_inverse, cgodot_transform_affine_inverse_allocs = x.TransformAffineInverse.PassRef()
	allocs57717e51.Borrow(cgodot_transform_affine_inverse_allocs)

	var cgodot_transform_orthonormalized_allocs *cgoAllocMap
	ref57717e51.godot_transform_orthonormalized, cgodot_transform_orthonormalized_allocs = x.TransformOrthonormalized.PassRef()
	allocs57717e51.Borrow(cgodot_transform_orthonormalized_allocs)

	var cgodot_transform_rotated_allocs *cgoAllocMap
	ref57717e51.godot_transform_rotated, cgodot_transform_rotated_allocs = x.TransformRotated.PassRef()
	allocs57717e51.Borrow(cgodot_transform_rotated_allocs)

	var cgodot_transform_scaled_allocs *cgoAllocMap
	ref57717e51.godot_transform_scaled, cgodot_transform_scaled_allocs = x.TransformScaled.PassRef()
	allocs57717e51.Borrow(cgodot_transform_scaled_allocs)

	var cgodot_transform_translated_allocs *cgoAllocMap
	ref57717e51.godot_transform_translated, cgodot_transform_translated_allocs = x.TransformTranslated.PassRef()
	allocs57717e51.Borrow(cgodot_transform_translated_allocs)

	var cgodot_transform_looking_at_allocs *cgoAllocMap
	ref57717e51.godot_transform_looking_at, cgodot_transform_looking_at_allocs = x.TransformLookingAt.PassRef()
	allocs57717e51.Borrow(cgodot_transform_looking_at_allocs)

	var cgodot_transform_xform_plane_allocs *cgoAllocMap
	ref57717e51.godot_transform_xform_plane, cgodot_transform_xform_plane_allocs = x.TransformXformPlane.PassRef()
	allocs57717e51.Borrow(cgodot_transform_xform_plane_allocs)

	var cgodot_transform_xform_inv_plane_allocs *cgoAllocMap
	ref57717e51.godot_transform_xform_inv_plane, cgodot_transform_xform_inv_plane_allocs = x.TransformXformInvPlane.PassRef()
	allocs57717e51.Borrow(cgodot_transform_xform_inv_plane_allocs)

	var cgodot_transform_new_identity_allocs *cgoAllocMap
	ref57717e51.godot_transform_new_identity, cgodot_transform_new_identity_allocs = x.TransformNewIdentity.PassRef()
	allocs57717e51.Borrow(cgodot_transform_new_identity_allocs)

	var cgodot_transform_operator_equal_allocs *cgoAllocMap
	ref57717e51.godot_transform_operator_equal, cgodot_transform_operator_equal_allocs = x.TransformOperatorEqual.PassRef()
	allocs57717e51.Borrow(cgodot_transform_operator_equal_allocs)

	var cgodot_transform_operator_multiply_allocs *cgoAllocMap
	ref57717e51.godot_transform_operator_multiply, cgodot_transform_operator_multiply_allocs = x.TransformOperatorMultiply.PassRef()
	allocs57717e51.Borrow(cgodot_transform_operator_multiply_allocs)

	var cgodot_transform_xform_vector3_allocs *cgoAllocMap
	ref57717e51.godot_transform_xform_vector3, cgodot_transform_xform_vector3_allocs = x.TransformXformVector3.PassRef()
	allocs57717e51.Borrow(cgodot_transform_xform_vector3_allocs)

	var cgodot_transform_xform_inv_vector3_allocs *cgoAllocMap
	ref57717e51.godot_transform_xform_inv_vector3, cgodot_transform_xform_inv_vector3_allocs = x.TransformXformInvVector3.PassRef()
	allocs57717e51.Borrow(cgodot_transform_xform_inv_vector3_allocs)

	var cgodot_transform_xform_aabb_allocs *cgoAllocMap
	ref57717e51.godot_transform_xform_aabb, cgodot_transform_xform_aabb_allocs = x.TransformXformAabb.PassRef()
	allocs57717e51.Borrow(cgodot_transform_xform_aabb_allocs)

	var cgodot_transform_xform_inv_aabb_allocs *cgoAllocMap
	ref57717e51.godot_transform_xform_inv_aabb, cgodot_transform_xform_inv_aabb_allocs = x.TransformXformInvAabb.PassRef()
	allocs57717e51.Borrow(cgodot_transform_xform_inv_aabb_allocs)

	var cgodot_transform2d_new_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_new, cgodot_transform2d_new_allocs = x.Transform2dNew.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_new_allocs)

	var cgodot_transform2d_new_axis_origin_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_new_axis_origin, cgodot_transform2d_new_axis_origin_allocs = x.Transform2dNewAxisOrigin.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_new_axis_origin_allocs)

	var cgodot_transform2d_as_string_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_as_string, cgodot_transform2d_as_string_allocs = x.Transform2dAsString.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_as_string_allocs)

	var cgodot_transform2d_inverse_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_inverse, cgodot_transform2d_inverse_allocs = x.Transform2dInverse.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_inverse_allocs)

	var cgodot_transform2d_affine_inverse_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_affine_inverse, cgodot_transform2d_affine_inverse_allocs = x.Transform2dAffineInverse.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_affine_inverse_allocs)

	var cgodot_transform2d_get_rotation_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_get_rotation, cgodot_transform2d_get_rotation_allocs = x.Transform2dGetRotation.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_get_rotation_allocs)

	var cgodot_transform2d_get_origin_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_get_origin, cgodot_transform2d_get_origin_allocs = x.Transform2dGetOrigin.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_get_origin_allocs)

	var cgodot_transform2d_get_scale_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_get_scale, cgodot_transform2d_get_scale_allocs = x.Transform2dGetScale.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_get_scale_allocs)

	var cgodot_transform2d_orthonormalized_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_orthonormalized, cgodot_transform2d_orthonormalized_allocs = x.Transform2dOrthonormalized.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_orthonormalized_allocs)

	var cgodot_transform2d_rotated_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_rotated, cgodot_transform2d_rotated_allocs = x.Transform2dRotated.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_rotated_allocs)

	var cgodot_transform2d_scaled_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_scaled, cgodot_transform2d_scaled_allocs = x.Transform2dScaled.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_scaled_allocs)

	var cgodot_transform2d_translated_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_translated, cgodot_transform2d_translated_allocs = x.Transform2dTranslated.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_translated_allocs)

	var cgodot_transform2d_xform_vector2_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_xform_vector2, cgodot_transform2d_xform_vector2_allocs = x.Transform2dXformVector2.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_xform_vector2_allocs)

	var cgodot_transform2d_xform_inv_vector2_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_xform_inv_vector2, cgodot_transform2d_xform_inv_vector2_allocs = x.Transform2dXformInvVector2.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_xform_inv_vector2_allocs)

	var cgodot_transform2d_basis_xform_vector2_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_basis_xform_vector2, cgodot_transform2d_basis_xform_vector2_allocs = x.Transform2dBasisXformVector2.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_basis_xform_vector2_allocs)

	var cgodot_transform2d_basis_xform_inv_vector2_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_basis_xform_inv_vector2, cgodot_transform2d_basis_xform_inv_vector2_allocs = x.Transform2dBasisXformInvVector2.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_basis_xform_inv_vector2_allocs)

	var cgodot_transform2d_interpolate_with_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_interpolate_with, cgodot_transform2d_interpolate_with_allocs = x.Transform2dInterpolateWith.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_interpolate_with_allocs)

	var cgodot_transform2d_operator_equal_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_operator_equal, cgodot_transform2d_operator_equal_allocs = x.Transform2dOperatorEqual.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_operator_equal_allocs)

	var cgodot_transform2d_operator_multiply_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_operator_multiply, cgodot_transform2d_operator_multiply_allocs = x.Transform2dOperatorMultiply.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_operator_multiply_allocs)

	var cgodot_transform2d_new_identity_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_new_identity, cgodot_transform2d_new_identity_allocs = x.Transform2dNewIdentity.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_new_identity_allocs)

	var cgodot_transform2d_xform_rect2_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_xform_rect2, cgodot_transform2d_xform_rect2_allocs = x.Transform2dXformRect2.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_xform_rect2_allocs)

	var cgodot_transform2d_xform_inv_rect2_allocs *cgoAllocMap
	ref57717e51.godot_transform2d_xform_inv_rect2, cgodot_transform2d_xform_inv_rect2_allocs = x.Transform2dXformInvRect2.PassRef()
	allocs57717e51.Borrow(cgodot_transform2d_xform_inv_rect2_allocs)

	var cgodot_variant_get_type_allocs *cgoAllocMap
	ref57717e51.godot_variant_get_type, cgodot_variant_get_type_allocs = x.VariantGetType.PassRef()
	allocs57717e51.Borrow(cgodot_variant_get_type_allocs)

	var cgodot_variant_new_copy_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_copy, cgodot_variant_new_copy_allocs = x.VariantNewCopy.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_copy_allocs)

	var cgodot_variant_new_nil_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_nil, cgodot_variant_new_nil_allocs = x.VariantNewNil.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_nil_allocs)

	var cgodot_variant_new_bool_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_bool, cgodot_variant_new_bool_allocs = x.VariantNewBool.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_bool_allocs)

	var cgodot_variant_new_uint_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_uint, cgodot_variant_new_uint_allocs = x.VariantNewUint.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_uint_allocs)

	var cgodot_variant_new_int_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_int, cgodot_variant_new_int_allocs = x.VariantNewInt.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_int_allocs)

	var cgodot_variant_new_real_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_real, cgodot_variant_new_real_allocs = x.VariantNewReal.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_real_allocs)

	var cgodot_variant_new_string_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_string, cgodot_variant_new_string_allocs = x.VariantNewString.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_string_allocs)

	var cgodot_variant_new_vector2_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_vector2, cgodot_variant_new_vector2_allocs = x.VariantNewVector2.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_vector2_allocs)

	var cgodot_variant_new_rect2_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_rect2, cgodot_variant_new_rect2_allocs = x.VariantNewRect2.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_rect2_allocs)

	var cgodot_variant_new_vector3_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_vector3, cgodot_variant_new_vector3_allocs = x.VariantNewVector3.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_vector3_allocs)

	var cgodot_variant_new_transform2d_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_transform2d, cgodot_variant_new_transform2d_allocs = x.VariantNewTransform2d.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_transform2d_allocs)

	var cgodot_variant_new_plane_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_plane, cgodot_variant_new_plane_allocs = x.VariantNewPlane.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_plane_allocs)

	var cgodot_variant_new_quat_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_quat, cgodot_variant_new_quat_allocs = x.VariantNewQuat.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_quat_allocs)

	var cgodot_variant_new_aabb_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_aabb, cgodot_variant_new_aabb_allocs = x.VariantNewAabb.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_aabb_allocs)

	var cgodot_variant_new_basis_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_basis, cgodot_variant_new_basis_allocs = x.VariantNewBasis.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_basis_allocs)

	var cgodot_variant_new_transform_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_transform, cgodot_variant_new_transform_allocs = x.VariantNewTransform.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_transform_allocs)

	var cgodot_variant_new_color_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_color, cgodot_variant_new_color_allocs = x.VariantNewColor.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_color_allocs)

	var cgodot_variant_new_node_path_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_node_path, cgodot_variant_new_node_path_allocs = x.VariantNewNodePath.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_node_path_allocs)

	var cgodot_variant_new_rid_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_rid, cgodot_variant_new_rid_allocs = x.VariantNewRid.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_rid_allocs)

	var cgodot_variant_new_object_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_object, cgodot_variant_new_object_allocs = x.VariantNewObject.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_object_allocs)

	var cgodot_variant_new_dictionary_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_dictionary, cgodot_variant_new_dictionary_allocs = x.VariantNewDictionary.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_dictionary_allocs)

	var cgodot_variant_new_array_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_array, cgodot_variant_new_array_allocs = x.VariantNewArray.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_array_allocs)

	var cgodot_variant_new_pool_byte_array_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_pool_byte_array, cgodot_variant_new_pool_byte_array_allocs = x.VariantNewPoolByteArray.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_pool_byte_array_allocs)

	var cgodot_variant_new_pool_int_array_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_pool_int_array, cgodot_variant_new_pool_int_array_allocs = x.VariantNewPoolIntArray.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_pool_int_array_allocs)

	var cgodot_variant_new_pool_real_array_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_pool_real_array, cgodot_variant_new_pool_real_array_allocs = x.VariantNewPoolRealArray.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_pool_real_array_allocs)

	var cgodot_variant_new_pool_string_array_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_pool_string_array, cgodot_variant_new_pool_string_array_allocs = x.VariantNewPoolStringArray.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_pool_string_array_allocs)

	var cgodot_variant_new_pool_vector2_array_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_pool_vector2_array, cgodot_variant_new_pool_vector2_array_allocs = x.VariantNewPoolVector2Array.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_pool_vector2_array_allocs)

	var cgodot_variant_new_pool_vector3_array_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_pool_vector3_array, cgodot_variant_new_pool_vector3_array_allocs = x.VariantNewPoolVector3Array.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_pool_vector3_array_allocs)

	var cgodot_variant_new_pool_color_array_allocs *cgoAllocMap
	ref57717e51.godot_variant_new_pool_color_array, cgodot_variant_new_pool_color_array_allocs = x.VariantNewPoolColorArray.PassRef()
	allocs57717e51.Borrow(cgodot_variant_new_pool_color_array_allocs)

	var cgodot_variant_as_bool_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_bool, cgodot_variant_as_bool_allocs = x.VariantAsBool.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_bool_allocs)

	var cgodot_variant_as_uint_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_uint, cgodot_variant_as_uint_allocs = x.VariantAsUint.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_uint_allocs)

	var cgodot_variant_as_int_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_int, cgodot_variant_as_int_allocs = x.VariantAsInt.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_int_allocs)

	var cgodot_variant_as_real_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_real, cgodot_variant_as_real_allocs = x.VariantAsReal.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_real_allocs)

	var cgodot_variant_as_string_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_string, cgodot_variant_as_string_allocs = x.VariantAsString.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_string_allocs)

	var cgodot_variant_as_vector2_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_vector2, cgodot_variant_as_vector2_allocs = x.VariantAsVector2.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_vector2_allocs)

	var cgodot_variant_as_rect2_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_rect2, cgodot_variant_as_rect2_allocs = x.VariantAsRect2.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_rect2_allocs)

	var cgodot_variant_as_vector3_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_vector3, cgodot_variant_as_vector3_allocs = x.VariantAsVector3.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_vector3_allocs)

	var cgodot_variant_as_transform2d_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_transform2d, cgodot_variant_as_transform2d_allocs = x.VariantAsTransform2d.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_transform2d_allocs)

	var cgodot_variant_as_plane_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_plane, cgodot_variant_as_plane_allocs = x.VariantAsPlane.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_plane_allocs)

	var cgodot_variant_as_quat_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_quat, cgodot_variant_as_quat_allocs = x.VariantAsQuat.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_quat_allocs)

	var cgodot_variant_as_aabb_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_aabb, cgodot_variant_as_aabb_allocs = x.VariantAsAabb.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_aabb_allocs)

	var cgodot_variant_as_basis_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_basis, cgodot_variant_as_basis_allocs = x.VariantAsBasis.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_basis_allocs)

	var cgodot_variant_as_transform_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_transform, cgodot_variant_as_transform_allocs = x.VariantAsTransform.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_transform_allocs)

	var cgodot_variant_as_color_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_color, cgodot_variant_as_color_allocs = x.VariantAsColor.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_color_allocs)

	var cgodot_variant_as_node_path_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_node_path, cgodot_variant_as_node_path_allocs = x.VariantAsNodePath.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_node_path_allocs)

	var cgodot_variant_as_rid_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_rid, cgodot_variant_as_rid_allocs = x.VariantAsRid.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_rid_allocs)

	var cgodot_variant_as_object_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_object, cgodot_variant_as_object_allocs = x.VariantAsObject.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_object_allocs)

	var cgodot_variant_as_dictionary_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_dictionary, cgodot_variant_as_dictionary_allocs = x.VariantAsDictionary.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_dictionary_allocs)

	var cgodot_variant_as_array_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_array, cgodot_variant_as_array_allocs = x.VariantAsArray.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_array_allocs)

	var cgodot_variant_as_pool_byte_array_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_pool_byte_array, cgodot_variant_as_pool_byte_array_allocs = x.VariantAsPoolByteArray.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_pool_byte_array_allocs)

	var cgodot_variant_as_pool_int_array_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_pool_int_array, cgodot_variant_as_pool_int_array_allocs = x.VariantAsPoolIntArray.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_pool_int_array_allocs)

	var cgodot_variant_as_pool_real_array_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_pool_real_array, cgodot_variant_as_pool_real_array_allocs = x.VariantAsPoolRealArray.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_pool_real_array_allocs)

	var cgodot_variant_as_pool_string_array_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_pool_string_array, cgodot_variant_as_pool_string_array_allocs = x.VariantAsPoolStringArray.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_pool_string_array_allocs)

	var cgodot_variant_as_pool_vector2_array_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_pool_vector2_array, cgodot_variant_as_pool_vector2_array_allocs = x.VariantAsPoolVector2Array.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_pool_vector2_array_allocs)

	var cgodot_variant_as_pool_vector3_array_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_pool_vector3_array, cgodot_variant_as_pool_vector3_array_allocs = x.VariantAsPoolVector3Array.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_pool_vector3_array_allocs)

	var cgodot_variant_as_pool_color_array_allocs *cgoAllocMap
	ref57717e51.godot_variant_as_pool_color_array, cgodot_variant_as_pool_color_array_allocs = x.VariantAsPoolColorArray.PassRef()
	allocs57717e51.Borrow(cgodot_variant_as_pool_color_array_allocs)

	var cgodot_variant_call_allocs *cgoAllocMap
	ref57717e51.godot_variant_call, cgodot_variant_call_allocs = x.VariantCall.PassRef()
	allocs57717e51.Borrow(cgodot_variant_call_allocs)

	var cgodot_variant_has_method_allocs *cgoAllocMap
	ref57717e51.godot_variant_has_method, cgodot_variant_has_method_allocs = x.VariantHasMethod.PassRef()
	allocs57717e51.Borrow(cgodot_variant_has_method_allocs)

	var cgodot_variant_operator_equal_allocs *cgoAllocMap
	ref57717e51.godot_variant_operator_equal, cgodot_variant_operator_equal_allocs = x.VariantOperatorEqual.PassRef()
	allocs57717e51.Borrow(cgodot_variant_operator_equal_allocs)

	var cgodot_variant_operator_less_allocs *cgoAllocMap
	ref57717e51.godot_variant_operator_less, cgodot_variant_operator_less_allocs = x.VariantOperatorLess.PassRef()
	allocs57717e51.Borrow(cgodot_variant_operator_less_allocs)

	var cgodot_variant_hash_compare_allocs *cgoAllocMap
	ref57717e51.godot_variant_hash_compare, cgodot_variant_hash_compare_allocs = x.VariantHashCompare.PassRef()
	allocs57717e51.Borrow(cgodot_variant_hash_compare_allocs)

	var cgodot_variant_booleanize_allocs *cgoAllocMap
	ref57717e51.godot_variant_booleanize, cgodot_variant_booleanize_allocs = x.VariantBooleanize.PassRef()
	allocs57717e51.Borrow(cgodot_variant_booleanize_allocs)

	var cgodot_variant_destroy_allocs *cgoAllocMap
	ref57717e51.godot_variant_destroy, cgodot_variant_destroy_allocs = x.VariantDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_variant_destroy_allocs)

	var cgodot_string_new_allocs *cgoAllocMap
	ref57717e51.godot_string_new, cgodot_string_new_allocs = x.StringNew.PassRef()
	allocs57717e51.Borrow(cgodot_string_new_allocs)

	var cgodot_string_new_copy_allocs *cgoAllocMap
	ref57717e51.godot_string_new_copy, cgodot_string_new_copy_allocs = x.StringNewCopy.PassRef()
	allocs57717e51.Borrow(cgodot_string_new_copy_allocs)

	var cgodot_string_new_data_allocs *cgoAllocMap
	ref57717e51.godot_string_new_data, cgodot_string_new_data_allocs = x.StringNewData.PassRef()
	allocs57717e51.Borrow(cgodot_string_new_data_allocs)

	var cgodot_string_new_unicode_data_allocs *cgoAllocMap
	ref57717e51.godot_string_new_unicode_data, cgodot_string_new_unicode_data_allocs = x.StringNewUnicodeData.PassRef()
	allocs57717e51.Borrow(cgodot_string_new_unicode_data_allocs)

	var cgodot_string_get_data_allocs *cgoAllocMap
	ref57717e51.godot_string_get_data, cgodot_string_get_data_allocs = x.StringGetData.PassRef()
	allocs57717e51.Borrow(cgodot_string_get_data_allocs)

	var cgodot_string_operator_index_allocs *cgoAllocMap
	ref57717e51.godot_string_operator_index, cgodot_string_operator_index_allocs = x.StringOperatorIndex.PassRef()
	allocs57717e51.Borrow(cgodot_string_operator_index_allocs)

	var cgodot_string_operator_index_const_allocs *cgoAllocMap
	ref57717e51.godot_string_operator_index_const, cgodot_string_operator_index_const_allocs = x.StringOperatorIndexConst.PassRef()
	allocs57717e51.Borrow(cgodot_string_operator_index_const_allocs)

	var cgodot_string_unicode_str_allocs *cgoAllocMap
	ref57717e51.godot_string_unicode_str, cgodot_string_unicode_str_allocs = x.StringUnicodeStr.PassRef()
	allocs57717e51.Borrow(cgodot_string_unicode_str_allocs)

	var cgodot_string_operator_equal_allocs *cgoAllocMap
	ref57717e51.godot_string_operator_equal, cgodot_string_operator_equal_allocs = x.StringOperatorEqual.PassRef()
	allocs57717e51.Borrow(cgodot_string_operator_equal_allocs)

	var cgodot_string_operator_less_allocs *cgoAllocMap
	ref57717e51.godot_string_operator_less, cgodot_string_operator_less_allocs = x.StringOperatorLess.PassRef()
	allocs57717e51.Borrow(cgodot_string_operator_less_allocs)

	var cgodot_string_operator_plus_allocs *cgoAllocMap
	ref57717e51.godot_string_operator_plus, cgodot_string_operator_plus_allocs = x.StringOperatorPlus.PassRef()
	allocs57717e51.Borrow(cgodot_string_operator_plus_allocs)

	var cgodot_string_length_allocs *cgoAllocMap
	ref57717e51.godot_string_length, cgodot_string_length_allocs = x.StringLength.PassRef()
	allocs57717e51.Borrow(cgodot_string_length_allocs)

	var cgodot_string_begins_with_allocs *cgoAllocMap
	ref57717e51.godot_string_begins_with, cgodot_string_begins_with_allocs = x.StringBeginsWith.PassRef()
	allocs57717e51.Borrow(cgodot_string_begins_with_allocs)

	var cgodot_string_begins_with_char_array_allocs *cgoAllocMap
	ref57717e51.godot_string_begins_with_char_array, cgodot_string_begins_with_char_array_allocs = x.StringBeginsWithCharArray.PassRef()
	allocs57717e51.Borrow(cgodot_string_begins_with_char_array_allocs)

	var cgodot_string_bigrams_allocs *cgoAllocMap
	ref57717e51.godot_string_bigrams, cgodot_string_bigrams_allocs = x.StringBigrams.PassRef()
	allocs57717e51.Borrow(cgodot_string_bigrams_allocs)

	var cgodot_string_chr_allocs *cgoAllocMap
	ref57717e51.godot_string_chr, cgodot_string_chr_allocs = x.StringChr.PassRef()
	allocs57717e51.Borrow(cgodot_string_chr_allocs)

	var cgodot_string_ends_with_allocs *cgoAllocMap
	ref57717e51.godot_string_ends_with, cgodot_string_ends_with_allocs = x.StringEndsWith.PassRef()
	allocs57717e51.Borrow(cgodot_string_ends_with_allocs)

	var cgodot_string_find_allocs *cgoAllocMap
	ref57717e51.godot_string_find, cgodot_string_find_allocs = x.StringFind.PassRef()
	allocs57717e51.Borrow(cgodot_string_find_allocs)

	var cgodot_string_find_from_allocs *cgoAllocMap
	ref57717e51.godot_string_find_from, cgodot_string_find_from_allocs = x.StringFindFrom.PassRef()
	allocs57717e51.Borrow(cgodot_string_find_from_allocs)

	var cgodot_string_findmk_allocs *cgoAllocMap
	ref57717e51.godot_string_findmk, cgodot_string_findmk_allocs = x.StringFindmk.PassRef()
	allocs57717e51.Borrow(cgodot_string_findmk_allocs)

	var cgodot_string_findmk_from_allocs *cgoAllocMap
	ref57717e51.godot_string_findmk_from, cgodot_string_findmk_from_allocs = x.StringFindmkFrom.PassRef()
	allocs57717e51.Borrow(cgodot_string_findmk_from_allocs)

	var cgodot_string_findmk_from_in_place_allocs *cgoAllocMap
	ref57717e51.godot_string_findmk_from_in_place, cgodot_string_findmk_from_in_place_allocs = x.StringFindmkFromInPlace.PassRef()
	allocs57717e51.Borrow(cgodot_string_findmk_from_in_place_allocs)

	var cgodot_string_findn_allocs *cgoAllocMap
	ref57717e51.godot_string_findn, cgodot_string_findn_allocs = x.StringFindn.PassRef()
	allocs57717e51.Borrow(cgodot_string_findn_allocs)

	var cgodot_string_findn_from_allocs *cgoAllocMap
	ref57717e51.godot_string_findn_from, cgodot_string_findn_from_allocs = x.StringFindnFrom.PassRef()
	allocs57717e51.Borrow(cgodot_string_findn_from_allocs)

	var cgodot_string_find_last_allocs *cgoAllocMap
	ref57717e51.godot_string_find_last, cgodot_string_find_last_allocs = x.StringFindLast.PassRef()
	allocs57717e51.Borrow(cgodot_string_find_last_allocs)

	var cgodot_string_format_allocs *cgoAllocMap
	ref57717e51.godot_string_format, cgodot_string_format_allocs = x.StringFormat.PassRef()
	allocs57717e51.Borrow(cgodot_string_format_allocs)

	var cgodot_string_format_with_custom_placeholder_allocs *cgoAllocMap
	ref57717e51.godot_string_format_with_custom_placeholder, cgodot_string_format_with_custom_placeholder_allocs = x.StringFormatWithCustomPlaceholder.PassRef()
	allocs57717e51.Borrow(cgodot_string_format_with_custom_placeholder_allocs)

	var cgodot_string_hex_encode_buffer_allocs *cgoAllocMap
	ref57717e51.godot_string_hex_encode_buffer, cgodot_string_hex_encode_buffer_allocs = x.StringHexEncodeBuffer.PassRef()
	allocs57717e51.Borrow(cgodot_string_hex_encode_buffer_allocs)

	var cgodot_string_hex_to_int_allocs *cgoAllocMap
	ref57717e51.godot_string_hex_to_int, cgodot_string_hex_to_int_allocs = x.StringHexToInt.PassRef()
	allocs57717e51.Borrow(cgodot_string_hex_to_int_allocs)

	var cgodot_string_hex_to_int_without_prefix_allocs *cgoAllocMap
	ref57717e51.godot_string_hex_to_int_without_prefix, cgodot_string_hex_to_int_without_prefix_allocs = x.StringHexToIntWithoutPrefix.PassRef()
	allocs57717e51.Borrow(cgodot_string_hex_to_int_without_prefix_allocs)

	var cgodot_string_insert_allocs *cgoAllocMap
	ref57717e51.godot_string_insert, cgodot_string_insert_allocs = x.StringInsert.PassRef()
	allocs57717e51.Borrow(cgodot_string_insert_allocs)

	var cgodot_string_is_numeric_allocs *cgoAllocMap
	ref57717e51.godot_string_is_numeric, cgodot_string_is_numeric_allocs = x.StringIsNumeric.PassRef()
	allocs57717e51.Borrow(cgodot_string_is_numeric_allocs)

	var cgodot_string_is_subsequence_of_allocs *cgoAllocMap
	ref57717e51.godot_string_is_subsequence_of, cgodot_string_is_subsequence_of_allocs = x.StringIsSubsequenceOf.PassRef()
	allocs57717e51.Borrow(cgodot_string_is_subsequence_of_allocs)

	var cgodot_string_is_subsequence_ofi_allocs *cgoAllocMap
	ref57717e51.godot_string_is_subsequence_ofi, cgodot_string_is_subsequence_ofi_allocs = x.StringIsSubsequenceOfi.PassRef()
	allocs57717e51.Borrow(cgodot_string_is_subsequence_ofi_allocs)

	var cgodot_string_lpad_allocs *cgoAllocMap
	ref57717e51.godot_string_lpad, cgodot_string_lpad_allocs = x.StringLpad.PassRef()
	allocs57717e51.Borrow(cgodot_string_lpad_allocs)

	var cgodot_string_lpad_with_custom_character_allocs *cgoAllocMap
	ref57717e51.godot_string_lpad_with_custom_character, cgodot_string_lpad_with_custom_character_allocs = x.StringLpadWithCustomCharacter.PassRef()
	allocs57717e51.Borrow(cgodot_string_lpad_with_custom_character_allocs)

	var cgodot_string_match_allocs *cgoAllocMap
	ref57717e51.godot_string_match, cgodot_string_match_allocs = x.StringMatch.PassRef()
	allocs57717e51.Borrow(cgodot_string_match_allocs)

	var cgodot_string_matchn_allocs *cgoAllocMap
	ref57717e51.godot_string_matchn, cgodot_string_matchn_allocs = x.StringMatchn.PassRef()
	allocs57717e51.Borrow(cgodot_string_matchn_allocs)

	var cgodot_string_md5_allocs *cgoAllocMap
	ref57717e51.godot_string_md5, cgodot_string_md5_allocs = x.StringMd5.PassRef()
	allocs57717e51.Borrow(cgodot_string_md5_allocs)

	var cgodot_string_num_allocs *cgoAllocMap
	ref57717e51.godot_string_num, cgodot_string_num_allocs = x.StringNum.PassRef()
	allocs57717e51.Borrow(cgodot_string_num_allocs)

	var cgodot_string_num_int64_allocs *cgoAllocMap
	ref57717e51.godot_string_num_int64, cgodot_string_num_int64_allocs = x.StringNumInt64.PassRef()
	allocs57717e51.Borrow(cgodot_string_num_int64_allocs)

	var cgodot_string_num_int64_capitalized_allocs *cgoAllocMap
	ref57717e51.godot_string_num_int64_capitalized, cgodot_string_num_int64_capitalized_allocs = x.StringNumInt64Capitalized.PassRef()
	allocs57717e51.Borrow(cgodot_string_num_int64_capitalized_allocs)

	var cgodot_string_num_real_allocs *cgoAllocMap
	ref57717e51.godot_string_num_real, cgodot_string_num_real_allocs = x.StringNumReal.PassRef()
	allocs57717e51.Borrow(cgodot_string_num_real_allocs)

	var cgodot_string_num_scientific_allocs *cgoAllocMap
	ref57717e51.godot_string_num_scientific, cgodot_string_num_scientific_allocs = x.StringNumScientific.PassRef()
	allocs57717e51.Borrow(cgodot_string_num_scientific_allocs)

	var cgodot_string_num_with_decimals_allocs *cgoAllocMap
	ref57717e51.godot_string_num_with_decimals, cgodot_string_num_with_decimals_allocs = x.StringNumWithDecimals.PassRef()
	allocs57717e51.Borrow(cgodot_string_num_with_decimals_allocs)

	var cgodot_string_pad_decimals_allocs *cgoAllocMap
	ref57717e51.godot_string_pad_decimals, cgodot_string_pad_decimals_allocs = x.StringPadDecimals.PassRef()
	allocs57717e51.Borrow(cgodot_string_pad_decimals_allocs)

	var cgodot_string_pad_zeros_allocs *cgoAllocMap
	ref57717e51.godot_string_pad_zeros, cgodot_string_pad_zeros_allocs = x.StringPadZeros.PassRef()
	allocs57717e51.Borrow(cgodot_string_pad_zeros_allocs)

	var cgodot_string_replace_first_allocs *cgoAllocMap
	ref57717e51.godot_string_replace_first, cgodot_string_replace_first_allocs = x.StringReplaceFirst.PassRef()
	allocs57717e51.Borrow(cgodot_string_replace_first_allocs)

	var cgodot_string_replace_allocs *cgoAllocMap
	ref57717e51.godot_string_replace, cgodot_string_replace_allocs = x.StringReplace.PassRef()
	allocs57717e51.Borrow(cgodot_string_replace_allocs)

	var cgodot_string_replacen_allocs *cgoAllocMap
	ref57717e51.godot_string_replacen, cgodot_string_replacen_allocs = x.StringReplacen.PassRef()
	allocs57717e51.Borrow(cgodot_string_replacen_allocs)

	var cgodot_string_rfind_allocs *cgoAllocMap
	ref57717e51.godot_string_rfind, cgodot_string_rfind_allocs = x.StringRfind.PassRef()
	allocs57717e51.Borrow(cgodot_string_rfind_allocs)

	var cgodot_string_rfindn_allocs *cgoAllocMap
	ref57717e51.godot_string_rfindn, cgodot_string_rfindn_allocs = x.StringRfindn.PassRef()
	allocs57717e51.Borrow(cgodot_string_rfindn_allocs)

	var cgodot_string_rfind_from_allocs *cgoAllocMap
	ref57717e51.godot_string_rfind_from, cgodot_string_rfind_from_allocs = x.StringRfindFrom.PassRef()
	allocs57717e51.Borrow(cgodot_string_rfind_from_allocs)

	var cgodot_string_rfindn_from_allocs *cgoAllocMap
	ref57717e51.godot_string_rfindn_from, cgodot_string_rfindn_from_allocs = x.StringRfindnFrom.PassRef()
	allocs57717e51.Borrow(cgodot_string_rfindn_from_allocs)

	var cgodot_string_rpad_allocs *cgoAllocMap
	ref57717e51.godot_string_rpad, cgodot_string_rpad_allocs = x.StringRpad.PassRef()
	allocs57717e51.Borrow(cgodot_string_rpad_allocs)

	var cgodot_string_rpad_with_custom_character_allocs *cgoAllocMap
	ref57717e51.godot_string_rpad_with_custom_character, cgodot_string_rpad_with_custom_character_allocs = x.StringRpadWithCustomCharacter.PassRef()
	allocs57717e51.Borrow(cgodot_string_rpad_with_custom_character_allocs)

	var cgodot_string_similarity_allocs *cgoAllocMap
	ref57717e51.godot_string_similarity, cgodot_string_similarity_allocs = x.StringSimilarity.PassRef()
	allocs57717e51.Borrow(cgodot_string_similarity_allocs)

	var cgodot_string_sprintf_allocs *cgoAllocMap
	ref57717e51.godot_string_sprintf, cgodot_string_sprintf_allocs = x.StringSprintf.PassRef()
	allocs57717e51.Borrow(cgodot_string_sprintf_allocs)

	var cgodot_string_substr_allocs *cgoAllocMap
	ref57717e51.godot_string_substr, cgodot_string_substr_allocs = x.StringSubstr.PassRef()
	allocs57717e51.Borrow(cgodot_string_substr_allocs)

	var cgodot_string_to_double_allocs *cgoAllocMap
	ref57717e51.godot_string_to_double, cgodot_string_to_double_allocs = x.StringToDouble.PassRef()
	allocs57717e51.Borrow(cgodot_string_to_double_allocs)

	var cgodot_string_to_float_allocs *cgoAllocMap
	ref57717e51.godot_string_to_float, cgodot_string_to_float_allocs = x.StringToFloat.PassRef()
	allocs57717e51.Borrow(cgodot_string_to_float_allocs)

	var cgodot_string_to_int_allocs *cgoAllocMap
	ref57717e51.godot_string_to_int, cgodot_string_to_int_allocs = x.StringToInt.PassRef()
	allocs57717e51.Borrow(cgodot_string_to_int_allocs)

	var cgodot_string_camelcase_to_underscore_allocs *cgoAllocMap
	ref57717e51.godot_string_camelcase_to_underscore, cgodot_string_camelcase_to_underscore_allocs = x.StringCamelcaseToUnderscore.PassRef()
	allocs57717e51.Borrow(cgodot_string_camelcase_to_underscore_allocs)

	var cgodot_string_camelcase_to_underscore_lowercased_allocs *cgoAllocMap
	ref57717e51.godot_string_camelcase_to_underscore_lowercased, cgodot_string_camelcase_to_underscore_lowercased_allocs = x.StringCamelcaseToUnderscoreLowercased.PassRef()
	allocs57717e51.Borrow(cgodot_string_camelcase_to_underscore_lowercased_allocs)

	var cgodot_string_capitalize_allocs *cgoAllocMap
	ref57717e51.godot_string_capitalize, cgodot_string_capitalize_allocs = x.StringCapitalize.PassRef()
	allocs57717e51.Borrow(cgodot_string_capitalize_allocs)

	var cgodot_string_char_to_double_allocs *cgoAllocMap
	ref57717e51.godot_string_char_to_double, cgodot_string_char_to_double_allocs = x.StringCharToDouble.PassRef()
	allocs57717e51.Borrow(cgodot_string_char_to_double_allocs)

	var cgodot_string_char_to_int_allocs *cgoAllocMap
	ref57717e51.godot_string_char_to_int, cgodot_string_char_to_int_allocs = x.StringCharToInt.PassRef()
	allocs57717e51.Borrow(cgodot_string_char_to_int_allocs)

	var cgodot_string_wchar_to_int_allocs *cgoAllocMap
	ref57717e51.godot_string_wchar_to_int, cgodot_string_wchar_to_int_allocs = x.StringWcharToInt.PassRef()
	allocs57717e51.Borrow(cgodot_string_wchar_to_int_allocs)

	var cgodot_string_char_to_int_with_len_allocs *cgoAllocMap
	ref57717e51.godot_string_char_to_int_with_len, cgodot_string_char_to_int_with_len_allocs = x.StringCharToIntWithLen.PassRef()
	allocs57717e51.Borrow(cgodot_string_char_to_int_with_len_allocs)

	var cgodot_string_char_to_int64_with_len_allocs *cgoAllocMap
	ref57717e51.godot_string_char_to_int64_with_len, cgodot_string_char_to_int64_with_len_allocs = x.StringCharToInt64WithLen.PassRef()
	allocs57717e51.Borrow(cgodot_string_char_to_int64_with_len_allocs)

	var cgodot_string_hex_to_int64_allocs *cgoAllocMap
	ref57717e51.godot_string_hex_to_int64, cgodot_string_hex_to_int64_allocs = x.StringHexToInt64.PassRef()
	allocs57717e51.Borrow(cgodot_string_hex_to_int64_allocs)

	var cgodot_string_hex_to_int64_with_prefix_allocs *cgoAllocMap
	ref57717e51.godot_string_hex_to_int64_with_prefix, cgodot_string_hex_to_int64_with_prefix_allocs = x.StringHexToInt64WithPrefix.PassRef()
	allocs57717e51.Borrow(cgodot_string_hex_to_int64_with_prefix_allocs)

	var cgodot_string_to_int64_allocs *cgoAllocMap
	ref57717e51.godot_string_to_int64, cgodot_string_to_int64_allocs = x.StringToInt64.PassRef()
	allocs57717e51.Borrow(cgodot_string_to_int64_allocs)

	var cgodot_string_unicode_char_to_double_allocs *cgoAllocMap
	ref57717e51.godot_string_unicode_char_to_double, cgodot_string_unicode_char_to_double_allocs = x.StringUnicodeCharToDouble.PassRef()
	allocs57717e51.Borrow(cgodot_string_unicode_char_to_double_allocs)

	var cgodot_string_get_slice_count_allocs *cgoAllocMap
	ref57717e51.godot_string_get_slice_count, cgodot_string_get_slice_count_allocs = x.StringGetSliceCount.PassRef()
	allocs57717e51.Borrow(cgodot_string_get_slice_count_allocs)

	var cgodot_string_get_slice_allocs *cgoAllocMap
	ref57717e51.godot_string_get_slice, cgodot_string_get_slice_allocs = x.StringGetSlice.PassRef()
	allocs57717e51.Borrow(cgodot_string_get_slice_allocs)

	var cgodot_string_get_slicec_allocs *cgoAllocMap
	ref57717e51.godot_string_get_slicec, cgodot_string_get_slicec_allocs = x.StringGetSlicec.PassRef()
	allocs57717e51.Borrow(cgodot_string_get_slicec_allocs)

	var cgodot_string_split_allocs *cgoAllocMap
	ref57717e51.godot_string_split, cgodot_string_split_allocs = x.StringSplit.PassRef()
	allocs57717e51.Borrow(cgodot_string_split_allocs)

	var cgodot_string_split_allow_empty_allocs *cgoAllocMap
	ref57717e51.godot_string_split_allow_empty, cgodot_string_split_allow_empty_allocs = x.StringSplitAllowEmpty.PassRef()
	allocs57717e51.Borrow(cgodot_string_split_allow_empty_allocs)

	var cgodot_string_split_floats_allocs *cgoAllocMap
	ref57717e51.godot_string_split_floats, cgodot_string_split_floats_allocs = x.StringSplitFloats.PassRef()
	allocs57717e51.Borrow(cgodot_string_split_floats_allocs)

	var cgodot_string_split_floats_allows_empty_allocs *cgoAllocMap
	ref57717e51.godot_string_split_floats_allows_empty, cgodot_string_split_floats_allows_empty_allocs = x.StringSplitFloatsAllowsEmpty.PassRef()
	allocs57717e51.Borrow(cgodot_string_split_floats_allows_empty_allocs)

	var cgodot_string_split_floats_mk_allocs *cgoAllocMap
	ref57717e51.godot_string_split_floats_mk, cgodot_string_split_floats_mk_allocs = x.StringSplitFloatsMk.PassRef()
	allocs57717e51.Borrow(cgodot_string_split_floats_mk_allocs)

	var cgodot_string_split_floats_mk_allows_empty_allocs *cgoAllocMap
	ref57717e51.godot_string_split_floats_mk_allows_empty, cgodot_string_split_floats_mk_allows_empty_allocs = x.StringSplitFloatsMkAllowsEmpty.PassRef()
	allocs57717e51.Borrow(cgodot_string_split_floats_mk_allows_empty_allocs)

	var cgodot_string_split_ints_allocs *cgoAllocMap
	ref57717e51.godot_string_split_ints, cgodot_string_split_ints_allocs = x.StringSplitInts.PassRef()
	allocs57717e51.Borrow(cgodot_string_split_ints_allocs)

	var cgodot_string_split_ints_allows_empty_allocs *cgoAllocMap
	ref57717e51.godot_string_split_ints_allows_empty, cgodot_string_split_ints_allows_empty_allocs = x.StringSplitIntsAllowsEmpty.PassRef()
	allocs57717e51.Borrow(cgodot_string_split_ints_allows_empty_allocs)

	var cgodot_string_split_ints_mk_allocs *cgoAllocMap
	ref57717e51.godot_string_split_ints_mk, cgodot_string_split_ints_mk_allocs = x.StringSplitIntsMk.PassRef()
	allocs57717e51.Borrow(cgodot_string_split_ints_mk_allocs)

	var cgodot_string_split_ints_mk_allows_empty_allocs *cgoAllocMap
	ref57717e51.godot_string_split_ints_mk_allows_empty, cgodot_string_split_ints_mk_allows_empty_allocs = x.StringSplitIntsMkAllowsEmpty.PassRef()
	allocs57717e51.Borrow(cgodot_string_split_ints_mk_allows_empty_allocs)

	var cgodot_string_split_spaces_allocs *cgoAllocMap
	ref57717e51.godot_string_split_spaces, cgodot_string_split_spaces_allocs = x.StringSplitSpaces.PassRef()
	allocs57717e51.Borrow(cgodot_string_split_spaces_allocs)

	var cgodot_string_char_lowercase_allocs *cgoAllocMap
	ref57717e51.godot_string_char_lowercase, cgodot_string_char_lowercase_allocs = x.StringCharLowercase.PassRef()
	allocs57717e51.Borrow(cgodot_string_char_lowercase_allocs)

	var cgodot_string_char_uppercase_allocs *cgoAllocMap
	ref57717e51.godot_string_char_uppercase, cgodot_string_char_uppercase_allocs = x.StringCharUppercase.PassRef()
	allocs57717e51.Borrow(cgodot_string_char_uppercase_allocs)

	var cgodot_string_to_lower_allocs *cgoAllocMap
	ref57717e51.godot_string_to_lower, cgodot_string_to_lower_allocs = x.StringToLower.PassRef()
	allocs57717e51.Borrow(cgodot_string_to_lower_allocs)

	var cgodot_string_to_upper_allocs *cgoAllocMap
	ref57717e51.godot_string_to_upper, cgodot_string_to_upper_allocs = x.StringToUpper.PassRef()
	allocs57717e51.Borrow(cgodot_string_to_upper_allocs)

	var cgodot_string_get_basename_allocs *cgoAllocMap
	ref57717e51.godot_string_get_basename, cgodot_string_get_basename_allocs = x.StringGetBasename.PassRef()
	allocs57717e51.Borrow(cgodot_string_get_basename_allocs)

	var cgodot_string_get_extension_allocs *cgoAllocMap
	ref57717e51.godot_string_get_extension, cgodot_string_get_extension_allocs = x.StringGetExtension.PassRef()
	allocs57717e51.Borrow(cgodot_string_get_extension_allocs)

	var cgodot_string_left_allocs *cgoAllocMap
	ref57717e51.godot_string_left, cgodot_string_left_allocs = x.StringLeft.PassRef()
	allocs57717e51.Borrow(cgodot_string_left_allocs)

	var cgodot_string_ord_at_allocs *cgoAllocMap
	ref57717e51.godot_string_ord_at, cgodot_string_ord_at_allocs = x.StringOrdAt.PassRef()
	allocs57717e51.Borrow(cgodot_string_ord_at_allocs)

	var cgodot_string_plus_file_allocs *cgoAllocMap
	ref57717e51.godot_string_plus_file, cgodot_string_plus_file_allocs = x.StringPlusFile.PassRef()
	allocs57717e51.Borrow(cgodot_string_plus_file_allocs)

	var cgodot_string_right_allocs *cgoAllocMap
	ref57717e51.godot_string_right, cgodot_string_right_allocs = x.StringRight.PassRef()
	allocs57717e51.Borrow(cgodot_string_right_allocs)

	var cgodot_string_strip_edges_allocs *cgoAllocMap
	ref57717e51.godot_string_strip_edges, cgodot_string_strip_edges_allocs = x.StringStripEdges.PassRef()
	allocs57717e51.Borrow(cgodot_string_strip_edges_allocs)

	var cgodot_string_strip_escapes_allocs *cgoAllocMap
	ref57717e51.godot_string_strip_escapes, cgodot_string_strip_escapes_allocs = x.StringStripEscapes.PassRef()
	allocs57717e51.Borrow(cgodot_string_strip_escapes_allocs)

	var cgodot_string_erase_allocs *cgoAllocMap
	ref57717e51.godot_string_erase, cgodot_string_erase_allocs = x.StringErase.PassRef()
	allocs57717e51.Borrow(cgodot_string_erase_allocs)

	var cgodot_string_ascii_allocs *cgoAllocMap
	ref57717e51.godot_string_ascii, cgodot_string_ascii_allocs = x.StringAscii.PassRef()
	allocs57717e51.Borrow(cgodot_string_ascii_allocs)

	var cgodot_string_ascii_extended_allocs *cgoAllocMap
	ref57717e51.godot_string_ascii_extended, cgodot_string_ascii_extended_allocs = x.StringAsciiExtended.PassRef()
	allocs57717e51.Borrow(cgodot_string_ascii_extended_allocs)

	var cgodot_string_utf8_allocs *cgoAllocMap
	ref57717e51.godot_string_utf8, cgodot_string_utf8_allocs = x.StringUtf8.PassRef()
	allocs57717e51.Borrow(cgodot_string_utf8_allocs)

	var cgodot_string_parse_utf8_allocs *cgoAllocMap
	ref57717e51.godot_string_parse_utf8, cgodot_string_parse_utf8_allocs = x.StringParseUtf8.PassRef()
	allocs57717e51.Borrow(cgodot_string_parse_utf8_allocs)

	var cgodot_string_parse_utf8_with_len_allocs *cgoAllocMap
	ref57717e51.godot_string_parse_utf8_with_len, cgodot_string_parse_utf8_with_len_allocs = x.StringParseUtf8WithLen.PassRef()
	allocs57717e51.Borrow(cgodot_string_parse_utf8_with_len_allocs)

	var cgodot_string_chars_to_utf8_allocs *cgoAllocMap
	ref57717e51.godot_string_chars_to_utf8, cgodot_string_chars_to_utf8_allocs = x.StringCharsToUtf8.PassRef()
	allocs57717e51.Borrow(cgodot_string_chars_to_utf8_allocs)

	var cgodot_string_chars_to_utf8_with_len_allocs *cgoAllocMap
	ref57717e51.godot_string_chars_to_utf8_with_len, cgodot_string_chars_to_utf8_with_len_allocs = x.StringCharsToUtf8WithLen.PassRef()
	allocs57717e51.Borrow(cgodot_string_chars_to_utf8_with_len_allocs)

	var cgodot_string_hash_allocs *cgoAllocMap
	ref57717e51.godot_string_hash, cgodot_string_hash_allocs = x.StringHash.PassRef()
	allocs57717e51.Borrow(cgodot_string_hash_allocs)

	var cgodot_string_hash64_allocs *cgoAllocMap
	ref57717e51.godot_string_hash64, cgodot_string_hash64_allocs = x.StringHash64.PassRef()
	allocs57717e51.Borrow(cgodot_string_hash64_allocs)

	var cgodot_string_hash_chars_allocs *cgoAllocMap
	ref57717e51.godot_string_hash_chars, cgodot_string_hash_chars_allocs = x.StringHashChars.PassRef()
	allocs57717e51.Borrow(cgodot_string_hash_chars_allocs)

	var cgodot_string_hash_chars_with_len_allocs *cgoAllocMap
	ref57717e51.godot_string_hash_chars_with_len, cgodot_string_hash_chars_with_len_allocs = x.StringHashCharsWithLen.PassRef()
	allocs57717e51.Borrow(cgodot_string_hash_chars_with_len_allocs)

	var cgodot_string_hash_utf8_chars_allocs *cgoAllocMap
	ref57717e51.godot_string_hash_utf8_chars, cgodot_string_hash_utf8_chars_allocs = x.StringHashUtf8Chars.PassRef()
	allocs57717e51.Borrow(cgodot_string_hash_utf8_chars_allocs)

	var cgodot_string_hash_utf8_chars_with_len_allocs *cgoAllocMap
	ref57717e51.godot_string_hash_utf8_chars_with_len, cgodot_string_hash_utf8_chars_with_len_allocs = x.StringHashUtf8CharsWithLen.PassRef()
	allocs57717e51.Borrow(cgodot_string_hash_utf8_chars_with_len_allocs)

	var cgodot_string_md5_buffer_allocs *cgoAllocMap
	ref57717e51.godot_string_md5_buffer, cgodot_string_md5_buffer_allocs = x.StringMd5Buffer.PassRef()
	allocs57717e51.Borrow(cgodot_string_md5_buffer_allocs)

	var cgodot_string_md5_text_allocs *cgoAllocMap
	ref57717e51.godot_string_md5_text, cgodot_string_md5_text_allocs = x.StringMd5Text.PassRef()
	allocs57717e51.Borrow(cgodot_string_md5_text_allocs)

	var cgodot_string_sha256_buffer_allocs *cgoAllocMap
	ref57717e51.godot_string_sha256_buffer, cgodot_string_sha256_buffer_allocs = x.StringSha256Buffer.PassRef()
	allocs57717e51.Borrow(cgodot_string_sha256_buffer_allocs)

	var cgodot_string_sha256_text_allocs *cgoAllocMap
	ref57717e51.godot_string_sha256_text, cgodot_string_sha256_text_allocs = x.StringSha256Text.PassRef()
	allocs57717e51.Borrow(cgodot_string_sha256_text_allocs)

	var cgodot_string_empty_allocs *cgoAllocMap
	ref57717e51.godot_string_empty, cgodot_string_empty_allocs = x.StringEmpty.PassRef()
	allocs57717e51.Borrow(cgodot_string_empty_allocs)

	var cgodot_string_get_base_dir_allocs *cgoAllocMap
	ref57717e51.godot_string_get_base_dir, cgodot_string_get_base_dir_allocs = x.StringGetBaseDir.PassRef()
	allocs57717e51.Borrow(cgodot_string_get_base_dir_allocs)

	var cgodot_string_get_file_allocs *cgoAllocMap
	ref57717e51.godot_string_get_file, cgodot_string_get_file_allocs = x.StringGetFile.PassRef()
	allocs57717e51.Borrow(cgodot_string_get_file_allocs)

	var cgodot_string_humanize_size_allocs *cgoAllocMap
	ref57717e51.godot_string_humanize_size, cgodot_string_humanize_size_allocs = x.StringHumanizeSize.PassRef()
	allocs57717e51.Borrow(cgodot_string_humanize_size_allocs)

	var cgodot_string_is_abs_path_allocs *cgoAllocMap
	ref57717e51.godot_string_is_abs_path, cgodot_string_is_abs_path_allocs = x.StringIsAbsPath.PassRef()
	allocs57717e51.Borrow(cgodot_string_is_abs_path_allocs)

	var cgodot_string_is_rel_path_allocs *cgoAllocMap
	ref57717e51.godot_string_is_rel_path, cgodot_string_is_rel_path_allocs = x.StringIsRelPath.PassRef()
	allocs57717e51.Borrow(cgodot_string_is_rel_path_allocs)

	var cgodot_string_is_resource_file_allocs *cgoAllocMap
	ref57717e51.godot_string_is_resource_file, cgodot_string_is_resource_file_allocs = x.StringIsResourceFile.PassRef()
	allocs57717e51.Borrow(cgodot_string_is_resource_file_allocs)

	var cgodot_string_path_to_allocs *cgoAllocMap
	ref57717e51.godot_string_path_to, cgodot_string_path_to_allocs = x.StringPathTo.PassRef()
	allocs57717e51.Borrow(cgodot_string_path_to_allocs)

	var cgodot_string_path_to_file_allocs *cgoAllocMap
	ref57717e51.godot_string_path_to_file, cgodot_string_path_to_file_allocs = x.StringPathToFile.PassRef()
	allocs57717e51.Borrow(cgodot_string_path_to_file_allocs)

	var cgodot_string_simplify_path_allocs *cgoAllocMap
	ref57717e51.godot_string_simplify_path, cgodot_string_simplify_path_allocs = x.StringSimplifyPath.PassRef()
	allocs57717e51.Borrow(cgodot_string_simplify_path_allocs)

	var cgodot_string_c_escape_allocs *cgoAllocMap
	ref57717e51.godot_string_c_escape, cgodot_string_c_escape_allocs = x.StringCEscape.PassRef()
	allocs57717e51.Borrow(cgodot_string_c_escape_allocs)

	var cgodot_string_c_escape_multiline_allocs *cgoAllocMap
	ref57717e51.godot_string_c_escape_multiline, cgodot_string_c_escape_multiline_allocs = x.StringCEscapeMultiline.PassRef()
	allocs57717e51.Borrow(cgodot_string_c_escape_multiline_allocs)

	var cgodot_string_c_unescape_allocs *cgoAllocMap
	ref57717e51.godot_string_c_unescape, cgodot_string_c_unescape_allocs = x.StringCUnescape.PassRef()
	allocs57717e51.Borrow(cgodot_string_c_unescape_allocs)

	var cgodot_string_http_escape_allocs *cgoAllocMap
	ref57717e51.godot_string_http_escape, cgodot_string_http_escape_allocs = x.StringHttpEscape.PassRef()
	allocs57717e51.Borrow(cgodot_string_http_escape_allocs)

	var cgodot_string_http_unescape_allocs *cgoAllocMap
	ref57717e51.godot_string_http_unescape, cgodot_string_http_unescape_allocs = x.StringHttpUnescape.PassRef()
	allocs57717e51.Borrow(cgodot_string_http_unescape_allocs)

	var cgodot_string_json_escape_allocs *cgoAllocMap
	ref57717e51.godot_string_json_escape, cgodot_string_json_escape_allocs = x.StringJsonEscape.PassRef()
	allocs57717e51.Borrow(cgodot_string_json_escape_allocs)

	var cgodot_string_word_wrap_allocs *cgoAllocMap
	ref57717e51.godot_string_word_wrap, cgodot_string_word_wrap_allocs = x.StringWordWrap.PassRef()
	allocs57717e51.Borrow(cgodot_string_word_wrap_allocs)

	var cgodot_string_xml_escape_allocs *cgoAllocMap
	ref57717e51.godot_string_xml_escape, cgodot_string_xml_escape_allocs = x.StringXmlEscape.PassRef()
	allocs57717e51.Borrow(cgodot_string_xml_escape_allocs)

	var cgodot_string_xml_escape_with_quotes_allocs *cgoAllocMap
	ref57717e51.godot_string_xml_escape_with_quotes, cgodot_string_xml_escape_with_quotes_allocs = x.StringXmlEscapeWithQuotes.PassRef()
	allocs57717e51.Borrow(cgodot_string_xml_escape_with_quotes_allocs)

	var cgodot_string_xml_unescape_allocs *cgoAllocMap
	ref57717e51.godot_string_xml_unescape, cgodot_string_xml_unescape_allocs = x.StringXmlUnescape.PassRef()
	allocs57717e51.Borrow(cgodot_string_xml_unescape_allocs)

	var cgodot_string_percent_decode_allocs *cgoAllocMap
	ref57717e51.godot_string_percent_decode, cgodot_string_percent_decode_allocs = x.StringPercentDecode.PassRef()
	allocs57717e51.Borrow(cgodot_string_percent_decode_allocs)

	var cgodot_string_percent_encode_allocs *cgoAllocMap
	ref57717e51.godot_string_percent_encode, cgodot_string_percent_encode_allocs = x.StringPercentEncode.PassRef()
	allocs57717e51.Borrow(cgodot_string_percent_encode_allocs)

	var cgodot_string_is_valid_float_allocs *cgoAllocMap
	ref57717e51.godot_string_is_valid_float, cgodot_string_is_valid_float_allocs = x.StringIsValidFloat.PassRef()
	allocs57717e51.Borrow(cgodot_string_is_valid_float_allocs)

	var cgodot_string_is_valid_hex_number_allocs *cgoAllocMap
	ref57717e51.godot_string_is_valid_hex_number, cgodot_string_is_valid_hex_number_allocs = x.StringIsValidHexNumber.PassRef()
	allocs57717e51.Borrow(cgodot_string_is_valid_hex_number_allocs)

	var cgodot_string_is_valid_html_color_allocs *cgoAllocMap
	ref57717e51.godot_string_is_valid_html_color, cgodot_string_is_valid_html_color_allocs = x.StringIsValidHtmlColor.PassRef()
	allocs57717e51.Borrow(cgodot_string_is_valid_html_color_allocs)

	var cgodot_string_is_valid_identifier_allocs *cgoAllocMap
	ref57717e51.godot_string_is_valid_identifier, cgodot_string_is_valid_identifier_allocs = x.StringIsValidIdentifier.PassRef()
	allocs57717e51.Borrow(cgodot_string_is_valid_identifier_allocs)

	var cgodot_string_is_valid_integer_allocs *cgoAllocMap
	ref57717e51.godot_string_is_valid_integer, cgodot_string_is_valid_integer_allocs = x.StringIsValidInteger.PassRef()
	allocs57717e51.Borrow(cgodot_string_is_valid_integer_allocs)

	var cgodot_string_is_valid_ip_address_allocs *cgoAllocMap
	ref57717e51.godot_string_is_valid_ip_address, cgodot_string_is_valid_ip_address_allocs = x.StringIsValidIpAddress.PassRef()
	allocs57717e51.Borrow(cgodot_string_is_valid_ip_address_allocs)

	var cgodot_string_destroy_allocs *cgoAllocMap
	ref57717e51.godot_string_destroy, cgodot_string_destroy_allocs = x.StringDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_string_destroy_allocs)

	var cgodot_string_name_new_allocs *cgoAllocMap
	ref57717e51.godot_string_name_new, cgodot_string_name_new_allocs = x.StringNameNew.PassRef()
	allocs57717e51.Borrow(cgodot_string_name_new_allocs)

	var cgodot_string_name_new_data_allocs *cgoAllocMap
	ref57717e51.godot_string_name_new_data, cgodot_string_name_new_data_allocs = x.StringNameNewData.PassRef()
	allocs57717e51.Borrow(cgodot_string_name_new_data_allocs)

	var cgodot_string_name_get_name_allocs *cgoAllocMap
	ref57717e51.godot_string_name_get_name, cgodot_string_name_get_name_allocs = x.StringNameGetName.PassRef()
	allocs57717e51.Borrow(cgodot_string_name_get_name_allocs)

	var cgodot_string_name_get_hash_allocs *cgoAllocMap
	ref57717e51.godot_string_name_get_hash, cgodot_string_name_get_hash_allocs = x.StringNameGetHash.PassRef()
	allocs57717e51.Borrow(cgodot_string_name_get_hash_allocs)

	var cgodot_string_name_get_data_unique_pointer_allocs *cgoAllocMap
	ref57717e51.godot_string_name_get_data_unique_pointer, cgodot_string_name_get_data_unique_pointer_allocs = x.StringNameGetDataUniquePointer.PassRef()
	allocs57717e51.Borrow(cgodot_string_name_get_data_unique_pointer_allocs)

	var cgodot_string_name_operator_equal_allocs *cgoAllocMap
	ref57717e51.godot_string_name_operator_equal, cgodot_string_name_operator_equal_allocs = x.StringNameOperatorEqual.PassRef()
	allocs57717e51.Borrow(cgodot_string_name_operator_equal_allocs)

	var cgodot_string_name_operator_less_allocs *cgoAllocMap
	ref57717e51.godot_string_name_operator_less, cgodot_string_name_operator_less_allocs = x.StringNameOperatorLess.PassRef()
	allocs57717e51.Borrow(cgodot_string_name_operator_less_allocs)

	var cgodot_string_name_destroy_allocs *cgoAllocMap
	ref57717e51.godot_string_name_destroy, cgodot_string_name_destroy_allocs = x.StringNameDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_string_name_destroy_allocs)

	var cgodot_object_destroy_allocs *cgoAllocMap
	ref57717e51.godot_object_destroy, cgodot_object_destroy_allocs = x.ObjectDestroy.PassRef()
	allocs57717e51.Borrow(cgodot_object_destroy_allocs)

	var cgodot_global_get_singleton_allocs *cgoAllocMap
	ref57717e51.godot_global_get_singleton, cgodot_global_get_singleton_allocs = x.GlobalGetSingleton.PassRef()
	allocs57717e51.Borrow(cgodot_global_get_singleton_allocs)

	var cgodot_method_bind_get_method_allocs *cgoAllocMap
	ref57717e51.godot_method_bind_get_method, cgodot_method_bind_get_method_allocs = x.MethodBindGetMethod.PassRef()
	allocs57717e51.Borrow(cgodot_method_bind_get_method_allocs)

	var cgodot_method_bind_ptrcall_allocs *cgoAllocMap
	ref57717e51.godot_method_bind_ptrcall, cgodot_method_bind_ptrcall_allocs = x.MethodBindPtrcall.PassRef()
	allocs57717e51.Borrow(cgodot_method_bind_ptrcall_allocs)

	var cgodot_method_bind_call_allocs *cgoAllocMap
	ref57717e51.godot_method_bind_call, cgodot_method_bind_call_allocs = x.MethodBindCall.PassRef()
	allocs57717e51.Borrow(cgodot_method_bind_call_allocs)

	var cgodot_get_class_constructor_allocs *cgoAllocMap
	ref57717e51.godot_get_class_constructor, cgodot_get_class_constructor_allocs = x.GetClassConstructor.PassValue()
	allocs57717e51.Borrow(cgodot_get_class_constructor_allocs)

	var cgodot_register_native_call_type_allocs *cgoAllocMap
	ref57717e51.godot_register_native_call_type, cgodot_register_native_call_type_allocs = x.RegisterNativeCallType.PassRef()
	allocs57717e51.Borrow(cgodot_register_native_call_type_allocs)

	var cgodot_alloc_allocs *cgoAllocMap
	ref57717e51.godot_alloc, cgodot_alloc_allocs = x.Alloc.PassRef()
	allocs57717e51.Borrow(cgodot_alloc_allocs)

	var cgodot_realloc_allocs *cgoAllocMap
	ref57717e51.godot_realloc, cgodot_realloc_allocs = x.Realloc.PassRef()
	allocs57717e51.Borrow(cgodot_realloc_allocs)

	var cgodot_free_allocs *cgoAllocMap
	ref57717e51.godot_free, cgodot_free_allocs = x.Free.PassRef()
	allocs57717e51.Borrow(cgodot_free_allocs)

	var cgodot_print_error_allocs *cgoAllocMap
	ref57717e51.godot_print_error, cgodot_print_error_allocs = x.PrintError.PassRef()
	allocs57717e51.Borrow(cgodot_print_error_allocs)

	var cgodot_print_warning_allocs *cgoAllocMap
	ref57717e51.godot_print_warning, cgodot_print_warning_allocs = x.PrintWarning.PassRef()
	allocs57717e51.Borrow(cgodot_print_warning_allocs)

	var cgodot_print_allocs *cgoAllocMap
	ref57717e51.godot_print, cgodot_print_allocs = x.Print.PassRef()
	allocs57717e51.Borrow(cgodot_print_allocs)

	x.ref57717e51 = ref57717e51
	x.allocs57717e51 = allocs57717e51
	return ref57717e51, allocs57717e51

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GdnativeCoreApiStruct) PassValue() (C.godot_gdnative_core_api_struct, *cgoAllocMap) {
	if x.ref57717e51 != nil {
		return *x.ref57717e51, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GdnativeCoreApiStruct) Deref() {
	if x.ref57717e51 == nil {
		return
	}
	x.Type = (uint32)(x.ref57717e51._type)
	x.Version = *NewGdnativeApiVersionRef(unsafe.Pointer(&x.ref57717e51.version))
	packSGdnativeApiStruct(x.Next, x.ref57717e51.next)
	x.NumExtensions = (uint32)(x.ref57717e51.num_extensions)
	packSSGdnativeApiStruct(x.Extensions, x.ref57717e51.extensions)
	x.ColorNewRgba = NewRef(unsafe.Pointer(x.ref57717e51.godot_color_new_rgba))
	x.ColorNewRgb = NewRef(unsafe.Pointer(x.ref57717e51.godot_color_new_rgb))
	x.ColorGetR = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_color_get_r))
	x.ColorSetR = NewRef(unsafe.Pointer(x.ref57717e51.godot_color_set_r))
	x.ColorGetG = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_color_get_g))
	x.ColorSetG = NewRef(unsafe.Pointer(x.ref57717e51.godot_color_set_g))
	x.ColorGetB = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_color_get_b))
	x.ColorSetB = NewRef(unsafe.Pointer(x.ref57717e51.godot_color_set_b))
	x.ColorGetA = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_color_get_a))
	x.ColorSetA = NewRef(unsafe.Pointer(x.ref57717e51.godot_color_set_a))
	x.ColorGetH = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_color_get_h))
	x.ColorGetS = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_color_get_s))
	x.ColorGetV = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_color_get_v))
	x.ColorAsString = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_color_as_string))
	x.ColorToRgba32 = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_color_to_rgba32))
	x.ColorToArgb32 = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_color_to_argb32))
	x.ColorGray = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_color_gray))
	x.ColorInverted = NewColorRef(unsafe.Pointer(x.ref57717e51.godot_color_inverted))
	x.ColorContrasted = NewColorRef(unsafe.Pointer(x.ref57717e51.godot_color_contrasted))
	x.ColorLinearInterpolate = NewColorRef(unsafe.Pointer(x.ref57717e51.godot_color_linear_interpolate))
	x.ColorBlend = NewColorRef(unsafe.Pointer(x.ref57717e51.godot_color_blend))
	x.ColorToHtml = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_color_to_html))
	x.ColorOperatorEqual = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_color_operator_equal))
	x.ColorOperatorLess = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_color_operator_less))
	x.Vector2New = NewRef(unsafe.Pointer(x.ref57717e51.godot_vector2_new))
	x.Vector2AsString = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_vector2_as_string))
	x.Vector2Normalized = NewVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_normalized))
	x.Vector2Length = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_vector2_length))
	x.Vector2Angle = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_vector2_angle))
	x.Vector2LengthSquared = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_vector2_length_squared))
	x.Vector2IsNormalized = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_vector2_is_normalized))
	x.Vector2DistanceTo = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_vector2_distance_to))
	x.Vector2DistanceSquaredTo = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_vector2_distance_squared_to))
	x.Vector2AngleTo = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_vector2_angle_to))
	x.Vector2AngleToPoint = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_vector2_angle_to_point))
	x.Vector2LinearInterpolate = NewVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_linear_interpolate))
	x.Vector2CubicInterpolate = NewVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_cubic_interpolate))
	x.Vector2Rotated = NewVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_rotated))
	x.Vector2Tangent = NewVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_tangent))
	x.Vector2Floor = NewVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_floor))
	x.Vector2Snapped = NewVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_snapped))
	x.Vector2Aspect = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_vector2_aspect))
	x.Vector2Dot = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_vector2_dot))
	x.Vector2Slide = NewVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_slide))
	x.Vector2Bounce = NewVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_bounce))
	x.Vector2Reflect = NewVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_reflect))
	x.Vector2Abs = NewVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_abs))
	x.Vector2Clamped = NewVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_clamped))
	x.Vector2OperatorAdd = NewVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_operator_add))
	x.Vector2OperatorSubstract = NewVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_operator_substract))
	x.Vector2OperatorMultiplyVector = NewVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_operator_multiply_vector))
	x.Vector2OperatorMultiplyScalar = NewVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_operator_multiply_scalar))
	x.Vector2OperatorDivideVector = NewVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_operator_divide_vector))
	x.Vector2OperatorDivideScalar = NewVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_operator_divide_scalar))
	x.Vector2OperatorEqual = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_vector2_operator_equal))
	x.Vector2OperatorLess = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_vector2_operator_less))
	x.Vector2OperatorNeg = NewVector2Ref(unsafe.Pointer(x.ref57717e51.godot_vector2_operator_neg))
	x.Vector2SetX = NewRef(unsafe.Pointer(x.ref57717e51.godot_vector2_set_x))
	x.Vector2SetY = NewRef(unsafe.Pointer(x.ref57717e51.godot_vector2_set_y))
	x.Vector2GetX = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_vector2_get_x))
	x.Vector2GetY = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_vector2_get_y))
	x.QuatNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_quat_new))
	x.QuatNewWithAxisAngle = NewRef(unsafe.Pointer(x.ref57717e51.godot_quat_new_with_axis_angle))
	x.QuatGetX = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_quat_get_x))
	x.QuatSetX = NewRef(unsafe.Pointer(x.ref57717e51.godot_quat_set_x))
	x.QuatGetY = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_quat_get_y))
	x.QuatSetY = NewRef(unsafe.Pointer(x.ref57717e51.godot_quat_set_y))
	x.QuatGetZ = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_quat_get_z))
	x.QuatSetZ = NewRef(unsafe.Pointer(x.ref57717e51.godot_quat_set_z))
	x.QuatGetW = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_quat_get_w))
	x.QuatSetW = NewRef(unsafe.Pointer(x.ref57717e51.godot_quat_set_w))
	x.QuatAsString = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_quat_as_string))
	x.QuatLength = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_quat_length))
	x.QuatLengthSquared = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_quat_length_squared))
	x.QuatNormalized = NewQuatRef(unsafe.Pointer(x.ref57717e51.godot_quat_normalized))
	x.QuatIsNormalized = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_quat_is_normalized))
	x.QuatInverse = NewQuatRef(unsafe.Pointer(x.ref57717e51.godot_quat_inverse))
	x.QuatDot = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_quat_dot))
	x.QuatXform = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_quat_xform))
	x.QuatSlerp = NewQuatRef(unsafe.Pointer(x.ref57717e51.godot_quat_slerp))
	x.QuatSlerpni = NewQuatRef(unsafe.Pointer(x.ref57717e51.godot_quat_slerpni))
	x.QuatCubicSlerp = NewQuatRef(unsafe.Pointer(x.ref57717e51.godot_quat_cubic_slerp))
	x.QuatOperatorMultiply = NewQuatRef(unsafe.Pointer(x.ref57717e51.godot_quat_operator_multiply))
	x.QuatOperatorAdd = NewQuatRef(unsafe.Pointer(x.ref57717e51.godot_quat_operator_add))
	x.QuatOperatorSubstract = NewQuatRef(unsafe.Pointer(x.ref57717e51.godot_quat_operator_substract))
	x.QuatOperatorDivide = NewQuatRef(unsafe.Pointer(x.ref57717e51.godot_quat_operator_divide))
	x.QuatOperatorEqual = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_quat_operator_equal))
	x.QuatOperatorNeg = NewQuatRef(unsafe.Pointer(x.ref57717e51.godot_quat_operator_neg))
	x.BasisNewWithRows = NewRef(unsafe.Pointer(x.ref57717e51.godot_basis_new_with_rows))
	x.BasisNewWithAxisAndAngle = NewRef(unsafe.Pointer(x.ref57717e51.godot_basis_new_with_axis_and_angle))
	x.BasisNewWithEuler = NewRef(unsafe.Pointer(x.ref57717e51.godot_basis_new_with_euler))
	x.BasisAsString = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_basis_as_string))
	x.BasisInverse = NewBasisRef(unsafe.Pointer(x.ref57717e51.godot_basis_inverse))
	x.BasisTransposed = NewBasisRef(unsafe.Pointer(x.ref57717e51.godot_basis_transposed))
	x.BasisOrthonormalized = NewBasisRef(unsafe.Pointer(x.ref57717e51.godot_basis_orthonormalized))
	x.BasisDeterminant = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_basis_determinant))
	x.BasisRotated = NewBasisRef(unsafe.Pointer(x.ref57717e51.godot_basis_rotated))
	x.BasisScaled = NewBasisRef(unsafe.Pointer(x.ref57717e51.godot_basis_scaled))
	x.BasisGetScale = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_basis_get_scale))
	x.BasisGetEuler = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_basis_get_euler))
	x.BasisTdotx = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_basis_tdotx))
	x.BasisTdoty = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_basis_tdoty))
	x.BasisTdotz = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_basis_tdotz))
	x.BasisXform = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_basis_xform))
	x.BasisXformInv = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_basis_xform_inv))
	x.BasisGetOrthogonalIndex = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_basis_get_orthogonal_index))
	x.BasisNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_basis_new))
	x.BasisNewWithEulerQuat = NewRef(unsafe.Pointer(x.ref57717e51.godot_basis_new_with_euler_quat))
	x.BasisGetElements = NewRef(unsafe.Pointer(x.ref57717e51.godot_basis_get_elements))
	x.BasisGetAxis = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_basis_get_axis))
	x.BasisSetAxis = NewRef(unsafe.Pointer(x.ref57717e51.godot_basis_set_axis))
	x.BasisGetRow = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_basis_get_row))
	x.BasisSetRow = NewRef(unsafe.Pointer(x.ref57717e51.godot_basis_set_row))
	x.BasisOperatorEqual = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_basis_operator_equal))
	x.BasisOperatorAdd = NewBasisRef(unsafe.Pointer(x.ref57717e51.godot_basis_operator_add))
	x.BasisOperatorSubstract = NewBasisRef(unsafe.Pointer(x.ref57717e51.godot_basis_operator_substract))
	x.BasisOperatorMultiplyVector = NewBasisRef(unsafe.Pointer(x.ref57717e51.godot_basis_operator_multiply_vector))
	x.BasisOperatorMultiplyScalar = NewBasisRef(unsafe.Pointer(x.ref57717e51.godot_basis_operator_multiply_scalar))
	x.Vector3New = NewRef(unsafe.Pointer(x.ref57717e51.godot_vector3_new))
	x.Vector3AsString = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_vector3_as_string))
	x.Vector3MinAxis = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_vector3_min_axis))
	x.Vector3MaxAxis = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_vector3_max_axis))
	x.Vector3Length = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_vector3_length))
	x.Vector3LengthSquared = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_vector3_length_squared))
	x.Vector3IsNormalized = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_vector3_is_normalized))
	x.Vector3Normalized = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_normalized))
	x.Vector3Inverse = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_inverse))
	x.Vector3Snapped = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_snapped))
	x.Vector3Rotated = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_rotated))
	x.Vector3LinearInterpolate = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_linear_interpolate))
	x.Vector3CubicInterpolate = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_cubic_interpolate))
	x.Vector3Dot = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_vector3_dot))
	x.Vector3Cross = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_cross))
	x.Vector3Outer = NewBasisRef(unsafe.Pointer(x.ref57717e51.godot_vector3_outer))
	x.Vector3ToDiagonalMatrix = NewBasisRef(unsafe.Pointer(x.ref57717e51.godot_vector3_to_diagonal_matrix))
	x.Vector3Abs = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_abs))
	x.Vector3Floor = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_floor))
	x.Vector3Ceil = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_ceil))
	x.Vector3DistanceTo = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_vector3_distance_to))
	x.Vector3DistanceSquaredTo = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_vector3_distance_squared_to))
	x.Vector3AngleTo = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_vector3_angle_to))
	x.Vector3Slide = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_slide))
	x.Vector3Bounce = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_bounce))
	x.Vector3Reflect = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_reflect))
	x.Vector3OperatorAdd = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_operator_add))
	x.Vector3OperatorSubstract = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_operator_substract))
	x.Vector3OperatorMultiplyVector = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_operator_multiply_vector))
	x.Vector3OperatorMultiplyScalar = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_operator_multiply_scalar))
	x.Vector3OperatorDivideVector = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_operator_divide_vector))
	x.Vector3OperatorDivideScalar = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_operator_divide_scalar))
	x.Vector3OperatorEqual = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_vector3_operator_equal))
	x.Vector3OperatorLess = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_vector3_operator_less))
	x.Vector3OperatorNeg = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_vector3_operator_neg))
	x.Vector3SetAxis = NewRef(unsafe.Pointer(x.ref57717e51.godot_vector3_set_axis))
	x.Vector3GetAxis = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_vector3_get_axis))
	x.PoolByteArrayNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_new))
	x.PoolByteArrayNewCopy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_new_copy))
	x.PoolByteArrayNewWithArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_new_with_array))
	x.PoolByteArrayAppend = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_append))
	x.PoolByteArrayAppendArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_append_array))
	x.PoolByteArrayInsert = NewErrorRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_insert))
	x.PoolByteArrayInvert = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_invert))
	x.PoolByteArrayPushBack = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_push_back))
	x.PoolByteArrayRemove = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_remove))
	x.PoolByteArrayResize = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_resize))
	x.PoolByteArrayRead = NewPoolByteArrayReadAccessRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_read))
	x.PoolByteArrayWrite = NewPoolByteArrayWriteAccessRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_write))
	x.PoolByteArraySet = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_set))
	x.PoolByteArrayGet = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_get))
	x.PoolByteArraySize = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_size))
	x.PoolByteArrayDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_destroy))
	x.PoolIntArrayNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_new))
	x.PoolIntArrayNewCopy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_new_copy))
	x.PoolIntArrayNewWithArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_new_with_array))
	x.PoolIntArrayAppend = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_append))
	x.PoolIntArrayAppendArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_append_array))
	x.PoolIntArrayInsert = NewErrorRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_insert))
	x.PoolIntArrayInvert = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_invert))
	x.PoolIntArrayPushBack = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_push_back))
	x.PoolIntArrayRemove = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_remove))
	x.PoolIntArrayResize = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_resize))
	x.PoolIntArrayRead = NewPoolIntArrayReadAccessRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_read))
	x.PoolIntArrayWrite = NewPoolIntArrayWriteAccessRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_write))
	x.PoolIntArraySet = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_set))
	x.PoolIntArrayGet = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_get))
	x.PoolIntArraySize = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_size))
	x.PoolIntArrayDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_destroy))
	x.PoolRealArrayNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_new))
	x.PoolRealArrayNewCopy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_new_copy))
	x.PoolRealArrayNewWithArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_new_with_array))
	x.PoolRealArrayAppend = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_append))
	x.PoolRealArrayAppendArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_append_array))
	x.PoolRealArrayInsert = NewErrorRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_insert))
	x.PoolRealArrayInvert = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_invert))
	x.PoolRealArrayPushBack = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_push_back))
	x.PoolRealArrayRemove = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_remove))
	x.PoolRealArrayResize = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_resize))
	x.PoolRealArrayRead = NewPoolRealArrayReadAccessRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_read))
	x.PoolRealArrayWrite = NewPoolRealArrayWriteAccessRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_write))
	x.PoolRealArraySet = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_set))
	x.PoolRealArrayGet = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_get))
	x.PoolRealArraySize = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_size))
	x.PoolRealArrayDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_destroy))
	x.PoolStringArrayNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_new))
	x.PoolStringArrayNewCopy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_new_copy))
	x.PoolStringArrayNewWithArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_new_with_array))
	x.PoolStringArrayAppend = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_append))
	x.PoolStringArrayAppendArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_append_array))
	x.PoolStringArrayInsert = NewErrorRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_insert))
	x.PoolStringArrayInvert = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_invert))
	x.PoolStringArrayPushBack = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_push_back))
	x.PoolStringArrayRemove = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_remove))
	x.PoolStringArrayResize = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_resize))
	x.PoolStringArrayRead = NewPoolStringArrayReadAccessRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_read))
	x.PoolStringArrayWrite = NewPoolStringArrayWriteAccessRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_write))
	x.PoolStringArraySet = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_set))
	x.PoolStringArrayGet = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_get))
	x.PoolStringArraySize = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_size))
	x.PoolStringArrayDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_destroy))
	x.PoolVector2ArrayNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_new))
	x.PoolVector2ArrayNewCopy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_new_copy))
	x.PoolVector2ArrayNewWithArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_new_with_array))
	x.PoolVector2ArrayAppend = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_append))
	x.PoolVector2ArrayAppendArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_append_array))
	x.PoolVector2ArrayInsert = NewErrorRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_insert))
	x.PoolVector2ArrayInvert = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_invert))
	x.PoolVector2ArrayPushBack = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_push_back))
	x.PoolVector2ArrayRemove = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_remove))
	x.PoolVector2ArrayResize = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_resize))
	x.PoolVector2ArrayRead = NewPoolVector2ArrayReadAccessRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_read))
	x.PoolVector2ArrayWrite = NewPoolVector2ArrayWriteAccessRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_write))
	x.PoolVector2ArraySet = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_set))
	x.PoolVector2ArrayGet = NewVector2Ref(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_get))
	x.PoolVector2ArraySize = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_size))
	x.PoolVector2ArrayDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_destroy))
	x.PoolVector3ArrayNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_new))
	x.PoolVector3ArrayNewCopy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_new_copy))
	x.PoolVector3ArrayNewWithArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_new_with_array))
	x.PoolVector3ArrayAppend = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_append))
	x.PoolVector3ArrayAppendArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_append_array))
	x.PoolVector3ArrayInsert = NewErrorRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_insert))
	x.PoolVector3ArrayInvert = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_invert))
	x.PoolVector3ArrayPushBack = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_push_back))
	x.PoolVector3ArrayRemove = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_remove))
	x.PoolVector3ArrayResize = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_resize))
	x.PoolVector3ArrayRead = NewPoolVector3ArrayReadAccessRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_read))
	x.PoolVector3ArrayWrite = NewPoolVector3ArrayWriteAccessRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_write))
	x.PoolVector3ArraySet = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_set))
	x.PoolVector3ArrayGet = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_get))
	x.PoolVector3ArraySize = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_size))
	x.PoolVector3ArrayDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_destroy))
	x.PoolColorArrayNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_new))
	x.PoolColorArrayNewCopy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_new_copy))
	x.PoolColorArrayNewWithArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_new_with_array))
	x.PoolColorArrayAppend = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_append))
	x.PoolColorArrayAppendArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_append_array))
	x.PoolColorArrayInsert = NewErrorRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_insert))
	x.PoolColorArrayInvert = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_invert))
	x.PoolColorArrayPushBack = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_push_back))
	x.PoolColorArrayRemove = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_remove))
	x.PoolColorArrayResize = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_resize))
	x.PoolColorArrayRead = NewPoolColorArrayReadAccessRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_read))
	x.PoolColorArrayWrite = NewPoolColorArrayWriteAccessRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_write))
	x.PoolColorArraySet = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_set))
	x.PoolColorArrayGet = NewColorRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_get))
	x.PoolColorArraySize = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_size))
	x.PoolColorArrayDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_destroy))
	x.PoolByteArrayReadAccessPtr = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_read_access_ptr))
	x.PoolByteArrayReadAccessOperatorAssign = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_read_access_operator_assign))
	x.PoolByteArrayReadAccessDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_read_access_destroy))
	x.PoolIntArrayReadAccessPtr = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_read_access_ptr))
	x.PoolIntArrayReadAccessOperatorAssign = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_read_access_operator_assign))
	x.PoolIntArrayReadAccessDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_read_access_destroy))
	x.PoolRealArrayReadAccessPtr = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_read_access_ptr))
	x.PoolRealArrayReadAccessOperatorAssign = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_read_access_operator_assign))
	x.PoolRealArrayReadAccessDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_read_access_destroy))
	x.PoolStringArrayReadAccessPtr = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_read_access_ptr))
	x.PoolStringArrayReadAccessOperatorAssign = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_read_access_operator_assign))
	x.PoolStringArrayReadAccessDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_read_access_destroy))
	x.PoolVector2ArrayReadAccessPtr = NewVector2Ref(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_read_access_ptr))
	x.PoolVector2ArrayReadAccessOperatorAssign = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_read_access_operator_assign))
	x.PoolVector2ArrayReadAccessDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_read_access_destroy))
	x.PoolVector3ArrayReadAccessPtr = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_read_access_ptr))
	x.PoolVector3ArrayReadAccessOperatorAssign = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_read_access_operator_assign))
	x.PoolVector3ArrayReadAccessDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_read_access_destroy))
	x.PoolColorArrayReadAccessPtr = NewColorRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_read_access_ptr))
	x.PoolColorArrayReadAccessOperatorAssign = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_read_access_operator_assign))
	x.PoolColorArrayReadAccessDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_read_access_destroy))
	x.PoolByteArrayWriteAccessPtr = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_write_access_ptr))
	x.PoolByteArrayWriteAccessOperatorAssign = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_write_access_operator_assign))
	x.PoolByteArrayWriteAccessDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_byte_array_write_access_destroy))
	x.PoolIntArrayWriteAccessPtr = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_write_access_ptr))
	x.PoolIntArrayWriteAccessOperatorAssign = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_write_access_operator_assign))
	x.PoolIntArrayWriteAccessDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_int_array_write_access_destroy))
	x.PoolRealArrayWriteAccessPtr = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_write_access_ptr))
	x.PoolRealArrayWriteAccessOperatorAssign = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_write_access_operator_assign))
	x.PoolRealArrayWriteAccessDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_real_array_write_access_destroy))
	x.PoolStringArrayWriteAccessPtr = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_write_access_ptr))
	x.PoolStringArrayWriteAccessOperatorAssign = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_write_access_operator_assign))
	x.PoolStringArrayWriteAccessDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_string_array_write_access_destroy))
	x.PoolVector2ArrayWriteAccessPtr = NewVector2Ref(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_write_access_ptr))
	x.PoolVector2ArrayWriteAccessOperatorAssign = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_write_access_operator_assign))
	x.PoolVector2ArrayWriteAccessDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector2_array_write_access_destroy))
	x.PoolVector3ArrayWriteAccessPtr = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_write_access_ptr))
	x.PoolVector3ArrayWriteAccessOperatorAssign = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_write_access_operator_assign))
	x.PoolVector3ArrayWriteAccessDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_vector3_array_write_access_destroy))
	x.PoolColorArrayWriteAccessPtr = NewColorRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_write_access_ptr))
	x.PoolColorArrayWriteAccessOperatorAssign = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_write_access_operator_assign))
	x.PoolColorArrayWriteAccessDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_pool_color_array_write_access_destroy))
	x.ArrayNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_new))
	x.ArrayNewCopy = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_new_copy))
	x.ArrayNewPoolColorArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_new_pool_color_array))
	x.ArrayNewPoolVector3Array = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_new_pool_vector3_array))
	x.ArrayNewPoolVector2Array = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_new_pool_vector2_array))
	x.ArrayNewPoolStringArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_new_pool_string_array))
	x.ArrayNewPoolRealArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_new_pool_real_array))
	x.ArrayNewPoolIntArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_new_pool_int_array))
	x.ArrayNewPoolByteArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_new_pool_byte_array))
	x.ArraySet = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_set))
	x.ArrayGet = NewVariantRef(unsafe.Pointer(x.ref57717e51.godot_array_get))
	x.ArrayOperatorIndex = NewVariantRef(unsafe.Pointer(x.ref57717e51.godot_array_operator_index))
	x.ArrayOperatorIndexConst = NewVariantRef(unsafe.Pointer(x.ref57717e51.godot_array_operator_index_const))
	x.ArrayAppend = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_append))
	x.ArrayClear = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_clear))
	x.ArrayCount = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_array_count))
	x.ArrayEmpty = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_array_empty))
	x.ArrayErase = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_erase))
	x.ArrayFront = NewVariantRef(unsafe.Pointer(x.ref57717e51.godot_array_front))
	x.ArrayBack = NewVariantRef(unsafe.Pointer(x.ref57717e51.godot_array_back))
	x.ArrayFind = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_array_find))
	x.ArrayFindLast = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_array_find_last))
	x.ArrayHas = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_array_has))
	x.ArrayHash = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_array_hash))
	x.ArrayInsert = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_insert))
	x.ArrayInvert = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_invert))
	x.ArrayPopBack = NewVariantRef(unsafe.Pointer(x.ref57717e51.godot_array_pop_back))
	x.ArrayPopFront = NewVariantRef(unsafe.Pointer(x.ref57717e51.godot_array_pop_front))
	x.ArrayPushBack = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_push_back))
	x.ArrayPushFront = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_push_front))
	x.ArrayRemove = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_remove))
	x.ArrayResize = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_resize))
	x.ArrayRfind = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_array_rfind))
	x.ArraySize = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_array_size))
	x.ArraySort = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_sort))
	x.ArraySortCustom = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_sort_custom))
	x.ArrayBsearch = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_array_bsearch))
	x.ArrayBsearchCustom = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_array_bsearch_custom))
	x.ArrayDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_array_destroy))
	x.DictionaryNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_new))
	x.DictionaryNewCopy = NewRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_new_copy))
	x.DictionaryDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_destroy))
	x.DictionarySize = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_size))
	x.DictionaryEmpty = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_empty))
	x.DictionaryClear = NewRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_clear))
	x.DictionaryHas = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_has))
	x.DictionaryHasAll = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_has_all))
	x.DictionaryErase = NewRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_erase))
	x.DictionaryHash = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_hash))
	x.DictionaryKeys = NewArrayRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_keys))
	x.DictionaryValues = NewArrayRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_values))
	x.DictionaryGet = NewVariantRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_get))
	x.DictionarySet = NewRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_set))
	x.DictionaryOperatorIndex = NewVariantRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_operator_index))
	x.DictionaryOperatorIndexConst = NewVariantRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_operator_index_const))
	x.DictionaryNext = NewVariantRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_next))
	x.DictionaryOperatorEqual = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_operator_equal))
	x.DictionaryToJson = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_dictionary_to_json))
	x.NodePathNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_node_path_new))
	x.NodePathNewCopy = NewRef(unsafe.Pointer(x.ref57717e51.godot_node_path_new_copy))
	x.NodePathDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_node_path_destroy))
	x.NodePathAsString = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_node_path_as_string))
	x.NodePathIsAbsolute = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_node_path_is_absolute))
	x.NodePathGetNameCount = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_node_path_get_name_count))
	x.NodePathGetName = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_node_path_get_name))
	x.NodePathGetSubnameCount = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_node_path_get_subname_count))
	x.NodePathGetSubname = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_node_path_get_subname))
	x.NodePathGetConcatenatedSubnames = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_node_path_get_concatenated_subnames))
	x.NodePathIsEmpty = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_node_path_is_empty))
	x.NodePathOperatorEqual = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_node_path_operator_equal))
	x.PlaneNewWithReals = NewRef(unsafe.Pointer(x.ref57717e51.godot_plane_new_with_reals))
	x.PlaneNewWithVectors = NewRef(unsafe.Pointer(x.ref57717e51.godot_plane_new_with_vectors))
	x.PlaneNewWithNormal = NewRef(unsafe.Pointer(x.ref57717e51.godot_plane_new_with_normal))
	x.PlaneAsString = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_plane_as_string))
	x.PlaneNormalized = NewPlaneRef(unsafe.Pointer(x.ref57717e51.godot_plane_normalized))
	x.PlaneCenter = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_plane_center))
	x.PlaneGetAnyPoint = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_plane_get_any_point))
	x.PlaneIsPointOver = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_plane_is_point_over))
	x.PlaneDistanceTo = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_plane_distance_to))
	x.PlaneHasPoint = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_plane_has_point))
	x.PlaneProject = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_plane_project))
	x.PlaneIntersect3 = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_plane_intersect_3))
	x.PlaneIntersectsRay = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_plane_intersects_ray))
	x.PlaneIntersectsSegment = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_plane_intersects_segment))
	x.PlaneOperatorNeg = NewPlaneRef(unsafe.Pointer(x.ref57717e51.godot_plane_operator_neg))
	x.PlaneOperatorEqual = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_plane_operator_equal))
	x.PlaneSetNormal = NewRef(unsafe.Pointer(x.ref57717e51.godot_plane_set_normal))
	x.PlaneGetNormal = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_plane_get_normal))
	x.PlaneGetD = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_plane_get_d))
	x.PlaneSetD = NewRef(unsafe.Pointer(x.ref57717e51.godot_plane_set_d))
	x.Rect2NewWithPositionAndSize = NewRef(unsafe.Pointer(x.ref57717e51.godot_rect2_new_with_position_and_size))
	x.Rect2New = NewRef(unsafe.Pointer(x.ref57717e51.godot_rect2_new))
	x.Rect2AsString = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_rect2_as_string))
	x.Rect2GetArea = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_rect2_get_area))
	x.Rect2Intersects = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_rect2_intersects))
	x.Rect2Encloses = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_rect2_encloses))
	x.Rect2HasNoArea = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_rect2_has_no_area))
	x.Rect2Clip = NewRect2Ref(unsafe.Pointer(x.ref57717e51.godot_rect2_clip))
	x.Rect2Merge = NewRect2Ref(unsafe.Pointer(x.ref57717e51.godot_rect2_merge))
	x.Rect2HasPoint = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_rect2_has_point))
	x.Rect2Grow = NewRect2Ref(unsafe.Pointer(x.ref57717e51.godot_rect2_grow))
	x.Rect2Expand = NewRect2Ref(unsafe.Pointer(x.ref57717e51.godot_rect2_expand))
	x.Rect2OperatorEqual = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_rect2_operator_equal))
	x.Rect2GetPosition = NewVector2Ref(unsafe.Pointer(x.ref57717e51.godot_rect2_get_position))
	x.Rect2GetSize = NewVector2Ref(unsafe.Pointer(x.ref57717e51.godot_rect2_get_size))
	x.Rect2SetPosition = NewRef(unsafe.Pointer(x.ref57717e51.godot_rect2_set_position))
	x.Rect2SetSize = NewRef(unsafe.Pointer(x.ref57717e51.godot_rect2_set_size))
	x.AabbNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_aabb_new))
	x.AabbGetPosition = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_aabb_get_position))
	x.AabbSetPosition = NewRef(unsafe.Pointer(x.ref57717e51.godot_aabb_set_position))
	x.AabbGetSize = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_aabb_get_size))
	x.AabbSetSize = NewRef(unsafe.Pointer(x.ref57717e51.godot_aabb_set_size))
	x.AabbAsString = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_aabb_as_string))
	x.AabbGetArea = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_aabb_get_area))
	x.AabbHasNoArea = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_aabb_has_no_area))
	x.AabbHasNoSurface = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_aabb_has_no_surface))
	x.AabbIntersects = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_aabb_intersects))
	x.AabbEncloses = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_aabb_encloses))
	x.AabbMerge = NewAabbRef(unsafe.Pointer(x.ref57717e51.godot_aabb_merge))
	x.AabbIntersection = NewAabbRef(unsafe.Pointer(x.ref57717e51.godot_aabb_intersection))
	x.AabbIntersectsPlane = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_aabb_intersects_plane))
	x.AabbIntersectsSegment = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_aabb_intersects_segment))
	x.AabbHasPoint = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_aabb_has_point))
	x.AabbGetSupport = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_aabb_get_support))
	x.AabbGetLongestAxis = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_aabb_get_longest_axis))
	x.AabbGetLongestAxisIndex = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_aabb_get_longest_axis_index))
	x.AabbGetLongestAxisSize = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_aabb_get_longest_axis_size))
	x.AabbGetShortestAxis = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_aabb_get_shortest_axis))
	x.AabbGetShortestAxisIndex = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_aabb_get_shortest_axis_index))
	x.AabbGetShortestAxisSize = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_aabb_get_shortest_axis_size))
	x.AabbExpand = NewAabbRef(unsafe.Pointer(x.ref57717e51.godot_aabb_expand))
	x.AabbGrow = NewAabbRef(unsafe.Pointer(x.ref57717e51.godot_aabb_grow))
	x.AabbGetEndpoint = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_aabb_get_endpoint))
	x.AabbOperatorEqual = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_aabb_operator_equal))
	x.RidNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_rid_new))
	x.RidGetId = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_rid_get_id))
	x.RidNewWithResource = NewRef(unsafe.Pointer(x.ref57717e51.godot_rid_new_with_resource))
	x.RidOperatorEqual = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_rid_operator_equal))
	x.RidOperatorLess = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_rid_operator_less))
	x.TransformNewWithAxisOrigin = NewRef(unsafe.Pointer(x.ref57717e51.godot_transform_new_with_axis_origin))
	x.TransformNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_transform_new))
	x.TransformGetBasis = NewBasisRef(unsafe.Pointer(x.ref57717e51.godot_transform_get_basis))
	x.TransformSetBasis = NewRef(unsafe.Pointer(x.ref57717e51.godot_transform_set_basis))
	x.TransformGetOrigin = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_transform_get_origin))
	x.TransformSetOrigin = NewRef(unsafe.Pointer(x.ref57717e51.godot_transform_set_origin))
	x.TransformAsString = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_transform_as_string))
	x.TransformInverse = NewTransformRef(unsafe.Pointer(x.ref57717e51.godot_transform_inverse))
	x.TransformAffineInverse = NewTransformRef(unsafe.Pointer(x.ref57717e51.godot_transform_affine_inverse))
	x.TransformOrthonormalized = NewTransformRef(unsafe.Pointer(x.ref57717e51.godot_transform_orthonormalized))
	x.TransformRotated = NewTransformRef(unsafe.Pointer(x.ref57717e51.godot_transform_rotated))
	x.TransformScaled = NewTransformRef(unsafe.Pointer(x.ref57717e51.godot_transform_scaled))
	x.TransformTranslated = NewTransformRef(unsafe.Pointer(x.ref57717e51.godot_transform_translated))
	x.TransformLookingAt = NewTransformRef(unsafe.Pointer(x.ref57717e51.godot_transform_looking_at))
	x.TransformXformPlane = NewPlaneRef(unsafe.Pointer(x.ref57717e51.godot_transform_xform_plane))
	x.TransformXformInvPlane = NewPlaneRef(unsafe.Pointer(x.ref57717e51.godot_transform_xform_inv_plane))
	x.TransformNewIdentity = NewRef(unsafe.Pointer(x.ref57717e51.godot_transform_new_identity))
	x.TransformOperatorEqual = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_transform_operator_equal))
	x.TransformOperatorMultiply = NewTransformRef(unsafe.Pointer(x.ref57717e51.godot_transform_operator_multiply))
	x.TransformXformVector3 = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_transform_xform_vector3))
	x.TransformXformInvVector3 = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_transform_xform_inv_vector3))
	x.TransformXformAabb = NewAabbRef(unsafe.Pointer(x.ref57717e51.godot_transform_xform_aabb))
	x.TransformXformInvAabb = NewAabbRef(unsafe.Pointer(x.ref57717e51.godot_transform_xform_inv_aabb))
	x.Transform2dNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_transform2d_new))
	x.Transform2dNewAxisOrigin = NewRef(unsafe.Pointer(x.ref57717e51.godot_transform2d_new_axis_origin))
	x.Transform2dAsString = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_transform2d_as_string))
	x.Transform2dInverse = NewTransform2dRef(unsafe.Pointer(x.ref57717e51.godot_transform2d_inverse))
	x.Transform2dAffineInverse = NewTransform2dRef(unsafe.Pointer(x.ref57717e51.godot_transform2d_affine_inverse))
	x.Transform2dGetRotation = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_transform2d_get_rotation))
	x.Transform2dGetOrigin = NewVector2Ref(unsafe.Pointer(x.ref57717e51.godot_transform2d_get_origin))
	x.Transform2dGetScale = NewVector2Ref(unsafe.Pointer(x.ref57717e51.godot_transform2d_get_scale))
	x.Transform2dOrthonormalized = NewTransform2dRef(unsafe.Pointer(x.ref57717e51.godot_transform2d_orthonormalized))
	x.Transform2dRotated = NewTransform2dRef(unsafe.Pointer(x.ref57717e51.godot_transform2d_rotated))
	x.Transform2dScaled = NewTransform2dRef(unsafe.Pointer(x.ref57717e51.godot_transform2d_scaled))
	x.Transform2dTranslated = NewTransform2dRef(unsafe.Pointer(x.ref57717e51.godot_transform2d_translated))
	x.Transform2dXformVector2 = NewVector2Ref(unsafe.Pointer(x.ref57717e51.godot_transform2d_xform_vector2))
	x.Transform2dXformInvVector2 = NewVector2Ref(unsafe.Pointer(x.ref57717e51.godot_transform2d_xform_inv_vector2))
	x.Transform2dBasisXformVector2 = NewVector2Ref(unsafe.Pointer(x.ref57717e51.godot_transform2d_basis_xform_vector2))
	x.Transform2dBasisXformInvVector2 = NewVector2Ref(unsafe.Pointer(x.ref57717e51.godot_transform2d_basis_xform_inv_vector2))
	x.Transform2dInterpolateWith = NewTransform2dRef(unsafe.Pointer(x.ref57717e51.godot_transform2d_interpolate_with))
	x.Transform2dOperatorEqual = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_transform2d_operator_equal))
	x.Transform2dOperatorMultiply = NewTransform2dRef(unsafe.Pointer(x.ref57717e51.godot_transform2d_operator_multiply))
	x.Transform2dNewIdentity = NewRef(unsafe.Pointer(x.ref57717e51.godot_transform2d_new_identity))
	x.Transform2dXformRect2 = NewRect2Ref(unsafe.Pointer(x.ref57717e51.godot_transform2d_xform_rect2))
	x.Transform2dXformInvRect2 = NewRect2Ref(unsafe.Pointer(x.ref57717e51.godot_transform2d_xform_inv_rect2))
	x.VariantGetType = NewVariantTypeRef(unsafe.Pointer(x.ref57717e51.godot_variant_get_type))
	x.VariantNewCopy = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_copy))
	x.VariantNewNil = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_nil))
	x.VariantNewBool = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_bool))
	x.VariantNewUint = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_uint))
	x.VariantNewInt = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_int))
	x.VariantNewReal = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_real))
	x.VariantNewString = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_string))
	x.VariantNewVector2 = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_vector2))
	x.VariantNewRect2 = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_rect2))
	x.VariantNewVector3 = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_vector3))
	x.VariantNewTransform2d = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_transform2d))
	x.VariantNewPlane = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_plane))
	x.VariantNewQuat = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_quat))
	x.VariantNewAabb = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_aabb))
	x.VariantNewBasis = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_basis))
	x.VariantNewTransform = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_transform))
	x.VariantNewColor = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_color))
	x.VariantNewNodePath = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_node_path))
	x.VariantNewRid = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_rid))
	x.VariantNewObject = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_object))
	x.VariantNewDictionary = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_dictionary))
	x.VariantNewArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_array))
	x.VariantNewPoolByteArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_pool_byte_array))
	x.VariantNewPoolIntArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_pool_int_array))
	x.VariantNewPoolRealArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_pool_real_array))
	x.VariantNewPoolStringArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_pool_string_array))
	x.VariantNewPoolVector2Array = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_pool_vector2_array))
	x.VariantNewPoolVector3Array = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_pool_vector3_array))
	x.VariantNewPoolColorArray = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_new_pool_color_array))
	x.VariantAsBool = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_bool))
	x.VariantAsUint = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_uint))
	x.VariantAsInt = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_int))
	x.VariantAsReal = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_real))
	x.VariantAsString = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_string))
	x.VariantAsVector2 = NewVector2Ref(unsafe.Pointer(x.ref57717e51.godot_variant_as_vector2))
	x.VariantAsRect2 = NewRect2Ref(unsafe.Pointer(x.ref57717e51.godot_variant_as_rect2))
	x.VariantAsVector3 = NewVector3Ref(unsafe.Pointer(x.ref57717e51.godot_variant_as_vector3))
	x.VariantAsTransform2d = NewTransform2dRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_transform2d))
	x.VariantAsPlane = NewPlaneRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_plane))
	x.VariantAsQuat = NewQuatRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_quat))
	x.VariantAsAabb = NewAabbRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_aabb))
	x.VariantAsBasis = NewBasisRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_basis))
	x.VariantAsTransform = NewTransformRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_transform))
	x.VariantAsColor = NewColorRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_color))
	x.VariantAsNodePath = NewNodePathRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_node_path))
	x.VariantAsRid = NewRidRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_rid))
	x.VariantAsObject = NewObjectRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_object))
	x.VariantAsDictionary = NewDictionaryRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_dictionary))
	x.VariantAsArray = NewArrayRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_array))
	x.VariantAsPoolByteArray = NewPoolByteArrayRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_pool_byte_array))
	x.VariantAsPoolIntArray = NewPoolIntArrayRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_pool_int_array))
	x.VariantAsPoolRealArray = NewPoolRealArrayRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_pool_real_array))
	x.VariantAsPoolStringArray = NewPoolStringArrayRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_pool_string_array))
	x.VariantAsPoolVector2Array = NewPoolVector2ArrayRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_pool_vector2_array))
	x.VariantAsPoolVector3Array = NewPoolVector3ArrayRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_pool_vector3_array))
	x.VariantAsPoolColorArray = NewPoolColorArrayRef(unsafe.Pointer(x.ref57717e51.godot_variant_as_pool_color_array))
	x.VariantCall = NewVariantRef(unsafe.Pointer(x.ref57717e51.godot_variant_call))
	x.VariantHasMethod = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_variant_has_method))
	x.VariantOperatorEqual = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_variant_operator_equal))
	x.VariantOperatorLess = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_variant_operator_less))
	x.VariantHashCompare = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_variant_hash_compare))
	x.VariantBooleanize = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_variant_booleanize))
	x.VariantDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_variant_destroy))
	x.StringNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_new))
	x.StringNewCopy = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_new_copy))
	x.StringNewData = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_new_data))
	x.StringNewUnicodeData = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_new_unicode_data))
	x.StringGetData = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_get_data))
	x.StringOperatorIndex = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_operator_index))
	x.StringOperatorIndexConst = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_operator_index_const))
	x.StringUnicodeStr = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_unicode_str))
	x.StringOperatorEqual = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_operator_equal))
	x.StringOperatorLess = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_operator_less))
	x.StringOperatorPlus = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_operator_plus))
	x.StringLength = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_string_length))
	x.StringBeginsWith = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_begins_with))
	x.StringBeginsWithCharArray = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_begins_with_char_array))
	x.StringBigrams = NewArrayRef(unsafe.Pointer(x.ref57717e51.godot_string_bigrams))
	x.StringChr = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_chr))
	x.StringEndsWith = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_ends_with))
	x.StringFind = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_string_find))
	x.StringFindFrom = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_string_find_from))
	x.StringFindmk = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_string_findmk))
	x.StringFindmkFrom = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_string_findmk_from))
	x.StringFindmkFromInPlace = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_string_findmk_from_in_place))
	x.StringFindn = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_string_findn))
	x.StringFindnFrom = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_string_findn_from))
	x.StringFindLast = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_string_find_last))
	x.StringFormat = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_format))
	x.StringFormatWithCustomPlaceholder = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_format_with_custom_placeholder))
	x.StringHexEncodeBuffer = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_hex_encode_buffer))
	x.StringHexToInt = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_string_hex_to_int))
	x.StringHexToIntWithoutPrefix = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_string_hex_to_int_without_prefix))
	x.StringInsert = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_insert))
	x.StringIsNumeric = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_is_numeric))
	x.StringIsSubsequenceOf = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_is_subsequence_of))
	x.StringIsSubsequenceOfi = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_is_subsequence_ofi))
	x.StringLpad = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_lpad))
	x.StringLpadWithCustomCharacter = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_lpad_with_custom_character))
	x.StringMatch = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_match))
	x.StringMatchn = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_matchn))
	x.StringMd5 = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_md5))
	x.StringNum = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_num))
	x.StringNumInt64 = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_num_int64))
	x.StringNumInt64Capitalized = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_num_int64_capitalized))
	x.StringNumReal = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_num_real))
	x.StringNumScientific = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_num_scientific))
	x.StringNumWithDecimals = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_num_with_decimals))
	x.StringPadDecimals = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_pad_decimals))
	x.StringPadZeros = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_pad_zeros))
	x.StringReplaceFirst = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_replace_first))
	x.StringReplace = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_replace))
	x.StringReplacen = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_replacen))
	x.StringRfind = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_string_rfind))
	x.StringRfindn = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_string_rfindn))
	x.StringRfindFrom = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_string_rfind_from))
	x.StringRfindnFrom = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_string_rfindn_from))
	x.StringRpad = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_rpad))
	x.StringRpadWithCustomCharacter = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_rpad_with_custom_character))
	x.StringSimilarity = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_string_similarity))
	x.StringSprintf = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_sprintf))
	x.StringSubstr = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_substr))
	x.StringToDouble = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_to_double))
	x.StringToFloat = NewRealRef(unsafe.Pointer(x.ref57717e51.godot_string_to_float))
	x.StringToInt = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_string_to_int))
	x.StringCamelcaseToUnderscore = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_camelcase_to_underscore))
	x.StringCamelcaseToUnderscoreLowercased = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_camelcase_to_underscore_lowercased))
	x.StringCapitalize = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_capitalize))
	x.StringCharToDouble = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_char_to_double))
	x.StringCharToInt = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_string_char_to_int))
	x.StringWcharToInt = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_wchar_to_int))
	x.StringCharToIntWithLen = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_string_char_to_int_with_len))
	x.StringCharToInt64WithLen = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_char_to_int64_with_len))
	x.StringHexToInt64 = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_hex_to_int64))
	x.StringHexToInt64WithPrefix = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_hex_to_int64_with_prefix))
	x.StringToInt64 = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_to_int64))
	x.StringUnicodeCharToDouble = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_unicode_char_to_double))
	x.StringGetSliceCount = NewIntRef(unsafe.Pointer(x.ref57717e51.godot_string_get_slice_count))
	x.StringGetSlice = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_get_slice))
	x.StringGetSlicec = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_get_slicec))
	x.StringSplit = NewArrayRef(unsafe.Pointer(x.ref57717e51.godot_string_split))
	x.StringSplitAllowEmpty = NewArrayRef(unsafe.Pointer(x.ref57717e51.godot_string_split_allow_empty))
	x.StringSplitFloats = NewArrayRef(unsafe.Pointer(x.ref57717e51.godot_string_split_floats))
	x.StringSplitFloatsAllowsEmpty = NewArrayRef(unsafe.Pointer(x.ref57717e51.godot_string_split_floats_allows_empty))
	x.StringSplitFloatsMk = NewArrayRef(unsafe.Pointer(x.ref57717e51.godot_string_split_floats_mk))
	x.StringSplitFloatsMkAllowsEmpty = NewArrayRef(unsafe.Pointer(x.ref57717e51.godot_string_split_floats_mk_allows_empty))
	x.StringSplitInts = NewArrayRef(unsafe.Pointer(x.ref57717e51.godot_string_split_ints))
	x.StringSplitIntsAllowsEmpty = NewArrayRef(unsafe.Pointer(x.ref57717e51.godot_string_split_ints_allows_empty))
	x.StringSplitIntsMk = NewArrayRef(unsafe.Pointer(x.ref57717e51.godot_string_split_ints_mk))
	x.StringSplitIntsMkAllowsEmpty = NewArrayRef(unsafe.Pointer(x.ref57717e51.godot_string_split_ints_mk_allows_empty))
	x.StringSplitSpaces = NewArrayRef(unsafe.Pointer(x.ref57717e51.godot_string_split_spaces))
	x.StringCharLowercase = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_char_lowercase))
	x.StringCharUppercase = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_char_uppercase))
	x.StringToLower = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_to_lower))
	x.StringToUpper = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_to_upper))
	x.StringGetBasename = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_get_basename))
	x.StringGetExtension = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_get_extension))
	x.StringLeft = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_left))
	x.StringOrdAt = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_ord_at))
	x.StringPlusFile = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_plus_file))
	x.StringRight = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_right))
	x.StringStripEdges = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_strip_edges))
	x.StringStripEscapes = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_strip_escapes))
	x.StringErase = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_erase))
	x.StringAscii = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_ascii))
	x.StringAsciiExtended = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_ascii_extended))
	x.StringUtf8 = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_utf8))
	x.StringParseUtf8 = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_parse_utf8))
	x.StringParseUtf8WithLen = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_parse_utf8_with_len))
	x.StringCharsToUtf8 = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_chars_to_utf8))
	x.StringCharsToUtf8WithLen = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_chars_to_utf8_with_len))
	x.StringHash = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_hash))
	x.StringHash64 = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_hash64))
	x.StringHashChars = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_hash_chars))
	x.StringHashCharsWithLen = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_hash_chars_with_len))
	x.StringHashUtf8Chars = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_hash_utf8_chars))
	x.StringHashUtf8CharsWithLen = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_hash_utf8_chars_with_len))
	x.StringMd5Buffer = NewPoolByteArrayRef(unsafe.Pointer(x.ref57717e51.godot_string_md5_buffer))
	x.StringMd5Text = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_md5_text))
	x.StringSha256Buffer = NewPoolByteArrayRef(unsafe.Pointer(x.ref57717e51.godot_string_sha256_buffer))
	x.StringSha256Text = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_sha256_text))
	x.StringEmpty = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_empty))
	x.StringGetBaseDir = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_get_base_dir))
	x.StringGetFile = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_get_file))
	x.StringHumanizeSize = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_humanize_size))
	x.StringIsAbsPath = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_is_abs_path))
	x.StringIsRelPath = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_is_rel_path))
	x.StringIsResourceFile = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_is_resource_file))
	x.StringPathTo = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_path_to))
	x.StringPathToFile = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_path_to_file))
	x.StringSimplifyPath = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_simplify_path))
	x.StringCEscape = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_c_escape))
	x.StringCEscapeMultiline = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_c_escape_multiline))
	x.StringCUnescape = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_c_unescape))
	x.StringHttpEscape = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_http_escape))
	x.StringHttpUnescape = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_http_unescape))
	x.StringJsonEscape = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_json_escape))
	x.StringWordWrap = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_word_wrap))
	x.StringXmlEscape = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_xml_escape))
	x.StringXmlEscapeWithQuotes = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_xml_escape_with_quotes))
	x.StringXmlUnescape = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_xml_unescape))
	x.StringPercentDecode = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_percent_decode))
	x.StringPercentEncode = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_percent_encode))
	x.StringIsValidFloat = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_is_valid_float))
	x.StringIsValidHexNumber = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_is_valid_hex_number))
	x.StringIsValidHtmlColor = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_is_valid_html_color))
	x.StringIsValidIdentifier = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_is_valid_identifier))
	x.StringIsValidInteger = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_is_valid_integer))
	x.StringIsValidIpAddress = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_is_valid_ip_address))
	x.StringDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_destroy))
	x.StringNameNew = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_name_new))
	x.StringNameNewData = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_name_new_data))
	x.StringNameGetName = NewStringRef(unsafe.Pointer(x.ref57717e51.godot_string_name_get_name))
	x.StringNameGetHash = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_name_get_hash))
	x.StringNameGetDataUniquePointer = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_name_get_data_unique_pointer))
	x.StringNameOperatorEqual = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_name_operator_equal))
	x.StringNameOperatorLess = NewBoolRef(unsafe.Pointer(x.ref57717e51.godot_string_name_operator_less))
	x.StringNameDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_string_name_destroy))
	x.ObjectDestroy = NewRef(unsafe.Pointer(x.ref57717e51.godot_object_destroy))
	x.GlobalGetSingleton = NewObjectRef(unsafe.Pointer(x.ref57717e51.godot_global_get_singleton))
	x.MethodBindGetMethod = NewMethodBindRef(unsafe.Pointer(x.ref57717e51.godot_method_bind_get_method))
	x.MethodBindPtrcall = NewRef(unsafe.Pointer(x.ref57717e51.godot_method_bind_ptrcall))
	x.MethodBindCall = NewVariantRef(unsafe.Pointer(x.ref57717e51.godot_method_bind_call))
	x.GetClassConstructor = *NewClassConstructorRef(unsafe.Pointer(&x.ref57717e51.godot_get_class_constructor))
	x.RegisterNativeCallType = NewRef(unsafe.Pointer(x.ref57717e51.godot_register_native_call_type))
	x.Alloc = NewRef(unsafe.Pointer(x.ref57717e51.godot_alloc))
	x.Realloc = NewRef(unsafe.Pointer(x.ref57717e51.godot_realloc))
	x.Free = NewRef(unsafe.Pointer(x.ref57717e51.godot_free))
	x.PrintError = NewRef(unsafe.Pointer(x.ref57717e51.godot_print_error))
	x.PrintWarning = NewRef(unsafe.Pointer(x.ref57717e51.godot_print_warning))
	x.Print = NewRef(unsafe.Pointer(x.ref57717e51.godot_print))
}

// allocMethodBindMemory allocates memory for type C.godot_method_bind in C.
// The caller is responsible for freeing the this memory via C.free.
func allocMethodBindMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfMethodBindValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfMethodBindValue = unsafe.Sizeof([1]C.godot_method_bind{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *MethodBind) Ref() *C.godot_method_bind {
	if x == nil {
		return nil
	}
	return x.ref3a05c0bc
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *MethodBind) Free() {
	if x != nil && x.allocs3a05c0bc != nil {
		x.allocs3a05c0bc.(*cgoAllocMap).Free()
		x.ref3a05c0bc = nil
	}
}

// NewMethodBindRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewMethodBindRef(ref unsafe.Pointer) *MethodBind {
	if ref == nil {
		return nil
	}
	obj := new(MethodBind)
	obj.ref3a05c0bc = (*C.godot_method_bind)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *MethodBind) PassRef() (*C.godot_method_bind, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref3a05c0bc != nil {
		return x.ref3a05c0bc, nil
	}
	mem3a05c0bc := allocMethodBindMemory(1)
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
func (x MethodBind) PassValue() (C.godot_method_bind, *cgoAllocMap) {
	if x.ref3a05c0bc != nil {
		return *x.ref3a05c0bc, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *MethodBind) Deref() {
	if x.ref3a05c0bc == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.ref3a05c0bc._dont_touch_that))
}

// allocGdnativeApiVersionMemory allocates memory for type C.godot_gdnative_api_version in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGdnativeApiVersionMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGdnativeApiVersionValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGdnativeApiVersionValue = unsafe.Sizeof([1]C.godot_gdnative_api_version{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GdnativeApiVersion) Ref() *C.godot_gdnative_api_version {
	if x == nil {
		return nil
	}
	return x.ref5eed2c27
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GdnativeApiVersion) Free() {
	if x != nil && x.allocs5eed2c27 != nil {
		x.allocs5eed2c27.(*cgoAllocMap).Free()
		x.ref5eed2c27 = nil
	}
}

// NewGdnativeApiVersionRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGdnativeApiVersionRef(ref unsafe.Pointer) *GdnativeApiVersion {
	if ref == nil {
		return nil
	}
	obj := new(GdnativeApiVersion)
	obj.ref5eed2c27 = (*C.godot_gdnative_api_version)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GdnativeApiVersion) PassRef() (*C.godot_gdnative_api_version, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref5eed2c27 != nil {
		return x.ref5eed2c27, nil
	}
	mem5eed2c27 := allocGdnativeApiVersionMemory(1)
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
func (x GdnativeApiVersion) PassValue() (C.godot_gdnative_api_version, *cgoAllocMap) {
	if x.ref5eed2c27 != nil {
		return *x.ref5eed2c27, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GdnativeApiVersion) Deref() {
	if x.ref5eed2c27 == nil {
		return
	}
	x.Major = (uint32)(x.ref5eed2c27.major)
	x.Minor = (uint32)(x.ref5eed2c27.minor)
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GdnativeApiStruct) Ref() *C.godot_gdnative_api_struct {
	if x == nil {
		return nil
	}
	return x.ref45f52b65
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GdnativeApiStruct) Free() {
	if x != nil && x.allocs45f52b65 != nil {
		x.allocs45f52b65.(*cgoAllocMap).Free()
		x.ref45f52b65 = nil
	}
}

// NewGdnativeApiStructRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGdnativeApiStructRef(ref unsafe.Pointer) *GdnativeApiStruct {
	if ref == nil {
		return nil
	}
	obj := new(GdnativeApiStruct)
	obj.ref45f52b65 = (*C.godot_gdnative_api_struct)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GdnativeApiStruct) PassRef() (*C.godot_gdnative_api_struct, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref45f52b65 != nil {
		return x.ref45f52b65, nil
	}
	mem45f52b65 := allocGdnativeApiStructMemory(1)
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
	ref45f52b65.next, cnext_allocs = unpackSGdnativeApiStruct(x.Next)
	allocs45f52b65.Borrow(cnext_allocs)

	x.ref45f52b65 = ref45f52b65
	x.allocs45f52b65 = allocs45f52b65
	return ref45f52b65, allocs45f52b65

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GdnativeApiStruct) PassValue() (C.godot_gdnative_api_struct, *cgoAllocMap) {
	if x.ref45f52b65 != nil {
		return *x.ref45f52b65, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GdnativeApiStruct) Deref() {
	if x.ref45f52b65 == nil {
		return
	}
	x.Type = (uint32)(x.ref45f52b65._type)
	x.Version = *NewGdnativeApiVersionRef(unsafe.Pointer(&x.ref45f52b65.version))
	packSGdnativeApiStruct(x.Next, x.ref45f52b65.next)
}

// allocGdnativeInitOptionsMemory allocates memory for type C.godot_gdnative_init_options in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGdnativeInitOptionsMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGdnativeInitOptionsValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGdnativeInitOptionsValue = unsafe.Sizeof([1]C.godot_gdnative_init_options{})

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

// unpackSGdnativeCoreApiStruct transforms a sliced Go data structure into plain C format.
func unpackSGdnativeCoreApiStruct(x []GdnativeCoreApiStruct) (unpacked *C.struct_godot_gdnative_core_api_struct, allocs *cgoAllocMap) {
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

// allocStringMemory allocates memory for type C.godot_string in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStringMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStringValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStringValue = unsafe.Sizeof([1]C.godot_string{})

// unpackSString transforms a sliced Go data structure into plain C format.
func unpackSString(x []String) (unpacked *C.godot_string, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_string) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocStringMemory(len0)
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

// packSGdnativeCoreApiStruct reads sliced Go data structure out from plain C format.
func packSGdnativeCoreApiStruct(v []GdnativeCoreApiStruct, ptr0 *C.struct_godot_gdnative_core_api_struct) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfStructGodotGdnativeCoreApiStructValue]C.struct_godot_gdnative_core_api_struct)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGdnativeCoreApiStructRef(unsafe.Pointer(&ptr1))
	}
}

// packSString reads sliced Go data structure out from plain C format.
func packSString(v []String, ptr0 *C.godot_string) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfStringValue]C.godot_string)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewStringRef(unsafe.Pointer(&ptr1))
	}
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GdnativeInitOptions) Ref() *C.godot_gdnative_init_options {
	if x == nil {
		return nil
	}
	return x.reff9d34929
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GdnativeInitOptions) Free() {
	if x != nil && x.allocsf9d34929 != nil {
		x.allocsf9d34929.(*cgoAllocMap).Free()
		x.reff9d34929 = nil
	}
}

// NewGdnativeInitOptionsRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGdnativeInitOptionsRef(ref unsafe.Pointer) *GdnativeInitOptions {
	if ref == nil {
		return nil
	}
	obj := new(GdnativeInitOptions)
	obj.reff9d34929 = (*C.godot_gdnative_init_options)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GdnativeInitOptions) PassRef() (*C.godot_gdnative_init_options, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.reff9d34929 != nil {
		return x.reff9d34929, nil
	}
	memf9d34929 := allocGdnativeInitOptionsMemory(1)
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
	reff9d34929.api_struct, capi_struct_allocs = unpackSGdnativeCoreApiStruct(x.ApiStruct)
	allocsf9d34929.Borrow(capi_struct_allocs)

	var cactive_library_path_allocs *cgoAllocMap
	reff9d34929.active_library_path, cactive_library_path_allocs = unpackSString(x.ActiveLibraryPath)
	allocsf9d34929.Borrow(cactive_library_path_allocs)

	x.reff9d34929 = reff9d34929
	x.allocsf9d34929 = allocsf9d34929
	return reff9d34929, allocsf9d34929

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x GdnativeInitOptions) PassValue() (C.godot_gdnative_init_options, *cgoAllocMap) {
	if x.reff9d34929 != nil {
		return *x.reff9d34929, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GdnativeInitOptions) Deref() {
	if x.reff9d34929 == nil {
		return
	}
	x.InEditor = (Bool)(x.reff9d34929.in_editor)
	x.CoreApiHash = (uint64)(x.reff9d34929.core_api_hash)
	x.EditorApiHash = (uint64)(x.reff9d34929.editor_api_hash)
	x.NoApiHash = (uint64)(x.reff9d34929.no_api_hash)
	x.ReportVersionMismatch = NewRef(unsafe.Pointer(x.reff9d34929.report_version_mismatch))
	x.ReportLoadingError = NewRef(unsafe.Pointer(x.reff9d34929.report_loading_error))
	x.GdNativeLibrary = (*Object)(unsafe.Pointer(x.reff9d34929.gd_native_library))
	packSGdnativeCoreApiStruct(x.ApiStruct, x.reff9d34929.api_struct)
	packSString(x.ActiveLibraryPath, x.reff9d34929.active_library_path)
}

// allocGdnativeTerminateOptionsMemory allocates memory for type C.godot_gdnative_terminate_options in C.
// The caller is responsible for freeing the this memory via C.free.
func allocGdnativeTerminateOptionsMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfGdnativeTerminateOptionsValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfGdnativeTerminateOptionsValue = unsafe.Sizeof([1]C.godot_gdnative_terminate_options{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *GdnativeTerminateOptions) Ref() *C.godot_gdnative_terminate_options {
	if x == nil {
		return nil
	}
	return x.ref80c63fdc
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *GdnativeTerminateOptions) Free() {
	if x != nil && x.allocs80c63fdc != nil {
		x.allocs80c63fdc.(*cgoAllocMap).Free()
		x.ref80c63fdc = nil
	}
}

// NewGdnativeTerminateOptionsRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewGdnativeTerminateOptionsRef(ref unsafe.Pointer) *GdnativeTerminateOptions {
	if ref == nil {
		return nil
	}
	obj := new(GdnativeTerminateOptions)
	obj.ref80c63fdc = (*C.godot_gdnative_terminate_options)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *GdnativeTerminateOptions) PassRef() (*C.godot_gdnative_terminate_options, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref80c63fdc != nil {
		return x.ref80c63fdc, nil
	}
	mem80c63fdc := allocGdnativeTerminateOptionsMemory(1)
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
func (x GdnativeTerminateOptions) PassValue() (C.godot_gdnative_terminate_options, *cgoAllocMap) {
	if x.ref80c63fdc != nil {
		return *x.ref80c63fdc, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *GdnativeTerminateOptions) Deref() {
	if x.ref80c63fdc == nil {
		return
	}
	x.InEditor = (Bool)(x.ref80c63fdc.in_editor)
}

func (x ClassConstructor) PassRef() (ref *C.godot_class_constructor, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if classConstructorC099ECDEFunc == nil {
		classConstructorC099ECDEFunc = x
	}
	return (*C.godot_class_constructor)(C.godot_class_constructor_c099ecde), nil
}

func (x ClassConstructor) PassValue() (ref C.godot_class_constructor, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if classConstructorC099ECDEFunc == nil {
		classConstructorC099ECDEFunc = x
	}
	return (C.godot_class_constructor)(C.godot_class_constructor_c099ecde), nil
}

func NewClassConstructorRef(ref unsafe.Pointer) *ClassConstructor {
	return (*ClassConstructor)(ref)
}

//export classConstructorC099ECDE
func classConstructorC099ECDE() C.godot_object {
	if classConstructorC099ECDEFunc != nil {
		retc099ecde := classConstructorC099ECDEFunc()
		ret, _ := (C.godot_object)(unsafe.Pointer(retc099ecde)), cgoAllocsUnknown
		return ret
	}
	panic("callback func has not been set (race?)")
}

var classConstructorC099ECDEFunc ClassConstructor

// packSGdnativeInitOptions reads sliced Go data structure out from plain C format.
func packSGdnativeInitOptions(v []GdnativeInitOptions, ptr0 *C.godot_gdnative_init_options) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGdnativeInitOptionsValue]C.godot_gdnative_init_options)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGdnativeInitOptionsRef(unsafe.Pointer(&ptr1))
	}
}

func (x GdnativeInitFn) PassRef() (ref *C.godot_gdnative_init_fn, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if gdnativeInitFnF0F9A077Func == nil {
		gdnativeInitFnF0F9A077Func = x
	}
	return (*C.godot_gdnative_init_fn)(C.godot_gdnative_init_fn_f0f9a077), nil
}

func (x GdnativeInitFn) PassValue() (ref C.godot_gdnative_init_fn, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if gdnativeInitFnF0F9A077Func == nil {
		gdnativeInitFnF0F9A077Func = x
	}
	return (C.godot_gdnative_init_fn)(C.godot_gdnative_init_fn_f0f9a077), nil
}

func NewGdnativeInitFnRef(ref unsafe.Pointer) *GdnativeInitFn {
	return (*GdnativeInitFn)(ref)
}

//export gdnativeInitFnF0F9A077
func gdnativeInitFnF0F9A077(carg0 *C.godot_gdnative_init_options) {
	if gdnativeInitFnF0F9A077Func != nil {
		var arg0f0f9a077 []GdnativeInitOptions
		packSGdnativeInitOptions(arg0f0f9a077, carg0)
		gdnativeInitFnF0F9A077Func(arg0f0f9a077)
		return
	}
	panic("callback func has not been set (race?)")
}

var gdnativeInitFnF0F9A077Func GdnativeInitFn

// packSGdnativeTerminateOptions reads sliced Go data structure out from plain C format.
func packSGdnativeTerminateOptions(v []GdnativeTerminateOptions, ptr0 *C.godot_gdnative_terminate_options) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfGdnativeTerminateOptionsValue]C.godot_gdnative_terminate_options)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewGdnativeTerminateOptionsRef(unsafe.Pointer(&ptr1))
	}
}

func (x GdnativeTerminateFn) PassRef() (ref *C.godot_gdnative_terminate_fn, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if gdnativeTerminateFn7F5590C3Func == nil {
		gdnativeTerminateFn7F5590C3Func = x
	}
	return (*C.godot_gdnative_terminate_fn)(C.godot_gdnative_terminate_fn_7f5590c3), nil
}

func (x GdnativeTerminateFn) PassValue() (ref C.godot_gdnative_terminate_fn, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if gdnativeTerminateFn7F5590C3Func == nil {
		gdnativeTerminateFn7F5590C3Func = x
	}
	return (C.godot_gdnative_terminate_fn)(C.godot_gdnative_terminate_fn_7f5590c3), nil
}

func NewGdnativeTerminateFnRef(ref unsafe.Pointer) *GdnativeTerminateFn {
	return (*GdnativeTerminateFn)(ref)
}

//export gdnativeTerminateFn7F5590C3
func gdnativeTerminateFn7F5590C3(carg0 *C.godot_gdnative_terminate_options) {
	if gdnativeTerminateFn7F5590C3Func != nil {
		var arg07f5590c3 []GdnativeTerminateOptions
		packSGdnativeTerminateOptions(arg07f5590c3, carg0)
		gdnativeTerminateFn7F5590C3Func(arg07f5590c3)
		return
	}
	panic("callback func has not been set (race?)")
}

var gdnativeTerminateFn7F5590C3Func GdnativeTerminateFn

// packSArray reads sliced Go data structure out from plain C format.
func packSArray(v []Array, ptr0 *C.godot_array) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfArrayValue]C.godot_array)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewArrayRef(unsafe.Pointer(&ptr1))
	}
}

func (x GdnativeProcedureFn) PassRef() (ref *C.godot_gdnative_procedure_fn, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if gdnativeProcedureFn5D45412Func == nil {
		gdnativeProcedureFn5D45412Func = x
	}
	return (*C.godot_gdnative_procedure_fn)(C.godot_gdnative_procedure_fn_5d45412), nil
}

func (x GdnativeProcedureFn) PassValue() (ref C.godot_gdnative_procedure_fn, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if gdnativeProcedureFn5D45412Func == nil {
		gdnativeProcedureFn5D45412Func = x
	}
	return (C.godot_gdnative_procedure_fn)(C.godot_gdnative_procedure_fn_5d45412), nil
}

func NewGdnativeProcedureFnRef(ref unsafe.Pointer) *GdnativeProcedureFn {
	return (*GdnativeProcedureFn)(ref)
}

//export gdnativeProcedureFn5D45412
func gdnativeProcedureFn5D45412(carg0 *C.godot_array) C.godot_variant {
	if gdnativeProcedureFn5D45412Func != nil {
		var arg05d45412 []Array
		packSArray(arg05d45412, carg0)
		ret5d45412 := gdnativeProcedureFn5D45412Func(arg05d45412)
		ret, _ := ret5d45412.PassValue()
		return ret
	}
	panic("callback func has not been set (race?)")
}

var gdnativeProcedureFn5D45412Func GdnativeProcedureFn

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *String) Ref() *C.godot_string {
	if x == nil {
		return nil
	}
	return x.ref6d1ede57
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *String) Free() {
	if x != nil && x.allocs6d1ede57 != nil {
		x.allocs6d1ede57.(*cgoAllocMap).Free()
		x.ref6d1ede57 = nil
	}
}

// NewStringRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewStringRef(ref unsafe.Pointer) *String {
	if ref == nil {
		return nil
	}
	obj := new(String)
	obj.ref6d1ede57 = (*C.godot_string)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *String) PassRef() (*C.godot_string, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref6d1ede57 != nil {
		return x.ref6d1ede57, nil
	}
	mem6d1ede57 := allocStringMemory(1)
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
func (x String) PassValue() (C.godot_string, *cgoAllocMap) {
	if x.ref6d1ede57 != nil {
		return *x.ref6d1ede57, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *String) Deref() {
	if x.ref6d1ede57 == nil {
		return
	}
	x.DontTouchThat = *(*[8]byte)(unsafe.Pointer(&x.ref6d1ede57._dont_touch_that))
}

// allocArrayMemory allocates memory for type C.godot_array in C.
// The caller is responsible for freeing the this memory via C.free.
func allocArrayMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfArrayValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfArrayValue = unsafe.Sizeof([1]C.godot_array{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *Array) Ref() *C.godot_array {
	if x == nil {
		return nil
	}
	return x.refb81158a5
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *Array) Free() {
	if x != nil && x.allocsb81158a5 != nil {
		x.allocsb81158a5.(*cgoAllocMap).Free()
		x.refb81158a5 = nil
	}
}

// NewArrayRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewArrayRef(ref unsafe.Pointer) *Array {
	if ref == nil {
		return nil
	}
	obj := new(Array)
	obj.refb81158a5 = (*C.godot_array)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *Array) PassRef() (*C.godot_array, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refb81158a5 != nil {
		return x.refb81158a5, nil
	}
	memb81158a5 := allocArrayMemory(1)
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
func (x Array) PassValue() (C.godot_array, *cgoAllocMap) {
	if x.refb81158a5 != nil {
		return *x.refb81158a5, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *Array) Deref() {
	if x.refb81158a5 == nil {
		return
	}
	x.DontTouchThat = *(*[8]byte)(unsafe.Pointer(&x.refb81158a5._dont_touch_that))
}

// allocPoolArrayReadAccessMemory allocates memory for type C.godot_pool_array_read_access in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPoolArrayReadAccessMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPoolArrayReadAccessValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPoolArrayReadAccessValue = unsafe.Sizeof([1]C.godot_pool_array_read_access{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PoolArrayReadAccess) Ref() *C.godot_pool_array_read_access {
	if x == nil {
		return nil
	}
	return x.ref172179be
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PoolArrayReadAccess) Free() {
	if x != nil && x.allocs172179be != nil {
		x.allocs172179be.(*cgoAllocMap).Free()
		x.ref172179be = nil
	}
}

// NewPoolArrayReadAccessRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPoolArrayReadAccessRef(ref unsafe.Pointer) *PoolArrayReadAccess {
	if ref == nil {
		return nil
	}
	obj := new(PoolArrayReadAccess)
	obj.ref172179be = (*C.godot_pool_array_read_access)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PoolArrayReadAccess) PassRef() (*C.godot_pool_array_read_access, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref172179be != nil {
		return x.ref172179be, nil
	}
	mem172179be := allocPoolArrayReadAccessMemory(1)
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
func (x PoolArrayReadAccess) PassValue() (C.godot_pool_array_read_access, *cgoAllocMap) {
	if x.ref172179be != nil {
		return *x.ref172179be, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PoolArrayReadAccess) Deref() {
	if x.ref172179be == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.ref172179be._dont_touch_that))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PoolByteArrayReadAccess) Ref() *C.godot_pool_array_read_access {
	if x == nil {
		return nil
	}
	return x.ref172179be
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PoolByteArrayReadAccess) Free() {
	if x != nil && x.allocs172179be != nil {
		x.allocs172179be.(*cgoAllocMap).Free()
		x.ref172179be = nil
	}
}

// NewPoolByteArrayReadAccessRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPoolByteArrayReadAccessRef(ref unsafe.Pointer) *PoolByteArrayReadAccess {
	if ref == nil {
		return nil
	}
	obj := new(PoolByteArrayReadAccess)
	obj.ref172179be = (*C.godot_pool_array_read_access)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PoolByteArrayReadAccess) PassRef() (*C.godot_pool_array_read_access, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref172179be != nil {
		return x.ref172179be, nil
	}
	mem172179be := allocPoolArrayReadAccessMemory(1)
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
func (x PoolByteArrayReadAccess) PassValue() (C.godot_pool_array_read_access, *cgoAllocMap) {
	if x.ref172179be != nil {
		return *x.ref172179be, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PoolByteArrayReadAccess) Deref() {
	if x.ref172179be == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.ref172179be._dont_touch_that))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PoolIntArrayReadAccess) Ref() *C.godot_pool_array_read_access {
	if x == nil {
		return nil
	}
	return x.ref172179be
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PoolIntArrayReadAccess) Free() {
	if x != nil && x.allocs172179be != nil {
		x.allocs172179be.(*cgoAllocMap).Free()
		x.ref172179be = nil
	}
}

// NewPoolIntArrayReadAccessRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPoolIntArrayReadAccessRef(ref unsafe.Pointer) *PoolIntArrayReadAccess {
	if ref == nil {
		return nil
	}
	obj := new(PoolIntArrayReadAccess)
	obj.ref172179be = (*C.godot_pool_array_read_access)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PoolIntArrayReadAccess) PassRef() (*C.godot_pool_array_read_access, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref172179be != nil {
		return x.ref172179be, nil
	}
	mem172179be := allocPoolArrayReadAccessMemory(1)
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
func (x PoolIntArrayReadAccess) PassValue() (C.godot_pool_array_read_access, *cgoAllocMap) {
	if x.ref172179be != nil {
		return *x.ref172179be, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PoolIntArrayReadAccess) Deref() {
	if x.ref172179be == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.ref172179be._dont_touch_that))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PoolRealArrayReadAccess) Ref() *C.godot_pool_array_read_access {
	if x == nil {
		return nil
	}
	return x.ref172179be
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PoolRealArrayReadAccess) Free() {
	if x != nil && x.allocs172179be != nil {
		x.allocs172179be.(*cgoAllocMap).Free()
		x.ref172179be = nil
	}
}

// NewPoolRealArrayReadAccessRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPoolRealArrayReadAccessRef(ref unsafe.Pointer) *PoolRealArrayReadAccess {
	if ref == nil {
		return nil
	}
	obj := new(PoolRealArrayReadAccess)
	obj.ref172179be = (*C.godot_pool_array_read_access)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PoolRealArrayReadAccess) PassRef() (*C.godot_pool_array_read_access, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref172179be != nil {
		return x.ref172179be, nil
	}
	mem172179be := allocPoolArrayReadAccessMemory(1)
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
func (x PoolRealArrayReadAccess) PassValue() (C.godot_pool_array_read_access, *cgoAllocMap) {
	if x.ref172179be != nil {
		return *x.ref172179be, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PoolRealArrayReadAccess) Deref() {
	if x.ref172179be == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.ref172179be._dont_touch_that))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PoolStringArrayReadAccess) Ref() *C.godot_pool_array_read_access {
	if x == nil {
		return nil
	}
	return x.ref172179be
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PoolStringArrayReadAccess) Free() {
	if x != nil && x.allocs172179be != nil {
		x.allocs172179be.(*cgoAllocMap).Free()
		x.ref172179be = nil
	}
}

// NewPoolStringArrayReadAccessRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPoolStringArrayReadAccessRef(ref unsafe.Pointer) *PoolStringArrayReadAccess {
	if ref == nil {
		return nil
	}
	obj := new(PoolStringArrayReadAccess)
	obj.ref172179be = (*C.godot_pool_array_read_access)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PoolStringArrayReadAccess) PassRef() (*C.godot_pool_array_read_access, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref172179be != nil {
		return x.ref172179be, nil
	}
	mem172179be := allocPoolArrayReadAccessMemory(1)
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
func (x PoolStringArrayReadAccess) PassValue() (C.godot_pool_array_read_access, *cgoAllocMap) {
	if x.ref172179be != nil {
		return *x.ref172179be, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PoolStringArrayReadAccess) Deref() {
	if x.ref172179be == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.ref172179be._dont_touch_that))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PoolVector2ArrayReadAccess) Ref() *C.godot_pool_array_read_access {
	if x == nil {
		return nil
	}
	return x.ref172179be
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PoolVector2ArrayReadAccess) Free() {
	if x != nil && x.allocs172179be != nil {
		x.allocs172179be.(*cgoAllocMap).Free()
		x.ref172179be = nil
	}
}

// NewPoolVector2ArrayReadAccessRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPoolVector2ArrayReadAccessRef(ref unsafe.Pointer) *PoolVector2ArrayReadAccess {
	if ref == nil {
		return nil
	}
	obj := new(PoolVector2ArrayReadAccess)
	obj.ref172179be = (*C.godot_pool_array_read_access)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PoolVector2ArrayReadAccess) PassRef() (*C.godot_pool_array_read_access, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref172179be != nil {
		return x.ref172179be, nil
	}
	mem172179be := allocPoolArrayReadAccessMemory(1)
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
func (x PoolVector2ArrayReadAccess) PassValue() (C.godot_pool_array_read_access, *cgoAllocMap) {
	if x.ref172179be != nil {
		return *x.ref172179be, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PoolVector2ArrayReadAccess) Deref() {
	if x.ref172179be == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.ref172179be._dont_touch_that))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PoolVector3ArrayReadAccess) Ref() *C.godot_pool_array_read_access {
	if x == nil {
		return nil
	}
	return x.ref172179be
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PoolVector3ArrayReadAccess) Free() {
	if x != nil && x.allocs172179be != nil {
		x.allocs172179be.(*cgoAllocMap).Free()
		x.ref172179be = nil
	}
}

// NewPoolVector3ArrayReadAccessRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPoolVector3ArrayReadAccessRef(ref unsafe.Pointer) *PoolVector3ArrayReadAccess {
	if ref == nil {
		return nil
	}
	obj := new(PoolVector3ArrayReadAccess)
	obj.ref172179be = (*C.godot_pool_array_read_access)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PoolVector3ArrayReadAccess) PassRef() (*C.godot_pool_array_read_access, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref172179be != nil {
		return x.ref172179be, nil
	}
	mem172179be := allocPoolArrayReadAccessMemory(1)
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
func (x PoolVector3ArrayReadAccess) PassValue() (C.godot_pool_array_read_access, *cgoAllocMap) {
	if x.ref172179be != nil {
		return *x.ref172179be, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PoolVector3ArrayReadAccess) Deref() {
	if x.ref172179be == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.ref172179be._dont_touch_that))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PoolColorArrayReadAccess) Ref() *C.godot_pool_array_read_access {
	if x == nil {
		return nil
	}
	return x.ref172179be
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PoolColorArrayReadAccess) Free() {
	if x != nil && x.allocs172179be != nil {
		x.allocs172179be.(*cgoAllocMap).Free()
		x.ref172179be = nil
	}
}

// NewPoolColorArrayReadAccessRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPoolColorArrayReadAccessRef(ref unsafe.Pointer) *PoolColorArrayReadAccess {
	if ref == nil {
		return nil
	}
	obj := new(PoolColorArrayReadAccess)
	obj.ref172179be = (*C.godot_pool_array_read_access)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PoolColorArrayReadAccess) PassRef() (*C.godot_pool_array_read_access, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref172179be != nil {
		return x.ref172179be, nil
	}
	mem172179be := allocPoolArrayReadAccessMemory(1)
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
func (x PoolColorArrayReadAccess) PassValue() (C.godot_pool_array_read_access, *cgoAllocMap) {
	if x.ref172179be != nil {
		return *x.ref172179be, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PoolColorArrayReadAccess) Deref() {
	if x.ref172179be == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.ref172179be._dont_touch_that))
}

// allocPoolArrayWriteAccessMemory allocates memory for type C.godot_pool_array_write_access in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPoolArrayWriteAccessMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPoolArrayWriteAccessValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPoolArrayWriteAccessValue = unsafe.Sizeof([1]C.godot_pool_array_write_access{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PoolArrayWriteAccess) Ref() *C.godot_pool_array_write_access {
	if x == nil {
		return nil
	}
	return x.refbe0648b
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PoolArrayWriteAccess) Free() {
	if x != nil && x.allocsbe0648b != nil {
		x.allocsbe0648b.(*cgoAllocMap).Free()
		x.refbe0648b = nil
	}
}

// NewPoolArrayWriteAccessRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPoolArrayWriteAccessRef(ref unsafe.Pointer) *PoolArrayWriteAccess {
	if ref == nil {
		return nil
	}
	obj := new(PoolArrayWriteAccess)
	obj.refbe0648b = (*C.godot_pool_array_write_access)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PoolArrayWriteAccess) PassRef() (*C.godot_pool_array_write_access, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refbe0648b != nil {
		return x.refbe0648b, nil
	}
	membe0648b := allocPoolArrayWriteAccessMemory(1)
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
func (x PoolArrayWriteAccess) PassValue() (C.godot_pool_array_write_access, *cgoAllocMap) {
	if x.refbe0648b != nil {
		return *x.refbe0648b, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PoolArrayWriteAccess) Deref() {
	if x.refbe0648b == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.refbe0648b._dont_touch_that))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PoolByteArrayWriteAccess) Ref() *C.godot_pool_array_write_access {
	if x == nil {
		return nil
	}
	return x.refbe0648b
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PoolByteArrayWriteAccess) Free() {
	if x != nil && x.allocsbe0648b != nil {
		x.allocsbe0648b.(*cgoAllocMap).Free()
		x.refbe0648b = nil
	}
}

// NewPoolByteArrayWriteAccessRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPoolByteArrayWriteAccessRef(ref unsafe.Pointer) *PoolByteArrayWriteAccess {
	if ref == nil {
		return nil
	}
	obj := new(PoolByteArrayWriteAccess)
	obj.refbe0648b = (*C.godot_pool_array_write_access)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PoolByteArrayWriteAccess) PassRef() (*C.godot_pool_array_write_access, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refbe0648b != nil {
		return x.refbe0648b, nil
	}
	membe0648b := allocPoolArrayWriteAccessMemory(1)
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
func (x PoolByteArrayWriteAccess) PassValue() (C.godot_pool_array_write_access, *cgoAllocMap) {
	if x.refbe0648b != nil {
		return *x.refbe0648b, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PoolByteArrayWriteAccess) Deref() {
	if x.refbe0648b == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.refbe0648b._dont_touch_that))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PoolIntArrayWriteAccess) Ref() *C.godot_pool_array_write_access {
	if x == nil {
		return nil
	}
	return x.refbe0648b
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PoolIntArrayWriteAccess) Free() {
	if x != nil && x.allocsbe0648b != nil {
		x.allocsbe0648b.(*cgoAllocMap).Free()
		x.refbe0648b = nil
	}
}

// NewPoolIntArrayWriteAccessRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPoolIntArrayWriteAccessRef(ref unsafe.Pointer) *PoolIntArrayWriteAccess {
	if ref == nil {
		return nil
	}
	obj := new(PoolIntArrayWriteAccess)
	obj.refbe0648b = (*C.godot_pool_array_write_access)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PoolIntArrayWriteAccess) PassRef() (*C.godot_pool_array_write_access, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refbe0648b != nil {
		return x.refbe0648b, nil
	}
	membe0648b := allocPoolArrayWriteAccessMemory(1)
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
func (x PoolIntArrayWriteAccess) PassValue() (C.godot_pool_array_write_access, *cgoAllocMap) {
	if x.refbe0648b != nil {
		return *x.refbe0648b, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PoolIntArrayWriteAccess) Deref() {
	if x.refbe0648b == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.refbe0648b._dont_touch_that))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PoolRealArrayWriteAccess) Ref() *C.godot_pool_array_write_access {
	if x == nil {
		return nil
	}
	return x.refbe0648b
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PoolRealArrayWriteAccess) Free() {
	if x != nil && x.allocsbe0648b != nil {
		x.allocsbe0648b.(*cgoAllocMap).Free()
		x.refbe0648b = nil
	}
}

// NewPoolRealArrayWriteAccessRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPoolRealArrayWriteAccessRef(ref unsafe.Pointer) *PoolRealArrayWriteAccess {
	if ref == nil {
		return nil
	}
	obj := new(PoolRealArrayWriteAccess)
	obj.refbe0648b = (*C.godot_pool_array_write_access)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PoolRealArrayWriteAccess) PassRef() (*C.godot_pool_array_write_access, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refbe0648b != nil {
		return x.refbe0648b, nil
	}
	membe0648b := allocPoolArrayWriteAccessMemory(1)
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
func (x PoolRealArrayWriteAccess) PassValue() (C.godot_pool_array_write_access, *cgoAllocMap) {
	if x.refbe0648b != nil {
		return *x.refbe0648b, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PoolRealArrayWriteAccess) Deref() {
	if x.refbe0648b == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.refbe0648b._dont_touch_that))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PoolStringArrayWriteAccess) Ref() *C.godot_pool_array_write_access {
	if x == nil {
		return nil
	}
	return x.refbe0648b
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PoolStringArrayWriteAccess) Free() {
	if x != nil && x.allocsbe0648b != nil {
		x.allocsbe0648b.(*cgoAllocMap).Free()
		x.refbe0648b = nil
	}
}

// NewPoolStringArrayWriteAccessRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPoolStringArrayWriteAccessRef(ref unsafe.Pointer) *PoolStringArrayWriteAccess {
	if ref == nil {
		return nil
	}
	obj := new(PoolStringArrayWriteAccess)
	obj.refbe0648b = (*C.godot_pool_array_write_access)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PoolStringArrayWriteAccess) PassRef() (*C.godot_pool_array_write_access, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refbe0648b != nil {
		return x.refbe0648b, nil
	}
	membe0648b := allocPoolArrayWriteAccessMemory(1)
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
func (x PoolStringArrayWriteAccess) PassValue() (C.godot_pool_array_write_access, *cgoAllocMap) {
	if x.refbe0648b != nil {
		return *x.refbe0648b, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PoolStringArrayWriteAccess) Deref() {
	if x.refbe0648b == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.refbe0648b._dont_touch_that))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PoolVector2ArrayWriteAccess) Ref() *C.godot_pool_array_write_access {
	if x == nil {
		return nil
	}
	return x.refbe0648b
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PoolVector2ArrayWriteAccess) Free() {
	if x != nil && x.allocsbe0648b != nil {
		x.allocsbe0648b.(*cgoAllocMap).Free()
		x.refbe0648b = nil
	}
}

// NewPoolVector2ArrayWriteAccessRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPoolVector2ArrayWriteAccessRef(ref unsafe.Pointer) *PoolVector2ArrayWriteAccess {
	if ref == nil {
		return nil
	}
	obj := new(PoolVector2ArrayWriteAccess)
	obj.refbe0648b = (*C.godot_pool_array_write_access)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PoolVector2ArrayWriteAccess) PassRef() (*C.godot_pool_array_write_access, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refbe0648b != nil {
		return x.refbe0648b, nil
	}
	membe0648b := allocPoolArrayWriteAccessMemory(1)
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
func (x PoolVector2ArrayWriteAccess) PassValue() (C.godot_pool_array_write_access, *cgoAllocMap) {
	if x.refbe0648b != nil {
		return *x.refbe0648b, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PoolVector2ArrayWriteAccess) Deref() {
	if x.refbe0648b == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.refbe0648b._dont_touch_that))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PoolVector3ArrayWriteAccess) Ref() *C.godot_pool_array_write_access {
	if x == nil {
		return nil
	}
	return x.refbe0648b
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PoolVector3ArrayWriteAccess) Free() {
	if x != nil && x.allocsbe0648b != nil {
		x.allocsbe0648b.(*cgoAllocMap).Free()
		x.refbe0648b = nil
	}
}

// NewPoolVector3ArrayWriteAccessRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPoolVector3ArrayWriteAccessRef(ref unsafe.Pointer) *PoolVector3ArrayWriteAccess {
	if ref == nil {
		return nil
	}
	obj := new(PoolVector3ArrayWriteAccess)
	obj.refbe0648b = (*C.godot_pool_array_write_access)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PoolVector3ArrayWriteAccess) PassRef() (*C.godot_pool_array_write_access, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refbe0648b != nil {
		return x.refbe0648b, nil
	}
	membe0648b := allocPoolArrayWriteAccessMemory(1)
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
func (x PoolVector3ArrayWriteAccess) PassValue() (C.godot_pool_array_write_access, *cgoAllocMap) {
	if x.refbe0648b != nil {
		return *x.refbe0648b, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PoolVector3ArrayWriteAccess) Deref() {
	if x.refbe0648b == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.refbe0648b._dont_touch_that))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PoolColorArrayWriteAccess) Ref() *C.godot_pool_array_write_access {
	if x == nil {
		return nil
	}
	return x.refbe0648b
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PoolColorArrayWriteAccess) Free() {
	if x != nil && x.allocsbe0648b != nil {
		x.allocsbe0648b.(*cgoAllocMap).Free()
		x.refbe0648b = nil
	}
}

// NewPoolColorArrayWriteAccessRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPoolColorArrayWriteAccessRef(ref unsafe.Pointer) *PoolColorArrayWriteAccess {
	if ref == nil {
		return nil
	}
	obj := new(PoolColorArrayWriteAccess)
	obj.refbe0648b = (*C.godot_pool_array_write_access)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PoolColorArrayWriteAccess) PassRef() (*C.godot_pool_array_write_access, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refbe0648b != nil {
		return x.refbe0648b, nil
	}
	membe0648b := allocPoolArrayWriteAccessMemory(1)
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
func (x PoolColorArrayWriteAccess) PassValue() (C.godot_pool_array_write_access, *cgoAllocMap) {
	if x.refbe0648b != nil {
		return *x.refbe0648b, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PoolColorArrayWriteAccess) Deref() {
	if x.refbe0648b == nil {
		return
	}
	x.DontTouchThat = *(*[1]byte)(unsafe.Pointer(&x.refbe0648b._dont_touch_that))
}

// allocPoolByteArrayMemory allocates memory for type C.godot_pool_byte_array in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPoolByteArrayMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPoolByteArrayValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPoolByteArrayValue = unsafe.Sizeof([1]C.godot_pool_byte_array{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PoolByteArray) Ref() *C.godot_pool_byte_array {
	if x == nil {
		return nil
	}
	return x.refbf60ed2
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PoolByteArray) Free() {
	if x != nil && x.allocsbf60ed2 != nil {
		x.allocsbf60ed2.(*cgoAllocMap).Free()
		x.refbf60ed2 = nil
	}
}

// NewPoolByteArrayRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPoolByteArrayRef(ref unsafe.Pointer) *PoolByteArray {
	if ref == nil {
		return nil
	}
	obj := new(PoolByteArray)
	obj.refbf60ed2 = (*C.godot_pool_byte_array)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PoolByteArray) PassRef() (*C.godot_pool_byte_array, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refbf60ed2 != nil {
		return x.refbf60ed2, nil
	}
	membf60ed2 := allocPoolByteArrayMemory(1)
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
func (x PoolByteArray) PassValue() (C.godot_pool_byte_array, *cgoAllocMap) {
	if x.refbf60ed2 != nil {
		return *x.refbf60ed2, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PoolByteArray) Deref() {
	if x.refbf60ed2 == nil {
		return
	}
	x.DontTouchThat = *(*[8]byte)(unsafe.Pointer(&x.refbf60ed2._dont_touch_that))
}

// allocPoolIntArrayMemory allocates memory for type C.godot_pool_int_array in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPoolIntArrayMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPoolIntArrayValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPoolIntArrayValue = unsafe.Sizeof([1]C.godot_pool_int_array{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PoolIntArray) Ref() *C.godot_pool_int_array {
	if x == nil {
		return nil
	}
	return x.ref5d4f26e6
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PoolIntArray) Free() {
	if x != nil && x.allocs5d4f26e6 != nil {
		x.allocs5d4f26e6.(*cgoAllocMap).Free()
		x.ref5d4f26e6 = nil
	}
}

// NewPoolIntArrayRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPoolIntArrayRef(ref unsafe.Pointer) *PoolIntArray {
	if ref == nil {
		return nil
	}
	obj := new(PoolIntArray)
	obj.ref5d4f26e6 = (*C.godot_pool_int_array)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PoolIntArray) PassRef() (*C.godot_pool_int_array, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref5d4f26e6 != nil {
		return x.ref5d4f26e6, nil
	}
	mem5d4f26e6 := allocPoolIntArrayMemory(1)
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
func (x PoolIntArray) PassValue() (C.godot_pool_int_array, *cgoAllocMap) {
	if x.ref5d4f26e6 != nil {
		return *x.ref5d4f26e6, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PoolIntArray) Deref() {
	if x.ref5d4f26e6 == nil {
		return
	}
	x.DontTouchThat = *(*[8]byte)(unsafe.Pointer(&x.ref5d4f26e6._dont_touch_that))
}

// allocPoolRealArrayMemory allocates memory for type C.godot_pool_real_array in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPoolRealArrayMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPoolRealArrayValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPoolRealArrayValue = unsafe.Sizeof([1]C.godot_pool_real_array{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PoolRealArray) Ref() *C.godot_pool_real_array {
	if x == nil {
		return nil
	}
	return x.refc76f44c3
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PoolRealArray) Free() {
	if x != nil && x.allocsc76f44c3 != nil {
		x.allocsc76f44c3.(*cgoAllocMap).Free()
		x.refc76f44c3 = nil
	}
}

// NewPoolRealArrayRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPoolRealArrayRef(ref unsafe.Pointer) *PoolRealArray {
	if ref == nil {
		return nil
	}
	obj := new(PoolRealArray)
	obj.refc76f44c3 = (*C.godot_pool_real_array)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PoolRealArray) PassRef() (*C.godot_pool_real_array, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refc76f44c3 != nil {
		return x.refc76f44c3, nil
	}
	memc76f44c3 := allocPoolRealArrayMemory(1)
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
func (x PoolRealArray) PassValue() (C.godot_pool_real_array, *cgoAllocMap) {
	if x.refc76f44c3 != nil {
		return *x.refc76f44c3, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PoolRealArray) Deref() {
	if x.refc76f44c3 == nil {
		return
	}
	x.DontTouchThat = *(*[8]byte)(unsafe.Pointer(&x.refc76f44c3._dont_touch_that))
}

// allocPoolStringArrayMemory allocates memory for type C.godot_pool_string_array in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPoolStringArrayMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPoolStringArrayValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPoolStringArrayValue = unsafe.Sizeof([1]C.godot_pool_string_array{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PoolStringArray) Ref() *C.godot_pool_string_array {
	if x == nil {
		return nil
	}
	return x.reff6fe5d9
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PoolStringArray) Free() {
	if x != nil && x.allocsf6fe5d9 != nil {
		x.allocsf6fe5d9.(*cgoAllocMap).Free()
		x.reff6fe5d9 = nil
	}
}

// NewPoolStringArrayRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPoolStringArrayRef(ref unsafe.Pointer) *PoolStringArray {
	if ref == nil {
		return nil
	}
	obj := new(PoolStringArray)
	obj.reff6fe5d9 = (*C.godot_pool_string_array)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PoolStringArray) PassRef() (*C.godot_pool_string_array, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.reff6fe5d9 != nil {
		return x.reff6fe5d9, nil
	}
	memf6fe5d9 := allocPoolStringArrayMemory(1)
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
func (x PoolStringArray) PassValue() (C.godot_pool_string_array, *cgoAllocMap) {
	if x.reff6fe5d9 != nil {
		return *x.reff6fe5d9, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PoolStringArray) Deref() {
	if x.reff6fe5d9 == nil {
		return
	}
	x.DontTouchThat = *(*[8]byte)(unsafe.Pointer(&x.reff6fe5d9._dont_touch_that))
}

// allocPoolVector2ArrayMemory allocates memory for type C.godot_pool_vector2_array in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPoolVector2ArrayMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPoolVector2ArrayValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPoolVector2ArrayValue = unsafe.Sizeof([1]C.godot_pool_vector2_array{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PoolVector2Array) Ref() *C.godot_pool_vector2_array {
	if x == nil {
		return nil
	}
	return x.ref7f6b2885
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PoolVector2Array) Free() {
	if x != nil && x.allocs7f6b2885 != nil {
		x.allocs7f6b2885.(*cgoAllocMap).Free()
		x.ref7f6b2885 = nil
	}
}

// NewPoolVector2ArrayRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPoolVector2ArrayRef(ref unsafe.Pointer) *PoolVector2Array {
	if ref == nil {
		return nil
	}
	obj := new(PoolVector2Array)
	obj.ref7f6b2885 = (*C.godot_pool_vector2_array)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PoolVector2Array) PassRef() (*C.godot_pool_vector2_array, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref7f6b2885 != nil {
		return x.ref7f6b2885, nil
	}
	mem7f6b2885 := allocPoolVector2ArrayMemory(1)
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
func (x PoolVector2Array) PassValue() (C.godot_pool_vector2_array, *cgoAllocMap) {
	if x.ref7f6b2885 != nil {
		return *x.ref7f6b2885, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PoolVector2Array) Deref() {
	if x.ref7f6b2885 == nil {
		return
	}
	x.DontTouchThat = *(*[8]byte)(unsafe.Pointer(&x.ref7f6b2885._dont_touch_that))
}

// allocPoolVector3ArrayMemory allocates memory for type C.godot_pool_vector3_array in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPoolVector3ArrayMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPoolVector3ArrayValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPoolVector3ArrayValue = unsafe.Sizeof([1]C.godot_pool_vector3_array{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PoolVector3Array) Ref() *C.godot_pool_vector3_array {
	if x == nil {
		return nil
	}
	return x.refd91c2331
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PoolVector3Array) Free() {
	if x != nil && x.allocsd91c2331 != nil {
		x.allocsd91c2331.(*cgoAllocMap).Free()
		x.refd91c2331 = nil
	}
}

// NewPoolVector3ArrayRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPoolVector3ArrayRef(ref unsafe.Pointer) *PoolVector3Array {
	if ref == nil {
		return nil
	}
	obj := new(PoolVector3Array)
	obj.refd91c2331 = (*C.godot_pool_vector3_array)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PoolVector3Array) PassRef() (*C.godot_pool_vector3_array, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refd91c2331 != nil {
		return x.refd91c2331, nil
	}
	memd91c2331 := allocPoolVector3ArrayMemory(1)
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
func (x PoolVector3Array) PassValue() (C.godot_pool_vector3_array, *cgoAllocMap) {
	if x.refd91c2331 != nil {
		return *x.refd91c2331, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PoolVector3Array) Deref() {
	if x.refd91c2331 == nil {
		return
	}
	x.DontTouchThat = *(*[8]byte)(unsafe.Pointer(&x.refd91c2331._dont_touch_that))
}

// allocPoolColorArrayMemory allocates memory for type C.godot_pool_color_array in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPoolColorArrayMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPoolColorArrayValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPoolColorArrayValue = unsafe.Sizeof([1]C.godot_pool_color_array{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PoolColorArray) Ref() *C.godot_pool_color_array {
	if x == nil {
		return nil
	}
	return x.refd5cae78e
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PoolColorArray) Free() {
	if x != nil && x.allocsd5cae78e != nil {
		x.allocsd5cae78e.(*cgoAllocMap).Free()
		x.refd5cae78e = nil
	}
}

// NewPoolColorArrayRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPoolColorArrayRef(ref unsafe.Pointer) *PoolColorArray {
	if ref == nil {
		return nil
	}
	obj := new(PoolColorArray)
	obj.refd5cae78e = (*C.godot_pool_color_array)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PoolColorArray) PassRef() (*C.godot_pool_color_array, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refd5cae78e != nil {
		return x.refd5cae78e, nil
	}
	memd5cae78e := allocPoolColorArrayMemory(1)
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
func (x PoolColorArray) PassValue() (C.godot_pool_color_array, *cgoAllocMap) {
	if x.refd5cae78e != nil {
		return *x.refd5cae78e, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PoolColorArray) Deref() {
	if x.refd5cae78e == nil {
		return
	}
	x.DontTouchThat = *(*[8]byte)(unsafe.Pointer(&x.refd5cae78e._dont_touch_that))
}

// allocColorMemory allocates memory for type C.godot_color in C.
// The caller is responsible for freeing the this memory via C.free.
func allocColorMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfColorValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfColorValue = unsafe.Sizeof([1]C.godot_color{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *Color) Ref() *C.godot_color {
	if x == nil {
		return nil
	}
	return x.ref7f4bfefb
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *Color) Free() {
	if x != nil && x.allocs7f4bfefb != nil {
		x.allocs7f4bfefb.(*cgoAllocMap).Free()
		x.ref7f4bfefb = nil
	}
}

// NewColorRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewColorRef(ref unsafe.Pointer) *Color {
	if ref == nil {
		return nil
	}
	obj := new(Color)
	obj.ref7f4bfefb = (*C.godot_color)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *Color) PassRef() (*C.godot_color, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref7f4bfefb != nil {
		return x.ref7f4bfefb, nil
	}
	mem7f4bfefb := allocColorMemory(1)
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
func (x Color) PassValue() (C.godot_color, *cgoAllocMap) {
	if x.ref7f4bfefb != nil {
		return *x.ref7f4bfefb, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *Color) Deref() {
	if x.ref7f4bfefb == nil {
		return
	}
	x.DontTouchThat = *(*[16]byte)(unsafe.Pointer(&x.ref7f4bfefb._dont_touch_that))
}

// allocVector2Memory allocates memory for type C.godot_vector2 in C.
// The caller is responsible for freeing the this memory via C.free.
func allocVector2Memory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfVector2Value))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfVector2Value = unsafe.Sizeof([1]C.godot_vector2{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *Vector2) Ref() *C.godot_vector2 {
	if x == nil {
		return nil
	}
	return x.refbc81274e
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *Vector2) Free() {
	if x != nil && x.allocsbc81274e != nil {
		x.allocsbc81274e.(*cgoAllocMap).Free()
		x.refbc81274e = nil
	}
}

// NewVector2Ref creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewVector2Ref(ref unsafe.Pointer) *Vector2 {
	if ref == nil {
		return nil
	}
	obj := new(Vector2)
	obj.refbc81274e = (*C.godot_vector2)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *Vector2) PassRef() (*C.godot_vector2, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refbc81274e != nil {
		return x.refbc81274e, nil
	}
	membc81274e := allocVector2Memory(1)
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
func (x Vector2) PassValue() (C.godot_vector2, *cgoAllocMap) {
	if x.refbc81274e != nil {
		return *x.refbc81274e, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *Vector2) Deref() {
	if x.refbc81274e == nil {
		return
	}
	x.DontTouchThat = *(*[8]byte)(unsafe.Pointer(&x.refbc81274e._dont_touch_that))
}

// allocVector3Memory allocates memory for type C.godot_vector3 in C.
// The caller is responsible for freeing the this memory via C.free.
func allocVector3Memory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfVector3Value))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfVector3Value = unsafe.Sizeof([1]C.godot_vector3{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *Vector3) Ref() *C.godot_vector3 {
	if x == nil {
		return nil
	}
	return x.refcb8617d8
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *Vector3) Free() {
	if x != nil && x.allocscb8617d8 != nil {
		x.allocscb8617d8.(*cgoAllocMap).Free()
		x.refcb8617d8 = nil
	}
}

// NewVector3Ref creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewVector3Ref(ref unsafe.Pointer) *Vector3 {
	if ref == nil {
		return nil
	}
	obj := new(Vector3)
	obj.refcb8617d8 = (*C.godot_vector3)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *Vector3) PassRef() (*C.godot_vector3, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refcb8617d8 != nil {
		return x.refcb8617d8, nil
	}
	memcb8617d8 := allocVector3Memory(1)
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
func (x Vector3) PassValue() (C.godot_vector3, *cgoAllocMap) {
	if x.refcb8617d8 != nil {
		return *x.refcb8617d8, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *Vector3) Deref() {
	if x.refcb8617d8 == nil {
		return
	}
	x.DontTouchThat = *(*[12]byte)(unsafe.Pointer(&x.refcb8617d8._dont_touch_that))
}

// allocBasisMemory allocates memory for type C.godot_basis in C.
// The caller is responsible for freeing the this memory via C.free.
func allocBasisMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfBasisValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfBasisValue = unsafe.Sizeof([1]C.godot_basis{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *Basis) Ref() *C.godot_basis {
	if x == nil {
		return nil
	}
	return x.ref94d3d325
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *Basis) Free() {
	if x != nil && x.allocs94d3d325 != nil {
		x.allocs94d3d325.(*cgoAllocMap).Free()
		x.ref94d3d325 = nil
	}
}

// NewBasisRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewBasisRef(ref unsafe.Pointer) *Basis {
	if ref == nil {
		return nil
	}
	obj := new(Basis)
	obj.ref94d3d325 = (*C.godot_basis)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *Basis) PassRef() (*C.godot_basis, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref94d3d325 != nil {
		return x.ref94d3d325, nil
	}
	mem94d3d325 := allocBasisMemory(1)
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
func (x Basis) PassValue() (C.godot_basis, *cgoAllocMap) {
	if x.ref94d3d325 != nil {
		return *x.ref94d3d325, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *Basis) Deref() {
	if x.ref94d3d325 == nil {
		return
	}
	x.DontTouchThat = *(*[36]byte)(unsafe.Pointer(&x.ref94d3d325._dont_touch_that))
}

// allocQuatMemory allocates memory for type C.godot_quat in C.
// The caller is responsible for freeing the this memory via C.free.
func allocQuatMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfQuatValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfQuatValue = unsafe.Sizeof([1]C.godot_quat{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *Quat) Ref() *C.godot_quat {
	if x == nil {
		return nil
	}
	return x.reffaf33e0b
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *Quat) Free() {
	if x != nil && x.allocsfaf33e0b != nil {
		x.allocsfaf33e0b.(*cgoAllocMap).Free()
		x.reffaf33e0b = nil
	}
}

// NewQuatRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewQuatRef(ref unsafe.Pointer) *Quat {
	if ref == nil {
		return nil
	}
	obj := new(Quat)
	obj.reffaf33e0b = (*C.godot_quat)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *Quat) PassRef() (*C.godot_quat, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.reffaf33e0b != nil {
		return x.reffaf33e0b, nil
	}
	memfaf33e0b := allocQuatMemory(1)
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
func (x Quat) PassValue() (C.godot_quat, *cgoAllocMap) {
	if x.reffaf33e0b != nil {
		return *x.reffaf33e0b, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *Quat) Deref() {
	if x.reffaf33e0b == nil {
		return
	}
	x.DontTouchThat = *(*[16]byte)(unsafe.Pointer(&x.reffaf33e0b._dont_touch_that))
}

// allocVariantMemory allocates memory for type C.godot_variant in C.
// The caller is responsible for freeing the this memory via C.free.
func allocVariantMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfVariantValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfVariantValue = unsafe.Sizeof([1]C.godot_variant{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *Variant) Ref() *C.godot_variant {
	if x == nil {
		return nil
	}
	return x.refabb5c0da
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *Variant) Free() {
	if x != nil && x.allocsabb5c0da != nil {
		x.allocsabb5c0da.(*cgoAllocMap).Free()
		x.refabb5c0da = nil
	}
}

// NewVariantRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewVariantRef(ref unsafe.Pointer) *Variant {
	if ref == nil {
		return nil
	}
	obj := new(Variant)
	obj.refabb5c0da = (*C.godot_variant)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *Variant) PassRef() (*C.godot_variant, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refabb5c0da != nil {
		return x.refabb5c0da, nil
	}
	memabb5c0da := allocVariantMemory(1)
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
func (x Variant) PassValue() (C.godot_variant, *cgoAllocMap) {
	if x.refabb5c0da != nil {
		return *x.refabb5c0da, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *Variant) Deref() {
	if x.refabb5c0da == nil {
		return
	}
	x.DontTouchThat = *(*[24]byte)(unsafe.Pointer(&x.refabb5c0da._dont_touch_that))
}

// allocVariantCallErrorMemory allocates memory for type C.godot_variant_call_error in C.
// The caller is responsible for freeing the this memory via C.free.
func allocVariantCallErrorMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfVariantCallErrorValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfVariantCallErrorValue = unsafe.Sizeof([1]C.godot_variant_call_error{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *VariantCallError) Ref() *C.godot_variant_call_error {
	if x == nil {
		return nil
	}
	return x.ref3ce71027
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *VariantCallError) Free() {
	if x != nil && x.allocs3ce71027 != nil {
		x.allocs3ce71027.(*cgoAllocMap).Free()
		x.ref3ce71027 = nil
	}
}

// NewVariantCallErrorRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewVariantCallErrorRef(ref unsafe.Pointer) *VariantCallError {
	if ref == nil {
		return nil
	}
	obj := new(VariantCallError)
	obj.ref3ce71027 = (*C.godot_variant_call_error)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *VariantCallError) PassRef() (*C.godot_variant_call_error, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref3ce71027 != nil {
		return x.ref3ce71027, nil
	}
	mem3ce71027 := allocVariantCallErrorMemory(1)
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
func (x VariantCallError) PassValue() (C.godot_variant_call_error, *cgoAllocMap) {
	if x.ref3ce71027 != nil {
		return *x.ref3ce71027, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *VariantCallError) Deref() {
	if x.ref3ce71027 == nil {
		return
	}
	x.Error = (VariantCallErrorError)(x.ref3ce71027.error)
	x.Argument = (int32)(x.ref3ce71027.argument)
	x.Expected = (VariantType)(x.ref3ce71027.expected)
}

// allocAabbMemory allocates memory for type C.godot_aabb in C.
// The caller is responsible for freeing the this memory via C.free.
func allocAabbMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfAabbValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfAabbValue = unsafe.Sizeof([1]C.godot_aabb{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *Aabb) Ref() *C.godot_aabb {
	if x == nil {
		return nil
	}
	return x.ref6e3c84aa
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *Aabb) Free() {
	if x != nil && x.allocs6e3c84aa != nil {
		x.allocs6e3c84aa.(*cgoAllocMap).Free()
		x.ref6e3c84aa = nil
	}
}

// NewAabbRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewAabbRef(ref unsafe.Pointer) *Aabb {
	if ref == nil {
		return nil
	}
	obj := new(Aabb)
	obj.ref6e3c84aa = (*C.godot_aabb)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *Aabb) PassRef() (*C.godot_aabb, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref6e3c84aa != nil {
		return x.ref6e3c84aa, nil
	}
	mem6e3c84aa := allocAabbMemory(1)
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
func (x Aabb) PassValue() (C.godot_aabb, *cgoAllocMap) {
	if x.ref6e3c84aa != nil {
		return *x.ref6e3c84aa, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *Aabb) Deref() {
	if x.ref6e3c84aa == nil {
		return
	}
	x.DontTouchThat = *(*[24]byte)(unsafe.Pointer(&x.ref6e3c84aa._dont_touch_that))
}

// allocPlaneMemory allocates memory for type C.godot_plane in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPlaneMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPlaneValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPlaneValue = unsafe.Sizeof([1]C.godot_plane{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *Plane) Ref() *C.godot_plane {
	if x == nil {
		return nil
	}
	return x.refd8ae9b92
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *Plane) Free() {
	if x != nil && x.allocsd8ae9b92 != nil {
		x.allocsd8ae9b92.(*cgoAllocMap).Free()
		x.refd8ae9b92 = nil
	}
}

// NewPlaneRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPlaneRef(ref unsafe.Pointer) *Plane {
	if ref == nil {
		return nil
	}
	obj := new(Plane)
	obj.refd8ae9b92 = (*C.godot_plane)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *Plane) PassRef() (*C.godot_plane, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refd8ae9b92 != nil {
		return x.refd8ae9b92, nil
	}
	memd8ae9b92 := allocPlaneMemory(1)
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
func (x Plane) PassValue() (C.godot_plane, *cgoAllocMap) {
	if x.refd8ae9b92 != nil {
		return *x.refd8ae9b92, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *Plane) Deref() {
	if x.refd8ae9b92 == nil {
		return
	}
	x.DontTouchThat = *(*[16]byte)(unsafe.Pointer(&x.refd8ae9b92._dont_touch_that))
}

// allocDictionaryMemory allocates memory for type C.godot_dictionary in C.
// The caller is responsible for freeing the this memory via C.free.
func allocDictionaryMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfDictionaryValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfDictionaryValue = unsafe.Sizeof([1]C.godot_dictionary{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *Dictionary) Ref() *C.godot_dictionary {
	if x == nil {
		return nil
	}
	return x.refb267a9b9
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *Dictionary) Free() {
	if x != nil && x.allocsb267a9b9 != nil {
		x.allocsb267a9b9.(*cgoAllocMap).Free()
		x.refb267a9b9 = nil
	}
}

// NewDictionaryRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewDictionaryRef(ref unsafe.Pointer) *Dictionary {
	if ref == nil {
		return nil
	}
	obj := new(Dictionary)
	obj.refb267a9b9 = (*C.godot_dictionary)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *Dictionary) PassRef() (*C.godot_dictionary, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refb267a9b9 != nil {
		return x.refb267a9b9, nil
	}
	memb267a9b9 := allocDictionaryMemory(1)
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
func (x Dictionary) PassValue() (C.godot_dictionary, *cgoAllocMap) {
	if x.refb267a9b9 != nil {
		return *x.refb267a9b9, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *Dictionary) Deref() {
	if x.refb267a9b9 == nil {
		return
	}
	x.DontTouchThat = *(*[8]byte)(unsafe.Pointer(&x.refb267a9b9._dont_touch_that))
}

// allocNodePathMemory allocates memory for type C.godot_node_path in C.
// The caller is responsible for freeing the this memory via C.free.
func allocNodePathMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfNodePathValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfNodePathValue = unsafe.Sizeof([1]C.godot_node_path{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *NodePath) Ref() *C.godot_node_path {
	if x == nil {
		return nil
	}
	return x.ref6c34dff3
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *NodePath) Free() {
	if x != nil && x.allocs6c34dff3 != nil {
		x.allocs6c34dff3.(*cgoAllocMap).Free()
		x.ref6c34dff3 = nil
	}
}

// NewNodePathRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewNodePathRef(ref unsafe.Pointer) *NodePath {
	if ref == nil {
		return nil
	}
	obj := new(NodePath)
	obj.ref6c34dff3 = (*C.godot_node_path)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *NodePath) PassRef() (*C.godot_node_path, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref6c34dff3 != nil {
		return x.ref6c34dff3, nil
	}
	mem6c34dff3 := allocNodePathMemory(1)
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
func (x NodePath) PassValue() (C.godot_node_path, *cgoAllocMap) {
	if x.ref6c34dff3 != nil {
		return *x.ref6c34dff3, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *NodePath) Deref() {
	if x.ref6c34dff3 == nil {
		return
	}
	x.DontTouchThat = *(*[8]byte)(unsafe.Pointer(&x.ref6c34dff3._dont_touch_that))
}

// allocRect2Memory allocates memory for type C.godot_rect2 in C.
// The caller is responsible for freeing the this memory via C.free.
func allocRect2Memory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfRect2Value))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfRect2Value = unsafe.Sizeof([1]C.godot_rect2{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *Rect2) Ref() *C.godot_rect2 {
	if x == nil {
		return nil
	}
	return x.ref99c06d9a
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *Rect2) Free() {
	if x != nil && x.allocs99c06d9a != nil {
		x.allocs99c06d9a.(*cgoAllocMap).Free()
		x.ref99c06d9a = nil
	}
}

// NewRect2Ref creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewRect2Ref(ref unsafe.Pointer) *Rect2 {
	if ref == nil {
		return nil
	}
	obj := new(Rect2)
	obj.ref99c06d9a = (*C.godot_rect2)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *Rect2) PassRef() (*C.godot_rect2, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref99c06d9a != nil {
		return x.ref99c06d9a, nil
	}
	mem99c06d9a := allocRect2Memory(1)
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
func (x Rect2) PassValue() (C.godot_rect2, *cgoAllocMap) {
	if x.ref99c06d9a != nil {
		return *x.ref99c06d9a, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *Rect2) Deref() {
	if x.ref99c06d9a == nil {
		return
	}
	x.DontTouchThat = *(*[16]byte)(unsafe.Pointer(&x.ref99c06d9a._dont_touch_that))
}

// allocRidMemory allocates memory for type C.godot_rid in C.
// The caller is responsible for freeing the this memory via C.free.
func allocRidMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfRidValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfRidValue = unsafe.Sizeof([1]C.godot_rid{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *Rid) Ref() *C.godot_rid {
	if x == nil {
		return nil
	}
	return x.ref67320fc7
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *Rid) Free() {
	if x != nil && x.allocs67320fc7 != nil {
		x.allocs67320fc7.(*cgoAllocMap).Free()
		x.ref67320fc7 = nil
	}
}

// NewRidRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewRidRef(ref unsafe.Pointer) *Rid {
	if ref == nil {
		return nil
	}
	obj := new(Rid)
	obj.ref67320fc7 = (*C.godot_rid)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *Rid) PassRef() (*C.godot_rid, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref67320fc7 != nil {
		return x.ref67320fc7, nil
	}
	mem67320fc7 := allocRidMemory(1)
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
func (x Rid) PassValue() (C.godot_rid, *cgoAllocMap) {
	if x.ref67320fc7 != nil {
		return *x.ref67320fc7, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *Rid) Deref() {
	if x.ref67320fc7 == nil {
		return
	}
	x.DontTouchThat = *(*[8]byte)(unsafe.Pointer(&x.ref67320fc7._dont_touch_that))
}

// allocTransformMemory allocates memory for type C.godot_transform in C.
// The caller is responsible for freeing the this memory via C.free.
func allocTransformMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfTransformValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfTransformValue = unsafe.Sizeof([1]C.godot_transform{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *Transform) Ref() *C.godot_transform {
	if x == nil {
		return nil
	}
	return x.refd77658c7
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *Transform) Free() {
	if x != nil && x.allocsd77658c7 != nil {
		x.allocsd77658c7.(*cgoAllocMap).Free()
		x.refd77658c7 = nil
	}
}

// NewTransformRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewTransformRef(ref unsafe.Pointer) *Transform {
	if ref == nil {
		return nil
	}
	obj := new(Transform)
	obj.refd77658c7 = (*C.godot_transform)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *Transform) PassRef() (*C.godot_transform, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refd77658c7 != nil {
		return x.refd77658c7, nil
	}
	memd77658c7 := allocTransformMemory(1)
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
func (x Transform) PassValue() (C.godot_transform, *cgoAllocMap) {
	if x.refd77658c7 != nil {
		return *x.refd77658c7, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *Transform) Deref() {
	if x.refd77658c7 == nil {
		return
	}
	x.DontTouchThat = *(*[48]byte)(unsafe.Pointer(&x.refd77658c7._dont_touch_that))
}

// allocTransform2dMemory allocates memory for type C.godot_transform2d in C.
// The caller is responsible for freeing the this memory via C.free.
func allocTransform2dMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfTransform2dValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfTransform2dValue = unsafe.Sizeof([1]C.godot_transform2d{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *Transform2d) Ref() *C.godot_transform2d {
	if x == nil {
		return nil
	}
	return x.ref77dacf6
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *Transform2d) Free() {
	if x != nil && x.allocs77dacf6 != nil {
		x.allocs77dacf6.(*cgoAllocMap).Free()
		x.ref77dacf6 = nil
	}
}

// NewTransform2dRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewTransform2dRef(ref unsafe.Pointer) *Transform2d {
	if ref == nil {
		return nil
	}
	obj := new(Transform2d)
	obj.ref77dacf6 = (*C.godot_transform2d)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *Transform2d) PassRef() (*C.godot_transform2d, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref77dacf6 != nil {
		return x.ref77dacf6, nil
	}
	mem77dacf6 := allocTransform2dMemory(1)
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
func (x Transform2d) PassValue() (C.godot_transform2d, *cgoAllocMap) {
	if x.ref77dacf6 != nil {
		return *x.ref77dacf6, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *Transform2d) Deref() {
	if x.ref77dacf6 == nil {
		return
	}
	x.DontTouchThat = *(*[24]byte)(unsafe.Pointer(&x.ref77dacf6._dont_touch_that))
}

// allocStringNameMemory allocates memory for type C.godot_string_name in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStringNameMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStringNameValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStringNameValue = unsafe.Sizeof([1]C.godot_string_name{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *StringName) Ref() *C.godot_string_name {
	if x == nil {
		return nil
	}
	return x.ref895548fc
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *StringName) Free() {
	if x != nil && x.allocs895548fc != nil {
		x.allocs895548fc.(*cgoAllocMap).Free()
		x.ref895548fc = nil
	}
}

// NewStringNameRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewStringNameRef(ref unsafe.Pointer) *StringName {
	if ref == nil {
		return nil
	}
	obj := new(StringName)
	obj.ref895548fc = (*C.godot_string_name)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *StringName) PassRef() (*C.godot_string_name, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref895548fc != nil {
		return x.ref895548fc, nil
	}
	mem895548fc := allocStringNameMemory(1)
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
func (x StringName) PassValue() (C.godot_string_name, *cgoAllocMap) {
	if x.ref895548fc != nil {
		return *x.ref895548fc, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *StringName) Deref() {
	if x.ref895548fc == nil {
		return
	}
	x.DontTouchThat = *(*[8]byte)(unsafe.Pointer(&x.ref895548fc._dont_touch_that))
}

// allocArvrInterfaceGdnativeMemory allocates memory for type C.godot_arvr_interface_gdnative in C.
// The caller is responsible for freeing the this memory via C.free.
func allocArvrInterfaceGdnativeMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfArvrInterfaceGdnativeValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfArvrInterfaceGdnativeValue = unsafe.Sizeof([1]C.godot_arvr_interface_gdnative{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *ArvrInterfaceGdnative) Ref() *C.godot_arvr_interface_gdnative {
	if x == nil {
		return nil
	}
	return x.reff96a0b88
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *ArvrInterfaceGdnative) Free() {
	if x != nil && x.allocsf96a0b88 != nil {
		x.allocsf96a0b88.(*cgoAllocMap).Free()
		x.reff96a0b88 = nil
	}
}

// NewArvrInterfaceGdnativeRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewArvrInterfaceGdnativeRef(ref unsafe.Pointer) *ArvrInterfaceGdnative {
	if ref == nil {
		return nil
	}
	obj := new(ArvrInterfaceGdnative)
	obj.reff96a0b88 = (*C.godot_arvr_interface_gdnative)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *ArvrInterfaceGdnative) PassRef() (*C.godot_arvr_interface_gdnative, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.reff96a0b88 != nil {
		return x.reff96a0b88, nil
	}
	memf96a0b88 := allocArvrInterfaceGdnativeMemory(1)
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
func (x ArvrInterfaceGdnative) PassValue() (C.godot_arvr_interface_gdnative, *cgoAllocMap) {
	if x.reff96a0b88 != nil {
		return *x.reff96a0b88, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *ArvrInterfaceGdnative) Deref() {
	if x.reff96a0b88 == nil {
		return
	}
	x.Constructor = NewRef(unsafe.Pointer(x.reff96a0b88.constructor))
	x.Destructor = NewRef(unsafe.Pointer(x.reff96a0b88.destructor))
	x.GetName = NewStringRef(unsafe.Pointer(x.reff96a0b88.get_name))
	x.GetCapabilities = NewIntRef(unsafe.Pointer(x.reff96a0b88.get_capabilities))
	x.GetAnchorDetectionIsEnabled = NewBoolRef(unsafe.Pointer(x.reff96a0b88.get_anchor_detection_is_enabled))
	x.SetAnchorDetectionIsEnabled = NewRef(unsafe.Pointer(x.reff96a0b88.set_anchor_detection_is_enabled))
	x.IsStereo = NewBoolRef(unsafe.Pointer(x.reff96a0b88.is_stereo))
	x.IsInitialized = NewBoolRef(unsafe.Pointer(x.reff96a0b88.is_initialized))
	x.Initialize = NewBoolRef(unsafe.Pointer(x.reff96a0b88.initialize))
	x.Uninitialize = NewRef(unsafe.Pointer(x.reff96a0b88.uninitialize))
	x.GetRenderTargetsize = NewVector2Ref(unsafe.Pointer(x.reff96a0b88.get_render_targetsize))
	x.GetTransformForEye = NewTransformRef(unsafe.Pointer(x.reff96a0b88.get_transform_for_eye))
	x.FillProjectionForEye = NewRef(unsafe.Pointer(x.reff96a0b88.fill_projection_for_eye))
	x.CommitForEye = NewRef(unsafe.Pointer(x.reff96a0b88.commit_for_eye))
	x.Process = NewRef(unsafe.Pointer(x.reff96a0b88.process))
}

// allocPropertyAttributesMemory allocates memory for type C.godot_property_attributes in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPropertyAttributesMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPropertyAttributesValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPropertyAttributesValue = unsafe.Sizeof([1]C.godot_property_attributes{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PropertyAttributes) Ref() *C.godot_property_attributes {
	if x == nil {
		return nil
	}
	return x.ref431c473b
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PropertyAttributes) Free() {
	if x != nil && x.allocs431c473b != nil {
		x.allocs431c473b.(*cgoAllocMap).Free()
		x.ref431c473b = nil
	}
}

// NewPropertyAttributesRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPropertyAttributesRef(ref unsafe.Pointer) *PropertyAttributes {
	if ref == nil {
		return nil
	}
	obj := new(PropertyAttributes)
	obj.ref431c473b = (*C.godot_property_attributes)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PropertyAttributes) PassRef() (*C.godot_property_attributes, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref431c473b != nil {
		return x.ref431c473b, nil
	}
	mem431c473b := allocPropertyAttributesMemory(1)
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
func (x PropertyAttributes) PassValue() (C.godot_property_attributes, *cgoAllocMap) {
	if x.ref431c473b != nil {
		return *x.ref431c473b, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PropertyAttributes) Deref() {
	if x.ref431c473b == nil {
		return
	}
	x.RsetType = (MethodRpcMode)(x.ref431c473b.rset_type)
	x.Type = (Int)(x.ref431c473b._type)
	x.Hint = (PropertyHint)(x.ref431c473b.hint)
	x.HintString = *NewStringRef(unsafe.Pointer(&x.ref431c473b.hint_string))
	x.Usage = (PropertyUsageFlags)(x.ref431c473b.usage)
	x.DefaultValue = *NewVariantRef(unsafe.Pointer(&x.ref431c473b.default_value))
}

// allocInstanceCreateFuncMemory allocates memory for type C.godot_instance_create_func in C.
// The caller is responsible for freeing the this memory via C.free.
func allocInstanceCreateFuncMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfInstanceCreateFuncValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfInstanceCreateFuncValue = unsafe.Sizeof([1]C.godot_instance_create_func{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *InstanceCreateFunc) Ref() *C.godot_instance_create_func {
	if x == nil {
		return nil
	}
	return x.ref70ecb5db
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *InstanceCreateFunc) Free() {
	if x != nil && x.allocs70ecb5db != nil {
		x.allocs70ecb5db.(*cgoAllocMap).Free()
		x.ref70ecb5db = nil
	}
}

// NewInstanceCreateFuncRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewInstanceCreateFuncRef(ref unsafe.Pointer) *InstanceCreateFunc {
	if ref == nil {
		return nil
	}
	obj := new(InstanceCreateFunc)
	obj.ref70ecb5db = (*C.godot_instance_create_func)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *InstanceCreateFunc) PassRef() (*C.godot_instance_create_func, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref70ecb5db != nil {
		return x.ref70ecb5db, nil
	}
	mem70ecb5db := allocInstanceCreateFuncMemory(1)
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
func (x InstanceCreateFunc) PassValue() (C.godot_instance_create_func, *cgoAllocMap) {
	if x.ref70ecb5db != nil {
		return *x.ref70ecb5db, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *InstanceCreateFunc) Deref() {
	if x.ref70ecb5db == nil {
		return
	}
	x.CreateFunc = NewRef(unsafe.Pointer(x.ref70ecb5db.create_func))
	x.MethodData = (unsafe.Pointer)(unsafe.Pointer(x.ref70ecb5db.method_data))
	x.FreeFunc = NewRef(unsafe.Pointer(x.ref70ecb5db.free_func))
}

// allocInstanceDestroyFuncMemory allocates memory for type C.godot_instance_destroy_func in C.
// The caller is responsible for freeing the this memory via C.free.
func allocInstanceDestroyFuncMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfInstanceDestroyFuncValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfInstanceDestroyFuncValue = unsafe.Sizeof([1]C.godot_instance_destroy_func{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *InstanceDestroyFunc) Ref() *C.godot_instance_destroy_func {
	if x == nil {
		return nil
	}
	return x.refd0d05668
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *InstanceDestroyFunc) Free() {
	if x != nil && x.allocsd0d05668 != nil {
		x.allocsd0d05668.(*cgoAllocMap).Free()
		x.refd0d05668 = nil
	}
}

// NewInstanceDestroyFuncRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewInstanceDestroyFuncRef(ref unsafe.Pointer) *InstanceDestroyFunc {
	if ref == nil {
		return nil
	}
	obj := new(InstanceDestroyFunc)
	obj.refd0d05668 = (*C.godot_instance_destroy_func)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *InstanceDestroyFunc) PassRef() (*C.godot_instance_destroy_func, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refd0d05668 != nil {
		return x.refd0d05668, nil
	}
	memd0d05668 := allocInstanceDestroyFuncMemory(1)
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
func (x InstanceDestroyFunc) PassValue() (C.godot_instance_destroy_func, *cgoAllocMap) {
	if x.refd0d05668 != nil {
		return *x.refd0d05668, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *InstanceDestroyFunc) Deref() {
	if x.refd0d05668 == nil {
		return
	}
	x.DestroyFunc = NewRef(unsafe.Pointer(x.refd0d05668.destroy_func))
	x.MethodData = (unsafe.Pointer)(unsafe.Pointer(x.refd0d05668.method_data))
	x.FreeFunc = NewRef(unsafe.Pointer(x.refd0d05668.free_func))
}

// allocMethodAttributesMemory allocates memory for type C.godot_method_attributes in C.
// The caller is responsible for freeing the this memory via C.free.
func allocMethodAttributesMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfMethodAttributesValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfMethodAttributesValue = unsafe.Sizeof([1]C.godot_method_attributes{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *MethodAttributes) Ref() *C.godot_method_attributes {
	if x == nil {
		return nil
	}
	return x.ref66a6c5c9
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *MethodAttributes) Free() {
	if x != nil && x.allocs66a6c5c9 != nil {
		x.allocs66a6c5c9.(*cgoAllocMap).Free()
		x.ref66a6c5c9 = nil
	}
}

// NewMethodAttributesRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewMethodAttributesRef(ref unsafe.Pointer) *MethodAttributes {
	if ref == nil {
		return nil
	}
	obj := new(MethodAttributes)
	obj.ref66a6c5c9 = (*C.godot_method_attributes)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *MethodAttributes) PassRef() (*C.godot_method_attributes, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref66a6c5c9 != nil {
		return x.ref66a6c5c9, nil
	}
	mem66a6c5c9 := allocMethodAttributesMemory(1)
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
func (x MethodAttributes) PassValue() (C.godot_method_attributes, *cgoAllocMap) {
	if x.ref66a6c5c9 != nil {
		return *x.ref66a6c5c9, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *MethodAttributes) Deref() {
	if x.ref66a6c5c9 == nil {
		return
	}
	x.RpcType = (MethodRpcMode)(x.ref66a6c5c9.rpc_type)
}

// allocInstanceMethodMemory allocates memory for type C.godot_instance_method in C.
// The caller is responsible for freeing the this memory via C.free.
func allocInstanceMethodMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfInstanceMethodValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfInstanceMethodValue = unsafe.Sizeof([1]C.godot_instance_method{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *InstanceMethod) Ref() *C.godot_instance_method {
	if x == nil {
		return nil
	}
	return x.ref10e1583e
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *InstanceMethod) Free() {
	if x != nil && x.allocs10e1583e != nil {
		x.allocs10e1583e.(*cgoAllocMap).Free()
		x.ref10e1583e = nil
	}
}

// NewInstanceMethodRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewInstanceMethodRef(ref unsafe.Pointer) *InstanceMethod {
	if ref == nil {
		return nil
	}
	obj := new(InstanceMethod)
	obj.ref10e1583e = (*C.godot_instance_method)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *InstanceMethod) PassRef() (*C.godot_instance_method, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref10e1583e != nil {
		return x.ref10e1583e, nil
	}
	mem10e1583e := allocInstanceMethodMemory(1)
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
func (x InstanceMethod) PassValue() (C.godot_instance_method, *cgoAllocMap) {
	if x.ref10e1583e != nil {
		return *x.ref10e1583e, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *InstanceMethod) Deref() {
	if x.ref10e1583e == nil {
		return
	}
	x.Method = NewVariantRef(unsafe.Pointer(x.ref10e1583e.method))
	x.MethodData = (unsafe.Pointer)(unsafe.Pointer(x.ref10e1583e.method_data))
	x.FreeFunc = NewRef(unsafe.Pointer(x.ref10e1583e.free_func))
}

// allocPropertySetFuncMemory allocates memory for type C.godot_property_set_func in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPropertySetFuncMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPropertySetFuncValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPropertySetFuncValue = unsafe.Sizeof([1]C.godot_property_set_func{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PropertySetFunc) Ref() *C.godot_property_set_func {
	if x == nil {
		return nil
	}
	return x.refc9844af
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PropertySetFunc) Free() {
	if x != nil && x.allocsc9844af != nil {
		x.allocsc9844af.(*cgoAllocMap).Free()
		x.refc9844af = nil
	}
}

// NewPropertySetFuncRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPropertySetFuncRef(ref unsafe.Pointer) *PropertySetFunc {
	if ref == nil {
		return nil
	}
	obj := new(PropertySetFunc)
	obj.refc9844af = (*C.godot_property_set_func)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PropertySetFunc) PassRef() (*C.godot_property_set_func, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refc9844af != nil {
		return x.refc9844af, nil
	}
	memc9844af := allocPropertySetFuncMemory(1)
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
func (x PropertySetFunc) PassValue() (C.godot_property_set_func, *cgoAllocMap) {
	if x.refc9844af != nil {
		return *x.refc9844af, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PropertySetFunc) Deref() {
	if x.refc9844af == nil {
		return
	}
	x.SetFunc = NewRef(unsafe.Pointer(x.refc9844af.set_func))
	x.MethodData = (unsafe.Pointer)(unsafe.Pointer(x.refc9844af.method_data))
	x.FreeFunc = NewRef(unsafe.Pointer(x.refc9844af.free_func))
}

// allocPropertyGetFuncMemory allocates memory for type C.godot_property_get_func in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPropertyGetFuncMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPropertyGetFuncValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPropertyGetFuncValue = unsafe.Sizeof([1]C.godot_property_get_func{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PropertyGetFunc) Ref() *C.godot_property_get_func {
	if x == nil {
		return nil
	}
	return x.reff4697b7e
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PropertyGetFunc) Free() {
	if x != nil && x.allocsf4697b7e != nil {
		x.allocsf4697b7e.(*cgoAllocMap).Free()
		x.reff4697b7e = nil
	}
}

// NewPropertyGetFuncRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPropertyGetFuncRef(ref unsafe.Pointer) *PropertyGetFunc {
	if ref == nil {
		return nil
	}
	obj := new(PropertyGetFunc)
	obj.reff4697b7e = (*C.godot_property_get_func)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PropertyGetFunc) PassRef() (*C.godot_property_get_func, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.reff4697b7e != nil {
		return x.reff4697b7e, nil
	}
	memf4697b7e := allocPropertyGetFuncMemory(1)
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
func (x PropertyGetFunc) PassValue() (C.godot_property_get_func, *cgoAllocMap) {
	if x.reff4697b7e != nil {
		return *x.reff4697b7e, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PropertyGetFunc) Deref() {
	if x.reff4697b7e == nil {
		return
	}
	x.GetFunc = NewVariantRef(unsafe.Pointer(x.reff4697b7e.get_func))
	x.MethodData = (unsafe.Pointer)(unsafe.Pointer(x.reff4697b7e.method_data))
	x.FreeFunc = NewRef(unsafe.Pointer(x.reff4697b7e.free_func))
}

// allocSignalArgumentMemory allocates memory for type C.godot_signal_argument in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSignalArgumentMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSignalArgumentValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSignalArgumentValue = unsafe.Sizeof([1]C.godot_signal_argument{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *SignalArgument) Ref() *C.godot_signal_argument {
	if x == nil {
		return nil
	}
	return x.refc21e72ac
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *SignalArgument) Free() {
	if x != nil && x.allocsc21e72ac != nil {
		x.allocsc21e72ac.(*cgoAllocMap).Free()
		x.refc21e72ac = nil
	}
}

// NewSignalArgumentRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewSignalArgumentRef(ref unsafe.Pointer) *SignalArgument {
	if ref == nil {
		return nil
	}
	obj := new(SignalArgument)
	obj.refc21e72ac = (*C.godot_signal_argument)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *SignalArgument) PassRef() (*C.godot_signal_argument, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refc21e72ac != nil {
		return x.refc21e72ac, nil
	}
	memc21e72ac := allocSignalArgumentMemory(1)
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
func (x SignalArgument) PassValue() (C.godot_signal_argument, *cgoAllocMap) {
	if x.refc21e72ac != nil {
		return *x.refc21e72ac, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *SignalArgument) Deref() {
	if x.refc21e72ac == nil {
		return
	}
	x.Name = *NewStringRef(unsafe.Pointer(&x.refc21e72ac.name))
	x.Type = (Int)(x.refc21e72ac._type)
	x.Hint = (PropertyHint)(x.refc21e72ac.hint)
	x.HintString = *NewStringRef(unsafe.Pointer(&x.refc21e72ac.hint_string))
	x.Usage = (PropertyUsageFlags)(x.refc21e72ac.usage)
	x.DefaultValue = *NewVariantRef(unsafe.Pointer(&x.refc21e72ac.default_value))
}

// allocSignalMemory allocates memory for type C.godot_signal in C.
// The caller is responsible for freeing the this memory via C.free.
func allocSignalMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfSignalValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfSignalValue = unsafe.Sizeof([1]C.godot_signal{})

// unpackSSignalArgument transforms a sliced Go data structure into plain C format.
func unpackSSignalArgument(x []SignalArgument) (unpacked *C.godot_signal_argument, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_signal_argument) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocSignalArgumentMemory(len0)
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

// unpackSVariant transforms a sliced Go data structure into plain C format.
func unpackSVariant(x []Variant) (unpacked *C.godot_variant, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_variant) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocVariantMemory(len0)
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

// packSSignalArgument reads sliced Go data structure out from plain C format.
func packSSignalArgument(v []SignalArgument, ptr0 *C.godot_signal_argument) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfSignalArgumentValue]C.godot_signal_argument)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewSignalArgumentRef(unsafe.Pointer(&ptr1))
	}
}

// packSVariant reads sliced Go data structure out from plain C format.
func packSVariant(v []Variant, ptr0 *C.godot_variant) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfVariantValue]C.godot_variant)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewVariantRef(unsafe.Pointer(&ptr1))
	}
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *Signal) Ref() *C.godot_signal {
	if x == nil {
		return nil
	}
	return x.ref87acf90b
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *Signal) Free() {
	if x != nil && x.allocs87acf90b != nil {
		x.allocs87acf90b.(*cgoAllocMap).Free()
		x.ref87acf90b = nil
	}
}

// NewSignalRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewSignalRef(ref unsafe.Pointer) *Signal {
	if ref == nil {
		return nil
	}
	obj := new(Signal)
	obj.ref87acf90b = (*C.godot_signal)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *Signal) PassRef() (*C.godot_signal, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref87acf90b != nil {
		return x.ref87acf90b, nil
	}
	mem87acf90b := allocSignalMemory(1)
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
	ref87acf90b.args, cargs_allocs = unpackSSignalArgument(x.Args)
	allocs87acf90b.Borrow(cargs_allocs)

	var cnum_default_args_allocs *cgoAllocMap
	ref87acf90b.num_default_args, cnum_default_args_allocs = (C.int)(x.NumDefaultArgs), cgoAllocsUnknown
	allocs87acf90b.Borrow(cnum_default_args_allocs)

	var cdefault_args_allocs *cgoAllocMap
	ref87acf90b.default_args, cdefault_args_allocs = unpackSVariant(x.DefaultArgs)
	allocs87acf90b.Borrow(cdefault_args_allocs)

	x.ref87acf90b = ref87acf90b
	x.allocs87acf90b = allocs87acf90b
	return ref87acf90b, allocs87acf90b

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x Signal) PassValue() (C.godot_signal, *cgoAllocMap) {
	if x.ref87acf90b != nil {
		return *x.ref87acf90b, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *Signal) Deref() {
	if x.ref87acf90b == nil {
		return
	}
	x.Name = *NewStringRef(unsafe.Pointer(&x.ref87acf90b.name))
	x.NumArgs = (int32)(x.ref87acf90b.num_args)
	packSSignalArgument(x.Args, x.ref87acf90b.args)
	x.NumDefaultArgs = (int32)(x.ref87acf90b.num_default_args)
	packSVariant(x.DefaultArgs, x.ref87acf90b.default_args)
}

// allocPluginscriptInstanceDescMemory allocates memory for type C.godot_pluginscript_instance_desc in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPluginscriptInstanceDescMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPluginscriptInstanceDescValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPluginscriptInstanceDescValue = unsafe.Sizeof([1]C.godot_pluginscript_instance_desc{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PluginscriptInstanceDesc) Ref() *C.godot_pluginscript_instance_desc {
	if x == nil {
		return nil
	}
	return x.refc0c19139
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PluginscriptInstanceDesc) Free() {
	if x != nil && x.allocsc0c19139 != nil {
		x.allocsc0c19139.(*cgoAllocMap).Free()
		x.refc0c19139 = nil
	}
}

// NewPluginscriptInstanceDescRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPluginscriptInstanceDescRef(ref unsafe.Pointer) *PluginscriptInstanceDesc {
	if ref == nil {
		return nil
	}
	obj := new(PluginscriptInstanceDesc)
	obj.refc0c19139 = (*C.godot_pluginscript_instance_desc)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PluginscriptInstanceDesc) PassRef() (*C.godot_pluginscript_instance_desc, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refc0c19139 != nil {
		return x.refc0c19139, nil
	}
	memc0c19139 := allocPluginscriptInstanceDescMemory(1)
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
func (x PluginscriptInstanceDesc) PassValue() (C.godot_pluginscript_instance_desc, *cgoAllocMap) {
	if x.refc0c19139 != nil {
		return *x.refc0c19139, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PluginscriptInstanceDesc) Deref() {
	if x.refc0c19139 == nil {
		return
	}
	x.Init = NewPluginscriptInstanceDataRef(unsafe.Pointer(x.refc0c19139.init))
	x.Finish = NewRef(unsafe.Pointer(x.refc0c19139.finish))
	x.SetProp = NewBoolRef(unsafe.Pointer(x.refc0c19139.set_prop))
	x.GetProp = NewBoolRef(unsafe.Pointer(x.refc0c19139.get_prop))
	x.CallMethod = NewVariantRef(unsafe.Pointer(x.refc0c19139.call_method))
	x.Notification = NewRef(unsafe.Pointer(x.refc0c19139.notification))
	x.GetRpcMode = NewMethodRpcModeRef(unsafe.Pointer(x.refc0c19139.get_rpc_mode))
	x.GetRsetMode = NewMethodRpcModeRef(unsafe.Pointer(x.refc0c19139.get_rset_mode))
	x.RefcountIncremented = NewRef(unsafe.Pointer(x.refc0c19139.refcount_incremented))
	x.RefcountDecremented = NewRef(unsafe.Pointer(x.refc0c19139.refcount_decremented))
}

// allocPluginscriptScriptManifestMemory allocates memory for type C.godot_pluginscript_script_manifest in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPluginscriptScriptManifestMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPluginscriptScriptManifestValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPluginscriptScriptManifestValue = unsafe.Sizeof([1]C.godot_pluginscript_script_manifest{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PluginscriptScriptManifest) Ref() *C.godot_pluginscript_script_manifest {
	if x == nil {
		return nil
	}
	return x.reffbf02dfd
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PluginscriptScriptManifest) Free() {
	if x != nil && x.allocsfbf02dfd != nil {
		x.allocsfbf02dfd.(*cgoAllocMap).Free()
		x.reffbf02dfd = nil
	}
}

// NewPluginscriptScriptManifestRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPluginscriptScriptManifestRef(ref unsafe.Pointer) *PluginscriptScriptManifest {
	if ref == nil {
		return nil
	}
	obj := new(PluginscriptScriptManifest)
	obj.reffbf02dfd = (*C.godot_pluginscript_script_manifest)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PluginscriptScriptManifest) PassRef() (*C.godot_pluginscript_script_manifest, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.reffbf02dfd != nil {
		return x.reffbf02dfd, nil
	}
	memfbf02dfd := allocPluginscriptScriptManifestMemory(1)
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
func (x PluginscriptScriptManifest) PassValue() (C.godot_pluginscript_script_manifest, *cgoAllocMap) {
	if x.reffbf02dfd != nil {
		return *x.reffbf02dfd, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PluginscriptScriptManifest) Deref() {
	if x.reffbf02dfd == nil {
		return
	}
	x.Data = (*PluginscriptScriptData)(unsafe.Pointer(x.reffbf02dfd.data))
	x.Name = *NewStringNameRef(unsafe.Pointer(&x.reffbf02dfd.name))
	x.IsTool = (Bool)(x.reffbf02dfd.is_tool)
	x.Base = *NewStringNameRef(unsafe.Pointer(&x.reffbf02dfd.base))
	x.MemberLines = *NewDictionaryRef(unsafe.Pointer(&x.reffbf02dfd.member_lines))
	x.Methods = *NewArrayRef(unsafe.Pointer(&x.reffbf02dfd.methods))
	x.Signals = *NewArrayRef(unsafe.Pointer(&x.reffbf02dfd.signals))
	x.Properties = *NewArrayRef(unsafe.Pointer(&x.reffbf02dfd.properties))
}

// allocPluginscriptScriptDescMemory allocates memory for type C.godot_pluginscript_script_desc in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPluginscriptScriptDescMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPluginscriptScriptDescValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPluginscriptScriptDescValue = unsafe.Sizeof([1]C.godot_pluginscript_script_desc{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PluginscriptScriptDesc) Ref() *C.godot_pluginscript_script_desc {
	if x == nil {
		return nil
	}
	return x.ref1aab3210
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PluginscriptScriptDesc) Free() {
	if x != nil && x.allocs1aab3210 != nil {
		x.allocs1aab3210.(*cgoAllocMap).Free()
		x.ref1aab3210 = nil
	}
}

// NewPluginscriptScriptDescRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPluginscriptScriptDescRef(ref unsafe.Pointer) *PluginscriptScriptDesc {
	if ref == nil {
		return nil
	}
	obj := new(PluginscriptScriptDesc)
	obj.ref1aab3210 = (*C.godot_pluginscript_script_desc)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PluginscriptScriptDesc) PassRef() (*C.godot_pluginscript_script_desc, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref1aab3210 != nil {
		return x.ref1aab3210, nil
	}
	mem1aab3210 := allocPluginscriptScriptDescMemory(1)
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
func (x PluginscriptScriptDesc) PassValue() (C.godot_pluginscript_script_desc, *cgoAllocMap) {
	if x.ref1aab3210 != nil {
		return *x.ref1aab3210, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PluginscriptScriptDesc) Deref() {
	if x.ref1aab3210 == nil {
		return
	}
	x.Init = NewPluginscriptScriptManifestRef(unsafe.Pointer(x.ref1aab3210.init))
	x.Finish = NewRef(unsafe.Pointer(x.ref1aab3210.finish))
	x.InstanceDesc = *NewPluginscriptInstanceDescRef(unsafe.Pointer(&x.ref1aab3210.instance_desc))
}

// allocPluginscriptProfilingDataMemory allocates memory for type C.godot_pluginscript_profiling_data in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPluginscriptProfilingDataMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPluginscriptProfilingDataValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPluginscriptProfilingDataValue = unsafe.Sizeof([1]C.godot_pluginscript_profiling_data{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PluginscriptProfilingData) Ref() *C.godot_pluginscript_profiling_data {
	if x == nil {
		return nil
	}
	return x.ref9c004e5a
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PluginscriptProfilingData) Free() {
	if x != nil && x.allocs9c004e5a != nil {
		x.allocs9c004e5a.(*cgoAllocMap).Free()
		x.ref9c004e5a = nil
	}
}

// NewPluginscriptProfilingDataRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPluginscriptProfilingDataRef(ref unsafe.Pointer) *PluginscriptProfilingData {
	if ref == nil {
		return nil
	}
	obj := new(PluginscriptProfilingData)
	obj.ref9c004e5a = (*C.godot_pluginscript_profiling_data)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PluginscriptProfilingData) PassRef() (*C.godot_pluginscript_profiling_data, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref9c004e5a != nil {
		return x.ref9c004e5a, nil
	}
	mem9c004e5a := allocPluginscriptProfilingDataMemory(1)
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
func (x PluginscriptProfilingData) PassValue() (C.godot_pluginscript_profiling_data, *cgoAllocMap) {
	if x.ref9c004e5a != nil {
		return *x.ref9c004e5a, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PluginscriptProfilingData) Deref() {
	if x.ref9c004e5a == nil {
		return
	}
	x.Signature = *NewStringNameRef(unsafe.Pointer(&x.ref9c004e5a.signature))
	x.CallCount = (Int)(x.ref9c004e5a.call_count)
	x.TotalTime = (Int)(x.ref9c004e5a.total_time)
	x.SelfTime = (Int)(x.ref9c004e5a.self_time)
}

// allocPluginscriptLanguageDescMemory allocates memory for type C.godot_pluginscript_language_desc in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPluginscriptLanguageDescMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPluginscriptLanguageDescValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPluginscriptLanguageDescValue = unsafe.Sizeof([1]C.godot_pluginscript_language_desc{})

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

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *PluginscriptLanguageDesc) Ref() *C.godot_pluginscript_language_desc {
	if x == nil {
		return nil
	}
	return x.refdac22bbe
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *PluginscriptLanguageDesc) Free() {
	if x != nil && x.allocsdac22bbe != nil {
		x.allocsdac22bbe.(*cgoAllocMap).Free()
		x.refdac22bbe = nil
	}
}

// NewPluginscriptLanguageDescRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewPluginscriptLanguageDescRef(ref unsafe.Pointer) *PluginscriptLanguageDesc {
	if ref == nil {
		return nil
	}
	obj := new(PluginscriptLanguageDesc)
	obj.refdac22bbe = (*C.godot_pluginscript_language_desc)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *PluginscriptLanguageDesc) PassRef() (*C.godot_pluginscript_language_desc, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refdac22bbe != nil {
		return x.refdac22bbe, nil
	}
	memdac22bbe := allocPluginscriptLanguageDescMemory(1)
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
func (x PluginscriptLanguageDesc) PassValue() (C.godot_pluginscript_language_desc, *cgoAllocMap) {
	if x.refdac22bbe != nil {
		return *x.refdac22bbe, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *PluginscriptLanguageDesc) Deref() {
	if x.refdac22bbe == nil {
		return
	}
	x.Name = packPCharString(x.refdac22bbe.name)
	x.Type = packPCharString(x.refdac22bbe._type)
	x.Extension = packPCharString(x.refdac22bbe.extension)
	packSString(x.RecognizedExtensions, x.refdac22bbe.recognized_extensions)
	x.Init = NewPluginscriptLanguageDataRef(unsafe.Pointer(x.refdac22bbe.init))
	x.Finish = NewRef(unsafe.Pointer(x.refdac22bbe.finish))
	packSString(x.ReservedWords, x.refdac22bbe.reserved_words)
	packSString(x.CommentDelimiters, x.refdac22bbe.comment_delimiters)
	packSString(x.StringDelimiters, x.refdac22bbe.string_delimiters)
	x.HasNamedClasses = (Bool)(x.refdac22bbe.has_named_classes)
	x.SupportsBuiltinMode = (Bool)(x.refdac22bbe.supports_builtin_mode)
	x.GetTemplateSourceCode = NewStringRef(unsafe.Pointer(x.refdac22bbe.get_template_source_code))
	x.Validate = NewBoolRef(unsafe.Pointer(x.refdac22bbe.validate))
	x.FindFunction = NewRef(unsafe.Pointer(x.refdac22bbe.find_function))
	x.MakeFunction = NewStringRef(unsafe.Pointer(x.refdac22bbe.make_function))
	x.CompleteCode = NewErrorRef(unsafe.Pointer(x.refdac22bbe.complete_code))
	x.AutoIndentCode = NewRef(unsafe.Pointer(x.refdac22bbe.auto_indent_code))
	x.AddGlobalConstant = NewRef(unsafe.Pointer(x.refdac22bbe.add_global_constant))
	x.DebugGetError = NewStringRef(unsafe.Pointer(x.refdac22bbe.debug_get_error))
	x.DebugGetStackLevelCount = NewRef(unsafe.Pointer(x.refdac22bbe.debug_get_stack_level_count))
	x.DebugGetStackLevelLine = NewRef(unsafe.Pointer(x.refdac22bbe.debug_get_stack_level_line))
	x.DebugGetStackLevelFunction = NewStringRef(unsafe.Pointer(x.refdac22bbe.debug_get_stack_level_function))
	x.DebugGetStackLevelSource = NewStringRef(unsafe.Pointer(x.refdac22bbe.debug_get_stack_level_source))
	x.DebugGetStackLevelLocals = NewRef(unsafe.Pointer(x.refdac22bbe.debug_get_stack_level_locals))
	x.DebugGetStackLevelMembers = NewRef(unsafe.Pointer(x.refdac22bbe.debug_get_stack_level_members))
	x.DebugGetGlobals = NewRef(unsafe.Pointer(x.refdac22bbe.debug_get_globals))
	x.DebugParseStackLevelExpression = NewStringRef(unsafe.Pointer(x.refdac22bbe.debug_parse_stack_level_expression))
	x.GetPublicFunctions = NewRef(unsafe.Pointer(x.refdac22bbe.get_public_functions))
	x.GetPublicConstants = NewRef(unsafe.Pointer(x.refdac22bbe.get_public_constants))
	x.ProfilingStart = NewRef(unsafe.Pointer(x.refdac22bbe.profiling_start))
	x.ProfilingStop = NewRef(unsafe.Pointer(x.refdac22bbe.profiling_stop))
	x.ProfilingGetAccumulatedData = NewRef(unsafe.Pointer(x.refdac22bbe.profiling_get_accumulated_data))
	x.ProfilingGetFrameData = NewRef(unsafe.Pointer(x.refdac22bbe.profiling_get_frame_data))
	x.ProfilingFrame = NewRef(unsafe.Pointer(x.refdac22bbe.profiling_frame))
	x.ScriptDesc = *NewPluginscriptScriptDescRef(unsafe.Pointer(&x.refdac22bbe.script_desc))
}

// unpackArgSMethodBind transforms a sliced Go data structure into plain C format.
func unpackArgSMethodBind(x []MethodBind) (unpacked *C.godot_method_bind, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_method_bind) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocMethodBindMemory(len0)
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

// packSMethodBind reads sliced Go data structure out from plain C format.
func packSMethodBind(v []MethodBind, ptr0 *C.godot_method_bind) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfMethodBindValue]C.godot_method_bind)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewMethodBindRef(unsafe.Pointer(&ptr1))
	}
}

// allocPVariantMemory allocates memory for type *C.godot_variant in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPVariantMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPVariantValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPVariantValue = unsafe.Sizeof([1]*C.godot_variant{})

// unpackArgSSVariant transforms a sliced Go data structure into plain C format.
func unpackArgSSVariant(x [][]Variant) (unpacked **C.godot_variant, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.godot_variant) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPVariantMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.godot_variant)(unsafe.Pointer(h0))
	for i0 := range x {
		len1 := len(x[i0])
		mem1 := allocVariantMemory(len1)
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

// packSSVariant reads sliced Go data structure out from plain C format.
func packSSVariant(v [][]Variant, ptr0 **C.godot_variant) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.godot_variant)(unsafe.Pointer(ptr0)))[i0]
		for i1 := range v[i0] {
			ptr2 := (*(*[m / sizeOfVariantValue]C.godot_variant)(unsafe.Pointer(ptr1)))[i1]
			v[i0][i1] = *NewVariantRef(unsafe.Pointer(&ptr2))
		}
	}
}

// unpackArgSVariantCallError transforms a sliced Go data structure into plain C format.
func unpackArgSVariantCallError(x []VariantCallError) (unpacked *C.godot_variant_call_error, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_variant_call_error) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocVariantCallErrorMemory(len0)
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

// packSVariantCallError reads sliced Go data structure out from plain C format.
func packSVariantCallError(v []VariantCallError, ptr0 *C.godot_variant_call_error) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfVariantCallErrorValue]C.godot_variant_call_error)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewVariantCallErrorRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSString transforms a sliced Go data structure into plain C format.
func unpackArgSString(x []String) (unpacked *C.godot_string, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_string) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocStringMemory(len0)
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

// unpackArgSArray transforms a sliced Go data structure into plain C format.
func unpackArgSArray(x []Array) (unpacked *C.godot_array, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_array) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocArrayMemory(len0)
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

// unpackArgSVariant transforms a sliced Go data structure into plain C format.
func unpackArgSVariant(x []Variant) (unpacked *C.godot_variant, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_variant) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocVariantMemory(len0)
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

// unpackArgSPoolColorArray transforms a sliced Go data structure into plain C format.
func unpackArgSPoolColorArray(x []PoolColorArray) (unpacked *C.godot_pool_color_array, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_color_array) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPoolColorArrayMemory(len0)
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

// packSPoolColorArray reads sliced Go data structure out from plain C format.
func packSPoolColorArray(v []PoolColorArray, ptr0 *C.godot_pool_color_array) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPoolColorArrayValue]C.godot_pool_color_array)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewPoolColorArrayRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSPoolVector3Array transforms a sliced Go data structure into plain C format.
func unpackArgSPoolVector3Array(x []PoolVector3Array) (unpacked *C.godot_pool_vector3_array, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_vector3_array) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPoolVector3ArrayMemory(len0)
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

// packSPoolVector3Array reads sliced Go data structure out from plain C format.
func packSPoolVector3Array(v []PoolVector3Array, ptr0 *C.godot_pool_vector3_array) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPoolVector3ArrayValue]C.godot_pool_vector3_array)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewPoolVector3ArrayRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSPoolVector2Array transforms a sliced Go data structure into plain C format.
func unpackArgSPoolVector2Array(x []PoolVector2Array) (unpacked *C.godot_pool_vector2_array, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_vector2_array) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPoolVector2ArrayMemory(len0)
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

// packSPoolVector2Array reads sliced Go data structure out from plain C format.
func packSPoolVector2Array(v []PoolVector2Array, ptr0 *C.godot_pool_vector2_array) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPoolVector2ArrayValue]C.godot_pool_vector2_array)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewPoolVector2ArrayRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSPoolStringArray transforms a sliced Go data structure into plain C format.
func unpackArgSPoolStringArray(x []PoolStringArray) (unpacked *C.godot_pool_string_array, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_string_array) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPoolStringArrayMemory(len0)
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

// packSPoolStringArray reads sliced Go data structure out from plain C format.
func packSPoolStringArray(v []PoolStringArray, ptr0 *C.godot_pool_string_array) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPoolStringArrayValue]C.godot_pool_string_array)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewPoolStringArrayRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSPoolRealArray transforms a sliced Go data structure into plain C format.
func unpackArgSPoolRealArray(x []PoolRealArray) (unpacked *C.godot_pool_real_array, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_real_array) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPoolRealArrayMemory(len0)
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

// packSPoolRealArray reads sliced Go data structure out from plain C format.
func packSPoolRealArray(v []PoolRealArray, ptr0 *C.godot_pool_real_array) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPoolRealArrayValue]C.godot_pool_real_array)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewPoolRealArrayRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSPoolIntArray transforms a sliced Go data structure into plain C format.
func unpackArgSPoolIntArray(x []PoolIntArray) (unpacked *C.godot_pool_int_array, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_int_array) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPoolIntArrayMemory(len0)
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

// packSPoolIntArray reads sliced Go data structure out from plain C format.
func packSPoolIntArray(v []PoolIntArray, ptr0 *C.godot_pool_int_array) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPoolIntArrayValue]C.godot_pool_int_array)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewPoolIntArrayRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSPoolByteArray transforms a sliced Go data structure into plain C format.
func unpackArgSPoolByteArray(x []PoolByteArray) (unpacked *C.godot_pool_byte_array, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_byte_array) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPoolByteArrayMemory(len0)
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

// packSPoolByteArray reads sliced Go data structure out from plain C format.
func packSPoolByteArray(v []PoolByteArray, ptr0 *C.godot_pool_byte_array) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPoolByteArrayValue]C.godot_pool_byte_array)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewPoolByteArrayRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSVector2 transforms a sliced Go data structure into plain C format.
func unpackArgSVector2(x []Vector2) (unpacked *C.godot_vector2, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_vector2) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocVector2Memory(len0)
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

// packSVector2 reads sliced Go data structure out from plain C format.
func packSVector2(v []Vector2, ptr0 *C.godot_vector2) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfVector2Value]C.godot_vector2)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewVector2Ref(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSVector3 transforms a sliced Go data structure into plain C format.
func unpackArgSVector3(x []Vector3) (unpacked *C.godot_vector3, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_vector3) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocVector3Memory(len0)
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

// packSVector3 reads sliced Go data structure out from plain C format.
func packSVector3(v []Vector3, ptr0 *C.godot_vector3) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfVector3Value]C.godot_vector3)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewVector3Ref(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSColor transforms a sliced Go data structure into plain C format.
func unpackArgSColor(x []Color) (unpacked *C.godot_color, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_color) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocColorMemory(len0)
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

// packSColor reads sliced Go data structure out from plain C format.
func packSColor(v []Color, ptr0 *C.godot_color) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfColorValue]C.godot_color)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewColorRef(unsafe.Pointer(&ptr1))
	}
}

// allocPoolByteArrayReadAccessMemory allocates memory for type C.godot_pool_byte_array_read_access in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPoolByteArrayReadAccessMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPoolByteArrayReadAccessValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPoolByteArrayReadAccessValue = unsafe.Sizeof([1]C.godot_pool_byte_array_read_access{})

// unpackArgSPoolByteArrayReadAccess transforms a sliced Go data structure into plain C format.
func unpackArgSPoolByteArrayReadAccess(x []PoolByteArrayReadAccess) (unpacked *C.godot_pool_byte_array_read_access, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_byte_array_read_access) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPoolByteArrayReadAccessMemory(len0)
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

// packSPoolByteArrayReadAccess reads sliced Go data structure out from plain C format.
func packSPoolByteArrayReadAccess(v []PoolByteArrayReadAccess, ptr0 *C.godot_pool_byte_array_read_access) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPoolByteArrayReadAccessValue]C.godot_pool_byte_array_read_access)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewPoolByteArrayReadAccessRef(unsafe.Pointer(&ptr1))
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

// allocPoolIntArrayReadAccessMemory allocates memory for type C.godot_pool_int_array_read_access in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPoolIntArrayReadAccessMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPoolIntArrayReadAccessValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPoolIntArrayReadAccessValue = unsafe.Sizeof([1]C.godot_pool_int_array_read_access{})

// unpackArgSPoolIntArrayReadAccess transforms a sliced Go data structure into plain C format.
func unpackArgSPoolIntArrayReadAccess(x []PoolIntArrayReadAccess) (unpacked *C.godot_pool_int_array_read_access, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_int_array_read_access) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPoolIntArrayReadAccessMemory(len0)
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

// packSPoolIntArrayReadAccess reads sliced Go data structure out from plain C format.
func packSPoolIntArrayReadAccess(v []PoolIntArrayReadAccess, ptr0 *C.godot_pool_int_array_read_access) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPoolIntArrayReadAccessValue]C.godot_pool_int_array_read_access)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewPoolIntArrayReadAccessRef(unsafe.Pointer(&ptr1))
	}
}

// allocPoolRealArrayReadAccessMemory allocates memory for type C.godot_pool_real_array_read_access in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPoolRealArrayReadAccessMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPoolRealArrayReadAccessValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPoolRealArrayReadAccessValue = unsafe.Sizeof([1]C.godot_pool_real_array_read_access{})

// unpackArgSPoolRealArrayReadAccess transforms a sliced Go data structure into plain C format.
func unpackArgSPoolRealArrayReadAccess(x []PoolRealArrayReadAccess) (unpacked *C.godot_pool_real_array_read_access, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_real_array_read_access) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPoolRealArrayReadAccessMemory(len0)
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

// packSPoolRealArrayReadAccess reads sliced Go data structure out from plain C format.
func packSPoolRealArrayReadAccess(v []PoolRealArrayReadAccess, ptr0 *C.godot_pool_real_array_read_access) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPoolRealArrayReadAccessValue]C.godot_pool_real_array_read_access)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewPoolRealArrayReadAccessRef(unsafe.Pointer(&ptr1))
	}
}

// allocPoolStringArrayReadAccessMemory allocates memory for type C.godot_pool_string_array_read_access in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPoolStringArrayReadAccessMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPoolStringArrayReadAccessValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPoolStringArrayReadAccessValue = unsafe.Sizeof([1]C.godot_pool_string_array_read_access{})

// unpackArgSPoolStringArrayReadAccess transforms a sliced Go data structure into plain C format.
func unpackArgSPoolStringArrayReadAccess(x []PoolStringArrayReadAccess) (unpacked *C.godot_pool_string_array_read_access, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_string_array_read_access) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPoolStringArrayReadAccessMemory(len0)
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

// packSPoolStringArrayReadAccess reads sliced Go data structure out from plain C format.
func packSPoolStringArrayReadAccess(v []PoolStringArrayReadAccess, ptr0 *C.godot_pool_string_array_read_access) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPoolStringArrayReadAccessValue]C.godot_pool_string_array_read_access)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewPoolStringArrayReadAccessRef(unsafe.Pointer(&ptr1))
	}
}

// allocPoolVector2ArrayReadAccessMemory allocates memory for type C.godot_pool_vector2_array_read_access in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPoolVector2ArrayReadAccessMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPoolVector2ArrayReadAccessValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPoolVector2ArrayReadAccessValue = unsafe.Sizeof([1]C.godot_pool_vector2_array_read_access{})

// unpackArgSPoolVector2ArrayReadAccess transforms a sliced Go data structure into plain C format.
func unpackArgSPoolVector2ArrayReadAccess(x []PoolVector2ArrayReadAccess) (unpacked *C.godot_pool_vector2_array_read_access, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_vector2_array_read_access) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPoolVector2ArrayReadAccessMemory(len0)
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

// packSPoolVector2ArrayReadAccess reads sliced Go data structure out from plain C format.
func packSPoolVector2ArrayReadAccess(v []PoolVector2ArrayReadAccess, ptr0 *C.godot_pool_vector2_array_read_access) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPoolVector2ArrayReadAccessValue]C.godot_pool_vector2_array_read_access)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewPoolVector2ArrayReadAccessRef(unsafe.Pointer(&ptr1))
	}
}

// allocPoolVector3ArrayReadAccessMemory allocates memory for type C.godot_pool_vector3_array_read_access in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPoolVector3ArrayReadAccessMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPoolVector3ArrayReadAccessValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPoolVector3ArrayReadAccessValue = unsafe.Sizeof([1]C.godot_pool_vector3_array_read_access{})

// unpackArgSPoolVector3ArrayReadAccess transforms a sliced Go data structure into plain C format.
func unpackArgSPoolVector3ArrayReadAccess(x []PoolVector3ArrayReadAccess) (unpacked *C.godot_pool_vector3_array_read_access, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_vector3_array_read_access) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPoolVector3ArrayReadAccessMemory(len0)
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

// packSPoolVector3ArrayReadAccess reads sliced Go data structure out from plain C format.
func packSPoolVector3ArrayReadAccess(v []PoolVector3ArrayReadAccess, ptr0 *C.godot_pool_vector3_array_read_access) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPoolVector3ArrayReadAccessValue]C.godot_pool_vector3_array_read_access)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewPoolVector3ArrayReadAccessRef(unsafe.Pointer(&ptr1))
	}
}

// allocPoolColorArrayReadAccessMemory allocates memory for type C.godot_pool_color_array_read_access in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPoolColorArrayReadAccessMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPoolColorArrayReadAccessValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPoolColorArrayReadAccessValue = unsafe.Sizeof([1]C.godot_pool_color_array_read_access{})

// unpackArgSPoolColorArrayReadAccess transforms a sliced Go data structure into plain C format.
func unpackArgSPoolColorArrayReadAccess(x []PoolColorArrayReadAccess) (unpacked *C.godot_pool_color_array_read_access, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_color_array_read_access) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPoolColorArrayReadAccessMemory(len0)
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

// packSPoolColorArrayReadAccess reads sliced Go data structure out from plain C format.
func packSPoolColorArrayReadAccess(v []PoolColorArrayReadAccess, ptr0 *C.godot_pool_color_array_read_access) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPoolColorArrayReadAccessValue]C.godot_pool_color_array_read_access)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewPoolColorArrayReadAccessRef(unsafe.Pointer(&ptr1))
	}
}

// allocPoolByteArrayWriteAccessMemory allocates memory for type C.godot_pool_byte_array_write_access in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPoolByteArrayWriteAccessMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPoolByteArrayWriteAccessValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPoolByteArrayWriteAccessValue = unsafe.Sizeof([1]C.godot_pool_byte_array_write_access{})

// unpackArgSPoolByteArrayWriteAccess transforms a sliced Go data structure into plain C format.
func unpackArgSPoolByteArrayWriteAccess(x []PoolByteArrayWriteAccess) (unpacked *C.godot_pool_byte_array_write_access, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_byte_array_write_access) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPoolByteArrayWriteAccessMemory(len0)
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

// packSPoolByteArrayWriteAccess reads sliced Go data structure out from plain C format.
func packSPoolByteArrayWriteAccess(v []PoolByteArrayWriteAccess, ptr0 *C.godot_pool_byte_array_write_access) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPoolByteArrayWriteAccessValue]C.godot_pool_byte_array_write_access)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewPoolByteArrayWriteAccessRef(unsafe.Pointer(&ptr1))
	}
}

// allocPoolIntArrayWriteAccessMemory allocates memory for type C.godot_pool_int_array_write_access in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPoolIntArrayWriteAccessMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPoolIntArrayWriteAccessValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPoolIntArrayWriteAccessValue = unsafe.Sizeof([1]C.godot_pool_int_array_write_access{})

// unpackArgSPoolIntArrayWriteAccess transforms a sliced Go data structure into plain C format.
func unpackArgSPoolIntArrayWriteAccess(x []PoolIntArrayWriteAccess) (unpacked *C.godot_pool_int_array_write_access, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_int_array_write_access) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPoolIntArrayWriteAccessMemory(len0)
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

// packSPoolIntArrayWriteAccess reads sliced Go data structure out from plain C format.
func packSPoolIntArrayWriteAccess(v []PoolIntArrayWriteAccess, ptr0 *C.godot_pool_int_array_write_access) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPoolIntArrayWriteAccessValue]C.godot_pool_int_array_write_access)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewPoolIntArrayWriteAccessRef(unsafe.Pointer(&ptr1))
	}
}

// allocPoolRealArrayWriteAccessMemory allocates memory for type C.godot_pool_real_array_write_access in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPoolRealArrayWriteAccessMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPoolRealArrayWriteAccessValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPoolRealArrayWriteAccessValue = unsafe.Sizeof([1]C.godot_pool_real_array_write_access{})

// unpackArgSPoolRealArrayWriteAccess transforms a sliced Go data structure into plain C format.
func unpackArgSPoolRealArrayWriteAccess(x []PoolRealArrayWriteAccess) (unpacked *C.godot_pool_real_array_write_access, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_real_array_write_access) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPoolRealArrayWriteAccessMemory(len0)
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

// packSPoolRealArrayWriteAccess reads sliced Go data structure out from plain C format.
func packSPoolRealArrayWriteAccess(v []PoolRealArrayWriteAccess, ptr0 *C.godot_pool_real_array_write_access) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPoolRealArrayWriteAccessValue]C.godot_pool_real_array_write_access)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewPoolRealArrayWriteAccessRef(unsafe.Pointer(&ptr1))
	}
}

// allocPoolStringArrayWriteAccessMemory allocates memory for type C.godot_pool_string_array_write_access in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPoolStringArrayWriteAccessMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPoolStringArrayWriteAccessValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPoolStringArrayWriteAccessValue = unsafe.Sizeof([1]C.godot_pool_string_array_write_access{})

// unpackArgSPoolStringArrayWriteAccess transforms a sliced Go data structure into plain C format.
func unpackArgSPoolStringArrayWriteAccess(x []PoolStringArrayWriteAccess) (unpacked *C.godot_pool_string_array_write_access, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_string_array_write_access) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPoolStringArrayWriteAccessMemory(len0)
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

// packSPoolStringArrayWriteAccess reads sliced Go data structure out from plain C format.
func packSPoolStringArrayWriteAccess(v []PoolStringArrayWriteAccess, ptr0 *C.godot_pool_string_array_write_access) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPoolStringArrayWriteAccessValue]C.godot_pool_string_array_write_access)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewPoolStringArrayWriteAccessRef(unsafe.Pointer(&ptr1))
	}
}

// allocPoolVector2ArrayWriteAccessMemory allocates memory for type C.godot_pool_vector2_array_write_access in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPoolVector2ArrayWriteAccessMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPoolVector2ArrayWriteAccessValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPoolVector2ArrayWriteAccessValue = unsafe.Sizeof([1]C.godot_pool_vector2_array_write_access{})

// unpackArgSPoolVector2ArrayWriteAccess transforms a sliced Go data structure into plain C format.
func unpackArgSPoolVector2ArrayWriteAccess(x []PoolVector2ArrayWriteAccess) (unpacked *C.godot_pool_vector2_array_write_access, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_vector2_array_write_access) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPoolVector2ArrayWriteAccessMemory(len0)
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

// packSPoolVector2ArrayWriteAccess reads sliced Go data structure out from plain C format.
func packSPoolVector2ArrayWriteAccess(v []PoolVector2ArrayWriteAccess, ptr0 *C.godot_pool_vector2_array_write_access) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPoolVector2ArrayWriteAccessValue]C.godot_pool_vector2_array_write_access)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewPoolVector2ArrayWriteAccessRef(unsafe.Pointer(&ptr1))
	}
}

// allocPoolVector3ArrayWriteAccessMemory allocates memory for type C.godot_pool_vector3_array_write_access in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPoolVector3ArrayWriteAccessMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPoolVector3ArrayWriteAccessValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPoolVector3ArrayWriteAccessValue = unsafe.Sizeof([1]C.godot_pool_vector3_array_write_access{})

// unpackArgSPoolVector3ArrayWriteAccess transforms a sliced Go data structure into plain C format.
func unpackArgSPoolVector3ArrayWriteAccess(x []PoolVector3ArrayWriteAccess) (unpacked *C.godot_pool_vector3_array_write_access, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_vector3_array_write_access) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPoolVector3ArrayWriteAccessMemory(len0)
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

// packSPoolVector3ArrayWriteAccess reads sliced Go data structure out from plain C format.
func packSPoolVector3ArrayWriteAccess(v []PoolVector3ArrayWriteAccess, ptr0 *C.godot_pool_vector3_array_write_access) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPoolVector3ArrayWriteAccessValue]C.godot_pool_vector3_array_write_access)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewPoolVector3ArrayWriteAccessRef(unsafe.Pointer(&ptr1))
	}
}

// allocPoolColorArrayWriteAccessMemory allocates memory for type C.godot_pool_color_array_write_access in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPoolColorArrayWriteAccessMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPoolColorArrayWriteAccessValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPoolColorArrayWriteAccessValue = unsafe.Sizeof([1]C.godot_pool_color_array_write_access{})

// unpackArgSPoolColorArrayWriteAccess transforms a sliced Go data structure into plain C format.
func unpackArgSPoolColorArrayWriteAccess(x []PoolColorArrayWriteAccess) (unpacked *C.godot_pool_color_array_write_access, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pool_color_array_write_access) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPoolColorArrayWriteAccessMemory(len0)
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

// packSPoolColorArrayWriteAccess reads sliced Go data structure out from plain C format.
func packSPoolColorArrayWriteAccess(v []PoolColorArrayWriteAccess, ptr0 *C.godot_pool_color_array_write_access) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPoolColorArrayWriteAccessValue]C.godot_pool_color_array_write_access)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewPoolColorArrayWriteAccessRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSBasis transforms a sliced Go data structure into plain C format.
func unpackArgSBasis(x []Basis) (unpacked *C.godot_basis, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_basis) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocBasisMemory(len0)
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

// packSBasis reads sliced Go data structure out from plain C format.
func packSBasis(v []Basis, ptr0 *C.godot_basis) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfBasisValue]C.godot_basis)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewBasisRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSQuat transforms a sliced Go data structure into plain C format.
func unpackArgSQuat(x []Quat) (unpacked *C.godot_quat, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_quat) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocQuatMemory(len0)
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

// packSQuat reads sliced Go data structure out from plain C format.
func packSQuat(v []Quat, ptr0 *C.godot_quat) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfQuatValue]C.godot_quat)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewQuatRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSRect2 transforms a sliced Go data structure into plain C format.
func unpackArgSRect2(x []Rect2) (unpacked *C.godot_rect2, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_rect2) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocRect2Memory(len0)
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

// packSRect2 reads sliced Go data structure out from plain C format.
func packSRect2(v []Rect2, ptr0 *C.godot_rect2) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfRect2Value]C.godot_rect2)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewRect2Ref(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSTransform2d transforms a sliced Go data structure into plain C format.
func unpackArgSTransform2d(x []Transform2d) (unpacked *C.godot_transform2d, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_transform2d) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocTransform2dMemory(len0)
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

// packSTransform2d reads sliced Go data structure out from plain C format.
func packSTransform2d(v []Transform2d, ptr0 *C.godot_transform2d) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfTransform2dValue]C.godot_transform2d)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewTransform2dRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSPlane transforms a sliced Go data structure into plain C format.
func unpackArgSPlane(x []Plane) (unpacked *C.godot_plane, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_plane) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPlaneMemory(len0)
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

// packSPlane reads sliced Go data structure out from plain C format.
func packSPlane(v []Plane, ptr0 *C.godot_plane) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPlaneValue]C.godot_plane)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewPlaneRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSAabb transforms a sliced Go data structure into plain C format.
func unpackArgSAabb(x []Aabb) (unpacked *C.godot_aabb, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_aabb) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocAabbMemory(len0)
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

// packSAabb reads sliced Go data structure out from plain C format.
func packSAabb(v []Aabb, ptr0 *C.godot_aabb) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfAabbValue]C.godot_aabb)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewAabbRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSTransform transforms a sliced Go data structure into plain C format.
func unpackArgSTransform(x []Transform) (unpacked *C.godot_transform, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_transform) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocTransformMemory(len0)
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

// packSTransform reads sliced Go data structure out from plain C format.
func packSTransform(v []Transform, ptr0 *C.godot_transform) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfTransformValue]C.godot_transform)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewTransformRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSNodePath transforms a sliced Go data structure into plain C format.
func unpackArgSNodePath(x []NodePath) (unpacked *C.godot_node_path, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_node_path) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocNodePathMemory(len0)
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

// packSNodePath reads sliced Go data structure out from plain C format.
func packSNodePath(v []NodePath, ptr0 *C.godot_node_path) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfNodePathValue]C.godot_node_path)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewNodePathRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSRid transforms a sliced Go data structure into plain C format.
func unpackArgSRid(x []Rid) (unpacked *C.godot_rid, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_rid) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocRidMemory(len0)
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

// packSRid reads sliced Go data structure out from plain C format.
func packSRid(v []Rid, ptr0 *C.godot_rid) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfRidValue]C.godot_rid)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewRidRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSDictionary transforms a sliced Go data structure into plain C format.
func unpackArgSDictionary(x []Dictionary) (unpacked *C.godot_dictionary, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_dictionary) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocDictionaryMemory(len0)
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

// packSDictionary reads sliced Go data structure out from plain C format.
func packSDictionary(v []Dictionary, ptr0 *C.godot_dictionary) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfDictionaryValue]C.godot_dictionary)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewDictionaryRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSStringName transforms a sliced Go data structure into plain C format.
func unpackArgSStringName(x []StringName) (unpacked *C.godot_string_name, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_string_name) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocStringNameMemory(len0)
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

// packSStringName reads sliced Go data structure out from plain C format.
func packSStringName(v []StringName, ptr0 *C.godot_string_name) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfStringNameValue]C.godot_string_name)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewStringNameRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSArvrInterfaceGdnative transforms a sliced Go data structure into plain C format.
func unpackArgSArvrInterfaceGdnative(x []ArvrInterfaceGdnative) (unpacked *C.godot_arvr_interface_gdnative, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_arvr_interface_gdnative) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocArvrInterfaceGdnativeMemory(len0)
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

// packSArvrInterfaceGdnative reads sliced Go data structure out from plain C format.
func packSArvrInterfaceGdnative(v []ArvrInterfaceGdnative, ptr0 *C.godot_arvr_interface_gdnative) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfArvrInterfaceGdnativeValue]C.godot_arvr_interface_gdnative)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewArvrInterfaceGdnativeRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSPropertyAttributes transforms a sliced Go data structure into plain C format.
func unpackArgSPropertyAttributes(x []PropertyAttributes) (unpacked *C.godot_property_attributes, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_property_attributes) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPropertyAttributesMemory(len0)
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

// packSPropertyAttributes reads sliced Go data structure out from plain C format.
func packSPropertyAttributes(v []PropertyAttributes, ptr0 *C.godot_property_attributes) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPropertyAttributesValue]C.godot_property_attributes)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewPropertyAttributesRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSSignal transforms a sliced Go data structure into plain C format.
func unpackArgSSignal(x []Signal) (unpacked *C.godot_signal, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_signal) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocSignalMemory(len0)
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

// packSSignal reads sliced Go data structure out from plain C format.
func packSSignal(v []Signal, ptr0 *C.godot_signal) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfSignalValue]C.godot_signal)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewSignalRef(unsafe.Pointer(&ptr1))
	}
}

// unpackArgSPluginscriptLanguageDesc transforms a sliced Go data structure into plain C format.
func unpackArgSPluginscriptLanguageDesc(x []PluginscriptLanguageDesc) (unpacked *C.godot_pluginscript_language_desc, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(**C.godot_pluginscript_language_desc) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPluginscriptLanguageDescMemory(len0)
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

// packSPluginscriptLanguageDesc reads sliced Go data structure out from plain C format.
func packSPluginscriptLanguageDesc(v []PluginscriptLanguageDesc, ptr0 *C.godot_pluginscript_language_desc) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPluginscriptLanguageDescValue]C.godot_pluginscript_language_desc)(unsafe.Pointer(ptr0)))[i0]
		v[i0] = *NewPluginscriptLanguageDescRef(unsafe.Pointer(&ptr1))
	}
}
