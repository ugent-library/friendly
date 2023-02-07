[![Go Reference](https://pkg.go.dev/badge/github.com/ugent-library/friendly.svg)](https://pkg.go.dev/github.com/ugent-library/friendly)

# ugent-library/friendly

Package friendly contains human friendly formatters for units.

## Install

```sh
go get -u github.com/ugent-library/friendly
```

## Examples

```go
friendly.Bytes(2_000_000_000)
// "2 GB"
```