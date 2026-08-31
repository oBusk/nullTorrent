package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/anacrolix/torrent"
)

func printProgress(t *torrent.Torrent) {
	lastStats := t.Stats()
	for {
		time.Sleep(time.Second)

		stats := t.Stats()
		rate := stats.BytesReadUsefulData.Int64() - lastStats.BytesReadUsefulData.Int64()
		lastStats = stats

		fmt.Println(t.BytesCompleted(), "/", t.Length(), "-", rate, "B/s -", stats.ActivePeers, "peers,", stats.ConnectedSeeders, "seeders")
	}
}

func serveStatus(t *torrent.Torrent) {

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
	go http.ListenAndServe(port, nil)
	fmt.Println("Server started at http://localhost" + port)
}

func main() {
	cfg := torrent.NewDefaultClientConfig()
	cfg.DataDir = "./downloads"

	client, _ := torrent.NewClient(cfg)
	defer client.Close()

	// t, _ := client.AddTorrentFromFile("bittorrent-v2-test.torrent")
	t, _ := client.AddTorrentFromFile("ReactOS-0.4.15-release-1-gdbb43bbaeb2-x86-iso.zip-hybrid.torrent")

	<-t.GotInfo()

	t.DownloadAll()

	// for !t.Complete().Bool() {
	// 	fmt.Println(t.BytesCompleted(), "/", t.Length())
	// 	time.Sleep(1 * time.Second)
	// }

	go printProgress(t)
	go serveStatus(t)

	select {}
}
