package andromeda

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_client_GetUserCompletedDepositTransactions(t *testing.T) {
	resp, err := testClient.GetUserCompletedDepositTransactions(context.Background(), 1, 1, 1)
	if assert.Equal(t, err, nil) {
		t.Logf("%+v", resp)
	}
}
