package andromeda

import (
	"context"
	"fmt"
	"net/http"
)

type GetDepositTransactionsCountReq struct {
	UserID int64 `json:"user_id"`
	Count  int64 `json:"count"`
}

func (c *client) GetDepositTransactionsCount(ctx context.Context, userID int64) (GetDepositTransactionsCountReq, error) {
	req := GetDepositTransactionsCountReq{}
	err := c.call(ctx, http.MethodGet, c.makeURL(fmt.Sprintf("/user/%d/deposit-transactions/count", userID)), nil, http.StatusOK, &req)
	return req, err
}
