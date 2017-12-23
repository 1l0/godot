package godot

//go:generate git pull -s subtree godot_headers master
//go:generate c-for-go -ccincl -out ../ godot.yml
