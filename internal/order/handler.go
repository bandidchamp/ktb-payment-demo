package order

import "github.com/gofiber/fiber/v2"

type Handler interface {
	CreateKtbOrder(c *fiber.Ctx) error
}
