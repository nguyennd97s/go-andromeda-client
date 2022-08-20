package andromeda

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_client_GetUserPendingWithdrawTransactions(t *testing.T) {
	resp, err := testClient.GetUserPendingWithdrawTransactions(context.Background(), 1)
	if assert.Equal(t, err, nil) {
		t.Logf("%+v", resp)
	} else {
		t.Error(err)
	}
}
