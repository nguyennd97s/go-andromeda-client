package andromeda

import (
	"context"
	"net/http"
)

type WithdrawReq struct {
	UserID    int64   `json:"user_id"`
	Address   string  `json:"address"`
	TokenName string  `json:"token_name"`
	Amount    float64 `json:"amount"`
	RequestID string  `json:"request_id"`
	Checksum  string  `json:"checksum"`
}

func (c *client) Withdraw(ctx context.Context, req WithdrawReq) error {
	return c.call(ctx, http.MethodPost, "/withdraw", req, http.StatusCreated, nil)
}
