package clickhouse

import (
	"database/sql"
	"errors"
	"github.com/ClickHouse/clickhouse-go/v2"
	_ "github.com/TechSir3n/analytics-platform/assistance"
	"os"
	"time"
)

type Transaction struct {
	ID     string
	Name   string
	Type   string
	Amount float64
	Date   time.Time
}

type Database interface {
	InsertData(id, name, tType string, amount, price, revenue, benefit float64, date string) error
	UpdateData() error
	GetData(id int64) error
	DeleteData(id int64) error
}

func NewClickHouseClient(databaseType string) (Database, error) {
	conn := clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{os.Getenv("CLICKHOUSE_ADDR")},
		Auth: clickhouse.Auth{
			Database: os.Getenv("CLICKHOUSE_DB"),
			Username: os.Getenv("CLICKHOUSE_USERNAME"),
		},

		DialTimeout: time.Second * 30,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},

		Debug:                true,
		BlockBufferSize:      10,
		MaxCompressionBuffer: 10240,

		ClientInfo: clickhouse.ClientInfo{
			Products: []struct {
				Name    string
				Version string
			}{
				{Name: "Analytics-Platform", Version: "0.1"},
			},
		},
	})

	conn.SetMaxIdleConns(5)
	conn.SetMaxOpenConns(10)
	conn.SetConnMaxLifetime(time.Hour)

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	var db Database

	switch databaseType {
	case "AnalyticsPlatform":
		db = &AnalyticsPlatformDB{db: conn}
	case "Transaction":
		db = &TransactionDB{db: conn}
	case "Product":
		db = &ProductDB{db: conn}
	default:
		return nil, errors.New("Invalid database type")
	}

	if err := createTables(conn, databaseType); err != nil {
		return nil, err
	}

	return db, nil
}

func createTables(conn *sql.DB, databaseType string) error {
	var query string

	switch databaseType {
	case "AnalyticsPlatform":
		query = `CREATE TABLE IF NOT EXISTS default.AnalyticsPlatform (
            id UInt64,
            revenue Float64,
            benefit Float64,
            date Date
            ) ENGINE = MergeTree()
            ORDER BY id`
	case "Transaction":
		query = `CREATE TABLE IF NOT EXISTS default.Transaction (
            order_id UInt64,
            customer_name String,
            transaction_type String,
            amount Float64,
            transaction_date Date
            ) ENGINE = SummingMergeTree()
            ORDER BY transaction_date`
	case "Product":
		query = `CREATE TABLE IF NOT EXISTS default.Product (
            product_id UInt64,
            product_name String,
            price Float64,
            quantity UInt64
            ) ENGINE = MergeTree()
            ORDER BY product_id`
	default:
		return errors.New("Invalid database type")
	}
	_, err := conn.Exec(query)
	return err
}
