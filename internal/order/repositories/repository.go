package repositories

import (
	"ktb-payment/internal/order"

	"github.com/jmoiron/sqlx"
)

type ProductRepo struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) order.RopoInterface {
	return &ProductRepo{db}
}

// func (conn *ProductRepo) GetTransectionById(ctx context.Context, transection_id string) error {
// 	span, ctx := apm.StartSpan(ctx, "ProductRepo.GetTransectionById", "usecase")
// 	defer span.End()
// 	query := "select * from ktb where 1"
// 	rows, err := conn.db.DB.QueryContext(ctx, query)
// 	if err != nil {
// 		return fmt.Errorf("failed to query rows with error: %w", err)
// 	}
// 	defer rows.Close()
// 	return nil
// }

// func (conn *ProductRepo) GetKtbPaymentById(ktbpayment_id string) error {

// 	return nil
// }
