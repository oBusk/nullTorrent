package main

import (
	"context"
	"io"

	"github.com/anacrolix/torrent/metainfo"
	"github.com/anacrolix/torrent/storage"
)

type memPiece struct {
	data     []byte
	complete bool
}

func (p *memPiece) ReadAt(b []byte, off int64) (int, error) {
	if off >= int64(len(p.data)) {
		return 0, io.EOF
	}
	n := copy(b, p.data[off:])
	if n < len(b) {
		return n, io.EOF
	}
	return n, nil
}

func (p *memPiece) WriteAt(b []byte, off int64) (int, error) {
	return copy(p.data[off:], b), nil
}

func (p *memPiece) MarkComplete() error {
	p.complete = true
	return nil
}

func (p *memPiece) MarkNotComplete() error {
	p.complete = false
	return nil
}

func (p *memPiece) Completion() storage.Completion {
	return storage.Completion{Ok: true, Complete: p.complete}
}

type memStorage struct{}

func newMemoryStorage() storage.ClientImpl {
	return memStorage{}
}

func (memStorage) OpenTorrent(ctx context.Context, info *metainfo.Info, infoHash metainfo.Hash) (storage.TorrentImpl, error) {
	pieces := make([]*memPiece, info.NumPieces())
	for i := range pieces {
		pieces[i] = &memPiece{data: make([]byte, info.Piece(i).Length())}
	}

	return storage.TorrentImpl{
		Piece: func(p metainfo.Piece) storage.PieceImpl {
			return pieces[p.Index()]
		},
		Close: func() error { return nil },
	}, nil
}
