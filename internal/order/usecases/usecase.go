package usecases

import (
	"context"
	"ktb-payment/internal/order"

	"go.elastic.co/apm"
)

type orderUC struct {
	orderRepo order.RopoInterface
}

func NewOrderUseCase(orderRepo order.RopoInterface) order.UCInterface {
	return &orderUC{
		orderRepo: orderRepo,
	}
}

func (o *orderUC) CaseCreateKtbOrder(ctx context.Context) error {
	span, ctx := apm.StartSpan(ctx, "orderUC.CaseCreateKtbOrder", "usecase")
	defer span.End()
	// o.orderRepo.GetTransectionById(ctx, "111DED4487S")

	return nil
}
