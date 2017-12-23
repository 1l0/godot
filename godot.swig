%module godot
%{
#include "gdnative_api_struct.gen.h"
%}

%rename("%(camelcase)s") "";

#define __attribute__(x)

%include "gdnative_api_struct.gen.h"
%include "gdnative/gdnative.h"

