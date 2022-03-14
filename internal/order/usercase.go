package order

import "context"

type UCInterface interface {
	CaseCreateKtbOrder(ctx context.Context) error
}
