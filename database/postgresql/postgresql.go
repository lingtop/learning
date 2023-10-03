package postgresql

import (
	"context"
	"fmt"
	"os"

	"template/service/logger"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/zapadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
)

var DatabasePool *pgxpool.Pool

func getDbLogLevel(level string) pgx.LogLevel {
	switch level {
	case "info":
		return pgx.LogLevelInfo
	case "warn":
		return pgx.LogLevelWarn
	case "debug":
		return pgx.LogLevelDebug
	case "error":
		return pgx.LogLevelError
	case "trace":
		return pgx.LogLevelTrace
	case "none":
		return pgx.LogLevelNone
	default:
		return pgx.LogLevelInfo
	}
}

func InitDatabase() {
	// Init Database Config
	userName := viper.GetString("Database.Username")
	password := viper.GetString("Database.Password")
	host := viper.GetString("Database.Host")
	port := viper.GetInt("Database.Port")
	databaseName := viper.GetString("Database.DatabaseName")
	databaseSchema := viper.GetString("Database.DatabaseSchema")
	connectionTimeout := viper.GetInt("Database.ConnectionTimeout")

	databaseUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?search_path=%s&connect_timeout=%d",
		userName, password, host, port, databaseName, databaseSchema, connectionTimeout,
	)

	databaseConfig, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		logger.Logger.Errorf("Unable to parse config for database: %s\n", err)
		os.Exit(1)
	}

	databaseConfig.MaxConns = viper.GetInt32("Database.MaxConnection")
	databaseConfig.MinConns = viper.GetInt32("Database.MinConnection")

	databaseConfig.ConnConfig.LogLevel = getDbLogLevel(viper.GetString("Database.LogLevel"))
	databaseConfig.ConnConfig.Logger = zapadapter.NewLogger(logger.Logger.Desugar())

	// Create PGX Pool
	logger.Logger.Infof("Database pool is starting")
	DatabasePool, err = pgxpool.ConnectConfig(context.Background(), databaseConfig)

	if err != nil {
		logger.Logger.Errorf("Unable to connect to database: %s\n", err)
		os.Exit(1)
	}
}

func ShutdownDatabase() {
	logger.Logger.Infof("Database pool is shutting down")
	if DatabasePool != nil {
		DatabasePool.Close()
	}
}
