package model

// GetBlockHeaderVerboseResult models the data from the getblockheader command when
// the verbose flag is set.  When the verbose flag is not set, getblockheader
// returns a hex-encoded string.
type GetBlockHeaderVerboseResult struct {
	Hash          string  `json:"hash"`
	Confirmations int64   `json:"confirmations"`
	Height        int32   `json:"height"`
	Version       int32   `json:"version"`
	VersionHex    string  `json:"versionHex"`
	MerkleRoot    string  `json:"merkleroot"`
	Time          int64   `json:"time"`
	Nonce         uint64  `json:"nonce"`
	Bits          string  `json:"bits"`
	Difficulty    float64 `json:"difficulty"`
	PreviousHash  string  `json:"previousblockhash,omitempty"`
	NextHash      string  `json:"nextblockhash,omitempty"`
}

// GetBlockVerboseResult models the data from the getblock command when the
// verbose flag is set.  When the verbose flag is not set, getblock returns a
// hex-encoded string.
type GetBlockVerboseResult struct {
	Hash          string `json:"hash"`
	Confirmations int64  `json:"confirmations"`
	StrippedSize  int32  `json:"strippedsize"`
	Size          int32  `json:"size"`
	Weight        int32  `json:"weight"`
	Height        int64  `json:"height"`
	Version       int32  `json:"version"`
	VersionHex    string `json:"versionHex"`
	MerkleRoot    string `json:"merkleroot"`
	// Tx            []string      `json:"tx,omitempty"`
	RawTx      []TxRawResult `json:"rawtx,omitempty"`
	Time       int64         `json:"time"`
	Nonce      uint32        `json:"nonce"`
	Bits       string        `json:"bits"`
	Difficulty float64       `json:"difficulty"`
	// PreviousHash  string        `json:"previousblockhash"`
	// NextHash      string        `json:"nextblockhash,omitempty"`

	Mediantime        uint64 `json:"mediantime,omitempty"` //(numeric) The median block time in seconds since epoch (Jan 1 1970 GMT)
	Chainwork         string `json:"chainwork,omitempty"`
	NTx               int    `json:"ntx,omitempty"` // (numeric) The number of transactions in the block.
	Previousblockhash string `json:"previousblockhash,omitempty"`
	Nextblockhash     string `json:"nextblockhash,omitempty"`
}

type GetBlockResultV1 struct {
	GetBlockVerboseResult
	Tx []string `json:"tx,omitempty"`
}
type GetBlockResultV2 struct {
	GetBlockVerboseResult
	Tx []TxRawResult `json:"tx,omitempty"`
}