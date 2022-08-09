package andromeda

type DepositPacket struct {
	UserID       int64   `json:"user_id"`
	TxHash       string  `json:"tx_hash"`
	Tokens       float64 `json:"tokens"`
	TokenName    string  `json:"token_name"`
	TokenAddress string  `json:"token_address"`
	Decimals     int64   `json:"decimals"`
	ChainID      int64   `json:"chain_id"`
	Checksum     string  `json:"checksum"`
}

type TransferredTokens struct {
	RequestID string  `json:"request_id"`
	UserID    int64   `json:"user_id"`
	TxHash    string  `json:"tx_hash"`
	Tokens    float64 `json:"tokens"`
	TokenName string  `json:"token_name"`
	Checksum  string  `json:"checksum"`
}

type DepositTx struct {
	ID               string  `json:"id,omitempty"`
	TxHash           string  `json:"tx_hash,omitempty"`
	UserID           int64   `json:"user_id,omitempty"`
	ContractAddress  string  `json:"contract_address,omitempty"`
	FromAddress      string  `json:"from_address,omitempty"`
	ToAddress        string  `json:"to_address,omitempty"`
	BlockNumber      uint64  `json:"block_number,omitempty"`
	BlockHash        string  `json:"block_hash,omitempty"`
	Confirmations    int64   `json:"confirmations"`
	MaxConfirmations int64   `json:"max_confirmations"`
	Value            string  `json:"value,omitempty"`
	TokenName        string  `json:"token_name,omitempty"`
	TokenAddress     string  `json:"token_address,omitempty"`
	Tokens           float64 `json:"tokens,omitempty"`
	ChainID          int64   `json:"chain_id,omitempty"`
	Decimals         int64   `json:"decimals,omitempty"`
	IsApproved       bool    `json:"is_approved,omitempty"`
	IsCollected      bool    `json:"is_collected,omitempty"`
	UpdatedAt        int64   `json:"updated_at,omitempty"`
	CreatedAt        int64   `json:"created_at,omitempty"`
}

type WithdrawTx struct {
	ID            string  `json:"id,omitempty"`
	RequestID     string  `json:"request_id,omitempty"`
	UserID        int64   `json:"user_id,omitempty"`
	Tokens        float64 `json:"tokens,omitempty"`
	TokenName     string  `json:"token_name,omitempty"`
	TokenAddress  string  `json:"token_address,omitempty"`
	ToAddress     string  `json:"to_address,omitempty"`
	IsTransferred bool    `json:"is_transferred,omitempty"`
	TxHash        string  `json:"tx_hash,omitempty"`
	CreatedAt     int64   `json:"created_at,omitempty"`
	UpdatedAt     int64   `json:"updated_at,omitempty"`
}
