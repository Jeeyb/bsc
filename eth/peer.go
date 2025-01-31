// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package eth

import (
	"math/big"
	"time"

	"github.com/Jeeyb/bsc/eth/protocols/trust"

	"github.com/Jeeyb/bsc/eth/protocols/diff"
	"github.com/Jeeyb/bsc/eth/protocols/eth"
	"github.com/Jeeyb/bsc/eth/protocols/snap"
)

// ethPeerInfo represents a short summary of the `eth` sub-protocol metadata known
// about a connected peer.
type ethPeerInfo struct {
	Version    uint     `json:"version"`    // Ethereum protocol version negotiated
	Difficulty *big.Int `json:"difficulty"` // Total difficulty of the peer's blockchain
	Head       string   `json:"head"`       // Hex hash of the peer's best owned block
}

// ethPeer is a wrapper around eth.Peer to maintain a few extra metadata.
type ethPeer struct {
	*eth.Peer
	snapExt  *snapPeer // Satellite `snap` connection
	diffExt  *diffPeer
	trustExt *trustPeer

	syncDrop *time.Timer   // Connection dropper if `eth` sync progress isn't validated in time
	snapWait chan struct{} // Notification channel for snap connections
}

// info gathers and returns some `eth` protocol metadata known about a peer.
func (p *ethPeer) info() *ethPeerInfo {
	hash, td := p.Head()

	return &ethPeerInfo{
		Version:    p.Version(),
		Difficulty: td,
		Head:       hash.Hex(),
	}
}

// snapPeerInfo represents a short summary of the `snap` sub-protocol metadata known
// about a connected peer.
type snapPeerInfo struct {
	Version uint `json:"version"` // Snapshot protocol version negotiated
}

// diffPeerInfo represents a short summary of the `diff` sub-protocol metadata known
// about a connected peer.
type diffPeerInfo struct {
	Version  uint `json:"version"` // diff protocol version negotiated
	DiffSync bool `json:"diff_sync"`
}

// trustPeerInfo represents a short summary of the `trust` sub-protocol metadata known
// about a connected peer.
type trustPeerInfo struct {
	Version uint `json:"version"` // Trust protocol version negotiated
}

// snapPeer is a wrapper around snap.Peer to maintain a few extra metadata.
type snapPeer struct {
	*snap.Peer
}

// diffPeer is a wrapper around diff.Peer to maintain a few extra metadata.
type diffPeer struct {
	*diff.Peer
}

// trustPeer is a wrapper around trust.Peer to maintain a few extra metadata.
type trustPeer struct {
	*trust.Peer
}

// info gathers and returns some `diff` protocol metadata known about a peer.
func (p *diffPeer) info() *diffPeerInfo {
	return &diffPeerInfo{
		Version:  p.Version(),
		DiffSync: p.DiffSync(),
	}
}

// info gathers and returns some `snap` protocol metadata known about a peer.
func (p *snapPeer) info() *snapPeerInfo {
	return &snapPeerInfo{
		Version: p.Version(),
	}
}

// info gathers and returns some `trust` protocol metadata known about a peer.
func (p *trustPeer) info() *trustPeerInfo {
	return &trustPeerInfo{
		Version: p.Version(),
	}
}
