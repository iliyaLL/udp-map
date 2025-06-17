# Go UDP Key-Value Store

A lightweight in-memory key-value store implementation in Go with support for key expiration.

## Learning Objectives
- TCP/IP
- UDP
- Memory Object Caching
- Networking Concepts

## Features

- Set key-value pairs with optional expiration time
- Get values by key
- Thread-safe operations using mutex locks
- Automatic key expiration handling

## Building and Testing

The program can be compiled by the following command in the project root directory:

```bash
$ go build -o kvstore ./cmd
```

The default port for the program is 8080, but it can be changed with the optional argument --port.
```bash
$ ./kvstore --port 9000
```

For debugging and testing your application, it is better to use the built-in linux utility [netcat](https://linux.die.net/man/1/nc).

    NAME
       nc - TCP/IP swiss army knife


    SYNOPSIS
       nc [-options] hostname port[s] [ports] ...
       nc -l -p port [-options] [hostname] [port]


    DESCRIPTION
        netcat  is  a  simple unix utility which reads and writes data across network connections, using TCP or UDP protocol. It is designed to be a reliable "back-end"  tool  that  can  be used  directly  or easily driven by other programs and scripts.  At the same time, it is a feature-rich network debugging and exploration tool, since it can create almost  any  kind of  connection  you would need and has several interesting built-in capabilities.  Netcat, or "nc" as the actual program is named, should have been supplied long ago as another  one of those cryptic but standard Unix tools.

## Usage

### PING

PING is one of the simplest commands. It's used to check whether a service is healthy. The response for the PING command is PONG.

```bash
$ nc -u 0.0.0.0 8080
PING
PONG
```

### Setting Values

```go
// Basic SET operation
SET key value

// SET with expiration (PX flag for milliseconds)
SET key value PX 10000  // Expires after 10 seconds
```

### Getting Values

```go
// GET operation
GET key  // Returns the value or (nil) if key doesn't exist or has expired
```

### Example

```go
SET mykey Hello World
GET mykey // Hello World

SET tempkey I will expire PX 5000  // Expires after 5 seconds
GET tempkey  // Returns: I will expire
// After 5 seconds...
GET tempkey  // Returns: (nil)
```

## Error Handling

The store returns error messages in the following cases:
- Wrong number of arguments
- Invalid expiration time format
- Invalid command syntax

## Thread Safety

All operations are thread-safe, implemented using sync.RWMutex for concurrent access.