package nexserver

import (
	"fmt"
	"net"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	RootDir  string `default:"/var/nex/data/"`
	Port     int    `default:"1900"`
	Listener net.Listener
	Conn     net.Conn
}

func (c *Config) Load(filepath string) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	err = toml.Unmarshal(file, c)
	if err != nil {
		panic(err)
	}
}

func (c *Config) StartServer() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", c.Port))
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		return
	}
	c.Listener = listener

	fmt.Printf("Nex server is listening on port %d\n", c.Port)

	for {
		conn, err := c.Listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %s\n", err)
			continue
		}
		c.Conn = conn
		go c.HandleConnection()
	}
}
