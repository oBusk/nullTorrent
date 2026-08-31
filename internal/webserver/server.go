package webserver

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/anacrolix/torrent"
)

func Serve(t *torrent.Torrent) error {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		stats := t.Stats()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(struct {
			BytesCompleted int64 `json:"bytesCompleted"`
			Length         int64 `json:"length"`
			ActivePeers    int   `json:"activePeers"`
			Seeders        int   `json:"seeders"`
		}{
			BytesCompleted: t.BytesCompleted(),
			Length:         t.Length(),
			ActivePeers:    stats.ActivePeers,
			Seeders:        stats.ConnectedSeeders,
		})
	})

	port := ":8080"
	ln, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	fmt.Println("Server started at http://localhost" + port)

	return http.Serve(ln, nil)
}
