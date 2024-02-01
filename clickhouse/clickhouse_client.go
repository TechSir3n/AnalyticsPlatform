package clickhouse

import (
	"os"

	"github.com/ClickHouse/clickhouse-go/v2"
	log "github.com/TechSir3n/analytics-platform/logging"
	_ "github.com/TechSir3n/analytics-platform/assistance"
	"github.com/jmoiron/sqlx"
)

type ClickHouseClient struct {
	db *sqlx.DB
}


func NewClickHouseClient() (*ClickHouseClient, error) {
	db, err := sqlx.Open("clickhouse", os.Getenv("CLICKHOUSE_URL"))
	if err != nil {
		return &ClickHouseClient{}, err
	}

	if err := db.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			log.Log.Printf("[%d] %s \n %s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			return &ClickHouseClient{}, err
		}
	}

	return &ClickHouseClient{db: db}, nil
}

func (c *ClickHouseClient) InsertData(data []string) error {

	return nil
}

func (c *ClickHouseClient) UpdateData() error {
	return nil
}

func (c *ClickHouseClient) GetData() error {
	return nil
}

func (c *ClickHouseClient) DeleteData() error {
	return nil
}
