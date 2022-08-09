package andromeda

import (
	"context"
	"net/http"
)

func (c *client) Ping(ctx context.Context) error {
	return c.call(ctx, http.MethodGet, c.makeURL("/ping"), nil, http.StatusOK, nil)
}
