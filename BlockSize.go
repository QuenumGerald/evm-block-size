package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	rpcClient, err := rpc.DialWebsocket(context.Background(), "wss:// ", "")
	if err != nil {
		log.Fatal("Error connecting to WebSocket: ", err)
	}
	defer rpcClient.Close()

	ethClient := ethclient.NewClient(rpcClient)

	headers := make(chan *types.Header)
	sub, err := ethClient.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal("Error subscribing to new block headers: ", err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case header := <-headers:
			block, err := ethClient.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Println("Error retrieving block: ", err)
				continue
			}

			var buf bytes.Buffer
			if err := block.EncodeRLP(&buf); err != nil {
				log.Println("Error encoding block RLP: ", err)
				continue
			}

			blockSize := buf.Len()
			fmt.Printf("Block number: %d, Block size: %d bytes\n", block.NumberU64(), blockSize)
		}
	}
}
