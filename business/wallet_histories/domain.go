package wallet_histories

import (
	"context"
	"time"
)

type Domain struct {
	ID            int
	WalletId      int
	TransactionId int
	Type          string
	Nominal       float64
	CreatedAt     time.Time
}

type Usecase interface {
	Create(ctx context.Context, domain Domain) error
}

type Repository interface {
	Create(ctx context.Context, domain Domain) error
}
