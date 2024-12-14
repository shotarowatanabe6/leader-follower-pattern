package handler

import "time"

type Transaction struct {
	ID        string    // トランザクションID
	Amount    float64   // 金額
	Timestamp time.Time // タイムスタンプ
	AccountID string    // アカウントID
}
