# Andromeda Client 

Thư viện client cho Andromeda Service được xây dựng dựa trên Document: https://github.com/oriS-casino/andromeda-bep20-payment-gateway-document


Example:

```go
package main

import (
	"context"
	andromeda "github.com/dangnguyendota-casino/go-andromeda-client"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	client := andromeda.NewClient("http://localhost:9711", "example-password")

	// lấy số lượng ví đã tạo 
	client.UserCount(context.Background())
	// Lấy địa chỉ của user có ID 1
	client.GetUserAddress(context.Background(), 1)
	// Lấy user id của user address 0xBC74a1868e84A3238948dEEbc59E81693DA1702a
	client.GetUserID(context.Background(), common.HexToAddress("0xBC74a1868e84A3238948dEEbc59E81693DA1702a"))
	// Lấy các giao dịch nạp đã hoàn thành của user id 1 tại page 3 có page size = 2.
	client.GetUserCompletedDepositTransactions(context.Background(), 1, 2, 3)
	// lấy tổng số lượng giao dịch nạp của user id 1
	client.GetDepositTransactionsCount(context.Background(), 1)
	// lấy các giao dịch nạp đang pending của user id 1
	client.GetUserPendingDepositTransactions(context.Background(), 1)
	// Lấy các giao dịch rút đã hoàn thành của user id 1 tại page 3 có page size = 2.
	client.GetUserCompletedWithdrawTransactions(context.Background(), 1, 2, 3)
	// lấy tổng số lượng giao dịch rút của user id 1
	client.GetWithdrawTransactionsCount(context.Background(), 1)
	// lấy các giao dịch rút đang pending của user id 1
	client.GetUserPendingWithdrawTransactions(context.Background(), 1)
	// kiểm tra giao dịch 
	client.CheckDepositTx(context.Background(), andromeda.CheckDepositTransactionReq{
		UserID:          1,
		TransactionHash: "0xccf29a252e8bf3caaacddf47ffac2422ab0cccd2135972e51186d00f8339a1d5",
	})
	// yêu cầu chuyển tiền 
	client.Withdraw(context.Background(), andromeda.WithdrawReq {
		UserID    : 1,
		Address   : "0xBC74a1868e84A3238948dEEbc59E81693DA1702a",
		TokenName : "usdt",
		Amount    : 0.145,
		RequestID : "123",
		Checksum  : "abcxyz",
	})
	
}
```
