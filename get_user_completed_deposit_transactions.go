package andromeda

import (
	"context"
	"fmt"
	"net/http"
)

func (c *client) GetUserCompletedDepositTransactions(ctx context.Context, userID, pageSize, page int64) ([]DepositTx, error) {
	resp := make([]DepositTx, 0)
	err := c.call(ctx, http.MethodGet, c.makeURL(fmt.Sprintf("/user/%d/deposit-transactions/%d/%d", userID, pageSize, page)), nil, http.StatusOK, &resp)
	return resp, err
}
