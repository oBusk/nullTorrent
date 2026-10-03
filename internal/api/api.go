// Package api defines the types sent over the HTTP API. They are the public
// contract with the web interface and other clients, and are converted to
// TypeScript for the SDK (see generate.go), so only add fields here that
// belong on the wire.
package api

type Status struct {
	BytesCompleted int64 `json:"bytesCompleted"`
	Length         int64 `json:"length"`
	ActivePeers    int   `json:"activePeers"`
	Seeders        int   `json:"seeders"`
}
