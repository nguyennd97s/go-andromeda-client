package andromeda

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_client_GetWithdrawTransactionsCount(t *testing.T) {
	resp, err := testClient.GetWithdrawTransactionsCount(context.Background(), 1)
	if assert.Equal(t, err, nil) {
		t.Logf("%+v", resp)
	}
}
