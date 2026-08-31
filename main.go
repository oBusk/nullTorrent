package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/oBusk/nullTorrent/internal/memstorage"
	"github.com/oBusk/nullTorrent/internal/webserver"

	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/storage"
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

func main() {
	inMemory := flag.Bool("memory", false, "keep downloaded data in memory instead of on disk")
	dataDir := flag.String("data-dir", "downloads", "directory to write downloaded data to")
	flag.Parse()

	cfg := torrent.NewDefaultClientConfig()
	if *inMemory {
		cfg.DefaultStorage = memstorage.New()
	} else {
		cfg.DataDir = *dataDir
		cfg.DefaultStorage = storage.NewFile(*dataDir)
	}

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
	go webserver.Serve(t)

	select {}
}
