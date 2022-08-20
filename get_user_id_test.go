package andromeda

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_client_GetUserID(t *testing.T) {
	resp, err := testClient.GetUserAddress(context.Background(), 1)
	if assert.Equal(t, err, nil) {
		t.Logf("address: %s, user id: %d", resp.Address, resp.UserID)
	} else {
		t.Error(err)
	}

	getUserIDResp, err := testClient.GetUserID(context.Background(), common.HexToAddress(resp.Address))
	if assert.Equal(t, err, nil) {
		assert.Equal(t, getUserIDResp.UserID, resp.UserID)
		assert.Equal(t, getUserIDResp.Address, resp.Address)
	} else {
		t.Error(err)
	}
}
