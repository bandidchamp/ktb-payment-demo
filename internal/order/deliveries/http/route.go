package http

import (
	"ktb-payment/internal/order"

	"github.com/gofiber/fiber/v2"
)

func MapOrderRoute(r fiber.Router, h order.Handler) {
	r.Post("/orders", h.CreateKtbOrder)
}
