package sui

import (
	"github.com/certusone/wormhole/node/pkg/common"
	gossipv1 "github.com/certusone/wormhole/node/pkg/proto/gossip/v1"
	"github.com/certusone/wormhole/node/pkg/query"
	"github.com/certusone/wormhole/node/pkg/supervisor"
	"github.com/certusone/wormhole/node/pkg/watchers"
	"github.com/certusone/wormhole/node/pkg/watchers/interfaces"
	"github.com/wormhole-foundation/wormhole/sdk/vaa"
)

type WatcherConfig struct {
	NetworkID        watchers.NetworkID // human readable name
	ChainID          vaa.ChainID        // ChainID
	Rpc              string
	SuiMoveEventType string
}

func (wc *WatcherConfig) GetNetworkID() watchers.NetworkID {
	return wc.NetworkID
}

func (wc *WatcherConfig) GetChainID() vaa.ChainID {
	return wc.ChainID
}

func (wc *WatcherConfig) RequiredL1Finalizer() watchers.NetworkID {
	return ""
}

func (wc *WatcherConfig) SetL1Finalizer(l1finalizer interfaces.L1Finalizer) {
	// empty
}

func (wc *WatcherConfig) Create(
	msgC chan<- *common.MessagePublication,
	obsvReqC <-chan *gossipv1.ObservationRequest,
	_ <-chan *query.PerChainQueryInternal,
	_ chan<- *query.PerChainQueryResponseInternal,
	_ chan<- *common.GuardianSet,
	env common.Environment,
) (interfaces.L1Finalizer, supervisor.Runnable, error) {
	var devMode bool = (env == common.UnsafeDevNet)

	return nil, NewWatcher(wc.Rpc, wc.SuiMoveEventType, devMode, msgC, obsvReqC).Run, nil
}
