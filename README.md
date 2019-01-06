# GetSetDB
An extremely optimised and lightweight NoSQL in memory database system written in Go.

### Systems Supported 
- Debian / Ubuntu 16.04 and up
- Fedora / Red Hat Enterprise Linux
- MacOS 

*note* - Windows support not determined as it's not been tested

### Installation

To see the instructions for building from source, [refer this](#building-from-source).

- for Ubuntu / Debian users - 
    - option 1 - download a pre-built binary and place it in `/usr/bin/`
    - option 2 - download a pre-built `.deb` package from here - [getsetdb-v1.deb](https://github.com/getsetdb/getsetdb/raw/master/packages/debian/getsetdb-v1.deb)
- for Fedora / Red Hat Enterprise Linux users - 
    - download a pre-built binary and place it in `/usr/bin/`
- for MacOS users
    - download a pre-built binary and place it in `/usr/bin/`
    
### Usage

Since GetSetDB is an extremely lightweight database, [A Tour of GetSetDB](https://medium.com/@mentix02/a-tour-of-getsetdb-8716c39e354d) should be enough for most people to get up and running with it.

If you wish to use it in conjunction with a programming language, please look out for [upcoming connectors for different languages](#connectors-in-development).

### Documentation

For a tutorial that covers everything from the structuring of GetSetDB to the syntax for the commands, check out [A Tour of GetSetDB](https://medium.com/@mentix02/a-tour-of-getsetdb-8716c39e354d)

API documentation for reference will be coming soon.

### Building From Source

GetSetDB is written purely in Go, so you need to have Go installed on your system to build it from souce. [Install it from here](https://golang.org/doc/install). Next, configure your `$GOPATH` - by default, it is your `home` directory. 

```sh
$ mkdir -p ~/go/src/
$ go get github.com/getsetdb/getsetdb
```

This should automatically build GetSetDB for you and will place the executable binary in `~/go/bin/`. To run it from anywhere, you'll have to add this binary to your `$PATH`. To do so, run the following - 

```sh
$ echo 'export PATH="$HOME/go/bin:$PATH"' >> ~/.bashrc
```

### Developing

If you're a developer looking for documentation on making connectors for a language that [GetSetDB doesn't already support](#connectors-in-development) , please be paitent - documentation for that will be coming soon.

Same goes for contributing to the database itself.

### Connectors In Development
1. [Python](https://github.com/getsetdb/getsetpy)
2. [Go](https://github.com/getsetdb/getsetgo)
3. [Ruby](https://github.com/getsetdb/getsetrub)
4. [C++](https://github.com/getsetdb/getsetc)
5. [Perl](https://github.com/getsetdb/getsetpl)
