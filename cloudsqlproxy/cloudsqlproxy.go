package cloudsqlproxy

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

type ConnectionPoolConfiguration struct {
	// Database name
	DatabaseName string
	// Database username
	DatabaseUsername string
	// Database password
	DatabasePassword string
	// TCP hostname. If set, the connection pool will use TCP. Otherwise, it
	// will use Unix sockets.
	TcpHost string
	// Instance connection name. If TcpHost is set, this value will be ignored.
	InstanceConnectionName string
	// Maximum number of connections in idle connection pool
	MaxIdleConns int
	// Maximum number of open connections to the database
	MaxOpenConns int
	// Maximum time (in seconds) that a connection can remain open
	ConnMaxLifetime time.Duration
}

func defaultConnectionPoolConfiguration() *ConnectionPoolConfiguration {
	return &ConnectionPoolConfiguration{
		DatabaseName:           os.Getenv("DB_NAME"),
		DatabaseUsername:       os.Getenv("DB_USER"),
		DatabasePassword:       os.Getenv("DB_PASS"),
		TcpHost:                os.Getenv("DB_TCP_HOST"),
		InstanceConnectionName: os.Getenv("INSTANCE_CONNECION_NAME"),
		MaxIdleConns:           5,
		MaxOpenConns:           7,
		ConnMaxLifetime:        1800,
	}
}

// NewConnectionPool initializes a connection pool based on the provided
// configuration.
func NewConnectionPool(config *ConnectionPoolConfiguration) (*sql.DB, error) {
	if config == nil {
		config = defaultConnectionPoolConfiguration()
	}

	if config.TcpHost != "" {
		return initTcpConnectionPool(config)
	} else {
		return initSocketConnectionPool(config)
	}
}

func initSocketConnectionPool(config *ConnectionPoolConfiguration) (*sql.DB, error) {
	connectionString := fmt.Sprintf(
		"%s:%s@unix(/cloudsql/%s)/%s",
		config.DatabaseUsername,
		config.DatabasePassword,
		config.InstanceConnectionName,
		config.DatabaseName,
	)

	pool, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}

	configureConnectionPool(pool, config)

	return pool, nil
}

func initTcpConnectionPool(config *ConnectionPoolConfiguration) (*sql.DB, error) {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		config.DatabaseUsername,
		config.DatabasePassword,
		config.TcpHost,
		config.DatabaseName,
	)

	pool, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}

	configureConnectionPool(pool, config)

	return pool, nil
}

func configureConnectionPool(dbPool *sql.DB, config *ConnectionPoolConfiguration) {
	if config.MaxIdleConns != 0 {
		dbPool.SetMaxIdleConns(config.MaxIdleConns)
	}
	if config.MaxOpenConns != 0 {
		dbPool.SetMaxOpenConns(config.MaxOpenConns)
	}
	if config.ConnMaxLifetime != 0 {
		dbPool.SetConnMaxLifetime(config.ConnMaxLifetime)
	}
}
