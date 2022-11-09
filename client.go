package andromeda

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"io"
	"log"
	"net/http"
	"reflect"
	"time"
)

type Client interface {
	// Ping ping ping ping
	Ping(ctx context.Context) error
	// Withdraw implement https://github.com/oriS-casino/andromeda-bep20-payment-gateway-document#withdraw
	Withdraw(ctx context.Context, req WithdrawReq) error
	// UserCount implement https://github.com/oriS-casino/andromeda-bep20-payment-gateway-document#user-count
	UserCount(ctx context.Context) (UserCountReq, error)
	// GetUserAddress implement https://github.com/oriS-casino/andromeda-bep20-payment-gateway-document#get-user-address
	GetUserAddress(ctx context.Context, userID int64) (GetUserAddressResp, error)
	// GetUserID implement https://github.com/oriS-casino/andromeda-bep20-payment-gateway-document#get-user-id
	GetUserID(ctx context.Context, address common.Address) (GetUserIDResp, error)
	// GetUserCompletedDepositTransactions implement https://github.com/oriS-casino/andromeda-bep20-payment-gateway-document#get-user-completed-deposit-transactions
	GetUserCompletedDepositTransactions(ctx context.Context, userID, pageSize, page int64) ([]DepositTx, error)
	// GetUserCompletedWithdrawTransactions implement https://github.com/oriS-casino/andromeda-bep20-payment-gateway-document#get-user-completed-withdraw-transactions
	GetUserCompletedWithdrawTransactions(ctx context.Context, userID, pageSize, page int64) ([]WithdrawTx, error)
	// GetDepositTransactionsCount implement https://github.com/oriS-casino/andromeda-bep20-payment-gateway-document#get-user-deposit-transactions-count
	GetDepositTransactionsCount(ctx context.Context, userID int64) (GetDepositTransactionsCountReq, error)
	// GetWithdrawTransactionsCount implement https://github.com/oriS-casino/andromeda-bep20-payment-gateway-document#get-user-withdraw-transactions-count
	GetWithdrawTransactionsCount(ctx context.Context, userID int64) (GetWithdrawTransactionsCountReq, error)
	// GetUserPendingDepositTransactions implement https://github.com/oriS-casino/andromeda-bep20-payment-gateway-document#get-user-pending-deposit-transactions
	GetUserPendingDepositTransactions(ctx context.Context, userID int64) ([]DepositTx, error)
	// GetUserPendingWithdrawTransactions implement https://github.com/oriS-casino/andromeda-bep20-payment-gateway-document#get-user-pending-withdraw-transactions
	GetUserPendingWithdrawTransactions(ctx context.Context, userID int64) ([]WithdrawTx, error)
	// CheckDepositTx implement https://github.com/oriS-casino/andromeda-bep20-payment-gateway-document#check-deposit-transaction
	CheckDepositTx(ctx context.Context, req CheckDepositTransactionReq) error
}

type client struct {
	client   http.Client
	password string
	baseURL  string
}

func NewClient(baseURL, password string) Client {
	return &client{
		client: http.Client{
			Timeout: time.Second * 30,
		},
		password: password,
		baseURL:  baseURL,
	}
}

// result must be pointer
func (c *client) call(ctx context.Context, method, path string, body interface{}, expectStatus int, result interface{}) error {
	log.Println(fmt.Sprintf("calling api %s %s", method, path))
	var payload *bytes.Buffer

	if !isNil(body) {
		b, err := json.Marshal(body)
		if err != nil {
			return errors.Wrap(err, "marshal transfer token req")
		}
		payload = bytes.NewBuffer(b)
	} else {
		payload = bytes.NewBuffer(nil)
	}

	r, err := http.NewRequestWithContext(ctx, method, path, payload)
	if err != nil {
		return errors.Wrap(err, "create request")
	}
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", "Password "+c.password)

	resp, err := c.client.Do(r)
	if err != nil {
		return errors.Wrap(err, "call api")
	}

	if resp.StatusCode != expectStatus {
		msg, _ := io.ReadAll(resp.Body)
		return errors.Errorf("invalid status code %d (msg: %s)", resp.StatusCode, string(msg))
	}

	if !isNil(result) {
		if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
			return errors.Wrap(err, "decode response")
		}
	}
	return nil
}

func (c *client) makeURL(path string) string {
	return c.baseURL + "/api/v1" + path
}

func isNil(i interface{}) bool {
	if i == nil || (reflect.ValueOf(i).Kind() == reflect.Ptr && reflect.ValueOf(i).IsNil()) {
		return true
	}
	return false
}
