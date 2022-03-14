package http

import (
	"ktb-payment/internal/order"

	"github.com/gofiber/fiber/v2"
)

type orderHandler struct {
	orderUC order.UCInterface
}

func NewOrderHandler(orderUC order.UCInterface) order.Handler {
	return &orderHandler{
		orderUC: orderUC,
	}
}

func (oh *orderHandler) CreateKtbOrder(c *fiber.Ctx) error {
	// span, ctx := apm.StartSpan(c.Context(), "orderHandler.CreateKtbOrder", "handler")
	// defer span.End()
	return c.Status(200).JSON(fiber.Map{"msg": "hello"})
}
