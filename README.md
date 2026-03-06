# rfidfyi-go

[![Go Reference](https://pkg.go.dev/badge/github.com/fyipedia/rfidfyi-go.svg)](https://pkg.go.dev/github.com/fyipedia/rfidfyi-go)

Go client for the [RFIDFYI](https://rfidfyi.com) API. Look up RFID tags, readers, families, frequencies, standards, EPC schemes, and use cases. Zero dependencies — stdlib only.

## Install

```bash
go get github.com/fyipedia/rfidfyi-go
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"

    rfidfyi "github.com/fyipedia/rfidfyi-go"
)

func main() {
    client := rfidfyi.NewClient()

    result, err := client.Search("uhf")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Found %d results for 'uhf'\n", result.Total)

    tag, err := client.Tag("epc-gen2")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s: %s\n", tag.Name, tag.Description)
}
```

## API Methods

| Method | Description |
|--------|-------------|
| `Search(query)` | Search tags, standards, and glossary |
| `Tag(slug)` | Get RFID tag details |
| `Reader(slug)` | Get RFID reader details |
| `Family(slug)` | Get RFID family details |
| `Frequency(slug)` | Get frequency band details |
| `Standard(slug)` | Get standard details |
| `EPC(slug)` | Get EPC scheme details |
| `UseCase(slug)` | Get use case details |
| `GlossaryTerm(slug)` | Get glossary term definition |
| `Compare(slugA, slugB)` | Compare two RFID tags |
| `Random()` | Get a random RFID tag |

## REST API

Free, no authentication required. Base URL: `https://rfidfyi.com/api`

```bash
curl https://rfidfyi.com/api/search/?q=uhf
curl https://rfidfyi.com/api/tag/epc-gen2/
curl https://rfidfyi.com/api/random/
```

Full OpenAPI spec: https://rfidfyi.com/api/openapi.json

## Also Available

| Language | Package | Install |
|----------|---------|---------|
| Python | [rfidfyi](https://pypi.org/project/rfidfyi/) | `pip install rfidfyi` |
| TypeScript | [rfidfyi](https://www.npmjs.com/package/rfidfyi) | `npm install rfidfyi` |
| Go | [rfidfyi-go](https://pkg.go.dev/github.com/fyipedia/rfidfyi-go) | `go get github.com/fyipedia/rfidfyi-go` |
| Rust | [rfidfyi](https://crates.io/crates/rfidfyi) | `cargo add rfidfyi` |
| Ruby | [rfidfyi](https://rubygems.org/gems/rfidfyi) | `gem install rfidfyi` |

## Code FYI Family

| Site | Domain | Focus |
|------|--------|-------|
| BarcodeFYI | [barcodefyi.com](https://barcodefyi.com) | Barcode symbologies and standards |
| QRCodeFYI | [qrcodefyi.com](https://qrcodefyi.com) | QR code types and encoding |
| NFCFYI | [nfcfyi.com](https://nfcfyi.com) | NFC chips and standards |
| BLEFYI | [blefyi.com](https://blefyi.com) | Bluetooth Low Energy profiles |
| RFIDFYI | [rfidfyi.com](https://rfidfyi.com) | RFID tags and frequencies |
| SmartCardFYI | [smartcardfyi.com](https://smartcardfyi.com) | Smart card platforms and standards |

## License

MIT
