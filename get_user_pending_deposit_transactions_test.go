package andromeda

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_client_GetUserPendingDepositTransactions(t *testing.T) {
	resp, err := testClient.GetUserPendingDepositTransactions(context.Background(), 1)
	if assert.Equal(t, err, nil) {
		t.Logf("%+v", resp)
	}
}
