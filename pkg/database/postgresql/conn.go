package postgresql

import (
	"fmt"
	"ktb-payment/configs"

	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v4/stdlib" //* load pgx driver for PostgreSQL
)

//* NewPostgreSQLDB func for connection to PostgreSQL database.
func NewPostgreSQLDB(c *configs.Config) (*sqlx.DB, error) {
	dataSoruce := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Posrgress.Host,
		c.Posrgress.Port,
		c.Posrgress.Username,
		c.Posrgress.Password,
		c.Posrgress.DBname,
		"disable",
	)

	//* Define database connection for PostgreSQL.
	db, err := sqlx.Connect("pgx", dataSoruce)
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	//* Try to ping database.
	if err := db.Ping(); err != nil {
		defer db.Close()
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	return db, nil
}
