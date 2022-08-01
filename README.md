# near-api-go

### This library is forked from [eteu-technologies/near-api-go](https://github.com/eteu-technologies/near-api-go)

**WARNING**: This library is still work in progress. While it covers about 90% of the use-cases, it does not have many batteries included.

## Usage

```
go get github.com/mapprotocol/near-api-go
```

### Notes

What this library does not (and probably won't) provide:
- Access key caching & management
- Retry solution for `TransactionSendAwait`

What this library doesn't have yet:
- Response types for RPC queries marked as experimental by NEAR (prefixed with `EXPERIMENTAL_`)
- Few type definitions (especially complex ones, for example it's not very comfortable to reflect enum types in Go)

## Examples

See [cmd/](cmd/) in this repo for more fully featured CLI examples.

### Query latest block on NEAR testnet
```go
package main

import (
	"context"
	"fmt"

	"github.com/mapprotocol/near-api-go/pkg/client"
	"github.com/mapprotocol/near-api-go/pkg/client/block"
)

func main() {
	rpc, err := client.NewClient("https://rpc.testnet.near.org")
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	res, err := rpc.BlockDetails(ctx, block.FinalityFinal())
	if err != nil {
		panic(err)
	}

	fmt.Println("latest block: ", res.Header.Hash)
}
```

### Transfer 1 NEAR token between accounts

```go
package main

import (
	"context"
	"fmt"

	"github.com/mapprotocol/near-api-go/pkg/client"
	"github.com/mapprotocol/near-api-go/pkg/types"
	"github.com/mapprotocol/near-api-go/pkg/types/action"
	"github.com/mapprotocol/near-api-go/pkg/types/key"
)

var (
	sender    = "mikroskeem.testnet"
	recipient = "mikroskeem2.testnet"

	senderPrivateKey = `ed25519:...`
)

func main() {
	rpc, err := client.NewClient("https://rpc.testnet.near.org")
	if err != nil {
		panic(err)
	}

	keyPair, err := key.NewBase58KeyPair(senderPrivateKey)
	if err != nil {
		panic(err)
	}

	ctx := client.ContextWithKeyPair(context.Background(), keyPair)
	res, err := rpc.TransactionSendAwait(ctx, sender, recipient, []action.Action{
		action.NewTransfer(types.NEARToYocto(1)),
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("https://rpc.testnet.near.org/transactions/%s\n", res.Transaction.Hash)
}
```

