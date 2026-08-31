package webserver

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"net"
	"net/http"

	"github.com/anacrolix/torrent"

	"github.com/oBusk/nullTorrent/internal/status"
)

func Serve(t *torrent.Torrent) error {
	dist, err := fs.Sub(distFS, "dist")
	if err != nil {
		return err
	}

	http.Handle("/", http.FileServer(http.FS(dist)))

	http.HandleFunc("/api/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(status.Of(t))
	})

	port := ":8080"
	ln, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	fmt.Println("Server started at http://localhost" + port)

	return http.Serve(ln, nil)
}
