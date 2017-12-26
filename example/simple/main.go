package main

import (
	"C"
	"log"

	"github.com/1l0/godot"
)

func main() {}

func init() {
	//runtime.LockOSThread()
	log.SetFlags(log.Lshortfile)
	godot.Init(nativescriptInit)
}

func nativescriptInit(handle godot.Handle) {
	version := godot.API.GetVersion()
	log.Printf("GDNative version: %d.%d", version.GetMajor(), version.GetMinor())
	version = godot.NativescriptAPI.GetVersion()
	log.Printf("Nativescript version: %d.%d", version.GetMajor(), version.GetMinor())
}
