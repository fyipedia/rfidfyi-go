# rfidfyi-go

[![Go Reference](https://pkg.go.dev/badge/github.com/fyipedia/rfidfyi-go.svg)](https://pkg.go.dev/github.com/fyipedia/rfidfyi-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Go client for the [RFIDFYI](https://rfidfyi.com) REST API. RFID tags, frequency bands, standards. Zero external dependencies beyond stdlib.

> **Try the interactive tools at [rfidfyi.com](https://rfidfyi.com)** — explore, search, and discover.

## Install

```bash
go get github.com/fyipedia/rfidfyi-go
```

## Quick Start

```go
package main

import (
    "fmt"
    rfidfyi "github.com/fyipedia/rfidfyi-go"
)

func main() {
    client := rfidfyi.NewClient()

    // Search across all content
    result, err := client.Search("query")
    if err != nil {
        panic(err)
    }
    fmt.Println(result)
}
```

## API Methods

| Method | Description |
|--------|-------------|
| `AntennaTypes()` | List antenna types |
| `EpcSchemes()` | List epc schemes |
| `Faqs()` | List faqs |
| `FrequencyBands()` | List frequency bands |
| `Glossary()` | List glossary |
| `Guides()` | List guides |
| `Industries()` | List industries |
| `Manufacturers()` | List manufacturers |
| `Readers()` | List readers |
| `Standards()` | List standards |
| `TagFamilies()` | List tag families |
| `Tags()` | List tags |
| `UseCases()` | List use cases |
| `Search(query)` | Search across all content |

## Also Available

| Platform | Package | Link |
|----------|---------|------|
| **Python** | `pip install rfidfyi` | [PyPI](https://pypi.org/project/rfidfyi/) |
| **npm** | `npm install rfidfyi` | [npm](https://www.npmjs.com/package/rfidfyi) |
| **Go** | `go get github.com/fyipedia/rfidfyi-go` | [pkg.go.dev](https://pkg.go.dev/github.com/fyipedia/rfidfyi-go) |
| **Rust** | `cargo add rfidfyi` | [crates.io](https://crates.io/crates/rfidfyi) |
| **Ruby** | `gem install rfidfyi` | [rubygems](https://rubygems.org/gems/rfidfyi) |

## Tag FYI Family

Part of the [FYIPedia](https://fyipedia.com) open-source developer tools ecosystem.

| Site | Domain | Focus |
|------|--------|-------|
| BarcodeFYI | [barcodefyi.com](https://barcodefyi.com) | Barcode formats, EAN, UPC, ISBN standards |
| BLEFYI | [blefyi.com](https://blefyi.com) | Bluetooth Low Energy, GATT, beacons |
| NFCFYI | [nfcfyi.com](https://nfcfyi.com) | NFC chips, tag types, NDEF, standards |
| QRCodeFYI | [qrcodefyi.com](https://qrcodefyi.com) | QR code types, versions, encoding modes |
| **RFIDFYI** | [rfidfyi.com](https://rfidfyi.com) | RFID tags, frequency bands, standards |
| SmartCardFYI | [smartcardfyi.com](https://smartcardfyi.com) | Smart cards, EMV, APDU, Java Card |

## Embed Widget

Embed [RFIDFYI](https://rfidfyi.com) widgets on any website with [rfidfyi-embed](https://widget.rfidfyi.com):

```html
<script src="https://cdn.jsdelivr.net/npm/rfidfyi-embed@1/dist/embed.min.js"></script>
<div data-rfidfyi="entity" data-slug="example"></div>
```

Zero dependencies · Shadow DOM · 4 themes (light/dark/sepia/auto) · [Widget docs](https://widget.rfidfyi.com)

## License

MIT
