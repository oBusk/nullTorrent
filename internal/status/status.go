package status

import (
	"github.com/anacrolix/torrent"

	"github.com/oBusk/nullTorrent/internal/api"
)

// Status is the API status plus fields that are only used internally.
type Status struct {
	api.Status
	BytesRead int64 `json:"-"`
}

func Of(t *torrent.Torrent) Status {
	stats := t.Stats()
	return Status{
		Status: api.Status{
			BytesCompleted: t.BytesCompleted(),
			Length:         t.Length(),
			ActivePeers:    stats.ActivePeers,
			Seeders:        stats.ConnectedSeeders,
		},
		BytesRead: stats.BytesReadUsefulData.Int64(),
	}
}
