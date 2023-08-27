package main

import (
	"flag"

	nexserver "github.com/ShreyanJain9/nex-server/pkg"
)

func main() {
	configPath := flag.String("config", "/etc/nex/config.toml", "Path to the config TOML file")
	flag.Parse()

	config := nexserver.Config{}
	config.Load(*configPath)
	config.StartServer()
}
