package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type Transaction struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
	Type   string  `json:"type"`
}

type TransactionRequest struct {
	RequestID    string        `json:"requestId"`
	Transactions []Transaction `json:"transactions"`
}

type TransactionResponse struct {
	Status      string  `json:"status"`
	TotalCredit float64 `json:"totalCredit"`
	Signature   string  `json:"signature"`
	ProcessedAt string  `json:"processedAt"`
}

func HandleRequest(ctx context.Context, req TransactionRequest) (TransactionResponse, error) {
	var totalCredit float64

	var idBuilder strings.Builder

	for _, tx := range req.Transactions {
		if tx.Type == "CREDIT" {
			totalCredit += tx.Amount
			idBuilder.WriteString(tx.ID)
		}
	}

	hash := sha256.Sum256([]byte(idBuilder.String()))

	signature := hex.EncodeToString(hash[:])

	return TransactionResponse{
		Status:      "success",
		TotalCredit: totalCredit,
		Signature:   signature,
		ProcessedAt: time.Now().UTC().Format(time.RFC3339),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
