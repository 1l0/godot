%module godot
%insert(cgo_comment_typedefs) %{
#cgo CFLAGS: -Igodot_headers -std=c11
#cgo linux LDFLAGS: -Wl,-unresolved-symbols=ignore-all
#cgo darwin LDFLAGS: -Wl,-undefined,dynamic_lookup
%}
%{
#include "gdnative_api_struct.gen.h"
%}

%include "godot_headers/gdnative_api_struct.gen.h"