package andromeda

import (
	"context"
	"fmt"
	"net/http"
)

type GetUserAddressResp struct {
	UserID  int64  `json:"user_id"`
	Address string `json:"address"`
}

func (c *client) GetUserAddress(ctx context.Context, userID int64) (GetUserAddressResp, error) {
	resp := GetUserAddressResp{}
	err := c.call(ctx, http.MethodGet, c.makeURL(fmt.Sprintf("/user/%d/address", userID)), nil, http.StatusOK, &resp)
	return resp, err
}
