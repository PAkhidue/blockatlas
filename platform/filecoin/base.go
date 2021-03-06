package filecoin

import (
	"github.com/trustwallet/blockatlas/internal"
	"github.com/trustwallet/blockatlas/platform/filecoin/explorer"
	"github.com/trustwallet/blockatlas/platform/filecoin/rpc"
	"github.com/trustwallet/golibs/coin"
)

type Platform struct {
	client   rpc.Client
	explorer explorer.Client
}

func Init(api, explorerApi string) *Platform {
	p := &Platform{
		client:   rpc.Client{Request: internal.InitClient(api)},
		explorer: explorer.Client{Request: internal.InitClient(explorerApi)},
	}
	return p
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.FIL]
}
