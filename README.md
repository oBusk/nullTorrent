# nullTorrent

An experimental BitTorrent client that stores files in a content-addressable store.

## What is this?

The idea is to have a BitTorrent client that natively stores all downloaded files in a Content-Addressable store (CAS). This should result in built-in deduplication, and easier integration for features like Cross Seeding.

This experiment will be built as a Go application which natively hosts an API and serves a web interface for managing.

## Requirements

- Go 1.26 or later
- Node.js and pnpm

Node is not optional. The web interface is compiled into the Go binary with
`go:embed`, and the built assets are not checked in, so the frontend must be
built at least once before the Go code will compile.

## Getting started

```bash
pnpm install
pnpm build
```

`pnpm build` builds the web interface and then the Go binary, in that order.

If you skip it and run `go run .` on a fresh clone, the build fails with:

```
internal/webserver/assets.go:5:12: pattern all:dist: no matching files found
```

The fix is `pnpm --dir webui build`.

## Development

```bash
pnpm dev
```

This starts two processes side by side:

- Vite on <http://localhost:5173>, which opens in your browser automatically
- The Go server on <http://localhost:8080>

Do your work against **:5173**. Vite serves `webui/index.html` straight from
source with automatic reloading, and proxies `/api` through to the Go server, so
frontend changes appear without rebuilding or restarting anything.

Port 8080 serves the version of the interface that was compiled into the binary.
That copy only changes when you rebuild, so use :8080 to verify a real build, not
to iterate.

### Running the Go server on its own

```bash
go run .
```

Serves the API and the embedded web interface on <http://localhost:8080>. Run
`pnpm --dir webui build` first if you have changed anything under `webui/`,
otherwise the previously built interface is served.

### Options

Keep downloaded data in memory instead of writing it to disk:

```bash
go run . -memory
```

Write downloaded data to a different directory (default `./downloads`):

```bash
go run . -data-dir custom_downloads
```

## Building a binary

```bash
pnpm build     # web interface, then ./nullTorrent
pnpm start     # the same, then runs the binary
```

## Project layout

```
main.go                 flag parsing and wiring
internal/
  memstorage/           in-memory torrent storage backend
  status/               torrent status shared by the API and the console output
  webserver/            HTTP API and web interface
    dist/               built web interface, embedded at compile time (generated)
webui/                  web interface source, built with Vite
```

## License

MIT License
