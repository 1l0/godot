%module godot
%{
#include "gdnative_api_struct.gen.h"
%}

%include "godot_headers/gdnative_api_struct.gen.h"
%include "godot_headers/gdnative/gdnative.h"
%include "godot_headers/arvr/godot_arvr.h"
%include "godot_headers/nativescript/godot_nativescript.h"
%include "godot_headers/pluginscript/godot_pluginscript.h"