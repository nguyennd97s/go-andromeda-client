package andromeda

import (
	"context"
	"fmt"
	"net/http"
)

func (c *client) GetUserPendingDepositTransactions(ctx context.Context, userID int64) ([]DepositTx, error) {
	resp := make([]DepositTx, 0)
	err := c.call(ctx, http.MethodGet, c.makeURL(fmt.Sprintf("/user/%d/pending-deposit-transactions", userID)), nil, http.StatusOK, &resp)
	return resp, err
}
