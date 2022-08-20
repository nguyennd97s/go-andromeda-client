package andromeda

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_client_CheckDepositTx(t *testing.T) {
	err := testClient.CheckDepositTx(context.Background(), CheckDepositTransactionReq{
		UserID:          1,
		TransactionHash: "0xccf29a252e8bf3caaacddf47ffac2422ab0cccd2135972e51186d00f8339a1d5",
	})

	if !assert.Equal(t, err, nil) {
		t.Error(err)
	}
}
