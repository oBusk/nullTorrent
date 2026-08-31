# nullTorrent

An experimental BitTorrent client that stores files in a content-addressable store.

## What is this?

The idea is to have a BitTorrent client that natively stores all downloaded files in a Content-Addressable store (CAS). This should result in built-in deduplication, and easier integration for features like Cross Seeding.

This experiment will be built as a Go application which natively hosts an API and serves a web interface for managing.

## Development

To run the project, use the following command:

```bash
go run .
```

You can optionally run it with a memory-backed storage by using the `-memory` flag:

```bash
go run . -memory
```

You can also specify a custom download directory using the `-data-dir` flag:

```bash
go run . -data-dir custom_downloads
```

By default, the download directory is set to `./downloads`.

## License

MIT License
