package iface

import (
	"context"
	path "github.com/ipweb-group/interface-go-ipws-core/path"

	"github.com/ipweb-group/interface-go-ipws-core/options"

	"github.com/libp2p/go-libp2p-peer"
	pstore "github.com/libp2p/go-libp2p-peerstore"
)

// DhtAPI specifies the interface to the DHT
// Note: This API will likely get deprecated in near future, see
// https://github.com/ipfs/interface-ipfs-core/issues/249 for more context.
type DhtAPI interface {
	// FindPeer queries the DHT for all of the multiaddresses associated with a
	// Peer ID
	FindPeer(context.Context, peer.ID) (pstore.PeerInfo, error)

	// FindProviders finds peers in the DHT who can provide a specific value
	// given a key.
	FindProviders(context.Context, path.Path, ...options.DhtFindProvidersOption) (<-chan pstore.PeerInfo, error)

	// Provide announces to the network that you are providing given values
	Provide(context.Context, path.Path, ...options.DhtProvideOption) error
}
