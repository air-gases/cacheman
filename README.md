# Cacheman

[![PkgGoDev](https://pkg.go.dev/badge/github.com/air-gases/cacheman)](https://pkg.go.dev/github.com/air-gases/cacheman)

A useful gas that used to manage the Cache-Control header for the web
applications built using [Air](https://github.com/aofei/air).

## Installation

Open your terminal and execute

```bash
$ go get github.com/air-gases/cacheman
```

done.

> The only requirement is the [Go](https://golang.org), at least v1.13.

## Usage

Create a file named `main.go`

```go
package main

import (
	"github.com/air-gases/cacheman"
	"github.com/aofei/air"
)

func main() {
	a := air.Default
	a.DebugMode = true
	a.GET("/", func(req *air.Request, res *air.Response) error {
		return res.WriteString("This message will last for an hour.")
	}, cacheman.Gas(cacheman.GasConfig{
		MaxAge:  3600,
		SMaxAge: -1,
	}))
	a.Serve()
}
```

and run it

```bash
$ go run main.go
```

then visit `http://localhost:8080`.

## Community

If you want to discuss Cacheman, or ask questions about it, simply post
questions or ideas [here](https://github.com/air-gases/cacheman/issues).

## Contributing

If you want to help build Cacheman, simply follow
[this](https://github.com/air-gases/cacheman/wiki/Contributing) to send pull
requests [here](https://github.com/air-gases/cacheman/pulls).

## License

This project is licensed under the MIT License.

License can be found [here](LICENSE).
