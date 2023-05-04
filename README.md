# BSC Block Size Monitor

This program retrieves the block size in bytes on EVM using node WebSocket URL. It subscribes to new block headers and fetches the full block for each new header. Then, it encodes the block and calculates the block size in bytes.

## Dependencies

- `bytes`: Provides functions to manipulate byte slices.
- `context`: Defines the `Context` package, which carries deadlines, cancellations, and other request-scoped values across API boundaries and between processes.
- `fmt`: Implements formatted I/O with functions analogous to C's printf and scanf.
- `log`: Implements a simple logging package for diagnostic messages.
- `github.com/ethereum/go-ethereum/core/types`: Contains data type definitions for blocks and transactions.
- `github.com/ethereum/go-ethereum/ethclient`: Provides an interface for interacting with an Ethereum or Binance Smart Chain node via JSON-RPC.
- `github.com/ethereum/go-ethereum/rpc`: Provides a client for interacting with a JSON-RPC server.

## Main Function

The `main()` function performs the following steps:

1. Connects to the BSC node via the WebSocket URL using `rpc.DialWebsocket()`.
2. Creates a new EthClient from the RPC client.
3. Subscribes to new block headers using `ethClient.SubscribeNewHead()`.
4. For each new block header, fetches the full block using `ethClient.BlockByHash()`.
5. Encodes the block and stores the data in a buffer.
6. Calculates the block size in bytes using `buf.Len()`.
7. Displays the block number and block size in bytes.

## Running the Program

To run this program, follow these steps:

1. Ensure you have Go installed on your computer. You can download it from [https://golang.org/dl/](https://golang.org/dl/).
2. Save the code in a file called `BlockSize.go`.
3. Open a terminal, navigate to the directory where the `BlockSize.go` file is located, and run the following command:go run BlockSize.go
4. The program will start displaying the block number and block size in bytes for each new block on Binance Smart Chain.

## Limitations and Improvements

This program has a few limitations and can be improved in several ways:

- It connects to a single BSC node, which may cause performance or availability issues. We may consider adding error handling and reconnection attempts to improve the program's robustness.
- The program only displays the block number and block size in bytes. We can add more information, such as block time, transactions included in the block, gas fees, etc., to get a more comprehensive view of each block.
