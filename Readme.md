# Andromeda Client 

Thư viện client cho Andromeda Service được xây dựng dựa trên Document: https://github.com/dangnguyendota-casino/andromeda-bep20-payment-gateway-document


Example:

```go
package main

import (
	"context"
	andromeda "github.com/dangnguyendota-casino/go-andromeda-client"
)

func main() {
	client := andromeda.NewClient("http://localhost:9711", "example-password")

	client.GetUserAddress(context.Background(), 1)
}
```