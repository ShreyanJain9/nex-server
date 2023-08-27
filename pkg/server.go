package nexserver

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (c *Config) HandleConnection() {
	defer c.Conn.Close()

	scanner := bufio.NewScanner(c.Conn)
	if scanner.Scan() {
		path := strings.TrimSpace(scanner.Text())
		c.ServeDocument(path)
	}
}

func (c *Config) ServeDocument(path string) {
	fullPath := filepath.Join(c.RootDir, path)
	isDirectory := path == "" || strings.HasSuffix(path, "/")

	if isDirectory {
		c.ServeDirectory(fullPath)
	} else {
		c.ServeFile(fullPath)
	}
}

func (c *Config) ServeDirectory(dirPath string) {
	fileInfos, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Fprintf(c.Conn, "Error reading directory: %s\n", err)
		return
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			fmt.Fprintf(c.Conn, "=> %s/\n", fileInfo.Name())
		} else {
			fmt.Fprintf(c.Conn, "=> %s\n", fileInfo.Name())
		}
	}
}

func (c *Config) ServeFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(c.Conn, "Error opening file: %s\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Fprintln(c.Conn, scanner.Text())
	}
}
