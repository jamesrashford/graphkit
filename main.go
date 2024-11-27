package main

import "github.com/jamesrashford/graphkit/webui"

//go:generate npm run build

func main() {
	addr := ":8080"
	webui.StartServer(addr)
}
