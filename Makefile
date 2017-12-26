all: gen patch install

gen:
	git pull -s subtree godot-headers master
	swig -go -cgo -intgosize 32 -Igodot_headers godot.swig_
patch:
	sed -i -e 's/Godot//g' godot.go
	sed -i -e 's/Godot//g' godot_wrap.c
	sed -i -e 's/ Struct_SS_godot_gdnative_api_struct/ GdnativeApiStruct/g' godot.go
	sed -i -e 's/(Struct_SS_godot_gdnative_api_struct/(GdnativeApiStruct/g' godot.go

	# after above do mod GetExtensions() to return []GdnativeApiStruct
#func (arg1 SwigcptrGdnativeCoreApiStruct) GetExtensions() (_swig_ret []GdnativeApiStruct) {
#	_swig_i_0 := arg1
#	var swig_r []SwigcptrStruct_SS_godot_gdnative_api_struct
#	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&swig_r)))
#	sliceHeader.Cap = 4
#	sliceHeader.Len = 4
#	sliceHeader.Data = uintptr(C._wrap_GdnativeCoreApiStruct_Extensions_get_godot_ab55971e86d21eb7(C.uintptr_t(_swig_i_0)))
#	var swig_r2 []GdnativeApiStruct
#	for _, s := range swig_r {
#		swig_r2 = append(swig_r2, GdnativeApiStruct(s))
#	}
#return swig_r2
#}
test:
	go build
clean:
	rm godot.go godot_wrap.c godot.go-e
install:
	go install
