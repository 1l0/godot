package main

import (
	"github.com/1l0/godot"
)

var (
	api godot.Godot_gdnative_core_api_struct
)

func init() {
	api = godot.NewGodot_gdnative_core_api_struct()
	defer godot.DeleteGodot_gdnative_core_api_struct(api)
}
func main() {}
