# example-go

[![coverage badge](https://img.shields.io/endpoint?url=https%3A%2F%2Fdemo.coveragetracker.dev%2Fapi%2Fbadge%2FCoverageTracker%2Fexample-go%2Fcoverage.json)](https://demo.coveragetracker.dev/CoverageTracker/example-go?metric=coverage)
[![complexity badge](https://img.shields.io/endpoint?url=https%3A%2F%2Fdemo.coveragetracker.dev%2Fapi%2Fbadge%2FCoverageTracker%2Fexample-go%2Fcomplexity.json)](https://demo.coveragetracker.dev/CoverageTracker/example-go?metric=complexity)

A small, idiomatic Go order-pricing library used as the Go reference example
for [Coverage Tracker](https://coveragetracker.dev). It exists to give the Go
row in the
[coverage report generation guide](https://coveragetracker.dev/docs/generating-coverage-reports)
a live, working reference, and to populate the
[demo dashboard](https://demo.coveragetracker.dev) with real trend data.

**This is a demo/marketing repo, not a test suite for Coverage Tracker
itself.** Go's native coverage profile doesn't carry branch data, so unlike
the other example repos this one won't show a branch coverage metric on the
dashboard.

## What's here

- `inventory/` — shipping cost, volume discount, and SKU validation logic for
  a small mail-order warehouse, each with real branching logic.
- `cmd/example-go/` — a thin CLI entry point that prices a sample order using
  the `inventory` package.
- `.github/workflows/coverage.yml` — runs `go test` with a coverage profile,
  generates a [gocyclo](https://github.com/fzipp/gocyclo) complexity report,
  then reports both to the demo instance via the `coverage-tracker` reporting
  Action.

## Running locally

```sh
go test -coverprofile=coverage.out ./...   # writes coverage.out
go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
gocyclo -avg -ignore '_test\.go' .
```
