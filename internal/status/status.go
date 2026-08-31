package status

import "github.com/anacrolix/torrent"

type Status struct {
	BytesCompleted int64 `json:"bytesCompleted"`
	Length         int64 `json:"length"`
	ActivePeers    int   `json:"activePeers"`
	Seeders        int   `json:"seeders"`
	BytesRead      int64 `json:"-"`
}

func Of(t *torrent.Torrent) Status {
	stats := t.Stats()
	return Status{
		BytesCompleted: t.BytesCompleted(),
		Length:         t.Length(),
		ActivePeers:    stats.ActivePeers,
		Seeders:        stats.ConnectedSeeders,
		BytesRead:      stats.BytesReadUsefulData.Int64(),
	}
}
