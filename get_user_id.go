package andromeda

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"net/http"
)

type GetUserIDResp struct {
	UserID  int64  `json:"user_id"`
	Address string `json:"address"`
}

func (c *client) GetUserID(ctx context.Context, address common.Address) (GetUserIDResp, error) {
	resp := GetUserIDResp{}
	err := c.call(ctx, http.MethodGet, c.makeURL(fmt.Sprintf("/user/%s/id", address.Hex())), nil, http.StatusOK, &resp)
	return resp, err
}
