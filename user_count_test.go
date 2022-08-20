package andromeda

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_client_UserCount(t *testing.T) {
	resp, err := testClient.UserCount(context.Background())
	if assert.Equal(t, err, nil) {
		t.Logf("%+v", resp)
	} else {
		t.Error(err)
	}
}
