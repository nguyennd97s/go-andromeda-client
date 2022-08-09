package andromeda

import (
	"context"
	"net/http"
)

type UserCountReq struct {
	Count int64 `json:"count"`
}

func (c *client) UserCount(ctx context.Context) (UserCountReq, error) {
	resp := UserCountReq{}
	err := c.call(ctx, http.MethodGet, c.makeURL("/user/count"), nil, http.StatusOK, &resp)
	return resp, err
}
