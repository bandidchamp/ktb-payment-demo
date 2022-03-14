package server

import (
	"fmt"
	"ktb-payment/configs"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	fiber     *fiber.App
	ktb_fiber *fiber.App
	config    *configs.Config
}

func NewServer(config *configs.Config) *Server {
	return &Server{
		fiber:     fiber.New(),
		ktb_fiber: fiber.New(),
		config:    config,
	}
}

func (s *Server) StartServer() error {

	// // map handle here.
	s.MapPaymentHandlers(s.fiber)

	s.fiber.Use(
		cors.New(cors.Config{
			AllowOrigins:     "example.com, dev.info, bandid.info, localhost:3000",
			AllowMethods:     "GET,POST,PUT,PATCH,DELETE",
			AllowCredentials: true,
		}),
	)

	if err := s.startServerWithInteruptShutdown(s.fiber, s.config.Fiber.Host, s.config.Fiber.Port); err != nil {
		return fmt.Errorf("failed to start server with graceful shutdown with error: %w", err)
	}

	return nil

}

func (s *Server) startServerWithInteruptShutdown(f *fiber.App, Host string, Port string) error {
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) //* Catch OS signals.
		<-sigint

		//* Received an interrupt signal, shutdown.
		if err := f.Shutdown(); err != nil {
			return
		}

		close(idleConnsClosed)
	}()

	err := f.Listen(fmt.Sprintf(`%s:%s`, Host, Port))
	if err != nil {
		return fmt.Errorf("oops... server is not running! reason: %v", err)
	}
	<-idleConnsClosed
	return nil
}
