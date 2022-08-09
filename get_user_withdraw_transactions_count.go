package andromeda

import (
	"context"
	"fmt"
	"net/http"
)

type GetWithdrawTransactionsCountReq struct {
	UserID int64 `json:"user_id"`
	Count  int64 `json:"count"`
}

func (c *client) GetWithdrawTransactionsCount(ctx context.Context, userID int64) (GetWithdrawTransactionsCountReq, error) {
	req := GetWithdrawTransactionsCountReq{}
	err := c.call(ctx, http.MethodGet, c.makeURL(fmt.Sprintf("/user/%d/withdraw-transactions/count", userID)), nil, http.StatusOK, &req)
	return req, err
}
