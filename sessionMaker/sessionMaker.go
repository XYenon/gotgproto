package sessionMaker

import (
	"context"

	"github.com/celestix/gotgproto/storage"
	"github.com/gotd/td/telegram"
)

func NewSessionStorage(ctx context.Context, sessionType SessionConstructor, inMemory bool) (*storage.PeerStorage, telegram.SessionStorage, error) {
	name, _, err := sessionType.loadSession()
	if err != nil {
		return nil, nil, err
	}
	sessDialect := name.(*sessionNameDialector)
	peerStorage := storage.NewPeerStorage(sessDialect.dialector, false)
	return peerStorage, &SessionStorage{
		data:        peerStorage.GetSession().Data,
		peerStorage: peerStorage,
	}, nil
}
