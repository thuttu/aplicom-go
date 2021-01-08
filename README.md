# Aplicom Go

[![PkgGoDev][pkg-badge]][pkg]
[![GoReportCard][report-badge]][report]
[![Codecov][codecov-badge]][codecov]

[pkg-badge]: https://pkg.go.dev/badge/go.einride.tech/aplicom
[pkg]: https://pkg.go.dev/go.einride.tech/aplicom
[report-badge]: https://goreportcard.com/badge/go.einride.tech/aplicom
[report]: https://goreportcard.com/report/go.einride.tech/aplicom
[codecov-badge]: https://codecov.io/gh/einride/aplicom-go/branch/master/graph/badge.svg
[codecov]: https://codecov.io/gh/einride/aplicom-go

Go SDK for [Aplicom][aplicom] telematics devices.

[aplicom]: https://www.aplicom.com

## Installing

```bash
$ go get go.einride.tech/aplicom
```

## Documentation

See the [Aplicom Extranet][aplicom-extranet] for device-specific and
protocol-specific documentation.

[aplicom-extranet]: https://www.aplicom.com/extranet/

## Examples

### Listening for D protocol packets

```go
package main

import (
	"fmt"
	"net"
	"time"

	"go.einride.tech/aplicom/dprotocol"
)

func main() {
	// Bind a TCP listener.
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err) // TODO: Handle error.
	}
	// Accept D protocol connections.
	for {
		conn, err := lis.Accept()
		if err != nil {
			panic(err) // TODO: Handle error.
		}
		// Scan D protocol packets.
		go func() {
			sc := dprotocol.NewScanner(conn)
			for sc.ScanPacket() {
				fmt.Printf(
					"Unit ID: %d Event ID: %d GPS Time: %s\n",
					sc.Packet().Header.UnitID,
					sc.Packet().EventID,
					sc.Packet().GPSTime.Format(time.RFC3339),
        )
			}
			if sc.Err() != nil {
				panic(err) // TODO: Handle error.
			}
		}()
	}
}
```
