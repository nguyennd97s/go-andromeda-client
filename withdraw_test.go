package andromeda

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_client_Withdraw(t *testing.T) {
	err := testClient.Withdraw(context.Background(), WithdrawReq{
		UserID:    1,
		Address:   "0x4eb36b083a57F054eEdFa4EBE258B48af936A03B",
		TokenName: "usdt",
		Amount:    0.001,
		RequestID: "a",
		Checksum:  "2435b7d643b0a78a987a74e0ca1b9633",
	})
	if !assert.Equal(t, err, nil) {
		t.Error(err)
	}
}
