# Nex Server

The Nex Server is a simple server that serves files and directory listings over a TCP connection. It is designed to be easy to use and can be customized using a configuration file in TOML format.

## Getting Started

1. **Installation**

   Clone this repository to your local machine:

   ```bash
   git clone https://github.com/your-username/nexserver.git
   cd nexserver
   cd build
   go build main.go
   sudo mv ./main /usr/local/bin/nex
   ```

   I'll add a Makefile soon.

2. **Configuration**

   The Nex Server can be configured using a TOML configuration file. By default, it looks for the configuration file at `/etc/nex/config.toml`. You can specify a different configuration file using the `-config` command-line flag.

   The configuration file should contain the following parameters:

   ```toml
   RootDir = "/path/to/your/files"
   Port = 1900
   ```

   - `RootDir`: The root directory from which files will be served.
   - `Port`: The port number on which the server will listen for incoming connections.

3. **Running the Server**

   To run the Nex Server with the default configuration file path:

   ```bash
   nex
   ```

   To run the server with a custom configuration file:

   ```bash
   nex -config /path/to/your/config.toml
   ```

## The Nex Protocol

The Nex protocol is a simple text-based protocol that allows clients to request files or directory listings from the Nex Server. Read more [here](https://nightfall.city/nex/info/)

1. **File Request**

   To request a file, the client should send the path of the file it wants to retrieve. The server will respond with the contents of the requested file.

   ```
   client: /path/to/file.txt
   server: Contents of the file...
   ```

2. **Directory Listing Request**

   To request a directory listing, the client should send the path of the directory it wants to list. The server will respond with a list of files and subdirectories in the requested directory.

   ```
   client: /path/to/directory/
   server: => file1.txt
   server: => subdirectory/
   ```

## Contributing

Contributions to the Nex Server are welcome! If you encounter any issues, have suggestions for improvements, or would like to add features, feel free to open an issue or submit a pull request.

## License

This project is licensed under the GNU AGPLv3. See the [LICENSE](LICENSE) file for details.

