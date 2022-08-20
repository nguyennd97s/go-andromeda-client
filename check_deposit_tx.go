package andromeda

import (
	"context"
	"net/http"
)

type CheckDepositTransactionReq struct {
	UserID          int64  `json:"user_id"`
	TransactionHash string `json:"transaction_hash"`
}

func (c *client) CheckDepositTx(ctx context.Context, req CheckDepositTransactionReq) error {
	return c.call(ctx, http.MethodPost, c.makeURL("/check-deposit-tx"), req, http.StatusAccepted, nil)
}
