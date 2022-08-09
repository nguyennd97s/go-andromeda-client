package andromeda

import (
	"context"
	"fmt"
	"net/http"
)

func (c *client) GetUserPendingWithdrawTransactions(ctx context.Context, userID int64) ([]WithdrawTx, error) {
	resp := make([]WithdrawTx, 0)
	err := c.call(ctx, http.MethodGet, c.makeURL(fmt.Sprintf("/user/%d/pending-withdraw-transactions", userID)), nil, http.StatusOK, &resp)
	return resp, err
}
