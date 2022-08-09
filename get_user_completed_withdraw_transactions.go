package andromeda

import (
	"context"
	"fmt"
	"net/http"
)

func (c *client) GetUserCompletedWithdrawTransactions(ctx context.Context, userID, pageSize, page int64) ([]WithdrawTx, error) {
	resp := make([]WithdrawTx, 0)
	err := c.call(ctx, http.MethodGet, c.makeURL(fmt.Sprintf("/user/%d/withdraw-transactions/%d/%d", userID, pageSize, page)), nil, http.StatusOK, &resp)
	return resp, err
}
