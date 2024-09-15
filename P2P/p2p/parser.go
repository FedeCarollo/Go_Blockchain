package p2p

import (
	"encoding/json"
	"p2p_network/p2p/btc_blockchain/blockchain"
)

func ParseJsonToBlockchain(data []byte) (*blockchain.Blockchain, error) {
	var blockchain blockchain.Blockchain

	err := json.Unmarshal([]byte(data), &blockchain)

	if err != nil {
		return nil, err
	}

	return &blockchain, nil

}
