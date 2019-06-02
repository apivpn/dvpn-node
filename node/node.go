package node

import (
	"fmt"
	"log"
	"net/http"
	"time"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/websocket"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/libs/common"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"

	"github.com/ironman0x7b2/vpn-node/database"
	"github.com/ironman0x7b2/vpn-node/tx"
	"github.com/ironman0x7b2/vpn-node/types"
)

type client struct {
	pubKey      crypto.PubKey
	conn        *websocket.Conn
	outMessages chan types.Msg
}

type Node struct {
	id      sdkTypes.ID
	address csdkTypes.AccAddress
	pubKey  crypto.PubKey

	tx      *tx.Tx
	db      *database.DB
	clients map[string]*client
	vpn     types.BaseVPN
}

func NewNode(id sdkTypes.ID, address csdkTypes.AccAddress, pubKey crypto.PubKey,
	tx *tx.Tx, db *database.DB, vpn types.BaseVPN) *Node {

	return &Node{
		id:      id,
		address: address,
		pubKey:  pubKey,

		tx:      tx,
		db:      db,
		clients: make(map[string]*client),
		vpn:     vpn,
	}
}

func (n Node) Start(apiPort uint16) {
	if err := n.vpn.Init(); err != nil {
		panic(err)
	}

	go func() {
		if err := n.vpn.Start(); err != nil {
			panic(err)
		}
	}()

	go func() {
		if err := n.updateNodeStatus(); err != nil {
			panic(err)
		}
	}()

	go func() {
		if err := n.updateAllSessionBandwidthsInfo(); err != nil {
			panic(err)
		}
	}()

	listenAddress := fmt.Sprintf("0.0.0.0:%d", apiPort)

	log.Printf("Listening the API server on address `%s`", listenAddress)
	if err := http.ListenAndServe(listenAddress, n.Router()); err != nil {
		panic(err)
	}
}

func (n Node) updateNodeStatus() error {
	log.Printf("Starting update node status ticker with interval `%s`",
		types.UpdateNodeStatusInterval.String())

	t := time.NewTicker(types.UpdateNodeStatusInterval)
	for ; ; <-t.C {
		msg := vpn.NewMsgUpdateNodeStatus(n.address, n.id, vpn.StatusActive)

		data, err := n.tx.CompleteAndSubscribeTx(msg)
		if err != nil {
			return err
		}

		log.Printf("Node status updated at block height `%d`, tx hash `%s`",
			data.Height, common.HexBytes(data.Tx.Hash()).String())
	}
}

func (n Node) updateAllSessionBandwidthsInfo() error {
	return nil
}

func (n Node) requestBandwidthSign(session *types.Session, bandwidth sdkTypes.Bandwidth) error {
	return nil
}