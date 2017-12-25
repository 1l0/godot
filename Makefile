gen:
	git pull -s subtree godot-headers master
	swig -go -cgo -intgosize 32 -Igodot_headers godot.swig_
patch:
	sed -i -e 's/ Struct_SS_godot_gdnative_api_struct/ SwigcptrStruct_SS_godot_gdnative_api_struct/g' godot.go
	sed -i -e 's/(Struct_SS_godot_gdnative_api_struct/(SwigcptrStruct_SS_godot_gdnative_api_struct/g' godot.go
	sed -i -e 's/Godot//g' godot.go
test:
	go build
clean:
	rm godot.go godot_wrap.c godot.go-e
