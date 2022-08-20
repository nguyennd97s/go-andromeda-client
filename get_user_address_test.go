package andromeda

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_client_GetUserAddress(t *testing.T) {
	resp, err := testClient.GetUserAddress(context.Background(), 1)
	if assert.Equal(t, err, nil) {
		t.Logf("address: %s, user id: %d", resp.Address, resp.UserID)
	} else {
		t.Error(err)
	}
}
