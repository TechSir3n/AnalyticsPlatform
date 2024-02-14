package clickhouse

import (
	"database/sql"
)

type AnalyticsPlatformDB struct {
	db *sql.DB
}

func (c *AnalyticsPlatformDB) InsertData(id, name, tType string, amount, price, revenue, benefit float64, date string) error {
	_, err := c.db.Exec("INSERT INTO AnalyticsPlatform(id,revenue,benefit,date)", id, revenue, benefit, date)
	if err != nil {
		return err
	}
	return nil
}

func (c *AnalyticsPlatformDB) UpdateData() error {
	return nil
}

func (c *AnalyticsPlatformDB) GetData(id int64) error {
	return nil
}

func (c *AnalyticsPlatformDB) DeleteData(id int64) error {
	return nil
}
